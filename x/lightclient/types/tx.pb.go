// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lightclient/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgUpdateVerificationFlags struct {
	Creator           string            `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	VerificationFlags VerificationFlags `protobuf:"bytes,2,opt,name=verification_flags,json=verificationFlags,proto3" json:"verification_flags"`
}

func (m *MsgUpdateVerificationFlags) Reset()         { *m = MsgUpdateVerificationFlags{} }
func (m *MsgUpdateVerificationFlags) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateVerificationFlags) ProtoMessage()    {}
func (*MsgUpdateVerificationFlags) Descriptor() ([]byte, []int) {
	return fileDescriptor_81fed8987f08d9c5, []int{0}
}
func (m *MsgUpdateVerificationFlags) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateVerificationFlags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateVerificationFlags.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateVerificationFlags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateVerificationFlags.Merge(m, src)
}
func (m *MsgUpdateVerificationFlags) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateVerificationFlags) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateVerificationFlags.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateVerificationFlags proto.InternalMessageInfo

func (m *MsgUpdateVerificationFlags) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateVerificationFlags) GetVerificationFlags() VerificationFlags {
	if m != nil {
		return m.VerificationFlags
	}
	return VerificationFlags{}
}

type MsgUpdateVerificationFlagsResponse struct {
}

func (m *MsgUpdateVerificationFlagsResponse) Reset()         { *m = MsgUpdateVerificationFlagsResponse{} }
func (m *MsgUpdateVerificationFlagsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateVerificationFlagsResponse) ProtoMessage()    {}
func (*MsgUpdateVerificationFlagsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_81fed8987f08d9c5, []int{1}
}
func (m *MsgUpdateVerificationFlagsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateVerificationFlagsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateVerificationFlagsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateVerificationFlagsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateVerificationFlagsResponse.Merge(m, src)
}
func (m *MsgUpdateVerificationFlagsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateVerificationFlagsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateVerificationFlagsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateVerificationFlagsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateVerificationFlags)(nil), "zetachain.zetacore.lightclient.MsgUpdateVerificationFlags")
	proto.RegisterType((*MsgUpdateVerificationFlagsResponse)(nil), "zetachain.zetacore.lightclient.MsgUpdateVerificationFlagsResponse")
}

func init() { proto.RegisterFile("lightclient/tx.proto", fileDescriptor_81fed8987f08d9c5) }

var fileDescriptor_81fed8987f08d9c5 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0xc9, 0x4c, 0xcf,
	0x28, 0x49, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0xd1, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0xab, 0x4a, 0x2d, 0x49, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x03, 0xb3, 0xf2, 0x8b,
	0x52, 0xf5, 0x90, 0x14, 0x4a, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x95, 0xea, 0x83, 0x58, 0x10,
	0x5d, 0x52, 0x2a, 0xc8, 0x66, 0x95, 0xa5, 0x16, 0x65, 0xa6, 0x65, 0x26, 0x27, 0x96, 0x64, 0xe6,
	0xe7, 0xc5, 0xa7, 0xe5, 0x24, 0xa6, 0x17, 0x43, 0x54, 0x29, 0xcd, 0x63, 0xe4, 0x92, 0xf2, 0x2d,
	0x4e, 0x0f, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0x0d, 0x43, 0x52, 0xe5, 0x06, 0x52, 0x24, 0x24, 0xc1,
	0xc5, 0x9e, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0x5f, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04,
	0xe3, 0x0a, 0xa5, 0x71, 0x09, 0x61, 0x1a, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0x64, 0xa8,
	0x87, 0xdf, 0xc5, 0x7a, 0x18, 0x16, 0x39, 0xb1, 0x9c, 0xb8, 0x27, 0xcf, 0x10, 0x24, 0x58, 0x86,
	0x2e, 0xa1, 0xa4, 0xc2, 0xa5, 0x84, 0xdb, 0x7d, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9,
	0x46, 0x0b, 0x19, 0xb9, 0x98, 0x7d, 0x8b, 0xd3, 0x85, 0x66, 0x32, 0x72, 0x89, 0xe3, 0xf2, 0x8b,
	0x15, 0x21, 0x57, 0xe1, 0xb6, 0x47, 0xca, 0x89, 0x7c, 0xbd, 0x30, 0x37, 0x3a, 0xf9, 0x9c, 0x78,
	0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c,
	0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x51, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92,
	0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xc8, 0x74, 0x5d, 0xb0, 0x45, 0xfa, 0x30, 0x8b, 0xf4, 0x2b, 0xf4,
	0x51, 0xd2, 0x45, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0xfe, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x5d, 0x1d, 0xfc, 0x77, 0x33, 0x02, 0x00, 0x00,
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
	UpdateVerificationFlags(ctx context.Context, in *MsgUpdateVerificationFlags, opts ...grpc.CallOption) (*MsgUpdateVerificationFlagsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateVerificationFlags(ctx context.Context, in *MsgUpdateVerificationFlags, opts ...grpc.CallOption) (*MsgUpdateVerificationFlagsResponse, error) {
	out := new(MsgUpdateVerificationFlagsResponse)
	err := c.cc.Invoke(ctx, "/zetachain.zetacore.lightclient.Msg/UpdateVerificationFlags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	UpdateVerificationFlags(context.Context, *MsgUpdateVerificationFlags) (*MsgUpdateVerificationFlagsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateVerificationFlags(ctx context.Context, req *MsgUpdateVerificationFlags) (*MsgUpdateVerificationFlagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVerificationFlags not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateVerificationFlags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateVerificationFlags)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateVerificationFlags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zetachain.zetacore.lightclient.Msg/UpdateVerificationFlags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateVerificationFlags(ctx, req.(*MsgUpdateVerificationFlags))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zetachain.zetacore.lightclient.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateVerificationFlags",
			Handler:    _Msg_UpdateVerificationFlags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lightclient/tx.proto",
}

func (m *MsgUpdateVerificationFlags) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateVerificationFlags) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateVerificationFlags) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.VerificationFlags.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateVerificationFlagsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateVerificationFlagsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateVerificationFlagsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgUpdateVerificationFlags) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.VerificationFlags.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateVerificationFlagsResponse) Size() (n int) {
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
func (m *MsgUpdateVerificationFlags) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateVerificationFlags: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateVerificationFlags: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerificationFlags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VerificationFlags.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *MsgUpdateVerificationFlagsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateVerificationFlagsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateVerificationFlagsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
