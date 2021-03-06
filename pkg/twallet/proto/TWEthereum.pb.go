// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.9.0
// source: TWEthereum.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EthTransactionMode int32

const (
	EthTransactionMode_Legacy    EthTransactionMode = 0 // Legacy transaction, pre-EIP2718/EIP1559; for fee gasPrice/gasLimit is used
	EthTransactionMode_Enveloped EthTransactionMode = 1 // Enveloped transaction EIP2718 (with type 0x2), fee is according to EIP1559 (base fee, inclusion fee, ...)
)

// Enum value maps for EthTransactionMode.
var (
	EthTransactionMode_name = map[int32]string{
		0: "Legacy",
		1: "Enveloped",
	}
	EthTransactionMode_value = map[string]int32{
		"Legacy":    0,
		"Enveloped": 1,
	}
)

func (x EthTransactionMode) Enum() *EthTransactionMode {
	p := new(EthTransactionMode)
	*p = x
	return p
}

func (x EthTransactionMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EthTransactionMode) Descriptor() protoreflect.EnumDescriptor {
	return file_TWEthereum_proto_enumTypes[0].Descriptor()
}

func (EthTransactionMode) Type() protoreflect.EnumType {
	return &file_TWEthereum_proto_enumTypes[0]
}

func (x EthTransactionMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EthTransactionMode.Descriptor instead.
func (EthTransactionMode) EnumDescriptor() ([]byte, []int) {
	return file_TWEthereum_proto_rawDescGZIP(), []int{0}
}

// Transaction (transfer, smart contract call, ...)
type EthTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to TransactionOneof:
	//	*EthTransaction_Transfer_
	//	*EthTransaction_Erc20Transfer
	//	*EthTransaction_Erc20Approve
	//	*EthTransaction_Erc721Transfer
	//	*EthTransaction_Erc1155Transfer
	//	*EthTransaction_ContractGeneric_
	TransactionOneof isEthTransaction_TransactionOneof `protobuf_oneof:"transaction_oneof"`
}

func (x *EthTransaction) Reset() {
	*x = EthTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TWEthereum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthTransaction) ProtoMessage() {}

