// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: observer/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	common "github.com/zeta-chain/zetacore/common"
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

type Policy_Type int32

const (
	Policy_Type_stop_inbound_cctx    Policy_Type = 0
	Policy_Type_deploy_fungible_coin Policy_Type = 1
	Policy_Type_update_client_params Policy_Type = 2
)

var Policy_Type_name = map[int32]string{
	0: "stop_inbound_cctx",
	1: "deploy_fungible_coin",
	2: "update_client_params",
}

var Policy_Type_value = map[string]int32{
	"stop_inbound_cctx":    0,
	"deploy_fungible_coin": 1,
	"update_client_params": 2,
}

func (x Policy_Type) String() string {
	return proto.EnumName(Policy_Type_name, int32(x))
}

func (Policy_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4542fa62877488a1, []int{0}
}

type ClientParams struct {
	ConfirmationCount int64 `protobuf:"varint,1,opt,name=confirmation_count,json=confirmationCount,proto3" json:"confirmation_count,omitempty"`
	GasPriceTicker    int64 `protobuf:"varint,2,opt,name=gas_price_ticker,json=gasPriceTicker,proto3" json:"gas_price_ticker,omitempty"`
}

func (m *ClientParams) Reset()         { *m = ClientParams{} }
func (m *ClientParams) String() string { return proto.CompactTextString(m) }
func (*ClientParams) ProtoMessage()    {}
func (*ClientParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_4542fa62877488a1, []int{0}
}
func (m *ClientParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientParams.Merge(m, src)
}
func (m *ClientParams) XXX_Size() int {
	return m.Size()
}
func (m *ClientParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientParams.DiscardUnknown(m)
}

var xxx_messageInfo_ClientParams proto.InternalMessageInfo

func (m *ClientParams) GetConfirmationCount() int64 {
	if m != nil {
		return m.ConfirmationCount
	}
	return 0
}

func (m *ClientParams) GetGasPriceTicker() int64 {
	if m != nil {
		return m.GasPriceTicker
	}
	return 0
}

type ObserverParams struct {
	Chain                 *common.Chain                          `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	BallotThreshold       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=ballot_threshold,json=ballotThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"ballot_threshold"`
	MinObserverDelegation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=min_observer_delegation,json=minObserverDelegation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_observer_delegation"`
	IsSupported           bool                                   `protobuf:"varint,5,opt,name=is_supported,json=isSupported,proto3" json:"is_supported,omitempty"`
}

func (m *ObserverParams) Reset()         { *m = ObserverParams{} }
func (m *ObserverParams) String() string { return proto.CompactTextString(m) }
func (*ObserverParams) ProtoMessage()    {}
func (*ObserverParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_4542fa62877488a1, []int{1}
}
func (m *ObserverParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ObserverParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ObserverParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ObserverParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObserverParams.Merge(m, src)
}
func (m *ObserverParams) XXX_Size() int {
	return m.Size()
}
func (m *ObserverParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ObserverParams.DiscardUnknown(m)
}

var xxx_messageInfo_ObserverParams proto.InternalMessageInfo

func (m *ObserverParams) GetChain() *common.Chain {
	if m != nil {
		return m.Chain
	}
	return nil
}

func (m *ObserverParams) GetIsSupported() bool {
	if m != nil {
		return m.IsSupported
	}
	return false
}

