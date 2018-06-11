// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc.proto

/*
Package rpcpb is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	GetAccountStateRequest
	GetAccountStateResponse
	GetBlockRequest
	BlockResponse
	NonParamsRequest
	GetMedStateResponse
	GetTransactionRequest
	SendTransactionRequest
	SendTransactionResponse
	TransactionData
	TransactionResponse
*/
package rpcpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type GetAccountStateRequest struct {
	// Hex string of the account addresss.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// block account state with height. Or the string "genesis", "confirmed", "tail".
	Height string `protobuf:"bytes,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *GetAccountStateRequest) Reset()                    { *m = GetAccountStateRequest{} }
func (m *GetAccountStateRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAccountStateRequest) ProtoMessage()               {}
func (*GetAccountStateRequest) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{0} }

func (m *GetAccountStateRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *GetAccountStateRequest) GetHeight() string {
	if m != nil {
		return m.Height
	}
	return ""
}

type GetAccountStateResponse struct {
	// Current balance in unit of 1/(10^18) nas.
	Balance string `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	// Current transaction count.
	Nonce uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Account type
	Type uint32 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (m *GetAccountStateResponse) Reset()                    { *m = GetAccountStateResponse{} }
func (m *GetAccountStateResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAccountStateResponse) ProtoMessage()               {}
func (*GetAccountStateResponse) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{1} }

func (m *GetAccountStateResponse) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

func (m *GetAccountStateResponse) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *GetAccountStateResponse) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type GetBlockRequest struct {
	// Block hash. Or the string "genesis", "confirmed", "tail".
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *GetBlockRequest) Reset()                    { *m = GetBlockRequest{} }
func (m *GetBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBlockRequest) ProtoMessage()               {}
func (*GetBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{2} }

func (m *GetBlockRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type BlockResponse struct {
	// Block hash
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// Block parent hash
	ParentHash string `protobuf:"bytes,2,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	// Block coinbase address
	Coinbase string `protobuf:"bytes,3,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	// Block timestamp
	Timestamp int64 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Block chain id
	ChainId uint32 `protobuf:"varint,5,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// Block signature algorithm
	Alg uint32 `protobuf:"varint,6,opt,name=alg,proto3" json:"alg,omitempty"`
	// Block signature
	Sign string `protobuf:"bytes,7,opt,name=sign,proto3" json:"sign,omitempty"`
	// Root hash of accounts trie
	AccsRoot string `protobuf:"bytes,8,opt,name=accs_root,json=accsRoot,proto3" json:"accs_root,omitempty"`
	// Root hash of transactions trie
	TxsRoot string `protobuf:"bytes,9,opt,name=txs_root,json=txsRoot,proto3" json:"txs_root,omitempty"`
	// Root hash of usage trie
	UsageRoot string `protobuf:"bytes,10,opt,name=usage_root,json=usageRoot,proto3" json:"usage_root,omitempty"`
	// Root hash of records trie
	RecordsRoot string `protobuf:"bytes,11,opt,name=records_root,json=recordsRoot,proto3" json:"records_root,omitempty"`
	// Root hash of consensus trie
	ConsensusRoot string `protobuf:"bytes,12,opt,name=consensus_root,json=consensusRoot,proto3" json:"consensus_root,omitempty"`
	// Transactions in block
	Transactions []*TransactionResponse `protobuf:"bytes,13,rep,name=transactions" json:"transactions,omitempty"`
	// Block height
	Height uint64 `protobuf:"varint,14,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *BlockResponse) Reset()                    { *m = BlockResponse{} }
func (m *BlockResponse) String() string            { return proto.CompactTextString(m) }
func (*BlockResponse) ProtoMessage()               {}
func (*BlockResponse) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{3} }

func (m *BlockResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *BlockResponse) GetParentHash() string {
	if m != nil {
		return m.ParentHash
	}
	return ""
}

func (m *BlockResponse) GetCoinbase() string {
	if m != nil {
		return m.Coinbase
	}
	return ""
}

func (m *BlockResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockResponse) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *BlockResponse) GetAlg() uint32 {
	if m != nil {
		return m.Alg
	}
	return 0
}

func (m *BlockResponse) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *BlockResponse) GetAccsRoot() string {
	if m != nil {
		return m.AccsRoot
	}
	return ""
}

