package trx

import (
	"crypto-transaction/config"
	"crypto-transaction/pkg/trongrid"
	"crypto-transaction/pkg/twallet"
	cryptoPb "crypto-transaction/pkg/twallet/proto"
	"crypto-transaction/services/coins/coin"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"math"
	"strconv"
	"time"
)

const CoinString = "TRX"

type Trx interface {
	CreateTransaction(amount string, fromAddress string, toAddress string, addressPrivateKey string) (proto.Message, error)
	SignTransaction(pb proto.Message) ([]byte, error)
	GetRawTransaction(res []byte) (string, error)
	BroadCastTransaction(hex string) (string, error)
}

type trx struct {
	c  config.Config
	tw twallet.TWallet
	tg trongrid.TronGrid
}

func (t *trx) CreateTransaction(amount string, fromAddress string, toAddress string, addressPrivateKey string) (proto.Message, error) {

	fee := t.calculateTransactionFeeLimit()

	sunAmount := t.getAmountInSun(amount)

	hexedPrivateKey, errPrivateKey := t.convertPrivateKeyToHexString(addressPrivateKey)

	if errPrivateKey != nil {
		return nil, errPrivateKey
	}

	blockHeaderResponseBody, errTransactionBlockHeader := t.getTransactionBlockHeader()

	if errTransactionBlockHeader != nil {
		return nil, errTransactionBlockHeader
	}

	blockHeader, errMakingTransactionBlockHeader := t.makeTransactionBlockHeader(blockHeaderResponseBody)

	if errMakingTransactionBlockHeader != nil {
		return nil, errMakingTransactionBlockHeader
	}

	now := time.Now()

	timestamp := now.Unix() * 1000

	expirationTimeStamp := blockHeader.Timestamp + 60*60*1000

	transferContract := &cryptoPb.TronTransferContract{
		OwnerAddress: fromAddress,
		ToAddress:    toAddress,
		Amount:       sunAmount,
	}

	txContract := &cryptoPb.TronTransaction_Transfer{
		Transfer: transferContract,
	}

	tx := &cryptoPb.TronTransaction{
		Timestamp:     timestamp,
		Expiration:    expirationTimeStamp,
		BlockHeader:   blockHeader,
		FeeLimit:      fee,
		ContractOneof: txContract,
	}

	si := &cryptoPb.TronSigningInput{
		Transaction: tx,
		PrivateKey:  hexedPrivateKey,
	}

	return si, nil
}

func (t *trx) SignTransaction(pb proto.Message) ([]byte, error) {
	return t.tw.SignTransaction(CoinString, pb)
}

func (t *trx) GetRawTransaction(res []byte) (string, error) {

	so := &cryptoPb.TronSigningOutput{}

	err := proto.Unmarshal(res, so)

	if err != nil {
		return "", err
	}

	return so.GetJson(), nil
}

func (t *trx) BroadCastTransaction(hex string) (string, error) {

	resp, err := t.tg.BroadcastTransaction(hex)

	if err != nil {
		return "", nil
	}

	if result, exist := resp["result"]; exist {
		if result.(bool) {
			return resp["txid"].(string), nil
		}
	}

	return "", errors.New(fmt.Sprintf("transaction broadcast error: %s", resp["code"].(string)))
}

func NewTrx(c config.Config, tw twallet.TWallet, tg trongrid.TronGrid) Trx {
	return &trx{c, tw, tg}
}

// ================================ trx specific ================================ //

func (t *trx) getAmountInSun(amount string) int64 {

	macroAmount, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		log.Fatal(err)
	}

	microAmount := macroAmount * math.Pow10(coin.GetCoinSubAmount(t.c, CoinString))

	return int64(microAmount)
}

func (t *trx) calculateTransactionFeeLimit() int64 {
	return coin.GetCoinFee(t.c, CoinString)
}

func (t *trx) convertPrivateKeyToHexString(privateKey string) ([]byte, error) {

	hexByte, err := coin.StringToHex(privateKey)

	if err != nil {
		return hexByte, err
	}

	return hexByte, nil
}

func (t *trx) getTransactionBlockHeader() (trongrid.BlockResponseBody, error) {

	respRaw, err := t.getCurrentBlockIndex()

	if err != nil {
		return trongrid.BlockResponseBody{}, err
	}

	getCurrentBlockIndexResponse, ok := respRaw.(trongrid.BlockResponseBody)

	if !ok {
		return trongrid.BlockResponseBody{}, errors.New("can not convert interface to tron.GetNowBlockResponseBody type")
	}

	return getCurrentBlockIndexResponse, nil
}

func (t *trx) getCurrentBlockIndex() (interface{}, error) {

	resp, err := t.tg.GetNowBlock()

	if err != nil {
		return "", err
	}

	return resp, nil
}

func (t *trx) makeTransactionBlockHeader(nowBlockResponseBody trongrid.BlockResponseBody) (*cryptoPb.TronBlockHeader, error) {

	blockHeaderRaw := nowBlockResponseBody.BlockHeader.RawData

	txTrieRootHex, errTxTrieRootHex := coin.StringToHex(blockHeaderRaw.TxTrieRoot)

	if errTxTrieRootHex != nil {
		return nil, errTxTrieRootHex
	}

	parentHash, errParentHash := coin.StringToHex(blockHeaderRaw.ParentHash)

	if errParentHash != nil {
		return nil, errParentHash
	}

	witnessAddress, errWitnessAddress := coin.StringToHex(blockHeaderRaw.WitnessAddress)

	if errWitnessAddress != nil {
		return nil, errWitnessAddress
	}

	blockHeader := &cryptoPb.TronBlockHeader{
		Timestamp:      blockHeaderRaw.Timestamp,
		TxTrieRoot:     txTrieRootHex,
		ParentHash:     parentHash,
		Number:         blockHeaderRaw.Number,
		WitnessAddress: witnessAddress,
		Version:        blockHeaderRaw.Version,
	}

	return blockHeader, nil
}