type Admin_Policy struct {
	PolicyType Policy_Type `protobuf:"varint,1,opt,name=policy_type,json=policyType,proto3,enum=zetachain.zetacore.observer.Policy_Type" json:"policy_type,omitempty"`
	Address    string      `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *Admin_Policy) Reset()         { *m = Admin_Policy{} }
func (m *Admin_Policy) String() string { return proto.CompactTextString(m) }
func (*Admin_Policy) ProtoMessage()    {}
func (*Admin_Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_4542fa62877488a1, []int{2}
}
func (m *Admin_Policy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Admin_Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Admin_Policy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Admin_Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Admin_Policy.Merge(m, src)
}
func (m *Admin_Policy) XXX_Size() int {
	return m.Size()
}
func (m *Admin_Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Admin_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Admin_Policy proto.InternalMessageInfo

func (m *Admin_Policy) GetPolicyType() Policy_Type {
	if m != nil {
		return m.PolicyType
	}
	return Policy_Type_stop_inbound_cctx
}

func (m *Admin_Policy) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// Params defines the parameters for the module.
type Params struct {
	ObserverParams []*ObserverParams `protobuf:"bytes,1,rep,name=observer_params,json=observerParams,proto3" json:"observer_params,omitempty"`
	AdminPolicy    []*Admin_Policy   `protobuf:"bytes,2,rep,name=admin_policy,json=adminPolicy,proto3" json:"admin_policy,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_4542fa62877488a1, []int{3}
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

func (m *Params) GetObserverParams() []*ObserverParams {
	if m != nil {
		return m.ObserverParams
	}
	return nil
}

func (m *Params) GetAdminPolicy() []*Admin_Policy {
	if m != nil {
		return m.AdminPolicy
	}
	return nil
}

func init() {
	proto.RegisterEnum("zetachain.zetacore.observer.Policy_Type", Policy_Type_name, Policy_Type_value)
	proto.RegisterType((*ClientParams)(nil), "zetachain.zetacore.observer.ClientParams")
	proto.RegisterType((*ObserverParams)(nil), "zetachain.zetacore.observer.ObserverParams")
	proto.RegisterType((*Admin_Policy)(nil), "zetachain.zetacore.observer.Admin_Policy")
	proto.RegisterType((*Params)(nil), "zetachain.zetacore.observer.Params")
}

func init() { proto.RegisterFile("observer/params.proto", fileDescriptor_4542fa62877488a1) }

var fileDescriptor_4542fa62877488a1 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xb6, 0xd3, 0xb4, 0xd0, 0x73, 0x48, 0x53, 0xd3, 0xa8, 0x56, 0x90, 0x9c, 0x10, 0x24, 0x64,
	0x40, 0xb1, 0xa5, 0xb0, 0xb1, 0xd1, 0x74, 0xa9, 0x84, 0x44, 0x64, 0xb2, 0xc0, 0xc0, 0xc9, 0x3e,
	0x5f, 0x9c, 0x53, 0xed, 0x3b, 0xcb, 0x77, 0x46, 0x0d, 0xbf, 0x82, 0x11, 0x89, 0x85, 0x81, 0x81,
	0x99, 0x5f, 0xd1, 0xb1, 0x23, 0x62, 0xa8, 0x50, 0xf2, 0x47, 0x90, 0xef, 0xec, 0x28, 0x5d, 0x32,
	0x30, 0xdd, 0xbb, 0xf7, 0xee, 0x7d, 0xef, 0x7d, 0xef, 0x7d, 0x07, 0xba, 0x2c, 0xe4, 0x38, 0xff,
	0x84, 0x73, 0x2f, 0x0b, 0xf2, 0x20, 0xe5, 0x6e, 0x96, 0x33, 0xc1, 0xcc, 0x47, 0x9f, 0xb1, 0x08,
	0xd0, 0x22, 0x20, 0xd4, 0x95, 0x16, 0xcb, 0xb1, 0x5b, 0xbf, 0xec, 0x9d, 0xc4, 0x2c, 0x66, 0xf2,
	0x9d, 0x57, 0x5a, 0x2a, 0xa5, 0x77, 0xba, 0x41, 0xaa, 0x8d, 0x2a, 0xf0, 0x10, 0xb1, 0x34, 0x65,
	0xd4, 0x53, 0x87, 0x72, 0x0e, 0x63, 0xd0, 0x9a, 0x24, 0x04, 0x53, 0x31, 0x95, 0x65, 0xcd, 0x11,
	0x30, 0x11, 0xa3, 0x73, 0x92, 0xa7, 0x81, 0x20, 0x8c, 0x42, 0xc4, 0x0a, 0x2a, 0x2c, 0x7d, 0xa0,
	0x3b, 0x7b, 0xfe, 0xf1, 0x76, 0x64, 0x52, 0x06, 0x4c, 0x07, 0x74, 0xe2, 0x80, 0xc3, 0x2c, 0x27,
	0x08, 0x43, 0x41, 0xd0, 0x25, 0xce, 0xad, 0x86, 0x7c, 0xdc, 0x8e, 0x03, 0x3e, 0x2d, 0xdd, 0x33,
	0xe9, 0x1d, 0x7e, 0x6b, 0x80, 0xf6, 0xdb, 0xaa, 0xa1, 0xaa, 0xd6, 0x13, 0xb0, 0x2f, 0xa9, 0x49,
	0x78, 0x63, 0xfc, 0xc0, 0xad, 0x3a, 0x9b, 0x94, 0x4e, 0x5f, 0xc5, 0xcc, 0xf7, 0xa0, 0x13, 0x06,
	0x49, 0xc2, 0x04, 0x14, 0x8b, 0x1c, 0xf3, 0x05, 0x4b, 0x22, 0x6b, 0x6f, 0xa0, 0x3b, 0x87, 0x67,
	0xee, 0xf5, 0x6d, 0x5f, 0xfb, 0x73, 0xdb, 0x7f, 0x1a, 0x13, 0xb1, 0x28, 0xc2, 0x32, 0xdb, 0x43,
	0x8c, 0xa7, 0x8c, 0x57, 0xc7, 0x88, 0x47, 0x97, 0x9e, 0x58, 0x66, 0x98, 0xbb, 0xe7, 0x18, 0xf9,
	0x47, 0x0a, 0x67, 0x56, 0xc3, 0x98, 0x73, 0x70, 0x9a, 0x12, 0x0a, 0xeb, 0x31, 0xc1, 0x08, 0x27,
	0x38, 0x96, 0xe4, 0xac, 0xe6, 0x7f, 0x55, 0xe8, 0xa6, 0x84, 0xd6, 0x1c, 0xcf, 0x37, 0x60, 0xe6,
	0x63, 0xd0, 0x22, 0x1c, 0xf2, 0x22, 0xcb, 0x58, 0x2e, 0x70, 0x64, 0xed, 0x0f, 0x74, 0xe7, 0xbe,
	0x6f, 0x10, 0xfe, 0xae, 0x76, 0x0d, 0x39, 0x68, 0xbd, 0x8e, 0xca, 0x66, 0xa6, 0x2c, 0x21, 0x68,
	0x69, 0x5e, 0x00, 0x23, 0x93, 0x16, 0x2c, 0xd1, 0xe5, 0x80, 0xda, 0x63, 0xc7, 0xdd, 0xa1, 0x06,
	0x57, 0x65, 0xc2, 0xd9, 0x32, 0xc3, 0x3e, 0x50, 0xc9, 0xa5, 0x6d, 0x5a, 0xe0, 0x5e, 0x10, 0x45,
	0x39, 0xe6, 0x5c, 0x6e, 0xe6, 0xd0, 0xaf, 0xaf, 0xc3, 0x5f, 0x3a, 0x38, 0xa8, 0x56, 0x31, 0x03,
	0x47, 0x9b, 0x31, 0x28, 0x01, 0x5a, 0xfa, 0x60, 0xcf, 0x31, 0xc6, 0x2f, 0x76, 0xd6, 0xbc, 0xbb,
	0x50, 0xbf, 0xcd, 0xee, 0x2e, 0xf8, 0x0d, 0x68, 0x05, 0x92, 0x95, 0x6a, 0xc7, 0x6a, 0x48, 0xc8,
	0x67, 0x3b, 0x21, 0xb7, 0xc7, 0xe0, 0x1b, 0x32, 0x5d, 0x5d, 0x5e, 0x35, 0xbf, 0x7e, 0xef, 0x6b,
	0xcf, 0x3f, 0x02, 0x63, 0x8b, 0xa9, 0xd9, 0x05, 0xc7, 0x5c, 0xb0, 0x0c, 0x12, 0x1a, 0xb2, 0x82,
	0x46, 0x10, 0x21, 0x71, 0xd5, 0xd1, 0x4c, 0x0b, 0x9c, 0x44, 0x38, 0x4b, 0xd8, 0x12, 0xce, 0x0b,
	0x1a, 0x93, 0x30, 0xc1, 0x10, 0x31, 0x42, 0x3b, 0x7a, 0x19, 0x29, 0xb2, 0x28, 0x10, 0x18, 0x22,
	0xa9, 0xfb, 0x8a, 0x6e, 0xa7, 0xd1, 0x6b, 0xfe, 0xfc, 0x61, 0xeb, 0x67, 0x17, 0xd7, 0x2b, 0x5b,
	0xbf, 0x59, 0xd9, 0xfa, 0xdf, 0x95, 0xad, 0x7f, 0x59, 0xdb, 0xda, 0xcd, 0xda, 0xd6, 0x7e, 0xaf,
	0x6d, 0xed, 0x83, 0xb7, 0xa5, 0x82, 0xb2, 0xef, 0x91, 0xa4, 0xe0, 0xd5, 0x14, 0xbc, 0xab, 0xcd,
	0x87, 0x53, 0x92, 0x08, 0x0f, 0xe4, 0x17, 0x7b, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xcd,
	0xdf, 0xa3, 0xdc, 0x03, 0x00, 0x00,
}

