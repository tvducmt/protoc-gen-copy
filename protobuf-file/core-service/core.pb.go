// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core.proto

package core_service

import (
	fmt "fmt"
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

type RefundTransLogApi int32

const (
	RefundTransLogApi_RTA_UNSPECIFIED RefundTransLogApi = 0
	RefundTransLogApi_RTA_FULL        RefundTransLogApi = 1
	RefundTransLogApi_RTA_PARTIAL     RefundTransLogApi = 2
)

var RefundTransLogApi_name = map[int32]string{
	0: "RTA_UNSPECIFIED",
	1: "RTA_FULL",
	2: "RTA_PARTIAL",
}

var RefundTransLogApi_value = map[string]int32{
	"RTA_UNSPECIFIED": 0,
	"RTA_FULL":        1,
	"RTA_PARTIAL":     2,
}

func (x RefundTransLogApi) String() string {
	return proto.EnumName(RefundTransLogApi_name, int32(x))
}

func (RefundTransLogApi) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{0}
}

// `ZaloPayType` is type transaction from app or gateway
type ZaloPayType int32

const (
	//Không xác định
	ZaloPayType_ZP_TYPE_UNSPECIFIED ZaloPayType = 0
	//Ứng dụng zaloPay
	ZaloPayType_ZP_TYPE_APP ZaloPayType = 1
	//Cổng ZaloPay
	ZaloPayType_ZP_TYPE_GATEWAY ZaloPayType = 2
	//Tất cả
	ZaloPayType_ZP_TYPE_ALL ZaloPayType = 3
)

var ZaloPayType_name = map[int32]string{
	0: "ZP_TYPE_UNSPECIFIED",
	1: "ZP_TYPE_APP",
	2: "ZP_TYPE_GATEWAY",
	3: "ZP_TYPE_ALL",
}

var ZaloPayType_value = map[string]int32{
	"ZP_TYPE_UNSPECIFIED": 0,
	"ZP_TYPE_APP":         1,
	"ZP_TYPE_GATEWAY":     2,
	"ZP_TYPE_ALL":         3,
}

func (x ZaloPayType) String() string {
	return proto.EnumName(ZaloPayType_name, int32(x))
}

func (ZaloPayType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{1}
}

// `TransStatus` is status of a transaction
type TransStatus int32

const (
	TransStatus_TS_UNSPECIFIED TransStatus = 0
	TransStatus_TS_FAIL        TransStatus = 1
	TransStatus_TS_SUCCESSFUL  TransStatus = 2
	TransStatus_TS_ALL         TransStatus = 3
	TransStatus_TS_NEW         TransStatus = 4
	TransStatus_TS_UNFINISHED  TransStatus = 5
)

var TransStatus_name = map[int32]string{
	0: "TS_UNSPECIFIED",
	1: "TS_FAIL",
	2: "TS_SUCCESSFUL",
	3: "TS_ALL",
	4: "TS_NEW",
	5: "TS_UNFINISHED",
}

var TransStatus_value = map[string]int32{
	"TS_UNSPECIFIED": 0,
	"TS_FAIL":        1,
	"TS_SUCCESSFUL":  2,
	"TS_ALL":         3,
	"TS_NEW":         4,
	"TS_UNFINISHED":  5,
}

func (x TransStatus) String() string {
	return proto.EnumName(TransStatus_name, int32(x))
}

func (TransStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{2}
}

//hình thức thanh toán
type PaymentMethod int32

const (
	PaymentMethod_PM_UNSPECIFIED PaymentMethod = 0
	PaymentMethod_PM_QR_PAY      PaymentMethod = 1
	PaymentMethod_PM_DIRECT      PaymentMethod = 2
	PaymentMethod_PM_QUICKPAY    PaymentMethod = 3
	PaymentMethod_PM_AUTODEBIT   PaymentMethod = 4
	PaymentMethod_PM_UNKNOWN     PaymentMethod = 5
)

var PaymentMethod_name = map[int32]string{
	0: "PM_UNSPECIFIED",
	1: "PM_QR_PAY",
	2: "PM_DIRECT",
	3: "PM_QUICKPAY",
	4: "PM_AUTODEBIT",
	5: "PM_UNKNOWN",
}

var PaymentMethod_value = map[string]int32{
	"PM_UNSPECIFIED": 0,
	"PM_QR_PAY":      1,
	"PM_DIRECT":      2,
	"PM_QUICKPAY":    3,
	"PM_AUTODEBIT":   4,
	"PM_UNKNOWN":     5,
}

func (x PaymentMethod) String() string {
	return proto.EnumName(PaymentMethod_name, int32(x))
}

func (PaymentMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{3}
}

// Profile p1 = 5;
// message Profile{
//     string Name =1;
// }
type BODetailReconciliation_VoucherCode int32

const (
	BODetailReconciliation_VOUCHER_CODE_UNSPECIFICED BODetailReconciliation_VoucherCode = 0
	BODetailReconciliation_VC_YES                    BODetailReconciliation_VoucherCode = 1
	BODetailReconciliation_VC_NO                     BODetailReconciliation_VoucherCode = 2
)

