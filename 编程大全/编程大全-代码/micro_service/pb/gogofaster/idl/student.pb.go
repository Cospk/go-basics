// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: student.proto

package student_service

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
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

type Student struct {
	Name      string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt int64              `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Locations []string           `protobuf:"bytes,4,rep,name=Locations,proto3" json:"Locations,omitempty"`
	Scores    map[string]float32 `protobuf:"bytes,3,rep,name=Scores,proto3" json:"Scores,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	Gender    bool               `protobuf:"varint,5,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Age       int32              `protobuf:"varint,6,opt,name=Age,proto3" json:"Age,omitempty"` // Deprecated: Do not use.
	Height    float32            `protobuf:"fixed32,7,opt,name=Height,proto3" json:"Height,omitempty"`
}

func (m *Student) Reset()         { *m = Student{} }
func (m *Student) String() string { return proto.CompactTextString(m) }
func (*Student) ProtoMessage()    {}
func (*Student) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{0}
}
func (m *Student) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Student) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Student.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Student) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Student.Merge(m, src)
}
func (m *Student) XXX_Size() int {
	return m.Size()
}
func (m *Student) XXX_DiscardUnknown() {
	xxx_messageInfo_Student.DiscardUnknown(m)
}

var xxx_messageInfo_Student proto.InternalMessageInfo

func (m *Student) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Student) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Student) GetLocations() []string {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *Student) GetScores() map[string]float32 {
	if m != nil {
		return m.Scores
	}
	return nil
}

func (m *Student) GetGender() bool {
	if m != nil {
		return m.Gender
	}
	return false
}

// Deprecated: Do not use.
func (m *Student) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Student) GetHeight() float32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*Student)(nil), "Student")
	proto.RegisterMapType((map[string]float32)(nil), "Student.ScoresEntry")
}

func init() { proto.RegisterFile("student.proto", fileDescriptor_94a1c1b032ad0c00) }

var fileDescriptor_94a1c1b032ad0c00 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x97, 0x64, 0xeb, 0xec, 0x3b, 0x04, 0x09, 0x53, 0x82, 0x68, 0x08, 0x9e, 0x72, 0x90,
	0x0e, 0xf4, 0xe2, 0x9f, 0xd3, 0x06, 0xa2, 0x07, 0x4f, 0xd9, 0xcd, 0xcb, 0xa8, 0xed, 0xcb, 0x2c,
	0xce, 0x56, 0xd2, 0x6c, 0xb0, 0x8f, 0xe0, 0xcd, 0x8f, 0xe5, 0x71, 0x47, 0x8f, 0xd2, 0x7e, 0x11,
	0x69, 0x1b, 0xd0, 0xdb, 0xfb, 0xfb, 0xf1, 0x90, 0xf0, 0x3c, 0xb0, 0x5f, 0xba, 0x75, 0x8a, 0xb9,
	0x8b, 0xde, 0x6d, 0xe1, 0x8a, 0xb3, 0x0f, 0x0a, 0xc3, 0x79, 0x67, 0x38, 0x87, 0x7e, 0x1e, 0xbf,
	0xa1, 0x20, 0x8a, 0xe8, 0xd0, 0xb4, 0x37, 0x3f, 0x05, 0x48, 0x2c, 0xc6, 0x0e, 0xd3, 0x45, 0xec,
	0x04, 0x55, 0x44, 0x33, 0x13, 0x7a, 0x33, 0x75, 0xfc, 0x04, 0xc2, 0xc7, 0x22, 0x89, 0x5d, 0x56,
	0xe4, 0xa5, 0xe8, 0x2b, 0xa6, 0x43, 0xf3, 0x27, 0xf8, 0x39, 0x04, 0xf3, 0xa4, 0xb0, 0x58, 0x0a,
	0xa6, 0x98, 0x1e, 0x5d, 0x8c, 0x23, 0xff, 0x55, 0xd4, 0xe9, 0xbb, 0xdc, 0xd9, 0xad, 0xf1, 0x19,
	0x7e, 0x04, 0xc1, 0x3d, 0xe6, 0x29, 0x5a, 0x31, 0x50, 0x44, 0xef, 0x19, 0x4f, 0x7c, 0x0c, 0x6c,
	0xba, 0x44, 0x11, 0x28, 0xa2, 0x07, 0x33, 0x2a, 0x88, 0x69, 0xb0, 0x49, 0x3f, 0x60, 0xb6, 0x7c,
	0x71, 0x62, 0xa8, 0x88, 0xa6, 0xc6, 0xd3, 0xf1, 0x35, 0x8c, 0xfe, 0x3d, 0xce, 0x0f, 0x80, 0xbd,
	0xe2, 0xd6, 0x57, 0x6a, 0x4e, 0x3e, 0x86, 0xc1, 0x26, 0x5e, 0xad, 0xb1, 0x2d, 0x43, 0x4d, 0x07,
	0x37, 0xf4, 0x8a, 0xcc, 0x26, 0x5f, 0x95, 0x24, 0xbb, 0x4a, 0x92, 0x9f, 0x4a, 0x92, 0xcf, 0x5a,
	0xf6, 0x76, 0xb5, 0xec, 0x7d, 0xd7, 0xb2, 0xf7, 0x74, 0x18, 0x4d, 0xb2, 0x74, 0x75, 0xeb, 0xa7,
	0x5b, 0x94, 0x68, 0x37, 0x59, 0x82, 0xcf, 0x41, 0xbb, 0xe1, 0xe5, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x4f, 0xbb, 0x85, 0xe2, 0x54, 0x01, 0x00, 0x00,
}

