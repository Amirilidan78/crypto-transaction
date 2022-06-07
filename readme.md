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

### TODOS
- add support for `ERC20` and `TRC20` tokens 

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
txId, err := cryptoService.CreateTransaction("TRON" ,"TRX" , "10", "TM1KhZrrwCXK9i5BY2JhuGpghp9SDn9EMR", "TLtqH1B8RogdFPf6ehNQuEDyc7XuJb1ug5", "7e9e3da4c22953f7be47c0131df77ba497fdb0ad3c141739adda8908faff8e7e")

fmt.Println(txId)  // nil 
fmt.Println(err) // nil 

```