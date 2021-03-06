// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/common.proto

package common

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

type TestModeKeys int32

const (
	TestModeKeys_api_test TestModeKeys = 0
)

var TestModeKeys_name = map[int32]string{
	0: "api_test",
}

var TestModeKeys_value = map[string]int32{
	"api_test": 0,
}

func (x TestModeKeys) String() string {
	return proto.EnumName(TestModeKeys_name, int32(x))
}

func (TestModeKeys) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{0}
}

// Logging verbosity level
type LogLevel_LogLevel int32

const (
	LogLevel_DEBUG    LogLevel_LogLevel = 0
	LogLevel_INFO     LogLevel_LogLevel = 1
	LogLevel_WARNING  LogLevel_LogLevel = 2
	LogLevel_ERROR    LogLevel_LogLevel = 3
	LogLevel_CRITICAL LogLevel_LogLevel = 4
	LogLevel_FATAL    LogLevel_LogLevel = 5
)

var LogLevel_LogLevel_name = map[int32]string{
	0: "DEBUG",
	1: "INFO",
	2: "WARNING",
	3: "ERROR",
	4: "CRITICAL",
	5: "FATAL",
}

var LogLevel_LogLevel_value = map[string]int32{
	"DEBUG":    0,
	"INFO":     1,
	"WARNING":  2,
	"ERROR":    3,
	"CRITICAL": 4,
	"FATAL":    5,
}

func (x LogLevel_LogLevel) String() string {
	return proto.EnumName(LogLevel_LogLevel_name, int32(x))
}

func (LogLevel_LogLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{2, 0}
}

// Administrative State
type AdminState_AdminState int32

const (
	// The administrative state of the device is unknown
	AdminState_UNKNOWN AdminState_AdminState = 0
	// The device is pre-provisioned into Voltha, but not contacted by it
	AdminState_PREPROVISIONED AdminState_AdminState = 1
	// The device is enabled for activation and operation
	AdminState_ENABLED AdminState_AdminState = 2
	// The device is disabled and shall not perform its intended forwarding
	// functions other than being available for re-activation.
	AdminState_DISABLED AdminState_AdminState = 3
	// The device is in the state of image download
	AdminState_DOWNLOADING_IMAGE AdminState_AdminState = 4
	// The device is marked to be deleted
	AdminState_DELETED AdminState_AdminState = 5
)

var AdminState_AdminState_name = map[int32]string{
	0: "UNKNOWN",
	1: "PREPROVISIONED",
	2: "ENABLED",
	3: "DISABLED",
	4: "DOWNLOADING_IMAGE",
	5: "DELETED",
}

var AdminState_AdminState_value = map[string]int32{
	"UNKNOWN":           0,
	"PREPROVISIONED":    1,
	"ENABLED":           2,
	"DISABLED":          3,
	"DOWNLOADING_IMAGE": 4,
	"DELETED":           5,
}

func (x AdminState_AdminState) String() string {
	return proto.EnumName(AdminState_AdminState_name, int32(x))
}

func (AdminState_AdminState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{3, 0}
}

// Operational Status
type OperStatus_OperStatus int32

const (
	// The status of the device is unknown at this point
	OperStatus_UNKNOWN OperStatus_OperStatus = 0
	// The device has been discovered, but not yet activated
	OperStatus_DISCOVERED OperStatus_OperStatus = 1
	// The device is being activated (booted, rebooted, upgraded, etc.)
	OperStatus_ACTIVATING OperStatus_OperStatus = 2
	// Service impacting tests are being conducted
	OperStatus_TESTING OperStatus_OperStatus = 3
	// The device is up and active
	OperStatus_ACTIVE OperStatus_OperStatus = 4
	// The device has failed and cannot fulfill its intended role
	OperStatus_FAILED OperStatus_OperStatus = 5
)

var OperStatus_OperStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "DISCOVERED",
	2: "ACTIVATING",
	3: "TESTING",
	4: "ACTIVE",
	5: "FAILED",
}

var OperStatus_OperStatus_value = map[string]int32{
	"UNKNOWN":    0,
	"DISCOVERED": 1,
	"ACTIVATING": 2,
	"TESTING":    3,
	"ACTIVE":     4,
	"FAILED":     5,
}

