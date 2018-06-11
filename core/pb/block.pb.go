// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block.proto

/*
Package corepb is a generated protocol buffer package.

It is generated from these files:
	block.proto

It has these top-level messages:
	BlockHeader
	Block
	DownloadParentBlock
	Data
	Transaction
*/
package corepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type BlockHeader struct {
	Hash                 []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ParentHash           []byte `protobuf:"bytes,2,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	Coinbase             []byte `protobuf:"bytes,3,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	Timestamp            int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ChainId              uint32 `protobuf:"varint,5,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Alg                  uint32 `protobuf:"varint,6,opt,name=alg,proto3" json:"alg,omitempty"`
	Sign                 []byte `protobuf:"bytes,7,opt,name=sign,proto3" json:"sign,omitempty"`
	AccsRoot             []byte `protobuf:"bytes,8,opt,name=accs_root,json=accsRoot,proto3" json:"accs_root,omitempty"`
	TxsRoot              []byte `protobuf:"bytes,9,opt,name=txs_root,json=txsRoot,proto3" json:"txs_root,omitempty"`
	UsageRoot            []byte `protobuf:"bytes,10,opt,name=usage_root,json=usageRoot,proto3" json:"usage_root,omitempty"`
	RecordsRoot          []byte `protobuf:"bytes,11,opt,name=records_root,json=recordsRoot,proto3" json:"records_root,omitempty"`
	CandidacyRoot        []byte `protobuf:"bytes,12,opt,name=candidacy_root,json=candidacyRoot,proto3" json:"candidacy_root,omitempty"`
	ConsensusRoot        []byte `protobuf:"bytes,13,opt,name=consensus_root,json=consensusRoot,proto3" json:"consensus_root,omitempty"`
	ReservationQueueHash []byte `protobuf:"bytes,14,opt,name=reservation_queue_hash,json=reservationQueueHash,proto3" json:"reservation_queue_hash,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{0} }

func (m *BlockHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockHeader) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *BlockHeader) GetCoinbase() []byte {
	if m != nil {
		return m.Coinbase
	}
	return nil
}

func (m *BlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *BlockHeader) GetAlg() uint32 {
	if m != nil {
		return m.Alg
	}
	return 0
}

func (m *BlockHeader) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *BlockHeader) GetAccsRoot() []byte {
	if m != nil {
		return m.AccsRoot
	}
	return nil
}

func (m *BlockHeader) GetTxsRoot() []byte {
	if m != nil {
		return m.TxsRoot
	}
	return nil
}

func (m *BlockHeader) GetUsageRoot() []byte {
	if m != nil {
		return m.UsageRoot
	}
	return nil
}

func (m *BlockHeader) GetRecordsRoot() []byte {
	if m != nil {
		return m.RecordsRoot
	}
	return nil
}

func (m *BlockHeader) GetCandidacyRoot() []byte {
	if m != nil {
		return m.CandidacyRoot
	}
	return nil
}

func (m *BlockHeader) GetConsensusRoot() []byte {
	if m != nil {
		return m.ConsensusRoot
	}
	return nil
}

func (m *BlockHeader) GetReservationQueueHash() []byte {
	if m != nil {
		return m.ReservationQueueHash
	}
	return nil
}

type Block struct {
	Header       *BlockHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
	Height       uint64         `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{1} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *Block) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type DownloadParentBlock struct {
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Sign []byte `protobuf:"bytes,2,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (m *DownloadParentBlock) Reset()                    { *m = DownloadParentBlock{} }
func (m *DownloadParentBlock) String() string            { return proto.CompactTextString(m) }
func (*DownloadParentBlock) ProtoMessage()               {}
func (*DownloadParentBlock) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{2} }

func (m *DownloadParentBlock) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *DownloadParentBlock) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type Data struct {
	Type    string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Data) Reset()                    { *m = Data{} }
func (m *Data) String() string            { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()               {}
func (*Data) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{3} }

func (m *Data) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Data) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Transaction struct {
	Hash      []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	From      []byte `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To        []byte `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Value     []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Data      *Data  `protobuf:"bytes,6,opt,name=data" json:"data,omitempty"`
	Nonce     uint64 `protobuf:"varint,7,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ChainId   uint32 `protobuf:"varint,8,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Alg       uint32 `protobuf:"varint,9,opt,name=alg,proto3" json:"alg,omitempty"`
	Sign      []byte `protobuf:"bytes,10,opt,name=sign,proto3" json:"sign,omitempty"`
	PayerSign []byte `protobuf:"bytes,11,opt,name=payerSign,proto3" json:"payerSign,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{4} }

func (m *Transaction) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Transaction) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Transaction) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Transaction) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Transaction) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Transaction) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Transaction) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Transaction) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *Transaction) GetAlg() uint32 {
	if m != nil {
		return m.Alg
	}
	return 0
}

func (m *Transaction) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *Transaction) GetPayerSign() []byte {
	if m != nil {
		return m.PayerSign
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockHeader)(nil), "corepb.BlockHeader")
	proto.RegisterType((*Block)(nil), "corepb.Block")
	proto.RegisterType((*DownloadParentBlock)(nil), "corepb.DownloadParentBlock")
	proto.RegisterType((*Data)(nil), "corepb.Data")
	proto.RegisterType((*Transaction)(nil), "corepb.Transaction")
}

func init() { proto.RegisterFile("block.proto", fileDescriptorBlock) }

var fileDescriptorBlock = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0xd3, 0xf4, 0x23, 0x93, 0xb4, 0x42, 0xde, 0xd5, 0x2a, 0xc0, 0x22, 0x4a, 0x25, 0xa4,
	0x4a, 0x48, 0x3d, 0x2c, 0x2b, 0x71, 0xe2, 0x82, 0xf6, 0xb0, 0xdc, 0xc0, 0x70, 0xaf, 0xa6, 0x8e,
	0x69, 0x23, 0x5a, 0x3b, 0xd8, 0xee, 0xb2, 0xfd, 0x01, 0xdc, 0xf9, 0x3f, 0xfc, 0x39, 0xe4, 0x71,
	0xd2, 0xa6, 0xb0, 0x7b, 0x9b, 0x79, 0xef, 0x79, 0x3c, 0xf1, 0x7b, 0x81, 0x74, 0xb9, 0xd1, 0xe2,
	0xfb, 0xbc, 0x32, 0xda, 0x69, 0xd6, 0x17, 0xda, 0xc8, 0x6a, 0x39, 0xfd, 0xd3, 0x85, 0xf4, 0x83,
	0xc7, 0x6f, 0x25, 0x16, 0xd2, 0x30, 0x06, 0xf1, 0x1a, 0xed, 0x3a, 0xef, 0x4c, 0x3a, 0xb3, 0x8c,
	0x53, 0xcd, 0x5e, 0x42, 0x5a, 0xa1, 0x91, 0xca, 0x2d, 0x88, 0x8a, 0x88, 0x82, 0x00, 0xdd, 0x7a,
	0xc1, 0x33, 0x18, 0x0a, 0x5d, 0xaa, 0x25, 0x5a, 0x99, 0x77, 0x89, 0x3d, 0xf4, 0xec, 0x12, 0x12,
	0x57, 0x6e, 0xa5, 0x75, 0xb8, 0xad, 0xf2, 0x78, 0xd2, 0x99, 0x75, 0xf9, 0x11, 0x60, 0x4f, 0x61,
	0x28, 0xd6, 0x58, 0xaa, 0x45, 0x59, 0xe4, 0xbd, 0x49, 0x67, 0x36, 0xe2, 0x03, 0xea, 0x3f, 0x16,
	0xec, 0x09, 0x74, 0x71, 0xb3, 0xca, 0xfb, 0x84, 0xfa, 0xd2, 0xef, 0x66, 0xcb, 0x95, 0xca, 0x07,
	0x61, 0x37, 0x5f, 0xb3, 0xe7, 0x90, 0xa0, 0x10, 0x76, 0x61, 0xb4, 0x76, 0xf9, 0x30, 0xdc, 0xed,
	0x01, 0xae, 0xb5, 0xf3, 0xd3, 0xdd, 0x7d, 0xcd, 0x25, 0xc4, 0x0d, 0xdc, 0x7d, 0xa0, 0x5e, 0x00,
	0xec, 0x2c, 0xae, 0x64, 0x20, 0x81, 0xc8, 0x84, 0x10, 0xa2, 0x5f, 0x41, 0x66, 0xa4, 0xd0, 0xa6,
	0xa8, 0x4f, 0xa7, 0x24, 0x48, 0x6b, 0x8c, 0x24, 0xaf, 0x61, 0x2c, 0x50, 0x15, 0x65, 0x81, 0x62,
	0x1f, 0x44, 0x19, 0x89, 0x46, 0x07, 0xf4, 0x20, 0xd3, 0xca, 0x4a, 0x65, 0x77, 0xf5, 0xac, 0x51,
	0x2d, 0x6b, 0x50, 0x92, 0x5d, 0xc3, 0x85, 0x91, 0x56, 0x9a, 0x3b, 0x74, 0xa5, 0x56, 0x8b, 0x1f,
	0x3b, 0xb9, 0x93, 0xe1, 0xb9, 0xc7, 0x24, 0x3f, 0x6f, 0xb1, 0x9f, 0x3d, 0xe9, 0x1f, 0x7e, 0xfa,
	0xab, 0x03, 0x3d, 0x72, 0x8f, 0xbd, 0x81, 0xfe, 0x9a, 0x1c, 0x24, 0xe7, 0xd2, 0xab, 0xb3, 0x79,
	0x30, 0x78, 0xde, 0x32, 0x97, 0xd7, 0x12, 0xf6, 0x0e, 0x32, 0x67, 0x50, 0x59, 0x14, 0x7e, 0x9c,
	0xcd, 0xa3, 0x49, 0xb7, 0x7d, 0xe4, 0xeb, 0x91, 0xe3, 0x27, 0x42, 0x76, 0xe1, 0x6f, 0x29, 0x57,
	0x6b, 0x47, 0x36, 0xc7, 0xbc, 0xee, 0xa6, 0xef, 0xe1, 0xec, 0x46, 0xff, 0x54, 0x1b, 0x8d, 0xc5,
	0x27, 0x8a, 0x45, 0x58, 0xea, 0xa1, 0x30, 0x35, 0x26, 0x46, 0x47, 0x13, 0xa7, 0xd7, 0x10, 0xdf,
	0xa0, 0x43, 0xcf, 0xb9, 0x7d, 0x25, 0x49, 0x9f, 0x70, 0xaa, 0x59, 0x0e, 0x83, 0x0a, 0xf7, 0x7e,
	0x72, 0x7d, 0xa4, 0x69, 0xa7, 0xbf, 0x23, 0x48, 0x5b, 0xab, 0x3e, 0x76, 0xdb, 0x37, 0xa3, 0xb7,
	0xcd, 0x6d, 0xbe, 0x66, 0x63, 0x88, 0x9c, 0xae, 0x73, 0x1a, 0x39, 0xcd, 0xce, 0xa1, 0x77, 0x87,
	0x9b, 0x9d, 0xa4, 0x74, 0x66, 0x3c, 0x34, 0xa7, 0xb9, 0xed, 0xfd, 0x9b, 0xdb, 0x09, 0xc4, 0x05,
	0x3a, 0xa4, 0x74, 0xa6, 0x57, 0x59, 0xf3, 0x72, 0xfe, 0x2b, 0x38, 0x31, 0x7e, 0xaa, 0xd2, 0x4a,
	0x48, 0x4a, 0x6b, 0xcc, 0x43, 0x73, 0x92, 0xf7, 0xe1, 0x83, 0x79, 0x4f, 0xfe, 0xcf, 0x3b, 0xb4,
	0xf2, 0x7e, 0x09, 0x49, 0x85, 0x7b, 0x69, 0xbe, 0x78, 0x22, 0xa4, 0xf2, 0x08, 0x2c, 0xfb, 0xf4,
	0x73, 0xbf, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x63, 0x6e, 0x72, 0x2b, 0xeb, 0x03, 0x00, 0x00,
}
