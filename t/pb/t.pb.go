// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: t.proto

/*
Package corepb is a generated protocol buffer package.

It is generated from these files:
	t.proto

It has these top-level messages:
	TransactionHashTarget
	DefaultPayload
	VotePayload
	AddCertificationPayload
	RevokeCertificationPayload
	AddRecordPayload
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

type TransactionHashTarget struct {
	TxType    string `protobuf:"bytes,1,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
	From      []byte `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To        []byte `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Value     []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Nonce     uint64 `protobuf:"varint,6,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ChainId   uint32 `protobuf:"varint,7,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*TransactionHashTarget_DefaultPayload
	//	*TransactionHashTarget_AddCertificationPayload
	//	*TransactionHashTarget_AddRecordPayload
	//	*TransactionHashTarget_RevokeCertificationPayload
	//	*TransactionHashTarget_VotePayload
	Payload isTransactionHashTarget_Payload `protobuf_oneof:"payload"`
}

func (m *TransactionHashTarget) Reset()                    { *m = TransactionHashTarget{} }
func (m *TransactionHashTarget) String() string            { return proto.CompactTextString(m) }
func (*TransactionHashTarget) ProtoMessage()               {}
func (*TransactionHashTarget) Descriptor() ([]byte, []int) { return fileDescriptorT, []int{0} }

type isTransactionHashTarget_Payload interface {
	isTransactionHashTarget_Payload()
}

type TransactionHashTarget_DefaultPayload struct {
	DefaultPayload *DefaultPayload `protobuf:"bytes,20,opt,name=default_payload,json=defaultPayload,oneof"`
}
type TransactionHashTarget_AddCertificationPayload struct {
	AddCertificationPayload *AddCertificationPayload `protobuf:"bytes,21,opt,name=add_certification_payload,json=addCertificationPayload,oneof"`
}
type TransactionHashTarget_AddRecordPayload struct {
	AddRecordPayload *AddRecordPayload `protobuf:"bytes,22,opt,name=add_record_payload,json=addRecordPayload,oneof"`
}
type TransactionHashTarget_RevokeCertificationPayload struct {
	RevokeCertificationPayload *RevokeCertificationPayload `protobuf:"bytes,23,opt,name=revoke_certification_payload,json=revokeCertificationPayload,oneof"`
}
type TransactionHashTarget_VotePayload struct {
	VotePayload *VotePayload `protobuf:"bytes,24,opt,name=vote_payload,json=votePayload,oneof"`
}

func (*TransactionHashTarget_DefaultPayload) isTransactionHashTarget_Payload()             {}
func (*TransactionHashTarget_AddCertificationPayload) isTransactionHashTarget_Payload()    {}
func (*TransactionHashTarget_AddRecordPayload) isTransactionHashTarget_Payload()           {}
func (*TransactionHashTarget_RevokeCertificationPayload) isTransactionHashTarget_Payload() {}
func (*TransactionHashTarget_VotePayload) isTransactionHashTarget_Payload()                {}

func (m *TransactionHashTarget) GetPayload() isTransactionHashTarget_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TransactionHashTarget) GetTxType() string {
	if m != nil {
		return m.TxType
	}
	return ""
}

func (m *TransactionHashTarget) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *TransactionHashTarget) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *TransactionHashTarget) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *TransactionHashTarget) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TransactionHashTarget) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *TransactionHashTarget) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *TransactionHashTarget) GetDefaultPayload() *DefaultPayload {
	if x, ok := m.GetPayload().(*TransactionHashTarget_DefaultPayload); ok {
		return x.DefaultPayload
	}
	return nil
}

func (m *TransactionHashTarget) GetAddCertificationPayload() *AddCertificationPayload {
	if x, ok := m.GetPayload().(*TransactionHashTarget_AddCertificationPayload); ok {
		return x.AddCertificationPayload
	}
	return nil
}

func (m *TransactionHashTarget) GetAddRecordPayload() *AddRecordPayload {
	if x, ok := m.GetPayload().(*TransactionHashTarget_AddRecordPayload); ok {
		return x.AddRecordPayload
	}
	return nil
}

func (m *TransactionHashTarget) GetRevokeCertificationPayload() *RevokeCertificationPayload {
	if x, ok := m.GetPayload().(*TransactionHashTarget_RevokeCertificationPayload); ok {
		return x.RevokeCertificationPayload
	}
	return nil
}

func (m *TransactionHashTarget) GetVotePayload() *VotePayload {
	if x, ok := m.GetPayload().(*TransactionHashTarget_VotePayload); ok {
		return x.VotePayload
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TransactionHashTarget) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TransactionHashTarget_OneofMarshaler, _TransactionHashTarget_OneofUnmarshaler, _TransactionHashTarget_OneofSizer, []interface{}{
		(*TransactionHashTarget_DefaultPayload)(nil),
		(*TransactionHashTarget_AddCertificationPayload)(nil),
		(*TransactionHashTarget_AddRecordPayload)(nil),
		(*TransactionHashTarget_RevokeCertificationPayload)(nil),
		(*TransactionHashTarget_VotePayload)(nil),
	}
}

func _TransactionHashTarget_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TransactionHashTarget)
	// payload
	switch x := m.Payload.(type) {
	case *TransactionHashTarget_DefaultPayload:
		_ = b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DefaultPayload); err != nil {
			return err
		}
	case *TransactionHashTarget_AddCertificationPayload:
		_ = b.EncodeVarint(21<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AddCertificationPayload); err != nil {
			return err
		}
	case *TransactionHashTarget_AddRecordPayload:
		_ = b.EncodeVarint(22<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AddRecordPayload); err != nil {
			return err
		}
	case *TransactionHashTarget_RevokeCertificationPayload:
		_ = b.EncodeVarint(23<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RevokeCertificationPayload); err != nil {
			return err
		}
	case *TransactionHashTarget_VotePayload:
		_ = b.EncodeVarint(24<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VotePayload); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TransactionHashTarget.Payload has unexpected type %T", x)
	}
	return nil
}

func _TransactionHashTarget_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TransactionHashTarget)
	switch tag {
	case 20: // payload.default_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DefaultPayload)
		err := b.DecodeMessage(msg)
		m.Payload = &TransactionHashTarget_DefaultPayload{msg}
		return true, err
	case 21: // payload.add_certification_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AddCertificationPayload)
		err := b.DecodeMessage(msg)
		m.Payload = &TransactionHashTarget_AddCertificationPayload{msg}
		return true, err
	case 22: // payload.add_record_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AddRecordPayload)
		err := b.DecodeMessage(msg)
		m.Payload = &TransactionHashTarget_AddRecordPayload{msg}
		return true, err
	case 23: // payload.revoke_certification_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RevokeCertificationPayload)
		err := b.DecodeMessage(msg)
		m.Payload = &TransactionHashTarget_RevokeCertificationPayload{msg}
		return true, err
	case 24: // payload.vote_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(VotePayload)
		err := b.DecodeMessage(msg)
		m.Payload = &TransactionHashTarget_VotePayload{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TransactionHashTarget_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TransactionHashTarget)
	// payload
	switch x := m.Payload.(type) {
	case *TransactionHashTarget_DefaultPayload:
		s := proto.Size(x.DefaultPayload)
		n += proto.SizeVarint(20<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TransactionHashTarget_AddCertificationPayload:
		s := proto.Size(x.AddCertificationPayload)
		n += proto.SizeVarint(21<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TransactionHashTarget_AddRecordPayload:
		s := proto.Size(x.AddRecordPayload)
		n += proto.SizeVarint(22<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TransactionHashTarget_RevokeCertificationPayload:
		s := proto.Size(x.RevokeCertificationPayload)
		n += proto.SizeVarint(23<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TransactionHashTarget_VotePayload:
		s := proto.Size(x.VotePayload)
		n += proto.SizeVarint(24<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DefaultPayload struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *DefaultPayload) Reset()                    { *m = DefaultPayload{} }
func (m *DefaultPayload) String() string            { return proto.CompactTextString(m) }
func (*DefaultPayload) ProtoMessage()               {}
func (*DefaultPayload) Descriptor() ([]byte, []int) { return fileDescriptorT, []int{1} }

func (m *DefaultPayload) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type VotePayload struct {
	Candidates [][]byte `protobuf:"bytes,1,rep,name=candidates" json:"candidates,omitempty"`
}

func (m *VotePayload) Reset()                    { *m = VotePayload{} }
func (m *VotePayload) String() string            { return proto.CompactTextString(m) }
func (*VotePayload) ProtoMessage()               {}
func (*VotePayload) Descriptor() ([]byte, []int) { return fileDescriptorT, []int{2} }

func (m *VotePayload) GetCandidates() [][]byte {
	if m != nil {
		return m.Candidates
	}
	return nil
}

type AddCertificationPayload struct {
	IssueTime      int64  `protobuf:"varint,1,opt,name=issue_time,json=issueTime,proto3" json:"issue_time,omitempty"`
	ExpirationTime int64  `protobuf:"varint,2,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	Hash           []byte `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *AddCertificationPayload) Reset()                    { *m = AddCertificationPayload{} }
func (m *AddCertificationPayload) String() string            { return proto.CompactTextString(m) }
func (*AddCertificationPayload) ProtoMessage()               {}
func (*AddCertificationPayload) Descriptor() ([]byte, []int) { return fileDescriptorT, []int{3} }

func (m *AddCertificationPayload) GetIssueTime() int64 {
	if m != nil {
		return m.IssueTime
	}
	return 0
}

func (m *AddCertificationPayload) GetExpirationTime() int64 {
	if m != nil {
		return m.ExpirationTime
	}
	return 0
}

func (m *AddCertificationPayload) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type RevokeCertificationPayload struct {
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *RevokeCertificationPayload) Reset()                    { *m = RevokeCertificationPayload{} }
func (m *RevokeCertificationPayload) String() string            { return proto.CompactTextString(m) }
func (*RevokeCertificationPayload) ProtoMessage()               {}
func (*RevokeCertificationPayload) Descriptor() ([]byte, []int) { return fileDescriptorT, []int{4} }

func (m *RevokeCertificationPayload) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type AddRecordPayload struct {
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *AddRecordPayload) Reset()                    { *m = AddRecordPayload{} }
func (m *AddRecordPayload) String() string            { return proto.CompactTextString(m) }
func (*AddRecordPayload) ProtoMessage()               {}
func (*AddRecordPayload) Descriptor() ([]byte, []int) { return fileDescriptorT, []int{5} }

func (m *AddRecordPayload) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func init() {
	proto.RegisterType((*TransactionHashTarget)(nil), "corepb.TransactionHashTarget")
	proto.RegisterType((*DefaultPayload)(nil), "corepb.DefaultPayload")
	proto.RegisterType((*VotePayload)(nil), "corepb.VotePayload")
	proto.RegisterType((*AddCertificationPayload)(nil), "corepb.AddCertificationPayload")
	proto.RegisterType((*RevokeCertificationPayload)(nil), "corepb.RevokeCertificationPayload")
	proto.RegisterType((*AddRecordPayload)(nil), "corepb.AddRecordPayload")
}

func init() { proto.RegisterFile("t.proto", fileDescriptorT) }

var fileDescriptorT = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x41, 0x6f, 0xd3, 0x4e,
	0x10, 0xc5, 0xff, 0x4e, 0xdc, 0xf8, 0x9f, 0x49, 0x48, 0xab, 0xa1, 0x6d, 0xb6, 0x55, 0x01, 0xcb,
	0x07, 0xb0, 0x90, 0x88, 0x10, 0x5c, 0xb8, 0x16, 0x38, 0x84, 0x1b, 0x5a, 0x45, 0xdc, 0x90, 0xb5,
	0xf5, 0x4e, 0x9a, 0x15, 0xb1, 0xd7, 0x5a, 0x6f, 0xa2, 0xe4, 0x6b, 0xf2, 0x89, 0x50, 0xd6, 0x75,
	0xec, 0x56, 0xf5, 0x6d, 0xdf, 0x9b, 0x37, 0xbf, 0x6c, 0x66, 0xc7, 0x10, 0xd8, 0x59, 0x61, 0xb4,
	0xd5, 0x38, 0x48, 0xb5, 0xa1, 0xe2, 0x2e, 0xfa, 0xeb, 0xc3, 0xc5, 0xc2, 0x88, 0xbc, 0x14, 0xa9,
	0x55, 0x3a, 0x9f, 0x8b, 0x72, 0xb5, 0x10, 0xe6, 0x9e, 0x2c, 0x4e, 0x21, 0xb0, 0xbb, 0xc4, 0xee,
	0x0b, 0x62, 0x5e, 0xe8, 0xc5, 0x43, 0x3e, 0xb0, 0xbb, 0xc5, 0xbe, 0x20, 0x44, 0xf0, 0x97, 0x46,
	0x67, 0xac, 0x17, 0x7a, 0xf1, 0x98, 0xbb, 0x33, 0x4e, 0xa0, 0x67, 0x35, 0xeb, 0x3b, 0xa7, 0x67,
	0x35, 0x9e, 0xc3, 0xc9, 0x56, 0xac, 0x37, 0xc4, 0x7c, 0x67, 0x55, 0x02, 0x6f, 0x60, 0x68, 0x55,
	0x46, 0xa5, 0x15, 0x59, 0xc1, 0x4e, 0x42, 0x2f, 0xee, 0xf3, 0xc6, 0x38, 0xf4, 0xe4, 0x3a, 0x4f,
	0x89, 0x0d, 0x42, 0x2f, 0xf6, 0x79, 0x25, 0xf0, 0x0a, 0xfe, 0x4f, 0x57, 0x42, 0xe5, 0x89, 0x92,
	0x2c, 0x08, 0xbd, 0xf8, 0x05, 0x0f, 0x9c, 0xfe, 0x21, 0xf1, 0x16, 0x4e, 0x25, 0x2d, 0xc5, 0x66,
	0x6d, 0x93, 0x42, 0xec, 0xd7, 0x5a, 0x48, 0x76, 0x1e, 0x7a, 0xf1, 0xe8, 0xd3, 0xe5, 0xac, 0xfa,
	0x77, 0xb3, 0xef, 0x55, 0xf9, 0x67, 0x55, 0x9d, 0xff, 0xc7, 0x27, 0xf2, 0x91, 0x83, 0xbf, 0xe1,
	0x4a, 0x48, 0x99, 0xa4, 0x64, 0xac, 0x5a, 0xaa, 0x54, 0x1c, 0x66, 0x70, 0x84, 0x5d, 0x38, 0xd8,
	0x9b, 0x1a, 0x76, 0x2b, 0xe5, 0xb7, 0x76, 0xae, 0xa1, 0x4e, 0xc5, 0xf3, 0x25, 0x9c, 0x03, 0x1e,
	0xf0, 0x86, 0x52, 0x6d, 0xe4, 0x91, 0x7b, 0xe9, 0xb8, 0xac, 0xc5, 0xe5, 0x2e, 0xd0, 0x00, 0xcf,
	0xc4, 0x13, 0x0f, 0x97, 0x70, 0x63, 0x68, 0xab, 0xff, 0x50, 0xc7, 0x5d, 0xa7, 0x8e, 0x19, 0xd5,
	0x4c, 0xee, 0xb2, 0x1d, 0xd7, 0xbd, 0x36, 0x9d, 0x55, 0xfc, 0x02, 0xe3, 0xad, 0xb6, 0x74, 0xe4,
	0x32, 0xc7, 0x7d, 0x59, 0x73, 0x7f, 0x69, 0x4b, 0x0d, 0x68, 0xb4, 0x6d, 0xe4, 0xd7, 0x21, 0x04,
	0x0f, 0x4d, 0xd1, 0x7b, 0x98, 0x3c, 0x9e, 0x3c, 0x32, 0x08, 0x32, 0x2a, 0x4b, 0x71, 0x5f, 0x2f,
	0x53, 0x2d, 0xa3, 0x0f, 0x30, 0x6a, 0x41, 0xf1, 0x35, 0x40, 0x2a, 0x72, 0xa9, 0xa4, 0xb0, 0x54,
	0x32, 0x2f, 0xec, 0xc7, 0x63, 0xde, 0x72, 0xa2, 0x0d, 0x4c, 0x3b, 0xde, 0x01, 0x5f, 0x01, 0xa8,
	0xb2, 0xdc, 0x50, 0x72, 0x58, 0x29, 0xf7, 0x33, 0x7d, 0x3e, 0x74, 0xce, 0x42, 0x65, 0x84, 0xef,
	0xe0, 0x94, 0x76, 0x85, 0x32, 0xd5, 0xdc, 0x5c, 0xa6, 0xe7, 0x32, 0x93, 0xc6, 0x76, 0x41, 0x04,
	0x7f, 0x25, 0xca, 0xd5, 0xc3, 0x36, 0xbb, 0x73, 0xf4, 0x11, 0xae, 0xbb, 0x47, 0x7a, 0xec, 0xf0,
	0x5a, 0x1d, 0x6f, 0xe1, 0xec, 0xe9, 0xc3, 0x3e, 0x97, 0xbb, 0x1b, 0xb8, 0xef, 0xf1, 0xf3, 0xbf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0x40, 0x4b, 0x1d, 0x9a, 0x03, 0x00, 0x00,
}
