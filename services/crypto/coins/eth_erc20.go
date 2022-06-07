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
	"math/big"
	"strconv"
)

const Erc20Blockchain = "ETH"
const ETHErc20ChainId = "01"
const ETHErc20GasLimit = 21000

type EthErc20 interface {
	CreateTransaction(coin string, amount string, fromAddress string, toAddress string, addressPrivateKey string) (proto.Message, error)
}

type ethErc20 struct {
	c  config.Config
	tw twallet.TWallet
	bb blockbook.HttpBlockBook
}

func (e *ethErc20) CreateTransaction(coin string, amount string, fromAddress string, toAddress string, addressPrivateKey string) (proto.Message, error) {

	fee := strconv.FormatInt(common.GetCoinFee(e.c, Erc20Blockchain), 10)

	_, _, gasPrice, gasLimit, errGasPrice := e.getRawGasPriceAndLimitFromConfig(fee)

	if errGasPrice != nil {

		_, _, gasPrice, gasLimit, errGasPrice = e.getRawGasPriceAndLimitFromBlockchain()

		if errGasPrice != nil {
			return nil, errGasPrice
		}

	}

	hexedGasPrice, errHexGasPrice := common.StringToHex(strconv.FormatInt(gasPrice, 64))

	hexedGasLimit, errHexGasLimit := common.StringToHex(strconv.FormatInt(gasLimit, 64))

	if errGasPrice != nil {
		return nil, errGasPrice
	}

	if errHexGasPrice != nil {
		return nil, errHexGasPrice
	}

	if errHexGasLimit != nil {
		return nil, errHexGasLimit
	}

	// converting private key to byte
	hexedPrivateKey, errPrivateKey := common.StringToHex(addressPrivateKey)

	if errPrivateKey != nil {
		return nil, errPrivateKey
	}

	// converting trxTrc20 to wei
	subAmount, errSubAmount := e.getSubAmount(coin, amount)

	if errSubAmount != nil {
		return nil, errSubAmount
	}

	// converting wei to byte
	hexedAmount, errAmount := common.StringToHex(subAmount)

	if errAmount != nil {
		return nil, errAmount
	}

	// getting address nonce
	nonce, errNonce := e.getNonce(fromAddress)

	if errNonce != nil {
		return nil, errNonce
	}

	// converting nonce to byte
	hexedNonce, errHexNonce := common.StringToHex(nonce)

	if errHexNonce != nil {
		return nil, errHexNonce
	}

	// getting contract address for token
	contractAddress, errContract := e.getErc20ContractAddress(coin)

	if errContract != nil {
		return nil, errContract
	}

	// getting chain id
	hexedChainId, errHexChainId := common.StringToHex(ETHErc20ChainId)

	if errHexChainId != nil {
		return nil, errHexChainId
	}

	// creating Eth transfer proto
	ethTransaction := &cryptoPb.EthTransaction{}

	// creating trxTrc20 transfer proto
	ethTransactionErc20Transfer := &cryptoPb.EthTransaction_ERC20Transfer{
		To:     toAddress,
		Amount: hexedAmount,
	}

	ethTransaction.TransactionOneof = &cryptoPb.EthTransaction_Erc20Transfer{Erc20Transfer: ethTransactionErc20Transfer}

	si := &cryptoPb.EthSigningInput{
		ChainId:     hexedChainId,
		Nonce:       hexedNonce,
		GasPrice:    hexedGasPrice,
		GasLimit:    hexedGasLimit,
		ToAddress:   contractAddress,
		PrivateKey:  hexedPrivateKey,
		Transaction: ethTransaction,
	}

	return si, nil
}

func NewEthErc20Coin(c config.Config, tw twallet.TWallet, bb blockbook.HttpBlockBook) EthErc20 {
	return &ethErc20{c, tw, bb}
}

func (e *ethErc20) getSubAmount(coin string, amount string) (string, error) {
	hex := ""
	floatAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return hex, err
	}
	floatWei := floatAmount * math.Pow10(common.GetTokenSubAmount(e.c, Erc20Blockchain, coin))
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

func (e *ethErc20) getNonce(address string) (string, error) {

	hexNonce := ""

	resp, err := e.bb.GetAddress(Erc20Blockchain, address)

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

func (e *ethErc20) getRawGasPriceAndLimitFromBlockchain() (int64, int64, int64, int64, error) {

	minGasPrice, maxGasPrice, avgGasPrice, err := e.estimateGasPriceFromLastBlock()

	if err != nil {
		return 0, 0, 0, 0, err
	}

	return minGasPrice, maxGasPrice, avgGasPrice, int64(ETHErc20GasLimit), nil
}

func (e *ethErc20) getRawGasPriceAndLimitFromConfig(fee string) (int64, int64, int64, int64, error) {

	feeFloat, err := strconv.ParseFloat(fee, 64)

	subAmountOfFee := feeFloat * math.Pow10(common.GetCoinSubAmount(e.c, Erc20Blockchain))

	if err != nil {
		return 0, 0, 0, 0, err
	}

	avgGasPrice := subAmountOfFee / ETHErc20GasLimit

	return 0, 0, int64(avgGasPrice), ETHErc20GasLimit, err
}

func (e *ethErc20) estimateGasPriceFromLastBlock() (int64, int64, int64, error) {

	resp, errStatus := e.bb.GetStatus(Erc20Blockchain)

	if errStatus != nil {
		return 0, 0, 0, errStatus
	}

	bestBlockHash := resp.Backend.BestBlockHash

	lastBlock, err := e.bb.GetBlock(Erc20Blockchain, bestBlockHash)
	if err != nil {
		return 0, 0, 0, err
	}

	minGasPrice := int64(0)
	maxGasPrice := int64(0)
	sumGasPrice := int64(0)

	for _, singleTx := range lastBlock.Txs {
		txGasPrice, err := strconv.ParseInt(singleTx.EthereumSpecific.GasPrice, 10, 64)
		if err != nil {
			return 0, 0, 0, err
		}

		sumGasPrice = sumGasPrice + txGasPrice

		if minGasPrice == 0 || txGasPrice < minGasPrice {
			minGasPrice = txGasPrice
		}

		if maxGasPrice == 0 || txGasPrice > maxGasPrice {
			maxGasPrice = txGasPrice
		}
	}

	avgGasPrice := sumGasPrice / lastBlock.TxCount

	return minGasPrice, maxGasPrice, avgGasPrice, nil
}

func (e *ethErc20) getErc20ContractAddress(coin string) (string, error) {

	contractAddress := common.GetCoinContractAddress(e.c, Erc20Blockchain, coin)

	if contractAddress == "" {
		return "", errors.New("contract address not found")
	}

	return contractAddress, nil
}
