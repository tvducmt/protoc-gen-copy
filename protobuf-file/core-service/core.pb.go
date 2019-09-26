// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core.proto

package core

import (
	fmt "fmt"
	_ "git.zapa.cloud/merchant-tools/helper/proto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type BODetailReconciliation struct {
	//     // message TimeAttribute {
	// //     //     google.protobuf.Timestamp retry_time = 1;
	// //     // }
	CountableAttribute   *BODetailReconciliation_CountableAttribute `protobuf:"bytes,14,opt,name=countable_attribute,json=countableAttribute,proto3" json:"countable_attribute,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *BODetailReconciliation) Reset()         { *m = BODetailReconciliation{} }
func (m *BODetailReconciliation) String() string { return proto.CompactTextString(m) }
func (*BODetailReconciliation) ProtoMessage()    {}
func (*BODetailReconciliation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{0}
}

func (m *BODetailReconciliation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BODetailReconciliation.Unmarshal(m, b)
}
func (m *BODetailReconciliation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BODetailReconciliation.Marshal(b, m, deterministic)
}
func (m *BODetailReconciliation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BODetailReconciliation.Merge(m, src)
}
func (m *BODetailReconciliation) XXX_Size() int {
	return xxx_messageInfo_BODetailReconciliation.Size(m)
}
func (m *BODetailReconciliation) XXX_DiscardUnknown() {
	xxx_messageInfo_BODetailReconciliation.DiscardUnknown(m)
}

var xxx_messageInfo_BODetailReconciliation proto.InternalMessageInfo

func (m *BODetailReconciliation) GetCountableAttribute() *BODetailReconciliation_CountableAttribute {
	if m != nil {
		return m.CountableAttribute
	}
	return nil
}

type BODetailReconciliation_CountableAttribute struct {
	// @inject_tag: es:"merchantRefundAmount"
	// int64 merchant_refund_amount = 19;
	// Data data = 20;
	//
	// message Data {
	//     string k1 =1;
	//     string k2=2;
	//     Hello hello=3;
	//     Hi hi =4;
	// }
	// message Hi{
	//     int32 hk =1;
	// }
	// @inject_tag: es:"refundPercentFeeNotVat"
	// double return_percent_fee_not_vat = 28;
	// // @inject_tag: es:"tpeBankCode"
	TpeBankCode string `protobuf:"bytes,36,opt,name=tpe_bank_code,json=tpeBankCode,proto3" json:"tpe_bank_code,omitempty"`
	Hello       *Hello `protobuf:"bytes,3,opt,name=hello,proto3" json:"hello,omitempty"`
	// @inject_tag: es:"itemCount"
	ItemCount            int32    `protobuf:"varint,38,opt,name=item_count,json=itemCount,proto3" json:"item_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BODetailReconciliation_CountableAttribute) Reset() {
	*m = BODetailReconciliation_CountableAttribute{}
}
func (m *BODetailReconciliation_CountableAttribute) String() string { return proto.CompactTextString(m) }
func (*BODetailReconciliation_CountableAttribute) ProtoMessage()    {}
func (*BODetailReconciliation_CountableAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{0, 0}
}

