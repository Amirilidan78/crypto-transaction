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

const LTCBlockchainString = "BTC"

type LiteCoin interface {
	SignTransaction(pb proto.Message) ([]byte, error)
	GetRawTransaction(res []byte) (string, error)
	BroadCastTransaction(hex string) (string, error)
}

type liteCoin struct {
	c  config.Config
	tw twallet.TWallet
	bb blockbook.HttpBlockBook
}

func (b *liteCoin) SignTransaction(pb proto.Message) ([]byte, error) {
	return b.tw.SignTransaction(LTCBlockchainString, pb)
}

func (b *liteCoin) GetRawTransaction(res []byte) (string, error) {

	so := &cryptoPb.BtcSigningOutput{}

	err := proto.Unmarshal(res, so)

	if err != nil {
		return "", err
	}

	if so.GetError() != cryptoPb.SigningError_OK {
		return "", errors.New("error in signing " + LTCBlockchainString + " transaction. error: " + so.GetError().String())
	}

	return hex.EncodeToString(so.GetEncoded()), nil
}

func (b *liteCoin) BroadCastTransaction(hex string) (string, error) {

	resp, err := b.bb.BroadcastTransaction(LTCBlockchainString, hex)

	if err != nil {
		return "", err
	}

	return resp.TxId, err
}

func NewLiteCoinBlockchain(c config.Config, tw twallet.TWallet, bb blockbook.HttpBlockBook) LiteCoin {
	return &liteCoin{c, tw, bb}
}
