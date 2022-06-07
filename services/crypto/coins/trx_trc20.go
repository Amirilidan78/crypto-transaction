package coins

import (
	"crypto-transaction/config"
	"crypto-transaction/pkg/trongrid"
	"crypto-transaction/pkg/twallet"
	cryptoPb "crypto-transaction/pkg/twallet/proto"
	"crypto-transaction/services/crypto/common"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"math"
	"strconv"
	"time"
)

const Trc20Blockchain = "TRX"

type TrxTrc20 interface {
	CreateTransaction(coin string, amount string, fromAddress string, toAddress string, addressPrivateKey string) (proto.Message, error)
}

type trxTrc20 struct {
	c  config.Config
	tw twallet.TWallet
	tg trongrid.TronGrid
}

func (t *trxTrc20) CreateTransaction(coin string, amount string, fromAddress string, toAddress string, addressPrivateKey string) (proto.Message, error) {

	fee := t.calculateTransactionFeeLimit(coin)

	subAmount, errSubAmount := t.getSubAmount(coin, amount)

	if errSubAmount != nil {
		return nil, errSubAmount
	}

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

	contractAddress, errContract := t.getTrc20ContractAddress(coin)

	if errContract != nil {
		return nil, errContract
	}

	transferTrc20Contract := &cryptoPb.TronTransferTRC20Contract{
		ContractAddress: contractAddress,
		OwnerAddress:    fromAddress,
		ToAddress:       toAddress,
		Amount:          subAmount,
	}

	txContract := &cryptoPb.TronTransaction_TransferTrc20Contract{
		TransferTrc20Contract: transferTrc20Contract,
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

func NewTrxTrc20Coin(c config.Config, tw twallet.TWallet, tg trongrid.TronGrid) TrxTrc20 {
	return &trxTrc20{c, tw, tg}
}

func (t *trxTrc20) getSubAmount(coin string, amount string) ([]byte, error) {

	macroAmount, errParse := strconv.ParseFloat(amount, 64)

	if errParse != nil {
		return nil, errParse
	}

	microAmount := macroAmount * math.Pow10(common.GetTokenSubAmount(t.c, Trc20Blockchain, coin))

	tokenMicroAmountHexStr := fmt.Sprintf("%x", microAmount)

	hexedSubAmount, errHex := common.StringToHex(tokenMicroAmountHexStr)

	if errHex != nil {
		return nil, errHex
	}

	return hexedSubAmount, nil
}

func (t *trxTrc20) calculateTransactionFeeLimit(coin string) int64 {
	return common.GetCoinFee(t.c, coin)
}

func (t *trxTrc20) convertPrivateKeyToHexString(privateKey string) ([]byte, error) {

	hexByte, err := common.StringToHex(privateKey)

	if err != nil {
		return hexByte, err
	}

	return hexByte, nil
}

func (t *trxTrc20) getTransactionBlockHeader() (trongrid.BlockResponseBody, error) {

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

func (t *trxTrc20) getCurrentBlockIndex() (interface{}, error) {

	resp, err := t.tg.GetNowBlock()

	if err != nil {
		return "", err
	}

	return resp, nil
}

func (t *trxTrc20) makeTransactionBlockHeader(nowBlockResponseBody trongrid.BlockResponseBody) (*cryptoPb.TronBlockHeader, error) {

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

func (t *trxTrc20) getTrc20ContractAddress(coin string) (string, error) {

	contractAddress := common.GetCoinContractAddress(t.c, Trc20Blockchain, coin)

	if contractAddress == "" {
		return "", errors.New("contract address not found")
	}

	return contractAddress, nil
}