func (m *BODetailReconciliation_CountableAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BODetailReconciliation_CountableAttribute.Unmarshal(m, b)
}
func (m *BODetailReconciliation_CountableAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BODetailReconciliation_CountableAttribute.Marshal(b, m, deterministic)
}
func (m *BODetailReconciliation_CountableAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BODetailReconciliation_CountableAttribute.Merge(m, src)
}
func (m *BODetailReconciliation_CountableAttribute) XXX_Size() int {
	return xxx_messageInfo_BODetailReconciliation_CountableAttribute.Size(m)
}
func (m *BODetailReconciliation_CountableAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_BODetailReconciliation_CountableAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_BODetailReconciliation_CountableAttribute proto.InternalMessageInfo

func (m *BODetailReconciliation_CountableAttribute) GetTpeBankCode() string {
	if m != nil {
		return m.TpeBankCode
	}
	return ""
}

func (m *BODetailReconciliation_CountableAttribute) GetHello() *Hello {
	if m != nil {
		return m.Hello
	}
	return nil
}

func (m *BODetailReconciliation_CountableAttribute) GetItemCount() int32 {
	if m != nil {
		return m.ItemCount
	}
	return 0
}

type Hello struct {
	H1                   string   `protobuf:"bytes,1,opt,name=h1,proto3" json:"h1,omitempty"`
	H2                   string   `protobuf:"bytes,2,opt,name=h2,proto3" json:"h2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hello) Reset()         { *m = Hello{} }
func (m *Hello) String() string { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()    {}
func (*Hello) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{1}
}

func (m *Hello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hello.Unmarshal(m, b)
}
func (m *Hello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hello.Marshal(b, m, deterministic)
}
func (m *Hello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hello.Merge(m, src)
}
func (m *Hello) XXX_Size() int {
	return xxx_messageInfo_Hello.Size(m)
}
func (m *Hello) XXX_DiscardUnknown() {
	xxx_messageInfo_Hello.DiscardUnknown(m)
}

var xxx_messageInfo_Hello proto.InternalMessageInfo

func (m *Hello) GetH1() string {
	if m != nil {
		return m.H1
	}
	return ""
}

func (m *Hello) GetH2() string {
	if m != nil {
		return m.H2
	}
	return ""
}

func init() {
	proto.RegisterType((*BODetailReconciliation)(nil), "core.BODetailReconciliation")
	proto.RegisterType((*BODetailReconciliation_CountableAttribute)(nil), "core.BODetailReconciliation.CountableAttribute")
	proto.RegisterType((*Hello)(nil), "core.Hello")
}

func init() { proto.RegisterFile("core.proto", fileDescriptor_f7e43720d1edc0fe) }

var fileDescriptor_f7e43720d1edc0fe = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x49, 0xff, 0xbf, 0x42, 0xa7, 0xd8, 0xc5, 0x08, 0x32, 0x14, 0xc4, 0x1a, 0x44, 0xb3,
	0x31, 0x43, 0x23, 0x3e, 0x80, 0xad, 0x0b, 0x77, 0x42, 0x5e, 0x20, 0x4e, 0x26, 0xd7, 0x64, 0xe8,
	0x64, 0xee, 0x90, 0xde, 0x2c, 0xec, 0x83, 0xbb, 0x96, 0x4c, 0xda, 0x55, 0xdd, 0xcd, 0x3d, 0xe7,
	0x7c, 0x9c, 0xc3, 0x30, 0xa6, 0xb1, 0x83, 0xd4, 0x77, 0x48, 0xc8, 0xff, 0x0f, 0xef, 0xe5, 0x4b,
	0x6d, 0x28, 0x3d, 0x28, 0xaf, 0x52, 0x6d, 0xb1, 0xaf, 0x64, 0x0b, 0x9d, 0x6e, 0x94, 0xa3, 0x27,
	0x42, 0xb4, 0x7b, 0xd9, 0x80, 0xf5, 0xd0, 0xc9, 0x40, 0x48, 0xfa, 0xf6, 0x47, 0x78, 0x79, 0x5b,
	0x23, 0xd6, 0x16, 0x46, 0xa3, 0xec, 0xbf, 0x24, 0x99, 0x16, 0xf6, 0xa4, 0x5a, 0x3f, 0x06, 0xe2,
	0x9f, 0x88, 0x5d, 0x6f, 0x3e, 0xde, 0x80, 0x94, 0xb1, 0x39, 0x68, 0x74, 0xda, 0x58, 0xa3, 0xc8,
	0xa0, 0xe3, 0x9f, 0xec, 0x4a, 0x63, 0xef, 0x48, 0x95, 0x16, 0x0a, 0x45, 0xd4, 0x99, 0xb2, 0x27,
	0x10, 0x8b, 0x55, 0x94, 0xcc, 0x33, 0x99, 0x86, 0x89, 0x7f, 0xa3, 0xe9, 0xf6, 0xc4, 0xbd, 0x9e,
	0xb0, 0x9c, 0xeb, 0x33, 0x6d, 0x79, 0x60, 0xfc, 0x3c, 0xc9, 0x63, 0x76, 0x49, 0x1e, 0x8a, 0x52,
	0xb9, 0x5d, 0xa1, 0xb1, 0x02, 0x71, 0xbf, 0x8a, 0x92, 0x59, 0x3e, 0x27, 0x0f, 0x1b, 0xe5, 0x76,
	0x5b, 0xac, 0x80, 0xdf, 0xb1, 0x69, 0x03, 0xd6, 0xa2, 0xf8, 0x17, 0xd6, 0xcc, 0xc7, 0x35, 0xef,
	0x83, 0x94, 0x8f, 0x0e, 0xbf, 0x61, 0xcc, 0x10, 0xb4, 0x45, 0xe8, 0x15, 0x0f, 0xab, 0x28, 0x99,
	0xe6, 0xb3, 0x41, 0x09, 0x95, 0xf1, 0x23, 0x9b, 0x86, 0x38, 0x5f, 0xb0, 0x49, 0xb3, 0x16, 0x51,
	0xe8, 0x98, 0x34, 0xeb, 0x70, 0x67, 0x62, 0x72, 0xbc, 0xb3, 0xf2, 0x22, 0x7c, 0xd4, 0xf3, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x9e, 0xa7, 0x1d, 0x94, 0x01, 0x00, 0x00,
}
