// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/aws_request_signing/v3/aws_request_signing.proto

package envoy_extensions_filters_http_aws_request_signing_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type AwsRequestSigning struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Region               string   `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsRequestSigning) Reset()         { *m = AwsRequestSigning{} }
func (m *AwsRequestSigning) String() string { return proto.CompactTextString(m) }
func (*AwsRequestSigning) ProtoMessage()    {}
func (*AwsRequestSigning) Descriptor() ([]byte, []int) {
	return fileDescriptor_c87dd6f1b9f622ab, []int{0}
}

func (m *AwsRequestSigning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsRequestSigning.Unmarshal(m, b)
}
func (m *AwsRequestSigning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsRequestSigning.Marshal(b, m, deterministic)
}
func (m *AwsRequestSigning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsRequestSigning.Merge(m, src)
}
func (m *AwsRequestSigning) XXX_Size() int {
	return xxx_messageInfo_AwsRequestSigning.Size(m)
}
func (m *AwsRequestSigning) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsRequestSigning.DiscardUnknown(m)
}

var xxx_messageInfo_AwsRequestSigning proto.InternalMessageInfo

func (m *AwsRequestSigning) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *AwsRequestSigning) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func init() {
	proto.RegisterType((*AwsRequestSigning)(nil), "envoy.extensions.filters.http.aws_request_signing.v3.AwsRequestSigning")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/aws_request_signing/v3/aws_request_signing.proto", fileDescriptor_c87dd6f1b9f622ab)
}

var fileDescriptor_c87dd6f1b9f622ab = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd0, 0xbd, 0x4e, 0x03, 0x31,
	0x0c, 0x07, 0x70, 0xa5, 0x42, 0x45, 0x04, 0x16, 0x6e, 0x80, 0xaa, 0x03, 0x14, 0x26, 0xc4, 0x90,
	0x48, 0x1c, 0x13, 0x1b, 0x19, 0x10, 0x0b, 0x55, 0x55, 0x76, 0x4e, 0xa6, 0x75, 0xaf, 0x91, 0xae,
	0x4e, 0x48, 0xd2, 0xb4, 0x7d, 0x03, 0x9e, 0x81, 0x97, 0xe0, 0x29, 0x78, 0x29, 0x26, 0x74, 0x1f,
	0x80, 0x50, 0x2b, 0x86, 0x6e, 0x91, 0x1d, 0xff, 0x64, 0xff, 0x79, 0x1f, 0x29, 0x9a, 0x95, 0xc4,
	0x65, 0x40, 0xf2, 0xda, 0x90, 0x97, 0x13, 0x5d, 0x04, 0x74, 0x5e, 0x4e, 0x43, 0xb0, 0x12, 0x16,
	0x3e, 0x73, 0xf8, 0x32, 0x47, 0x1f, 0x32, 0xaf, 0x73, 0xd2, 0x94, 0xcb, 0x98, 0x6e, 0x2a, 0x0b,
	0xeb, 0x4c, 0x30, 0xc9, 0x75, 0xe5, 0x89, 0x5f, 0x4f, 0x34, 0x9e, 0x28, 0x3d, 0xb1, 0x69, 0x30,
	0xa6, 0xdd, 0xb3, 0xf9, 0xd8, 0x82, 0x04, 0x22, 0x13, 0x20, 0x54, 0x5b, 0x44, 0x74, 0xe5, 0xf8,
	0x0f, 0xdc, 0x3d, 0x8e, 0x50, 0xe8, 0x31, 0x04, 0x94, 0xdf, 0x8f, 0xba, 0x71, 0xfe, 0xce, 0xf8,
	0xe1, 0xed, 0xc2, 0x0f, 0x6b, 0xf5, 0xb1, 0x46, 0x93, 0x4b, 0x7e, 0xe0, 0xd1, 0x45, 0x3d, 0xc2,
	0x8c, 0x60, 0x86, 0x1d, 0xd6, 0x63, 0x17, 0x7b, 0x6a, 0xf7, 0x53, 0xed, 0xb8, 0x56, 0x8f, 0x0d,
	0xf7, 0x9b, 0x66, 0x1f, 0x66, 0x98, 0x9c, 0xf2, 0xb6, 0xc3, 0x5c, 0x1b, 0xea, 0xb4, 0xfe, 0xfe,
	0x6a, 0xca, 0x37, 0x0f, 0x6f, 0x1f, 0xaf, 0x27, 0xf7, 0xfc, 0xae, 0xbe, 0x6d, 0x64, 0x68, 0xa2,
	0xf3, 0xe6, 0xae, 0x7f, 0xce, 0xba, 0x82, 0xc2, 0x4e, 0x41, 0xac, 0xed, 0xa6, 0x9e, 0xb8, 0xd2,
	0x46, 0x54, 0x98, 0x75, 0x66, 0xb9, 0x12, 0xdb, 0x64, 0xa6, 0x8e, 0xd6, 0xe0, 0x41, 0x99, 0xc7,
	0x80, 0x3d, 0xb7, 0xab, 0x60, 0xd2, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xab, 0xbe, 0x15,
	0xdc, 0x01, 0x00, 0x00,
}