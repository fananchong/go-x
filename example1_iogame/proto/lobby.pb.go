// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lobby.proto

/*
	Package proto is a generated protocol buffer package.

	It is generated from these files:
		lobby.proto

	It has these top-level messages:
		EnumCreatePlayer
		MsgCreatePlayer
		MsgCreatePlayerResult
		EnumPlayerBaseInfo
		MsgPlayerBaseInfo
		MsgPlayerBaseInfoResult
*/
package proto

import proto1 "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MsgTypeCmd_Lobby int32

const (
	MsgTypeCmd_Lobby_UNSPECIFIED    MsgTypeCmd_Lobby = 0
	MsgTypeCmd_Lobby_CreatePlayer   MsgTypeCmd_Lobby = 10001
	MsgTypeCmd_Lobby_PlayerBaseInfo MsgTypeCmd_Lobby = 10002
)

var MsgTypeCmd_Lobby_name = map[int32]string{
	0:     "UNSPECIFIED",
	10001: "CreatePlayer",
	10002: "PlayerBaseInfo",
}
var MsgTypeCmd_Lobby_value = map[string]int32{
	"UNSPECIFIED":    0,
	"CreatePlayer":   10001,
	"PlayerBaseInfo": 10002,
}

func (x MsgTypeCmd_Lobby) String() string {
	return proto1.EnumName(MsgTypeCmd_Lobby_name, int32(x))
}
func (MsgTypeCmd_Lobby) EnumDescriptor() ([]byte, []int) { return fileDescriptorLobby, []int{0} }

type EnumCreatePlayer_Error int32

const (
	EnumCreatePlayer_NoErr    EnumCreatePlayer_Error = 0
	EnumCreatePlayer_ErrDB    EnumCreatePlayer_Error = 1
	EnumCreatePlayer_ErrExist EnumCreatePlayer_Error = 2
)

var EnumCreatePlayer_Error_name = map[int32]string{
	0: "NoErr",
	1: "ErrDB",
	2: "ErrExist",
}
var EnumCreatePlayer_Error_value = map[string]int32{
	"NoErr":    0,
	"ErrDB":    1,
	"ErrExist": 2,
}

func (x EnumCreatePlayer_Error) String() string {
	return proto1.EnumName(EnumCreatePlayer_Error_name, int32(x))
}
func (EnumCreatePlayer_Error) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorLobby, []int{0, 0}
}

type EnumPlayerBaseInfo_Error int32

const (
	EnumPlayerBaseInfo_NoErr      EnumPlayerBaseInfo_Error = 0
	EnumPlayerBaseInfo_ErrDB      EnumPlayerBaseInfo_Error = 1
	EnumPlayerBaseInfo_ErrNoExist EnumPlayerBaseInfo_Error = 2
)

var EnumPlayerBaseInfo_Error_name = map[int32]string{
	0: "NoErr",
	1: "ErrDB",
	2: "ErrNoExist",
}
var EnumPlayerBaseInfo_Error_value = map[string]int32{
	"NoErr":      0,
	"ErrDB":      1,
	"ErrNoExist": 2,
}

func (x EnumPlayerBaseInfo_Error) String() string {
	return proto1.EnumName(EnumPlayerBaseInfo_Error_name, int32(x))
}
func (EnumPlayerBaseInfo_Error) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorLobby, []int{3, 0}
}

type EnumCreatePlayer struct {
}

func (m *EnumCreatePlayer) Reset()                    { *m = EnumCreatePlayer{} }
func (m *EnumCreatePlayer) String() string            { return proto1.CompactTextString(m) }
func (*EnumCreatePlayer) ProtoMessage()               {}
func (*EnumCreatePlayer) Descriptor() ([]byte, []int) { return fileDescriptorLobby, []int{0} }

type MsgCreatePlayer struct {
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Sex  int32  `protobuf:"varint,2,opt,name=Sex,proto3" json:"Sex,omitempty"`
}

func (m *MsgCreatePlayer) Reset()                    { *m = MsgCreatePlayer{} }
func (m *MsgCreatePlayer) String() string            { return proto1.CompactTextString(m) }
func (*MsgCreatePlayer) ProtoMessage()               {}
func (*MsgCreatePlayer) Descriptor() ([]byte, []int) { return fileDescriptorLobby, []int{1} }

