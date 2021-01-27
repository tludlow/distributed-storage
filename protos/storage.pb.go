// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/storage.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type ReadRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d058d87b753bae70, []int{0}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type WriteRequest struct {
	Id                   int32           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Value                *_struct.Struct `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d058d87b753bae70, []int{1}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *WriteRequest) GetValue() *_struct.Struct {
	if m != nil {
		return m.Value
	}
	return nil
}

type ReadResponse struct {
	Value                *_struct.Struct `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d058d87b753bae70, []int{2}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetValue() *_struct.Struct {
	if m != nil {
		return m.Value
	}
	return nil
}

type WriteResponse struct {
	Value                *_struct.Struct `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d058d87b753bae70, []int{3}
}

func (m *WriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteResponse.Unmarshal(m, b)
}
func (m *WriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteResponse.Marshal(b, m, deterministic)
}
func (m *WriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponse.Merge(m, src)
}
func (m *WriteResponse) XXX_Size() int {
	return xxx_messageInfo_WriteResponse.Size(m)
}
func (m *WriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponse proto.InternalMessageInfo

func (m *WriteResponse) GetValue() *_struct.Struct {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*ReadRequest)(nil), "ReadRequest")
	proto.RegisterType((*WriteRequest)(nil), "WriteRequest")
	proto.RegisterType((*ReadResponse)(nil), "ReadResponse")
	proto.RegisterType((*WriteResponse)(nil), "WriteResponse")
}

func init() { proto.RegisterFile("protos/storage.proto", fileDescriptor_d058d87b753bae70) }

var fileDescriptor_d058d87b753bae70 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0x4b, 0xf3, 0x40,
	0x10, 0xc5, 0x49, 0xf8, 0xf2, 0x09, 0xd3, 0xa4, 0x87, 0x45, 0xb0, 0x04, 0x85, 0x12, 0x41, 0x7a,
	0xe9, 0x2c, 0xd6, 0xb3, 0x1e, 0x44, 0x3c, 0x14, 0xbc, 0xa4, 0xa0, 0xe0, 0x2d, 0xe9, 0x8e, 0xe9,
	0x42, 0x75, 0xeb, 0xce, 0xac, 0xfe, 0xfb, 0x62, 0xb6, 0x62, 0x7b, 0x10, 0xc4, 0xd3, 0x32, 0xcc,
	0x3e, 0x7e, 0x6f, 0xde, 0x83, 0xc3, 0x8d, 0x77, 0xe2, 0x58, 0xb3, 0x38, 0xdf, 0x74, 0x84, 0xfd,
	0x58, 0x1e, 0x77, 0xce, 0x75, 0x6b, 0xd2, 0xfd, 0xd4, 0x86, 0x27, 0xcd, 0xe2, 0xc3, 0x52, 0xe2,
	0xb6, 0x3a, 0x81, 0x41, 0x4d, 0x8d, 0xa9, 0xe9, 0x35, 0x10, 0x8b, 0x1a, 0x42, 0x6a, 0xcd, 0x28,
	0x19, 0x27, 0x93, 0xac, 0x4e, 0xad, 0xa9, 0xee, 0x20, 0x7f, 0xf0, 0x56, 0xe8, 0x87, 0xbd, 0x9a,
	0x42, 0xf6, 0xd6, 0xac, 0x03, 0x8d, 0xd2, 0x71, 0x32, 0x19, 0xcc, 0x8e, 0x30, 0xc2, 0xf0, 0x0b,
	0x86, 0x8b, 0x1e, 0x56, 0xc7, 0x5f, 0xd5, 0x25, 0xe4, 0x91, 0xc6, 0x1b, 0xf7, 0xc2, 0xf4, 0x2d,
	0x4f, 0x7e, 0x25, 0xbf, 0x82, 0x62, 0xeb, 0xe6, 0x4f, 0xfa, 0xd9, 0x3d, 0x1c, 0x2c, 0x62, 0x36,
	0xea, 0x14, 0xfe, 0x79, 0x6a, 0x8c, 0xca, 0x71, 0xe7, 0xfc, 0xb2, 0xc0, 0x3d, 0x7b, 0x67, 0x90,
	0xbd, 0x7f, 0xf2, 0x54, 0x81, 0xbb, 0x29, 0x94, 0x43, 0xdc, 0xb3, 0x71, 0x7d, 0xfe, 0xa8, 0x3b,
	0x2b, 0xab, 0xd0, 0xe2, 0xd2, 0x3d, 0xeb, 0xf9, 0xcd, 0xfc, 0xd6, 0xf2, 0x8a, 0xbc, 0x36, 0x96,
	0xc5, 0xdb, 0x36, 0x08, 0x99, 0xe9, 0xb6, 0x94, 0x58, 0x03, 0xb7, 0xff, 0xfb, 0xf7, 0xe2, 0x23,
	0x00, 0x00, 0xff, 0xff, 0x8c, 0x0f, 0x45, 0x08, 0xb4, 0x01, 0x00, 0x00,
}