// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language-agent/JVMMetric.proto

package v3

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	v3 "github.com/SkyAPM/SkyAPM-php-sdk/reporter/network/common/v3"
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

type PoolType int32

const (
	PoolType_CODE_CACHE_USAGE PoolType = 0
	PoolType_NEWGEN_USAGE     PoolType = 1
	PoolType_OLDGEN_USAGE     PoolType = 2
	PoolType_SURVIVOR_USAGE   PoolType = 3
	PoolType_PERMGEN_USAGE    PoolType = 4
	PoolType_METASPACE_USAGE  PoolType = 5
)

var PoolType_name = map[int32]string{
	0: "CODE_CACHE_USAGE",
	1: "NEWGEN_USAGE",
	2: "OLDGEN_USAGE",
	3: "SURVIVOR_USAGE",
	4: "PERMGEN_USAGE",
	5: "METASPACE_USAGE",
}

var PoolType_value = map[string]int32{
	"CODE_CACHE_USAGE": 0,
	"NEWGEN_USAGE":     1,
	"OLDGEN_USAGE":     2,
	"SURVIVOR_USAGE":   3,
	"PERMGEN_USAGE":    4,
	"METASPACE_USAGE":  5,
}

func (x PoolType) String() string {
	return proto.EnumName(PoolType_name, int32(x))
}

func (PoolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0cf5d9e7fa9a007c, []int{0}
}

type GCPhrase int32

const (
	GCPhrase_NEW GCPhrase = 0
	GCPhrase_OLD GCPhrase = 1
)

var GCPhrase_name = map[int32]string{
	0: "NEW",
	1: "OLD",
}

var GCPhrase_value = map[string]int32{
	"NEW": 0,
	"OLD": 1,
}

func (x GCPhrase) String() string {
	return proto.EnumName(GCPhrase_name, int32(x))
}

func (GCPhrase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0cf5d9e7fa9a007c, []int{1}
}

type JVMMetricCollection struct {
	Metrics              []*JVMMetric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	Service              string       `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	ServiceInstance      string       `protobuf:"bytes,3,opt,name=serviceInstance,proto3" json:"serviceInstance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *JVMMetricCollection) Reset()         { *m = JVMMetricCollection{} }
func (m *JVMMetricCollection) String() string { return proto.CompactTextString(m) }
func (*JVMMetricCollection) ProtoMessage()    {}
func (*JVMMetricCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cf5d9e7fa9a007c, []int{0}
}

func (m *JVMMetricCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JVMMetricCollection.Unmarshal(m, b)
}
func (m *JVMMetricCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JVMMetricCollection.Marshal(b, m, deterministic)
}
func (m *JVMMetricCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JVMMetricCollection.Merge(m, src)
}
func (m *JVMMetricCollection) XXX_Size() int {
	return xxx_messageInfo_JVMMetricCollection.Size(m)
}
func (m *JVMMetricCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_JVMMetricCollection.DiscardUnknown(m)
}

var xxx_messageInfo_JVMMetricCollection proto.InternalMessageInfo

func (m *JVMMetricCollection) GetMetrics() []*JVMMetric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *JVMMetricCollection) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *JVMMetricCollection) GetServiceInstance() string {
	if m != nil {
		return m.ServiceInstance
	}
	return ""
}

