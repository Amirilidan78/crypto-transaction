package blockchains

import (
	"crypto-transaction/config"
	"crypto-transaction/pkg/blockbook"
	"crypto-transaction/pkg/twallet"
	cryptoPb "crypto-transaction/pkg/twallet/proto"
	"encoding/hex"
	"errors"
	"github.com/golang/protobuf/proto"
)

const BTCBlockchainString = "BTC"

type Bitcoin interface {
	SignTransaction(pb proto.Message) ([]byte, error)
	GetRawTransaction(res []byte) (string, error)
	BroadCastTransaction(hex string) (string, error)
}

type bitcoin struct {
	c  config.Config
	tw twallet.TWallet
	bb blockbook.HttpBlockBook
}

func (b *bitcoin) SignTransaction(pb proto.Message) ([]byte, error) {
	return b.tw.SignTransaction(BTCBlockchainString, pb)
}

func (b *bitcoin) GetRawTransaction(res []byte) (string, error) {

	so := &cryptoPb.BtcSigningOutput{}

	err := proto.Unmarshal(res, so)

	if err != nil {
		return "", err
	}

	if so.GetError() != cryptoPb.SigningError_OK {
		return "", errors.New("error in signing " + BTCBlockchainString + " transaction. error: " + so.GetError().String())
	}

	return hex.EncodeToString(so.GetEncoded()), nil
}

func (b *bitcoin) BroadCastTransaction(hex string) (string, error) {

	resp, err := b.bb.BroadcastTransaction(BTCBlockchainString, hex)

	if err != nil {
		return "", err
	}

	return resp.TxId, err
}

func NewBitcoinBlockchain(c config.Config, tw twallet.TWallet, bb blockbook.HttpBlockBook) Bitcoin {
	return &bitcoin{c, tw, bb}
}
