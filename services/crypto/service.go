package crypto

import (
	"crypto-transaction/config"
	"crypto-transaction/pkg/blockbook"
	"crypto-transaction/pkg/trongrid"
	"crypto-transaction/pkg/twallet"
	"crypto-transaction/services/crypto/blockchains"
	"crypto-transaction/services/crypto/coins"
	coinConfig "crypto-transaction/services/crypto/common"
	"strings"
)

type Crypto interface {
	CreateTransaction(blockchain string, coin string, amount string, fromAddress string, toAddress string, addressPrivateKey string) (string, error)
}

type crypto struct {
	c  config.Config
	tw twallet.TWallet
	bb blockbook.HttpBlockBook
	tg trongrid.TronGrid
}

func (t *crypto) getBlockchain(blockchain string) coinConfig.Blockchain {
	switch strings.ToUpper(blockchain) {
	case "BITCOIN":
		return blockchains.NewBitcoinBlockchain(t.c, t.tw, t.bb)
	case "ETHEREUM":
		return blockchains.NewEthereumBlockchain(t.c, t.tw, t.bb)
	case "TRON":
		return blockchains.NewTronBlockchain(t.c, t.tw, t.tg)
	default:
		panic("Blockchain is not implemented")
	}
}

func (t *crypto) getCoin(coin string) coinConfig.Coin {
	switch strings.ToUpper(coin) {
	case "BTC":
		return coins.NewBtcCoin(t.c, t.tw, t.bb)
	case "ETH":
		return coins.NewEthCoin(t.c, t.tw, t.bb)
	case "TRX":
		return coins.NewTrxCoin(t.c, t.tw, t.tg)
	default:
		panic("Coin is not implemented")
	}
}

func (t *crypto) CreateTransaction(blockchainName string, coinName string, amount string, fromAddress string, toAddress string, addressPrivateKey string) (string, error) {

	blockchain := t.getBlockchain(blockchainName)
	coin := t.getCoin(coinName)

	// tx hash
	tx, txErr := coin.CreateTransaction(amount, fromAddress, toAddress, addressPrivateKey)

	if txErr != nil {
		return "", txErr
	}

	// singed tx in byte
	signedTx, errSign := blockchain.SignTransaction(tx)

	if errSign != nil {
		return "", errSign
	}

	// hex tx string
	hex, rawErr := blockchain.GetRawTransaction(signedTx)

	if rawErr != nil {
		return "", rawErr
	}

	// submitted crypto id
	txId, broadCastErr := blockchain.BroadCastTransaction(hex)

	if broadCastErr != nil {
		return "", broadCastErr
	}

	return txId, nil
}

func NewCryptoService(c config.Config, tw twallet.TWallet, bb blockbook.HttpBlockBook, tg trongrid.TronGrid) Crypto {
	return &crypto{c, tw, bb, tg}
}
