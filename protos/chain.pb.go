// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/chain.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Provide the name of the service which is joining (docker service name), allows the network to then send messages back by using the docker DNS
type RegisterRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9621718efbbc1f14, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type RegisterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9621718efbbc1f14, []int{1}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "RegisterResponse")
}

func init() { proto.RegisterFile("protos/chain.proto", fileDescriptor_9621718efbbc1f14) }

var fileDescriptor_9621718efbbc1f14 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x03, 0x73, 0x94, 0xb4, 0xb9, 0xf8, 0x83,
	0x52, 0xd3, 0x33, 0x8b, 0x4b, 0x52, 0x8b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24,
	0xb8, 0xd8, 0x13, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x60, 0x5c, 0x25, 0x21, 0x2e, 0x01, 0x84, 0xe2, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x23, 0x0b,
	0x2e, 0x56, 0x67, 0x90, 0x79, 0x42, 0xfa, 0x5c, 0x1c, 0x30, 0x49, 0x21, 0x01, 0x3d, 0x34, 0x43,
	0xa5, 0x04, 0xf5, 0xd0, 0x75, 0x3a, 0x19, 0x46, 0xe9, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9,
	0x25, 0xe7, 0xe7, 0xea, 0x7b, 0xb9, 0x78, 0xb9, 0x65, 0x16, 0x67, 0xa4, 0x16, 0xe9, 0xa7, 0x64,
	0x16, 0x97, 0x14, 0x65, 0x26, 0x95, 0x96, 0xa4, 0xa6, 0xe8, 0x16, 0x97, 0xe4, 0x17, 0x25, 0xa6,
	0xa7, 0xea, 0x43, 0x5c, 0x9e, 0xc4, 0x06, 0xa6, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x80,
	0x7e, 0xa0, 0xf1, 0xca, 0x00, 0x00, 0x00,
}
