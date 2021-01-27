// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/network.proto

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

type NetworkJoinResponse_NodeType int32

const (
	NetworkJoinResponse_HEAD   NetworkJoinResponse_NodeType = 0
	NetworkJoinResponse_TAIL   NetworkJoinResponse_NodeType = 1
	NetworkJoinResponse_NORMAL NetworkJoinResponse_NodeType = 2
)

var NetworkJoinResponse_NodeType_name = map[int32]string{
	0: "HEAD",
	1: "TAIL",
	2: "NORMAL",
}

var NetworkJoinResponse_NodeType_value = map[string]int32{
	"HEAD":   0,
	"TAIL":   1,
	"NORMAL": 2,
}

func (x NetworkJoinResponse_NodeType) String() string {
	return proto.EnumName(NetworkJoinResponse_NodeType_name, int32(x))
}

func (NetworkJoinResponse_NodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9898f5d59e04eeea, []int{1, 0}
}

type NetworkJoinRequest struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkJoinRequest) Reset()         { *m = NetworkJoinRequest{} }
func (m *NetworkJoinRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkJoinRequest) ProtoMessage()    {}
func (*NetworkJoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9898f5d59e04eeea, []int{0}
}

func (m *NetworkJoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkJoinRequest.Unmarshal(m, b)
}
func (m *NetworkJoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkJoinRequest.Marshal(b, m, deterministic)
}
func (m *NetworkJoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkJoinRequest.Merge(m, src)
}
func (m *NetworkJoinRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkJoinRequest.Size(m)
}
func (m *NetworkJoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkJoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkJoinRequest proto.InternalMessageInfo

func (m *NetworkJoinRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

type NetworkJoinResponse struct {
	Type                 NetworkJoinResponse_NodeType `protobuf:"varint,2,opt,name=type,proto3,enum=NetworkJoinResponse_NodeType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *NetworkJoinResponse) Reset()         { *m = NetworkJoinResponse{} }
func (m *NetworkJoinResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkJoinResponse) ProtoMessage()    {}
func (*NetworkJoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9898f5d59e04eeea, []int{1}
}

func (m *NetworkJoinResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkJoinResponse.Unmarshal(m, b)
}
func (m *NetworkJoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkJoinResponse.Marshal(b, m, deterministic)
}
func (m *NetworkJoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkJoinResponse.Merge(m, src)
}
func (m *NetworkJoinResponse) XXX_Size() int {
	return xxx_messageInfo_NetworkJoinResponse.Size(m)
}
func (m *NetworkJoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkJoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkJoinResponse proto.InternalMessageInfo

func (m *NetworkJoinResponse) GetType() NetworkJoinResponse_NodeType {
	if m != nil {
		return m.Type
	}
	return NetworkJoinResponse_HEAD
}

func init() {
	proto.RegisterEnum("NetworkJoinResponse_NodeType", NetworkJoinResponse_NodeType_name, NetworkJoinResponse_NodeType_value)
	proto.RegisterType((*NetworkJoinRequest)(nil), "NetworkJoinRequest")
	proto.RegisterType((*NetworkJoinResponse)(nil), "NetworkJoinResponse")
}

func init() { proto.RegisterFile("protos/network.proto", fileDescriptor_9898f5d59e04eeea) }

var fileDescriptor_9898f5d59e04eeea = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x03, 0x73, 0x95, 0xcc, 0xb8,
	0x84, 0xfc, 0x20, 0x02, 0x5e, 0xf9, 0x99, 0x79, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42,
	0x0a, 0x5c, 0xdc, 0xc5, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x7e, 0x89, 0xb9, 0xa9, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xc8, 0x42, 0x4a, 0x25, 0x5c, 0xc2, 0x28, 0xfa, 0x8a, 0x0b, 0xf2,
	0xf3, 0x8a, 0x53, 0x85, 0x0c, 0xb9, 0x58, 0x4a, 0x2a, 0x0b, 0x52, 0x25, 0x98, 0x14, 0x18, 0x35,
	0xf8, 0x8c, 0x64, 0xf5, 0xb0, 0xa8, 0xd1, 0xf3, 0xcb, 0x4f, 0x49, 0x0d, 0xa9, 0x2c, 0x48, 0x0d,
	0x02, 0x2b, 0x55, 0xd2, 0xe2, 0xe2, 0x80, 0x89, 0x08, 0x71, 0x70, 0xb1, 0x78, 0xb8, 0x3a, 0xba,
	0x08, 0x30, 0x80, 0x58, 0x21, 0x8e, 0x9e, 0x3e, 0x02, 0x8c, 0x42, 0x5c, 0x5c, 0x6c, 0x7e, 0xfe,
	0x41, 0xbe, 0x8e, 0x3e, 0x02, 0x4c, 0x46, 0xce, 0x5c, 0xec, 0x50, 0x13, 0x85, 0x2c, 0xb8, 0xb8,
	0x41, 0xa6, 0xc2, 0xb8, 0xc2, 0x7a, 0x98, 0xde, 0x90, 0x12, 0xc1, 0x66, 0xbf, 0x93, 0x61, 0x94,
	0x7e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe, 0x97, 0x8b, 0x97, 0x5b,
	0x66, 0x71, 0x46, 0x6a, 0x91, 0x7e, 0x4a, 0x66, 0x71, 0x49, 0x51, 0x66, 0x52, 0x69, 0x49, 0x6a,
	0x8a, 0x6e, 0x71, 0x49, 0x7e, 0x51, 0x62, 0x7a, 0xaa, 0x3e, 0x24, 0xcc, 0x92, 0xd8, 0xc0, 0xb4,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x98, 0x03, 0xd0, 0x34, 0x44, 0x01, 0x00, 0x00,
}
