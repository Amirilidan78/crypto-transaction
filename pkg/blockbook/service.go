package blockbook

import (
	"crypto-transaction/config"
	"crypto-transaction/pkg/httpClient"
)

type HttpBlockBook interface {
	get(coin string, path string, res interface{}) error
	getHost(coin string) string
	GetStatus(coin string) (StatusResponse, error)
	GetAddress(coin string, address string) (AddressResponse, error)
	GetBlock(coin string, hash string) (BlockResponse, error)
	GetTransaction(coin string, txId string) (TransactionResponse, error)
	GetAddressUTXO(coin string, address string) ([]Utxo, error)
	BroadcastTransaction(coin string, hex string) (BroadcastTransactionResponse, error)
}

type httpBlockBook struct {
	c  config.Config
	hc httpClient.HttpClient
}

func (b *httpBlockBook) getHost(coin string) string {

	url := b.c.GetString("block-book." + coin + ".node")

	if url == "" {
		panic("error in getting block book node")
	}

	return url
}

func (b *httpBlockBook) get(coin string, path string, res interface{}) error {

	host := b.getHost(coin)

	url := host + path

	err := b.hc.SimpleGet(url, res)

	return err
}

func (b *httpBlockBook) GetStatus(coin string) (StatusResponse, error) {

	res := StatusResponse{}

	path := StatusPath

	err := b.get(coin, path, &res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (b *httpBlockBook) GetAddress(coin string, address string) (AddressResponse, error) {

	res := AddressResponse{}

	path := AddressPath + address

	err := b.get(coin, path, &res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (b *httpBlockBook) GetBlock(coin string, hash string) (BlockResponse, error) {

	res := BlockResponse{}

	path := BlockPath + hash

	err := b.get(coin, path, &res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (b *httpBlockBook) GetTransaction(coin string, txId string) (TransactionResponse, error) {

	res := TransactionResponse{}

	path := TXPath + txId

	err := b.get(coin, path, &res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (b *httpBlockBook) GetAddressUTXO(coin string, address string) ([]Utxo, error) {

	res := make([]Utxo, 0)

	path := UTXOPath + address

	err := b.get(coin, path, &res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (b *httpBlockBook) BroadcastTransaction(coin string, hex string) (BroadcastTransactionResponse, error) {

	res := BroadcastTransactionResponse{}

	path := BroadcastPath + hex

	err := b.get(coin, path, &res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func NewHttpBlockBookService(c config.Config, hc httpClient.HttpClient) HttpBlockBook {
	return &httpBlockBook{c, hc}
}