func (x OperStatus_OperStatus) String() string {
	return proto.EnumName(OperStatus_OperStatus_name, int32(x))
}

func (OperStatus_OperStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{4, 0}
}

// Connectivity Status
type ConnectStatus_ConnectStatus int32

const (
	// The device connectivity status is unknown
	ConnectStatus_UNKNOWN ConnectStatus_ConnectStatus = 0
	// The device cannot be reached by Voltha
	ConnectStatus_UNREACHABLE ConnectStatus_ConnectStatus = 1
	// There is live communication between device and Voltha
	ConnectStatus_REACHABLE ConnectStatus_ConnectStatus = 2
)

var ConnectStatus_ConnectStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "UNREACHABLE",
	2: "REACHABLE",
}

var ConnectStatus_ConnectStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"UNREACHABLE": 1,
	"REACHABLE":   2,
}

func (x ConnectStatus_ConnectStatus) String() string {
	return proto.EnumName(ConnectStatus_ConnectStatus_name, int32(x))
}

func (ConnectStatus_ConnectStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{5, 0}
}

type OperationResp_OperationReturnCode int32

const (
	OperationResp_OPERATION_SUCCESS     OperationResp_OperationReturnCode = 0
	OperationResp_OPERATION_FAILURE     OperationResp_OperationReturnCode = 1
	OperationResp_OPERATION_UNSUPPORTED OperationResp_OperationReturnCode = 2
)

var OperationResp_OperationReturnCode_name = map[int32]string{
	0: "OPERATION_SUCCESS",
	1: "OPERATION_FAILURE",
	2: "OPERATION_UNSUPPORTED",
}

var OperationResp_OperationReturnCode_value = map[string]int32{
	"OPERATION_SUCCESS":     0,
	"OPERATION_FAILURE":     1,
	"OPERATION_UNSUPPORTED": 2,
}

func (x OperationResp_OperationReturnCode) String() string {
	return proto.EnumName(OperationResp_OperationReturnCode_name, int32(x))
}

func (OperationResp_OperationReturnCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{6, 0}
}