func (m *BlockResponse) GetTxsRoot() string {
	if m != nil {
		return m.TxsRoot
	}
	return ""
}

func (m *BlockResponse) GetUsageRoot() string {
	if m != nil {
		return m.UsageRoot
	}
	return ""
}

func (m *BlockResponse) GetRecordsRoot() string {
	if m != nil {
		return m.RecordsRoot
	}
	return ""
}

func (m *BlockResponse) GetConsensusRoot() string {
	if m != nil {
		return m.ConsensusRoot
	}
	return ""
}

func (m *BlockResponse) GetTransactions() []*TransactionResponse {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *BlockResponse) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type NonParamsRequest struct {
}

func (m *NonParamsRequest) Reset()                    { *m = NonParamsRequest{} }
func (m *NonParamsRequest) String() string            { return proto.CompactTextString(m) }
func (*NonParamsRequest) ProtoMessage()               {}
func (*NonParamsRequest) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{4} }

type GetMedStateResponse struct {
	// Block chain id
	ChainId uint32 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// Current tail block hash
	Tail string `protobuf:"bytes,2,opt,name=tail,proto3" json:"tail,omitempty"`
	// Current tail block height
	Height uint64 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	// The current med protocol version.
	ProtocolVersion string `protobuf:"bytes,7,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
	// Med version
	Version string `protobuf:"bytes,8,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *GetMedStateResponse) Reset()                    { *m = GetMedStateResponse{} }
func (m *GetMedStateResponse) String() string            { return proto.CompactTextString(m) }
func (*GetMedStateResponse) ProtoMessage()               {}
func (*GetMedStateResponse) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{5} }

func (m *GetMedStateResponse) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *GetMedStateResponse) GetTail() string {
	if m != nil {
		return m.Tail
	}
	return ""
}

func (m *GetMedStateResponse) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *GetMedStateResponse) GetProtocolVersion() string {
	if m != nil {
		return m.ProtocolVersion
	}
	return ""
}

func (m *GetMedStateResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type GetTransactionRequest struct {
	// Transaction hash
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *GetTransactionRequest) Reset()                    { *m = GetTransactionRequest{} }
func (m *GetTransactionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTransactionRequest) ProtoMessage()               {}
func (*GetTransactionRequest) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{6} }