func (m *ClientParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GasPriceTicker != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.GasPriceTicker))
		i--
		dAtA[i] = 0x10
	}
	if m.ConfirmationCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ConfirmationCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ObserverParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObserverParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ObserverParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsSupported {
		i--
		if m.IsSupported {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.MinObserverDelegation.Size()
		i -= size
		if _, err := m.MinObserverDelegation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.BallotThreshold.Size()
		i -= size
		if _, err := m.BallotThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Chain != nil {
		{
			size, err := m.Chain.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Admin_Policy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Admin_Policy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Admin_Policy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.PolicyType != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PolicyType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
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
	if len(m.AdminPolicy) > 0 {
		for iNdEx := len(m.AdminPolicy) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AdminPolicy[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ObserverParams) > 0 {
		for iNdEx := len(m.ObserverParams) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ObserverParams[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *ClientParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConfirmationCount != 0 {
		n += 1 + sovParams(uint64(m.ConfirmationCount))
	}
	if m.GasPriceTicker != 0 {
		n += 1 + sovParams(uint64(m.GasPriceTicker))
	}
	return n
}

func (m *ObserverParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Chain != nil {
		l = m.Chain.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.BallotThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MinObserverDelegation.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.IsSupported {
		n += 2
	}
	return n
}

func (m *Admin_Policy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PolicyType != 0 {
		n += 1 + sovParams(uint64(m.PolicyType))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ObserverParams) > 0 {
		for _, e := range m.ObserverParams {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.AdminPolicy) > 0 {
		for _, e := range m.AdminPolicy {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ClientParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfirmationCount", wireType)
			}
			m.ConfirmationCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConfirmationCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPriceTicker", wireType)
			}
			m.GasPriceTicker = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasPriceTicker |= int64(b&0x7F) << shift
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
func (m *ObserverParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ObserverParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObserverParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Chain == nil {
				m.Chain = &common.Chain{}
			}
			if err := m.Chain.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BallotThreshold", wireType)
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
			if err := m.BallotThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinObserverDelegation", wireType)
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
			if err := m.MinObserverDelegation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsSupported", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.IsSupported = bool(v != 0)
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
func (m *Admin_Policy) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Admin_Policy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Admin_Policy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PolicyType", wireType)
			}
			m.PolicyType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PolicyType |= Policy_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
				return fmt.Errorf("proto: wrong wireType = %d for field ObserverParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObserverParams = append(m.ObserverParams, &ObserverParams{})
			if err := m.ObserverParams[len(m.ObserverParams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminPolicy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdminPolicy = append(m.AdminPolicy, &Admin_Policy{})
			if err := m.AdminPolicy[len(m.AdminPolicy)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
