// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core.proto

package core

import (
	fmt "fmt"
	proto1 "git.zapa.cloud/merchant-tools/helper/proto"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	FromDate *proto1.Date `protobuf:"bytes,1,opt,name=from_date,json=fromDate,proto3" json:"from_date,omitempty"`
	//     // message TimeAttribute {
	// //     //     google.protobuf.Timestamp retry_time = 1;
	// //     // }
	CountableAttribute *BODetailReconciliation_CountableAttribute `protobuf:"bytes,14,opt,name=countable_attribute,json=countableAttribute,proto3" json:"countable_attribute,omitempty"`
	// enum VoucherCode{
	//     VOUCHER_CODE_UNSPECIFICED    = 0;
	//     VC_YES   = 1;
	//     VC_NO    = 2;
	// }
	TransTime            *timestamp.Timestamp `protobuf:"bytes,24,opt,name=trans_time,json=transTime,proto3" json:"trans_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
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

func (m *BODetailReconciliation) GetFromDate() *proto1.Date {
	if m != nil {
		return m.FromDate
	}
	return nil
}

func (m *BODetailReconciliation) GetCountableAttribute() *BODetailReconciliation_CountableAttribute {
	if m != nil {
		return m.CountableAttribute
	}
	return nil
}

func (m *BODetailReconciliation) GetTransTime() *timestamp.Timestamp {
	if m != nil {
		return m.TransTime
	}
	return nil
}

type BODetailReconciliation_CountableAttribute struct {
	// @inject_tag: es:"merchantRefundAmount"
	// int64 merchant_refund_amount = 19;
	Data *BODetailReconciliation_CountableAttribute_Data `protobuf:"bytes,20,opt,name=data,proto3" json:"data,omitempty"`
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

func (m *BODetailReconciliation_CountableAttribute) GetData() *BODetailReconciliation_CountableAttribute_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

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

type BODetailReconciliation_CountableAttribute_Data struct {
	K1                   string                                        `protobuf:"bytes,1,opt,name=k1,proto3" json:"k1,omitempty"`
	K2                   string                                        `protobuf:"bytes,2,opt,name=k2,proto3" json:"k2,omitempty"`
	Hello                *Hello                                        `protobuf:"bytes,3,opt,name=hello,proto3" json:"hello,omitempty"`
	Hi                   *BODetailReconciliation_CountableAttribute_Hi `protobuf:"bytes,4,opt,name=hi,proto3" json:"hi,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *BODetailReconciliation_CountableAttribute_Data) Reset() {
	*m = BODetailReconciliation_CountableAttribute_Data{}
}
func (m *BODetailReconciliation_CountableAttribute_Data) String() string {
	return proto.CompactTextString(m)
}
func (*BODetailReconciliation_CountableAttribute_Data) ProtoMessage() {}
func (*BODetailReconciliation_CountableAttribute_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{0, 0, 0}
}

func (m *BODetailReconciliation_CountableAttribute_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BODetailReconciliation_CountableAttribute_Data.Unmarshal(m, b)
}
func (m *BODetailReconciliation_CountableAttribute_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BODetailReconciliation_CountableAttribute_Data.Marshal(b, m, deterministic)
}
func (m *BODetailReconciliation_CountableAttribute_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BODetailReconciliation_CountableAttribute_Data.Merge(m, src)
}
func (m *BODetailReconciliation_CountableAttribute_Data) XXX_Size() int {
	return xxx_messageInfo_BODetailReconciliation_CountableAttribute_Data.Size(m)
}
func (m *BODetailReconciliation_CountableAttribute_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_BODetailReconciliation_CountableAttribute_Data.DiscardUnknown(m)
}

var xxx_messageInfo_BODetailReconciliation_CountableAttribute_Data proto.InternalMessageInfo

func (m *BODetailReconciliation_CountableAttribute_Data) GetK1() string {
	if m != nil {
		return m.K1
	}
	return ""
}

func (m *BODetailReconciliation_CountableAttribute_Data) GetK2() string {
	if m != nil {
		return m.K2
	}
	return ""
}

