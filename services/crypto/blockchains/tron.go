package blockchains

import (
	"crypto-transaction/config"
	"crypto-transaction/pkg/trongrid"
	"crypto-transaction/pkg/twallet"
	cryptoPb "crypto-transaction/pkg/twallet/proto"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
)

const TRXBlockchainString = "TRX"

type Tron interface {
	SignTransaction(pb proto.Message) ([]byte, error)
	GetRawTransaction(res []byte) (string, error)
	BroadCastTransaction(hex string) (string, error)
}

type tron struct {
	c  config.Config
	tw twallet.TWallet
	tg trongrid.TronGrid
}

func (t *tron) SignTransaction(pb proto.Message) ([]byte, error) {
	return t.tw.SignTransaction(TRXBlockchainString, pb)
}

func (t *tron) GetRawTransaction(res []byte) (string, error) {

	so := &cryptoPb.TronSigningOutput{}

	err := proto.Unmarshal(res, so)

	if err != nil {
		return "", err
	}

	return so.GetJson(), nil
}

func (t *tron) BroadCastTransaction(hex string) (string, error) {

	resp, err := t.tg.BroadcastTransaction(hex)

	if err != nil {
		return "", nil
	}

	if result, exist := resp["result"]; exist {
		if result.(bool) {
			return resp["txid"].(string), nil
		}
	}

	return "", errors.New(fmt.Sprintf("transaction broadcast error: %s", resp["code"].(string)))
}

func NewTronBlockchain(c config.Config, tw twallet.TWallet, tg trongrid.TronGrid) Tron {
	return &tron{c, tw, tg}
}
