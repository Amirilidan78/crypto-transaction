package common

import "github.com/golang/protobuf/proto"

type Blockchain interface {
	SignTransaction(pb proto.Message) ([]byte, error)
	GetRawTransaction(res []byte) (string, error)
	BroadCastTransaction(hex string) (string, error)
}