func (x *EthTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_TWEthereum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthTransaction.ProtoReflect.Descriptor instead.
func (*EthTransaction) Descriptor() ([]byte, []int) {
	return file_TWEthereum_proto_rawDescGZIP(), []int{0}
}

func (m *EthTransaction) GetTransactionOneof() isEthTransaction_TransactionOneof {
	if m != nil {
		return m.TransactionOneof
	}
	return nil
}

func (x *EthTransaction) GetTransfer() *EthTransaction_Transfer {
	if x, ok := x.GetTransactionOneof().(*EthTransaction_Transfer_); ok {
		return x.Transfer
	}
	return nil
}

func (x *EthTransaction) GetErc20Transfer() *EthTransaction_ERC20Transfer {
	if x, ok := x.GetTransactionOneof().(*EthTransaction_Erc20Transfer); ok {
		return x.Erc20Transfer
	}
	return nil
}

func (x *EthTransaction) GetErc20Approve() *EthTransaction_ERC20Approve {
	if x, ok := x.GetTransactionOneof().(*EthTransaction_Erc20Approve); ok {
		return x.Erc20Approve
	}
	return nil
}

func (x *EthTransaction) GetErc721Transfer() *EthTransaction_ERC721Transfer {
	if x, ok := x.GetTransactionOneof().(*EthTransaction_Erc721Transfer); ok {
		return x.Erc721Transfer
	}
	return nil
}

func (x *EthTransaction) GetErc1155Transfer() *EthTransaction_ERC1155Transfer {
	if x, ok := x.GetTransactionOneof().(*EthTransaction_Erc1155Transfer); ok {
		return x.Erc1155Transfer
	}
	return nil
}

func (x *EthTransaction) GetContractGeneric() *EthTransaction_ContractGeneric {
	if x, ok := x.GetTransactionOneof().(*EthTransaction_ContractGeneric_); ok {
		return x.ContractGeneric
	}
	return nil
}

type isEthTransaction_TransactionOneof interface {
	isEthTransaction_TransactionOneof()
}

type EthTransaction_Transfer_ struct {
	Transfer *EthTransaction_Transfer `protobuf:"bytes,1,opt,name=transfer,proto3,oneof"`
}

type EthTransaction_Erc20Transfer struct {
	Erc20Transfer *EthTransaction_ERC20Transfer `protobuf:"bytes,2,opt,name=erc20_transfer,json=erc20Transfer,proto3,oneof"`
}

type EthTransaction_Erc20Approve struct {
	Erc20Approve *EthTransaction_ERC20Approve `protobuf:"bytes,3,opt,name=erc20_approve,json=erc20Approve,proto3,oneof"`
}

type EthTransaction_Erc721Transfer struct {
	Erc721Transfer *EthTransaction_ERC721Transfer `protobuf:"bytes,4,opt,name=erc721_transfer,json=erc721Transfer,proto3,oneof"`
}

type EthTransaction_Erc1155Transfer struct {
	Erc1155Transfer *EthTransaction_ERC1155Transfer `protobuf:"bytes,5,opt,name=erc1155_transfer,json=erc1155Transfer,proto3,oneof"`
}

type EthTransaction_ContractGeneric_ struct {
	ContractGeneric *EthTransaction_ContractGeneric `protobuf:"bytes,6,opt,name=contract_generic,json=contractGeneric,proto3,oneof"`
}

func (*EthTransaction_Transfer_) isEthTransaction_TransactionOneof() {}

func (*EthTransaction_Erc20Transfer) isEthTransaction_TransactionOneof() {}

func (*EthTransaction_Erc20Approve) isEthTransaction_TransactionOneof() {}

func (*EthTransaction_Erc721Transfer) isEthTransaction_TransactionOneof() {}

func (*EthTransaction_Erc1155Transfer) isEthTransaction_TransactionOneof() {}

func (*EthTransaction_ContractGeneric_) isEthTransaction_TransactionOneof() {}

// Input data necessary to create a signed transaction.
// Legacy and EIP2718/EIP1559 transactions supported, see TransactionMode.
type EthSigningInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chain identifier (256-bit number)
	ChainId []byte `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// Nonce (256-bit number)
	Nonce []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Transaction version selector: Legacy or enveloped, has impact on fee structure.
	// Default is Legacy (value 0)
	TxMode EthTransactionMode `protobuf:"varint,3,opt,name=tx_mode,json=txMode,proto3,enum=proto.EthTransactionMode" json:"tx_mode,omitempty"`
	// Gas price (256-bit number)
	// Relevant for legacy transactions only (disregarded for enveloped/EIP1559)
	GasPrice []byte `protobuf:"bytes,4,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	// Gas limit (256-bit number)
	GasLimit []byte `protobuf:"bytes,5,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	// Maxinmum optional inclusion fee (aka tip) (256-bit number)
	// Relevant for enveloped/EIP1559 transactions only, tx_mode=Enveloped, (disregarded for legacy)
	MaxInclusionFeePerGas []byte `protobuf:"bytes,6,opt,name=max_inclusion_fee_per_gas,json=maxInclusionFeePerGas,proto3" json:"max_inclusion_fee_per_gas,omitempty"`
	// Maxinmum fee (256-bit number)
	// Relevant for enveloped/EIP1559 transactions only, tx_mode=Enveloped, (disregarded for legacy)
	MaxFeePerGas []byte `protobuf:"bytes,7,opt,name=max_fee_per_gas,json=maxFeePerGas,proto3" json:"max_fee_per_gas,omitempty"`
	// Recipient's address.
	ToAddress string `protobuf:"bytes,8,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	// Private key.
	PrivateKey  []byte          `protobuf:"bytes,9,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	Transaction *EthTransaction `protobuf:"bytes,10,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *EthSigningInput) Reset() {
	*x = EthSigningInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TWEthereum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthSigningInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthSigningInput) ProtoMessage() {}

func (x *EthSigningInput) ProtoReflect() protoreflect.Message {
	mi := &file_TWEthereum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthSigningInput.ProtoReflect.Descriptor instead.
func (*EthSigningInput) Descriptor() ([]byte, []int) {
	return file_TWEthereum_proto_rawDescGZIP(), []int{1}
}

func (x *EthSigningInput) GetChainId() []byte {
	if x != nil {
		return x.ChainId
	}
	return nil
}

func (x *EthSigningInput) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *EthSigningInput) GetTxMode() EthTransactionMode {
	if x != nil {
		return x.TxMode
	}
	return EthTransactionMode_Legacy
}

func (x *EthSigningInput) GetGasPrice() []byte {
	if x != nil {
		return x.GasPrice
	}
	return nil
}

func (x *EthSigningInput) GetGasLimit() []byte {
	if x != nil {
		return x.GasLimit
	}
	return nil
}

func (x *EthSigningInput) GetMaxInclusionFeePerGas() []byte {
	if x != nil {
		return x.MaxInclusionFeePerGas
	}
	return nil
}

func (x *EthSigningInput) GetMaxFeePerGas() []byte {
	if x != nil {
		return x.MaxFeePerGas
	}
	return nil
}

