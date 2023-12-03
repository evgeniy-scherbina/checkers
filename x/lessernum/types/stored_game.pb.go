// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lessernum/stored_game.proto

package types

import (
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

type StoredGame struct {
	Index        string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Player1      string `protobuf:"bytes,2,opt,name=player1,proto3" json:"player1,omitempty"`
	Player2      string `protobuf:"bytes,3,opt,name=player2,proto3" json:"player2,omitempty"`
	PlayerToMove uint64 `protobuf:"varint,4,opt,name=playerToMove,proto3" json:"playerToMove,omitempty"`
	Move1        uint64 `protobuf:"varint,5,opt,name=move1,proto3" json:"move1,omitempty"`
	Move2        uint64 `protobuf:"varint,6,opt,name=move2,proto3" json:"move2,omitempty"`
	Winner       string `protobuf:"bytes,7,opt,name=winner,proto3" json:"winner,omitempty"`
	Wager        uint64 `protobuf:"varint,8,opt,name=wager,proto3" json:"wager,omitempty"`
}

func (m *StoredGame) Reset()         { *m = StoredGame{} }
func (m *StoredGame) String() string { return proto.CompactTextString(m) }
func (*StoredGame) ProtoMessage()    {}
func (*StoredGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba9fe6ef1253ec15, []int{0}
}
func (m *StoredGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoredGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoredGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoredGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredGame.Merge(m, src)
}
func (m *StoredGame) XXX_Size() int {
	return m.Size()
}
func (m *StoredGame) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredGame.DiscardUnknown(m)
}

var xxx_messageInfo_StoredGame proto.InternalMessageInfo

func (m *StoredGame) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *StoredGame) GetPlayer1() string {
	if m != nil {
		return m.Player1
	}
	return ""
}

func (m *StoredGame) GetPlayer2() string {
	if m != nil {
		return m.Player2
	}
	return ""
}

func (m *StoredGame) GetPlayerToMove() uint64 {
	if m != nil {
		return m.PlayerToMove
	}
	return 0
}

func (m *StoredGame) GetMove1() uint64 {
	if m != nil {
		return m.Move1
	}
	return 0
}

func (m *StoredGame) GetMove2() uint64 {
	if m != nil {
		return m.Move2
	}
	return 0
}

func (m *StoredGame) GetWinner() string {
	if m != nil {
		return m.Winner
	}
	return ""
}

func (m *StoredGame) GetWager() uint64 {
	if m != nil {
		return m.Wager
	}
	return 0
}

func init() {
	proto.RegisterType((*StoredGame)(nil), "alice.checkers.lessernum.StoredGame")
}

func init() { proto.RegisterFile("lessernum/stored_game.proto", fileDescriptor_ba9fe6ef1253ec15) }

