// Code generated by protoc-gen-go. DO NOT EDIT.
// source: terminal.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type TerminalUI struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminalUI) Reset()         { *m = TerminalUI{} }
func (m *TerminalUI) String() string { return proto.CompactTextString(m) }
func (*TerminalUI) ProtoMessage()    {}
func (*TerminalUI) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff8b8260c8ef16ad, []int{0}
}

func (m *TerminalUI) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminalUI.Unmarshal(m, b)
}
func (m *TerminalUI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminalUI.Marshal(b, m, deterministic)
}
func (m *TerminalUI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminalUI.Merge(m, src)
}
func (m *TerminalUI) XXX_Size() int {
	return xxx_messageInfo_TerminalUI.Size(m)
}
func (m *TerminalUI) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminalUI.DiscardUnknown(m)
}

var xxx_messageInfo_TerminalUI proto.InternalMessageInfo

type TerminalUI_OutputRequest struct {
	Lines                []string `protobuf:"bytes,1,rep,name=lines,proto3" json:"lines,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminalUI_OutputRequest) Reset()         { *m = TerminalUI_OutputRequest{} }
func (m *TerminalUI_OutputRequest) String() string { return proto.CompactTextString(m) }
func (*TerminalUI_OutputRequest) ProtoMessage()    {}
func (*TerminalUI_OutputRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff8b8260c8ef16ad, []int{0, 0}
}

func (m *TerminalUI_OutputRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminalUI_OutputRequest.Unmarshal(m, b)
}
func (m *TerminalUI_OutputRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminalUI_OutputRequest.Marshal(b, m, deterministic)
}
func (m *TerminalUI_OutputRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminalUI_OutputRequest.Merge(m, src)
}
func (m *TerminalUI_OutputRequest) XXX_Size() int {
	return xxx_messageInfo_TerminalUI_OutputRequest.Size(m)
}
func (m *TerminalUI_OutputRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminalUI_OutputRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TerminalUI_OutputRequest proto.InternalMessageInfo

func (m *TerminalUI_OutputRequest) GetLines() []string {
	if m != nil {
		return m.Lines
	}
	return nil
}

func init() {
	proto.RegisterType((*TerminalUI)(nil), "proto.TerminalUI")
	proto.RegisterType((*TerminalUI_OutputRequest)(nil), "proto.TerminalUI.OutputRequest")
}

func init() {
	proto.RegisterFile("terminal.proto", fileDescriptor_ff8b8260c8ef16ad)
}

var fileDescriptor_ff8b8260c8ef16ad = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x49, 0x2d, 0xca,
	0xcd, 0xcc, 0x4b, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0xd2,
	0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x60, 0x5e, 0x52, 0x69, 0x9a, 0x7e, 0x6a, 0x6e, 0x41,
	0x49, 0x25, 0x44, 0x8d, 0x92, 0x31, 0x17, 0x57, 0x08, 0x54, 0x57, 0xa8, 0xa7, 0x94, 0x2a, 0x17,
	0xaf, 0x7f, 0x69, 0x49, 0x41, 0x69, 0x49, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x08,
	0x17, 0x6b, 0x4e, 0x66, 0x5e, 0x6a, 0xb1, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0x67, 0x10, 0x84, 0x63,
	0x14, 0xc6, 0x25, 0x88, 0xd0, 0x14, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xe4, 0xc8, 0xc5,
	0x06, 0xd1, 0x2b, 0x24, 0x0f, 0x31, 0x5b, 0x0f, 0xa1, 0x46, 0x0f, 0xc5, 0x54, 0x29, 0x31, 0x3d,
	0x88, 0x93, 0xf4, 0x60, 0x4e, 0xd2, 0x73, 0x05, 0x39, 0x29, 0x89, 0x0d, 0xcc, 0x37, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x71, 0x51, 0x6f, 0xa7, 0xc9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TerminalUIServiceClient is the client API for TerminalUIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TerminalUIServiceClient interface {
	Output(ctx context.Context, in *TerminalUI_OutputRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type terminalUIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTerminalUIServiceClient(cc grpc.ClientConnInterface) TerminalUIServiceClient {
	return &terminalUIServiceClient{cc}
}

func (c *terminalUIServiceClient) Output(ctx context.Context, in *TerminalUI_OutputRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.TerminalUIService/Output", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TerminalUIServiceServer is the server API for TerminalUIService service.
type TerminalUIServiceServer interface {
	Output(context.Context, *TerminalUI_OutputRequest) (*empty.Empty, error)
}

// UnimplementedTerminalUIServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTerminalUIServiceServer struct {
}

func (*UnimplementedTerminalUIServiceServer) Output(ctx context.Context, req *TerminalUI_OutputRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Output not implemented")
}

func RegisterTerminalUIServiceServer(s *grpc.Server, srv TerminalUIServiceServer) {
	s.RegisterService(&_TerminalUIService_serviceDesc, srv)
}

func _TerminalUIService_Output_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminalUI_OutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerminalUIServiceServer).Output(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TerminalUIService/Output",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerminalUIServiceServer).Output(ctx, req.(*TerminalUI_OutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TerminalUIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TerminalUIService",
	HandlerType: (*TerminalUIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Output",
			Handler:    _TerminalUIService_Output_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "terminal.proto",
}