package majsoulgo

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

	MAX_MSG_INDEX = 1 << 16

	PING_INTERVAL    = 3 //seconds
	TIMEOUT_INTERVAL = 2 //seconds
)

type MajsoulChannel struct {
	Connection *websocket.Conn

	mutexMsgIndex *sync.Mutex
	mutexMsgWrite *sync.Mutex
	mutexMsgRead  *sync.Mutex

	//Requests sent to the server must be indexed. When a request is sent
	//to the server, the second byte of the message must contain the msgIndex.
	//That way, responses received from the server can be traced back to the
	//request that initiated the response.
	msgIndex  int16
	responses map[int16](chan []byte)

	timeLastPingSent     time.Time
	timeLastPongReceived time.Time

	//`MajsoulChannel.stop` is a buffered channel of capacity 1 and does two tasks
	//First, it acts as a signaler to any goroutines that the channel
	//has been closed. Whenever an error value is sent through the channel, that
	//indicates the channel has been closed.
	//
	//Second, it provides a way to retrieve any errors that resulted in
	//the termination of the channel's connection.
	stop chan error

	open chan struct{}
}

func NewMajsoulChannel() *MajsoulChannel {
	m := new(MajsoulChannel)

	m.mutexMsgIndex = new(sync.Mutex)
	m.mutexMsgRead = new(sync.Mutex)
	m.mutexMsgWrite = new(sync.Mutex)

	m.responses = make(map[int16](chan []byte))
	m.open = make(chan struct{}, 1)
	m.stop = make(chan error, 1)
	return m
}

/*
UNFINISHED
*/
func (m *MajsoulChannel) Send(msg []byte) {
	//wait until the channel has successfully opened a connection
	s := <-m.open
	m.open <- s

	m.mutexMsgIndex.Lock()
	index := m.msgIndex
	m.msgIndex = int16(MAX_MSG_INDEX) % (index + 1)
	m.mutexMsgIndex.Unlock()

	m.responses[index] = make(chan []byte)

	h1 := make([]byte, 1)
	h2 := make([]byte, 2)

	h1[0] = byte(MSG_TYPE_REQUEST)
	binary.LittleEndian.PutUint16(h2, uint16(index))

	var msgBuffer bytes.Buffer
	msgBuffer.Write(h1)
	msgBuffer.Write(h2)
	msgBuffer.Write(msg)

	m.mutexMsgWrite.Lock()
	err := m.Connection.WriteMessage(websocket.TextMessage, msgBuffer.Bytes())
	log.Println("Sent message", msgBuffer.String())
	m.mutexMsgWrite.Unlock()

	if err != nil {
		log.Println("ERROR: ", err)
	}
}

/*
Connects to the server specified by the url string. The url string of a
MahjongSoul server can be obtained by using the functions defined in majsoul.go

Once connected, one goroutine will be spawned that listens for incoming messages
and one will be spawned that routinely sends ping messages in order to maintain the
connection.

After successfully connecting, the channel can be closed by calling MajsoulChannel.Close()
*/
func (m *MajsoulChannel) Connect(url string) {

	var err error
	m.Connection, _, err = websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		m.stop <- err
		return
	}

	defer m.Connection.Close()

	log.Println("Successfully connected to", url)
	m.open <- struct{}{}

	m.Connection.SetPongHandler(m.pongHandler)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go m.listen()
	go m.keepAlive()

	for {
		select {
		case <-interrupt:
			m.stop <- errors.New("connection terminated by interrupt signal")
			return
		case err := <-m.stop:
			m.stop <- err
			return
		}
	}
}

/*
Closes the channel. The way channels are closed is by passing an error
into the close channel. A nil value passed into the channel indicates that the
channel was closed by the application and not due to encountering any errors.
*/
func (m *MajsoulChannel) Close() {
	if !m.IsClosed() {
		m.stop <- nil
	}
}

func (m *MajsoulChannel) IsClosed() bool {
	//A channel with length 1 indicates a stored error and therefore
	//that the channel was closed
	return len(m.stop) == 1
}

/*
ExitValue is the function for user applications to retrieve the error that
caused the MajsoulChannel to close. A nil value means that the channel was closed
by the user.
*/
func (m *MajsoulChannel) ExitValue() error {
	err := <-m.stop
	m.stop <- err
	return err
}

/*
Looping function that sends ping messages to the server every PING_INTERVAL seconds.
A timeout is implemented such that if the corresponding pong message is
not received within TIMEOUT_INTERVAL seconds, then the MajsoulChannel is automatically
closed.
*/
func (m *MajsoulChannel) keepAlive() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ticker := time.NewTicker(time.Duration(PING_INTERVAL) * time.Second)
	defer ticker.Stop()

	var timer (<-chan time.Time)
	for {
		select {
		case <-interrupt:
			return
		case err := <-m.stop:
			m.stop <- err
			return
		case <-ticker.C:
			log.Println("Ping sent")
			m.mutexMsgWrite.Lock()
			err := m.Connection.WriteMessage(websocket.PingMessage, []byte("Ping"))
			m.mutexMsgWrite.Unlock()
			if err != nil {
				m.stop <- err
				return
			}
			m.timeLastPingSent = time.Now()
			timer = time.After(time.Duration(TIMEOUT_INTERVAL) * time.Second)
		case <-timer:
			if time.Since(m.timeLastPongReceived).Seconds() > float64(TIMEOUT_INTERVAL) {
				err := errors.New("timeout: pong not received within timeout duration")
				m.stop <- err
				return
			}
		}
	}
}

/*
UNFINISHED
*/
func (m *MajsoulChannel) listen() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			return
		case err := <-m.stop:
			m.stop <- err
			return
		default:
			m.mutexMsgRead.Lock()
			_, r, err := m.Connection.NextReader()
			m.mutexMsgRead.Unlock()
			if err != nil {
				if !m.IsClosed() {
					m.stop <- err
				}
				return
			}

			b := bytes.Buffer{}
			b.ReadFrom(r)
			log.Println("Received message: ", b.String())
		}
	}
}

func (m *MajsoulChannel) pongHandler(appData string) error {
	m.timeLastPongReceived = time.Now()
	log.Println("Pong received")
	return nil
}
