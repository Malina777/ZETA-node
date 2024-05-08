// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/zetacore/emissions/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the module.
type Params struct {
	MaxBondFactor               string                                 `protobuf:"bytes,1,opt,name=max_bond_factor,json=maxBondFactor,proto3" json:"max_bond_factor,omitempty"`
	MinBondFactor               string                                 `protobuf:"bytes,2,opt,name=min_bond_factor,json=minBondFactor,proto3" json:"min_bond_factor,omitempty"`
	AvgBlockTime                string                                 `protobuf:"bytes,3,opt,name=avg_block_time,json=avgBlockTime,proto3" json:"avg_block_time,omitempty"`
	TargetBondRatio             string                                 `protobuf:"bytes,4,opt,name=target_bond_ratio,json=targetBondRatio,proto3" json:"target_bond_ratio,omitempty"`
	ValidatorEmissionPercentage string                                 `protobuf:"bytes,5,opt,name=validator_emission_percentage,json=validatorEmissionPercentage,proto3" json:"validator_emission_percentage,omitempty"`
	ObserverEmissionPercentage  string                                 `protobuf:"bytes,6,opt,name=observer_emission_percentage,json=observerEmissionPercentage,proto3" json:"observer_emission_percentage,omitempty"`
	TssSignerEmissionPercentage string                                 `protobuf:"bytes,7,opt,name=tss_signer_emission_percentage,json=tssSignerEmissionPercentage,proto3" json:"tss_signer_emission_percentage,omitempty"`
	DurationFactorConstant      string                                 `protobuf:"bytes,8,opt,name=duration_factor_constant,json=durationFactorConstant,proto3" json:"duration_factor_constant,omitempty"`
	ObserverSlashAmount         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=observer_slash_amount,json=observerSlashAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"observer_slash_amount"`
	BallotMaturityBlocks        int64                                  `protobuf:"varint,10,opt,name=ballot_maturity_blocks,json=ballotMaturityBlocks,proto3" json:"ballot_maturity_blocks,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_259272924aec0acf, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxBondFactor() string {
	if m != nil {
		return m.MaxBondFactor
	}
	return ""
}

func (m *Params) GetMinBondFactor() string {
	if m != nil {
		return m.MinBondFactor
	}
	return ""
}

func (m *Params) GetAvgBlockTime() string {
	if m != nil {
		return m.AvgBlockTime
	}
	return ""
}

func (m *Params) GetTargetBondRatio() string {
	if m != nil {
		return m.TargetBondRatio
	}
	return ""
}

func (m *Params) GetValidatorEmissionPercentage() string {
	if m != nil {
		return m.ValidatorEmissionPercentage
	}
	return ""
}

func (m *Params) GetObserverEmissionPercentage() string {
	if m != nil {
		return m.ObserverEmissionPercentage
	}
	return ""
}

func (m *Params) GetTssSignerEmissionPercentage() string {
	if m != nil {
		return m.TssSignerEmissionPercentage
	}
	return ""
}

func (m *Params) GetDurationFactorConstant() string {
	if m != nil {
		return m.DurationFactorConstant
	}
	return ""
}

func (m *Params) GetBallotMaturityBlocks() int64 {
	if m != nil {
		return m.BallotMaturityBlocks
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "zetachain.zetacore.emissions.Params")
}

func init() {
	proto.RegisterFile("zetachain/zetacore/emissions/params.proto", fileDescriptor_259272924aec0acf)
}

