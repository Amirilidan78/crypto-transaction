package coins

import (
	"crypto-transaction/config"
	"crypto-transaction/pkg/blockbook"
	"crypto-transaction/pkg/twallet"
	cryptoPb "crypto-transaction/pkg/twallet/proto"
	"crypto-transaction/services/crypto/common"
	"errors"
	"github.com/golang/protobuf/proto"
	"math"
	"strconv"
)

type Ltc interface {
	CreateTransaction(coin string, amount string, fromAddress string, toAddress string, addressPrivateKey string) (proto.Message, error)
}

type ltc struct {
	c  config.Config
	tw twallet.TWallet
	bb blockbook.HttpBlockBook
}

func (l *ltc) CreateTransaction(coin string, amount string, fromAddress string, toAddress string, addressPrivateKey string) (proto.Message, error) {

	byteFee := common.GetCoinFee(l.c, coin)

	satoshiAmount, errAmount := l.getAmountInLitoshi(coin, amount)

	if errAmount != nil {
		return nil, errAmount
	}

	hexedPrivateKey, errPrivateKey := l.convertPrivateKeyToHexStringArray(addressPrivateKey)

	if errPrivateKey != nil {
		return nil, errPrivateKey
	}

	utxos, totalUTXOsAmount, utxoErr := l.getUnspentTransactions(coin, fromAddress)

	if utxoErr != nil {
		return nil, utxoErr
	}

	if totalUTXOsAmount >= satoshiAmount {
		return nil, errors.New("not enough utxo for creating transaction")
	}

	si := &cryptoPb.BtcSigningInput{
		HashType:      common.GetCoinHashType(l.c, coin),
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

func (l *ltc) getAmountInLitoshi(coin string, amount string) (int64, error) {

	btcAmount, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		return 0, err
	}

	satoshiAmount := btcAmount * math.Pow10(int(common.GetCoinSubAmount(l.c, coin)))

	return int64(satoshiAmount), nil
}

func (l *ltc) convertPrivateKeyToHexStringArray(privateKey string) ([][]byte, error) {

	var hxPrivateKeysArr [][]byte

	hexByte, err := common.StringToHex(privateKey)

	if err != nil {
		return hxPrivateKeysArr, err
	}

	if err != nil {
		return hxPrivateKeysArr, err

	}

	hxPrivateKeysArr = append(hxPrivateKeysArr, hexByte)

	return hxPrivateKeysArr, nil
}

func (l *ltc) getUnspentTransactions(coin string, address string) ([]*cryptoPb.BtcUnspentTransaction, int64, error) {

	var result []*cryptoPb.BtcUnspentTransaction
	var totalAmount int64

	utxos, err := l.bb.GetAddressUTXO(coin, address)

	if err != nil {
		return result, totalAmount, err
	}

	for _, utxo := range utxos {
		if utxo.Confirmations >= common.GetCoinUTXOMinConfirmation(l.c, coin) {

			utxId := utxo.Txid
			utxIdHex, err := common.StringToHex(utxId)
			if err != nil {
				return result, totalAmount, err
			}
			utxIdReversed := common.ReverseByte(utxIdHex)

			utxoOutPoit := cryptoPb.BtcOutPoint{
				Hash:     utxIdReversed,
				Index:    utxo.Vout,
				Sequence: common.GetCoinSequenceUnitMax(l.c, coin),
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

func NewLtcCOin(c config.Config, tw twallet.TWallet, bb blockbook.HttpBlockBook) Ltc {
	return &ltc{c, tw, bb}
}