func (m *BODetailReconciliation_CountableAttribute_Data) GetHello() *Hello {
	if m != nil {
		return m.Hello
	}
	return nil
}

func (m *BODetailReconciliation_CountableAttribute_Data) GetHi() *BODetailReconciliation_CountableAttribute_Hi {
	if m != nil {
		return m.Hi
	}
	return nil
}

type BODetailReconciliation_CountableAttribute_Hi struct {
	Hk                   int32    `protobuf:"varint,1,opt,name=hk,proto3" json:"hk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BODetailReconciliation_CountableAttribute_Hi) Reset() {
	*m = BODetailReconciliation_CountableAttribute_Hi{}
}
func (m *BODetailReconciliation_CountableAttribute_Hi) String() string {
	return proto.CompactTextString(m)
}
func (*BODetailReconciliation_CountableAttribute_Hi) ProtoMessage() {}
func (*BODetailReconciliation_CountableAttribute_Hi) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{0, 0, 1}
}

func (m *BODetailReconciliation_CountableAttribute_Hi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BODetailReconciliation_CountableAttribute_Hi.Unmarshal(m, b)
}
func (m *BODetailReconciliation_CountableAttribute_Hi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BODetailReconciliation_CountableAttribute_Hi.Marshal(b, m, deterministic)
}
func (m *BODetailReconciliation_CountableAttribute_Hi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BODetailReconciliation_CountableAttribute_Hi.Merge(m, src)
}
func (m *BODetailReconciliation_CountableAttribute_Hi) XXX_Size() int {
	return xxx_messageInfo_BODetailReconciliation_CountableAttribute_Hi.Size(m)
}
func (m *BODetailReconciliation_CountableAttribute_Hi) XXX_DiscardUnknown() {
	xxx_messageInfo_BODetailReconciliation_CountableAttribute_Hi.DiscardUnknown(m)
}

var xxx_messageInfo_BODetailReconciliation_CountableAttribute_Hi proto.InternalMessageInfo

func (m *BODetailReconciliation_CountableAttribute_Hi) GetHk() int32 {
	if m != nil {
		return m.Hk
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
	proto.RegisterType((*BODetailReconciliation_CountableAttribute_Data)(nil), "core.BODetailReconciliation.CountableAttribute.Data")
	proto.RegisterType((*BODetailReconciliation_CountableAttribute_Hi)(nil), "core.BODetailReconciliation.CountableAttribute.Hi")
	proto.RegisterType((*Hello)(nil), "core.Hello")
}

func init() { proto.RegisterFile("core.proto", fileDescriptor_f7e43720d1edc0fe) }

var fileDescriptor_f7e43720d1edc0fe = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x8e, 0x94, 0x40,
	0x10, 0xc6, 0x03, 0x32, 0x46, 0x6a, 0xe3, 0x26, 0xb6, 0x1b, 0xd3, 0x21, 0x31, 0x8e, 0x13, 0xa3,
	0x73, 0x11, 0xb2, 0xa8, 0x07, 0x8f, 0xce, 0xee, 0x61, 0x6e, 0x26, 0x1d, 0xef, 0xd8, 0x40, 0xed,
	0x74, 0x87, 0x86, 0x26, 0x4c, 0xcd, 0x41, 0xdf, 0xc1, 0x97, 0xf3, 0x31, 0x7c, 0x0a, 0xd3, 0xc5,
	0xb2, 0x97, 0xf5, 0xe0, 0xde, 0xa8, 0x3f, 0x5f, 0xfd, 0xbe, 0x8f, 0x06, 0x68, 0xfc, 0x84, 0xf9,
	0x38, 0x79, 0xf2, 0x22, 0x09, 0xdf, 0xd9, 0xa7, 0x83, 0xa5, 0xfc, 0xa7, 0x1e, 0x75, 0xde, 0x38,
	0x7f, 0x6a, 0x8b, 0x1e, 0xa7, 0xc6, 0xe8, 0x81, 0xde, 0x93, 0xf7, 0xee, 0x58, 0x18, 0x74, 0x23,
	0x4e, 0x05, 0x2b, 0x0a, 0xfa, 0x31, 0xde, 0x8a, 0xb3, 0x57, 0x07, 0xef, 0x0f, 0x0e, 0xe7, 0x41,
	0x7d, 0xba, 0x29, 0xc8, 0xf6, 0x78, 0x24, 0xdd, 0x8f, 0xf3, 0xc2, 0xe6, 0x77, 0x02, 0x2f, 0x76,
	0x5f, 0xaf, 0x91, 0xb4, 0x75, 0x0a, 0x1b, 0x3f, 0x34, 0xd6, 0x59, 0x4d, 0xd6, 0x0f, 0x22, 0x87,
	0xf4, 0x66, 0xf2, 0x7d, 0xd5, 0x6a, 0x42, 0x19, 0xad, 0xa3, 0xed, 0x59, 0xf9, 0x2c, 0x9f, 0xef,
	0xe5, 0x8c, 0xb8, 0xd6, 0x84, 0xea, 0x49, 0xd8, 0x09, 0x5f, 0xe2, 0x3b, 0x3c, 0x6f, 0xfc, 0x69,
	0x20, 0x5d, 0x3b, 0xac, 0x34, 0xd1, 0x64, 0xeb, 0x13, 0xa1, 0x3c, 0x67, 0x65, 0x91, 0x73, 0xa4,
	0x7f, 0xa3, 0xf2, 0xab, 0x45, 0xf7, 0x65, 0x91, 0x29, 0xd1, 0xdc, 0xeb, 0x89, 0xcf, 0x00, 0x34,
	0xe9, 0xe1, 0x58, 0x85, 0x14, 0x52, 0xf2, 0xe1, 0x6c, 0xb1, 0xb4, 0x44, 0xcc, 0xbf, 0x2d, 0x11,
	0x55, 0xca, 0xdb, 0xa1, 0xce, 0xfe, 0xc4, 0x20, 0xee, 0x53, 0xc4, 0x1e, 0x92, 0x56, 0x93, 0x96,
	0x17, 0x7c, 0xeb, 0xe3, 0x03, 0x4d, 0x86, 0x3f, 0xa0, 0x15, 0x5f, 0x10, 0x1b, 0x78, 0x4a, 0x23,
	0x56, 0xb5, 0x1e, 0xba, 0xaa, 0xf1, 0x2d, 0xca, 0x37, 0xeb, 0x68, 0x9b, 0xaa, 0x33, 0x1a, 0x71,
	0xa7, 0x87, 0xee, 0xca, 0xb7, 0x28, 0x5e, 0xc3, 0xca, 0xa0, 0x73, 0x5e, 0x3e, 0x62, 0xdc, 0xd9,
	0x8c, 0xdb, 0x87, 0x96, 0x9a, 0x27, 0xe2, 0x25, 0x80, 0x25, 0xec, 0x2b, 0x4e, 0x2f, 0xdf, 0xae,
	0xa3, 0xed, 0x4a, 0xa5, 0xa1, 0xc3, 0xf4, 0xec, 0x57, 0x04, 0x49, 0x80, 0x8a, 0x73, 0x88, 0xbb,
	0x4b, 0x7e, 0x95, 0x54, 0xc5, 0xdd, 0x25, 0xd7, 0xa5, 0x8c, 0x6f, 0xeb, 0xf2, 0x7f, 0x50, 0x3b,
	0x88, 0x8d, 0x95, 0x09, 0xcf, 0xcb, 0x87, 0x26, 0xdf, 0x5b, 0x15, 0x1b, 0x9b, 0x5d, 0x40, 0xbc,
	0xb7, 0x01, 0x6e, 0x3a, 0x36, 0xb3, 0x52, 0xb1, 0xe9, 0x36, 0xef, 0x60, 0xc5, 0x24, 0x1e, 0xdc,
	0xb9, 0x34, 0xec, 0xd2, 0xdc, 0xb9, 0x34, 0x65, 0xfd, 0x98, 0x1f, 0xed, 0xc3, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x2c, 0x33, 0xa9, 0xde, 0xf0, 0x02, 0x00, 0x00,
}
