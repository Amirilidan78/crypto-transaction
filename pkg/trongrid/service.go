package trongrid

import (
	"crypto-transaction/config"
	"crypto-transaction/pkg/httpClient"
	"encoding/json"
)

type TronGrid interface {
	GetAccount(address string) (GetAccountResponseBody, error)
	GetNowBlock() (BlockResponseBody, error)
	BroadcastTransaction(hex string) (map[string]interface{}, error)
}

type tronGrid struct {
	c  config.Config
	hc httpClient.HttpClient
}

func (t *tronGrid) getHost() string {

	url := t.c.GetString("coins.trx.node")

	if url == "" {
		panic("error in getting trx node")
	}

	return url
}

func (t *tronGrid) post(path string, body interface{}, res interface{}) error {

	host := t.getHost()

	url := host + path

	err := t.hc.SimplePost(url, body, res)

	return err
}

func (t *tronGrid) GetAccount(address string) (GetAccountResponseBody, error) {

	body := GetAccountRequestBody{
		Address: address,
		Visible: true,
	}

	res := GetAccountResponseBody{}

	err := t.post(GetAccountUri, body, &res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (t *tronGrid) GetNowBlock() (BlockResponseBody, error) {

	res := BlockResponseBody{}

	err := t.post(GetNowBlockUri, nil, &res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (t *tronGrid) BroadcastTransaction(hex string) (map[string]interface{}, error) {

	var txJsonBodyMap map[string]interface{}

	errJson := json.Unmarshal([]byte(hex), &txJsonBodyMap)

	if errJson != nil {
		return nil, errJson
	}

	var res map[string]interface{}

	err := t.post(BroadcastTransactionUri, txJsonBodyMap, &res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func NewTronGrid(c config.Config, hc httpClient.HttpClient) TronGrid {
	return &tronGrid{c, hc}
}
