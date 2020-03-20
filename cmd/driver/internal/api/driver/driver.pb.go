// Code generated by protoc-gen-go. DO NOT EDIT.
// source: driver/driver.proto

package driver

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

type CommandRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandRequest) Reset()         { *m = CommandRequest{} }
func (m *CommandRequest) String() string { return proto.CompactTextString(m) }
func (*CommandRequest) ProtoMessage()    {}
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_651b72783301f7a9, []int{0}
}

func (m *CommandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandRequest.Unmarshal(m, b)
}
func (m *CommandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandRequest.Marshal(b, m, deterministic)
}
func (m *CommandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandRequest.Merge(m, src)
}
func (m *CommandRequest) XXX_Size() int {
	return xxx_messageInfo_CommandRequest.Size(m)
}
func (m *CommandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommandRequest proto.InternalMessageInfo

type CommandReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandReply) Reset()         { *m = CommandReply{} }
func (m *CommandReply) String() string { return proto.CompactTextString(m) }
func (*CommandReply) ProtoMessage()    {}
func (*CommandReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_651b72783301f7a9, []int{1}
}

func (m *CommandReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandReply.Unmarshal(m, b)
}
func (m *CommandReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandReply.Marshal(b, m, deterministic)
}
func (m *CommandReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandReply.Merge(m, src)
}
func (m *CommandReply) XXX_Size() int {
	return xxx_messageInfo_CommandReply.Size(m)
}
func (m *CommandReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandReply.DiscardUnknown(m)
}

var xxx_messageInfo_CommandReply proto.InternalMessageInfo

func (m *CommandReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CommandRequest)(nil), "driver.CommandRequest")
	proto.RegisterType((*CommandReply)(nil), "driver.CommandReply")
}

func init() {
	proto.RegisterFile("driver/driver.proto", fileDescriptor_651b72783301f7a9)
}

var fileDescriptor_651b72783301f7a9 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x29, 0xca, 0x2c,
	0x4b, 0x2d, 0xd2, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x92,
	0x00, 0x17, 0x9f, 0x73, 0x7e, 0x6e, 0x6e, 0x62, 0x5e, 0x4a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71,
	0x89, 0x92, 0x06, 0x17, 0x0f, 0x5c, 0xa4, 0x20, 0xa7, 0x52, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5,
	0xb8, 0x38, 0x31, 0x3d, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6, 0x35, 0x3a, 0xca,
	0xc2, 0xc5, 0xe6, 0x02, 0x36, 0x46, 0xc8, 0x8c, 0x8b, 0xc5, 0x27, 0x35, 0xad, 0x44, 0x48, 0x4c,
	0x0f, 0x6a, 0x0b, 0xaa, 0xa1, 0x52, 0x22, 0x18, 0xe2, 0x05, 0x39, 0x95, 0x4a, 0x0c, 0x42, 0x26,
	0x5c, 0x4c, 0xa1, 0x05, 0x24, 0xeb, 0x32, 0xe7, 0x62, 0x0d, 0xca, 0x4c, 0xcf, 0x20, 0xdd, 0x3a,
	0x33, 0x2e, 0x16, 0x97, 0xfc, 0xf2, 0x3c, 0x92, 0xf5, 0x59, 0x70, 0xb1, 0x85, 0x16, 0x90, 0xe5,
	0x41, 0x2b, 0x2e, 0x0e, 0x90, 0x8d, 0x64, 0xe9, 0xb5, 0xe4, 0x62, 0x0f, 0x2d, 0x20, 0xcf, 0xa3,
	0xd6, 0x5c, 0x9c, 0x20, 0x6b, 0xc9, 0x0e, 0xa5, 0xe0, 0x92, 0xfc, 0x02, 0x72, 0xf4, 0xb9, 0x65,
	0x16, 0xa5, 0x92, 0xaa, 0x2f, 0x89, 0x0d, 0x9c, 0x24, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xec, 0x48, 0x24, 0xd0, 0xa9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DriverClient is the client API for Driver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DriverClient interface {
	Left(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error)
	Up(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error)
	Right(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error)
	Down(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error)
	UpLeft(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error)
	DownLeft(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error)
	UpRight(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error)
	DownRight(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error)
	Stop(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error)
	Fire(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error)
}

type driverClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverClient(cc grpc.ClientConnInterface) DriverClient {
	return &driverClient{cc}
}

