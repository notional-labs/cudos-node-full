// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: marketplace/collection.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Collection listed for sale in the marketplace
type Collection struct {
	Id              uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	DenomId         string    `protobuf:"bytes,2,opt,name=denomId,proto3" json:"denomId,omitempty"`
	MintRoyalties   []Royalty `protobuf:"bytes,3,rep,name=mintRoyalties,proto3" json:"mintRoyalties"`
	ResaleRoyalties []Royalty `protobuf:"bytes,4,rep,name=resaleRoyalties,proto3" json:"resaleRoyalties"`
	Verified        bool      `protobuf:"varint,5,opt,name=verified,proto3" json:"verified"`
	Owner           string    `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *Collection) Reset()         { *m = Collection{} }
func (m *Collection) String() string { return proto.CompactTextString(m) }
func (*Collection) ProtoMessage()    {}
func (*Collection) Descriptor() ([]byte, []int) {
	return fileDescriptor_2874de85607b5723, []int{0}
}
func (m *Collection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Collection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Collection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Collection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collection.Merge(m, src)
}
func (m *Collection) XXX_Size() int {
	return m.Size()
}
func (m *Collection) XXX_DiscardUnknown() {
	xxx_messageInfo_Collection.DiscardUnknown(m)
}

var xxx_messageInfo_Collection proto.InternalMessageInfo

func (m *Collection) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Collection) GetDenomId() string {
	if m != nil {
		return m.DenomId
	}
	return ""
}

func (m *Collection) GetMintRoyalties() []Royalty {
	if m != nil {
		return m.MintRoyalties
	}
	return nil
}

func (m *Collection) GetResaleRoyalties() []Royalty {
	if m != nil {
		return m.ResaleRoyalties
	}
	return nil
}

func (m *Collection) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

func (m *Collection) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func init() {
	proto.RegisterType((*Collection)(nil), "cudoventures.cudosnode.marketplace.Collection")
}

func init() { proto.RegisterFile("marketplace/collection.proto", fileDescriptor_2874de85607b5723) }

var fileDescriptor_2874de85607b5723 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x33, 0xe9, 0x9f, 0xaf, 0xdf, 0xa8, 0x08, 0x43, 0x91, 0x58, 0x24, 0x0d, 0x5d, 0x05,
	0xc4, 0x04, 0x14, 0x7c, 0x80, 0x74, 0xe5, 0x4e, 0xb2, 0x50, 0xd0, 0x55, 0x9a, 0xb9, 0xd6, 0xc1,
	0x64, 0x6e, 0x99, 0x4c, 0xaa, 0x79, 0x0b, 0x1f, 0xc6, 0x87, 0xe8, 0xb2, 0x4b, 0x57, 0x45, 0xda,
	0x5d, 0x9f, 0x42, 0x9a, 0x68, 0x8d, 0x6e, 0x04, 0x77, 0xe7, 0x5c, 0x38, 0xbf, 0x7b, 0x2e, 0x97,
	0x1e, 0xa5, 0x91, 0x7a, 0x00, 0x3d, 0x49, 0xa2, 0x18, 0xfc, 0x18, 0x93, 0x04, 0x62, 0x2d, 0x50,
	0x7a, 0x13, 0x85, 0x1a, 0xd9, 0x20, 0xce, 0x39, 0x4e, 0x41, 0xea, 0x5c, 0x41, 0xe6, 0x6d, 0x4c,
	0x26, 0x91, 0x83, 0x57, 0x0b, 0xf5, 0x0e, 0xeb, 0x04, 0x85, 0x45, 0x94, 0xe8, 0xa2, 0x8a, 0xf7,
	0xba, 0x63, 0x1c, 0x63, 0x29, 0xfd, 0x8d, 0xaa, 0xa6, 0x83, 0x17, 0x93, 0xd2, 0xe1, 0x76, 0x13,
	0x3b, 0xa0, 0xa6, 0xe0, 0x16, 0x71, 0x88, 0xdb, 0x0c, 0xda, 0xeb, 0x45, 0xdf, 0x14, 0x3c, 0x34,
	0x05, 0x67, 0x16, 0xfd, 0xc7, 0x41, 0x62, 0x7a, 0xc1, 0x2d, 0xd3, 0x21, 0xee, 0xff, 0xf0, 0xd3,
	0xb2, 0x6b, 0xba, 0x97, 0x0a, 0xa9, 0xc3, 0x72, 0x97, 0x80, 0xcc, 0x6a, 0x38, 0x0d, 0x77, 0xe7,
	0xf4, 0xd8, 0xfb, 0xbd, 0xad, 0x57, 0x85, 0x8a, 0xa0, 0x39, 0x5b, 0xf4, 0x8d, 0xf0, 0x3b, 0x87,
	0xdd, 0xd2, 0x7d, 0x05, 0x59, 0x94, 0xc0, 0x17, 0xba, 0xf9, 0x57, 0xf4, 0x4f, 0x12, 0x73, 0x69,
	0x67, 0x0a, 0x4a, 0xdc, 0x09, 0xe0, 0x56, 0xcb, 0x21, 0x6e, 0x27, 0xd8, 0x5d, 0x2f, 0xfa, 0xdb,
	0x59, 0xb8, 0x55, 0xac, 0x4b, 0x5b, 0xf8, 0x28, 0x41, 0x59, 0xed, 0xf2, 0xee, 0xca, 0x04, 0x97,
	0xb3, 0xa5, 0x4d, 0xe6, 0x4b, 0x9b, 0xbc, 0x2d, 0x6d, 0xf2, 0xbc, 0xb2, 0x8d, 0xf9, 0xca, 0x36,
	0x5e, 0x57, 0xb6, 0x71, 0x73, 0x3e, 0x16, 0xfa, 0x3e, 0x1f, 0x79, 0x31, 0xa6, 0xfe, 0x30, 0xe7,
	0x78, 0xf5, 0xd1, 0xd3, 0x2f, 0x7b, 0x9e, 0x6c, 0x8a, 0xfa, 0x4f, 0x7e, 0xfd, 0x4d, 0xba, 0x98,
	0x40, 0x36, 0x6a, 0x97, 0xff, 0x38, 0x7b, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xce, 0xec, 0xbb, 0x9f,
	0x04, 0x02, 0x00, 0x00,
}

func (m *Collection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Collection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Collection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x32
	}
	if m.Verified {
		i--
		if m.Verified {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.ResaleRoyalties) > 0 {
		for iNdEx := len(m.ResaleRoyalties) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ResaleRoyalties[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.MintRoyalties) > 0 {
		for iNdEx := len(m.MintRoyalties) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MintRoyalties[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DenomId) > 0 {
		i -= len(m.DenomId)
		copy(dAtA[i:], m.DenomId)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.DenomId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintCollection(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCollection(dAtA []byte, offset int, v uint64) int {
	offset -= sovCollection(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Collection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovCollection(uint64(m.Id))
	}
	l = len(m.DenomId)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	if len(m.MintRoyalties) > 0 {
		for _, e := range m.MintRoyalties {
			l = e.Size()
			n += 1 + l + sovCollection(uint64(l))
		}
	}
	if len(m.ResaleRoyalties) > 0 {
		for _, e := range m.ResaleRoyalties {
			l = e.Size()
			n += 1 + l + sovCollection(uint64(l))
		}
	}
	if m.Verified {
		n += 2
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	return n
}

func sovCollection(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCollection(x uint64) (n int) {
	return sovCollection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Collection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollection
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Collection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Collection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintRoyalties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintRoyalties = append(m.MintRoyalties, Royalty{})
			if err := m.MintRoyalties[len(m.MintRoyalties)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResaleRoyalties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResaleRoyalties = append(m.ResaleRoyalties, Royalty{})
			if err := m.ResaleRoyalties[len(m.ResaleRoyalties)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Verified", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Verified = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollection
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
func skipCollection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCollection
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
					return 0, ErrIntOverflowCollection
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCollection
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
			if length < 0 {
				return 0, ErrInvalidLengthCollection
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCollection
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCollection
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCollection        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCollection          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCollection = fmt.Errorf("proto: unexpected end of group")
)
