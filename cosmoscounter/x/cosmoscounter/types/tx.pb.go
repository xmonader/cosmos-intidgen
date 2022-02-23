// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmoscounter/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
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

func init() { proto.RegisterFile("cosmoscounter/tx.proto", fileDescriptor_2b2ba0fc77f7db5a) }

var fileDescriptor_2b2ba0fc77f7db5a = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x49, 0x2d, 0xd2, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x52, 0xa9, 0xc8, 0xcd, 0xcf, 0x4b, 0x4c, 0x49, 0x2d, 0xd2, 0x43, 0x51, 0x80,
	0xca, 0x33, 0x62, 0xe5, 0x62, 0xf6, 0x2d, 0x4e, 0x77, 0x0a, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2,
	0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1,
	0xc6, 0x63, 0x39, 0x86, 0x28, 0xf3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c,
	0x7d, 0x98, 0x89, 0xfa, 0xa8, 0x56, 0x56, 0xa0, 0xf1, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8,
	0xc0, 0xce, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x73, 0x64, 0xb4, 0x07, 0xa0, 0x00, 0x00,
	0x00,
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
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "xmonader.cosmoscounter.cosmoscounter.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "cosmoscounter/tx.proto",
}