func (c *driverClient) Left(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error) {
	out := new(CommandReply)
	err := c.cc.Invoke(ctx, "/driver.Driver/Left", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Up(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error) {
	out := new(CommandReply)
	err := c.cc.Invoke(ctx, "/driver.Driver/Up", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Right(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error) {
	out := new(CommandReply)
	err := c.cc.Invoke(ctx, "/driver.Driver/Right", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Down(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error) {
	out := new(CommandReply)
	err := c.cc.Invoke(ctx, "/driver.Driver/Down", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) UpLeft(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error) {
	out := new(CommandReply)
	err := c.cc.Invoke(ctx, "/driver.Driver/UpLeft", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) DownLeft(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error) {
	out := new(CommandReply)
	err := c.cc.Invoke(ctx, "/driver.Driver/DownLeft", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) UpRight(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error) {
	out := new(CommandReply)
	err := c.cc.Invoke(ctx, "/driver.Driver/UpRight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) DownRight(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error) {
	out := new(CommandReply)
	err := c.cc.Invoke(ctx, "/driver.Driver/DownRight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Stop(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error) {
	out := new(CommandReply)
	err := c.cc.Invoke(ctx, "/driver.Driver/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Fire(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error) {
	out := new(CommandReply)
	err := c.cc.Invoke(ctx, "/driver.Driver/Fire", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverServer is the server API for Driver service.
type DriverServer interface {
	Left(context.Context, *CommandRequest) (*CommandReply, error)
	Up(context.Context, *CommandRequest) (*CommandReply, error)
	Right(context.Context, *CommandRequest) (*CommandReply, error)
	Down(context.Context, *CommandRequest) (*CommandReply, error)
	UpLeft(context.Context, *CommandRequest) (*CommandReply, error)
	DownLeft(context.Context, *CommandRequest) (*CommandReply, error)
	UpRight(context.Context, *CommandRequest) (*CommandReply, error)
	DownRight(context.Context, *CommandRequest) (*CommandReply, error)
	Stop(context.Context, *CommandRequest) (*CommandReply, error)
	Fire(context.Context, *CommandRequest) (*CommandReply, error)
}

// UnimplementedDriverServer can be embedded to have forward compatible implementations.
type UnimplementedDriverServer struct {
}

func (*UnimplementedDriverServer) Left(ctx context.Context, req *CommandRequest) (*CommandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Left not implemented")
}
func (*UnimplementedDriverServer) Up(ctx context.Context, req *CommandRequest) (*CommandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Up not implemented")
}
func (*UnimplementedDriverServer) Right(ctx context.Context, req *CommandRequest) (*CommandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Right not implemented")
}
func (*UnimplementedDriverServer) Down(ctx context.Context, req *CommandRequest) (*CommandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Down not implemented")
}
func (*UnimplementedDriverServer) UpLeft(ctx context.Context, req *CommandRequest) (*CommandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpLeft not implemented")
}
func (*UnimplementedDriverServer) DownLeft(ctx context.Context, req *CommandRequest) (*CommandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownLeft not implemented")
}
func (*UnimplementedDriverServer) UpRight(ctx context.Context, req *CommandRequest) (*CommandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpRight not implemented")
}
func (*UnimplementedDriverServer) DownRight(ctx context.Context, req *CommandRequest) (*CommandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownRight not implemented")
}
func (*UnimplementedDriverServer) Stop(ctx context.Context, req *CommandRequest) (*CommandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (*UnimplementedDriverServer) Fire(ctx context.Context, req *CommandRequest) (*CommandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fire not implemented")
}

func RegisterDriverServer(s *grpc.Server, srv DriverServer) {
	s.RegisterService(&_Driver_serviceDesc, srv)
}

func _Driver_Left_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Left(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driver.Driver/Left",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Left(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Up_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Up(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driver.Driver/Up",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Up(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Right_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Right(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driver.Driver/Right",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Right(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Down_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Down(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driver.Driver/Down",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Down(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_UpLeft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).UpLeft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driver.Driver/UpLeft",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).UpLeft(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_DownLeft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).DownLeft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driver.Driver/DownLeft",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).DownLeft(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_UpRight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).UpRight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driver.Driver/UpRight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).UpRight(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_DownRight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).DownRight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driver.Driver/DownRight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).DownRight(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driver.Driver/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Stop(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Fire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Fire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driver.Driver/Fire",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Fire(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Driver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "driver.Driver",
	HandlerType: (*DriverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Left",
			Handler:    _Driver_Left_Handler,
		},
		{
			MethodName: "Up",
			Handler:    _Driver_Up_Handler,
		},
		{
			MethodName: "Right",
			Handler:    _Driver_Right_Handler,
		},
		{
			MethodName: "Down",
			Handler:    _Driver_Down_Handler,
		},
		{
			MethodName: "UpLeft",
			Handler:    _Driver_UpLeft_Handler,
		},
		{
			MethodName: "DownLeft",
			Handler:    _Driver_DownLeft_Handler,
		},
		{
			MethodName: "UpRight",
			Handler:    _Driver_UpRight_Handler,
		},
		{
			MethodName: "DownRight",
			Handler:    _Driver_DownRight_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Driver_Stop_Handler,
		},
		{
			MethodName: "Fire",
			Handler:    _Driver_Fire_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "driver/driver.proto",
}