var BODetailReconciliation_VoucherCode_name = map[int32]string{
	0: "VOUCHER_CODE_UNSPECIFICED",
	1: "VC_YES",
	2: "VC_NO",
}

var BODetailReconciliation_VoucherCode_value = map[string]int32{
	"VOUCHER_CODE_UNSPECIFICED": 0,
	"VC_YES":                    1,
	"VC_NO":                     2,
}

func (x BODetailReconciliation_VoucherCode) String() string {
	return proto.EnumName(BODetailReconciliation_VoucherCode_name, int32(x))
}

func (BODetailReconciliation_VoucherCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{1, 0}
}

type OutnerField struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutnerField) Reset()         { *m = OutnerField{} }
func (m *OutnerField) String() string { return proto.CompactTextString(m) }
func (*OutnerField) ProtoMessage()    {}
func (*OutnerField) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{0}
}

func (m *OutnerField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutnerField.Unmarshal(m, b)
}
func (m *OutnerField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutnerField.Marshal(b, m, deterministic)
}
func (m *OutnerField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutnerField.Merge(m, src)
}
func (m *OutnerField) XXX_Size() int {
	return xxx_messageInfo_OutnerField.Size(m)
}
func (m *OutnerField) XXX_DiscardUnknown() {
	xxx_messageInfo_OutnerField.DiscardUnknown(m)
}

var xxx_messageInfo_OutnerField proto.InternalMessageInfo