func (x *EthSigningInput) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *EthSigningInput) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *EthSigningInput) GetTransaction() *EthTransaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

// Transaction signing output.
type EthSigningOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Signed and encoded transaction bytes.
	Encoded []byte `protobuf:"bytes,1,opt,name=encoded,proto3" json:"encoded,omitempty"`
	V       []byte `protobuf:"bytes,2,opt,name=v,proto3" json:"v,omitempty"`
	R       []byte `protobuf:"bytes,3,opt,name=r,proto3" json:"r,omitempty"`
	S       []byte `protobuf:"bytes,4,opt,name=s,proto3" json:"s,omitempty"`
	// The payload part, supplied in the input or assembled from input parameters
	Data []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EthSigningOutput) Reset() {
	*x = EthSigningOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TWEthereum_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthSigningOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthSigningOutput) ProtoMessage() {}

func (x *EthSigningOutput) ProtoReflect() protoreflect.Message {
	mi := &file_TWEthereum_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthSigningOutput.ProtoReflect.Descriptor instead.
func (*EthSigningOutput) Descriptor() ([]byte, []int) {
	return file_TWEthereum_proto_rawDescGZIP(), []int{2}
}

func (x *EthSigningOutput) GetEncoded() []byte {
	if x != nil {
		return x.Encoded
	}
	return nil
}

func (x *EthSigningOutput) GetV() []byte {
	if x != nil {
		return x.V
	}
	return nil
}

func (x *EthSigningOutput) GetR() []byte {
	if x != nil {
		return x.R
	}
	return nil
}

func (x *EthSigningOutput) GetS() []byte {
	if x != nil {
		return x.S
	}
	return nil
}

func (x *EthSigningOutput) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Native coin transfer transaction
type EthTransaction_Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Amount to send in wei (256-bit number)
	Amount []byte `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// Optional payload data
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EthTransaction_Transfer) Reset() {
	*x = EthTransaction_Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TWEthereum_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthTransaction_Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthTransaction_Transfer) ProtoMessage() {}

func (x *EthTransaction_Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_TWEthereum_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthTransaction_Transfer.ProtoReflect.Descriptor instead.
func (*EthTransaction_Transfer) Descriptor() ([]byte, []int) {
	return file_TWEthereum_proto_rawDescGZIP(), []int{0, 0}
}

func (x *EthTransaction_Transfer) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *EthTransaction_Transfer) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// ERC20 token transfer transaction
type EthTransaction_ERC20Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	// Amount to send (256-bit number)
	Amount []byte `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *EthTransaction_ERC20Transfer) Reset() {
	*x = EthTransaction_ERC20Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TWEthereum_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthTransaction_ERC20Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthTransaction_ERC20Transfer) ProtoMessage() {}

func (x *EthTransaction_ERC20Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_TWEthereum_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthTransaction_ERC20Transfer.ProtoReflect.Descriptor instead.
func (*EthTransaction_ERC20Transfer) Descriptor() ([]byte, []int) {
	return file_TWEthereum_proto_rawDescGZIP(), []int{0, 1}
}

func (x *EthTransaction_ERC20Transfer) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *EthTransaction_ERC20Transfer) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

// ERC20 approve transaction
type EthTransaction_ERC20Approve struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spender string `protobuf:"bytes,1,opt,name=spender,proto3" json:"spender,omitempty"`
	// Amount to send (256-bit number)
	Amount []byte `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *EthTransaction_ERC20Approve) Reset() {
	*x = EthTransaction_ERC20Approve{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TWEthereum_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthTransaction_ERC20Approve) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthTransaction_ERC20Approve) ProtoMessage() {}

func (x *EthTransaction_ERC20Approve) ProtoReflect() protoreflect.Message {
	mi := &file_TWEthereum_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthTransaction_ERC20Approve.ProtoReflect.Descriptor instead.
func (*EthTransaction_ERC20Approve) Descriptor() ([]byte, []int) {
	return file_TWEthereum_proto_rawDescGZIP(), []int{0, 2}
}

func (x *EthTransaction_ERC20Approve) GetSpender() string {
	if x != nil {
		return x.Spender
	}
	return ""
}

func (x *EthTransaction_ERC20Approve) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

// ERC721 NFT transfer transaction
type EthTransaction_ERC721Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// ID of the token (256-bit number)
	TokenId []byte `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
}

func (x *EthTransaction_ERC721Transfer) Reset() {
	*x = EthTransaction_ERC721Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TWEthereum_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthTransaction_ERC721Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthTransaction_ERC721Transfer) ProtoMessage() {}

