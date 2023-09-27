// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fungible/foreign_coins.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	common "github.com/zeta-chain/zetacore/common"
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

type ForeignCoins struct {
	//  string index = 1;
	Zrc20ContractAddress string                                 `protobuf:"bytes,2,opt,name=zrc20_contract_address,json=zrc20ContractAddress,proto3" json:"zrc20_contract_address,omitempty"`
	Asset                string                                 `protobuf:"bytes,3,opt,name=asset,proto3" json:"asset,omitempty"`
	ForeignChainId       int64                                  `protobuf:"varint,4,opt,name=foreign_chain_id,json=foreignChainId,proto3" json:"foreign_chain_id,omitempty"`
	Decimals             uint32                                 `protobuf:"varint,5,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Name                 string                                 `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Symbol               string                                 `protobuf:"bytes,7,opt,name=symbol,proto3" json:"symbol,omitempty"`
	CoinType             common.CoinType                        `protobuf:"varint,8,opt,name=coin_type,json=coinType,proto3,enum=common.CoinType" json:"coin_type,omitempty"`
	GasLimit             uint64                                 `protobuf:"varint,9,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	Paused               bool                                   `protobuf:"varint,10,opt,name=paused,proto3" json:"paused,omitempty"`
	LiquidityCap         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,11,opt,name=liquidity_cap,json=liquidityCap,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"liquidity_cap"`
}

func (m *ForeignCoins) Reset()         { *m = ForeignCoins{} }
func (m *ForeignCoins) String() string { return proto.CompactTextString(m) }
func (*ForeignCoins) ProtoMessage()    {}
func (*ForeignCoins) Descriptor() ([]byte, []int) {
	return fileDescriptor_5285bb476cecbbf8, []int{0}
}
func (m *ForeignCoins) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ForeignCoins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ForeignCoins.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ForeignCoins) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForeignCoins.Merge(m, src)
}
func (m *ForeignCoins) XXX_Size() int {
	return m.Size()
}
func (m *ForeignCoins) XXX_DiscardUnknown() {
	xxx_messageInfo_ForeignCoins.DiscardUnknown(m)
}

var xxx_messageInfo_ForeignCoins proto.InternalMessageInfo

func (m *ForeignCoins) GetZrc20ContractAddress() string {
	if m != nil {
		return m.Zrc20ContractAddress
	}
	return ""
}

func (m *ForeignCoins) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *ForeignCoins) GetForeignChainId() int64 {
	if m != nil {
		return m.ForeignChainId
	}
	return 0
}

func (m *ForeignCoins) GetDecimals() uint32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *ForeignCoins) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ForeignCoins) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ForeignCoins) GetCoinType() common.CoinType {
	if m != nil {
		return m.CoinType
	}
	return common.CoinType_Zeta
}

func (m *ForeignCoins) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *ForeignCoins) GetPaused() bool {
	if m != nil {
		return m.Paused
	}
	return false
}

func init() {
	proto.RegisterType((*ForeignCoins)(nil), "zetachain.zetacore.fungible.ForeignCoins")
}

func init() { proto.RegisterFile("fungible/foreign_coins.proto", fileDescriptor_5285bb476cecbbf8) }

