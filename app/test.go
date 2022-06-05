package main

import (
	"crypto-transaction/config"
	"crypto-transaction/pkg/blockbook"
	"crypto-transaction/pkg/httpClient"
	"crypto-transaction/pkg/twallet"
	"crypto-transaction/services/transaction"
	"fmt"
)

func main() {
	c := config.NewConfig()
	hc := httpClient.NewHttpClient()
	bb := blockbook.NewHttpBlockBookService(c, hc)
	tw := twallet.NewTWallet()

	txService := transaction.NewTransactionService(c, tw, bb)

	res, err := txService.CreateTransaction("BTC", "0.00001", "bc1qmkzl7kj9m359tsj3t2kr9g7rsuv975ljzp26h2", "bc1qn7wfusd9ekded9cjr5xlncx8wz9zv5vuh7hjw6", "ba69e597c0316fbea91e5aaef58079d9a11e0c3eb4a5546995c166a04211830c")

	fmt.Println(res)
	fmt.Println(err)

	select {}
}
