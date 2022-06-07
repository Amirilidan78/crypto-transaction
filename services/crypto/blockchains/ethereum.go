package blockchains

import (
	"crypto-transaction/config"
	"crypto-transaction/pkg/blockbook"
	"crypto-transaction/pkg/twallet"
	cryptoPb "crypto-transaction/pkg/twallet/proto"
	"encoding/hex"
	"github.com/golang/protobuf/proto"
)

const ETHBlockchainString = "ETH"

type Ethereum interface {
	SignTransaction(pb proto.Message) ([]byte, error)
	GetRawTransaction(res []byte) (string, error)
	BroadCastTransaction(hex string) (string, error)
}

type ethereum struct {
	c  config.Config
	tw twallet.TWallet
	bb blockbook.HttpBlockBook
}

func (e *ethereum) SignTransaction(pb proto.Message) ([]byte, error) {
	return e.tw.SignTransaction(ETHBlockchainString, pb)
}

func (e *ethereum) GetRawTransaction(res []byte) (string, error) {

	so := &cryptoPb.EthSigningOutput{}

	err := proto.Unmarshal(res, so)

	if err != nil {
		return "", err
	}

	return string("0x" + hex.EncodeToString(so.GetEncoded())), err
}

func (e *ethereum) BroadCastTransaction(hex string) (string, error) {

	resp, err := e.bb.BroadcastTransaction(ETHBlockchainString, hex)

	if err != nil {
		return "", err
	}

	return resp.TxId, err
}

func NewEthereumBlockchain(c config.Config, tw twallet.TWallet, bb blockbook.HttpBlockBook) Ethereum {
	return &ethereum{c, tw, bb}
}
