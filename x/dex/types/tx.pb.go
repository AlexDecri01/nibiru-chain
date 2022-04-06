// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgCreatePool struct {
	Creator    string      `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	PoolParams *PoolParams `protobuf:"bytes,2,opt,name=poolParams,proto3" json:"poolParams,omitempty" yaml:"pool_params"`
	PoolAssets []PoolAsset `protobuf:"bytes,3,rep,name=poolAssets,proto3" json:"poolAssets"`
}

func (m *MsgCreatePool) Reset()         { *m = MsgCreatePool{} }
func (m *MsgCreatePool) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePool) ProtoMessage()    {}
func (*MsgCreatePool) Descriptor() ([]byte, []int) {
	return fileDescriptor_18e8aa85ff669608, []int{0}
}
func (m *MsgCreatePool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePool.Merge(m, src)
}
func (m *MsgCreatePool) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePool proto.InternalMessageInfo

func (m *MsgCreatePool) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreatePool) GetPoolParams() *PoolParams {
	if m != nil {
		return m.PoolParams
	}
	return nil
}

func (m *MsgCreatePool) GetPoolAssets() []PoolAsset {
	if m != nil {
		return m.PoolAssets
	}
	return nil
}

type MsgCreatePoolResponse struct {
	PoolId uint64 `protobuf:"varint,1,opt,name=poolId,proto3" json:"poolId,omitempty"`
}

func (m *MsgCreatePoolResponse) Reset()         { *m = MsgCreatePoolResponse{} }
func (m *MsgCreatePoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePoolResponse) ProtoMessage()    {}
func (*MsgCreatePoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_18e8aa85ff669608, []int{1}
}
func (m *MsgCreatePoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePoolResponse.Merge(m, src)
}
func (m *MsgCreatePoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePoolResponse proto.InternalMessageInfo

func (m *MsgCreatePoolResponse) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

//
//Message to join a pool (identified by poolId) with a set of tokens to deposit.
type MsgJoinPool struct {
	Sender   string       `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	PoolId   uint64       `protobuf:"varint,2,opt,name=poolId,proto3" json:"poolId,omitempty" yaml:"pool_id"`
	TokensIn []types.Coin `protobuf:"bytes,3,rep,name=tokensIn,proto3" json:"tokensIn" yaml:"tokens_in"`
}

func (m *MsgJoinPool) Reset()         { *m = MsgJoinPool{} }
func (m *MsgJoinPool) String() string { return proto.CompactTextString(m) }
func (*MsgJoinPool) ProtoMessage()    {}
func (*MsgJoinPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_18e8aa85ff669608, []int{2}
}
func (m *MsgJoinPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgJoinPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgJoinPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgJoinPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgJoinPool.Merge(m, src)
}
func (m *MsgJoinPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgJoinPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgJoinPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgJoinPool proto.InternalMessageInfo

func (m *MsgJoinPool) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgJoinPool) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *MsgJoinPool) GetTokensIn() []types.Coin {
	if m != nil {
		return m.TokensIn
	}
	return nil
}

//
//Response when a user joins a pool.
type MsgJoinPoolResponse struct {
	// the final state of the pool after a join
	Pool *Pool `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool,omitempty"`
	// sum of LP tokens minted from the join
	NumPoolSharesOut types.Coin `protobuf:"bytes,2,opt,name=numPoolSharesOut,proto3" json:"numPoolSharesOut" yaml:"num_pool_shares_out"`
	// remaining tokens from attempting to join the pool
	RemainingCoins []types.Coin `protobuf:"bytes,3,rep,name=remainingCoins,proto3" json:"remainingCoins" yaml:"tokens_in"`
}