func (m *MsgCreatePlayer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgCreatePlayer) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

type MsgCreatePlayerResult struct {
	Err EnumCreatePlayer_Error `protobuf:"varint,1,opt,name=Err,proto3,enum=proto.EnumCreatePlayer_Error" json:"Err,omitempty"`
}

func (m *MsgCreatePlayerResult) Reset()                    { *m = MsgCreatePlayerResult{} }
func (m *MsgCreatePlayerResult) String() string            { return proto1.CompactTextString(m) }
func (*MsgCreatePlayerResult) ProtoMessage()               {}
func (*MsgCreatePlayerResult) Descriptor() ([]byte, []int) { return fileDescriptorLobby, []int{2} }

func (m *MsgCreatePlayerResult) GetErr() EnumCreatePlayer_Error {
	if m != nil {
		return m.Err
	}
	return EnumCreatePlayer_NoErr
}

type EnumPlayerBaseInfo struct {
}

func (m *EnumPlayerBaseInfo) Reset()                    { *m = EnumPlayerBaseInfo{} }
func (m *EnumPlayerBaseInfo) String() string            { return proto1.CompactTextString(m) }
func (*EnumPlayerBaseInfo) ProtoMessage()               {}
func (*EnumPlayerBaseInfo) Descriptor() ([]byte, []int) { return fileDescriptorLobby, []int{3} }

type MsgPlayerBaseInfo struct {
}

func (m *MsgPlayerBaseInfo) Reset()                    { *m = MsgPlayerBaseInfo{} }
func (m *MsgPlayerBaseInfo) String() string            { return proto1.CompactTextString(m) }
func (*MsgPlayerBaseInfo) ProtoMessage()               {}
func (*MsgPlayerBaseInfo) Descriptor() ([]byte, []int) { return fileDescriptorLobby, []int{4} }

type MsgPlayerBaseInfoResult struct {
	Err  EnumPlayerBaseInfo_Error `protobuf:"varint,1,opt,name=Err,proto3,enum=proto.EnumPlayerBaseInfo_Error" json:"Err,omitempty"`
	Name string                   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Sex  int32                    `protobuf:"varint,3,opt,name=Sex,proto3" json:"Sex,omitempty"`
}

func (m *MsgPlayerBaseInfoResult) Reset()                    { *m = MsgPlayerBaseInfoResult{} }
func (m *MsgPlayerBaseInfoResult) String() string            { return proto1.CompactTextString(m) }
func (*MsgPlayerBaseInfoResult) ProtoMessage()               {}
func (*MsgPlayerBaseInfoResult) Descriptor() ([]byte, []int) { return fileDescriptorLobby, []int{5} }

func (m *MsgPlayerBaseInfoResult) GetErr() EnumPlayerBaseInfo_Error {
	if m != nil {
		return m.Err
	}
	return EnumPlayerBaseInfo_NoErr
}

func (m *MsgPlayerBaseInfoResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgPlayerBaseInfoResult) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func init() {
	proto1.RegisterType((*EnumCreatePlayer)(nil), "proto.EnumCreatePlayer")
	proto1.RegisterType((*MsgCreatePlayer)(nil), "proto.MsgCreatePlayer")
	proto1.RegisterType((*MsgCreatePlayerResult)(nil), "proto.MsgCreatePlayerResult")
	proto1.RegisterType((*EnumPlayerBaseInfo)(nil), "proto.EnumPlayerBaseInfo")
	proto1.RegisterType((*MsgPlayerBaseInfo)(nil), "proto.MsgPlayerBaseInfo")
	proto1.RegisterType((*MsgPlayerBaseInfoResult)(nil), "proto.MsgPlayerBaseInfoResult")
	proto1.RegisterEnum("proto.MsgTypeCmd_Lobby", MsgTypeCmd_Lobby_name, MsgTypeCmd_Lobby_value)
	proto1.RegisterEnum("proto.EnumCreatePlayer_Error", EnumCreatePlayer_Error_name, EnumCreatePlayer_Error_value)
	proto1.RegisterEnum("proto.EnumPlayerBaseInfo_Error", EnumPlayerBaseInfo_Error_name, EnumPlayerBaseInfo_Error_value)
}
func (m *EnumCreatePlayer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnumCreatePlayer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *MsgCreatePlayer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePlayer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLobby(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Sex != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintLobby(dAtA, i, uint64(m.Sex))
	}
	return i, nil
}

