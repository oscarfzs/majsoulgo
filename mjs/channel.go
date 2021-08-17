package mjs

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	MSG_TYPE_NOTIFY   = 1
	MSG_TYPE_REQUEST  = 2
	MSG_TYPE_RESPONSE = 3

	PING_INTERVAL    = 3 //seconds
	TIMEOUT_INTERVAL = 2 //seconds

	OUTGOING_CHANNEL_SIZE      = 20
	NOTIFICATIONS_CHANNEL_SIZE = 100
)

type MajsoulChannel struct {
	//Underlying websocket connection
	Connection *websocket.Conn

	//Mutex for coarse grained locking
	mutexAll *sync.Mutex

	//Mutex for ensuring that only one goroutine can
	//read and increment the msgIndex field
	mutexMsgIndex *sync.Mutex

	//Requests sent to the server must be indexed. When a request is sent
	//to the server, the second and third byte of the message contains the message index.
	//That way, responses received from the server can be traced back to the
	//request that initiated the response.
	msgIndex      uint16
	requests      map[uint16](chan []byte)
	responses     map[uint16](chan []byte)
	notifications chan []byte

	//Buffered channel for outgoing requests
	outgoing chan struct {
		MsgType int
		Msg     []byte
	}

	timeLastPingSent     time.Time
	timeLastPongReceived time.Time

	//Buffered channel for signaling that the Majsoul channel has stopped
	stop chan struct{}

	exitValue error
	open      bool
}

//Function for creating a new MajsoulChannel. Handles all field initalizations.
func NewMajsoulChannel() *MajsoulChannel {
	m := new(MajsoulChannel)

	m.mutexAll = new(sync.Mutex)
	m.mutexMsgIndex = new(sync.Mutex)
	m.responses = make(map[uint16](chan []byte))
	m.notifications = make(chan []byte, NOTIFICATIONS_CHANNEL_SIZE)

	m.outgoing = make(chan struct {
		MsgType int
		Msg     []byte
	}, OUTGOING_CHANNEL_SIZE)

	m.stop = make(chan struct{}, 1)
	return m
}

/*
Connects to the server specified by the url string. The url string of a
MahjongSoul server can be obtained by using the functions defined in majsoul.go

After successfully connecting, the channel can be closed by calling MajsoulChannel.Close()
*/
func (m *MajsoulChannel) Connect(url string) error {
	var err error
	m.Connection, _, err = websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		return err
	}

	defer m.Connection.Close()

	log.Println("Successfully connected to", url)

	m.open = true
	m.stop = make(chan struct{}, 1)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	m.Connection.SetPongHandler(m.pongHandler)

	go m.recvMsg(m.stop)
	go m.sendMsg(m.stop)
	go m.sustain(m.stop)

	for {
		select {
		case <-interrupt:
			m.Close(errors.New("connection terminated by interrupt signal"))
			return m.ExitValue()
		case <-m.stop:
			return m.ExitValue()
		}
	}
}

//Processes and pushes a message to the outgoing channel for sendMsg() to handle
func (m *MajsoulChannel) Send(msg []byte) []byte {
	m.mutexMsgIndex.Lock()
	index := m.msgIndex
	m.msgIndex += 1
	m.mutexMsgIndex.Unlock()

	m.responses[index] = make(chan []byte, 1)

	h1 := make([]byte, 1)
	h2 := make([]byte, 2)

	h1[0] = byte(MSG_TYPE_REQUEST)
	binary.LittleEndian.PutUint16(h2, index)

	var msgBuffer bytes.Buffer
	msgBuffer.Write(h1)
	msgBuffer.Write(h2)
	msgBuffer.Write(msg)

	m.writeMsg(websocket.BinaryMessage, msgBuffer.Bytes())

	//Wait for the response to come through the channel
	res := <-m.responses[index]
	return res
}