func (m *OutnerField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type BODetailReconciliation struct {
	OutnerFiled *OutnerField `protobuf:"bytes,1,opt,name=outner_filed,json=outnerFiled,proto3" json:"outner_filed,omitempty"`
	// google.type.Date from_date               = 1;
	TransTime            *timestamp.Timestamp                       `protobuf:"bytes,2,opt,name=trans_time,json=transTime,proto3" json:"trans_time,omitempty"`
	ZpSystemName         string                                     `protobuf:"bytes,3,opt,name=zp_system_name,json=zpSystemName,proto3" json:"zp_system_name,omitempty"`
	VoucherCode          BODetailReconciliation_VoucherCode         `protobuf:"varint,10,opt,name=voucher_code,json=voucherCode,proto3,enum=core_service.BODetailReconciliation_VoucherCode" json:"voucher_code,omitempty"`
	TimeAttribute        *BODetailReconciliation_TimeAttribute      `protobuf:"bytes,13,opt,name=time_attribute,json=timeAttribute,proto3" json:"time_attribute,omitempty"`
	CountableAttribute   *BODetailReconciliation_CountableAttribute `protobuf:"bytes,14,opt,name=countable_attribute,json=countableAttribute,proto3" json:"countable_attribute,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *BODetailReconciliation) Reset()         { *m = BODetailReconciliation{} }
func (m *BODetailReconciliation) String() string { return proto.CompactTextString(m) }
func (*BODetailReconciliation) ProtoMessage()    {}
func (*BODetailReconciliation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{1}
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

func (m *BODetailReconciliation) GetOutnerFiled() *OutnerField {
	if m != nil {
		return m.OutnerFiled
	}
	return nil
}

func (m *BODetailReconciliation) GetTransTime() *timestamp.Timestamp {
	if m != nil {
		return m.TransTime
	}
	return nil
}

func (m *BODetailReconciliation) GetZpSystemName() string {
	if m != nil {
		return m.ZpSystemName
	}
	return ""
}

func (m *BODetailReconciliation) GetVoucherCode() BODetailReconciliation_VoucherCode {
	if m != nil {
		return m.VoucherCode
	}
	return BODetailReconciliation_VOUCHER_CODE_UNSPECIFICED
}

func (m *BODetailReconciliation) GetTimeAttribute() *BODetailReconciliation_TimeAttribute {
	if m != nil {
		return m.TimeAttribute
	}
	return nil
}

func (m *BODetailReconciliation) GetCountableAttribute() *BODetailReconciliation_CountableAttribute {
	if m != nil {
		return m.CountableAttribute
	}
	return nil
}

type BODetailReconciliation_TimeAttribute struct {
	RetryTime            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=retry_time,json=retryTime,proto3" json:"retry_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BODetailReconciliation_TimeAttribute) Reset()         { *m = BODetailReconciliation_TimeAttribute{} }
func (m *BODetailReconciliation_TimeAttribute) String() string { return proto.CompactTextString(m) }
func (*BODetailReconciliation_TimeAttribute) ProtoMessage()    {}
func (*BODetailReconciliation_TimeAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{1, 0}
}

func (m *BODetailReconciliation_TimeAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BODetailReconciliation_TimeAttribute.Unmarshal(m, b)
}
func (m *BODetailReconciliation_TimeAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BODetailReconciliation_TimeAttribute.Marshal(b, m, deterministic)
}
func (m *BODetailReconciliation_TimeAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BODetailReconciliation_TimeAttribute.Merge(m, src)
}
func (m *BODetailReconciliation_TimeAttribute) XXX_Size() int {
	return xxx_messageInfo_BODetailReconciliation_TimeAttribute.Size(m)
}
func (m *BODetailReconciliation_TimeAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_BODetailReconciliation_TimeAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_BODetailReconciliation_TimeAttribute proto.InternalMessageInfo

func (m *BODetailReconciliation_TimeAttribute) GetRetryTime() *timestamp.Timestamp {
	if m != nil {
		return m.RetryTime
	}
	return nil
}

type BODetailReconciliation_CountableAttribute struct {
	RefundApi            *RefundApi                                             `protobuf:"bytes,18,opt,name=refund_api,json=refundApi,proto3" json:"refund_api,omitempty"`
	Data                 *BODetailReconciliation_CountableAttribute_Transaction `protobuf:"bytes,20,opt,name=data,proto3" json:"data,omitempty"`
	TpeBankCode          string                                                 `protobuf:"bytes,36,opt,name=tpe_bank_code,json=tpeBankCode,proto3" json:"tpe_bank_code,omitempty"`
	Stm                  *Statement                                             `protobuf:"bytes,3,opt,name=stm,proto3" json:"stm,omitempty"`
	ItemCount            int32                                                  `protobuf:"varint,38,opt,name=item_count,json=itemCount,proto3" json:"item_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                               `json:"-"`
	XXX_unrecognized     []byte                                                 `json:"-"`
	XXX_sizecache        int32                                                  `json:"-"`
}

func (m *BODetailReconciliation_CountableAttribute) Reset() {
	*m = BODetailReconciliation_CountableAttribute{}
}
func (m *BODetailReconciliation_CountableAttribute) String() string { return proto.CompactTextString(m) }
func (*BODetailReconciliation_CountableAttribute) ProtoMessage()    {}
func (*BODetailReconciliation_CountableAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{1, 1}
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

func (m *BODetailReconciliation_CountableAttribute) GetRefundApi() *RefundApi {
	if m != nil {
		return m.RefundApi
	}
	return nil
}

func (m *BODetailReconciliation_CountableAttribute) GetData() *BODetailReconciliation_CountableAttribute_Transaction {
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

func (m *BODetailReconciliation_CountableAttribute) GetStm() *Statement {
	if m != nil {
		return m.Stm
	}
	return nil
}

func (m *BODetailReconciliation_CountableAttribute) GetItemCount() int32 {
	if m != nil {
		return m.ItemCount
	}
	return 0
}

type BODetailReconciliation_CountableAttribute_Transaction struct {
	K1                   string     `protobuf:"bytes,1,opt,name=k1,proto3" json:"k1,omitempty"`
	K2                   string     `protobuf:"bytes,2,opt,name=k2,proto3" json:"k2,omitempty"`
	Stm                  *Statement `protobuf:"bytes,3,opt,name=stm,proto3" json:"stm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BODetailReconciliation_CountableAttribute_Transaction) Reset() {
	*m = BODetailReconciliation_CountableAttribute_Transaction{}
}
func (m *BODetailReconciliation_CountableAttribute_Transaction) String() string {
	return proto.CompactTextString(m)
}
func (*BODetailReconciliation_CountableAttribute_Transaction) ProtoMessage() {}
func (*BODetailReconciliation_CountableAttribute_Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{1, 1, 0}
}

func (m *BODetailReconciliation_CountableAttribute_Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BODetailReconciliation_CountableAttribute_Transaction.Unmarshal(m, b)
}
func (m *BODetailReconciliation_CountableAttribute_Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BODetailReconciliation_CountableAttribute_Transaction.Marshal(b, m, deterministic)
}
func (m *BODetailReconciliation_CountableAttribute_Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BODetailReconciliation_CountableAttribute_Transaction.Merge(m, src)
}
func (m *BODetailReconciliation_CountableAttribute_Transaction) XXX_Size() int {
	return xxx_messageInfo_BODetailReconciliation_CountableAttribute_Transaction.Size(m)
}
func (m *BODetailReconciliation_CountableAttribute_Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_BODetailReconciliation_CountableAttribute_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_BODetailReconciliation_CountableAttribute_Transaction proto.InternalMessageInfo

func (m *BODetailReconciliation_CountableAttribute_Transaction) GetK1() string {
	if m != nil {
		return m.K1
	}
	return ""
}

func (m *BODetailReconciliation_CountableAttribute_Transaction) GetK2() string {
	if m != nil {
		return m.K2
	}
	return ""
}

func (m *BODetailReconciliation_CountableAttribute_Transaction) GetStm() *Statement {
	if m != nil {
		return m.Stm
	}
	return nil
}

type Statement struct {
	H1                   string   `protobuf:"bytes,1,opt,name=h1,proto3" json:"h1,omitempty"`
	H2                   string   `protobuf:"bytes,2,opt,name=h2,proto3" json:"h2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Statement) Reset()         { *m = Statement{} }
func (m *Statement) String() string { return proto.CompactTextString(m) }
func (*Statement) ProtoMessage()    {}
func (*Statement) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{2}
}

func (m *Statement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Statement.Unmarshal(m, b)
}
func (m *Statement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Statement.Marshal(b, m, deterministic)
}
func (m *Statement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Statement.Merge(m, src)
}
func (m *Statement) XXX_Size() int {
	return xxx_messageInfo_Statement.Size(m)
}
func (m *Statement) XXX_DiscardUnknown() {
	xxx_messageInfo_Statement.DiscardUnknown(m)
}

var xxx_messageInfo_Statement proto.InternalMessageInfo

func (m *Statement) GetH1() string {
	if m != nil {
		return m.H1
	}
	return ""
}

func (m *Statement) GetH2() string {
	if m != nil {
		return m.H2
	}
	return ""
}

type Statement1 struct {
	H1                   string   `protobuf:"bytes,1,opt,name=h1,proto3" json:"h1,omitempty"`
	H2                   string   `protobuf:"bytes,2,opt,name=h2,proto3" json:"h2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Statement1) Reset()         { *m = Statement1{} }
func (m *Statement1) String() string { return proto.CompactTextString(m) }
func (*Statement1) ProtoMessage()    {}
func (*Statement1) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{3}
}

func (m *Statement1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Statement1.Unmarshal(m, b)
}
func (m *Statement1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Statement1.Marshal(b, m, deterministic)
}
func (m *Statement1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Statement1.Merge(m, src)
}
func (m *Statement1) XXX_Size() int {
	return xxx_messageInfo_Statement1.Size(m)
}
func (m *Statement1) XXX_DiscardUnknown() {
	xxx_messageInfo_Statement1.DiscardUnknown(m)
}

var xxx_messageInfo_Statement1 proto.InternalMessageInfo

func (m *Statement1) GetH1() string {
	if m != nil {
		return m.H1
	}
	return ""
}

func (m *Statement1) GetH2() string {
	if m != nil {
		return m.H2
	}
	return ""
}

type RefundApi struct {
	Stm1                 *Statement1 `protobuf:"bytes,3,opt,name=stm1,proto3" json:"stm1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RefundApi) Reset()         { *m = RefundApi{} }
func (m *RefundApi) String() string { return proto.CompactTextString(m) }
func (*RefundApi) ProtoMessage()    {}
func (*RefundApi) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{4}
}

func (m *RefundApi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefundApi.Unmarshal(m, b)
}
func (m *RefundApi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefundApi.Marshal(b, m, deterministic)
}
func (m *RefundApi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefundApi.Merge(m, src)
}
func (m *RefundApi) XXX_Size() int {
	return xxx_messageInfo_RefundApi.Size(m)
}
func (m *RefundApi) XXX_DiscardUnknown() {
	xxx_messageInfo_RefundApi.DiscardUnknown(m)
}

var xxx_messageInfo_RefundApi proto.InternalMessageInfo

func (m *RefundApi) GetStm1() *Statement1 {
	if m != nil {
		return m.Stm1
	}
	return nil
}

type Fish struct {
	Statement1           []*Statement1 `protobuf:"bytes,1,rep,name=statement1,proto3" json:"statement1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Fish) Reset()         { *m = Fish{} }
func (m *Fish) String() string { return proto.CompactTextString(m) }
func (*Fish) ProtoMessage()    {}
func (*Fish) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{5}
}

func (m *Fish) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fish.Unmarshal(m, b)
}
func (m *Fish) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fish.Marshal(b, m, deterministic)
}
func (m *Fish) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fish.Merge(m, src)
}
func (m *Fish) XXX_Size() int {
	return xxx_messageInfo_Fish.Size(m)
}
func (m *Fish) XXX_DiscardUnknown() {
	xxx_messageInfo_Fish.DiscardUnknown(m)
}

var xxx_messageInfo_Fish proto.InternalMessageInfo

func (m *Fish) GetStatement1() []*Statement1 {
	if m != nil {
		return m.Statement1
	}
	return nil
}

type NestedFish struct {
	Fish                 []*Fish  `protobuf:"bytes,1,rep,name=fish,proto3" json:"fish,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NestedFish) Reset()         { *m = NestedFish{} }
func (m *NestedFish) String() string { return proto.CompactTextString(m) }
func (*NestedFish) ProtoMessage()    {}
func (*NestedFish) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{6}
}

func (m *NestedFish) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NestedFish.Unmarshal(m, b)
}
func (m *NestedFish) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NestedFish.Marshal(b, m, deterministic)
}
func (m *NestedFish) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NestedFish.Merge(m, src)
}
func (m *NestedFish) XXX_Size() int {
	return xxx_messageInfo_NestedFish.Size(m)
}
func (m *NestedFish) XXX_DiscardUnknown() {
	xxx_messageInfo_NestedFish.DiscardUnknown(m)
}

