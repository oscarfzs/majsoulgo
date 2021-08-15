package dhs

import (
	"majsoulgo"

	"google.golang.org/protobuf/proto"
)

type ContestManagerClient struct{ majsoulgo.MajsoulChannel }

func NewContestManagerClient() *ContestManagerClient {
	c := new(ContestManagerClient)
	c.MajsoulChannel = *majsoulgo.NewMajsoulChannel()
	return c
}

func (client *ContestManagerClient) Call(pbReq proto.Message) ([]byte, error) {
	name, err := GetMethodFullName(pbReq)
	if err != nil {
		return nil, err
	}

	wrapped, err := Wrap(name, pbReq)
	if err != nil {
		return nil, err
	}
	res := client.Send(wrapped)

	unwrapped, err := Unwrap(res)
	if err != nil {
		return nil, err
	}

	return unwrapped, nil
}
