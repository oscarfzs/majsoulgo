package majsoulgo

import "google.golang.org/protobuf/proto"

type MajsoulClient interface {
	Call(fullName string, pbMsg proto.Message) error
}
