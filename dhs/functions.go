package dhs

import (
	"errors"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func GetMethodFullName(pbMsg proto.Message) (string, error) {
	services := File_proto_dhs_dhs_proto.Services()
	for i := 0; i < services.Len(); i++ {
		s := services.Get(i)
		methods := s.Methods()
		for j := 0; j < methods.Len(); j++ {
			m := methods.Get(j)
			msgName := pbMsg.ProtoReflect().Descriptor().Name()
			if msgName == m.Input().Name() {
				return string(m.FullName()), nil
			}
			if msgName == m.Output().Name() {
				return string(m.FullName()), nil
			}
		}
	}
	return "", errors.New("could not find message name")
}

func Wrap(fullName string, pbMsg proto.Message) ([]byte, error) {
	msg, err := proto.Marshal(pbMsg)
	if err != nil {
		return nil, err
	}

	wrapped := &Wrapper{
		Name: fmt.Sprintf(".%s", fullName),
		Data: msg,
	}

	out, err := proto.Marshal(wrapped)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func Unwrap(msg []byte) ([]byte, error) {
	wrapped := &Wrapper{}
	err := proto.Unmarshal(msg, wrapped)
	if err != nil {
		return nil, err
	}

	return wrapped.Data, nil
}
