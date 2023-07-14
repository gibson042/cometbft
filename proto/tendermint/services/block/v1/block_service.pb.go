// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/services/block/v1/block_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

func init() {
	proto.RegisterFile("tendermint/services/block/v1/block_service.proto", fileDescriptor_1488dadaa3ae41e3)
}

var fileDescriptor_1488dadaa3ae41e3 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x28, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6,
	0x4f, 0xca, 0xc9, 0x4f, 0xce, 0xd6, 0x2f, 0x33, 0x84, 0x30, 0xe2, 0xa1, 0xe2, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0x32, 0x08, 0x1d, 0x7a, 0x30, 0x1d, 0x7a, 0x60, 0x85, 0x7a, 0x65, 0x86,
	0x52, 0x1a, 0x84, 0xcd, 0x83, 0x98, 0x63, 0xd4, 0xc3, 0xc8, 0xc5, 0xe3, 0x04, 0xe2, 0x07, 0x43,
	0x94, 0x09, 0xd5, 0x70, 0x71, 0xbb, 0xa7, 0x96, 0x38, 0x55, 0x7a, 0xa4, 0x66, 0xa6, 0x67, 0x94,
	0x08, 0x99, 0xea, 0xe1, 0xb3, 0x48, 0x0f, 0xa4, 0x14, 0xc4, 0x86, 0xa9, 0x0f, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0x91, 0x32, 0x23, 0x55, 0x5b, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x53, 0xe4,
	0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c,
	0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xd9, 0xa7, 0x67, 0x96, 0x64, 0x94,
	0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0xe7, 0xe7, 0xa6, 0x96, 0x24, 0xa5, 0x95, 0x20, 0x18,
	0x60, 0xcf, 0xe8, 0xe3, 0xf3, 0x75, 0x12, 0x1b, 0x58, 0x8d, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0xf8, 0xb1, 0x06, 0x00, 0x6c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlockServiceClient is the client API for BlockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockServiceClient interface {
	// GetBlock retrieves the block information at a particular height,
	// if height '0' (zero) is specified it returns the latest height
	GetByHeight(ctx context.Context, in *GetBlockByHeightRequest, opts ...grpc.CallOption) (*GetBlockByHeightResponse, error)
}

type blockServiceClient struct {
	cc grpc1.ClientConn
}

func NewBlockServiceClient(cc grpc1.ClientConn) BlockServiceClient {
	return &blockServiceClient{cc}
}

func (c *blockServiceClient) GetByHeight(ctx context.Context, in *GetBlockByHeightRequest, opts ...grpc.CallOption) (*GetBlockByHeightResponse, error) {
	out := new(GetBlockByHeightResponse)
	err := c.cc.Invoke(ctx, "/tendermint.services.block.v1.BlockService/GetByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockServiceServer is the server API for BlockService service.
type BlockServiceServer interface {
	// GetBlock retrieves the block information at a particular height,
	// if height '0' (zero) is specified it returns the latest height
	GetByHeight(context.Context, *GetBlockByHeightRequest) (*GetBlockByHeightResponse, error)
}

// UnimplementedBlockServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlockServiceServer struct {
}

func (*UnimplementedBlockServiceServer) GetByHeight(ctx context.Context, req *GetBlockByHeightRequest) (*GetBlockByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByHeight not implemented")
}

func RegisterBlockServiceServer(s grpc1.Server, srv BlockServiceServer) {
	s.RegisterService(&_BlockService_serviceDesc, srv)
}

func _BlockService_GetByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockServiceServer).GetByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.services.block.v1.BlockService/GetByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockServiceServer).GetByHeight(ctx, req.(*GetBlockByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlockService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tendermint.services.block.v1.BlockService",
	HandlerType: (*BlockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByHeight",
			Handler:    _BlockService_GetByHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tendermint/services/block/v1/block_service.proto",
}