func (m *Student) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Student) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Student) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Height))))
		i--
		dAtA[i] = 0x3d
	}
	if m.Age != 0 {
		i = encodeVarintStudent(dAtA, i, uint64(m.Age))
		i--
		dAtA[i] = 0x30
	}
	if m.Gender {
		i--
		if m.Gender {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Locations) > 0 {
		for iNdEx := len(m.Locations) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Locations[iNdEx])
			copy(dAtA[i:], m.Locations[iNdEx])
			i = encodeVarintStudent(dAtA, i, uint64(len(m.Locations[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Scores) > 0 {
		for k := range m.Scores {
			v := m.Scores[k]
			baseI := i
			i -= 4
			encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(v))))
			i--
			dAtA[i] = 0x15
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintStudent(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintStudent(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.CreatedAt != 0 {
		i = encodeVarintStudent(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintStudent(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStudent(dAtA []byte, offset int, v uint64) int {
	offset -= sovStudent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Student) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovStudent(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovStudent(uint64(m.CreatedAt))
	}
	if len(m.Scores) > 0 {
		for k, v := range m.Scores {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovStudent(uint64(len(k))) + 1 + 4
			n += mapEntrySize + 1 + sovStudent(uint64(mapEntrySize))
		}
	}
	if len(m.Locations) > 0 {
		for _, s := range m.Locations {
			l = len(s)
			n += 1 + l + sovStudent(uint64(l))
		}
	}
	if m.Gender {
		n += 2
	}
	if m.Age != 0 {
		n += 1 + sovStudent(uint64(m.Age))
	}
	if m.Height != 0 {
		n += 5
	}
	return n
}

func sovStudent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStudent(x uint64) (n int) {
	return sovStudent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Student) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStudent
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
			return fmt.Errorf("proto: Student: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Student: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStudent
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
				return ErrInvalidLengthStudent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStudent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStudent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scores", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStudent
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
				return ErrInvalidLengthStudent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStudent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Scores == nil {
				m.Scores = make(map[string]float32)
			}
			var mapkey string
			var mapvalue float32
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStudent
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowStudent
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthStudent
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthStudent
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapvaluetemp uint32
					if (iNdEx + 4) > l {
						return io.ErrUnexpectedEOF
					}
					mapvaluetemp = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
					iNdEx += 4
					mapvalue = math.Float32frombits(mapvaluetemp)
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipStudent(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthStudent
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Scores[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locations", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStudent
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
				return ErrInvalidLengthStudent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStudent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Locations = append(m.Locations, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gender", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStudent
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
			m.Gender = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStudent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Height = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipStudent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStudent
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
func skipStudent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStudent
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
					return 0, ErrIntOverflowStudent
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
					return 0, ErrIntOverflowStudent
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
				return 0, ErrInvalidLengthStudent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStudent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStudent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStudent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStudent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStudent = fmt.Errorf("proto: unexpected end of group")
)
