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

type Data struct {
	Market               string    `protobuf:"bytes,1,opt,name=market,proto3" json:"market,omitempty"`
	Size                 int64     `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	PeriodSec            int64     `protobuf:"varint,3,opt,name=period_sec,json=periodSec,proto3" json:"period_sec,omitempty"`
	Values               []float64 `protobuf:"fixed64,4,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{1}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

func (m *Data) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Data) GetPeriodSec() int64 {
	if m != nil {
		return m.PeriodSec
	}
	return 0
}

func (m *Data) GetValues() []float64 {
	if m != nil {
		return m.Values
	}
	return nil
}

type Market struct {
	Market               string   `protobuf:"bytes,1,opt,name=market,proto3" json:"market,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Market) Reset()         { *m = Market{} }
func (m *Market) String() string { return proto.CompactTextString(m) }
func (*Market) ProtoMessage()    {}
func (*Market) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{2}
}

func (m *Market) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Market.Unmarshal(m, b)
}
func (m *Market) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Market.Marshal(b, m, deterministic)
}
func (m *Market) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Market.Merge(m, src)
}
func (m *Market) XXX_Size() int {
	return xxx_messageInfo_Market.Size(m)
}
func (m *Market) XXX_DiscardUnknown() {
	xxx_messageInfo_Market.DiscardUnknown(m)
}

var xxx_messageInfo_Market proto.InternalMessageInfo

func (m *Market) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

type Double struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Double) Reset()         { *m = Double{} }
func (m *Double) String() string { return proto.CompactTextString(m) }
func (*Double) ProtoMessage()    {}
func (*Double) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{3}
}

func (m *Double) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Double.Unmarshal(m, b)
}
func (m *Double) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Double.Marshal(b, m, deterministic)
}
func (m *Double) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Double.Merge(m, src)
}
func (m *Double) XXX_Size() int {
	return xxx_messageInfo_Double.Size(m)
}
func (m *Double) XXX_DiscardUnknown() {
	xxx_messageInfo_Double.DiscardUnknown(m)
}

var xxx_messageInfo_Double proto.InternalMessageInfo

func (m *Double) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "main.Empty")
	proto.RegisterType((*Data)(nil), "main.Data")
	proto.RegisterType((*Market)(nil), "main.Market")
	proto.RegisterType((*Double)(nil), "main.Double")
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor_f80abaa17e25ccc8) }

var fileDescriptor_f80abaa17e25ccc8 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xbf, 0x0e, 0xc2, 0x20,
	0x10, 0x87, 0x83, 0xa5, 0x68, 0x6f, 0xbc, 0x18, 0xc3, 0xa2, 0x21, 0x4c, 0x9d, 0x5c, 0x7c, 0x85,
	0x3a, 0xba, 0xe0, 0x03, 0x18, 0x5a, 0x6f, 0x20, 0xb6, 0x42, 0x5a, 0x6a, 0xa2, 0x4f, 0x6f, 0x84,
	0xae, 0x6e, 0xf7, 0xdd, 0x9f, 0xef, 0xf2, 0x83, 0x4d, 0x68, 0x8f, 0x61, 0xf4, 0xd1, 0x23, 0x1f,
	0xac, 0x7b, 0xea, 0x35, 0x94, 0xe7, 0x21, 0xc4, 0xb7, 0x76, 0xc0, 0x1b, 0x1b, 0x2d, 0xee, 0x40,
	0x0c, 0x76, 0x7c, 0x50, 0x94, 0x4c, 0xb1, 0xba, 0x32, 0x0b, 0x21, 0x02, 0x9f, 0xdc, 0x87, 0xe4,
	0x4a, 0xb1, 0xba, 0x30, 0xa9, 0xc6, 0x3d, 0x40, 0xa0, 0xd1, 0xf9, 0xfb, 0x6d, 0xa2, 0x4e, 0x16,
	0x69, 0x52, 0xe5, 0xce, 0x95, 0xba, 0x9f, 0xea, 0x65, 0xfb, 0x99, 0x26, 0xc9, 0x55, 0x51, 0x33,
	0xb3, 0x90, 0x56, 0x20, 0x2e, 0x59, 0xfa, 0xe7, 0x99, 0x3e, 0x80, 0x68, 0xfc, 0xdc, 0xf6, 0x84,
	0x5b, 0x28, 0xd3, 0x55, 0x5a, 0x60, 0x26, 0x43, 0x2b, 0x52, 0x84, 0xd3, 0x37, 0x00, 0x00, 0xff,
	0xff, 0xfb, 0x0f, 0xb8, 0x4e, 0xce, 0x00, 0x00, 0x00,
}
