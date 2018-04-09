// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usage.proto

package corepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Usage struct {
	Timestamps []*TxTimestamp `protobuf:"bytes,1,rep,name=timestamps" json:"timestamps,omitempty"`
}

func (m *Usage) Reset()                    { *m = Usage{} }
func (m *Usage) String() string            { return proto.CompactTextString(m) }
func (*Usage) ProtoMessage()               {}
func (*Usage) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Usage) GetTimestamps() []*TxTimestamp {
	if m != nil {
		return m.Timestamps
	}
	return nil
}

type TxTimestamp struct {
	Hash      []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *TxTimestamp) Reset()                    { *m = TxTimestamp{} }
func (m *TxTimestamp) String() string            { return proto.CompactTextString(m) }
func (*TxTimestamp) ProtoMessage()               {}
func (*TxTimestamp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *TxTimestamp) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *TxTimestamp) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Usage)(nil), "corepb.Usage")
	proto.RegisterType((*TxTimestamp)(nil), "corepb.TxTimestamp")
}

func init() { proto.RegisterFile("usage.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2d, 0x4e, 0x4c,
	0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xce, 0x2f, 0x4a, 0x2d, 0x48, 0x52,
	0xb2, 0xe1, 0x62, 0x0d, 0x05, 0x09, 0x0b, 0x19, 0x73, 0x71, 0x95, 0x64, 0xe6, 0xa6, 0x16, 0x97,
	0x24, 0xe6, 0x16, 0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x09, 0xeb, 0x41, 0x54, 0xe9,
	0x85, 0x54, 0x84, 0xc0, 0xe4, 0x82, 0x90, 0x94, 0x29, 0xd9, 0x73, 0x71, 0x23, 0x49, 0x09, 0x09,
	0x71, 0xb1, 0x64, 0x24, 0x16, 0x67, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x81, 0xd9, 0x42,
	0x32, 0x5c, 0x9c, 0x70, 0x0d, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x08, 0x81, 0x24, 0x36,
	0xb0, 0x6b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x22, 0x9d, 0x3e, 0x62, 0x9c, 0x00, 0x00,
	0x00,
}