var xxx_messageInfo_NestedFish proto.InternalMessageInfo

func (m *NestedFish) GetFish() []*Fish {
	if m != nil {
		return m.Fish
	}
	return nil
}

type ListTransactionResponse struct {
	Data                 []*Transaction `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListTransactionResponse) Reset()         { *m = ListTransactionResponse{} }
func (m *ListTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*ListTransactionResponse) ProtoMessage()    {}
func (*ListTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{7}
}

func (m *ListTransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTransactionResponse.Unmarshal(m, b)
}
func (m *ListTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTransactionResponse.Marshal(b, m, deterministic)
}
func (m *ListTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTransactionResponse.Merge(m, src)
}
func (m *ListTransactionResponse) XXX_Size() int {
	return xxx_messageInfo_ListTransactionResponse.Size(m)
}
func (m *ListTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTransactionResponse proto.InternalMessageInfo

func (m *ListTransactionResponse) GetData() []*Transaction {
	if m != nil {
		return m.Data
	}
	return nil
}

type Department struct {
	// @inject_tag: es:"branch"
	Branch *Department_Object `protobuf:"bytes,1,opt,name=branch,proto3" json:"branch,omitempty"`
	// @inject_tag: es:"store"
	Store *Department_Object `protobuf:"bytes,2,opt,name=store,proto3" json:"store,omitempty"`
	// @inject_tag: es:"cashier"
	Cashier              *Department_Object `protobuf:"bytes,3,opt,name=cashier,proto3" json:"cashier,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Department) Reset()         { *m = Department{} }
func (m *Department) String() string { return proto.CompactTextString(m) }
func (*Department) ProtoMessage()    {}
func (*Department) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{8}
}

