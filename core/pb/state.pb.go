// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: state.proto

/*
Package corepb is a generated protocol buffer package.

It is generated from these files:
	state.proto

It has these top-level messages:
	Account
	DataState
	Record
	Certification
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

type Account struct {
	Address         []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance         []byte `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Nonce           uint64 `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Vesting         []byte `protobuf:"bytes,11,opt,name=vesting,proto3" json:"vesting,omitempty"`
	VotedRootHash   []byte `protobuf:"bytes,12,opt,name=voted_root_hash,json=votedRootHash,proto3" json:"voted_root_hash,omitempty"`
	Bandwidth       []byte `protobuf:"bytes,13,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	LastBandwidthTs int64  `protobuf:"varint,14,opt,name=last_bandwidth_ts,json=lastBandwidthTs,proto3" json:"last_bandwidth_ts,omitempty"`
	Unstaking       []byte `protobuf:"bytes,15,opt,name=unstaking,proto3" json:"unstaking,omitempty"`
	LastUnstakingTs int64  `protobuf:"varint,16,opt,name=last_unstaking_ts,json=lastUnstakingTs,proto3" json:"last_unstaking_ts,omitempty"`
	Collateral      []byte `protobuf:"bytes,21,opt,name=collateral,proto3" json:"collateral,omitempty"`
	VotersRootHash  []byte `protobuf:"bytes,22,opt,name=voters_root_hash,json=votersRootHash,proto3" json:"voters_root_hash,omitempty"`
	VotePower       []byte `protobuf:"bytes,23,opt,name=vote_power,json=votePower,proto3" json:"vote_power,omitempty"`
	TxsFromRootHash []byte `protobuf:"bytes,31,opt,name=txs_from_root_hash,json=txsFromRootHash,proto3" json:"txs_from_root_hash,omitempty"`
	TxsToRootHash   []byte `protobuf:"bytes,32,opt,name=txs_to_root_hash,json=txsToRootHash,proto3" json:"txs_to_root_hash,omitempty"`
	DataRootHash    []byte `protobuf:"bytes,40,opt,name=data_root_hash,json=dataRootHash,proto3" json:"data_root_hash,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptorState, []int{0} }

func (m *Account) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Account) GetBalance() []byte {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *Account) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Account) GetVesting() []byte {
	if m != nil {
		return m.Vesting
	}
	return nil
}

func (m *Account) GetVotedRootHash() []byte {
	if m != nil {
		return m.VotedRootHash
	}
	return nil
}

func (m *Account) GetBandwidth() []byte {
	if m != nil {
		return m.Bandwidth
	}
	return nil
}

func (m *Account) GetLastBandwidthTs() int64 {
	if m != nil {
		return m.LastBandwidthTs
	}
	return 0
}

func (m *Account) GetUnstaking() []byte {
	if m != nil {
		return m.Unstaking
	}
	return nil
}

func (m *Account) GetLastUnstakingTs() int64 {
	if m != nil {
		return m.LastUnstakingTs
	}
	return 0
}

func (m *Account) GetCollateral() []byte {
	if m != nil {
		return m.Collateral
	}
	return nil
}

func (m *Account) GetVotersRootHash() []byte {
	if m != nil {
		return m.VotersRootHash
	}
	return nil
}

func (m *Account) GetVotePower() []byte {
	if m != nil {
		return m.VotePower
	}
	return nil
}

func (m *Account) GetTxsFromRootHash() []byte {
	if m != nil {
		return m.TxsFromRootHash
	}
	return nil
}

func (m *Account) GetTxsToRootHash() []byte {
	if m != nil {
		return m.TxsToRootHash
	}
	return nil
}

func (m *Account) GetDataRootHash() []byte {
	if m != nil {
		return m.DataRootHash
	}
	return nil
}

type DataState struct {
	TxStateRootHash            []byte `protobuf:"bytes,1,opt,name=tx_state_root_hash,json=txStateRootHash,proto3" json:"tx_state_root_hash,omitempty"`
	RecordStateRootHash        []byte `protobuf:"bytes,2,opt,name=record_state_root_hash,json=recordStateRootHash,proto3" json:"record_state_root_hash,omitempty"`
	CertificationStateRootHash []byte `protobuf:"bytes,3,opt,name=certification_state_root_hash,json=certificationStateRootHash,proto3" json:"certification_state_root_hash,omitempty"`
}

func (m *DataState) Reset()                    { *m = DataState{} }
func (m *DataState) String() string            { return proto.CompactTextString(m) }
func (*DataState) ProtoMessage()               {}
func (*DataState) Descriptor() ([]byte, []int) { return fileDescriptorState, []int{1} }

func (m *DataState) GetTxStateRootHash() []byte {
	if m != nil {
		return m.TxStateRootHash
	}
	return nil
}

func (m *DataState) GetRecordStateRootHash() []byte {
	if m != nil {
		return m.RecordStateRootHash
	}
	return nil
}

func (m *DataState) GetCertificationStateRootHash() []byte {
	if m != nil {
		return m.CertificationStateRootHash
	}
	return nil
}

type Record struct {
	RecordHash []byte `protobuf:"bytes,1,opt,name=record_hash,json=recordHash,proto3" json:"record_hash,omitempty"`
	Owner      []byte `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Timestamp  int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptorState, []int{2} }

func (m *Record) GetRecordHash() []byte {
	if m != nil {
		return m.RecordHash
	}
	return nil
}

func (m *Record) GetOwner() []byte {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Record) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Certification struct {
	CertificateHash []byte `protobuf:"bytes,1,opt,name=certificate_hash,json=certificateHash,proto3" json:"certificate_hash,omitempty"`
	Issuer          []byte `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Certified       []byte `protobuf:"bytes,3,opt,name=certified,proto3" json:"certified,omitempty"`
	IssueTime       int64  `protobuf:"varint,4,opt,name=issue_time,json=issueTime,proto3" json:"issue_time,omitempty"`
	ExpirationTime  int64  `protobuf:"varint,5,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	RevocationTime  int64  `protobuf:"varint,6,opt,name=revocation_time,json=revocationTime,proto3" json:"revocation_time,omitempty"`
}

func (m *Certification) Reset()                    { *m = Certification{} }
func (m *Certification) String() string            { return proto.CompactTextString(m) }
func (*Certification) ProtoMessage()               {}
func (*Certification) Descriptor() ([]byte, []int) { return fileDescriptorState, []int{3} }

func (m *Certification) GetCertificateHash() []byte {
	if m != nil {
		return m.CertificateHash
	}
	return nil
}

func (m *Certification) GetIssuer() []byte {
	if m != nil {
		return m.Issuer
	}
	return nil
}

func (m *Certification) GetCertified() []byte {
	if m != nil {
		return m.Certified
	}
	return nil
}

func (m *Certification) GetIssueTime() int64 {
	if m != nil {
		return m.IssueTime
	}
	return 0
}

func (m *Certification) GetExpirationTime() int64 {
	if m != nil {
		return m.ExpirationTime
	}
	return 0
}

func (m *Certification) GetRevocationTime() int64 {
	if m != nil {
		return m.RevocationTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Account)(nil), "corepb.Account")
	proto.RegisterType((*DataState)(nil), "corepb.DataState")
	proto.RegisterType((*Record)(nil), "corepb.Record")
	proto.RegisterType((*Certification)(nil), "corepb.Certification")
}

func init() { proto.RegisterFile("state.proto", fileDescriptorState) }

var fileDescriptorState = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x94, 0xdb, 0x8e, 0xd3, 0x3c,
	0x10, 0xc7, 0x95, 0xaf, 0x87, 0xd5, 0x4e, 0x0f, 0xe9, 0x67, 0x96, 0x62, 0x21, 0x96, 0xad, 0x2a,
	0xc4, 0x06, 0x90, 0xb8, 0xd9, 0x27, 0x58, 0x40, 0x88, 0x4b, 0x14, 0xca, 0x25, 0x8a, 0xdc, 0xc4,
	0x4b, 0x23, 0x92, 0x4c, 0x64, 0x4f, 0x0f, 0x4f, 0xc1, 0xb3, 0xf0, 0x4a, 0xbc, 0x09, 0xb2, 0x9d,
	0x83, 0xab, 0xbd, 0x9c, 0x9f, 0x7f, 0x9e, 0xf9, 0x3b, 0x1a, 0x05, 0x26, 0x9a, 0x04, 0xc9, 0xf7,
	0xb5, 0x42, 0x42, 0x36, 0x4e, 0x51, 0xc9, 0x7a, 0xbb, 0xfe, 0x3d, 0x84, 0x8b, 0xfb, 0x34, 0xc5,
	0x7d, 0x45, 0x8c, 0xc3, 0x85, 0xc8, 0x32, 0x25, 0xb5, 0xe6, 0xc1, 0x2a, 0x88, 0xa6, 0x71, 0x5b,
	0x9a, 0x93, 0xad, 0x28, 0x44, 0x95, 0x4a, 0xfe, 0x9f, 0x3b, 0x69, 0x4a, 0x76, 0x05, 0xa3, 0x0a,
	0x0d, 0x1f, 0xac, 0x82, 0x68, 0x18, 0xbb, 0xc2, 0xf8, 0x07, 0xa9, 0x29, 0xaf, 0x7e, 0xf2, 0x89,
	0xf3, 0x9b, 0x92, 0xbd, 0x86, 0xf0, 0x80, 0x24, 0xb3, 0x44, 0x21, 0x52, 0xb2, 0x13, 0x7a, 0xc7,
	0xa7, 0xd6, 0x98, 0x59, 0x1c, 0x23, 0xd2, 0x17, 0xa1, 0x77, 0xec, 0x05, 0x5c, 0x6e, 0x45, 0x95,
	0x1d, 0xf3, 0x8c, 0x76, 0x7c, 0x66, 0x8d, 0x1e, 0xb0, 0xb7, 0xf0, 0x7f, 0x21, 0x34, 0x25, 0x1d,
	0x49, 0x48, 0xf3, 0xf9, 0x2a, 0x88, 0x06, 0x71, 0x68, 0x0e, 0x3e, 0xb4, 0x7c, 0xa3, 0x4d, 0xa7,
	0x7d, 0xa5, 0x49, 0xfc, 0x32, 0x69, 0x42, 0xd7, 0xa9, 0x03, 0x5d, 0xa7, 0x8e, 0x98, 0x4e, 0x8b,
	0xbe, 0xd3, 0xf7, 0x96, 0x6f, 0x34, 0x7b, 0x09, 0x90, 0x62, 0x51, 0x08, 0x92, 0x4a, 0x14, 0xfc,
	0xa9, 0x6d, 0xe5, 0x11, 0x16, 0xc1, 0xc2, 0x3c, 0x42, 0x69, 0xef, 0x71, 0x4b, 0x6b, 0xcd, 0x1d,
	0xef, 0x5e, 0x77, 0x0d, 0x60, 0x48, 0x52, 0xe3, 0x51, 0x2a, 0xfe, 0xcc, 0x85, 0x32, 0xe4, 0xab,
	0x01, 0xec, 0x1d, 0x30, 0x3a, 0xe9, 0xe4, 0x41, 0x61, 0xe9, 0xb5, 0xba, 0xb1, 0x5a, 0x48, 0x27,
	0xfd, 0x59, 0x61, 0xd9, 0xf5, 0xba, 0x85, 0x85, 0x91, 0x09, 0x3d, 0x75, 0xe5, 0x3e, 0x29, 0x9d,
	0xf4, 0x06, 0x3b, 0xf1, 0x15, 0xcc, 0x33, 0x41, 0xc2, 0xd3, 0x22, 0xab, 0x4d, 0x0d, 0x6d, 0xad,
	0xf5, 0x9f, 0x00, 0x2e, 0x3f, 0x09, 0x12, 0xdf, 0xcc, 0xb2, 0xb8, 0x24, 0x89, 0x5d, 0x1c, 0xef,
	0x5e, 0xd0, 0x26, 0xb1, 0x52, 0x37, 0xe0, 0x0e, 0x96, 0x4a, 0xa6, 0xa8, 0xb2, 0x47, 0x17, 0xdc,
	0xd2, 0x3c, 0x71, 0xa7, 0xe7, 0x97, 0xee, 0xe1, 0x3a, 0x95, 0x8a, 0xf2, 0x87, 0x3c, 0x15, 0x94,
	0x63, 0xf5, 0xe8, 0xee, 0xc0, 0xde, 0x7d, 0x7e, 0x26, 0x9d, 0xb5, 0x58, 0xff, 0x80, 0x71, 0x6c,
	0x3b, 0xb3, 0x1b, 0x98, 0x34, 0x09, 0xbc, 0x9c, 0xe0, 0x90, 0x9d, 0x76, 0x05, 0x23, 0x3c, 0x56,
	0x52, 0x35, 0x89, 0x5c, 0x61, 0x56, 0x84, 0xf2, 0x52, 0x6a, 0x12, 0x65, 0x6d, 0xe7, 0x0d, 0xe2,
	0x1e, 0xac, 0xff, 0x06, 0x30, 0xfb, 0xe8, 0x4f, 0x67, 0x6f, 0x60, 0xd1, 0xc7, 0x91, 0x67, 0xdf,
	0xc4, 0xe3, 0x76, 0xe0, 0x12, 0xc6, 0xb9, 0xd6, 0xfb, 0x6e, 0x62, 0x53, 0x99, 0x91, 0x8d, 0x2a,
	0xb3, 0xe6, 0x89, 0x3d, 0x30, 0xfb, 0x61, 0xbd, 0xc4, 0xa4, 0xe0, 0x43, 0x97, 0xc8, 0x92, 0x4d,
	0x5e, 0x4a, 0x76, 0x0b, 0xa1, 0x3c, 0xd5, 0xb9, 0x72, 0x1f, 0xcc, 0x3a, 0x23, 0xeb, 0xcc, 0x7b,
	0xdc, 0x8a, 0x4a, 0x1e, 0x30, 0xf5, 0xc4, 0xb1, 0x13, 0x7b, 0x6c, 0xc4, 0xed, 0xd8, 0xfe, 0x15,
	0xee, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x13, 0xad, 0xfe, 0x22, 0x24, 0x04, 0x00, 0x00,
}
