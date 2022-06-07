package common

import "github.com/golang/protobuf/proto"

type Coin interface {
	CreateTransaction(amount string, fromAddress string, toAddress string, addressPrivateKey string) (proto.Message, error)
}