func (m *Department) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Department.Unmarshal(m, b)
}
func (m *Department) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Department.Marshal(b, m, deterministic)
}
func (m *Department) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Department.Merge(m, src)
}
func (m *Department) XXX_Size() int {
	return xxx_messageInfo_Department.Size(m)
}
func (m *Department) XXX_DiscardUnknown() {
	xxx_messageInfo_Department.DiscardUnknown(m)
}

var xxx_messageInfo_Department proto.InternalMessageInfo

func (m *Department) GetBranch() *Department_Object {
	if m != nil {
		return m.Branch
	}
	return nil
}

func (m *Department) GetStore() *Department_Object {
	if m != nil {
		return m.Store
	}
	return nil
}

func (m *Department) GetCashier() *Department_Object {
	if m != nil {
		return m.Cashier
	}
	return nil
}

type Department_Object struct {
	// @inject_tag: es:"id"
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Department_Object) Reset()         { *m = Department_Object{} }
func (m *Department_Object) String() string { return proto.CompactTextString(m) }
func (*Department_Object) ProtoMessage()    {}
func (*Department_Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{8, 0}
}

func (m *Department_Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Department_Object.Unmarshal(m, b)
}
func (m *Department_Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Department_Object.Marshal(b, m, deterministic)
}
func (m *Department_Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Department_Object.Merge(m, src)
}
func (m *Department_Object) XXX_Size() int {
	return xxx_messageInfo_Department_Object.Size(m)
}
func (m *Department_Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Department_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Department_Object proto.InternalMessageInfo

func (m *Department_Object) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Department_Object) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// `DataTransLog` represents data will be reported as table on client in all trasactions.
type Transaction struct {
	// // @inject_tag: es:"reqDate"
	// google.protobuf.Timestamp create_time = 1;
	// @inject_tag: es:"appTransId" query:"*%*"
	AppTransId           string   `protobuf:"bytes,2,opt,name=app_trans_id,json=appTransId,proto3" json:"app_trans_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{9}
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

func (m *Transaction) GetAppTransId() string {
	if m != nil {
		return m.AppTransId
	}
	return ""
}

func init() {
	proto.RegisterEnum("core_service.RefundTransLogApi", RefundTransLogApi_name, RefundTransLogApi_value)
	proto.RegisterEnum("core_service.ZaloPayType", ZaloPayType_name, ZaloPayType_value)
	proto.RegisterEnum("core_service.TransStatus", TransStatus_name, TransStatus_value)
	proto.RegisterEnum("core_service.PaymentMethod", PaymentMethod_name, PaymentMethod_value)
	proto.RegisterEnum("core_service.BODetailReconciliation_VoucherCode", BODetailReconciliation_VoucherCode_name, BODetailReconciliation_VoucherCode_value)
	proto.RegisterType((*OutnerField)(nil), "core_service.OutnerField")
	proto.RegisterType((*BODetailReconciliation)(nil), "core_service.BODetailReconciliation")
	proto.RegisterType((*BODetailReconciliation_TimeAttribute)(nil), "core_service.BODetailReconciliation.TimeAttribute")
	proto.RegisterType((*BODetailReconciliation_CountableAttribute)(nil), "core_service.BODetailReconciliation.CountableAttribute")
	proto.RegisterType((*BODetailReconciliation_CountableAttribute_Transaction)(nil), "core_service.BODetailReconciliation.CountableAttribute.Transaction")
	proto.RegisterType((*Statement)(nil), "core_service.Statement")
	proto.RegisterType((*Statement1)(nil), "core_service.Statement1")
	proto.RegisterType((*RefundApi)(nil), "core_service.RefundApi")
	proto.RegisterType((*Fish)(nil), "core_service.Fish")
	proto.RegisterType((*NestedFish)(nil), "core_service.NestedFish")
	proto.RegisterType((*ListTransactionResponse)(nil), "core_service.ListTransactionResponse")
	proto.RegisterType((*Department)(nil), "core_service.Department")
	proto.RegisterType((*Department_Object)(nil), "core_service.Department.Object")
	proto.RegisterType((*Transaction)(nil), "core_service.Transaction")
}

func init() { proto.RegisterFile("core.proto", fileDescriptor_f7e43720d1edc0fe) }

var fileDescriptor_f7e43720d1edc0fe = []byte{
	// 985 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xed, 0x6e, 0xe3, 0x44,
	0x14, 0xc5, 0x69, 0xda, 0x25, 0xd7, 0x49, 0xea, 0x9d, 0x22, 0x9a, 0x8d, 0xb4, 0xda, 0x62, 0xad,
	0x56, 0xa5, 0x94, 0x94, 0x86, 0x8f, 0xa5, 0x12, 0x3f, 0x70, 0x1d, 0x67, 0x6b, 0x36, 0x4d, 0x8c,
	0xed, 0xb4, 0x74, 0x25, 0x34, 0x9a, 0xd8, 0xd3, 0xda, 0x34, 0xb1, 0x2d, 0x7b, 0x52, 0x29, 0xfb,
	0x98, 0xbc, 0x03, 0x3f, 0x79, 0x07, 0x34, 0x63, 0xe7, 0xab, 0x05, 0xba, 0xf0, 0x6f, 0xee, 0xdc,
	0x73, 0xcf, 0x9c, 0x3b, 0xf7, 0x78, 0x0c, 0xe0, 0xc5, 0x29, 0x6d, 0x25, 0x69, 0xcc, 0x62, 0x54,
	0xe5, 0x6b, 0x9c, 0xd1, 0xf4, 0x2e, 0xf4, 0x68, 0xf3, 0xc5, 0x4d, 0x1c, 0xdf, 0x8c, 0xe9, 0x91,
	0xc8, 0x8d, 0xa6, 0xd7, 0x47, 0x2c, 0x9c, 0xd0, 0x8c, 0x91, 0x49, 0x92, 0xc3, 0xd5, 0xcf, 0x40,
	0x1e, 0x4c, 0x59, 0x44, 0xd3, 0x6e, 0x48, 0xc7, 0x3e, 0x42, 0x50, 0x8e, 0xc8, 0x84, 0x36, 0xa4,
	0x3d, 0x69, 0xbf, 0x62, 0x8b, 0xb5, 0xfa, 0xfb, 0x13, 0xf8, 0xf4, 0x74, 0xd0, 0xa1, 0x8c, 0x84,
	0x63, 0x9b, 0x7a, 0x71, 0xe4, 0x85, 0xe3, 0x90, 0xb0, 0x30, 0x8e, 0xd0, 0x0f, 0x50, 0x8d, 0x45,
	0x35, 0xbe, 0x0e, 0xc7, 0xd4, 0x17, 0x65, 0x72, 0xfb, 0x59, 0x6b, 0x55, 0x43, 0x6b, 0x85, 0xdf,
	0x96, 0xe3, 0x22, 0x18, 0x53, 0x1f, 0x9d, 0x00, 0xb0, 0x94, 0x44, 0x19, 0xe6, 0xa2, 0x1a, 0x25,
	0x51, 0xdb, 0x6c, 0xe5, 0x8a, 0x5b, 0x73, 0xc5, 0x2d, 0x77, 0xae, 0xd8, 0xae, 0x08, 0x34, 0x8f,
	0xd1, 0x4b, 0xa8, 0xbf, 0x4f, 0x70, 0x36, 0xcb, 0x18, 0x9d, 0x60, 0xa1, 0x78, 0x43, 0x28, 0xae,
	0xbe, 0x4f, 0x1c, 0xb1, 0xd9, 0x27, 0x13, 0x8a, 0x1c, 0xa8, 0xde, 0xc5, 0x53, 0x2f, 0xa0, 0x29,
	0xf6, 0x62, 0x9f, 0x36, 0x60, 0x4f, 0xda, 0xaf, 0xb7, 0xbf, 0x5a, 0x97, 0xf7, 0xf7, 0xad, 0xb5,
	0x2e, 0xf2, 0x42, 0x3d, 0xf6, 0xa9, 0x2d, 0xdf, 0x2d, 0x03, 0x74, 0x05, 0x75, 0xae, 0x17, 0x13,
	0xc6, 0xd2, 0x70, 0x34, 0x65, 0xb4, 0x51, 0x13, 0xca, 0xdb, 0x1f, 0x44, 0xcb, 0xd5, 0x6b, 0xf3,
	0x4a, 0xbb, 0xc6, 0x56, 0x43, 0x14, 0xc0, 0x8e, 0x17, 0x4f, 0x23, 0x46, 0x46, 0xe3, 0x55, 0xfe,
	0xba, 0xe0, 0x7f, 0xfd, 0x41, 0xfc, 0xfa, 0xbc, 0x7e, 0x79, 0x08, 0xf2, 0x1e, 0xec, 0x35, 0x7f,
	0x82, 0xda, 0x9a, 0x12, 0x3e, 0x8b, 0x94, 0xb2, 0x74, 0x96, 0xcf, 0x42, 0x7a, 0x7c, 0x16, 0x02,
	0xcd, 0xe3, 0xe6, 0x9f, 0x25, 0x40, 0x0f, 0x8f, 0x45, 0xdf, 0x71, 0xc6, 0xeb, 0x69, 0xe4, 0x63,
	0x92, 0x84, 0x0d, 0x24, 0x18, 0x77, 0xd7, 0x7b, 0xb0, 0x45, 0x5e, 0x4b, 0x42, 0x4e, 0x57, 0x2c,
	0xd1, 0x25, 0x94, 0x7d, 0xc2, 0x48, 0xe3, 0x13, 0x51, 0xa1, 0xff, 0xcf, 0xae, 0x5b, 0x2e, 0xf7,
	0x0a, 0xf1, 0x78, 0xda, 0x16, 0x84, 0x48, 0x85, 0x1a, 0x4b, 0x28, 0x1e, 0x91, 0xe8, 0x36, 0xb7,
	0xc3, 0x4b, 0x61, 0x19, 0x99, 0x25, 0xf4, 0x94, 0x44, 0xb7, 0x62, 0xb8, 0x9f, 0xc3, 0x46, 0xc6,
	0x26, 0xc2, 0x4c, 0x0f, 0xd4, 0x3a, 0x8c, 0x30, 0x3a, 0xa1, 0x11, 0xb3, 0x39, 0x06, 0x3d, 0x07,
	0x08, 0xb9, 0xfb, 0xc4, 0xed, 0x36, 0x5e, 0xed, 0x49, 0xfb, 0x9b, 0x76, 0x85, 0xef, 0x08, 0x31,
	0xcd, 0x5f, 0x40, 0x5e, 0x91, 0x80, 0xea, 0x50, 0xba, 0x3d, 0x2e, 0x3e, 0xab, 0xd2, 0xed, 0xb1,
	0x88, 0xdb, 0xc2, 0xf3, 0x3c, 0x6e, 0xff, 0x87, 0x83, 0x55, 0x1d, 0xe4, 0x15, 0x73, 0xa2, 0xe7,
	0xf0, 0xec, 0x62, 0x30, 0xd4, 0xcf, 0x0c, 0x1b, 0xeb, 0x83, 0x8e, 0x81, 0x87, 0x7d, 0xc7, 0x32,
	0x74, 0xb3, 0x6b, 0xea, 0x46, 0x47, 0xf9, 0x08, 0x01, 0x6c, 0x5d, 0xe8, 0xf8, 0xca, 0x70, 0x14,
	0x09, 0x55, 0x60, 0xf3, 0x42, 0xc7, 0xfd, 0x81, 0x52, 0x52, 0xbf, 0x80, 0xca, 0x82, 0x96, 0x8b,
	0x09, 0x16, 0xe2, 0x02, 0x21, 0x2e, 0x58, 0x88, 0x0b, 0xda, 0xea, 0x21, 0xc0, 0x02, 0x7c, 0xfc,
	0x28, 0xfa, 0x04, 0x2a, 0x8b, 0xc1, 0xa2, 0x43, 0x28, 0x67, 0x6c, 0x72, 0x5c, 0x34, 0xd6, 0xf8,
	0x87, 0xc6, 0x8e, 0x6d, 0x81, 0x52, 0x7f, 0x84, 0x72, 0x37, 0xcc, 0x02, 0xf4, 0x3d, 0x40, 0xb6,
	0xc8, 0x35, 0xa4, 0xbd, 0x8d, 0x7f, 0xad, 0x5d, 0xc1, 0xaa, 0xdf, 0x00, 0xf4, 0x69, 0xc6, 0xa8,
	0x2f, 0x78, 0x5e, 0x41, 0xf9, 0x3a, 0xcc, 0x82, 0x82, 0x01, 0xad, 0x33, 0x70, 0x84, 0x2d, 0xf2,
	0xea, 0x19, 0xec, 0xf6, 0xc2, 0x8c, 0xad, 0x7a, 0x86, 0x66, 0x49, 0x1c, 0x65, 0x14, 0x7d, 0x59,
	0xd8, 0x31, 0xa7, 0xb8, 0xf7, 0xb4, 0x3d, 0x30, 0x99, 0xfa, 0x87, 0x04, 0xd0, 0xa1, 0x09, 0x49,
	0x99, 0xb8, 0xd9, 0xd7, 0xb0, 0x35, 0x4a, 0x49, 0xe4, 0x05, 0xc5, 0x27, 0xf5, 0x62, 0xbd, 0x7e,
	0x89, 0x6c, 0x0d, 0x46, 0xbf, 0x51, 0x8f, 0xd9, 0x05, 0x1c, 0x7d, 0x0b, 0x9b, 0x19, 0x8b, 0xd3,
	0xf9, 0xb3, 0xf8, 0x68, 0x5d, 0x8e, 0x46, 0x27, 0xf0, 0xc4, 0x23, 0x59, 0x10, 0xd2, 0xb4, 0xb8,
	0xf1, 0x47, 0x0b, 0xe7, 0xf8, 0xe6, 0x21, 0x6c, 0xe5, 0x5b, 0x7c, 0xa0, 0xa1, 0x3f, 0x1f, 0x70,
	0xb8, 0xfc, 0x29, 0x94, 0x56, 0x7e, 0x0a, 0x47, 0xeb, 0xf6, 0xde, 0x83, 0x2a, 0x49, 0x12, 0x9c,
	0x3f, 0xe7, 0xa1, 0x5f, 0x40, 0x81, 0x24, 0x89, 0x40, 0x99, 0xfe, 0xc1, 0x1b, 0x78, 0x9a, 0xbb,
	0x42, 0x6c, 0xf4, 0xe2, 0x1b, 0xee, 0x8e, 0x1d, 0xd8, 0xb6, 0x5d, 0x6d, 0x69, 0x59, 0xe1, 0xd8,
	0x2a, 0x7c, 0xcc, 0x37, 0xbb, 0xc3, 0x5e, 0x4f, 0x91, 0xd0, 0x36, 0xc8, 0x3c, 0xb2, 0x34, 0xdb,
	0x35, 0xb5, 0x9e, 0x52, 0x3a, 0xf8, 0x15, 0xe4, 0x77, 0x64, 0x1c, 0x5b, 0x64, 0xe6, 0xce, 0x12,
	0x8a, 0x76, 0x61, 0xe7, 0x9d, 0x85, 0xdd, 0x2b, 0xcb, 0xb8, 0x47, 0xb3, 0x0d, 0xf2, 0x3c, 0xa1,
	0x59, 0x96, 0x22, 0xf1, 0xc3, 0xe6, 0x1b, 0x6f, 0x34, 0xd7, 0xb8, 0xd4, 0xae, 0x94, 0xd2, 0x1a,
	0xaa, 0xd7, 0x53, 0x36, 0x0e, 0xc6, 0x45, 0x63, 0xdc, 0x5f, 0xd3, 0x0c, 0x21, 0xa8, 0xbb, 0xce,
	0x3d, 0x66, 0x19, 0x9e, 0xb8, 0x0e, 0xee, 0x6a, 0x26, 0xd7, 0xf7, 0x14, 0x6a, 0xae, 0x83, 0x9d,
	0xa1, 0xae, 0x1b, 0x8e, 0xd3, 0x1d, 0xf6, 0x94, 0x12, 0xff, 0xe4, 0x5c, 0x27, 0xa7, 0x2b, 0xd6,
	0x7d, 0xe3, 0x52, 0x29, 0x17, 0xd0, 0x61, 0xbf, 0x6b, 0xf6, 0x4d, 0xe7, 0xcc, 0xe8, 0x28, 0x9b,
	0x07, 0x0c, 0x6a, 0x16, 0x99, 0xf1, 0x79, 0x9c, 0x53, 0x16, 0xc4, 0xfc, 0xae, 0xeb, 0xd6, 0xf9,
	0xbd, 0xf3, 0x6a, 0x50, 0xb1, 0xce, 0xf1, 0xcf, 0x36, 0xb6, 0xb4, 0x2b, 0x45, 0x2a, 0xc2, 0x8e,
	0x69, 0x1b, 0xba, 0x9b, 0x77, 0xc0, 0xb3, 0x43, 0x53, 0x7f, 0xcb, 0xf3, 0x1b, 0x48, 0x81, 0xaa,
	0x75, 0x8e, 0xb5, 0xa1, 0x3b, 0xe8, 0x18, 0xa7, 0xa6, 0xab, 0x94, 0x51, 0x1d, 0x40, 0x90, 0xbe,
	0xed, 0x0f, 0x2e, 0xfb, 0xca, 0xe6, 0x68, 0x4b, 0x3c, 0xe8, 0x5f, 0xff, 0x15, 0x00, 0x00, 0xff,
	0xff, 0x82, 0x7e, 0xd0, 0xec, 0x38, 0x08, 0x00, 0x00,
}
