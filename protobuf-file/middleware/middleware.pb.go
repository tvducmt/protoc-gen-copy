// Code generated by protoc-gen-go. DO NOT EDIT.
// source: middleware.proto

package middleware

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

type Hello struct {
	KA                   string   `protobuf:"bytes,1,opt,name=k_a,json=kA,proto3" json:"k_a,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hello) Reset()         { *m = Hello{} }
func (m *Hello) String() string { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()    {}
func (*Hello) Descriptor() ([]byte, []int) {
	return fileDescriptor_10e8b4c446957f50, []int{0}
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

func (m *Hello) GetKA() string {
	if m != nil {
		return m.KA
	}
	return ""
}

type Transaction struct {
	MA                   string   `protobuf:"bytes,1,opt,name=m_a,json=mA,proto3" json:"m_a,omitempty"`
	ZpB                  string   `protobuf:"bytes,2,opt,name=zp_b,json=zpB,proto3" json:"zp_b,omitempty"`
	Hello                *Hello   `protobuf:"bytes,3,opt,name=hello,proto3" json:"hello,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_10e8b4c446957f50, []int{1}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetMA() string {
	if m != nil {
		return m.MA
	}
	return ""
}

func (m *Transaction) GetZpB() string {
	if m != nil {
		return m.ZpB
	}
	return ""
}

func (m *Transaction) GetHello() *Hello {
	if m != nil {
		return m.Hello
	}
	return nil
}

type ListCITransactionsRequest struct {
	MTransId             string       `protobuf:"bytes,1,opt,name=m_trans_id,json=mTransId,proto3" json:"m_trans_id,omitempty"`
	ZpTransId            string       `protobuf:"bytes,2,opt,name=zp_trans_id,json=zpTransId,proto3" json:"zp_trans_id,omitempty"`
	MerchantId           string       `protobuf:"bytes,3,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	MId                  string       `protobuf:"bytes,4,opt,name=m_id,json=mId,proto3" json:"m_id,omitempty"`
	Data                 *Transaction `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListCITransactionsRequest) Reset()         { *m = ListCITransactionsRequest{} }
func (m *ListCITransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*ListCITransactionsRequest) ProtoMessage()    {}
func (*ListCITransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_10e8b4c446957f50, []int{2}
}

func (m *ListCITransactionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCITransactionsRequest.Unmarshal(m, b)
}
func (m *ListCITransactionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCITransactionsRequest.Marshal(b, m, deterministic)
}
func (m *ListCITransactionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCITransactionsRequest.Merge(m, src)
}
func (m *ListCITransactionsRequest) XXX_Size() int {
	return xxx_messageInfo_ListCITransactionsRequest.Size(m)
}
func (m *ListCITransactionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCITransactionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCITransactionsRequest proto.InternalMessageInfo

func (m *ListCITransactionsRequest) GetMTransId() string {
	if m != nil {
		return m.MTransId
	}
	return ""
}

func (m *ListCITransactionsRequest) GetZpTransId() string {
	if m != nil {
		return m.ZpTransId
	}
	return ""
}

func (m *ListCITransactionsRequest) GetMerchantId() string {
	if m != nil {
		return m.MerchantId
	}
	return ""
}

func (m *ListCITransactionsRequest) GetMId() string {
	if m != nil {
		return m.MId
	}
	return ""
}

func (m *ListCITransactionsRequest) GetData() *Transaction {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Hello)(nil), "middleware.Hello")
	proto.RegisterType((*Transaction)(nil), "middleware.Transaction")
	proto.RegisterType((*ListCITransactionsRequest)(nil), "middleware.ListCITransactionsRequest")
}

func init() { proto.RegisterFile("middleware.proto", fileDescriptor_10e8b4c446957f50) }

var fileDescriptor_10e8b4c446957f50 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xe9, 0xda, 0x8a, 0xfb, 0x7a, 0xd0, 0xe5, 0x62, 0x04, 0xd1, 0xd1, 0x8b, 0x03, 0x61,
	0x07, 0xfd, 0x05, 0xd3, 0x8b, 0x05, 0x4f, 0xc5, 0x93, 0x97, 0x92, 0x2d, 0x81, 0x85, 0xf5, 0x6b,
	0x62, 0x12, 0x11, 0xf2, 0xd3, 0xfc, 0x75, 0x92, 0x74, 0xa5, 0xb9, 0xbe, 0xcf, 0x47, 0x9e, 0x87,
	0xc0, 0x35, 0x4a, 0xce, 0x7b, 0xf1, 0xcb, 0x8c, 0xd8, 0x6a, 0xa3, 0x9c, 0x22, 0x30, 0x2f, 0x35,
	0x85, 0xf2, 0x5d, 0xf4, 0xbd, 0x22, 0x57, 0x90, 0x9f, 0x3a, 0x46, 0xb3, 0x75, 0xb6, 0x59, 0xb6,
	0x8b, 0xd3, 0xae, 0xfe, 0x82, 0xea, 0xd3, 0xb0, 0xc1, 0xb2, 0x83, 0x93, 0x6a, 0x08, 0x1c, 0x67,
	0x8e, 0x3b, 0xb2, 0x82, 0xc2, 0xeb, 0x6e, 0x4f, 0x17, 0x71, 0xc9, 0xbd, 0x7e, 0x25, 0x8f, 0x50,
	0x1e, 0xc3, 0x63, 0x34, 0x5f, 0x67, 0x9b, 0xea, 0x79, 0xb5, 0x4d, 0xd4, 0xd1, 0xd2, 0x8e, 0xbc,
	0xfe, 0xcb, 0xe0, 0xf6, 0x43, 0x5a, 0xf7, 0xd6, 0x24, 0x0a, 0xdb, 0x8a, 0xef, 0x1f, 0x61, 0x1d,
	0xb9, 0x03, 0xc0, 0xce, 0x05, 0xd0, 0x49, 0x7e, 0x36, 0x5e, 0x62, 0xbc, 0x6c, 0x38, 0xb9, 0x87,
	0xca, 0xeb, 0x19, 0x8f, 0xfa, 0xa5, 0xd7, 0x13, 0x7f, 0x80, 0x0a, 0x85, 0x39, 0x1c, 0xd9, 0xe0,
	0x02, 0xcf, 0x23, 0x87, 0x69, 0x6a, 0x78, 0x08, 0xc7, 0x40, 0x8a, 0x31, 0x1c, 0x1b, 0x4e, 0x9e,
	0xa0, 0xe0, 0xcc, 0x31, 0x5a, 0xc6, 0xee, 0x9b, 0xb4, 0x3b, 0x09, 0x6c, 0xe3, 0xd1, 0xfe, 0x22,
	0xfe, 0xe2, 0xcb, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0xc2, 0x8d, 0xbf, 0x59, 0x01, 0x00,
	0x00,
}
