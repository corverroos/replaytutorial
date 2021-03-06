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

type JobStatus int32

const (
	JobStatus_UNKNOWN   JobStatus = 0
	JobStatus_PENDING   JobStatus = 1
	JobStatus_CLAIMED   JobStatus = 2
	JobStatus_COMPLETED JobStatus = 3
	JobStatus_FAILED    JobStatus = 4
)

var JobStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "PENDING",
	2: "CLAIMED",
	3: "COMPLETED",
	4: "FAILED",
}

var JobStatus_value = map[string]int32{
	"UNKNOWN":   0,
	"PENDING":   1,
	"CLAIMED":   2,
	"COMPLETED": 3,
	"FAILED":    4,
}

func (x JobStatus) String() string {
	return proto.EnumName(JobStatus_name, int32(x))
}

func (JobStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{0}
}

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

type JobRequest struct {
	Status     JobStatus `protobuf:"varint,1,opt,name=status,proto3,enum=main.JobStatus" json:"status,omitempty"`
	ClaimedBy  string    `protobuf:"bytes,2,opt,name=claimed_by,json=claimedBy,proto3" json:"claimed_by,omitempty"`
	FailReason string    `protobuf:"bytes,3,opt,name=fail_reason,json=failReason,proto3" json:"fail_reason,omitempty"`
	// Input data
	IsFoo                bool     `protobuf:"varint,4,opt,name=is_foo,json=isFoo,proto3" json:"is_foo,omitempty"`
	IsBar                bool     `protobuf:"varint,5,opt,name=is_bar,json=isBar,proto3" json:"is_bar,omitempty"`
	Value                string   `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobRequest) Reset()         { *m = JobRequest{} }
func (m *JobRequest) String() string { return proto.CompactTextString(m) }
func (*JobRequest) ProtoMessage()    {}
func (*JobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{1}
}

func (m *JobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobRequest.Unmarshal(m, b)
}
func (m *JobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobRequest.Marshal(b, m, deterministic)
}
func (m *JobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobRequest.Merge(m, src)
}
func (m *JobRequest) XXX_Size() int {
	return xxx_messageInfo_JobRequest.Size(m)
}
func (m *JobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JobRequest proto.InternalMessageInfo

func (m *JobRequest) GetStatus() JobStatus {
	if m != nil {
		return m.Status
	}
	return JobStatus_UNKNOWN
}

func (m *JobRequest) GetClaimedBy() string {
	if m != nil {
		return m.ClaimedBy
	}
	return ""
}

func (m *JobRequest) GetFailReason() string {
	if m != nil {
		return m.FailReason
	}
	return ""
}

func (m *JobRequest) GetIsFoo() bool {
	if m != nil {
		return m.IsFoo
	}
	return false
}

func (m *JobRequest) GetIsBar() bool {
	if m != nil {
		return m.IsBar
	}
	return false
}

func (m *JobRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ClaimJob struct {
	By                   string   `protobuf:"bytes,1,opt,name=by,proto3" json:"by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClaimJob) Reset()         { *m = ClaimJob{} }
func (m *ClaimJob) String() string { return proto.CompactTextString(m) }
func (*ClaimJob) ProtoMessage()    {}
func (*ClaimJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{2}
}

func (m *ClaimJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClaimJob.Unmarshal(m, b)
}
func (m *ClaimJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClaimJob.Marshal(b, m, deterministic)
}
func (m *ClaimJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimJob.Merge(m, src)
}
func (m *ClaimJob) XXX_Size() int {
	return xxx_messageInfo_ClaimJob.Size(m)
}
func (m *ClaimJob) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimJob.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimJob proto.InternalMessageInfo

func (m *ClaimJob) GetBy() string {
	if m != nil {
		return m.By
	}
	return ""
}

type CompleteJob struct {
	Fail                 bool     `protobuf:"varint,1,opt,name=fail,proto3" json:"fail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompleteJob) Reset()         { *m = CompleteJob{} }
func (m *CompleteJob) String() string { return proto.CompactTextString(m) }
func (*CompleteJob) ProtoMessage()    {}
func (*CompleteJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{3}
}

func (m *CompleteJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteJob.Unmarshal(m, b)
}
func (m *CompleteJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteJob.Marshal(b, m, deterministic)
}
func (m *CompleteJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteJob.Merge(m, src)
}
func (m *CompleteJob) XXX_Size() int {
	return xxx_messageInfo_CompleteJob.Size(m)
}
func (m *CompleteJob) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteJob.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteJob proto.InternalMessageInfo

func (m *CompleteJob) GetFail() bool {
	if m != nil {
		return m.Fail
	}
	return false
}

func init() {
	proto.RegisterEnum("main.JobStatus", JobStatus_name, JobStatus_value)
	proto.RegisterType((*Empty)(nil), "main.Empty")
	proto.RegisterType((*JobRequest)(nil), "main.JobRequest")
	proto.RegisterType((*ClaimJob)(nil), "main.ClaimJob")
	proto.RegisterType((*CompleteJob)(nil), "main.CompleteJob")
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor_f80abaa17e25ccc8) }

var fileDescriptor_f80abaa17e25ccc8 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0x6d, 0xd7, 0x76, 0xed, 0x1d, 0xce, 0x12, 0x14, 0x82, 0x20, 0xce, 0xbe, 0x38, 0x7c,
	0xd8, 0x83, 0xfe, 0x82, 0xad, 0xcd, 0x64, 0x73, 0xeb, 0x46, 0x54, 0x7c, 0x1c, 0x89, 0x66, 0x50,
	0x68, 0x97, 0xda, 0x64, 0x42, 0x7f, 0x9a, 0xff, 0x4e, 0x12, 0xab, 0x6f, 0x39, 0xdf, 0x39, 0xe4,
	0xdc, 0x7b, 0x21, 0xac, 0xf9, 0xa4, 0x6e, 0xa4, 0x96, 0xc8, 0xab, 0x58, 0x71, 0x48, 0xfa, 0xe0,
	0x93, 0xaa, 0xd6, 0x6d, 0xf2, 0xed, 0x00, 0x2c, 0x25, 0xa7, 0xe2, 0xf3, 0x28, 0x94, 0x46, 0xb7,
	0x10, 0x28, 0xcd, 0xf4, 0x51, 0x61, 0x67, 0xe4, 0x8c, 0x87, 0xf7, 0x67, 0x13, 0x13, 0x9f, 0x2c,
	0x25, 0x7f, 0xb6, 0x98, 0x76, 0x36, 0xba, 0x02, 0x78, 0x2f, 0x59, 0x51, 0x89, 0x8f, 0x1d, 0x6f,
	0xb1, 0x3b, 0x72, 0xc6, 0x11, 0x8d, 0x3a, 0x32, 0x6b, 0xd1, 0x35, 0x0c, 0xf6, 0xac, 0x28, 0x77,
	0x8d, 0x60, 0x4a, 0x1e, 0x70, 0xcf, 0xfa, 0x60, 0x10, 0xb5, 0x04, 0x5d, 0x40, 0x50, 0xa8, 0xdd,
	0x5e, 0x4a, 0xec, 0x8d, 0x9c, 0x71, 0x48, 0xfd, 0x42, 0xcd, 0xa5, 0xec, 0x30, 0x67, 0x0d, 0xf6,
	0xff, 0xf0, 0x8c, 0x35, 0xe8, 0x1c, 0xfc, 0x2f, 0x56, 0x1e, 0x05, 0x0e, 0xec, 0x47, 0xbf, 0x22,
	0xb9, 0x84, 0x30, 0x35, 0x8d, 0x4b, 0xc9, 0xd1, 0x10, 0x5c, 0xde, 0xda, 0xa1, 0x23, 0xea, 0xf2,
	0x36, 0xb9, 0x81, 0x41, 0x2a, 0xab, 0xba, 0x14, 0x5a, 0x18, 0x1b, 0x81, 0x67, 0xca, 0x6d, 0x20,
	0xa4, 0xf6, 0x7d, 0xb7, 0x86, 0xe8, 0x7f, 0x2f, 0x34, 0x80, 0xfe, 0x6b, 0xfe, 0x94, 0x6f, 0xde,
	0xf2, 0xf8, 0xc4, 0x88, 0x2d, 0xc9, 0xb3, 0x45, 0xfe, 0x18, 0x3b, 0x46, 0xa4, 0xab, 0xe9, 0x62,
	0x4d, 0xb2, 0xd8, 0x45, 0xa7, 0x10, 0xa5, 0x9b, 0xf5, 0x76, 0x45, 0x5e, 0x48, 0x16, 0xf7, 0x10,
	0x40, 0x30, 0x9f, 0x2e, 0x56, 0x24, 0x8b, 0x3d, 0x1e, 0xd8, 0xfb, 0x3e, 0xfc, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xf9, 0x3f, 0x53, 0xa9, 0x6b, 0x01, 0x00, 0x00,
}