func (m *MsgCreatePlayerResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePlayerResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Err != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLobby(dAtA, i, uint64(m.Err))
	}
	return i, nil
}

func (m *EnumPlayerBaseInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnumPlayerBaseInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *MsgPlayerBaseInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlayerBaseInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *MsgPlayerBaseInfoResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlayerBaseInfoResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Err != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLobby(dAtA, i, uint64(m.Err))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLobby(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Sex != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintLobby(dAtA, i, uint64(m.Sex))
	}
	return i, nil
}

func encodeFixed64Lobby(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Lobby(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintLobby(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EnumCreatePlayer) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *MsgCreatePlayer) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovLobby(uint64(l))
	}
	if m.Sex != 0 {
		n += 1 + sovLobby(uint64(m.Sex))
	}
	return n
}

func (m *MsgCreatePlayerResult) Size() (n int) {
	var l int
	_ = l
	if m.Err != 0 {
		n += 1 + sovLobby(uint64(m.Err))
	}
	return n
}

func (m *EnumPlayerBaseInfo) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *MsgPlayerBaseInfo) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *MsgPlayerBaseInfoResult) Size() (n int) {
	var l int
	_ = l
	if m.Err != 0 {
		n += 1 + sovLobby(uint64(m.Err))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovLobby(uint64(l))
	}
	if m.Sex != 0 {
		n += 1 + sovLobby(uint64(m.Sex))
	}
	return n
}