func (m *MsgJoinPoolResponse) Reset()         { *m = MsgJoinPoolResponse{} }
func (m *MsgJoinPoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgJoinPoolResponse) ProtoMessage()    {}
func (*MsgJoinPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_18e8aa85ff669608, []int{3}
}
func (m *MsgJoinPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgJoinPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgJoinPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgJoinPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgJoinPoolResponse.Merge(m, src)
}
func (m *MsgJoinPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgJoinPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgJoinPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgJoinPoolResponse proto.InternalMessageInfo

func (m *MsgJoinPoolResponse) GetPool() *Pool {
	if m != nil {
		return m.Pool
	}
	return nil
}

func (m *MsgJoinPoolResponse) GetNumPoolSharesOut() types.Coin {
	if m != nil {
		return m.NumPoolSharesOut
	}
	return types.Coin{}
}

func (m *MsgJoinPoolResponse) GetRemainingCoins() []types.Coin {
	if m != nil {
		return m.RemainingCoins
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgCreatePool)(nil), "matrix.dex.v1.MsgCreatePool")
	proto.RegisterType((*MsgCreatePoolResponse)(nil), "matrix.dex.v1.MsgCreatePoolResponse")
	proto.RegisterType((*MsgJoinPool)(nil), "matrix.dex.v1.MsgJoinPool")
	proto.RegisterType((*MsgJoinPoolResponse)(nil), "matrix.dex.v1.MsgJoinPoolResponse")
}

func init() { proto.RegisterFile("dex/v1/tx.proto", fileDescriptor_18e8aa85ff669608) }

