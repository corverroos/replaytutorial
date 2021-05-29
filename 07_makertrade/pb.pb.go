// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb.proto

package main

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Order struct {
	Market               string   `protobuf:"bytes,1,opt,name=market,proto3" json:"market,omitempty"`
	Amount               float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	IsBuy                bool     `protobuf:"varint,3,opt,name=is_buy,json=isBuy,proto3" json:"is_buy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{1}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

func (m *Order) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Order) GetIsBuy() bool {
	if m != nil {
		return m.IsBuy
	}
	return false
}

type OrderRef struct {
	Market               string   `protobuf:"bytes,1,opt,name=market,proto3" json:"market,omitempty"`
	IsBuy                bool     `protobuf:"varint,2,opt,name=is_buy,json=isBuy,proto3" json:"is_buy,omitempty"`
	Price                float64  `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	ExtId                string   `protobuf:"bytes,4,opt,name=ext_id,json=extId,proto3" json:"ext_id,omitempty"`
	TimeUnixMilli        int64    `protobuf:"varint,5,opt,name=time_unix_milli,json=timeUnixMilli,proto3" json:"time_unix_milli,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderRef) Reset()         { *m = OrderRef{} }
func (m *OrderRef) String() string { return proto.CompactTextString(m) }
func (*OrderRef) ProtoMessage()    {}
func (*OrderRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{2}
}

func (m *OrderRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRef.Unmarshal(m, b)
}
func (m *OrderRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRef.Marshal(b, m, deterministic)
}
func (m *OrderRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRef.Merge(m, src)
}
func (m *OrderRef) XXX_Size() int {
	return xxx_messageInfo_OrderRef.Size(m)
}
func (m *OrderRef) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRef.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRef proto.InternalMessageInfo

func (m *OrderRef) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

func (m *OrderRef) GetIsBuy() bool {
	if m != nil {
		return m.IsBuy
	}
	return false
}

func (m *OrderRef) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderRef) GetExtId() string {
	if m != nil {
		return m.ExtId
	}
	return ""
}

func (m *OrderRef) GetTimeUnixMilli() int64 {
	if m != nil {
		return m.TimeUnixMilli
	}
	return 0
}

type OrderState struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	CurrentPrice         float64  `protobuf:"fixed64,2,opt,name=current_price,json=currentPrice,proto3" json:"current_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderState) Reset()         { *m = OrderState{} }
func (m *OrderState) String() string { return proto.CompactTextString(m) }
func (*OrderState) ProtoMessage()    {}
func (*OrderState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{3}
}

func (m *OrderState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderState.Unmarshal(m, b)
}
func (m *OrderState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderState.Marshal(b, m, deterministic)
}
func (m *OrderState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderState.Merge(m, src)
}
func (m *OrderState) XXX_Size() int {
	return xxx_messageInfo_OrderState.Size(m)
}
func (m *OrderState) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderState.DiscardUnknown(m)
}

var xxx_messageInfo_OrderState proto.InternalMessageInfo

func (m *OrderState) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *OrderState) GetCurrentPrice() float64 {
	if m != nil {
		return m.CurrentPrice
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "main.Empty")
	proto.RegisterType((*Order)(nil), "main.Order")
	proto.RegisterType((*OrderRef)(nil), "main.OrderRef")
	proto.RegisterType((*OrderState)(nil), "main.OrderState")
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor_f80abaa17e25ccc8) }

var fileDescriptor_f80abaa17e25ccc8 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0xc9, 0xcc, 0xa4, 0xd6, 0x8b, 0x83, 0x10, 0x54, 0xb2, 0x2c, 0x15, 0xa4, 0x2b, 0x37,
	0xbe, 0x81, 0xe0, 0x62, 0x16, 0xfe, 0x10, 0x71, 0x1d, 0xd2, 0x69, 0x84, 0x8b, 0x93, 0xb4, 0xa4,
	0x37, 0xd0, 0xbe, 0x84, 0xcf, 0x2c, 0x49, 0x0b, 0xba, 0x71, 0x79, 0x3e, 0x72, 0x0e, 0x5f, 0x2e,
	0x94, 0x43, 0x7b, 0x3f, 0x84, 0x9e, 0x7a, 0xb1, 0x73, 0x06, 0x7d, 0x7d, 0x06, 0xfc, 0xc9, 0x0d,
	0x34, 0xd7, 0x2f, 0xc0, 0x5f, 0x43, 0x67, 0x83, 0xb8, 0x81, 0xc2, 0x99, 0xf0, 0x65, 0x49, 0xb2,
	0x8a, 0x35, 0xe7, 0x6a, 0x4d, 0x89, 0x1b, 0xd7, 0x47, 0x4f, 0x72, 0x53, 0xb1, 0x86, 0xa9, 0x35,
	0x89, 0x6b, 0x28, 0x70, 0xd4, 0x6d, 0x9c, 0xe5, 0xb6, 0x62, 0x4d, 0xa9, 0x38, 0x8e, 0x8f, 0x71,
	0xae, 0xbf, 0x19, 0x94, 0x79, 0x50, 0xd9, 0xcf, 0x7f, 0x37, 0x7f, 0xbb, 0x9b, 0x3f, 0x5d, 0x71,
	0x05, 0x7c, 0x08, 0x78, 0xb4, 0x79, 0x91, 0xa9, 0x25, 0xa4, 0xc7, 0x76, 0x22, 0x8d, 0x9d, 0xdc,
	0xe5, 0x11, 0x6e, 0x27, 0x3a, 0x74, 0xe2, 0x0e, 0x2e, 0x09, 0x9d, 0xd5, 0xd1, 0xe3, 0xa4, 0x1d,
	0x9e, 0x4e, 0x28, 0x79, 0xc5, 0x9a, 0xad, 0xda, 0x27, 0xfc, 0xe1, 0x71, 0x7a, 0x4e, 0xb0, 0x3e,
	0x00, 0x64, 0x9f, 0x77, 0x32, 0x64, 0x93, 0xd1, 0x48, 0x86, 0xe2, 0x98, 0x8d, 0xb8, 0x5a, 0x93,
	0xb8, 0x85, 0xfd, 0x31, 0x86, 0x60, 0x3d, 0xe9, 0x45, 0x61, 0xf9, 0xec, 0xc5, 0x0a, 0xdf, 0x12,
	0x6b, 0x8b, 0x7c, 0xc1, 0x87, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0x4e, 0x6b, 0x98, 0x4d,
	0x01, 0x00, 0x00,
}
