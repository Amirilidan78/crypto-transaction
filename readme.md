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


I am using trezor public nodes for btc and eth check `/config/config.yml` `coins.btc.node`

and using shasta public api for trx `/config/config.yml` `coins.trx.node`


### TODOS
- add support for `ETH` and `TRX` tokens 
- add `BCH` &`LTC` &`DOGE` &`ZEC` &`DASH` &`ETC`blockchains 