// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetaobserver/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/zeta-chain/zetacore/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgSetSupportedChains struct {
	Creator   string          `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Chainlist []ObserverChain `protobuf:"varint,2,rep,packed,name=Chainlist,proto3,enum=zetachain.zetacore.zetaobserver.ObserverChain" json:"Chainlist,omitempty"`
}

func (m *MsgSetSupportedChains) Reset()         { *m = MsgSetSupportedChains{} }
func (m *MsgSetSupportedChains) String() string { return proto.CompactTextString(m) }
func (*MsgSetSupportedChains) ProtoMessage()    {}
func (*MsgSetSupportedChains) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae97d3beb5919c1b, []int{0}
}
func (m *MsgSetSupportedChains) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetSupportedChains) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetSupportedChains.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetSupportedChains) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetSupportedChains.Merge(m, src)
}
func (m *MsgSetSupportedChains) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetSupportedChains) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetSupportedChains.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetSupportedChains proto.InternalMessageInfo

func (m *MsgSetSupportedChains) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSetSupportedChains) GetChainlist() []ObserverChain {
	if m != nil {
		return m.Chainlist
	}
	return nil
}

type MsgSetSupportedChainsResponse struct {
}

func (m *MsgSetSupportedChainsResponse) Reset()         { *m = MsgSetSupportedChainsResponse{} }
func (m *MsgSetSupportedChainsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetSupportedChainsResponse) ProtoMessage()    {}
func (*MsgSetSupportedChainsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae97d3beb5919c1b, []int{1}
}
func (m *MsgSetSupportedChainsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetSupportedChainsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetSupportedChainsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetSupportedChainsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetSupportedChainsResponse.Merge(m, src)
}
func (m *MsgSetSupportedChainsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetSupportedChainsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetSupportedChainsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetSupportedChainsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSetSupportedChains)(nil), "zetachain.zetacore.zetaobserver.MsgSetSupportedChains")
	proto.RegisterType((*MsgSetSupportedChainsResponse)(nil), "zetachain.zetacore.zetaobserver.MsgSetSupportedChainsResponse")
}

func init() { proto.RegisterFile("zetaobserver/tx.proto", fileDescriptor_ae97d3beb5919c1b) }

var fileDescriptor_ae97d3beb5919c1b = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xad, 0x4a, 0x2d, 0x49,
	0xcc, 0x4f, 0x2a, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x07, 0x09, 0x27, 0x67, 0x24, 0x66, 0xe6, 0xe9, 0x81, 0x59, 0xf9, 0x45, 0xa9,
	0x7a, 0xc8, 0x2a, 0xa5, 0x84, 0x93, 0xf3, 0x73, 0x73, 0xf3, 0xf3, 0xf4, 0x21, 0x14, 0x44, 0x97,
	0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x98, 0xa9, 0x0f, 0x62, 0x41, 0x45, 0xa5, 0x51, 0xac, 0x80,
	0x31, 0x20, 0x92, 0x4a, 0xf5, 0x5c, 0xa2, 0xbe, 0xc5, 0xe9, 0xc1, 0xa9, 0x25, 0xc1, 0xa5, 0x05,
	0x05, 0xf9, 0x45, 0x25, 0xa9, 0x29, 0xce, 0x20, 0x5b, 0x8b, 0x85, 0x24, 0xb8, 0xd8, 0x93, 0x8b,
	0x52, 0x13, 0x4b, 0xf2, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x1f,
	0x2e, 0x4e, 0xb0, 0x9a, 0x9c, 0xcc, 0xe2, 0x12, 0x09, 0x26, 0x05, 0x66, 0x0d, 0x3e, 0x23, 0x3d,
	0x3d, 0x02, 0xee, 0xd5, 0xf3, 0x87, 0x32, 0xc0, 0x3a, 0x83, 0x10, 0x06, 0x28, 0xc9, 0x73, 0xc9,
	0x62, 0x75, 0x40, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0xd1, 0x14, 0x46, 0x2e, 0x66,
	0xdf, 0xe2, 0x74, 0xa1, 0x1e, 0x46, 0x2e, 0x21, 0x2c, 0xee, 0x34, 0x23, 0x68, 0x35, 0x56, 0xe3,
	0xa5, 0xec, 0xc8, 0xd3, 0x07, 0x73, 0x96, 0x93, 0xef, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9,
	0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e,
	0xcb, 0x31, 0x44, 0x19, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x83,
	0x0c, 0xd4, 0x05, 0x5b, 0xa2, 0x0f, 0xb3, 0x44, 0xbf, 0x42, 0x1f, 0x35, 0xce, 0x2b, 0x0b, 0x52,
	0x8b, 0x93, 0xd8, 0xc0, 0xd1, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x43, 0xdb, 0x91,
	0x10, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	SetSupportedChains(ctx context.Context, in *MsgSetSupportedChains, opts ...grpc.CallOption) (*MsgSetSupportedChainsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetSupportedChains(ctx context.Context, in *MsgSetSupportedChains, opts ...grpc.CallOption) (*MsgSetSupportedChainsResponse, error) {
	out := new(MsgSetSupportedChainsResponse)
	err := c.cc.Invoke(ctx, "/zetachain.zetacore.zetaobserver.Msg/SetSupportedChains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SetSupportedChains(context.Context, *MsgSetSupportedChains) (*MsgSetSupportedChainsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SetSupportedChains(ctx context.Context, req *MsgSetSupportedChains) (*MsgSetSupportedChainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSupportedChains not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SetSupportedChains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetSupportedChains)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetSupportedChains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zetachain.zetacore.zetaobserver.Msg/SetSupportedChains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetSupportedChains(ctx, req.(*MsgSetSupportedChains))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zetachain.zetacore.zetaobserver.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetSupportedChains",
			Handler:    _Msg_SetSupportedChains_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zetaobserver/tx.proto",
}

func (m *MsgSetSupportedChains) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetSupportedChains) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetSupportedChains) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chainlist) > 0 {
		dAtA2 := make([]byte, len(m.Chainlist)*10)
		var j1 int
		for _, num := range m.Chainlist {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintTx(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetSupportedChainsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetSupportedChainsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetSupportedChainsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSetSupportedChains) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Chainlist) > 0 {
		l = 0
		for _, e := range m.Chainlist {
			l += sovTx(uint64(e))
		}
		n += 1 + sovTx(uint64(l)) + l
	}
	return n
}

func (m *MsgSetSupportedChainsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSetSupportedChains) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetSupportedChains: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetSupportedChains: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v ObserverChain
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= ObserverChain(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Chainlist = append(m.Chainlist, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTx
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTx
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Chainlist) == 0 {
					m.Chainlist = make([]ObserverChain, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v ObserverChain
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTx
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= ObserverChain(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Chainlist = append(m.Chainlist, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Chainlist", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSetSupportedChainsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetSupportedChainsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetSupportedChainsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