/*
Closes the channel. The way channels are closed is by passing an error
into the close channel. A nil value passed into the channel indicates that the
channel was closed by the application and not due to encountering any errors.
*/
func (m *MajsoulChannel) Close(err error) {
	m.mutexAll.Lock()
	if m.open {
		m.exitValue = err
		m.open = false
		close(m.stop)
		log.Println("Majsoul channel closed.")
	}
	m.mutexAll.Unlock()
}

func (m *MajsoulChannel) IsOpen() bool {
	return m.open
}

/*
ExitValue is the function for user applications to retrieve the error that
caused the MajsoulChannel to close. A nil value means that the channel was closed
by the user.
*/
func (m *MajsoulChannel) ExitValue() error {
	<-m.stop
	return m.exitValue
}

/*
Looping function that sends ping messages to the server every PING_INTERVAL seconds.
A timeout is implemented such that if the corresponding pong message is
not received within TIMEOUT_INTERVAL seconds, then the MajsoulChannel is automatically
closed.
*/
func (m *MajsoulChannel) sustain(stop chan struct{}) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ticker := time.NewTicker(time.Duration(PING_INTERVAL) * time.Second)
	defer ticker.Stop()

	var timer (<-chan time.Time)
	for {
		select {
		case <-interrupt:
			return
		case <-stop:
			return
		case <-ticker.C:
			m.writeMsg(websocket.PingMessage, nil)

			timer = time.After(time.Duration(TIMEOUT_INTERVAL) * time.Second)
		case <-timer:
			if time.Since(m.timeLastPongReceived).Seconds() > float64(TIMEOUT_INTERVAL) {
				m.Close(errors.New("timeout: pong not received within timeout duration"))
				return
			}
		}
	}
}

func (m *MajsoulChannel) writeMsg(msgType int, msg []byte) {
	m.outgoing <- struct {
		MsgType int
		Msg     []byte
	}{msgType, msg}
}

//Looping goroutine that reads outgoing messages from the outgoing channel
//and sends them through the websocket connection. Only this function may call
//the ReadMessage() function of the websocket connection.
func (m *MajsoulChannel) sendMsg(stop chan struct{}) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case <-interrupt:
			return
		case <-stop:
			return
		case out := <-m.outgoing:
			err := m.Connection.WriteMessage(out.MsgType, out.Msg)
			if err != nil {
				m.Close(err)
				return
			}
			switch out.MsgType {
			case websocket.PingMessage:
				m.timeLastPingSent = time.Now()
				log.Println("Ping sent")
			case websocket.PongMessage:
				log.Println("Pong sent")
			case websocket.CloseMessage:
				log.Println("Close sent")
			default:
				log.Println("Message sent: ", out.Msg)
			}
		}
	}
}

//Looping goroutine that reads incoming messages from the websocket connection.
//Only this function may call the WriteMessage() function of the websocket connection.
func (m *MajsoulChannel) recvMsg(stop chan struct{}) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			return
		case <-stop:
			return
		default:
			_, msg, err := m.Connection.ReadMessage()
			if err != nil {
				m.Close(err)
				return
			}

			m.processMsg(msg)
		}
	}
}

//Processes incoming messages to see if they are responses or notifications.
//Pushes the message to the corresponding channel (responses or notifications).
func (m *MajsoulChannel) processMsg(msg []byte) {
	reqType := int(msg[0])

	switch reqType {
	case MSG_TYPE_RESPONSE:
		var index uint16
		buf := bytes.NewReader(msg[1:3])
		err := binary.Read(buf, binary.LittleEndian, &index)
		if err != nil {
			log.Println(err)
			return
		}

		_, ok := m.responses[index]
		if !ok {
			log.Println(errors.New("response received with unexpected request index"))
			return
		}

		log.Println("Response received: ", msg)
		m.responses[index] <- msg[3:]
	case MSG_TYPE_NOTIFY:
		log.Println("Notification received: ", msg)
		if len(m.notifications) == cap(m.notifications) {
			<-m.notifications
		}
		m.notifications <- msg[1:]
	default:
		log.Println("Unknown message type received")
	}
}

func (m *MajsoulChannel) pongHandler(appData string) error {
	m.timeLastPongReceived = time.Now()
	log.Println("Pong received")
	return nil
}
