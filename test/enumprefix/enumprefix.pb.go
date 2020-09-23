// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enumprefix.proto

package enumprefix

import (
	fmt "fmt"
	_ "github.com/akqp2019/protobuf/gogoproto"
	proto "github.com/akqp2019/protobuf/proto"
	test "github.com/akqp2019/protobuf/test"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MyMessage struct {
	TheField             test.TheTestEnum `protobuf:"varint,1,opt,name=TheField,enum=test.TheTestEnum" json:"TheField"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MyMessage) Reset()         { *m = MyMessage{} }
func (m *MyMessage) String() string { return proto.CompactTextString(m) }
func (*MyMessage) ProtoMessage()    {}
func (*MyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0d23e6cb4323eb3, []int{0}
}
func (m *MyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyMessage.Unmarshal(m, b)
}
func (m *MyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyMessage.Marshal(b, m, deterministic)
}
func (m *MyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyMessage.Merge(m, src)
}
func (m *MyMessage) XXX_Size() int {
	return xxx_messageInfo_MyMessage.Size(m)
}
func (m *MyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_MyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_MyMessage proto.InternalMessageInfo

func (m *MyMessage) GetTheField() test.TheTestEnum {
	if m != nil {
		return m.TheField
	}
	return test.A
}

func init() {
	proto.RegisterType((*MyMessage)(nil), "enumprefix.MyMessage")
}

func init() { proto.RegisterFile("enumprefix.proto", fileDescriptor_d0d23e6cb4323eb3) }

var fileDescriptor_d0d23e6cb4323eb3 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xcd, 0x2b, 0xcd,
	0x2d, 0x28, 0x4a, 0x4d, 0xcb, 0xac, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x48, 0x69, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7,
	0xe7, 0xeb, 0x83, 0x95, 0x24, 0x95, 0xa6, 0xe9, 0x97, 0xa4, 0x16, 0x97, 0xe8, 0x97, 0x64, 0xa4,
	0x82, 0x68, 0x88, 0x46, 0x29, 0x5d, 0x9c, 0x8a, 0x41, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x28, 0x57,
	0x72, 0xe0, 0xe2, 0xf4, 0xad, 0xf4, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x32, 0xe6, 0xe2,
	0x08, 0xc9, 0x48, 0x75, 0xcb, 0x4c, 0xcd, 0x49, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0x12,
	0xd4, 0x03, 0x1b, 0x1d, 0x92, 0x91, 0x1a, 0x92, 0x5a, 0x5c, 0xe2, 0x9a, 0x57, 0x9a, 0xeb, 0xc4,
	0x72, 0xe2, 0x9e, 0x3c, 0x43, 0x10, 0x5c, 0x21, 0x20, 0x00, 0x00, 0xff, 0xff, 0xda, 0xd1, 0x3c,
	0x03, 0xbc, 0x00, 0x00, 0x00,
}
