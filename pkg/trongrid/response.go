package trongrid

type GetAccountRequestBody struct {
	Address string `json:"address"`
	Visible bool   `json:"visible"`
}

type GetAccountResponseBody struct {
	Address               string `json:"address"`
	Balance               uint64 `json:"balance"`
	CreateTime            uint64 `json:"create_time"`
	LatestOprationTime    uint64 `json:"latest_opration_time"`
	LatestConsumeFreeTime uint64 `json:"latest_consume_free_time"`
}

type BlockResponseBody struct {
	BlockID      string        `json:"blockID"`
	BlockHeader  BlockHeader   `json:"block_header"`
	Transactions []Transaction `json:"transactions"`
}

type BlockHeader struct {
	WitnessSignature string             `json:"witness_signature"`
	RawData          BlockHeaderRawData `json:"raw_data"`
}

type BlockHeaderRawData struct {
	Number         int64  `json:"number"`
	TxTrieRoot     string `json:"txTrieRoot"`
	WitnessAddress string `json:"witness_address"`
	ParentHash     string `json:"parentHash"`
	Version        int32  `json:"version"`
	Timestamp      int64  `json:"timestamp"`
}

type Transaction struct {
	Ret        []TransactionRet   `json:"ret"`
	Signature  []string           `json:"signature"`
	TxID       string             `json:"txID"`
	RawData    TransactionRawData `json:"raw_data"`
	RawDataHex string             `json:"raw_data_hex"`
}

type TransactionRet struct {
	ContractRet string `json:"contractRet"`
}

type TransactionRawData struct {
	Contract      []TransactionContract `json:"contract"`
	RefBlockBytes string                `json:"ref_block_bytes"`
	RefBlockHash  string                `json:"ref_block_hash"`
	Expiration    uint64                `json:"expiration"`
	FeeLimit      uint64                `json:"fee_limit"`
	Timestamp     uint64                `json:"timestamp"`
}

type TransactionContract struct {
	Parameter TransactionContractParameter `json:"parameter"`
	Type      string                       `json:"type"`
}

type TransactionContractParameter struct {
	Value   TransactionContractParameterValue `json:"value"`
	TypeUrl string                            `json:"type_url"`
}

type TransactionContractParameterValue struct {
	Amount          uint64 `json:"amount"`
	OwnerAddress    string `json:"owner_address"`
	ToAddress       string `json:"to_address"`
	Data            string `json:"data"`
	ContractAddress string `json:"contract_address"`
}
