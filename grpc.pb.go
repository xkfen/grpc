// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	grpc.proto

It has these top-level messages:
	ApprovalRpcReqMsg
	ApprovalReplyMsg
*/
package pb

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

type ApprovalRpcReqMsg struct {
	Method       string `protobuf:"bytes,1,opt,name=Method" json:"Method,omitempty"`
	RpcRequest   string `protobuf:"bytes,2,opt,name=RpcRequest" json:"RpcRequest,omitempty"`
	ApprovalType string `protobuf:"bytes,3,opt,name=ApprovalType" json:"ApprovalType,omitempty"`
	Username     string `protobuf:"bytes,4,opt,name=Username" json:"Username,omitempty"`
	Name         string `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	IsAuth       bool   `protobuf:"varint,6,opt,name=IsAuth" json:"IsAuth,omitempty"`
}

func (m *ApprovalRpcReqMsg) Reset()                    { *m = ApprovalRpcReqMsg{} }
func (m *ApprovalRpcReqMsg) String() string            { return proto.CompactTextString(m) }
func (*ApprovalRpcReqMsg) ProtoMessage()               {}
func (*ApprovalRpcReqMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ApprovalRpcReqMsg) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *ApprovalRpcReqMsg) GetRpcRequest() string {
	if m != nil {
		return m.RpcRequest
	}
	return ""
}

func (m *ApprovalRpcReqMsg) GetApprovalType() string {
	if m != nil {
		return m.ApprovalType
	}
	return ""
}

func (m *ApprovalRpcReqMsg) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ApprovalRpcReqMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApprovalRpcReqMsg) GetIsAuth() bool {
	if m != nil {
		return m.IsAuth
	}
	return false
}

type ApprovalReplyMsg struct {
	RpcReply string `protobuf:"bytes,1,opt,name=RpcReply" json:"RpcReply,omitempty"`
}

func (m *ApprovalReplyMsg) Reset()                    { *m = ApprovalReplyMsg{} }
func (m *ApprovalReplyMsg) String() string            { return proto.CompactTextString(m) }
func (*ApprovalReplyMsg) ProtoMessage()               {}
func (*ApprovalReplyMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ApprovalReplyMsg) GetRpcReply() string {
	if m != nil {
		return m.RpcReply
	}
	return ""
}

func init() {
	proto.RegisterType((*ApprovalRpcReqMsg)(nil), "pb.ApprovalRpcReqMsg")
	proto.RegisterType((*ApprovalReplyMsg)(nil), "pb.ApprovalReplyMsg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Approval service

type ApprovalClient interface {
	ApprovalRpcHandler(ctx context.Context, in *ApprovalRpcReqMsg, opts ...grpc.CallOption) (*ApprovalReplyMsg, error)
}

type approvalClient struct {
	cc *grpc.ClientConn
}

func NewApprovalClient(cc *grpc.ClientConn) ApprovalClient {
	return &approvalClient{cc}
}

func (c *approvalClient) ApprovalRpcHandler(ctx context.Context, in *ApprovalRpcReqMsg, opts ...grpc.CallOption) (*ApprovalReplyMsg, error) {
	out := new(ApprovalReplyMsg)
	err := grpc.Invoke(ctx, "/pb.Approval/ApprovalRpcHandler", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Approval service

type ApprovalServer interface {
	ApprovalRpcHandler(context.Context, *ApprovalRpcReqMsg) (*ApprovalReplyMsg, error)
}

func RegisterApprovalServer(s *grpc.Server, srv ApprovalServer) {
	s.RegisterService(&_Approval_serviceDesc, srv)
}

func _Approval_ApprovalRpcHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApprovalRpcReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalServer).ApprovalRpcHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Approval/ApprovalRpcHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalServer).ApprovalRpcHandler(ctx, req.(*ApprovalRpcReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Approval_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Approval",
	HandlerType: (*ApprovalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApprovalRpcHandler",
			Handler:    _Approval_ApprovalRpcHandler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xda, 0xce, 0xc8, 0x25, 0xe8,
	0x58, 0x50, 0x50, 0x94, 0x5f, 0x96, 0x98, 0x13, 0x54, 0x90, 0x1c, 0x94, 0x5a, 0xe8, 0x5b, 0x9c,
	0x2e, 0x24, 0xc6, 0xc5, 0xe6, 0x9b, 0x5a, 0x92, 0x91, 0x9f, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x19, 0x04, 0xe5, 0x09, 0xc9, 0x71, 0x71, 0x41, 0x14, 0x95, 0xa6, 0x16, 0x97, 0x48, 0x30, 0x81,
	0xe5, 0x90, 0x44, 0x84, 0x94, 0xb8, 0x78, 0x60, 0x86, 0x85, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x83,
	0x55, 0xa0, 0x88, 0x09, 0x49, 0x71, 0x71, 0x84, 0x16, 0xa7, 0x16, 0xe5, 0x25, 0xe6, 0xa6, 0x4a,
	0xb0, 0x80, 0xe5, 0xe1, 0x7c, 0x21, 0x21, 0x2e, 0x16, 0xb0, 0x38, 0x2b, 0x58, 0x1c, 0xcc, 0x06,
	0xb9, 0xc5, 0xb3, 0xd8, 0xb1, 0xb4, 0x24, 0x43, 0x82, 0x4d, 0x81, 0x51, 0x83, 0x23, 0x08, 0xca,
	0x53, 0xd2, 0xe3, 0x12, 0x80, 0x3b, 0x3c, 0xb5, 0x20, 0xa7, 0x12, 0xe4, 0x6e, 0x29, 0x2e, 0x0e,
	0xb0, 0x6b, 0x0a, 0x72, 0x2a, 0xa1, 0x2e, 0x87, 0xf3, 0x8d, 0xfc, 0xb9, 0x38, 0x60, 0xea, 0x85,
	0x9c, 0xb9, 0x84, 0x90, 0x3c, 0xed, 0x91, 0x98, 0x97, 0x92, 0x93, 0x5a, 0x24, 0x24, 0xaa, 0x57,
	0x90, 0xa4, 0x87, 0x11, 0x18, 0x52, 0x22, 0x28, 0xc2, 0x50, 0xab, 0x94, 0x18, 0x92, 0xd8, 0xc0,
	0xa1, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xec, 0xdf, 0x0e, 0x04, 0x53, 0x01, 0x00, 0x00,
}
