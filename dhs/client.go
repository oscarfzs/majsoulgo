package dhs

import (
	"fmt"
	"log"
	"strings"

	"github.com/oscarfzs/majsoulgo/mjs"
	"github.com/oscarfzs/majsoulgo/mjsproto"

	"google.golang.org/protobuf/proto"
)

var (
	NOTIFICATION_HANDLER_CAPACITY = 10
)

type ContestManagerClient struct {
	mjs.MajsoulChannel

	notificationHandlers map[string][]func(pbMsg proto.Message)
}

func NewContestManagerClient() *ContestManagerClient {
	client := new(ContestManagerClient)
	client.MajsoulChannel = *mjs.NewMajsoulChannel()
	client.notificationHandlers = make(map[string][]func(pbMsg proto.Message))
	client.SetNotificationHandler(client.processNotification)
	return client
}

func (client *ContestManagerClient) CallMethod(methodFullName string, pbReq proto.Message) (proto.Message, error) {
	msg, err := proto.Marshal(pbReq)
	if err != nil {
		return nil, err
	}

	wrapped, err := wrap(methodFullName, msg)
	if err != nil {
		return nil, err
	}

	res := client.Send(wrapped)

	_, data, err := unwrap(res)
	if err != nil {
		return nil, err
	}

	outputName, err := mjsproto.FindMethodOutputName(File_dhs_proto, methodFullName)
	if err != nil {
		return nil, err
	}

	resMsg := NewMessageByName[outputName]()

	err = proto.Unmarshal(data, resMsg)
	if err != nil {
		return nil, err
	}

	return resMsg, nil
}

func (client *ContestManagerClient) AddNotificationHandler(pbMsg proto.Message, h func(pbMsg proto.Message)) {
	name := string(pbMsg.ProtoReflect().Descriptor().Name())
	handlers, ok := client.notificationHandlers[name]
	if !ok {
		handlers = []func(pbMsg proto.Message){}
	}
	client.notificationHandlers[name] = append(handlers, h)
}

func (client *ContestManagerClient) processNotification(msg []byte) {
	name, data, err := unwrap(msg)
	if err != nil {
		log.Println(err)
		return
	}

	_, ok := client.notificationHandlers[name]
	if !ok {
		log.Println("processNotification: no handler found")
		return
	}

	f, ok := NewMessageByName[name]
	if !ok {
		log.Println("processNotification: unable to find message " + name)
		return
	}

	pbMsg := f()
	err = proto.Unmarshal(data, pbMsg)
	if err != nil {
		log.Println("processNotification: unable to unmarshal message")
		return
	}

	for _, h := range client.notificationHandlers[name] {
		go h(pbMsg)
	}
}

func wrap(name string, msg []byte) ([]byte, error) {
	wrapped := &Wrapper{
		Name: fmt.Sprintf(".%s", name),
		Data: msg,
	}

	out, err := proto.Marshal(wrapped)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func unwrap(msg []byte) (string, []byte, error) {
	wrapped := &Wrapper{}
	err := proto.Unmarshal(msg, wrapped)
	if err != nil {
		return "", nil, err
	}

	split := strings.Split(wrapped.Name, ".")
	if len(split) == 0 {
		return "", wrapped.Data, nil
	}

	name := split[len(split)-1]
	return name, wrapped.Data, nil
}
