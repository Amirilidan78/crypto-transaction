# Crypto transaction 
Simple crypto transaction repository for creating and submitting crypto transactions 


### Packages and services 
- Golang https://go.dev/
- Trust wallet core https://developer.trustwallet.com/wallet-core


### Setup
- `docker-compose up -d`


### Supported blockchains
- BTC
- ETH
- TRX


### Supported coins
- BTC
- ETH
- TRX
- BNB (ERC20)
- USDT (ERC20)
- USDC (ERC20)
- TUSD (ERC20)
- BUSD (ERC20)
- WBTC (ERC20)
- DAI (ERC20)
- SHIBA (ERC20)
- LINK (ERC20)
- UNI (ERC20)
- SAND (ERC20)
- USDT (TRC20)
- TUSD (TRC20)
- USDC (TRC20)
- USDD (TRC20)
- USDJ (TRC20)
- JST (TRC20)
- BTT (TRC20)

### Add more token ERC20 and TRC20 token

Erc20 token : add token `name` ,`subAmount` and `contract address` in `/config/config.yml` `coins.eth.tokens`

Trc20 token : add token `name` ,`subAmount` and `contract address` in `/config/config.yml` `coins.trx.tokens`


### Usage 
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

### Attention 

this project is for learning purpose only

this project is using trezor public nodes and shasta public node 

`/config/config.yml` `coins.***.node`
