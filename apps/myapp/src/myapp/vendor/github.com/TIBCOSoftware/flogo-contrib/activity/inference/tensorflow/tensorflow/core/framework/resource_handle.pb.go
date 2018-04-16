// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/resource_handle.proto

/*
Package tensorflow is a generated protocol buffer package.

It is generated from these files:
	tensorflow/core/framework/resource_handle.proto

It has these top-level messages:
	ResourceHandleProto
*/
package tffw

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type ResourceHandleProto struct {
	// Unique name for the device containing the resource.
	Device string `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
	// Container in which this resource is placed.
	Container string `protobuf:"bytes,2,opt,name=container" json:"container,omitempty"`
	// Unique name of this resource.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// Hash code for the type of the resource. Is only valid in the same device
	// and in the same execution.
	HashCode uint64 `protobuf:"varint,4,opt,name=hash_code,json=hashCode" json:"hash_code,omitempty"`
	// For debug-only, the name of the type pointed to by this handle, if
	// available.
	MaybeTypeName string `protobuf:"bytes,5,opt,name=maybe_type_name,json=maybeTypeName" json:"maybe_type_name,omitempty"`
}

func (m *ResourceHandleProto) Reset()                    { *m = ResourceHandleProto{} }
func (m *ResourceHandleProto) String() string            { return proto.CompactTextString(m) }
func (*ResourceHandleProto) ProtoMessage()               {}
func (*ResourceHandleProto) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *ResourceHandleProto) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *ResourceHandleProto) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *ResourceHandleProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceHandleProto) GetHashCode() uint64 {
	if m != nil {
		return m.HashCode
	}
	return 0
}

func (m *ResourceHandleProto) GetMaybeTypeName() string {
	if m != nil {
		return m.MaybeTypeName
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceHandleProto)(nil), "tensorflow.ResourceHandleProto")
}

func init() { proto.RegisterFile("tensorflow/core/framework/resource_handle.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x59, 0x8d, 0xc5, 0x0c, 0xa8, 0xb0, 0x82, 0x2c, 0xe8, 0xa1, 0x78, 0x90, 0x9e, 0xb2,
	0x07, 0xff, 0x41, 0xbd, 0x78, 0x12, 0x09, 0xde, 0xc3, 0x76, 0xf3, 0x6a, 0x8a, 0xcd, 0x4e, 0x98,
	0x44, 0x4b, 0xfe, 0x8f, 0x3f, 0xd2, 0xa3, 0x38, 0x04, 0x03, 0xbd, 0xcd, 0xbc, 0xf7, 0x3e, 0x78,
	0x8f, 0xfc, 0x80, 0xd4, 0xb3, 0x6c, 0xf7, 0x7c, 0xf0, 0x91, 0x05, 0x7e, 0x2b, 0xa1, 0xc5, 0x81,
	0xe5, 0xc3, 0x0b, 0x7a, 0xfe, 0x94, 0x88, 0xaa, 0x09, 0xa9, 0xde, 0xa3, 0xe8, 0x84, 0x07, 0xb6,
	0x34, 0x03, 0xf7, 0xdf, 0x86, 0xae, 0xcb, 0x29, 0xf5, 0xac, 0xa1, 0x57, 0xcd, 0xdc, 0xd0, 0xa2,
	0xc6, 0xd7, 0x2e, 0xc2, 0x99, 0xa5, 0x59, 0xe5, 0xe5, 0xf4, 0xd9, 0x3b, 0xca, 0x23, 0xa7, 0x21,
	0xec, 0x12, 0xc4, 0x9d, 0xa8, 0x35, 0x0b, 0xd6, 0x52, 0x96, 0x42, 0x0b, 0x77, 0xaa, 0x86, 0xde,
	0xf6, 0x96, 0xf2, 0x26, 0xf4, 0x4d, 0x15, 0xb9, 0x86, 0xcb, 0x96, 0x66, 0x95, 0x95, 0xe7, 0x7f,
	0xc2, 0x13, 0xd7, 0xb0, 0x0f, 0x74, 0xd5, 0x86, 0x71, 0x83, 0x6a, 0x18, 0x3b, 0x54, 0xca, 0x9e,
	0x29, 0x7b, 0xa1, 0xf2, 0xdb, 0xd8, 0xe1, 0x25, 0xb4, 0x58, 0x7b, 0x72, 0x2c, 0xef, 0xc5, 0x5c,
	0xbc, 0xf8, 0x1f, 0xb9, 0xbe, 0x3c, 0xea, 0x6f, 0x7e, 0x8c, 0xd9, 0x2c, 0x74, 0xea, 0xe3, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0x2a, 0x69, 0xfc, 0x1d, 0x01, 0x00, 0x00,
}
