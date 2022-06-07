package coins

import (
	"crypto-transaction/config"
	"crypto-transaction/pkg/trongrid"
	"crypto-transaction/pkg/twallet"
	cryptoPb "crypto-transaction/pkg/twallet/proto"
	"crypto-transaction/services/crypto/common"
	"errors"
	"github.com/golang/protobuf/proto"
	"log"
	"math"
	"strconv"
	"time"
)

type Trx interface {
	CreateTransaction(coin string, amount string, fromAddress string, toAddress string, addressPrivateKey string) (proto.Message, error)
}

type trx struct {
	c  config.Config
	tw twallet.TWallet
	tg trongrid.TronGrid
}

func (t *trx) CreateTransaction(coin string, amount string, fromAddress string, toAddress string, addressPrivateKey string) (proto.Message, error) {

	fee := t.calculateTransactionFeeLimit(coin)

	sunAmount := t.getAmountInSun(coin, amount)

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

func NewTrxCoin(c config.Config, tw twallet.TWallet, tg trongrid.TronGrid) Trx {
	return &trx{c, tw, tg}
}

func (t *trx) getAmountInSun(coin string, amount string) int64 {

	macroAmount, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		log.Fatal(err)
	}

	microAmount := macroAmount * math.Pow10(common.GetCoinSubAmount(t.c, coin))

	return int64(microAmount)
}

func (t *trx) calculateTransactionFeeLimit(coin string) int64 {
	return common.GetCoinFee(t.c, coin)
}

func (t *trx) convertPrivateKeyToHexString(privateKey string) ([]byte, error) {

	hexByte, err := common.StringToHex(privateKey)

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

	txTrieRootHex, errTxTrieRootHex := common.StringToHex(blockHeaderRaw.TxTrieRoot)

	if errTxTrieRootHex != nil {
		return nil, errTxTrieRootHex
	}

	parentHash, errParentHash := common.StringToHex(blockHeaderRaw.ParentHash)

	if errParentHash != nil {
		return nil, errParentHash
	}

	witnessAddress, errWitnessAddress := common.StringToHex(blockHeaderRaw.WitnessAddress)

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
