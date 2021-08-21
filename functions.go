package majsoulgo

import (
	"errors"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func GetMessageFullName(pbMsg proto.Message, fd protoreflect.FileDescriptor) (string, error) {
	services := fd.Services()
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

func FindMethodInputName(fd protoreflect.FileDescriptor, methodFullName string) (string, error) {
	split := strings.Split(methodFullName, ".")
	if len(split) < 3 {
		return "", errors.New("GetMethodInput: incomplete name")
	}

	serviceName := split[1]
	methodName := split[2]

	service := fd.Services().ByName(protoreflect.Name(split[1]))
	if service == nil {
		return "", errors.New("GetMethodInput: no service named " + serviceName)
	}

	method := service.Methods().ByName(protoreflect.Name(split[2]))
	if method == nil {
		return "", errors.New("GetMethodInput: no method name " + methodName)
	}
	return string(method.Input().Name()), nil
}

func FindMethodOutputName(fd protoreflect.FileDescriptor, methodFullName string) (string, error) {
	split := strings.Split(methodFullName, ".")
	if len(split) < 3 {
		return "", errors.New("GetMethodOutput: incomplete name")
	}

	serviceName := split[1]
	methodName := split[2]

	service := fd.Services().ByName(protoreflect.Name(split[1]))
	if service == nil {
		return "", errors.New("GetMethodOutput: no service named " + serviceName)
	}

	method := service.Methods().ByName(protoreflect.Name(split[2]))
	if method == nil {
		return "", errors.New("GetMethodOutput: no method name " + methodName)
	}
	return string(method.Output().Name()), nil
}
