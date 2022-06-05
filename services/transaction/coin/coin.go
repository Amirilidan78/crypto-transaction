package coin

import "github.com/golang/protobuf/proto"

type Coin interface {
	CreateTransaction(amount string, fromAddress string, toAddress string, addressPrivateKey string) (proto.Message, error)
	SignTransaction(pb proto.Message) ([]byte, error)
	GetRawTransaction(res []byte) (string, error)
	BroadCastTransaction(hex string) (string, error)
}
