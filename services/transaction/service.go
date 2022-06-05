package transaction

import (
	"crypto-transaction/config"
	"crypto-transaction/pkg/blockbook"
	"crypto-transaction/pkg/twallet"
	"crypto-transaction/services/transaction/btc"
	coinConfig "crypto-transaction/services/transaction/coin"
	"strings"
)

type Transaction interface {
	CreateTransaction(blockchain string, amount string, fromAddress string, toAddress string, addressPrivateKey string) (string, error)
}

type transaction struct {
	c  config.Config
	tw twallet.TWallet
	bb blockbook.HttpBlockBook
}

func (t *transaction) getCoin(blockchain string) coinConfig.Coin {
	switch strings.ToUpper(blockchain) {
	case "BTC":
		return btc.NewBtc(t.c, t.tw, t.bb)
	default:
		panic("Blockchain is not implemented")
	}
}

func (t *transaction) CreateTransaction(blockchain string, amount string, fromAddress string, toAddress string, addressPrivateKey string) (string, error) {

	coin := t.getCoin(blockchain)

	// tx hash
	tx, txErr := coin.CreateTransaction(amount, fromAddress, toAddress, addressPrivateKey)

	if txErr != nil {
		return "", txErr
	}

	// singed tx in byte
	signedTx, errSign := coin.SignTransaction(tx)

	if errSign != nil {
		return "", errSign
	}

	// hex tx string
	hex, rawErr := coin.GetRawTransaction(signedTx)

	if errSign != nil {
		return "", rawErr
	}

	// submitted transaction id
	txId, broadCastErr := coin.BroadCastTransaction(hex)

	if broadCastErr != nil {
		return "", broadCastErr
	}

	return txId, nil
}

func NewTransactionService(c config.Config, tw twallet.TWallet, bb blockbook.HttpBlockBook) Transaction {
	return &transaction{c, tw, bb}
}