// Convey a resource identifier
type ID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{0}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Represents a list of IDs
type IDs struct {
	Items                []*ID    `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDs) Reset()         { *m = IDs{} }
func (m *IDs) String() string { return proto.CompactTextString(m) }
func (*IDs) ProtoMessage()    {}
func (*IDs) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{1}
}

func (m *IDs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDs.Unmarshal(m, b)
}
func (m *IDs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDs.Marshal(b, m, deterministic)
}
func (m *IDs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDs.Merge(m, src)
}
func (m *IDs) XXX_Size() int {
	return xxx_messageInfo_IDs.Size(m)
}
func (m *IDs) XXX_DiscardUnknown() {
	xxx_messageInfo_IDs.DiscardUnknown(m)
}

var xxx_messageInfo_IDs proto.InternalMessageInfo

func (m *IDs) GetItems() []*ID {
	if m != nil {
		return m.Items
	}
	return nil
}

type LogLevel struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogLevel) Reset()         { *m = LogLevel{} }
func (m *LogLevel) String() string { return proto.CompactTextString(m) }
func (*LogLevel) ProtoMessage()    {}
func (*LogLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{2}
}

func (m *LogLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogLevel.Unmarshal(m, b)
}
func (m *LogLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogLevel.Marshal(b, m, deterministic)
}
func (m *LogLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogLevel.Merge(m, src)
}
func (m *LogLevel) XXX_Size() int {
	return xxx_messageInfo_LogLevel.Size(m)
}
func (m *LogLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_LogLevel.DiscardUnknown(m)
}

var xxx_messageInfo_LogLevel proto.InternalMessageInfo

type AdminState struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdminState) Reset()         { *m = AdminState{} }
func (m *AdminState) String() string { return proto.CompactTextString(m) }
func (*AdminState) ProtoMessage()    {}
func (*AdminState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{3}
}

func (m *AdminState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminState.Unmarshal(m, b)
}
func (m *AdminState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminState.Marshal(b, m, deterministic)
}
func (m *AdminState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminState.Merge(m, src)
}
func (m *AdminState) XXX_Size() int {
	return xxx_messageInfo_AdminState.Size(m)
}
func (m *AdminState) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminState.DiscardUnknown(m)
}

var xxx_messageInfo_AdminState proto.InternalMessageInfo

type OperStatus struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperStatus) Reset()         { *m = OperStatus{} }
func (m *OperStatus) String() string { return proto.CompactTextString(m) }
func (*OperStatus) ProtoMessage()    {}
func (*OperStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{4}
}

func (m *OperStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperStatus.Unmarshal(m, b)
}
func (m *OperStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperStatus.Marshal(b, m, deterministic)
}
func (m *OperStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperStatus.Merge(m, src)
}
func (m *OperStatus) XXX_Size() int {
	return xxx_messageInfo_OperStatus.Size(m)
}
func (m *OperStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_OperStatus.DiscardUnknown(m)
}

var xxx_messageInfo_OperStatus proto.InternalMessageInfo

type ConnectStatus struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectStatus) Reset()         { *m = ConnectStatus{} }
func (m *ConnectStatus) String() string { return proto.CompactTextString(m) }
func (*ConnectStatus) ProtoMessage()    {}
func (*ConnectStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{5}
}

func (m *ConnectStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectStatus.Unmarshal(m, b)
}
func (m *ConnectStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectStatus.Marshal(b, m, deterministic)
}
func (m *ConnectStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectStatus.Merge(m, src)
}
func (m *ConnectStatus) XXX_Size() int {
	return xxx_messageInfo_ConnectStatus.Size(m)
}
func (m *ConnectStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectStatus proto.InternalMessageInfo

type OperationResp struct {
	// Return code
	Code OperationResp_OperationReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=voltha.OperationResp_OperationReturnCode" json:"code,omitempty"`
	// Additional Info
	AdditionalInfo       string   `protobuf:"bytes,2,opt,name=additional_info,json=additionalInfo,proto3" json:"additional_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationResp) Reset()         { *m = OperationResp{} }
func (m *OperationResp) String() string { return proto.CompactTextString(m) }
func (*OperationResp) ProtoMessage()    {}
func (*OperationResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{6}
}

func (m *OperationResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationResp.Unmarshal(m, b)
}
func (m *OperationResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationResp.Marshal(b, m, deterministic)
}
func (m *OperationResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationResp.Merge(m, src)
}
func (m *OperationResp) XXX_Size() int {
	return xxx_messageInfo_OperationResp.Size(m)
}
func (m *OperationResp) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationResp.DiscardUnknown(m)
}

var xxx_messageInfo_OperationResp proto.InternalMessageInfo

func (m *OperationResp) GetCode() OperationResp_OperationReturnCode {
	if m != nil {
		return m.Code
	}
	return OperationResp_OPERATION_SUCCESS
}

func (m *OperationResp) GetAdditionalInfo() string {
	if m != nil {
		return m.AdditionalInfo
	}
	return ""
}

func init() {
	proto.RegisterEnum("voltha.TestModeKeys", TestModeKeys_name, TestModeKeys_value)
	proto.RegisterEnum("voltha.LogLevel_LogLevel", LogLevel_LogLevel_name, LogLevel_LogLevel_value)
	proto.RegisterEnum("voltha.AdminState_AdminState", AdminState_AdminState_name, AdminState_AdminState_value)
	proto.RegisterEnum("voltha.OperStatus_OperStatus", OperStatus_OperStatus_name, OperStatus_OperStatus_value)
	proto.RegisterEnum("voltha.ConnectStatus_ConnectStatus", ConnectStatus_ConnectStatus_name, ConnectStatus_ConnectStatus_value)
	proto.RegisterEnum("voltha.OperationResp_OperationReturnCode", OperationResp_OperationReturnCode_name, OperationResp_OperationReturnCode_value)
	proto.RegisterType((*ID)(nil), "voltha.ID")
	proto.RegisterType((*IDs)(nil), "voltha.IDs")
	proto.RegisterType((*LogLevel)(nil), "voltha.LogLevel")
	proto.RegisterType((*AdminState)(nil), "voltha.AdminState")
	proto.RegisterType((*OperStatus)(nil), "voltha.OperStatus")
	proto.RegisterType((*ConnectStatus)(nil), "voltha.ConnectStatus")
	proto.RegisterType((*OperationResp)(nil), "voltha.OperationResp")
}

func init() { proto.RegisterFile("voltha_protos/common.proto", fileDescriptor_c2e3fd231961e826) }

var fileDescriptor_c2e3fd231961e826 = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x5f, 0x4f, 0xdb, 0x3e,
	0x14, 0x6d, 0xd2, 0x3f, 0xc0, 0x2d, 0x94, 0xfc, 0xfc, 0x1b, 0x12, 0x43, 0x9b, 0x54, 0xe5, 0x05,
	0xb6, 0x89, 0x22, 0xb1, 0xb7, 0x69, 0x7b, 0x30, 0x89, 0xe9, 0x2c, 0x82, 0x5d, 0x39, 0x09, 0x48,
	0x7b, 0xa0, 0x0a, 0x8d, 0x29, 0x91, 0xda, 0x38, 0x6a, 0x0c, 0x12, 0x8f, 0xfb, 0x80, 0xfb, 0x0a,
	0xfb, 0x0c, 0x7b, 0xda, 0xf3, 0xe4, 0x04, 0xd4, 0x76, 0xda, 0x5b, 0xce, 0x39, 0xd7, 0xf7, 0xf8,
	0xe4, 0x5e, 0xc3, 0xc1, 0xa3, 0x9a, 0xe9, 0xfb, 0x64, 0x5c, 0x2c, 0x94, 0x56, 0xe5, 0xc9, 0x44,
	0xcd, 0xe7, 0x2a, 0x1f, 0x54, 0x08, 0x75, 0x6a, 0xed, 0xa0, 0xbf, 0x5e, 0xf3, 0x94, 0xe4, 0xd3,
	0xb1, 0x2a, 0x74, 0xa6, 0xf2, 0xb2, 0xae, 0x74, 0x5f, 0x81, 0x4d, 0x7d, 0xd4, 0x03, 0x3b, 0x4b,
	0xf7, 0xad, 0xbe, 0x75, 0xb4, 0x25, 0xec, 0x2c, 0x75, 0x0f, 0xa1, 0x49, 0xfd, 0x12, 0xf5, 0xa1,
	0x9d, 0x69, 0x39, 0x2f, 0xf7, 0xad, 0x7e, 0xf3, 0xa8, 0x7b, 0x0a, 0x83, 0xba, 0xdd, 0x80, 0xfa,
	0xa2, 0x16, 0xdc, 0x09, 0x6c, 0x06, 0x6a, 0x1a, 0xc8, 0x47, 0x39, 0x73, 0x47, 0xcb, 0x6f, 0xb4,
	0x05, 0x6d, 0x9f, 0x9c, 0xc5, 0x43, 0xa7, 0x81, 0x36, 0xa1, 0x45, 0xd9, 0x39, 0x77, 0x2c, 0xd4,
	0x85, 0x8d, 0x6b, 0x2c, 0x18, 0x65, 0x43, 0xc7, 0x36, 0x15, 0x44, 0x08, 0x2e, 0x9c, 0x26, 0xda,
	0x86, 0x4d, 0x4f, 0xd0, 0x88, 0x7a, 0x38, 0x70, 0x5a, 0x46, 0x38, 0xc7, 0x11, 0x0e, 0x9c, 0xf6,
	0xa7, 0xf6, 0xaf, 0xdf, 0x3f, 0xde, 0x36, 0xdc, 0xef, 0x16, 0x00, 0x4e, 0xe7, 0x59, 0x1e, 0xea,
	0x44, 0x4b, 0x77, 0xb6, 0x8a, 0x4c, 0xd3, 0x98, 0x5d, 0x30, 0x7e, 0xcd, 0x9c, 0x06, 0x42, 0xd0,
	0x1b, 0x09, 0x32, 0x12, 0xfc, 0x8a, 0x86, 0x94, 0x33, 0xe2, 0xd7, 0xae, 0x84, 0xe1, 0xb3, 0x80,
	0xf8, 0x8e, 0x6d, 0xac, 0x7c, 0x1a, 0xd6, 0xa8, 0x89, 0xf6, 0xe0, 0x3f, 0x9f, 0x5f, 0xb3, 0x80,
	0x63, 0x9f, 0xb2, 0xe1, 0x98, 0x5e, 0xe2, 0x21, 0x71, 0x5a, 0xe6, 0x84, 0x4f, 0x02, 0x12, 0x11,
	0x7f, 0x79, 0x87, 0x12, 0x80, 0x17, 0x72, 0x61, 0x3c, 0x1f, 0x4a, 0xf7, 0x66, 0x15, 0xad, 0x5f,
	0xa1, 0x07, 0xe0, 0xd3, 0xd0, 0xe3, 0x57, 0x44, 0x54, 0xf6, 0x3d, 0x00, 0xec, 0x45, 0xf4, 0x0a,
	0x47, 0x75, 0xee, 0x2e, 0x6c, 0x44, 0x24, 0xac, 0x40, 0x13, 0x01, 0x74, 0x2a, 0xd1, 0xb8, 0x02,
	0x74, 0xce, 0x31, 0x0d, 0x56, 0x4d, 0x23, 0xd8, 0xf1, 0x54, 0x9e, 0xcb, 0x89, 0x7e, 0xf6, 0xfd,
	0xfc, 0x17, 0xb1, 0x6e, 0xbd, 0x0b, 0xdd, 0x98, 0x09, 0x82, 0xbd, 0xaf, 0x26, 0xa0, 0x63, 0xa1,
	0x1d, 0xd8, 0x5a, 0x42, 0xfb, 0xa5, 0xeb, 0x4f, 0x0b, 0x76, 0xcc, 0xed, 0x13, 0xb3, 0x07, 0x42,
	0x96, 0x05, 0xfa, 0x02, 0xad, 0x89, 0x4a, 0x65, 0xb5, 0x00, 0xbd, 0xd3, 0x77, 0x2f, 0x63, 0x5e,
	0x2b, 0x5a, 0x45, 0xfa, 0x61, 0x91, 0x7b, 0x2a, 0x95, 0xa2, 0x3a, 0x86, 0x0e, 0x61, 0x37, 0x49,
	0xd3, 0xcc, 0x68, 0xc9, 0x6c, 0x9c, 0xe5, 0x77, 0x6a, 0xdf, 0xae, 0x56, 0xa9, 0xb7, 0xa4, 0x69,
	0x7e, 0xa7, 0xdc, 0x1b, 0xf8, 0xff, 0x1f, 0x5d, 0xcc, 0x18, 0xf8, 0x88, 0x08, 0x1c, 0x51, 0xce,
	0xc6, 0x61, 0xec, 0x79, 0x24, 0x0c, 0x9d, 0xc6, 0x3a, 0x6d, 0x7e, 0x4d, 0x2c, 0x4c, 0xa8, 0xd7,
	0xb0, 0xb7, 0xa4, 0x63, 0x16, 0xc6, 0xa3, 0x11, 0x17, 0x66, 0x56, 0x2f, 0x01, 0xdf, 0xbf, 0x81,
	0xed, 0x48, 0x96, 0xfa, 0x52, 0xa5, 0xf2, 0x42, 0x3e, 0x95, 0x66, 0xe8, 0x49, 0x91, 0x8d, 0xb5,
	0x2c, 0xb5, 0xd3, 0x38, 0x3b, 0xfe, 0xf6, 0x61, 0x9a, 0xe9, 0xfb, 0x87, 0xdb, 0xc1, 0x44, 0xcd,
	0x4f, 0x54, 0x21, 0xf3, 0x89, 0x5a, 0xa4, 0x27, 0x75, 0xe6, 0xe3, 0xe7, 0x97, 0x32, 0x55, 0xcf,
	0x0f, 0xea, 0xb6, 0x53, 0x31, 0x1f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x0d, 0x0e, 0x19,
	0x6f, 0x03, 0x00, 0x00,
}
