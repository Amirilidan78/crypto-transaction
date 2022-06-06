package main

import (
	"crypto-transaction/config"
	"crypto-transaction/pkg/blockbook"
	"crypto-transaction/pkg/httpClient"
	"crypto-transaction/pkg/trongrid"
	"crypto-transaction/pkg/twallet"
	"crypto-transaction/services/coins"
	"fmt"
)

func main() {

	c := config.NewConfig()
	hc := httpClient.NewHttpClient()
	bb := blockbook.NewHttpBlockBookService(c, hc)
	tw := twallet.NewTWallet()
	tg := trongrid.NewTronGrid(c, hc)
	txService := coins.NewTransactionService(c, tw, bb, tg)

	res, err := txService.CreateTransaction("TRX", "10", "TM1KhZrrwCXK9i5BY2JhuGpghp9SDn9EMR", "TLtqH1B8RogdFPf6ehNQuEDyc7XuJb1ug5", "7e9e3da4c22953f7be47c0131df77ba497fdb0ad3c141739adda8908faff8e7e")

	fmt.Println("===========================================================")
	fmt.Println(res)
	fmt.Println(err)

	select {}
}