var fileDescriptor_5285bb476cecbbf8 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x52, 0xc1, 0x6e, 0xd4, 0x30,
	0x14, 0x5c, 0xb3, 0xdb, 0x25, 0x6b, 0xda, 0xaa, 0x32, 0xab, 0xca, 0xda, 0xa2, 0x34, 0xe2, 0x80,
	0x72, 0xd9, 0x18, 0x15, 0x7e, 0x80, 0x46, 0x42, 0x5a, 0x89, 0x53, 0xe0, 0xc4, 0x25, 0x72, 0x6c,
	0x37, 0xb5, 0x48, 0xec, 0x10, 0x3b, 0x12, 0xe9, 0x9d, 0x3b, 0x9f, 0xd5, 0x63, 0x8f, 0x88, 0x43,
	0x85, 0x76, 0x7f, 0x04, 0xd9, 0x49, 0x23, 0x4e, 0x9e, 0x99, 0x17, 0xcf, 0x9b, 0xbc, 0x67, 0xf8,
	0xea, 0xa6, 0x53, 0xa5, 0x2c, 0x2a, 0x41, 0x6e, 0x74, 0x2b, 0x64, 0xa9, 0x72, 0xa6, 0xa5, 0x32,
	0x49, 0xd3, 0x6a, 0xab, 0xd1, 0xc5, 0x9d, 0xb0, 0x94, 0xdd, 0x52, 0xa9, 0x12, 0x8f, 0x74, 0x2b,
	0x92, 0xa7, 0x0b, 0x9b, 0x97, 0x4c, 0xd7, 0xb5, 0x56, 0x64, 0x38, 0x86, 0x1b, 0x9b, 0x75, 0xa9,
	0x4b, 0xed, 0x21, 0x71, 0x68, 0x50, 0x5f, 0xff, 0x9c, 0xc3, 0xe3, 0x8f, 0x83, 0x7f, 0xea, 0xec,
	0xd1, 0x7b, 0x78, 0x7e, 0xd7, 0xb2, 0xab, 0xb7, 0x39, 0xd3, 0xca, 0xb6, 0x94, 0xd9, 0x9c, 0x72,
	0xde, 0x0a, 0x63, 0xf0, 0xb3, 0x08, 0xc4, 0xab, 0x6c, 0xed, 0xab, 0xe9, 0x58, 0xfc, 0x30, 0xd4,
	0xd0, 0x1a, 0x1e, 0x51, 0x63, 0x84, 0xc5, 0x73, 0xff, 0xd1, 0x40, 0x50, 0x0c, 0xcf, 0xa6, 0xec,
	0x2e, 0x6a, 0x2e, 0x39, 0x5e, 0x44, 0x20, 0x9e, 0x67, 0xa7, 0xa3, 0x9e, 0x3a, 0x79, 0xc7, 0xd1,
	0x06, 0x06, 0x5c, 0x30, 0x59, 0xd3, 0xca, 0xe0, 0xa3, 0x08, 0xc4, 0x27, 0xd9, 0xc4, 0x11, 0x82,
	0x0b, 0x45, 0x6b, 0x81, 0x97, 0xde, 0xda, 0x63, 0x74, 0x0e, 0x97, 0xa6, 0xaf, 0x0b, 0x5d, 0xe1,
	0xe7, 0x5e, 0x1d, 0x19, 0xda, 0xc2, 0x95, 0x9b, 0x52, 0x6e, 0xfb, 0x46, 0xe0, 0x20, 0x02, 0xf1,
	0xe9, 0xd5, 0x59, 0x32, 0x8e, 0xc1, 0xfd, 0xdf, 0x97, 0xbe, 0x11, 0x59, 0xc0, 0x46, 0x84, 0x2e,
	0xe0, 0xaa, 0xa4, 0x26, 0xaf, 0x64, 0x2d, 0x2d, 0x5e, 0x45, 0x20, 0x5e, 0x64, 0x41, 0x49, 0xcd,
	0x27, 0xc7, 0x5d, 0x8f, 0x86, 0x76, 0x46, 0x70, 0x0c, 0x23, 0x10, 0x07, 0xd9, 0xc8, 0xd0, 0x67,
	0x78, 0x52, 0xc9, 0xef, 0x9d, 0xe4, 0xd2, 0xf6, 0x39, 0xa3, 0x0d, 0x7e, 0xe1, 0x22, 0x5c, 0x27,
	0xf7, 0x8f, 0x97, 0xb3, 0x3f, 0x8f, 0x97, 0x6f, 0x4a, 0x69, 0x6f, 0xbb, 0xc2, 0x75, 0x25, 0x4c,
	0x9b, 0x5a, 0x9b, 0xf1, 0xd8, 0x1a, 0xfe, 0x8d, 0xb8, 0x60, 0x26, 0xd9, 0x29, 0x9b, 0x1d, 0x4f,
	0x26, 0x29, 0x6d, 0xae, 0x77, 0xf7, 0xfb, 0x10, 0x3c, 0xec, 0x43, 0xf0, 0x77, 0x1f, 0x82, 0x5f,
	0x87, 0x70, 0xf6, 0x70, 0x08, 0x67, 0xbf, 0x0f, 0xe1, 0xec, 0x2b, 0xf9, 0xcf, 0xcf, 0xad, 0x7a,
	0xeb, 0x47, 0x49, 0x9e, 0xb6, 0x4e, 0x7e, 0x90, 0xe9, 0xa1, 0x78, 0xf3, 0x62, 0xe9, 0x37, 0xfb,
	0xee, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0x33, 0x9c, 0x87, 0x41, 0x02, 0x00, 0x00,
}