type JVMMetric struct {
	Time                 int64         `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Cpu                  *v3.CPU       `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory               []*Memory     `protobuf:"bytes,3,rep,name=memory,proto3" json:"memory,omitempty"`
	MemoryPool           []*MemoryPool `protobuf:"bytes,4,rep,name=memoryPool,proto3" json:"memoryPool,omitempty"`
	Gc                   []*GC         `protobuf:"bytes,5,rep,name=gc,proto3" json:"gc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *JVMMetric) Reset()         { *m = JVMMetric{} }
func (m *JVMMetric) String() string { return proto.CompactTextString(m) }
func (*JVMMetric) ProtoMessage()    {}
func (*JVMMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cf5d9e7fa9a007c, []int{1}
}

func (m *JVMMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JVMMetric.Unmarshal(m, b)
}
func (m *JVMMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JVMMetric.Marshal(b, m, deterministic)
}
func (m *JVMMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JVMMetric.Merge(m, src)
}
func (m *JVMMetric) XXX_Size() int {
	return xxx_messageInfo_JVMMetric.Size(m)
}
func (m *JVMMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_JVMMetric.DiscardUnknown(m)
}

var xxx_messageInfo_JVMMetric proto.InternalMessageInfo

func (m *JVMMetric) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *JVMMetric) GetCpu() *v3.CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *JVMMetric) GetMemory() []*Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *JVMMetric) GetMemoryPool() []*MemoryPool {
	if m != nil {
		return m.MemoryPool
	}
	return nil
}

func (m *JVMMetric) GetGc() []*GC {
	if m != nil {
		return m.Gc
	}
	return nil
}

type Memory struct {
	IsHeap               bool     `protobuf:"varint,1,opt,name=isHeap,proto3" json:"isHeap,omitempty"`
	Init                 int64    `protobuf:"varint,2,opt,name=init,proto3" json:"init,omitempty"`
	Max                  int64    `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
	Used                 int64    `protobuf:"varint,4,opt,name=used,proto3" json:"used,omitempty"`
	Committed            int64    `protobuf:"varint,5,opt,name=committed,proto3" json:"committed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Memory) Reset()         { *m = Memory{} }
func (m *Memory) String() string { return proto.CompactTextString(m) }
func (*Memory) ProtoMessage()    {}
func (*Memory) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cf5d9e7fa9a007c, []int{2}
}

func (m *Memory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Memory.Unmarshal(m, b)
}
func (m *Memory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Memory.Marshal(b, m, deterministic)
}
func (m *Memory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Memory.Merge(m, src)
}
func (m *Memory) XXX_Size() int {
	return xxx_messageInfo_Memory.Size(m)
}
func (m *Memory) XXX_DiscardUnknown() {
	xxx_messageInfo_Memory.DiscardUnknown(m)
}

var xxx_messageInfo_Memory proto.InternalMessageInfo

func (m *Memory) GetIsHeap() bool {
	if m != nil {
		return m.IsHeap
	}
	return false
}

func (m *Memory) GetInit() int64 {
	if m != nil {
		return m.Init
	}
	return 0
}

func (m *Memory) GetMax() int64 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *Memory) GetUsed() int64 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *Memory) GetCommitted() int64 {
	if m != nil {
		return m.Committed
	}
	return 0
}

type MemoryPool struct {
	Type                 PoolType `protobuf:"varint,1,opt,name=type,proto3,enum=PoolType" json:"type,omitempty"`
	Init                 int64    `protobuf:"varint,2,opt,name=init,proto3" json:"init,omitempty"`
	Max                  int64    `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
	Used                 int64    `protobuf:"varint,4,opt,name=used,proto3" json:"used,omitempty"`
	Committed            int64    `protobuf:"varint,5,opt,name=committed,proto3" json:"committed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemoryPool) Reset()         { *m = MemoryPool{} }
func (m *MemoryPool) String() string { return proto.CompactTextString(m) }
func (*MemoryPool) ProtoMessage()    {}
func (*MemoryPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cf5d9e7fa9a007c, []int{3}
}

func (m *MemoryPool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemoryPool.Unmarshal(m, b)
}
func (m *MemoryPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemoryPool.Marshal(b, m, deterministic)
}
func (m *MemoryPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemoryPool.Merge(m, src)
}
func (m *MemoryPool) XXX_Size() int {
	return xxx_messageInfo_MemoryPool.Size(m)
}
func (m *MemoryPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MemoryPool.DiscardUnknown(m)
}

var xxx_messageInfo_MemoryPool proto.InternalMessageInfo

func (m *MemoryPool) GetType() PoolType {
	if m != nil {
		return m.Type
	}
	return PoolType_CODE_CACHE_USAGE
}

func (m *MemoryPool) GetInit() int64 {
	if m != nil {
		return m.Init
	}
	return 0
}

func (m *MemoryPool) GetMax() int64 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *MemoryPool) GetUsed() int64 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *MemoryPool) GetCommitted() int64 {
	if m != nil {
		return m.Committed
	}
	return 0
}

type GC struct {
	Phrase               GCPhrase `protobuf:"varint,1,opt,name=phrase,proto3,enum=GCPhrase" json:"phrase,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Time                 int64    `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GC) Reset()         { *m = GC{} }
func (m *GC) String() string { return proto.CompactTextString(m) }
func (*GC) ProtoMessage()    {}
func (*GC) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cf5d9e7fa9a007c, []int{4}
}

