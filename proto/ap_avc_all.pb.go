// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ap_avc_all.proto

package proto

import (
	fmt "fmt"
	github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type APAVCStats struct {
	//*
	// @property
	// @aggregation
	// @description
	// @since 5.1.1
	Version *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	//*
	// @property
	// @aggregation
	// @description
	// @since 5.1.1
	ArcMessage           []*ArcMessage `protobuf:"bytes,2,rep,name=arc_message,json=arcMessage" json:"arc_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *APAVCStats) Reset()         { *m = APAVCStats{} }
func (m *APAVCStats) String() string { return proto.CompactTextString(m) }
func (*APAVCStats) ProtoMessage()    {}
func (*APAVCStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e53c40df2ef5d7f, []int{0}
}
func (m *APAVCStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *APAVCStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_APAVCStats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *APAVCStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APAVCStats.Merge(m, src)
}
func (m *APAVCStats) XXX_Size() int {
	return m.Size()
}
func (m *APAVCStats) XXX_DiscardUnknown() {
	xxx_messageInfo_APAVCStats.DiscardUnknown(m)
}

var xxx_messageInfo_APAVCStats proto.InternalMessageInfo

func (m *APAVCStats) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *APAVCStats) GetArcMessage() []*ArcMessage {
	if m != nil {
		return m.ArcMessage
	}
	return nil
}

func init() {
	proto.RegisterType((*APAVCStats)(nil), "proto.APAVCStats")
}

func init() { proto.RegisterFile("ap_avc_all.proto", fileDescriptor_6e53c40df2ef5d7f) }

var fileDescriptor_6e53c40df2ef5d7f = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2c, 0x88, 0x4f,
	0x2c, 0x4b, 0x8e, 0x4f, 0xcc, 0xc9, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53,
	0x52, 0x3c, 0x10, 0x09, 0x88, 0xa0, 0x52, 0x14, 0x17, 0x97, 0x63, 0x80, 0x63, 0x98, 0x73, 0x70,
	0x49, 0x62, 0x49, 0xb1, 0x90, 0x04, 0x17, 0x7b, 0x59, 0x6a, 0x51, 0x71, 0x66, 0x7e, 0x9e, 0x04,
	0xa3, 0x02, 0x93, 0x06, 0x6f, 0x10, 0x8c, 0x2b, 0x64, 0xc4, 0xc5, 0x9d, 0x58, 0x94, 0x1c, 0x9f,
	0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x2a, 0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x6d, 0x24, 0x08, 0x31,
	0x44, 0xcf, 0xb1, 0x28, 0xd9, 0x17, 0x22, 0x11, 0xc4, 0x95, 0x08, 0x67, 0x3b, 0x19, 0x9e, 0x78,
	0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x70,
	0xc9, 0x27, 0xe7, 0xe7, 0xea, 0x15, 0x95, 0x26, 0x67, 0x97, 0x16, 0x97, 0x67, 0x16, 0xa5, 0xe6,
	0xa4, 0x16, 0x17, 0xeb, 0x15, 0x27, 0xa7, 0x43, 0x8c, 0x49, 0x2a, 0x4d, 0x03, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xf4, 0xcc, 0x00, 0x9b, 0xb6, 0x00, 0x00, 0x00,
}

func (m *APAVCStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *APAVCStats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *APAVCStats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ArcMessage) > 0 {
		for iNdEx := len(m.ArcMessage) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ArcMessage[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintApAvcAll(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Version == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i = encodeVarintApAvcAll(dAtA, i, uint64(*m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintApAvcAll(dAtA []byte, offset int, v uint64) int {
	offset -= sovApAvcAll(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *APAVCStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != nil {
		n += 1 + sovApAvcAll(uint64(*m.Version))
	}
	if len(m.ArcMessage) > 0 {
		for _, e := range m.ArcMessage {
			l = e.Size()
			n += 1 + l + sovApAvcAll(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApAvcAll(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApAvcAll(x uint64) (n int) {
	return sovApAvcAll(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *APAVCStats) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApAvcAll
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
			return fmt.Errorf("proto: APAVCStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: APAVCStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApAvcAll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Version = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArcMessage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApAvcAll
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
				return ErrInvalidLengthApAvcAll
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApAvcAll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ArcMessage = append(m.ArcMessage, &ArcMessage{})
			if err := m.ArcMessage[len(m.ArcMessage)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApAvcAll(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApAvcAll
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApAvcAll
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApAvcAll(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApAvcAll
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
					return 0, ErrIntOverflowApAvcAll
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
					return 0, ErrIntOverflowApAvcAll
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
				return 0, ErrInvalidLengthApAvcAll
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApAvcAll
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApAvcAll
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApAvcAll        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApAvcAll          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApAvcAll = fmt.Errorf("proto: unexpected end of group")
)