func (m *GetTransactionRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type SendTransactionRequest struct {
	// Transaction hash
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// Hex string of the sender account addresss.
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	// Hex string of the receiver account addresss.
	To string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	// Amount of value sending with this transaction.
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	// Transaction timestamp.
	Timestamp int64 `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Transaction Data type.
	Data *TransactionData `protobuf:"bytes,6,opt,name=data" json:"data,omitempty"`
	// Transaction nonce.
	Nonce uint64 `protobuf:"varint,7,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Transaction chain ID.
	ChainId uint32 `protobuf:"varint,8,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// Transaction algorithm.
	Alg uint32 `protobuf:"varint,9,opt,name=alg,proto3" json:"alg,omitempty"`
	// Transaction sign.
	Sign string `protobuf:"bytes,10,opt,name=sign,proto3" json:"sign,omitempty"`
	// Transaction payer's sign.
	PayerSign string `protobuf:"bytes,11,opt,name=payer_sign,json=payerSign,proto3" json:"payer_sign,omitempty"`
}

func (m *SendTransactionRequest) Reset()                    { *m = SendTransactionRequest{} }
func (m *SendTransactionRequest) String() string            { return proto.CompactTextString(m) }
func (*SendTransactionRequest) ProtoMessage()               {}
func (*SendTransactionRequest) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{7} }

func (m *SendTransactionRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *SendTransactionRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *SendTransactionRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *SendTransactionRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *SendTransactionRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *SendTransactionRequest) GetData() *TransactionData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SendTransactionRequest) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *SendTransactionRequest) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *SendTransactionRequest) GetAlg() uint32 {
	if m != nil {
		return m.Alg
	}
	return 0
}

func (m *SendTransactionRequest) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *SendTransactionRequest) GetPayerSign() string {
	if m != nil {
		return m.PayerSign
	}
	return ""
}

type SendTransactionResponse struct {
	// Hex string of transaction hash.
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *SendTransactionResponse) Reset()                    { *m = SendTransactionResponse{} }
func (m *SendTransactionResponse) String() string            { return proto.CompactTextString(m) }
func (*SendTransactionResponse) ProtoMessage()               {}
func (*SendTransactionResponse) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{8} }

func (m *SendTransactionResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type TransactionData struct {
	// Transaction data type.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Transaction data payload.
	Payload string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *TransactionData) Reset()                    { *m = TransactionData{} }
func (m *TransactionData) String() string            { return proto.CompactTextString(m) }
func (*TransactionData) ProtoMessage()               {}
func (*TransactionData) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{9} }

func (m *TransactionData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TransactionData) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type TransactionResponse struct {
	// Transaction hash
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// Hex string of the sender account addresss.
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	// Hex string of the receiver account addresss.
	To string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	// Amount of value sending with this transaction.
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	// Transaction timestamp.
	Timestamp int64 `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Transaction Data type.
	Data *TransactionData `protobuf:"bytes,6,opt,name=data" json:"data,omitempty"`
	// Transaction nonce.
	Nonce uint64 `protobuf:"varint,7,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Transaction chain ID.
	ChainId uint32 `protobuf:"varint,8,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// Transaction algorithm.
	Alg uint32 `protobuf:"varint,9,opt,name=alg,proto3" json:"alg,omitempty"`
	// Transaction sign.
	Sign string `protobuf:"bytes,10,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (m *TransactionResponse) Reset()                    { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string            { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()               {}
func (*TransactionResponse) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{10} }

func (m *TransactionResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *TransactionResponse) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *TransactionResponse) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *TransactionResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *TransactionResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TransactionResponse) GetData() *TransactionData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *TransactionResponse) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *TransactionResponse) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *TransactionResponse) GetAlg() uint32 {
	if m != nil {
		return m.Alg
	}
	return 0
}

func (m *TransactionResponse) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAccountStateRequest)(nil), "rpcpb.GetAccountStateRequest")
	proto.RegisterType((*GetAccountStateResponse)(nil), "rpcpb.GetAccountStateResponse")
	proto.RegisterType((*GetBlockRequest)(nil), "rpcpb.GetBlockRequest")
	proto.RegisterType((*BlockResponse)(nil), "rpcpb.BlockResponse")
	proto.RegisterType((*NonParamsRequest)(nil), "rpcpb.NonParamsRequest")
	proto.RegisterType((*GetMedStateResponse)(nil), "rpcpb.GetMedStateResponse")
	proto.RegisterType((*GetTransactionRequest)(nil), "rpcpb.GetTransactionRequest")
	proto.RegisterType((*SendTransactionRequest)(nil), "rpcpb.SendTransactionRequest")
	proto.RegisterType((*SendTransactionResponse)(nil), "rpcpb.SendTransactionResponse")
	proto.RegisterType((*TransactionData)(nil), "rpcpb.TransactionData")
	proto.RegisterType((*TransactionResponse)(nil), "rpcpb.TransactionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ApiService service

type ApiServiceClient interface {
	GetAccountState(ctx context.Context, in *GetAccountStateRequest, opts ...grpc.CallOption) (*GetAccountStateResponse, error)
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*BlockResponse, error)
	GetMedState(ctx context.Context, in *NonParamsRequest, opts ...grpc.CallOption) (*GetMedStateResponse, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
	SendTransaction(ctx context.Context, in *SendTransactionRequest, opts ...grpc.CallOption) (*SendTransactionResponse, error)
}

type apiServiceClient struct {
	cc *grpc.ClientConn
}

func NewApiServiceClient(cc *grpc.ClientConn) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) GetAccountState(ctx context.Context, in *GetAccountStateRequest, opts ...grpc.CallOption) (*GetAccountStateResponse, error) {
	out := new(GetAccountStateResponse)
	err := grpc.Invoke(ctx, "/rpcpb.ApiService/GetAccountState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*BlockResponse, error) {
	out := new(BlockResponse)
	err := grpc.Invoke(ctx, "/rpcpb.ApiService/GetBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetMedState(ctx context.Context, in *NonParamsRequest, opts ...grpc.CallOption) (*GetMedStateResponse, error) {
	out := new(GetMedStateResponse)
	err := grpc.Invoke(ctx, "/rpcpb.ApiService/GetMedState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := grpc.Invoke(ctx, "/rpcpb.ApiService/GetTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) SendTransaction(ctx context.Context, in *SendTransactionRequest, opts ...grpc.CallOption) (*SendTransactionResponse, error) {
	out := new(SendTransactionResponse)
	err := grpc.Invoke(ctx, "/rpcpb.ApiService/SendTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApiService service

type ApiServiceServer interface {
	GetAccountState(context.Context, *GetAccountStateRequest) (*GetAccountStateResponse, error)
	GetBlock(context.Context, *GetBlockRequest) (*BlockResponse, error)
	GetMedState(context.Context, *NonParamsRequest) (*GetMedStateResponse, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*TransactionResponse, error)
	SendTransaction(context.Context, *SendTransactionRequest) (*SendTransactionResponse, error)
}

func RegisterApiServiceServer(s *grpc.Server, srv ApiServiceServer) {
	s.RegisterService(&_ApiService_serviceDesc, srv)
}

func _ApiService_GetAccountState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetAccountState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.ApiService/GetAccountState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetAccountState(ctx, req.(*GetAccountStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.ApiService/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetBlock(ctx, req.(*GetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetMedState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NonParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetMedState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.ApiService/GetMedState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetMedState(ctx, req.(*NonParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.ApiService/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_SendTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).SendTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.ApiService/SendTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).SendTransaction(ctx, req.(*SendTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcpb.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountState",
			Handler:    _ApiService_GetAccountState_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _ApiService_GetBlock_Handler,
		},
		{
			MethodName: "GetMedState",
			Handler:    _ApiService_GetMedState_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _ApiService_GetTransaction_Handler,
		},
		{
			MethodName: "SendTransaction",
			Handler:    _ApiService_SendTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptorRpc) }

var fileDescriptorRpc = []byte{
	// 818 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x55, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0x96, 0x7f, 0x92, 0x78, 0xca, 0xb1, 0x9d, 0x74, 0xb2, 0xf6, 0xec, 0x10, 0x83, 0x19, 0x69,
	0x25, 0xb3, 0x88, 0x58, 0x84, 0x1b, 0x07, 0xd0, 0x22, 0x24, 0xf3, 0x23, 0x10, 0x1a, 0x23, 0x6e,
	0x2b, 0xab, 0xdd, 0xd3, 0xd8, 0x23, 0xc6, 0xdd, 0xc3, 0x74, 0xdb, 0xda, 0x5c, 0x11, 0x6f, 0xc0,
	0x13, 0x70, 0xe1, 0x19, 0x78, 0x0f, 0x5e, 0x81, 0xc7, 0xe0, 0x80, 0xba, 0xba, 0xc7, 0x1e, 0xff,
	0xa1, 0x9c, 0xb9, 0x75, 0x55, 0x7d, 0xfe, 0xaa, 0xa6, 0xea, 0xab, 0x32, 0x78, 0x79, 0xc6, 0xee,
	0xb3, 0x5c, 0x6a, 0x49, 0xce, 0xf2, 0x8c, 0x65, 0xb3, 0xe0, 0x6e, 0x2e, 0xe5, 0x3c, 0xe5, 0x23,
	0x9a, 0x25, 0x23, 0x2a, 0x84, 0xd4, 0x54, 0x27, 0x52, 0x28, 0x0b, 0x0a, 0xbf, 0x82, 0xee, 0x98,
	0xeb, 0x57, 0x8c, 0xc9, 0x95, 0xd0, 0x13, 0x4d, 0x35, 0x8f, 0xf8, 0xcf, 0x2b, 0xae, 0x34, 0xf1,
	0xe1, 0x82, 0xc6, 0x71, 0xce, 0x95, 0xf2, 0x2b, 0x83, 0xca, 0xd0, 0x8b, 0x0a, 0x93, 0x74, 0xe1,
	0x7c, 0xc1, 0x93, 0xf9, 0x42, 0xfb, 0x55, 0x0c, 0x38, 0x2b, 0x7c, 0x0d, 0xbd, 0x03, 0x2e, 0x95,
	0x49, 0xa1, 0xb8, 0x21, 0x9b, 0xd1, 0x94, 0x0a, 0xc6, 0x0b, 0x32, 0x67, 0x92, 0x5b, 0x38, 0x13,
	0xd2, 0xf8, 0x0d, 0x57, 0x3d, 0xb2, 0x06, 0x21, 0x50, 0xd7, 0x8f, 0x19, 0xf7, 0x6b, 0x83, 0xca,
	0xb0, 0x15, 0xe1, 0x3b, 0x7c, 0x01, 0x9d, 0x31, 0xd7, 0x9f, 0xa5, 0x92, 0xfd, 0x54, 0xd4, 0x48,
	0xa0, 0xbe, 0xa0, 0x6a, 0xe1, 0x38, 0xf1, 0x1d, 0xfe, 0x59, 0x83, 0x96, 0x03, 0xb9, 0xe4, 0x47,
	0x50, 0xe4, 0x1d, 0x68, 0x66, 0x34, 0xe7, 0x42, 0x4f, 0x31, 0x64, 0x3f, 0x04, 0xac, 0xeb, 0x0b,
	0x03, 0x08, 0xa0, 0xc1, 0x64, 0x22, 0x66, 0x54, 0xd9, 0x2a, 0xbc, 0x68, 0x63, 0x93, 0x3b, 0xf0,
	0x74, 0xb2, 0xe4, 0x4a, 0xd3, 0x65, 0xe6, 0xd7, 0x07, 0x95, 0x61, 0x2d, 0xda, 0x3a, 0xc8, 0x73,
	0x68, 0xb0, 0x05, 0x4d, 0xc4, 0x34, 0x89, 0xfd, 0x33, 0xac, 0xff, 0x02, 0xed, 0x2f, 0x63, 0x72,
	0x05, 0x35, 0x9a, 0xce, 0xfd, 0x73, 0xf4, 0x9a, 0xa7, 0xa9, 0x4d, 0x25, 0x73, 0xe1, 0x5f, 0xd8,
	0xda, 0xcc, 0x9b, 0xbc, 0x05, 0x1e, 0x65, 0x4c, 0x4d, 0x73, 0x29, 0xb5, 0xdf, 0xb0, 0xb9, 0x8d,
	0x23, 0x92, 0x52, 0x1b, 0x76, 0xfd, 0xc6, 0xc5, 0x3c, 0xdb, 0x4a, 0xfd, 0xc6, 0x86, 0xfa, 0x00,
	0x2b, 0x45, 0xe7, 0xdc, 0x06, 0x01, 0x83, 0x1e, 0x7a, 0x30, 0xfc, 0x2e, 0x5c, 0xe6, 0x9c, 0xc9,
	0x3c, 0x76, 0xbf, 0x6e, 0x22, 0xa0, 0xe9, 0x7c, 0x08, 0x79, 0x01, 0x6d, 0x66, 0x5a, 0x26, 0xd4,
	0xca, 0x81, 0x2e, 0x11, 0xd4, 0xda, 0x78, 0x11, 0xf6, 0x09, 0x5c, 0xea, 0x9c, 0x0a, 0x45, 0x19,
	0x4a, 0xc9, 0x6f, 0x0d, 0x6a, 0xc3, 0xe6, 0x43, 0x70, 0x8f, 0x82, 0xbb, 0xff, 0x7e, 0x1b, 0x2a,
	0x46, 0x10, 0xed, 0xe0, 0x4b, 0x02, 0x6a, 0xe3, 0xd0, 0x0b, 0x01, 0x11, 0xb8, 0xfa, 0x56, 0x8a,
	0xef, 0x68, 0x4e, 0x97, 0xca, 0x8d, 0x38, 0xfc, 0xbd, 0x02, 0x37, 0x63, 0xae, 0xbf, 0xe1, 0xf1,
	0xae, 0xa2, 0xca, 0x5d, 0xae, 0xec, 0x76, 0xd9, 0x88, 0x87, 0x26, 0xa9, 0x1b, 0x2a, 0xbe, 0x4b,
	0x29, 0x6b, 0xe5, 0x94, 0xe4, 0x3d, 0xb8, 0xc2, 0x45, 0x60, 0x32, 0x9d, 0xae, 0x79, 0xae, 0x12,
	0x59, 0xcc, 0xa2, 0x53, 0xf8, 0x7f, 0xb0, 0x6e, 0xa3, 0xe1, 0x02, 0x61, 0x87, 0x52, 0x98, 0xe1,
	0xfb, 0xf0, 0x6c, 0xcc, 0xf5, 0xce, 0x77, 0x9f, 0xd6, 0xe7, 0x1f, 0x55, 0xe8, 0x4e, 0xb8, 0x88,
	0x9f, 0x06, 0x37, 0xbe, 0x1f, 0x73, 0xb9, 0x2c, 0x3e, 0xc6, 0xbc, 0x49, 0x1b, 0xaa, 0x5a, 0x3a,
	0x55, 0x56, 0xb5, 0x34, 0x3b, 0xb4, 0xa6, 0xe9, 0x8a, 0xa3, 0x16, 0xbd, 0xc8, 0x1a, 0xbb, 0x2a,
	0x3d, 0xdb, 0x57, 0xe9, 0x4b, 0xa8, 0xc7, 0x54, 0x53, 0xd4, 0x62, 0xf3, 0xa1, 0x7b, 0x38, 0xbb,
	0xcf, 0xa9, 0xa6, 0x11, 0x62, 0xb6, 0x3b, 0x7a, 0x51, 0xde, 0xd1, 0xf2, 0x04, 0x1a, 0x47, 0x75,
	0xee, 0x1d, 0xea, 0x1c, 0x4a, 0x3a, 0xef, 0x03, 0x64, 0xf4, 0x91, 0xe7, 0x53, 0x8c, 0x58, 0x39,
	0x7a, 0xe8, 0x99, 0x24, 0x73, 0x11, 0x7e, 0x00, 0xbd, 0x83, 0x3e, 0x9d, 0xde, 0xe8, 0xf0, 0x53,
	0xe8, 0xec, 0x55, 0xbf, 0xb9, 0x22, 0x0e, 0x66, 0xde, 0x66, 0x8a, 0x19, 0x7d, 0x4c, 0x25, 0x8d,
	0x5d, 0x4b, 0x0b, 0x33, 0xfc, 0xb5, 0x0a, 0x37, 0x4f, 0x4c, 0xf6, 0x3f, 0x9e, 0xca, 0xc3, 0x3f,
	0x35, 0x80, 0x57, 0x59, 0x32, 0xe1, 0xf9, 0x3a, 0x61, 0x9c, 0x48, 0xbc, 0xba, 0xe5, 0xa3, 0x4e,
	0xfa, 0xae, 0xac, 0xe3, 0x7f, 0x1c, 0xc1, 0xdb, 0xa7, 0xc2, 0xb6, 0x9f, 0x61, 0xff, 0x97, 0xbf,
	0xfe, 0xfe, 0xad, 0xda, 0x23, 0xcf, 0x46, 0xeb, 0x0f, 0x47, 0x2b, 0xc5, 0xf3, 0x11, 0xb5, 0x30,
	0x85, 0xec, 0x5f, 0x43, 0xa3, 0x38, 0xf3, 0xa4, 0xbb, 0xa5, 0x2a, 0xdf, 0xfd, 0xe0, 0xd6, 0xf9,
	0x77, 0xee, 0x7c, 0x78, 0x8d, 0xc4, 0x4d, 0xe2, 0x19, 0xe2, 0x19, 0x12, 0xbc, 0x86, 0x66, 0xe9,
	0x78, 0x90, 0x9e, 0xfb, 0xdd, 0xfe, 0x95, 0x09, 0x82, 0x6d, 0xa2, 0xfd, 0x4b, 0x13, 0x3e, 0x47,
	0xda, 0x1b, 0x72, 0x6d, 0x68, 0x85, 0x8c, 0xf9, 0x68, 0xc9, 0x63, 0x5b, 0x2b, 0x83, 0xf6, 0xee,
	0xe2, 0x93, 0xbb, 0x2d, 0xd1, 0xe1, 0x82, 0x07, 0xff, 0x71, 0x22, 0xc3, 0x1e, 0xa6, 0xb9, 0x26,
	0x1d, 0x93, 0xa6, 0x74, 0x2e, 0x49, 0x0a, 0x9d, 0xbd, 0x3d, 0xd8, 0x4c, 0xe0, 0xf8, 0x1d, 0xd9,
	0x4c, 0xe0, 0xc4, 0xfa, 0x84, 0x01, 0xa6, 0xba, 0x0d, 0xf7, 0x53, 0x7d, 0x5c, 0x79, 0x39, 0x3b,
	0xc7, 0xb3, 0xf7, 0xd1, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xc7, 0xce, 0x79, 0x49, 0x08,
	0x00, 0x00,
}