func (m *GC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GC.Unmarshal(m, b)
}
func (m *GC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GC.Marshal(b, m, deterministic)
}
func (m *GC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GC.Merge(m, src)
}
func (m *GC) XXX_Size() int {
	return xxx_messageInfo_GC.Size(m)
}
func (m *GC) XXX_DiscardUnknown() {
	xxx_messageInfo_GC.DiscardUnknown(m)
}

var xxx_messageInfo_GC proto.InternalMessageInfo

func (m *GC) GetPhrase() GCPhrase {
	if m != nil {
		return m.Phrase
	}
	return GCPhrase_NEW
}

func (m *GC) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GC) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterEnum("PoolType", PoolType_name, PoolType_value)
	proto.RegisterEnum("GCPhrase", GCPhrase_name, GCPhrase_value)
	proto.RegisterType((*JVMMetricCollection)(nil), "JVMMetricCollection")
	proto.RegisterType((*JVMMetric)(nil), "JVMMetric")
	proto.RegisterType((*Memory)(nil), "Memory")
	proto.RegisterType((*MemoryPool)(nil), "MemoryPool")
	proto.RegisterType((*GC)(nil), "GC")
}

func init() { proto.RegisterFile("language-agent/JVMMetric.proto", fileDescriptor_0cf5d9e7fa9a007c) }

