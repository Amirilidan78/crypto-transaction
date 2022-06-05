# Crypto transaction 
simple crypto transaction repository for creating and submitting crypto transactions 

### Packages and services 
- Golang https://go.dev/
- Trust wallet core https://developer.trustwallet.com/wallet-core

### Quick start 
- `docker-compose up -d `
- `docker exec -it api-crypto-transaction-container go run ./app/test.go`


### Supported blockchains 
- BTC -> implemented 
- ETH -> implemented
- TRX -> working on it ...


used trezor blockbook public nodes for submiting transactions ,getting utxos , getting address nonce and getting block details .
BTC : btc1.trezor.io
ETH : eth1.trezor.io