var fileDescriptor_ba9fe6ef1253ec15 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0x49, 0x2d, 0x2e,
	0x4e, 0x2d, 0xca, 0x2b, 0xcd, 0xd5, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4d, 0x89, 0x4f, 0x4f, 0xcc,
	0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x48, 0xcc, 0xc9, 0x4c, 0x4e, 0xd5, 0x4b,
	0xce, 0x48, 0x4d, 0xce, 0x4e, 0x2d, 0x2a, 0xd6, 0x83, 0xab, 0x55, 0xba, 0xc2, 0xc8, 0xc5, 0x15,
	0x0c, 0x56, 0xef, 0x9e, 0x98, 0x9b, 0x2a, 0x24, 0xc2, 0xc5, 0x9a, 0x99, 0x97, 0x92, 0x5a, 0x21,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x08, 0x49, 0x70, 0xb1, 0x17, 0xe4, 0x24, 0x56,
	0xa6, 0x16, 0x19, 0x4a, 0x30, 0x81, 0xc5, 0x61, 0x5c, 0x84, 0x8c, 0x91, 0x04, 0x33, 0xb2, 0x8c,
	0x91, 0x90, 0x12, 0x17, 0x0f, 0x84, 0x19, 0x92, 0xef, 0x9b, 0x5f, 0x96, 0x2a, 0xc1, 0xa2, 0xc0,
	0xa8, 0xc1, 0x12, 0x84, 0x22, 0x06, 0xb2, 0x2d, 0x37, 0xbf, 0x2c, 0xd5, 0x50, 0x82, 0x15, 0x2c,
	0x09, 0xe1, 0xc0, 0x44, 0x8d, 0x24, 0xd8, 0x10, 0xa2, 0x46, 0x42, 0x62, 0x5c, 0x6c, 0xe5, 0x99,
	0x79, 0x79, 0xa9, 0x45, 0x12, 0xec, 0x60, 0x8b, 0xa0, 0x3c, 0x90, 0xea, 0xf2, 0xc4, 0xf4, 0xd4,
	0x22, 0x09, 0x0e, 0x88, 0x6a, 0x30, 0xc7, 0xc9, 0xf5, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4,
	0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f,
	0xe5, 0x18, 0xa2, 0xb4, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xc1,
	0xa1, 0xa2, 0x0f, 0x0b, 0x15, 0xfd, 0x0a, 0x7d, 0x44, 0x18, 0x96, 0x54, 0x16, 0xa4, 0x16, 0x27,
	0xb1, 0x81, 0x83, 0xcf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x12, 0x3e, 0x77, 0x5d, 0x01,
	0x00, 0x00,
}

func (m *StoredGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoredGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoredGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Wager != 0 {
		i = encodeVarintStoredGame(dAtA, i, uint64(m.Wager))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Winner) > 0 {
		i -= len(m.Winner)
		copy(dAtA[i:], m.Winner)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Winner)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Move2 != 0 {
		i = encodeVarintStoredGame(dAtA, i, uint64(m.Move2))
		i--
		dAtA[i] = 0x30
	}
	if m.Move1 != 0 {
		i = encodeVarintStoredGame(dAtA, i, uint64(m.Move1))
		i--
		dAtA[i] = 0x28
	}
	if m.PlayerToMove != 0 {
		i = encodeVarintStoredGame(dAtA, i, uint64(m.PlayerToMove))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Player2) > 0 {
		i -= len(m.Player2)
		copy(dAtA[i:], m.Player2)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Player2)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Player1) > 0 {
		i -= len(m.Player1)
		copy(dAtA[i:], m.Player1)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Player1)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStoredGame(dAtA []byte, offset int, v uint64) int {
	offset -= sovStoredGame(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoredGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Player1)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Player2)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	if m.PlayerToMove != 0 {
		n += 1 + sovStoredGame(uint64(m.PlayerToMove))
	}
	if m.Move1 != 0 {
		n += 1 + sovStoredGame(uint64(m.Move1))
	}
	if m.Move2 != 0 {
		n += 1 + sovStoredGame(uint64(m.Move2))
	}
	l = len(m.Winner)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	if m.Wager != 0 {
		n += 1 + sovStoredGame(uint64(m.Wager))
	}
	return n
}

func sovStoredGame(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStoredGame(x uint64) (n int) {
	return sovStoredGame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoredGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStoredGame
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
			return fmt.Errorf("proto: StoredGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoredGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Player1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Player1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Player2", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Player2 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerToMove", wireType)
			}
			m.PlayerToMove = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlayerToMove |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Move1", wireType)
			}
			m.Move1 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Move1 |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Move2", wireType)
			}
			m.Move2 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Move2 |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Winner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Winner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wager", wireType)
			}
			m.Wager = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Wager |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStoredGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStoredGame
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
func skipStoredGame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStoredGame
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
					return 0, ErrIntOverflowStoredGame
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
					return 0, ErrIntOverflowStoredGame
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
				return 0, ErrInvalidLengthStoredGame
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStoredGame
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStoredGame
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStoredGame        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStoredGame          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStoredGame = fmt.Errorf("proto: unexpected end of group")
)