var fileDescriptor_0cf5d9e7fa9a007c = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xdd, 0x6e, 0xd3, 0x3e,
	0x18, 0xc6, 0x97, 0xba, 0x1f, 0xeb, 0xbb, 0xff, 0x7f, 0x0b, 0xee, 0x34, 0x45, 0xd3, 0x80, 0x51,
	0xed, 0xa0, 0x1a, 0xe0, 0x4a, 0xed, 0x15, 0x94, 0x2c, 0x74, 0x43, 0x6b, 0x1b, 0xdc, 0x7d, 0x48,
	0x9c, 0x4c, 0xc1, 0xb3, 0xb2, 0xa8, 0x49, 0x1c, 0x25, 0x6e, 0xb7, 0x72, 0x02, 0x12, 0x77, 0xc1,
	0x25, 0x70, 0x95, 0xc8, 0x4e, 0xd2, 0x20, 0xc4, 0x29, 0x47, 0x79, 0xdf, 0xdf, 0xfb, 0xc8, 0x7e,
	0x1c, 0xfb, 0x81, 0x17, 0xa1, 0x17, 0xfb, 0x4b, 0xcf, 0xe7, 0x6f, 0x3d, 0x9f, 0xc7, 0xb2, 0xff,
	0xe1, 0x66, 0x32, 0xe1, 0x32, 0x0d, 0x18, 0x49, 0x52, 0x21, 0xc5, 0x61, 0x87, 0x89, 0x28, 0x12,
	0x71, 0xdf, 0xd6, 0x9f, 0x1c, 0x76, 0xbf, 0x42, 0x67, 0xa3, 0xb3, 0x45, 0x18, 0x72, 0x26, 0x03,
	0x11, 0xe3, 0x13, 0x68, 0x45, 0x9a, 0x65, 0x96, 0x71, 0x8c, 0x7a, 0x3b, 0x03, 0x20, 0x1b, 0x19,
	0x2d, 0x47, 0xd8, 0x82, 0x56, 0xc6, 0xd3, 0x55, 0xc0, 0xb8, 0x55, 0x3b, 0x36, 0x7a, 0x6d, 0x5a,
	0xb6, 0xb8, 0x07, 0x7b, 0x45, 0x79, 0x11, 0x67, 0xd2, 0x8b, 0x19, 0xb7, 0x90, 0x56, 0xfc, 0x89,
	0xbb, 0x3f, 0x0c, 0x68, 0x6f, 0x96, 0xc6, 0x18, 0xea, 0x32, 0x88, 0xb8, 0x65, 0x1c, 0x1b, 0x3d,
	0x44, 0x75, 0x8d, 0x0f, 0x00, 0xb1, 0x64, 0xa9, 0x77, 0xd8, 0x19, 0xd4, 0x89, 0xed, 0x5e, 0x53,
	0x05, 0xf0, 0x4b, 0x68, 0x46, 0x3c, 0x12, 0xe9, 0xda, 0x42, 0xda, 0x62, 0x8b, 0x4c, 0x74, 0x4b,
	0x0b, 0x8c, 0x5f, 0x03, 0xe4, 0x95, 0x2b, 0x44, 0x68, 0xd5, 0xb5, 0x68, 0xa7, 0x10, 0x29, 0x44,
	0x7f, 0x1b, 0xe3, 0x0e, 0xd4, 0x7c, 0x66, 0x35, 0xb4, 0x08, 0x91, 0xb1, 0x4d, 0x6b, 0x3e, 0xeb,
	0x3e, 0x41, 0x33, 0x97, 0xe3, 0x03, 0x68, 0x06, 0xd9, 0x39, 0xf7, 0x12, 0x6d, 0x6d, 0x9b, 0x16,
	0x9d, 0x32, 0x1c, 0xc4, 0x81, 0xd4, 0xee, 0x10, 0xd5, 0x35, 0x36, 0x01, 0x45, 0xde, 0x93, 0x3e,
	0x30, 0xa2, 0xaa, 0x54, 0xaa, 0x65, 0xc6, 0xef, 0xad, 0x7a, 0xae, 0x52, 0x35, 0x3e, 0x82, 0xb6,
	0xba, 0x90, 0x40, 0x4a, 0x7e, 0x6f, 0x35, 0xf4, 0xa0, 0x02, 0xdd, 0xef, 0x06, 0x40, 0xe5, 0x14,
	0x3f, 0x87, 0xba, 0x5c, 0x27, 0xf9, 0x7f, 0xd9, 0x1d, 0xb4, 0x89, 0x82, 0x57, 0xeb, 0x84, 0x53,
	0x8d, 0xff, 0x99, 0x8b, 0x8f, 0x50, 0x1b, 0xdb, 0xf8, 0x15, 0x34, 0x93, 0x87, 0xd4, 0xcb, 0xaa,
	0xed, 0xc7, 0xb6, 0xab, 0x01, 0x2d, 0x06, 0x78, 0x1f, 0x1a, 0x4c, 0x2c, 0xe3, 0xd2, 0x41, 0xde,
	0x6c, 0x6e, 0x13, 0x55, 0xb7, 0x79, 0xfa, 0xcd, 0x80, 0xed, 0xd2, 0x3d, 0xde, 0x07, 0xd3, 0x9e,
	0x9d, 0x39, 0x77, 0xf6, 0xc8, 0x3e, 0x77, 0xee, 0xae, 0xe7, 0xa3, 0xb1, 0x63, 0x6e, 0x61, 0x13,
	0xfe, 0x9b, 0x3a, 0xb7, 0x63, 0x67, 0x5a, 0x10, 0x43, 0x91, 0xd9, 0xe5, 0x59, 0x45, 0x6a, 0x18,
	0xc3, 0xee, 0xfc, 0x9a, 0xde, 0x5c, 0xdc, 0xcc, 0x68, 0xc1, 0x10, 0x7e, 0x06, 0xff, 0xbb, 0x0e,
	0x9d, 0x54, 0xb2, 0x3a, 0xee, 0xc0, 0xde, 0xc4, 0xb9, 0x1a, 0xcd, 0xdd, 0x91, 0x5d, 0xae, 0xdf,
	0x38, 0x3d, 0x82, 0xed, 0xf2, 0x00, 0xb8, 0x05, 0x68, 0xea, 0xdc, 0x9a, 0x5b, 0xaa, 0x98, 0x5d,
	0x9e, 0x99, 0xc6, 0xe0, 0x3d, 0x1c, 0x54, 0x4f, 0x9d, 0x27, 0x22, 0x95, 0xf3, 0xe2, 0x51, 0xbf,
	0x81, 0x16, 0xcb, 0x23, 0x82, 0xf7, 0xc9, 0x5f, 0x52, 0x73, 0xd8, 0x26, 0x2a, 0x5b, 0x5e, 0x7c,
	0x9f, 0x75, 0xb7, 0xde, 0x7d, 0x81, 0xa1, 0x48, 0x7d, 0xe2, 0x25, 0x1e, 0x7b, 0xe0, 0x24, 0x5b,
	0xac, 0x1f, 0xbd, 0x70, 0x11, 0xc4, 0x8a, 0x44, 0x24, 0xe6, 0xf2, 0x51, 0xa4, 0x0b, 0x52, 0x46,
	0x96, 0xe8, 0xc8, 0x92, 0xd5, 0xd0, 0x35, 0x3e, 0x9d, 0x54, 0xda, 0x7e, 0xa1, 0xeb, 0x97, 0xba,
	0x7e, 0x1e, 0xed, 0xd5, 0xf0, 0x67, 0xed, 0x70, 0xbe, 0x58, 0xdf, 0x16, 0x4b, 0x4e, 0x73, 0x99,
	0xab, 0x22, 0xcd, 0x44, 0xf8, 0xb9, 0xa9, 0xc3, 0x3d, 0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0x1d,
	0x82, 0xd0, 0xa5, 0x13, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JVMMetricReportServiceClient is the client API for JVMMetricReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JVMMetricReportServiceClient interface {
	Collect(ctx context.Context, in *JVMMetricCollection, opts ...grpc.CallOption) (*v3.Commands, error)
}