var fileDescriptor_18e8aa85ff669608 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xbf, 0x6e, 0x13, 0x31,
	0x1c, 0xc7, 0x73, 0x6d, 0x14, 0x8a, 0xa3, 0x94, 0xd6, 0x81, 0xea, 0x7a, 0x42, 0x97, 0xc8, 0x42,
	0x6a, 0x60, 0x38, 0x2b, 0x61, 0x63, 0x40, 0xea, 0x95, 0xa5, 0x95, 0xa2, 0x86, 0x63, 0x83, 0x21,
	0x72, 0x12, 0xeb, 0x6a, 0x91, 0xb3, 0xa3, 0xb3, 0x13, 0xa5, 0x2f, 0xc0, 0xcc, 0x3b, 0xf0, 0x0c,
	0x6c, 0x3c, 0x40, 0xc7, 0x8e, 0x4c, 0x11, 0x4a, 0xde, 0x20, 0x4f, 0x80, 0x6c, 0x5f, 0x4e, 0x97,
	0xfe, 0x61, 0x61, 0xb3, 0x7f, 0x7f, 0x3e, 0xfe, 0x7e, 0xfd, 0xb3, 0xc1, 0xb3, 0x11, 0x9d, 0xe3,
	0x59, 0x1b, 0xab, 0x79, 0x30, 0x49, 0x85, 0x12, 0xb0, 0x96, 0x10, 0x95, 0xb2, 0x79, 0x30, 0xa2,
	0xf3, 0x60, 0xd6, 0xf6, 0xea, 0x59, 0x7e, 0x42, 0x52, 0x92, 0x48, 0x5b, 0xe3, 0x1d, 0x6e, 0x82,
	0x42, 0x8c, 0xb3, 0xd0, 0xf3, 0x58, 0xc4, 0xc2, 0x2c, 0xb1, 0x5e, 0x65, 0x51, 0x7f, 0x28, 0x64,
	0x22, 0x24, 0x1e, 0x10, 0x49, 0xf1, 0xac, 0x3d, 0xa0, 0x8a, 0xb4, 0xf1, 0x50, 0x30, 0x6e, 0xf3,
	0xe8, 0x97, 0x03, 0x6a, 0x5d, 0x19, 0x9f, 0xa5, 0x94, 0x28, 0xda, 0x13, 0x62, 0x0c, 0x5d, 0xf0,
	0x64, 0xa8, 0x77, 0x22, 0x75, 0x9d, 0xa6, 0xd3, 0x7a, 0x1a, 0x6d, 0xb6, 0xf0, 0x23, 0x00, 0xfa,
	0xbc, 0x9e, 0x11, 0xe2, 0xee, 0x34, 0x9d, 0x56, 0xb5, 0x73, 0x1c, 0x6c, 0xa9, 0x0d, 0x7a, 0x79,
	0x41, 0x78, 0xb4, 0x5e, 0x34, 0xe0, 0x35, 0x49, 0xc6, 0xef, 0x90, 0x6e, 0xeb, 0x5b, 0x03, 0x28,
	0x2a, 0x40, 0xe0, 0x7b, 0x8b, 0x3c, 0x95, 0x92, 0x2a, 0xe9, 0xee, 0x36, 0x77, 0x5b, 0xd5, 0x8e,
	0xfb, 0x00, 0xd2, 0x14, 0x84, 0xe5, 0x9b, 0x45, 0xa3, 0x14, 0x15, 0x3a, 0x10, 0x06, 0x2f, 0xb6,
	0xd4, 0x47, 0x54, 0x4e, 0x04, 0x97, 0x14, 0x1e, 0x81, 0x8a, 0x2e, 0x3b, 0x1f, 0x19, 0x13, 0xe5,
	0x28, 0xdb, 0xa1, 0x9f, 0x0e, 0xa8, 0x76, 0x65, 0x7c, 0x21, 0x18, 0x37, 0x6e, 0x5f, 0x83, 0x8a,
	0xa4, 0x7c, 0x44, 0x33, 0xb3, 0xe1, 0xe1, 0x7a, 0xd1, 0xa8, 0x59, 0xd1, 0x36, 0x8e, 0xa2, 0xac,
	0x00, 0xbe, 0xc9, 0x91, 0xda, 0x7a, 0x39, 0x84, 0xeb, 0x45, 0x63, 0xbf, 0xe0, 0x8f, 0x8d, 0xd0,
	0xe6, 0x18, 0x78, 0x09, 0xf6, 0x94, 0xf8, 0x4a, 0xb9, 0x3c, 0xe7, 0x99, 0xab, 0xe3, 0xc0, 0x4e,
	0x22, 0xd0, 0x93, 0x08, 0xb2, 0x49, 0x04, 0x67, 0x82, 0xf1, 0xd0, 0xd5, 0xb6, 0xd6, 0x8b, 0xc6,
	0x81, 0x85, 0xd9, 0xc6, 0x3e, 0xe3, 0x28, 0xca, 0x21, 0xe8, 0xdb, 0x0e, 0xa8, 0x17, 0x74, 0xe7,
	0x3e, 0x4f, 0x40, 0x59, 0x1f, 0x69, 0xd4, 0x57, 0x3b, 0xf5, 0x07, 0xae, 0x2e, 0x32, 0x05, 0x90,
	0x81, 0x03, 0x3e, 0x4d, 0x74, 0xe0, 0xd3, 0x15, 0x49, 0xa9, 0xbc, 0x9c, 0xaa, 0x7c, 0x84, 0x8f,
	0x2a, 0x43, 0x99, 0x32, 0xcf, 0x2a, 0xe3, 0xd3, 0xa4, 0x6f, 0xac, 0x4a, 0x83, 0xe8, 0x8b, 0xa9,
	0x42, 0xd1, 0x3d, 0x2c, 0xfc, 0x02, 0xf6, 0x53, 0x9a, 0x10, 0xc6, 0x19, 0x8f, 0x35, 0x46, 0xfe,
	0xcf, 0x15, 0xdc, 0x41, 0x75, 0x7e, 0x38, 0x60, 0xb7, 0x2b, 0x63, 0xd8, 0x03, 0xa0, 0xf0, 0x68,
	0x5f, 0xde, 0x31, 0xbe, 0xf5, 0x28, 0xbc, 0x57, 0xff, 0xca, 0xe6, 0x57, 0x79, 0x01, 0xf6, 0xf2,
	0x67, 0xe1, 0xdd, 0xef, 0xd8, 0xe4, 0x3c, 0xf4, 0x78, 0x6e, 0xc3, 0x0a, 0x4f, 0x6f, 0x96, 0xbe,
	0x73, 0xbb, 0xf4, 0x9d, 0x3f, 0x4b, 0xdf, 0xf9, 0xbe, 0xf2, 0x4b, 0xb7, 0x2b, 0xbf, 0xf4, 0x7b,
	0xe5, 0x97, 0x3e, 0x9f, 0xc4, 0x4c, 0x5d, 0x4d, 0x07, 0xc1, 0x50, 0x24, 0xb8, 0x6b, 0x38, 0x1f,
	0x88, 0xc0, 0x96, 0x88, 0xe7, 0x58, 0xff, 0x6b, 0x75, 0x3d, 0xa1, 0x72, 0x50, 0x31, 0x1f, 0xf4,
	0xed, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x38, 0x60, 0x1b, 0x20, 0x04, 0x00, 0x00,
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
	// Used to create a pool.
	CreatePool(ctx context.Context, in *MsgCreatePool, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error)
	// Join a pool as a liquidity provider.
	JoinPool(ctx context.Context, in *MsgJoinPool, opts ...grpc.CallOption) (*MsgJoinPoolResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreatePool(ctx context.Context, in *MsgCreatePool, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error) {
	out := new(MsgCreatePoolResponse)
	err := c.cc.Invoke(ctx, "/matrix.dex.v1.Msg/CreatePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) JoinPool(ctx context.Context, in *MsgJoinPool, opts ...grpc.CallOption) (*MsgJoinPoolResponse, error) {
	out := new(MsgJoinPoolResponse)
	err := c.cc.Invoke(ctx, "/matrix.dex.v1.Msg/JoinPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Used to create a pool.
	CreatePool(context.Context, *MsgCreatePool) (*MsgCreatePoolResponse, error)
	// Join a pool as a liquidity provider.
	JoinPool(context.Context, *MsgJoinPool) (*MsgJoinPoolResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreatePool(ctx context.Context, req *MsgCreatePool) (*MsgCreatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePool not implemented")
}
func (*UnimplementedMsgServer) JoinPool(ctx context.Context, req *MsgJoinPool) (*MsgJoinPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinPool not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matrix.dex.v1.Msg/CreatePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePool(ctx, req.(*MsgCreatePool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_JoinPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgJoinPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).JoinPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matrix.dex.v1.Msg/JoinPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).JoinPool(ctx, req.(*MsgJoinPool))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "matrix.dex.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePool",
			Handler:    _Msg_CreatePool_Handler,
		},
		{
			MethodName: "JoinPool",
			Handler:    _Msg_JoinPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dex/v1/tx.proto",
}

func (m *MsgCreatePool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolAssets) > 0 {
		for iNdEx := len(m.PoolAssets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolAssets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.PoolParams != nil {
		{
			size, err := m.PoolParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
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

func (m *MsgCreatePoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgJoinPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgJoinPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgJoinPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokensIn) > 0 {
		for iNdEx := len(m.TokensIn) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokensIn[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.PoolId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgJoinPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgJoinPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgJoinPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RemainingCoins) > 0 {
		for iNdEx := len(m.RemainingCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RemainingCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.NumPoolSharesOut.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Pool != nil {
		{
			size, err := m.Pool.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
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
func (m *MsgCreatePool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PoolParams != nil {
		l = m.PoolParams.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.PoolAssets) > 0 {
		for _, e := range m.PoolAssets {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgCreatePoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovTx(uint64(m.PoolId))
	}
	return n
}

func (m *MsgJoinPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PoolId != 0 {
		n += 1 + sovTx(uint64(m.PoolId))
	}
	if len(m.TokensIn) > 0 {
		for _, e := range m.TokensIn {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgJoinPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pool != nil {
		l = m.Pool.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.NumPoolSharesOut.Size()
	n += 1 + l + sovTx(uint64(l))
	if len(m.RemainingCoins) > 0 {
		for _, e := range m.RemainingCoins {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreatePool) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreatePool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePool: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field PoolParams", wireType)
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
			if m.PoolParams == nil {
				m.PoolParams = &PoolParams{}
			}
			if err := m.PoolParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAssets", wireType)
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
			m.PoolAssets = append(m.PoolAssets, PoolAsset{})
			if err := m.PoolAssets[len(m.PoolAssets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgCreatePoolResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreatePoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
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
func (m *MsgJoinPool) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgJoinPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgJoinPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokensIn", wireType)
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
			m.TokensIn = append(m.TokensIn, types.Coin{})
			if err := m.TokensIn[len(m.TokensIn)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgJoinPoolResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgJoinPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgJoinPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
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
			if m.Pool == nil {
				m.Pool = &Pool{}
			}
			if err := m.Pool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumPoolSharesOut", wireType)
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
			if err := m.NumPoolSharesOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainingCoins", wireType)
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
			m.RemainingCoins = append(m.RemainingCoins, types.Coin{})
			if err := m.RemainingCoins[len(m.RemainingCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
