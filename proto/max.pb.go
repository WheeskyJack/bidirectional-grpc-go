// Code generated by protoc-gen-go. DO NOT EDIT.
// source: max.proto

package service_findint_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	// inp is the input integer to server
	Inp                  int32    `protobuf:"varint,1,opt,name=inp,proto3" json:"inp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd713ec412f37532, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetInp() int32 {
	if m != nil {
		return m.Inp
	}
	return 0
}

type Response struct {
	// out is the max integer server returns
	Out                  int32    `protobuf:"varint,1,opt,name=out,proto3" json:"out,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd713ec412f37532, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetOut() int32 {
	if m != nil {
		return m.Out
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "service.findint.v1.Request")
	proto.RegisterType((*Response)(nil), "service.findint.v1.Response")
}

func init() { proto.RegisterFile("max.proto", fileDescriptor_dd713ec412f37532) }

var fileDescriptor_dd713ec412f37532 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x4d, 0xac, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4b,
	0xcb, 0xcc, 0x4b, 0xc9, 0xcc, 0x2b, 0xd1, 0x2b, 0x33, 0x54, 0x92, 0xe6, 0x62, 0x0f, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe0, 0x62, 0xce, 0xcc, 0x2b, 0x90, 0x60, 0x54, 0x60, 0xd4,
	0x60, 0x0d, 0x02, 0x31, 0x95, 0x64, 0xb8, 0x38, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53,
	0x41, 0xb2, 0xf9, 0xa5, 0x25, 0x30, 0xd9, 0xfc, 0xd2, 0x12, 0x23, 0x1f, 0x2e, 0x16, 0x90, 0x41,
	0x42, 0x2e, 0x5c, 0xcc, 0xb9, 0x89, 0x15, 0x42, 0xd2, 0x7a, 0x98, 0xc6, 0xeb, 0x41, 0xcd, 0x96,
	0x92, 0xc1, 0x2e, 0x09, 0x31, 0x5b, 0x83, 0xd1, 0x80, 0xd1, 0x49, 0x24, 0x0a, 0xe6, 0xbc, 0x78,
	0xa8, 0x92, 0xf8, 0x32, 0xc3, 0x24, 0x36, 0xb0, 0xcb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x5f, 0xa4, 0x88, 0xfc, 0xc6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FindClient is the client API for Find service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FindClient interface {
	// max finds the max number provided to server
	Max(ctx context.Context, opts ...grpc.CallOption) (Find_MaxClient, error)
}

type findClient struct {
	cc *grpc.ClientConn
}

func NewFindClient(cc *grpc.ClientConn) FindClient {
	return &findClient{cc}
}

func (c *findClient) Max(ctx context.Context, opts ...grpc.CallOption) (Find_MaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Find_serviceDesc.Streams[0], "/service.findint.v1.find/max", opts...)
	if err != nil {
		return nil, err
	}
	x := &findMaxClient{stream}
	return x, nil
}

type Find_MaxClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type findMaxClient struct {
	grpc.ClientStream
}

func (x *findMaxClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *findMaxClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FindServer is the server API for Find service.
type FindServer interface {
	// max finds the max number provided to server
	Max(Find_MaxServer) error
}

// UnimplementedFindServer can be embedded to have forward compatible implementations.
type UnimplementedFindServer struct {
}

func (*UnimplementedFindServer) Max(srv Find_MaxServer) error {
	return status.Errorf(codes.Unimplemented, "method Max not implemented")
}

func RegisterFindServer(s *grpc.Server, srv FindServer) {
	s.RegisterService(&_Find_serviceDesc, srv)
}

func _Find_Max_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FindServer).Max(&findMaxServer{stream})
}

type Find_MaxServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type findMaxServer struct {
	grpc.ServerStream
}

func (x *findMaxServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *findMaxServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Find_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.findint.v1.find",
	HandlerType: (*FindServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "max",
			Handler:       _Find_Max_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "max.proto",
}