type jVMMetricReportServiceClient struct {
	cc *grpc.ClientConn
}

func NewJVMMetricReportServiceClient(cc *grpc.ClientConn) JVMMetricReportServiceClient {
	return &jVMMetricReportServiceClient{cc}
}

func (c *jVMMetricReportServiceClient) Collect(ctx context.Context, in *JVMMetricCollection, opts ...grpc.CallOption) (*v3.Commands, error) {
	out := new(v3.Commands)
	err := c.cc.Invoke(ctx, "/JVMMetricReportService/collect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JVMMetricReportServiceServer is the server API for JVMMetricReportService service.
type JVMMetricReportServiceServer interface {
	Collect(context.Context, *JVMMetricCollection) (*v3.Commands, error)
}

// UnimplementedJVMMetricReportServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJVMMetricReportServiceServer struct {
}

func (*UnimplementedJVMMetricReportServiceServer) Collect(ctx context.Context, req *JVMMetricCollection) (*v3.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}

func RegisterJVMMetricReportServiceServer(s *grpc.Server, srv JVMMetricReportServiceServer) {
	s.RegisterService(&_JVMMetricReportService_serviceDesc, srv)
}

func _JVMMetricReportService_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JVMMetricCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JVMMetricReportServiceServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JVMMetricReportService/Collect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JVMMetricReportServiceServer).Collect(ctx, req.(*JVMMetricCollection))
	}
	return interceptor(ctx, in, info, handler)
}

var _JVMMetricReportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "JVMMetricReportService",
	HandlerType: (*JVMMetricReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "collect",
			Handler:    _JVMMetricReportService_Collect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "language-agent/JVMMetric.proto",
}