func (x *EthTransaction_ERC721Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_TWEthereum_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthTransaction_ERC721Transfer.ProtoReflect.Descriptor instead.
func (*EthTransaction_ERC721Transfer) Descriptor() ([]byte, []int) {
	return file_TWEthereum_proto_rawDescGZIP(), []int{0, 3}
}

func (x *EthTransaction_ERC721Transfer) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *EthTransaction_ERC721Transfer) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *EthTransaction_ERC721Transfer) GetTokenId() []byte {
	if x != nil {
		return x.TokenId
	}
	return nil
}

// ERC1155 NFT transfer transaction
type EthTransaction_ERC1155Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// ID of the token (256-bit number)
	TokenId []byte `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	// The amount of tokens being transferred
	Value []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Data  []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EthTransaction_ERC1155Transfer) Reset() {
	*x = EthTransaction_ERC1155Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TWEthereum_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthTransaction_ERC1155Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthTransaction_ERC1155Transfer) ProtoMessage() {}

func (x *EthTransaction_ERC1155Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_TWEthereum_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthTransaction_ERC1155Transfer.ProtoReflect.Descriptor instead.
func (*EthTransaction_ERC1155Transfer) Descriptor() ([]byte, []int) {
	return file_TWEthereum_proto_rawDescGZIP(), []int{0, 4}
}

func (x *EthTransaction_ERC1155Transfer) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *EthTransaction_ERC1155Transfer) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *EthTransaction_ERC1155Transfer) GetTokenId() []byte {
	if x != nil {
		return x.TokenId
	}
	return nil
}

func (x *EthTransaction_ERC1155Transfer) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *EthTransaction_ERC1155Transfer) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Generic smart contract transaction
type EthTransaction_ContractGeneric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Amount to send in wei (256-bit number)
	Amount []byte `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// Contract call payload data
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EthTransaction_ContractGeneric) Reset() {
	*x = EthTransaction_ContractGeneric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TWEthereum_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthTransaction_ContractGeneric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthTransaction_ContractGeneric) ProtoMessage() {}

func (x *EthTransaction_ContractGeneric) ProtoReflect() protoreflect.Message {
	mi := &file_TWEthereum_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthTransaction_ContractGeneric.ProtoReflect.Descriptor instead.
func (*EthTransaction_ContractGeneric) Descriptor() ([]byte, []int) {
	return file_TWEthereum_proto_rawDescGZIP(), []int{0, 5}
}

func (x *EthTransaction_ContractGeneric) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *EthTransaction_ContractGeneric) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_TWEthereum_proto protoreflect.FileDescriptor

var file_TWEthereum_proto_rawDesc = []byte{
	0x0a, 0x10, 0x54, 0x57, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x07, 0x0a, 0x0e, 0x45, 0x74,
	0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x08,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x74, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x00,
	0x52, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x0e, 0x65, 0x72,
	0x63, 0x32, 0x30, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x74, 0x68, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x52, 0x43, 0x32, 0x30, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0d, 0x65, 0x72, 0x63, 0x32, 0x30,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x0d, 0x65, 0x72, 0x63, 0x32,
	0x30, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x74, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x52, 0x43, 0x32, 0x30, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x65, 0x72, 0x63, 0x32, 0x30, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x12, 0x4f, 0x0a, 0x0f, 0x65, 0x72, 0x63, 0x37, 0x32, 0x31, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x74, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x52, 0x43, 0x37, 0x32, 0x31, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x48, 0x00, 0x52, 0x0e, 0x65, 0x72, 0x63, 0x37, 0x32, 0x31, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x10, 0x65, 0x72, 0x63, 0x31, 0x31, 0x35, 0x35, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x74, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x52, 0x43, 0x31, 0x31, 0x35, 0x35, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0f, 0x65, 0x72, 0x63, 0x31, 0x31, 0x35, 0x35,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x74, 0x68, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x1a, 0x36, 0x0a, 0x08,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x0d, 0x45, 0x52, 0x43, 0x32, 0x30, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x40, 0x0a,
	0x0c, 0x45, 0x52, 0x43, 0x32, 0x30, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a,
	0x4f, 0x0a, 0x0e, 0x45, 0x52, 0x43, 0x37, 0x32, 0x31, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64,
	0x1a, 0x7a, 0x0a, 0x0f, 0x45, 0x52, 0x43, 0x31, 0x31, 0x35, 0x35, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3d, 0x0a, 0x0f,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x13, 0x0a, 0x11, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x22, 0x8a, 0x03, 0x0a, 0x0f, 0x45, 0x74, 0x68, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x74, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x06, 0x74, 0x78, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x67, 0x61,
	0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x38, 0x0a, 0x19, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x15, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x63, 0x6c, 0x75,
	0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x12, 0x25, 0x0a,
	0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x46, 0x65, 0x65, 0x50, 0x65,
	0x72, 0x47, 0x61, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x74, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6a, 0x0a,
	0x10, 0x45, 0x74, 0x68, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x76,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x76, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x01, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x2f, 0x0a, 0x12, 0x45, 0x74, 0x68,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x0a, 0x0a, 0x06, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x45,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x64, 0x10, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x6b,
	0x67, 0x2f, 0x74, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TWEthereum_proto_rawDescOnce sync.Once
	file_TWEthereum_proto_rawDescData = file_TWEthereum_proto_rawDesc
)

