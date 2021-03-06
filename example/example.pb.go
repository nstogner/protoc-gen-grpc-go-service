// Code generated by protoc-gen-go.
// source: example.proto
// DO NOT EDIT!

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	example.proto

It has these top-level messages:
	InputMessage
	OutputMessage
*/
package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InputMessage struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *InputMessage) Reset()                    { *m = InputMessage{} }
func (m *InputMessage) String() string            { return proto.CompactTextString(m) }
func (*InputMessage) ProtoMessage()               {}
func (*InputMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type OutputMessage struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *OutputMessage) Reset()                    { *m = OutputMessage{} }
func (m *OutputMessage) String() string            { return proto.CompactTextString(m) }
func (*OutputMessage) ProtoMessage()               {}
func (*OutputMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*InputMessage)(nil), "example.InputMessage")
	proto.RegisterType((*OutputMessage)(nil), "example.OutputMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ExampleService service

type ExampleServiceClient interface {
	Echo(ctx context.Context, in *InputMessage, opts ...grpc.CallOption) (*OutputMessage, error)
	EchoStreamOut(ctx context.Context, in *InputMessage, opts ...grpc.CallOption) (ExampleService_EchoStreamOutClient, error)
	EchoStreamIn(ctx context.Context, opts ...grpc.CallOption) (ExampleService_EchoStreamInClient, error)
	EchoStreamInOut(ctx context.Context, opts ...grpc.CallOption) (ExampleService_EchoStreamInOutClient, error)
}

type exampleServiceClient struct {
	cc *grpc.ClientConn
}

func NewExampleServiceClient(cc *grpc.ClientConn) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) Echo(ctx context.Context, in *InputMessage, opts ...grpc.CallOption) (*OutputMessage, error) {
	out := new(OutputMessage)
	err := grpc.Invoke(ctx, "/example.ExampleService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) EchoStreamOut(ctx context.Context, in *InputMessage, opts ...grpc.CallOption) (ExampleService_EchoStreamOutClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ExampleService_serviceDesc.Streams[0], c.cc, "/example.ExampleService/EchoStreamOut", opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleServiceEchoStreamOutClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExampleService_EchoStreamOutClient interface {
	Recv() (*OutputMessage, error)
	grpc.ClientStream
}

type exampleServiceEchoStreamOutClient struct {
	grpc.ClientStream
}

func (x *exampleServiceEchoStreamOutClient) Recv() (*OutputMessage, error) {
	m := new(OutputMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exampleServiceClient) EchoStreamIn(ctx context.Context, opts ...grpc.CallOption) (ExampleService_EchoStreamInClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ExampleService_serviceDesc.Streams[1], c.cc, "/example.ExampleService/EchoStreamIn", opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleServiceEchoStreamInClient{stream}
	return x, nil
}

type ExampleService_EchoStreamInClient interface {
	Send(*InputMessage) error
	CloseAndRecv() (*OutputMessage, error)
	grpc.ClientStream
}

type exampleServiceEchoStreamInClient struct {
	grpc.ClientStream
}

func (x *exampleServiceEchoStreamInClient) Send(m *InputMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *exampleServiceEchoStreamInClient) CloseAndRecv() (*OutputMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(OutputMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exampleServiceClient) EchoStreamInOut(ctx context.Context, opts ...grpc.CallOption) (ExampleService_EchoStreamInOutClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ExampleService_serviceDesc.Streams[2], c.cc, "/example.ExampleService/EchoStreamInOut", opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleServiceEchoStreamInOutClient{stream}
	return x, nil
}

type ExampleService_EchoStreamInOutClient interface {
	Send(*InputMessage) error
	Recv() (*OutputMessage, error)
	grpc.ClientStream
}

type exampleServiceEchoStreamInOutClient struct {
	grpc.ClientStream
}

func (x *exampleServiceEchoStreamInOutClient) Send(m *InputMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *exampleServiceEchoStreamInOutClient) Recv() (*OutputMessage, error) {
	m := new(OutputMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ExampleService service

type ExampleServiceServer interface {
	Echo(context.Context, *InputMessage) (*OutputMessage, error)
	EchoStreamOut(*InputMessage, ExampleService_EchoStreamOutServer) error
	EchoStreamIn(ExampleService_EchoStreamInServer) error
	EchoStreamInOut(ExampleService_EchoStreamInOutServer) error
}

func RegisterExampleServiceServer(s *grpc.Server, srv ExampleServiceServer) {
	s.RegisterService(&_ExampleService_serviceDesc, srv)
}

func _ExampleService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.ExampleService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).Echo(ctx, req.(*InputMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_EchoStreamOut_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InputMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExampleServiceServer).EchoStreamOut(m, &exampleServiceEchoStreamOutServer{stream})
}

type ExampleService_EchoStreamOutServer interface {
	Send(*OutputMessage) error
	grpc.ServerStream
}

type exampleServiceEchoStreamOutServer struct {
	grpc.ServerStream
}

func (x *exampleServiceEchoStreamOutServer) Send(m *OutputMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _ExampleService_EchoStreamIn_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleServiceServer).EchoStreamIn(&exampleServiceEchoStreamInServer{stream})
}

type ExampleService_EchoStreamInServer interface {
	SendAndClose(*OutputMessage) error
	Recv() (*InputMessage, error)
	grpc.ServerStream
}

type exampleServiceEchoStreamInServer struct {
	grpc.ServerStream
}

func (x *exampleServiceEchoStreamInServer) SendAndClose(m *OutputMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *exampleServiceEchoStreamInServer) Recv() (*InputMessage, error) {
	m := new(InputMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ExampleService_EchoStreamInOut_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleServiceServer).EchoStreamInOut(&exampleServiceEchoStreamInOutServer{stream})
}

type ExampleService_EchoStreamInOutServer interface {
	Send(*OutputMessage) error
	Recv() (*InputMessage, error)
	grpc.ServerStream
}

type exampleServiceEchoStreamInOutServer struct {
	grpc.ServerStream
}

func (x *exampleServiceEchoStreamInOutServer) Send(m *OutputMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *exampleServiceEchoStreamInOutServer) Recv() (*InputMessage, error) {
	m := new(InputMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ExampleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _ExampleService_Echo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EchoStreamOut",
			Handler:       _ExampleService_EchoStreamOut_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EchoStreamIn",
			Handler:       _ExampleService_EchoStreamIn_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "EchoStreamInOut",
			Handler:       _ExampleService_EchoStreamInOut_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("example.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x54, 0xb8,
	0x78, 0x3c, 0xf3, 0x0a, 0x4a, 0x4b, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0x44, 0xb8,
	0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x25,
	0x55, 0x2e, 0x5e, 0xff, 0xd2, 0x12, 0x42, 0xca, 0x8c, 0x66, 0x30, 0x71, 0xf1, 0xb9, 0x42, 0x0c,
	0x0e, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x32, 0xe7, 0x62, 0x71, 0x4d, 0xce, 0xc8, 0x17,
	0x12, 0xd5, 0x83, 0x39, 0x00, 0xd9, 0x3a, 0x29, 0x31, 0xb8, 0x30, 0x8a, 0xf9, 0x4a, 0x0c, 0x42,
	0x4e, 0x5c, 0xbc, 0x20, 0x8d, 0xc1, 0x25, 0x45, 0xa9, 0x89, 0xb9, 0xfe, 0xa5, 0x25, 0x24, 0x9b,
	0x60, 0xc0, 0x28, 0xe4, 0xc8, 0xc5, 0x83, 0x30, 0xc3, 0x33, 0x8f, 0x64, 0x23, 0x34, 0x18, 0x85,
	0xdc, 0xb8, 0xf8, 0x91, 0x8d, 0x20, 0xc7, 0x21, 0x1a, 0x8c, 0x06, 0x8c, 0x49, 0x6c, 0xe0, 0x70,
	0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x34, 0xeb, 0x52, 0x88, 0x01, 0x00, 0x00,
}
