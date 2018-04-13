// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: record.proto

package corepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Reader struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	EncKey  []byte `protobuf:"bytes,2,opt,name=enc_key,json=encKey,proto3" json:"enc_key,omitempty"`
	Seed    []byte `protobuf:"bytes,3,opt,name=seed,proto3" json:"seed,omitempty"`
}

func (m *Reader) Reset()                    { *m = Reader{} }
func (m *Reader) String() string            { return proto.CompactTextString(m) }
func (*Reader) ProtoMessage()               {}
func (*Reader) Descriptor() ([]byte, []int) { return fileDescriptorRecord, []int{0} }

func (m *Reader) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Reader) GetEncKey() []byte {
	if m != nil {
		return m.EncKey
	}
	return nil
}

func (m *Reader) GetSeed() []byte {
	if m != nil {
		return m.Seed
	}
	return nil
}

type Record struct {
	Hash      []byte    `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Storage   string    `protobuf:"bytes,2,opt,name=storage,proto3" json:"storage,omitempty"`
	Owner     []byte    `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Timestamp int64     `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Readers   []*Reader `protobuf:"bytes,5,rep,name=readers" json:"readers,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptorRecord, []int{1} }

func (m *Record) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Record) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
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

func (m *Record) GetReaders() []*Reader {
	if m != nil {
		return m.Readers
	}
	return nil
}

func init() {
	proto.RegisterType((*Reader)(nil), "corepb.Reader")
	proto.RegisterType((*Record)(nil), "corepb.Record")
}
func (m *Reader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Reader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if len(m.EncKey) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRecord(dAtA, i, uint64(len(m.EncKey)))
		i += copy(dAtA[i:], m.EncKey)
	}
	if len(m.Seed) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Seed)))
		i += copy(dAtA[i:], m.Seed)
	}
	return i, nil
}

func (m *Record) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Hash)))
		i += copy(dAtA[i:], m.Hash)
	}
	if len(m.Storage) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Storage)))
		i += copy(dAtA[i:], m.Storage)
	}
	if len(m.Owner) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Owner)))
		i += copy(dAtA[i:], m.Owner)
	}
	if m.Timestamp != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintRecord(dAtA, i, uint64(m.Timestamp))
	}
	if len(m.Readers) > 0 {
		for _, msg := range m.Readers {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintRecord(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintRecord(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Reader) Size() (n int) {
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	l = len(m.EncKey)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	l = len(m.Seed)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	return n
}

func (m *Record) Size() (n int) {
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	l = len(m.Storage)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovRecord(uint64(m.Timestamp))
	}
	if len(m.Readers) > 0 {
		for _, e := range m.Readers {
			l = e.Size()
			n += 1 + l + sovRecord(uint64(l))
		}
	}
	return n
}

func sovRecord(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRecord(x uint64) (n int) {
	return sovRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Reader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecord
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Reader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Reader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncKey = append(m.EncKey[:0], dAtA[iNdEx:postIndex]...)
			if m.EncKey == nil {
				m.EncKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seed", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seed = append(m.Seed[:0], dAtA[iNdEx:postIndex]...)
			if m.Seed == nil {
				m.Seed = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRecord
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Record) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecord
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Readers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Readers = append(m.Readers, &Reader{})
			if err := m.Readers[len(m.Readers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRecord
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRecord
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRecord
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRecord
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRecord(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRecord = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRecord   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("record.proto", fileDescriptorRecord) }

var fileDescriptorRecord = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x89, 0x9d, 0x49, 0x99, 0x67, 0x71, 0xf1, 0x10, 0xcc, 0x42, 0x4a, 0x99, 0x55, 0x56,
	0x5d, 0xe8, 0x0d, 0xdc, 0xba, 0x10, 0x72, 0x01, 0xc9, 0x24, 0x0f, 0x47, 0x64, 0x9a, 0xf2, 0x12,
	0x90, 0xde, 0xc3, 0x43, 0xb9, 0xf4, 0x08, 0xd2, 0x93, 0x48, 0xd2, 0x16, 0x77, 0xef, 0xff, 0x7f,
	0xf8, 0xfe, 0xfc, 0x81, 0x86, 0xc9, 0x05, 0xf6, 0xfd, 0xc8, 0x21, 0x05, 0x94, 0x2e, 0x30, 0x8d,
	0xa7, 0xe3, 0x0b, 0x48, 0x43, 0xd6, 0x13, 0xa3, 0x82, 0xda, 0x7a, 0xcf, 0x14, 0xa3, 0x12, 0x9d,
	0xd0, 0x8d, 0xd9, 0x24, 0xde, 0x41, 0x4d, 0x83, 0x7b, 0xfd, 0xa0, 0x49, 0x5d, 0x95, 0x44, 0xd2,
	0xe0, 0x9e, 0x69, 0x42, 0x84, 0x5d, 0x24, 0xf2, 0xaa, 0x2a, 0x6e, 0xb9, 0x8f, 0x5f, 0x22, 0x13,
	0x73, 0x53, 0x8e, 0xcf, 0x36, 0x9e, 0x57, 0x5c, 0xb9, 0x73, 0x4b, 0x4c, 0x81, 0xed, 0x1b, 0x15,
	0xd6, 0xc1, 0x6c, 0x12, 0x6f, 0x61, 0x1f, 0x3e, 0x07, 0xe2, 0x95, 0xb6, 0x08, 0xbc, 0x87, 0x43,
	0x7a, 0xbf, 0x50, 0x4c, 0xf6, 0x32, 0xaa, 0x5d, 0x27, 0x74, 0x65, 0xfe, 0x0d, 0xd4, 0x50, 0x73,
	0x79, 0x7d, 0x54, 0xfb, 0xae, 0xd2, 0xd7, 0x0f, 0x37, 0xfd, 0xb2, 0xab, 0x5f, 0x46, 0x99, 0x2d,
	0x7e, 0x6a, 0xbe, 0xe7, 0x56, 0xfc, 0xcc, 0xad, 0xf8, 0x9d, 0x5b, 0x71, 0x92, 0xe5, 0x13, 0x1e,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x70, 0x17, 0xb8, 0x7b, 0x14, 0x01, 0x00, 0x00,
}