func file_TWEthereum_proto_rawDescGZIP() []byte {
	file_TWEthereum_proto_rawDescOnce.Do(func() {
		file_TWEthereum_proto_rawDescData = protoimpl.X.CompressGZIP(file_TWEthereum_proto_rawDescData)
	})
	return file_TWEthereum_proto_rawDescData
}

var file_TWEthereum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_TWEthereum_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_TWEthereum_proto_goTypes = []interface{}{
	(EthTransactionMode)(0),                // 0: proto.EthTransactionMode
	(*EthTransaction)(nil),                 // 1: proto.EthTransaction
	(*EthSigningInput)(nil),                // 2: proto.EthSigningInput
	(*EthSigningOutput)(nil),               // 3: proto.EthSigningOutput
	(*EthTransaction_Transfer)(nil),        // 4: proto.EthTransaction.Transfer
	(*EthTransaction_ERC20Transfer)(nil),   // 5: proto.EthTransaction.ERC20Transfer
	(*EthTransaction_ERC20Approve)(nil),    // 6: proto.EthTransaction.ERC20Approve
	(*EthTransaction_ERC721Transfer)(nil),  // 7: proto.EthTransaction.ERC721Transfer
	(*EthTransaction_ERC1155Transfer)(nil), // 8: proto.EthTransaction.ERC1155Transfer
	(*EthTransaction_ContractGeneric)(nil), // 9: proto.EthTransaction.ContractGeneric
}
var file_TWEthereum_proto_depIdxs = []int32{
	4, // 0: proto.EthTransaction.transfer:type_name -> proto.EthTransaction.Transfer
	5, // 1: proto.EthTransaction.erc20_transfer:type_name -> proto.EthTransaction.ERC20Transfer
	6, // 2: proto.EthTransaction.erc20_approve:type_name -> proto.EthTransaction.ERC20Approve
	7, // 3: proto.EthTransaction.erc721_transfer:type_name -> proto.EthTransaction.ERC721Transfer
	8, // 4: proto.EthTransaction.erc1155_transfer:type_name -> proto.EthTransaction.ERC1155Transfer
	9, // 5: proto.EthTransaction.contract_generic:type_name -> proto.EthTransaction.ContractGeneric
	0, // 6: proto.EthSigningInput.tx_mode:type_name -> proto.EthTransactionMode
	1, // 7: proto.EthSigningInput.transaction:type_name -> proto.EthTransaction
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_TWEthereum_proto_init() }
func file_TWEthereum_proto_init() {
	if File_TWEthereum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TWEthereum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_TWEthereum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthSigningInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_TWEthereum_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthSigningOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_TWEthereum_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthTransaction_Transfer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_TWEthereum_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthTransaction_ERC20Transfer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_TWEthereum_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthTransaction_ERC20Approve); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_TWEthereum_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthTransaction_ERC721Transfer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_TWEthereum_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthTransaction_ERC1155Transfer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_TWEthereum_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthTransaction_ContractGeneric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_TWEthereum_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*EthTransaction_Transfer_)(nil),
		(*EthTransaction_Erc20Transfer)(nil),
		(*EthTransaction_Erc20Approve)(nil),
		(*EthTransaction_Erc721Transfer)(nil),
		(*EthTransaction_Erc1155Transfer)(nil),
		(*EthTransaction_ContractGeneric_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_TWEthereum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TWEthereum_proto_goTypes,
		DependencyIndexes: file_TWEthereum_proto_depIdxs,
		EnumInfos:         file_TWEthereum_proto_enumTypes,
		MessageInfos:      file_TWEthereum_proto_msgTypes,
	}.Build()
	File_TWEthereum_proto = out.File
	file_TWEthereum_proto_rawDesc = nil
	file_TWEthereum_proto_goTypes = nil
	file_TWEthereum_proto_depIdxs = nil
}