var fileDescriptor_259272924aec0acf = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x6a, 0x14, 0x4d,
	0x10, 0xc0, 0x77, 0xbe, 0xec, 0xb7, 0x9a, 0x46, 0x0d, 0x8e, 0x31, 0x0c, 0x31, 0xce, 0x06, 0x91,
	0x10, 0x85, 0xcc, 0x08, 0x7a, 0x10, 0x4f, 0x3a, 0x41, 0x41, 0x41, 0x08, 0x1b, 0x4f, 0x5e, 0x9a,
	0x9e, 0x99, 0x76, 0xb6, 0xc9, 0x74, 0xd7, 0xd2, 0x55, 0xbb, 0x6c, 0x7c, 0x0a, 0x8f, 0x1e, 0x7d,
	0x9c, 0xdc, 0xcc, 0x51, 0x3c, 0x04, 0xd9, 0x7d, 0x11, 0x99, 0x9e, 0x3f, 0xac, 0xb2, 0x9e, 0xba,
	0xa9, 0xfa, 0xd5, 0x8f, 0xaa, 0xee, 0x62, 0x8f, 0x3e, 0x4b, 0x12, 0xd9, 0x58, 0x28, 0x13, 0xbb,
	0x1b, 0x58, 0x19, 0x4b, 0xad, 0x10, 0x15, 0x18, 0x8c, 0x27, 0xc2, 0x0a, 0x8d, 0xd1, 0xc4, 0x02,
	0x81, 0xbf, 0xd7, 0xa1, 0x51, 0x8b, 0x46, 0x1d, 0xba, 0xbb, 0x5d, 0x40, 0x01, 0x0e, 0x8c, 0xab,
	0x5b, 0x5d, 0xf3, 0xe0, 0x7b, 0x9f, 0x0d, 0x4e, 0x9c, 0xc4, 0x3f, 0x60, 0x5b, 0x5a, 0xcc, 0x79,
	0x0a, 0x26, 0xe7, 0x9f, 0x44, 0x46, 0x60, 0x03, 0x6f, 0xdf, 0x3b, 0xdc, 0x1c, 0xdd, 0xd4, 0x62,
	0x9e, 0x80, 0xc9, 0xdf, 0xb8, 0xa0, 0xe3, 0x94, 0xf9, 0x83, 0xfb, 0xaf, 0xe1, 0x94, 0x59, 0xe1,
	0x1e, 0xb2, 0x5b, 0x62, 0x56, 0xf0, 0xb4, 0x84, 0xec, 0x8c, 0x93, 0xd2, 0x32, 0xd8, 0x70, 0xd8,
	0x0d, 0x31, 0x2b, 0x92, 0x2a, 0xf8, 0x41, 0x69, 0xe9, 0x3f, 0x66, 0xb7, 0x49, 0xd8, 0x42, 0x52,
	0x2d, 0xb4, 0x82, 0x14, 0x04, 0x7d, 0x07, 0x6e, 0xd5, 0x89, 0x4a, 0x39, 0xaa, 0xc2, 0x7e, 0xc2,
	0xee, 0xcf, 0x44, 0xa9, 0x72, 0x41, 0x60, 0x79, 0x3b, 0x19, 0x9f, 0x48, 0x9b, 0x49, 0x43, 0xa2,
	0x90, 0xc1, 0xff, 0xae, 0xee, 0x5e, 0x07, 0xbd, 0x6e, 0x98, 0x93, 0x0e, 0xf1, 0x5f, 0xb2, 0x3d,
	0x48, 0x51, 0xda, 0x99, 0x5c, 0xaf, 0x18, 0x38, 0xc5, 0x6e, 0xcb, 0xac, 0x31, 0x1c, 0xb3, 0x90,
	0x10, 0x39, 0xaa, 0xc2, 0xfc, 0xc3, 0x71, 0xad, 0x6e, 0x83, 0x10, 0x4f, 0x1d, 0xb4, 0x46, 0xf2,
	0x9c, 0x05, 0xf9, 0xd4, 0x0d, 0x6b, 0x9a, 0x47, 0xe4, 0x19, 0x18, 0x24, 0x61, 0x28, 0xb8, 0xee,
	0xca, 0x77, 0xda, 0x7c, 0xfd, 0x9c, 0xc7, 0x4d, 0xd6, 0x4f, 0xd9, 0xdd, 0x6e, 0x00, 0x2c, 0x05,
	0x8e, 0xb9, 0xd0, 0x30, 0x35, 0x14, 0x6c, 0x56, 0x65, 0x49, 0x74, 0x71, 0x35, 0xec, 0xfd, 0xbc,
	0x1a, 0x1e, 0x14, 0x8a, 0xc6, 0xd3, 0x34, 0xca, 0x40, 0xc7, 0x19, 0xa0, 0x06, 0x6c, 0x8e, 0x23,
	0xcc, 0xcf, 0x62, 0x3a, 0x9f, 0x48, 0x8c, 0xde, 0x1a, 0x1a, 0xdd, 0x69, 0x65, 0xa7, 0x95, 0xeb,
	0x95, 0x53, 0xf9, 0xcf, 0xd8, 0x4e, 0x2a, 0xca, 0x12, 0x88, 0x6b, 0x41, 0x53, 0xab, 0xe8, 0xbc,
	0xfe, 0x46, 0x0c, 0xd8, 0xbe, 0x77, 0xb8, 0x31, 0xda, 0xae, 0xb3, 0xef, 0x9b, 0xa4, 0xfb, 0x4d,
	0x7c, 0xd1, 0xff, 0xfa, 0x6d, 0xd8, 0x4b, 0xde, 0x5d, 0x2c, 0x42, 0xef, 0x72, 0x11, 0x7a, 0xbf,
	0x16, 0xa1, 0xf7, 0x65, 0x19, 0xf6, 0x2e, 0x97, 0x61, 0xef, 0xc7, 0x32, 0xec, 0x7d, 0x7c, 0xb2,
	0xd2, 0x52, 0xb5, 0xa0, 0x47, 0x7f, 0xad, 0xf5, 0x7c, 0x65, 0xb1, 0x5d, 0x83, 0xe9, 0xc0, 0x2d,
	0xe9, 0xd3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x6f, 0x44, 0x4d, 0x05, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BallotMaturityBlocks != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BallotMaturityBlocks))
		i--
		dAtA[i] = 0x50
	}
	{
		size := m.ObserverSlashAmount.Size()
		i -= size
		if _, err := m.ObserverSlashAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.DurationFactorConstant) > 0 {
		i -= len(m.DurationFactorConstant)
		copy(dAtA[i:], m.DurationFactorConstant)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DurationFactorConstant)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.TssSignerEmissionPercentage) > 0 {
		i -= len(m.TssSignerEmissionPercentage)
		copy(dAtA[i:], m.TssSignerEmissionPercentage)
		i = encodeVarintParams(dAtA, i, uint64(len(m.TssSignerEmissionPercentage)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ObserverEmissionPercentage) > 0 {
		i -= len(m.ObserverEmissionPercentage)
		copy(dAtA[i:], m.ObserverEmissionPercentage)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ObserverEmissionPercentage)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ValidatorEmissionPercentage) > 0 {
		i -= len(m.ValidatorEmissionPercentage)
		copy(dAtA[i:], m.ValidatorEmissionPercentage)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ValidatorEmissionPercentage)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TargetBondRatio) > 0 {
		i -= len(m.TargetBondRatio)
		copy(dAtA[i:], m.TargetBondRatio)
		i = encodeVarintParams(dAtA, i, uint64(len(m.TargetBondRatio)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AvgBlockTime) > 0 {
		i -= len(m.AvgBlockTime)
		copy(dAtA[i:], m.AvgBlockTime)
		i = encodeVarintParams(dAtA, i, uint64(len(m.AvgBlockTime)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MinBondFactor) > 0 {
		i -= len(m.MinBondFactor)
		copy(dAtA[i:], m.MinBondFactor)
		i = encodeVarintParams(dAtA, i, uint64(len(m.MinBondFactor)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MaxBondFactor) > 0 {
		i -= len(m.MaxBondFactor)
		copy(dAtA[i:], m.MaxBondFactor)
		i = encodeVarintParams(dAtA, i, uint64(len(m.MaxBondFactor)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MaxBondFactor)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.MinBondFactor)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.AvgBlockTime)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.TargetBondRatio)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ValidatorEmissionPercentage)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ObserverEmissionPercentage)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.TssSignerEmissionPercentage)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DurationFactorConstant)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.ObserverSlashAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.BallotMaturityBlocks != 0 {
		n += 1 + sovParams(uint64(m.BallotMaturityBlocks))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBondFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxBondFactor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBondFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinBondFactor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvgBlockTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AvgBlockTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetBondRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetBondRatio = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorEmissionPercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorEmissionPercentage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObserverEmissionPercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObserverEmissionPercentage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TssSignerEmissionPercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TssSignerEmissionPercentage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationFactorConstant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DurationFactorConstant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObserverSlashAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObserverSlashAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BallotMaturityBlocks", wireType)
			}
			m.BallotMaturityBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BallotMaturityBlocks |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
