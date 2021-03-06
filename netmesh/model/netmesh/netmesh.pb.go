// Code generated by protoc-gen-go. DO NOT EDIT.
// source: netmesh.proto

package netmesh

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

type Label struct {
	Selector             map[string]string `protobuf:"bytes,1,rep,name=selector,proto3" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Label) Reset()         { *m = Label{} }
func (m *Label) String() string { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()    {}
func (*Label) Descriptor() ([]byte, []int) {
	return fileDescriptor_netmesh_8c0dd4f2c491bf60, []int{0}
}
func (m *Label) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Label.Unmarshal(m, b)
}
func (m *Label) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Label.Marshal(b, m, deterministic)
}
func (dst *Label) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Label.Merge(dst, src)
}
func (m *Label) XXX_Size() int {
	return xxx_messageInfo_Label.Size(m)
}
func (m *Label) XXX_DiscardUnknown() {
	xxx_messageInfo_Label.DiscardUnknown(m)
}

var xxx_messageInfo_Label proto.InternalMessageInfo

func (m *Label) GetSelector() map[string]string {
	if m != nil {
		return m.Selector
	}
	return nil
}

type Metadata struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Labels               *Label   `protobuf:"bytes,3,opt,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_netmesh_8c0dd4f2c491bf60, []int{1}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (dst *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(dst, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metadata) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Metadata) GetLabels() *Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

type NetworkServiceEndpoint struct {
	Metadata             *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *NetworkServiceEndpoint) Reset()         { *m = NetworkServiceEndpoint{} }
func (m *NetworkServiceEndpoint) String() string { return proto.CompactTextString(m) }
func (*NetworkServiceEndpoint) ProtoMessage()    {}
func (*NetworkServiceEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_netmesh_8c0dd4f2c491bf60, []int{2}
}
func (m *NetworkServiceEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServiceEndpoint.Unmarshal(m, b)
}
func (m *NetworkServiceEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServiceEndpoint.Marshal(b, m, deterministic)
}
func (dst *NetworkServiceEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServiceEndpoint.Merge(dst, src)
}
func (m *NetworkServiceEndpoint) XXX_Size() int {
	return xxx_messageInfo_NetworkServiceEndpoint.Size(m)
}
func (m *NetworkServiceEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServiceEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServiceEndpoint proto.InternalMessageInfo

func (m *NetworkServiceEndpoint) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type NetworkService struct {
	Metadata             *Metadata                        `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Channels             []*NetworkService_NetmeshChannel `protobuf:"bytes,2,rep,name=channels,proto3" json:"channels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *NetworkService) Reset()         { *m = NetworkService{} }
func (m *NetworkService) String() string { return proto.CompactTextString(m) }
func (*NetworkService) ProtoMessage()    {}
func (*NetworkService) Descriptor() ([]byte, []int) {
	return fileDescriptor_netmesh_8c0dd4f2c491bf60, []int{3}
}
func (m *NetworkService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkService.Unmarshal(m, b)
}
func (m *NetworkService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkService.Marshal(b, m, deterministic)
}
func (dst *NetworkService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkService.Merge(dst, src)
}
func (m *NetworkService) XXX_Size() int {
	return xxx_messageInfo_NetworkService.Size(m)
}
func (m *NetworkService) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkService.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkService proto.InternalMessageInfo

func (m *NetworkService) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *NetworkService) GetChannels() []*NetworkService_NetmeshChannel {
	if m != nil {
		return m.Channels
	}
	return nil
}

type NetworkService_NetmeshChannel struct {
	Metadata             *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Payload              string    `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *NetworkService_NetmeshChannel) Reset()         { *m = NetworkService_NetmeshChannel{} }
func (m *NetworkService_NetmeshChannel) String() string { return proto.CompactTextString(m) }
func (*NetworkService_NetmeshChannel) ProtoMessage()    {}
func (*NetworkService_NetmeshChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_netmesh_8c0dd4f2c491bf60, []int{3, 0}
}
func (m *NetworkService_NetmeshChannel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkService_NetmeshChannel.Unmarshal(m, b)
}
func (m *NetworkService_NetmeshChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkService_NetmeshChannel.Marshal(b, m, deterministic)
}
func (dst *NetworkService_NetmeshChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkService_NetmeshChannel.Merge(dst, src)
}
func (m *NetworkService_NetmeshChannel) XXX_Size() int {
	return xxx_messageInfo_NetworkService_NetmeshChannel.Size(m)
}
func (m *NetworkService_NetmeshChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkService_NetmeshChannel.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkService_NetmeshChannel proto.InternalMessageInfo

func (m *NetworkService_NetmeshChannel) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *NetworkService_NetmeshChannel) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func init() {
	proto.RegisterType((*Label)(nil), "netmesh.Label")
	proto.RegisterMapType((map[string]string)(nil), "netmesh.Label.SelectorEntry")
	proto.RegisterType((*Metadata)(nil), "netmesh.Metadata")
	proto.RegisterType((*NetworkServiceEndpoint)(nil), "netmesh.NetworkServiceEndpoint")
	proto.RegisterType((*NetworkService)(nil), "netmesh.NetworkService")
	proto.RegisterType((*NetworkService_NetmeshChannel)(nil), "netmesh.NetworkService.NetmeshChannel")
}

func init() { proto.RegisterFile("netmesh.proto", fileDescriptor_netmesh_8c0dd4f2c491bf60) }

var fileDescriptor_netmesh_8c0dd4f2c491bf60 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x3f, 0x6f, 0xbb, 0x30,
	0x10, 0x95, 0xe1, 0x97, 0x84, 0x1c, 0x4a, 0xf4, 0xab, 0x55, 0x55, 0x28, 0xca, 0x80, 0x18, 0x22,
	0x96, 0x32, 0xd0, 0x25, 0x6a, 0xb7, 0x56, 0x51, 0x97, 0xb6, 0x03, 0x99, 0x3a, 0x3a, 0x70, 0x52,
	0xa2, 0x80, 0x8d, 0xc0, 0x4d, 0xc5, 0xd2, 0x8f, 0xda, 0xcf, 0x52, 0xd9, 0x31, 0x6e, 0x19, 0x33,
	0xe1, 0x77, 0xef, 0xcf, 0x3d, 0x9d, 0x80, 0x19, 0x47, 0x59, 0x61, 0xbb, 0x4f, 0xea, 0x46, 0x48,
	0x41, 0x27, 0x06, 0x46, 0x5f, 0x30, 0x7a, 0x61, 0x3b, 0x2c, 0xe9, 0x1a, 0xbc, 0x16, 0x4b, 0xcc,
	0xa5, 0x68, 0x02, 0x12, 0xba, 0xb1, 0x9f, 0x2e, 0x93, 0xde, 0xa3, 0x15, 0xc9, 0xd6, 0xd0, 0x1b,
	0x2e, 0x9b, 0x2e, 0xb3, 0xea, 0xc5, 0x03, 0xcc, 0x06, 0x14, 0xfd, 0x0f, 0xee, 0x11, 0xbb, 0x80,
	0x84, 0x24, 0x9e, 0x66, 0xea, 0x49, 0xaf, 0x61, 0x74, 0x62, 0xe5, 0x07, 0x06, 0x8e, 0x9e, 0x9d,
	0xc1, 0xbd, 0xb3, 0x26, 0x51, 0x01, 0xde, 0x2b, 0x4a, 0x56, 0x30, 0xc9, 0x28, 0x85, 0x7f, 0x9c,
	0x55, 0x68, 0x8c, 0xfa, 0x4d, 0x97, 0x30, 0x55, 0xdf, 0xb6, 0x66, 0x79, 0xef, 0xfe, 0x1d, 0xd0,
	0x15, 0x8c, 0x4b, 0xd5, 0xad, 0x0d, 0xdc, 0x90, 0xc4, 0x7e, 0x3a, 0x1f, 0x56, 0xce, 0x0c, 0x1b,
	0x3d, 0xc3, 0xcd, 0x1b, 0xca, 0x4f, 0xd1, 0x1c, 0xb7, 0xd8, 0x9c, 0x0e, 0x39, 0x6e, 0x78, 0x51,
	0x8b, 0x03, 0x97, 0xf4, 0x16, 0xbc, 0xca, 0xec, 0xd7, 0x7b, 0xfd, 0xf4, 0xca, 0x66, 0xf4, 0xc5,
	0x32, 0x2b, 0x89, 0xbe, 0x09, 0xcc, 0x87, 0x49, 0x17, 0x26, 0xd0, 0x47, 0xf0, 0xf2, 0x3d, 0xe3,
	0x5c, 0x95, 0x76, 0xf4, 0x9d, 0x57, 0x56, 0x3e, 0x4c, 0x56, 0x50, 0x8d, 0x9f, 0xce, 0xf2, 0xcc,
	0xfa, 0x16, 0xef, 0xba, 0xc4, 0x1f, 0xee, 0xd2, 0x12, 0x01, 0x4c, 0x6a, 0xd6, 0x95, 0x82, 0x15,
	0xe6, 0xa6, 0x3d, 0xdc, 0x8d, 0xf5, 0xff, 0x71, 0xf7, 0x13, 0x00, 0x00, 0xff, 0xff, 0x52, 0x08,
	0xa7, 0x9f, 0x30, 0x02, 0x00, 0x00,
}
