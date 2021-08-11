package majsoulgo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	MSG_TYPE_NOTIFY   = 1
	MSG_TYPE_REQUEST  = 2
	MSG_TYPE_RESPONSE = 3

	MAX_MSG_INDEX = 1 << 16

	PING_INTERVAL    = 5 //seconds
	TIMEOUT_INTERVAL = 3 //seconds
)

type ChannelMethods interface {
	Connect(url string)
	KeepAlive(ping int, timeout int)
}

type MajsoulChannel struct {
	Connection   *websocket.Conn
	Url          string
	LastResponse time.Time

	//Mutexes
	mutexMsgIndex sync.Mutex
	mutexMsgWrite sync.Mutex

	msgIndex int16

	responses map[int16](chan []byte)
}

func (ch *MajsoulChannel) Send(msg []byte) {
	ch.mutexMsgIndex.Lock()

	index := ch.msgIndex
	ch.msgIndex = (ch.msgIndex + int16(1)) % int16(MAX_MSG_INDEX)

	ch.mutexMsgIndex.Unlock()

	ch.responses[index] = make(chan []byte)

	h1 := make([]byte, 1)
	h2 := make([]byte, 2)

	binary.LittleEndian.PutUint16(h1, uint16(MSG_TYPE_REQUEST))
	binary.LittleEndian.PutUint16(h2, uint16(index))

	var msgBuffer bytes.Buffer
	msgBuffer.Write(h1)
	msgBuffer.Write(h2)
	msgBuffer.Write(msg)

	ch.mutexMsgWrite.Lock()

	err := ch.Connection.WriteMessage(websocket.TextMessage, msgBuffer.Bytes())

	ch.mutexMsgWrite.Unlock()

	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func (ch *MajsoulChannel) Connect(url string) {
	var err error
	ch.Connection, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Successfully connected to ", url)

	ch.responses = make(map[int16](chan []byte))

	go ch.KeepAlive(PING_INTERVAL)
	go ch.Listen()

}

func (ch *MajsoulChannel) KeepAlive(ping int) {
	ticker := time.NewTicker(time.Duration(PING_INTERVAL) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("tick")
			ch.mutexMsgWrite.Lock()
			err := ch.Connection.WriteMessage(websocket.PingMessage, []byte("Ping"))
			ch.mutexMsgWrite.Unlock()
			if err != nil {
				return
			}
		}
	}
}

func (ch *MajsoulChannel) Listen() {
	for {
		msgType, m, _ := ch.Connection.ReadMessage()
		if msgType == websocket.TextMessage {
			log.Println("Received message: ", string(m))
		}
	}
}
