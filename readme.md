# Crypto transaction 
simple crypto transaction repository for creating and submitting crypto transactions 

### Packages and services 
- Golang https://go.dev/
- Trust wallet core https://developer.trustwallet.com/wallet-core

### Quick start 
- `docker-compose up -d `
- `docker exec -it api-crypto-transaction-container go run ./app/test.go`

### Supported blockchains
- BTC
- ETH
- TRX

### Supported coins
- BTC
- ETH
- TRX

I am using trezor public nodes for btc and eth check `/config/config.yml` `coins.btc.node`

and using shasta public api for trx `/config/config.yml` `coins.trx.node`

To add more token add token `name` ,`subAmount` and `contract address` in `/config/config.yml` `coins.eth.tokens`

### TODOS
- add support for `TRC20` tokens 

### Example 
```

// create dependencies 
c := config.NewConfig()
hc := httpClient.NewHttpClient()
bb := blockbook.NewHttpBlockBookService(c, hc)
tw := twallet.NewTWallet()
tg := trongrid.NewTronGrid(c, hc)

// create service
cryptoService := crypto.NewCryptoService(c, tw, bb, tg)

// generate and submit transaction 
txId, err := cryptoService.CreateTransaction("TRON" ,"TRX" , "10", "TM1KhZrrwCXK9i5BY2JhuGpghp9SDn9EMR", "TLtqH1B8RogdFPf6ehNQuEDyc7XuJb1ug5", "privateKeyhere")

// token example  
txId, err := cryptoService.CreateTransaction("ETHEREUM" ,"USDT" , "10", "0x7fca062d4c1f7118b6e34fad8a95ec92e1753a4f", "0xf3e36ad56aa85abdacc18c02d19509ae4f7d5899", "privateKeyhere")

```