func (m *ForeignCoins) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ForeignCoins) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ForeignCoins) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.LiquidityCap.Size()
		i -= size
		if _, err := m.LiquidityCap.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintForeignCoins(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	if m.Paused {
		i--
		if m.Paused {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if m.GasLimit != 0 {
		i = encodeVarintForeignCoins(dAtA, i, uint64(m.GasLimit))
		i--
		dAtA[i] = 0x48
	}
	if m.CoinType != 0 {
		i = encodeVarintForeignCoins(dAtA, i, uint64(m.CoinType))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintForeignCoins(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintForeignCoins(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x32
	}
	if m.Decimals != 0 {
		i = encodeVarintForeignCoins(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x28
	}
	if m.ForeignChainId != 0 {
		i = encodeVarintForeignCoins(dAtA, i, uint64(m.ForeignChainId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintForeignCoins(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Zrc20ContractAddress) > 0 {
		i -= len(m.Zrc20ContractAddress)
		copy(dAtA[i:], m.Zrc20ContractAddress)
		i = encodeVarintForeignCoins(dAtA, i, uint64(len(m.Zrc20ContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func encodeVarintForeignCoins(dAtA []byte, offset int, v uint64) int {
	offset -= sovForeignCoins(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ForeignCoins) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Zrc20ContractAddress)
	if l > 0 {
		n += 1 + l + sovForeignCoins(uint64(l))
	}
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovForeignCoins(uint64(l))
	}
	if m.ForeignChainId != 0 {
		n += 1 + sovForeignCoins(uint64(m.ForeignChainId))
	}
	if m.Decimals != 0 {
		n += 1 + sovForeignCoins(uint64(m.Decimals))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovForeignCoins(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovForeignCoins(uint64(l))
	}
	if m.CoinType != 0 {
		n += 1 + sovForeignCoins(uint64(m.CoinType))
	}
	if m.GasLimit != 0 {
		n += 1 + sovForeignCoins(uint64(m.GasLimit))
	}
	if m.Paused {
		n += 2
	}
	l = m.LiquidityCap.Size()
	n += 1 + l + sovForeignCoins(uint64(l))
	return n
}

func sovForeignCoins(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozForeignCoins(x uint64) (n int) {
	return sovForeignCoins(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ForeignCoins) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowForeignCoins
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
			return fmt.Errorf("proto: ForeignCoins: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ForeignCoins: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Zrc20ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowForeignCoins
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
				return ErrInvalidLengthForeignCoins
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthForeignCoins
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Zrc20ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowForeignCoins
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
				return ErrInvalidLengthForeignCoins
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthForeignCoins
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForeignChainId", wireType)
			}
			m.ForeignChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowForeignCoins
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ForeignChainId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowForeignCoins
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowForeignCoins
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
				return ErrInvalidLengthForeignCoins
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthForeignCoins
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowForeignCoins
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
				return ErrInvalidLengthForeignCoins
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthForeignCoins
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinType", wireType)
			}
			m.CoinType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowForeignCoins
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoinType |= common.CoinType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowForeignCoins
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paused", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowForeignCoins
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
			m.Paused = bool(v != 0)
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityCap", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowForeignCoins
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
				return ErrInvalidLengthForeignCoins
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthForeignCoins
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidityCap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipForeignCoins(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthForeignCoins
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
func skipForeignCoins(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowForeignCoins
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
					return 0, ErrIntOverflowForeignCoins
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
					return 0, ErrIntOverflowForeignCoins
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
				return 0, ErrInvalidLengthForeignCoins
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupForeignCoins
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthForeignCoins
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthForeignCoins        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowForeignCoins          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupForeignCoins = fmt.Errorf("proto: unexpected end of group")
)
