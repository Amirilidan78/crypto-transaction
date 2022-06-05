package blockbook

type StatusResponse struct {
	BlockBook StatusBlockBook `json:"blockbook"`
	Backend   StatusBackend   `json:"backend"`
}

type StatusBlockBook struct {
	Coin          string `json:"coin"`
	Host          string `json:"host"`
	InSync        bool   `json:"inSync"`
	InSyncMempool bool   `json:"inSyncMempool"`
}

type StatusBackend struct {
	Chain         string `json:"chain"`
	BestBlockHash string `json:"bestBlockHash"`
	Blocks        int64  `json:"blocks"`
	Warnings      string `json:"warnings,omitempty"`
}

type AddressResponse struct {
	Page               int32          `json:"page"`
	TotalPages         int32          `json:"totalPages"`
	ItemsOnPage        int32          `json:"itemsOnPage"`
	Address            string         `json:"address"`
	Balance            string         `json:"balance"`
	TotalReceived      string         `json:"totalReceived"`
	TotalSent          string         `json:"totalSent"`
	UnconfirmedBalance string         `json:"unconfirmedBalance"`
	UnconfirmedTxs     int32          `json:"unconfirmedTxs"`
	Txs                int32          `json:"txs"`
	TxIds              []string       `json:"txids"`
	Nonce              string         `json:"nonce"`
	NonTokenTxs        int32          `json:"nonTokenTxs"`
	Tokens             []AddressToken `json:"tokens"`
}

type AddressToken struct {
	TokenType string `json:"type"`
	Name      string `json:"name"`
	Contract  string `json:"contract"`
	Transfers int32  `json:"transfers"`
	Symbol    string `json:"symbol"`
	SubUnit   int    `json:"decimals"`
	Balance   string `json:"balance"`
}

// ========================================== //

type Utxo struct {
	Txid          string `json:"txid"`
	Vout          uint32 `json:"vout"`
	Value         string `json:"value"`
	Height        int64  `json:"height"`
	Confirmations int64  `json:"confirmations"`
}

type BroadcastTransactionResponse struct {
	TxId string `json:"result"`
}

type TransactionResponse struct {
	TxId             string                    `json:"txid"`
	VIn              []ResponseTxVIn           `json:"vin"`
	VOut             []ResponseTxVOut          `json:"vout"`
	BlockHash        string                    `json:"blockHash"`
	BlockHeight      int64                     `json:"blockHeight"`
	Confirmations    int64                     `json:"confirmations"`
	BlockTime        int64                     `json:"blockTime"`
	Value            string                    `json:"value"`
	Fees             string                    `json:"fees"`
	TokenTransfers   []*TokenTransfers         `json:"tokenTransfers,omitempty"`
	EthereumSpecific *ResponseEthereumSpecific `json:"ethereumSpecific"`
}

type ResponseEthereumSpecific struct {
	Status   int    `json:"status"`
	Nonce    int64  `json:"nonce"`
	GasLimit int64  `json:"gasLimit"`
	GasUsed  int64  `json:"gasUsed"`
	GasPrice string `json:"gasPrice"`
}

type TokenTransfers struct {
	Type    string `json:"type"`
	From    string `json:"from"`
	To      string `json:"to"`
	Token   string `json:"token"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	SubUnit int    `json:"decimals"`
	Value   string `json:"value"`
}

type ResponseTxVIn struct {
	N         int64    `json:"n"`
	Addresses []string `json:"addresses"`
	IsAddress bool     `json:"isAddress"`
}

type ResponseTxVOut struct {
	Value     string   `json:"value"`
	N         int64    `json:"n"`
	Addresses []string `json:"addresses"`
	IsAddress bool     `json:"isAddress"`
}

// ==========================================

// ========================================== //

type WsResponse struct {
	Id   string                   `json:"id"`
	Data WsNewTransactionResponse `json:"data"`
}

type WsNewTransactionResponse struct {
	Address string              `json:"address"`
	Tx      TransactionResponse `json:"tx"`
}

type WsSubscribeNewTransaction struct {
	Id     string      `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}
