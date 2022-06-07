package common

import (
	"crypto-transaction/config"
	"encoding/hex"
	"strings"
)

func StringToHex(str string) ([]byte, error) {

	if len(str)%2 == 1 {
		str = "0" + str
	}

	return hex.DecodeString(str)
}

func ReverseByte(input []byte) []byte {

	temp := len(input)

	var output []byte

	for i := temp - 1; i >= 0; i-- {
		output = append(output, input[i])
	}

	return output
}

func GetCoinSubAmount(c config.Config, coin string) int {
	return c.GetInt("coins." + coin + ".subAmount")
}

func GetTokenSubAmount(c config.Config, blockchain string, coin string) int {
	return c.GetInt("coins." + blockchain + "tokens." + coin + ".subAmount")
}

func GetCoinHashType(c config.Config, coin string) uint32 {
	return uint32(c.GetInt("coins." + coin + ".hashType"))
}

func GetCoinUTXOMinConfirmation(c config.Config, coin string) int64 {
	return int64(c.GetInt("coins." + coin + ".minUTXOConfirmation"))
}

func GetCoinFee(c config.Config, coin string) int64 {
	return int64(c.GetInt("coins." + coin + ".fee"))
}

func GetCoinSequenceUnitMax(c config.Config, coin string) uint32 {
	return uint32(c.GetInt("coins." + coin + ".sequenceUnitMax"))
}

func GetCoinContractAddress(c config.Config, blockchain string, coin string) string {
	return c.GetString("coins." + strings.ToLower(blockchain) + ".tokens." + strings.ToUpper(coin))
}
