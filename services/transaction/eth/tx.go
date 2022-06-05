package eth

import (
	"crypto-transaction/config"
	"crypto-transaction/pkg/blockbook"
	"crypto-transaction/pkg/twallet"
	cryptoPb "crypto-transaction/pkg/twallet/proto"
	"crypto-transaction/services/transaction/coin"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"math"
	"math/big"
	"strconv"
)

const CoinString = "ETH"
const GasLimit = 21000

type Eth interface {
	CreateTransaction(amount string, fromAddress string, toAddress string, addressPrivateKey string) (proto.Message, error)
	SignTransaction(pb proto.Message) ([]byte, error)
	GetRawTransaction(res []byte) (string, error)
	BroadCastTransaction(hex string) (string, error)
}

type eth struct {
	c  config.Config
	tw twallet.TWallet
	bb blockbook.HttpBlockBook
}

func (e *eth) CreateTransaction(amount string, fromAddress string, toAddress string, addressPrivateKey string) (proto.Message, error) {

	fee := strconv.FormatInt(coin.GetCoinFee(e.c, CoinString), 10)

	// getting gas price
	_, _, gasPrice, gasLimit, err := e.GetRawGasPriceAndLimit(CoinString, fee)

	// converting private key to byte
	hexedPrivateKey, errPrivateKey := coin.StringToHex(addressPrivateKey)

	if errPrivateKey != nil {
		return nil, errPrivateKey
	}

	// converting eth to wei
	weiAmount, errWeiAmount := e.getAmountInWei(amount)

	if errWeiAmount != nil {
		return nil, errWeiAmount
	}

	// converting wei to byte
	hexedAmount, errAmount := coin.StringToHex(weiAmount)

	if errAmount != nil {
		return nil, errAmount
	}

	// getting address nonce
	nonce, errNonce := e.getNonce(fromAddress)

	if errNonce != nil {
		return nil, errNonce
	}

	// converting nonce to byte
	hexedNonce, errHexNonce := coin.StringToHex(nonce)

	if errHexNonce != nil {
		return nil, errHexNonce
	}

	// creating eth transfer proto
	ethTransaction := &cryptoPb.EthTransaction{}

	ethTransactionTransfer := &cryptoPb.EthTransaction_Transfer{
		Amount: hexedAmount,
		Data:   []byte(""),
	}

	ethTransactionTransfer_ := &cryptoPb.EthTransaction_Transfer_{Transfer: ethTransactionTransfer}

	ethTransaction.TransactionOneof = ethTransactionTransfer_

	fmt.Println(hexedPrivateKey)
	fmt.Println(hexedAmount)
	fmt.Println(hexedNonce)
	panic("1")
}

func (e *eth) SignTransaction(pb proto.Message) ([]byte, error) {
	return e.tw.SignTransaction(CoinString, pb)
}

func (e *eth) GetRawTransaction(res []byte) (string, error) {
	panic("1")
}

func (e *eth) BroadCastTransaction(hex string) (string, error) {

	resp, err := e.bb.BroadcastTransaction(CoinString, hex)

	if err != nil {
		return "", err
	}

	return resp.TxId, err
}

func NewEth(c config.Config, tw twallet.TWallet, bb blockbook.HttpBlockBook) Eth {
	return &eth{c, tw, bb}
}

// ================================ eth specific ================================ //

func (e *eth) getAmountInWei(amount string) (string, error) {
	hex := ""
	floatAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return hex, err
	}
	floatWei := floatAmount * math.Pow10(coin.GetCoinSubAmount(e.c, CoinString))
	strWei := strconv.FormatFloat(floatWei, 'f', 0, 64)
	bi, assigned := new(big.Int).SetString(strWei, 10)
	if !assigned {
		return hex, errors.New("strWei can not assigned to bi big number")
	}
	si, assigned := new(big.Int).SetString("16", 10)
	if !assigned {
		return hex, errors.New("strWei can not assigned to si big number")
	}
	mod := new(big.Int)
	for bi.Sign() > 0 {
		mod.Mod(bi, si)
		hex = strconv.FormatInt(mod.Int64(), 16) + hex
		bi.Sub(bi, mod)
		bi.Div(bi, si)
	}

	return hex, nil
}

func (e *eth) getNonce(address string) (string, error) {

	hexNonce := ""

	resp, err := e.bb.GetAddress(CoinString, address)

	if err != nil {
		return "", err
	}

	nonce, errConvertToString := strconv.ParseInt(resp.Nonce, 10, 64)

	if errConvertToString != nil {
		return hexNonce, errConvertToString
	}

	hexNonce = strconv.FormatInt(nonce, 16)

	if nonce < 16 {
		hexNonce = "0" + hexNonce
	}

	return hexNonce, nil
}

func (e *eth) getGasPrice(address string) (string, error) {
	panic("1")
}

// ================================ should be out of this module ================================ //
