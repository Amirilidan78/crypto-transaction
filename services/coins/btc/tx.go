package btc

import (
	"crypto-transaction/config"
	"crypto-transaction/pkg/blockbook"
	"crypto-transaction/pkg/twallet"
	cryptoPb "crypto-transaction/pkg/twallet/proto"
	"crypto-transaction/services/coins/coin"
	"encoding/hex"
	"errors"
	"github.com/golang/protobuf/proto"
	"math"
	"strconv"
)

const CoinString = "BTC"

type Btc interface {
	CreateTransaction(amount string, fromAddress string, toAddress string, addressPrivateKey string) (proto.Message, error)
	SignTransaction(pb proto.Message) ([]byte, error)
	GetRawTransaction(res []byte) (string, error)
	BroadCastTransaction(hex string) (string, error)
}

type btc struct {
	c  config.Config
	tw twallet.TWallet
	bb blockbook.HttpBlockBook
}

func (b *btc) CreateTransaction(amount string, fromAddress string, toAddress string, addressPrivateKey string) (proto.Message, error) {

	byteFee := coin.GetCoinFee(b.c, CoinString)

	satoshiAmount, errAmount := b.getAmountInSatoshi(amount)

	if errAmount != nil {
		return nil, errAmount
	}

	hexedPrivateKey, errPrivateKey := b.convertPrivateKeyToHexStringArray(addressPrivateKey)

	if errPrivateKey != nil {
		return nil, errPrivateKey
	}

	utxos, totalUTXOsAmount, utxoErr := b.getUnspentTransactions(fromAddress)

	if utxoErr != nil {
		return nil, utxoErr
	}

	if totalUTXOsAmount >= satoshiAmount {
		return nil, errors.New("not enough utxo for creating transaction")
	}

	si := &cryptoPb.BtcSigningInput{
		HashType:      coin.GetCoinHashType(b.c, CoinString),
		Amount:        satoshiAmount,
		ByteFee:       byteFee,
		ToAddress:     toAddress,
		ChangeAddress: fromAddress,
		PrivateKey:    hexedPrivateKey,
		Scripts:       nil,
		Utxo:          utxos,
		UseMaxAmount:  false,
		CoinType:      0,
	}

	return si, nil
}

func (b *btc) SignTransaction(pb proto.Message) ([]byte, error) {
	return b.tw.SignTransaction(CoinString, pb)
}

func (b *btc) GetRawTransaction(res []byte) (string, error) {

	so := &cryptoPb.BtcSigningOutput{}

	err := proto.Unmarshal(res, so)

	if err != nil {
		return "", err
	}

	if so.GetError() != cryptoPb.SigningError_OK {
		return "", errors.New("error in signing " + CoinString + " transaction. error: " + so.GetError().String())
	}

	return hex.EncodeToString(so.GetEncoded()), nil
}

func (b *btc) BroadCastTransaction(hex string) (string, error) {

	resp, err := b.bb.BroadcastTransaction(CoinString, hex)

	if err != nil {
		return "", err
	}

	return resp.TxId, err
}

func NewBtc(c config.Config, tw twallet.TWallet, bb blockbook.HttpBlockBook) Btc {
	return &btc{c, tw, bb}
}

// ================================ btc specific ================================ //

func (b *btc) getAmountInSatoshi(amount string) (int64, error) {

	btcAmount, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		return 0, err
	}

	satoshiAmount := btcAmount * math.Pow10(coin.GetCoinSubAmount(b.c, CoinString))

	return int64(satoshiAmount), nil
}

func (b *btc) convertPrivateKeyToHexStringArray(privateKey string) ([][]byte, error) {

	var hxPrivateKeysArr [][]byte

	hexByte, err := coin.StringToHex(privateKey)

	if err != nil {
		return hxPrivateKeysArr, err
	}

	if err != nil {
		return hxPrivateKeysArr, err

	}

	hxPrivateKeysArr = append(hxPrivateKeysArr, hexByte)

	return hxPrivateKeysArr, nil
}

func (b *btc) getUnspentTransactions(address string) ([]*cryptoPb.BtcUnspentTransaction, int64, error) {

	var result []*cryptoPb.BtcUnspentTransaction
	var totalAmount int64

	utxos, err := b.bb.GetAddressUTXO(CoinString, address)

	if err != nil {
		return result, totalAmount, err
	}

	for _, utxo := range utxos {
		if utxo.Confirmations >= coin.GetCoinUTXOMinConfirmation(b.c, CoinString) {

			utxId := utxo.Txid
			utxIdHex, err := coin.StringToHex(utxId)
			if err != nil {
				return result, totalAmount, err
			}
			utxIdReversed := coin.ReverseByte(utxIdHex)

			utxoOutPoit := cryptoPb.BtcOutPoint{
				Hash:     utxIdReversed,
				Index:    utxo.Vout,
				Sequence: coin.GetCoinSequenceUnitMax(b.c, CoinString),
			}

			utxoAmount, err := strconv.ParseInt(utxo.Value, 10, 64)

			if err != nil {
				return result, totalAmount, err
			}

			utxoItem := cryptoPb.BtcUnspentTransaction{
				OutPoint: &utxoOutPoit,
				Script:   nil, // TODO : TransactionInputScript
				Amount:   utxoAmount,
			}

			totalAmount = totalAmount + utxoAmount

			result = append(result, &utxoItem)
		}
	}

	return result, totalAmount, nil
}

// ================================ should be out of this module ================================ //