func sovLobby(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLobby(x uint64) (n int) {
	return sovLobby(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EnumCreatePlayer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLobby
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
			return fmt.Errorf("proto: EnumCreatePlayer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnumCreatePlayer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLobby(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLobby
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
func (m *MsgCreatePlayer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLobby
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
			return fmt.Errorf("proto: MsgCreatePlayer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePlayer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLobby
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
				return ErrInvalidLengthLobby
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sex", wireType)
			}
			m.Sex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLobby
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sex |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLobby(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLobby
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
func (m *MsgCreatePlayerResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLobby
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
			return fmt.Errorf("proto: MsgCreatePlayerResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePlayerResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Err", wireType)
			}
			m.Err = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLobby
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Err |= (EnumCreatePlayer_Error(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLobby(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLobby
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
func (m *EnumPlayerBaseInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLobby
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
			return fmt.Errorf("proto: EnumPlayerBaseInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnumPlayerBaseInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLobby(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLobby
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
func (m *MsgPlayerBaseInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLobby
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
			return fmt.Errorf("proto: MsgPlayerBaseInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlayerBaseInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLobby(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLobby
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
func (m *MsgPlayerBaseInfoResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLobby
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
			return fmt.Errorf("proto: MsgPlayerBaseInfoResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlayerBaseInfoResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Err", wireType)
			}
			m.Err = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLobby
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Err |= (EnumPlayerBaseInfo_Error(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLobby
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
				return ErrInvalidLengthLobby
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sex", wireType)
			}
			m.Sex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLobby
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sex |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLobby(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLobby
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
func skipLobby(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLobby
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
					return 0, ErrIntOverflowLobby
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
					return 0, ErrIntOverflowLobby
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
				return 0, ErrInvalidLengthLobby
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLobby
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
				next, err := skipLobby(dAtA[start:])
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
	ErrInvalidLengthLobby = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLobby   = fmt.Errorf("proto: integer overflow")
)

func init() { proto1.RegisterFile("lobby.proto", fileDescriptorLobby) }

var fileDescriptorLobby = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xd1, 0x4a, 0x3a, 0x41,
	0x14, 0xc6, 0x1d, 0xfd, 0xef, 0x9f, 0x3c, 0x8a, 0x8e, 0x23, 0x91, 0x37, 0x6d, 0x32, 0x57, 0x52,
	0x64, 0x54, 0x17, 0x5d, 0x06, 0xae, 0x13, 0x49, 0x39, 0xc8, 0x5a, 0xd7, 0xb1, 0xd2, 0x24, 0xc1,
	0xae, 0x23, 0x67, 0x56, 0x70, 0xdf, 0xa2, 0x7a, 0xaa, 0x2e, 0x7b, 0x84, 0xd8, 0x5e, 0x24, 0x76,
	0x2d, 0x59, 0x57, 0xe8, 0x6a, 0x3e, 0xe6, 0x9c, 0x6f, 0xe6, 0x77, 0xbe, 0x03, 0x15, 0x5f, 0x4f,
	0x26, 0x51, 0x77, 0x8e, 0x3a, 0xd4, 0xcc, 0x4a, 0x0f, 0x7e, 0x09, 0x54, 0xcc, 0x16, 0x81, 0x83,
	0xca, 0x0b, 0xd5, 0xc8, 0xf7, 0x22, 0x85, 0xfc, 0x08, 0x2c, 0x81, 0xa8, 0x91, 0x95, 0xc1, 0x92,
	0x5a, 0x20, 0xd2, 0x42, 0x22, 0x05, 0x62, 0xbf, 0x47, 0x09, 0xab, 0xc2, 0x8e, 0x40, 0x14, 0xcb,
	0x67, 0x13, 0xd2, 0x22, 0xbf, 0x80, 0xfa, 0xd0, 0x4c, 0xb3, 0x7e, 0xc6, 0xe0, 0x9f, 0xf4, 0x02,
	0xd5, 0x22, 0x6d, 0xd2, 0x29, 0xbb, 0xa9, 0x66, 0x14, 0x4a, 0x63, 0xb5, 0x6c, 0x15, 0xdb, 0xa4,
	0x63, 0xb9, 0x89, 0xe4, 0xd7, 0xb0, 0x9b, 0x33, 0xba, 0xca, 0x2c, 0xfc, 0x90, 0x9d, 0x40, 0x49,
	0x20, 0xa6, 0xee, 0xda, 0xd9, 0xfe, 0x0a, 0xb7, 0x9b, 0x87, 0xec, 0xa6, 0x84, 0x6e, 0xd2, 0xc9,
	0x1d, 0x60, 0x49, 0x79, 0x55, 0xe8, 0x79, 0x46, 0x0d, 0x66, 0x4f, 0x9a, 0x1f, 0xff, 0x3d, 0x45,
	0x0d, 0x40, 0x20, 0x4a, 0xfd, 0x3b, 0x47, 0x13, 0x1a, 0x43, 0x33, 0xcd, 0xbd, 0x81, 0xb0, 0xb7,
	0x75, 0xf9, 0x43, 0x79, 0x9a, 0xa5, 0x3c, 0xc8, 0x50, 0x6e, 0x76, 0x67, 0x38, 0xd7, 0xb9, 0x14,
	0xb7, 0x73, 0x29, 0xad, 0x73, 0x39, 0xbc, 0x01, 0x3a, 0x34, 0xd3, 0xbb, 0x68, 0xae, 0x9c, 0xe0,
	0xf1, 0xe1, 0x36, 0x59, 0x19, 0xab, 0x43, 0xe5, 0x5e, 0x8e, 0x47, 0xc2, 0x19, 0x5c, 0x0d, 0x44,
	0x9f, 0x16, 0x58, 0x03, 0xaa, 0xd9, 0x34, 0xe8, 0xab, 0x64, 0x4d, 0xa8, 0x6d, 0x7e, 0x4d, 0xdf,
	0x64, 0x8f, 0xbe, 0xc7, 0x36, 0xf9, 0x88, 0x6d, 0xf2, 0x19, 0xdb, 0xe4, 0xe5, 0xcb, 0x2e, 0x4c,
	0xfe, 0xa7, 0xa4, 0xe7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x92, 0x6e, 0xb8, 0x6c, 0x0d, 0x02,
	0x00, 0x00,
}
