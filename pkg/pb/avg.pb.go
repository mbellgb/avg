// Code generated by protoc-gen-go. DO NOT EDIT.
// source: avg.proto

package avg

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

type IntArrayGreeting struct {
	Data                 []int32  `protobuf:"varint,1,rep,packed,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntArrayGreeting) Reset()         { *m = IntArrayGreeting{} }
func (m *IntArrayGreeting) String() string { return proto.CompactTextString(m) }
func (*IntArrayGreeting) ProtoMessage()    {}
func (*IntArrayGreeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_b889feec73713d05, []int{0}
}

func (m *IntArrayGreeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntArrayGreeting.Unmarshal(m, b)
}
func (m *IntArrayGreeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntArrayGreeting.Marshal(b, m, deterministic)
}
func (m *IntArrayGreeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntArrayGreeting.Merge(m, src)
}
func (m *IntArrayGreeting) XXX_Size() int {
	return xxx_messageInfo_IntArrayGreeting.Size(m)
}
func (m *IntArrayGreeting) XXX_DiscardUnknown() {
	xxx_messageInfo_IntArrayGreeting.DiscardUnknown(m)
}

var xxx_messageInfo_IntArrayGreeting proto.InternalMessageInfo

func (m *IntArrayGreeting) GetData() []int32 {
	if m != nil {
		return m.Data
	}
	return nil
}

type IntReply struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntReply) Reset()         { *m = IntReply{} }
func (m *IntReply) String() string { return proto.CompactTextString(m) }
func (*IntReply) ProtoMessage()    {}
func (*IntReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b889feec73713d05, []int{1}
}

func (m *IntReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntReply.Unmarshal(m, b)
}
func (m *IntReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntReply.Marshal(b, m, deterministic)
}
func (m *IntReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntReply.Merge(m, src)
}
func (m *IntReply) XXX_Size() int {
	return xxx_messageInfo_IntReply.Size(m)
}
func (m *IntReply) XXX_DiscardUnknown() {
	xxx_messageInfo_IntReply.DiscardUnknown(m)
}

var xxx_messageInfo_IntReply proto.InternalMessageInfo

func (m *IntReply) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type FloatArrayGreeting struct {
	Data                 []float32 `protobuf:"fixed32,1,rep,packed,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FloatArrayGreeting) Reset()         { *m = FloatArrayGreeting{} }
func (m *FloatArrayGreeting) String() string { return proto.CompactTextString(m) }
func (*FloatArrayGreeting) ProtoMessage()    {}
func (*FloatArrayGreeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_b889feec73713d05, []int{2}
}

func (m *FloatArrayGreeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FloatArrayGreeting.Unmarshal(m, b)
}
func (m *FloatArrayGreeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FloatArrayGreeting.Marshal(b, m, deterministic)
}
func (m *FloatArrayGreeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FloatArrayGreeting.Merge(m, src)
}
func (m *FloatArrayGreeting) XXX_Size() int {
	return xxx_messageInfo_FloatArrayGreeting.Size(m)
}
func (m *FloatArrayGreeting) XXX_DiscardUnknown() {
	xxx_messageInfo_FloatArrayGreeting.DiscardUnknown(m)
}

var xxx_messageInfo_FloatArrayGreeting proto.InternalMessageInfo

func (m *FloatArrayGreeting) GetData() []float32 {
	if m != nil {
		return m.Data
	}
	return nil
}

type FloatReply struct {
	Result               float32  `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FloatReply) Reset()         { *m = FloatReply{} }
func (m *FloatReply) String() string { return proto.CompactTextString(m) }
func (*FloatReply) ProtoMessage()    {}
func (*FloatReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b889feec73713d05, []int{3}
}

func (m *FloatReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FloatReply.Unmarshal(m, b)
}
func (m *FloatReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FloatReply.Marshal(b, m, deterministic)
}
func (m *FloatReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FloatReply.Merge(m, src)
}
func (m *FloatReply) XXX_Size() int {
	return xxx_messageInfo_FloatReply.Size(m)
}
func (m *FloatReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FloatReply.DiscardUnknown(m)
}

var xxx_messageInfo_FloatReply proto.InternalMessageInfo

func (m *FloatReply) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*IntArrayGreeting)(nil), "avg.IntArrayGreeting")
	proto.RegisterType((*IntReply)(nil), "avg.IntReply")
	proto.RegisterType((*FloatArrayGreeting)(nil), "avg.FloatArrayGreeting")
	proto.RegisterType((*FloatReply)(nil), "avg.FloatReply")
}

func init() { proto.RegisterFile("avg.proto", fileDescriptor_b889feec73713d05) }

var fileDescriptor_b889feec73713d05 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0x4b, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0x4b, 0x57, 0x52, 0xe3, 0x12, 0xf0, 0xcc,
	0x2b, 0x71, 0x2c, 0x2a, 0x4a, 0xac, 0x74, 0x2f, 0x4a, 0x4d, 0x2d, 0xc9, 0xcc, 0x4b, 0x17, 0x12,
	0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd6, 0x60, 0x0d, 0x02, 0xb3, 0x95,
	0x94, 0xb8, 0x38, 0x3c, 0xf3, 0x4a, 0x82, 0x52, 0x0b, 0x72, 0x2a, 0x85, 0xc4, 0xb8, 0xd8, 0x8a,
	0x52, 0x8b, 0x4b, 0x73, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xa0, 0x3c, 0x25, 0x0d,
	0x2e, 0x21, 0xb7, 0x9c, 0xfc, 0x44, 0x3c, 0xa6, 0x31, 0x41, 0x4d, 0x53, 0xe1, 0xe2, 0x02, 0xab,
	0xc4, 0x66, 0x1e, 0x13, 0xcc, 0x3c, 0xa3, 0x1c, 0x2e, 0x76, 0xb0, 0x29, 0xa9, 0x45, 0x42, 0x06,
	0x5c, 0x2c, 0xbe, 0xa9, 0x89, 0x79, 0x42, 0xa2, 0x7a, 0x20, 0xf7, 0xa3, 0xbb, 0x58, 0x8a, 0x1f,
	0x2c, 0x8c, 0x30, 0x52, 0x89, 0x41, 0x48, 0x9f, 0x8b, 0x35, 0x28, 0x31, 0x2f, 0x3d, 0x15, 0x97,
	0x16, 0x5e, 0x98, 0x30, 0x54, 0x43, 0x12, 0x1b, 0x38, 0x54, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x38, 0xd7, 0x03, 0x56, 0x22, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Calculates mean average
	Mean(ctx context.Context, in *IntArrayGreeting, opts ...grpc.CallOption) (*FloatReply, error)
	Range(ctx context.Context, in *IntArrayGreeting, opts ...grpc.CallOption) (*IntReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Mean(ctx context.Context, in *IntArrayGreeting, opts ...grpc.CallOption) (*FloatReply, error) {
	out := new(FloatReply)
	err := c.cc.Invoke(ctx, "/avg.Greeter/Mean", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Range(ctx context.Context, in *IntArrayGreeting, opts ...grpc.CallOption) (*IntReply, error) {
	out := new(IntReply)
	err := c.cc.Invoke(ctx, "/avg.Greeter/Range", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Calculates mean average
	Mean(context.Context, *IntArrayGreeting) (*FloatReply, error)
	Range(context.Context, *IntArrayGreeting) (*IntReply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) Mean(ctx context.Context, req *IntArrayGreeting) (*FloatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mean not implemented")
}
func (*UnimplementedGreeterServer) Range(ctx context.Context, req *IntArrayGreeting) (*IntReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Range not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Mean_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntArrayGreeting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Mean(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/avg.Greeter/Mean",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Mean(ctx, req.(*IntArrayGreeting))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Range_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntArrayGreeting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Range(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/avg.Greeter/Range",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Range(ctx, req.(*IntArrayGreeting))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "avg.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Mean",
			Handler:    _Greeter_Mean_Handler,
		},
		{
			MethodName: "Range",
			Handler:    _Greeter_Range_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "avg.proto",
}
