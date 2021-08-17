package dhs

import (
	"fmt"
	"majsoulgo/mjs"
	"majsoulgo/mjsproto"

	"google.golang.org/protobuf/proto"
)

type ContestManagerClient struct{ mjs.MajsoulChannel }

func NewContestManagerClient() *ContestManagerClient {
	c := new(ContestManagerClient)
	c.MajsoulChannel = *mjs.NewMajsoulChannel()
	return c
}

func (client *ContestManagerClient) CallMethod(methodFullName string, pbReq proto.Message) (proto.Message, error) {
	msg, err := proto.Marshal(pbReq)
	if err != nil {
		return nil, err
	}

	wrapped, err := Wrap(methodFullName, msg)
	if err != nil {
		return nil, err
	}

	res := client.Send(wrapped)

	_, data, err := Unwrap(res)
	if err != nil {
		return nil, err
	}

	outputName, err := mjsproto.FindMethodOutputName(File_proto_dhs_dhs_proto, methodFullName)
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

func Wrap(name string, msg []byte) ([]byte, error) {
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

func Unwrap(msg []byte) (string, []byte, error) {
	wrapped := &Wrapper{}
	err := proto.Unmarshal(msg, wrapped)
	if err != nil {
		return "", nil, err
	}

	return wrapped.Name, wrapped.Data, nil
}
