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

type NeighbourInfo struct {
	Predecessor          string   `protobuf:"bytes,1,opt,name=predecessor,proto3" json:"predecessor,omitempty"`
	Successor            string   `protobuf:"bytes,2,opt,name=successor,proto3" json:"successor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NeighbourInfo) Reset()         { *m = NeighbourInfo{} }
func (m *NeighbourInfo) String() string { return proto.CompactTextString(m) }
func (*NeighbourInfo) ProtoMessage()    {}
func (*NeighbourInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9621718efbbc1f14, []int{1}
}

func (m *NeighbourInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NeighbourInfo.Unmarshal(m, b)
}
func (m *NeighbourInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NeighbourInfo.Marshal(b, m, deterministic)
}
func (m *NeighbourInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NeighbourInfo.Merge(m, src)
}
func (m *NeighbourInfo) XXX_Size() int {
	return xxx_messageInfo_NeighbourInfo.Size(m)
}
func (m *NeighbourInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NeighbourInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NeighbourInfo proto.InternalMessageInfo

func (m *NeighbourInfo) GetPredecessor() string {
	if m != nil {
		return m.Predecessor
	}
	return ""
}

func (m *NeighbourInfo) GetSuccessor() string {
	if m != nil {
		return m.Successor
	}
	return ""
}

type OkReponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OkReponse) Reset()         { *m = OkReponse{} }
func (m *OkReponse) String() string { return proto.CompactTextString(m) }
func (*OkReponse) ProtoMessage()    {}
func (*OkReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9621718efbbc1f14, []int{2}
}

func (m *OkReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OkReponse.Unmarshal(m, b)
}
func (m *OkReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OkReponse.Marshal(b, m, deterministic)
}
func (m *OkReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OkReponse.Merge(m, src)
}
func (m *OkReponse) XXX_Size() int {
	return xxx_messageInfo_OkReponse.Size(m)
}
func (m *OkReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OkReponse.DiscardUnknown(m)
}

var xxx_messageInfo_OkReponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "RegisterRequest")
	proto.RegisterType((*NeighbourInfo)(nil), "NeighbourInfo")
	proto.RegisterType((*OkReponse)(nil), "OkReponse")
}

func init() { proto.RegisterFile("protos/chain.proto", fileDescriptor_9621718efbbc1f14) }

var fileDescriptor_9621718efbbc1f14 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xa9, 0xe0, 0x9f, 0x4c, 0x51, 0xcb, 0x9e, 0x4a, 0xf1, 0x50, 0x72, 0x12, 0xd4, 0x5d,
	0xd4, 0x6f, 0xa0, 0x22, 0xd8, 0x83, 0x85, 0x80, 0x17, 0x6f, 0x49, 0xf6, 0x99, 0x2c, 0x62, 0x36,
	0xce, 0xcc, 0x7e, 0x7f, 0x31, 0xb6, 0x55, 0x73, 0x5a, 0xde, 0xef, 0x3d, 0xd8, 0x1f, 0x43, 0xa6,
	0xe7, 0xa8, 0x51, 0x5c, 0xdd, 0x96, 0xa1, 0xb3, 0x43, 0xc8, 0x2f, 0xe8, 0xb4, 0x40, 0x13, 0x44,
	0xc1, 0x05, 0x3e, 0x13, 0x44, 0xcd, 0x9c, 0x0e, 0x4b, 0xef, 0x19, 0x22, 0xf3, 0xc9, 0x72, 0x72,
	0x9e, 0x15, 0xdb, 0x98, 0xaf, 0xe9, 0xf8, 0x19, 0xa1, 0x69, 0xab, 0x98, 0xf8, 0xa9, 0x7b, 0x8b,
	0x66, 0x49, 0xd3, 0x9e, 0xe1, 0x51, 0x43, 0x24, 0xf2, 0x66, 0xfe, 0x17, 0x99, 0x33, 0xca, 0x24,
	0xd5, 0x9b, 0x7e, 0x6f, 0xe8, 0x7f, 0x41, 0x3e, 0xa5, 0x6c, 0xfd, 0x5e, 0xa0, 0x8f, 0x9d, 0xe0,
	0x06, 0xb4, 0x7f, 0xff, 0x6d, 0x66, 0x2e, 0xe9, 0x68, 0xeb, 0x64, 0x66, 0x76, 0xa4, 0xb7, 0x38,
	0xb1, 0xff, 0x1d, 0x2c, 0xcd, 0x5e, 0x7a, 0x5f, 0x2a, 0x76, 0x58, 0xcc, 0x68, 0xb3, 0x20, 0xbb,
	0xfb, 0xe6, 0xee, 0xfa, 0xd5, 0x35, 0x41, 0xdb, 0x54, 0xd9, 0x3a, 0x7e, 0xb8, 0xd5, 0xc3, 0xea,
	0x31, 0x48, 0x0b, 0x76, 0x3e, 0x88, 0x72, 0xa8, 0x92, 0xc2, 0x5f, 0x89, 0x46, 0x2e, 0x1b, 0xb8,
	0x9f, 0x83, 0x55, 0x07, 0xc3, 0x7b, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x72, 0x8b, 0x2e, 0x5c,
	0x41, 0x01, 0x00, 0x00,
}
