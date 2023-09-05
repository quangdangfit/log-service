// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: log/log.proto

package log

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogLevel int32

const (
	LogLevel_LOG_LEVEL_INVALID LogLevel = 0
	LogLevel_LOG_LEVEL_DEBUG   LogLevel = 1
	LogLevel_LOG_LEVEL_INFO    LogLevel = 2
	LogLevel_LOG_LEVEL_WARNING LogLevel = 3
	LogLevel_LOG_LEVEL_ERROR   LogLevel = 4
	LogLevel_LOG_LEVEL_FATAL   LogLevel = 5
)

// Enum value maps for LogLevel.
var (
	LogLevel_name = map[int32]string{
		0: "LOG_LEVEL_INVALID",
		1: "LOG_LEVEL_DEBUG",
		2: "LOG_LEVEL_INFO",
		3: "LOG_LEVEL_WARNING",
		4: "LOG_LEVEL_ERROR",
		5: "LOG_LEVEL_FATAL",
	}
	LogLevel_value = map[string]int32{
		"LOG_LEVEL_INVALID": 0,
		"LOG_LEVEL_DEBUG":   1,
		"LOG_LEVEL_INFO":    2,
		"LOG_LEVEL_WARNING": 3,
		"LOG_LEVEL_ERROR":   4,
		"LOG_LEVEL_FATAL":   5,
	}
)

func (x LogLevel) Enum() *LogLevel {
	p := new(LogLevel)
	*p = x
	return p
}

func (x LogLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_log_log_proto_enumTypes[0].Descriptor()
}

func (LogLevel) Type() protoreflect.EnumType {
	return &file_log_log_proto_enumTypes[0]
}

func (x LogLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogLevel.Descriptor instead.
func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return file_log_log_proto_rawDescGZIP(), []int{0}
}

type LogInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Level       string `protobuf:"bytes,3,opt,name=level,proto3" json:"level,omitempty"`
	Content     []byte `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	TraceId     string `protobuf:"bytes,5,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	CreatedAt   string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *LogInfo) Reset() {
	*x = LogInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogInfo) ProtoMessage() {}

func (x *LogInfo) ProtoReflect() protoreflect.Message {
	mi := &file_log_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogInfo.ProtoReflect.Descriptor instead.
func (*LogInfo) Descriptor() ([]byte, []int) {
	return file_log_log_proto_rawDescGZIP(), []int{0}
}

func (x *LogInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LogInfo) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *LogInfo) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *LogInfo) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *LogInfo) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *LogInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type AddLogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string   `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Level       LogLevel `protobuf:"varint,2,opt,name=level,proto3,enum=log.LogLevel" json:"level,omitempty"`
	Content     []byte   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	TraceId     string   `protobuf:"bytes,4,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
}

func (x *AddLogReq) Reset() {
	*x = AddLogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLogReq) ProtoMessage() {}

func (x *AddLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_log_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLogReq.ProtoReflect.Descriptor instead.
func (*AddLogReq) Descriptor() ([]byte, []int) {
	return file_log_log_proto_rawDescGZIP(), []int{1}
}

func (x *AddLogReq) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *AddLogReq) GetLevel() LogLevel {
	if x != nil {
		return x.Level
	}
	return LogLevel_LOG_LEVEL_INVALID
}

func (x *AddLogReq) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *AddLogReq) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

type AddLogRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Log *LogInfo `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *AddLogRes) Reset() {
	*x = AddLogRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddLogRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLogRes) ProtoMessage() {}

func (x *AddLogRes) ProtoReflect() protoreflect.Message {
	mi := &file_log_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLogRes.ProtoReflect.Descriptor instead.
func (*AddLogRes) Descriptor() ([]byte, []int) {
	return file_log_log_proto_rawDescGZIP(), []int{2}
}

func (x *AddLogRes) GetLog() *LogInfo {
	if x != nil {
		return x.Log
	}
	return nil
}

type GetLogsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string         `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Level       LogLevel       `protobuf:"varint,2,opt,name=level,proto3,enum=log.LogLevel" json:"level,omitempty"`
	TraceId     string         `protobuf:"bytes,3,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	CreatedFrom int64          `protobuf:"varint,4,opt,name=created_from,json=createdFrom,proto3" json:"created_from,omitempty"`
	CreatedTo   int64          `protobuf:"varint,5,opt,name=created_to,json=createdTo,proto3" json:"created_to,omitempty"`
	Pagination  *PaginationReq `protobuf:"bytes,6,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Sort        *SortReq       `protobuf:"bytes,7,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *GetLogsReq) Reset() {
	*x = GetLogsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogsReq) ProtoMessage() {}

func (x *GetLogsReq) ProtoReflect() protoreflect.Message {
	mi := &file_log_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogsReq.ProtoReflect.Descriptor instead.
func (*GetLogsReq) Descriptor() ([]byte, []int) {
	return file_log_log_proto_rawDescGZIP(), []int{3}
}

func (x *GetLogsReq) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *GetLogsReq) GetLevel() LogLevel {
	if x != nil {
		return x.Level
	}
	return LogLevel_LOG_LEVEL_INVALID
}

func (x *GetLogsReq) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *GetLogsReq) GetCreatedFrom() int64 {
	if x != nil {
		return x.CreatedFrom
	}
	return 0
}

func (x *GetLogsReq) GetCreatedTo() int64 {
	if x != nil {
		return x.CreatedTo
	}
	return 0
}

func (x *GetLogsReq) GetPagination() *PaginationReq {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *GetLogsReq) GetSort() *SortReq {
	if x != nil {
		return x.Sort
	}
	return nil
}

type GetLogsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs       []*LogInfo     `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
	Pagination *PaginationRes `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetLogsRes) Reset() {
	*x = GetLogsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogsRes) ProtoMessage() {}

func (x *GetLogsRes) ProtoReflect() protoreflect.Message {
	mi := &file_log_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogsRes.ProtoReflect.Descriptor instead.
func (*GetLogsRes) Descriptor() ([]byte, []int) {
	return file_log_log_proto_rawDescGZIP(), []int{4}
}

func (x *GetLogsRes) GetLogs() []*LogInfo {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *GetLogsRes) GetPagination() *PaginationRes {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_log_log_proto protoreflect.FileDescriptor

var file_log_log_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x6c, 0x6f, 0x67, 0x1a, 0x0e, 0x6c, 0x6f, 0x67, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x88, 0x01,
	0x0a, 0x09, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e,
	0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x03, 0x6c, 0x6f, 0x67, 0x22, 0x87, 0x02, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x12, 0x32, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6c, 0x6f, 0x67, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f,
	0x67, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22,
	0x62, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x12, 0x20, 0x0a,
	0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f,
	0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12,
	0x32, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2a, 0x8b, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x15, 0x0a, 0x11, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x02,
	0x12, 0x15, 0x0a, 0x11, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x57, 0x41,
	0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f,
	0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10,
	0x05, 0x32, 0x63, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x28, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x2e,
	0x41, 0x64, 0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x2e,
	0x41, 0x64, 0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x67, 0x73, 0x12, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x6c, 0x6f, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_log_log_proto_rawDescOnce sync.Once
	file_log_log_proto_rawDescData = file_log_log_proto_rawDesc
)

func file_log_log_proto_rawDescGZIP() []byte {
	file_log_log_proto_rawDescOnce.Do(func() {
		file_log_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_log_proto_rawDescData)
	})
	return file_log_log_proto_rawDescData
}

var file_log_log_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_log_log_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_log_log_proto_goTypes = []interface{}{
	(LogLevel)(0),         // 0: log.LogLevel
	(*LogInfo)(nil),       // 1: log.LogInfo
	(*AddLogReq)(nil),     // 2: log.AddLogReq
	(*AddLogRes)(nil),     // 3: log.AddLogRes
	(*GetLogsReq)(nil),    // 4: log.GetLogsReq
	(*GetLogsRes)(nil),    // 5: log.GetLogsRes
	(*PaginationReq)(nil), // 6: log.PaginationReq
	(*SortReq)(nil),       // 7: log.SortReq
	(*PaginationRes)(nil), // 8: log.PaginationRes
}
var file_log_log_proto_depIdxs = []int32{
	0, // 0: log.AddLogReq.level:type_name -> log.LogLevel
	1, // 1: log.AddLogRes.log:type_name -> log.LogInfo
	0, // 2: log.GetLogsReq.level:type_name -> log.LogLevel
	6, // 3: log.GetLogsReq.pagination:type_name -> log.PaginationReq
	7, // 4: log.GetLogsReq.sort:type_name -> log.SortReq
	1, // 5: log.GetLogsRes.logs:type_name -> log.LogInfo
	8, // 6: log.GetLogsRes.pagination:type_name -> log.PaginationRes
	2, // 7: log.LogService.AddLog:input_type -> log.AddLogReq
	4, // 8: log.LogService.GetLogs:input_type -> log.GetLogsReq
	3, // 9: log.LogService.AddLog:output_type -> log.AddLogRes
	5, // 10: log.LogService.GetLogs:output_type -> log.GetLogsRes
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_log_log_proto_init() }
func file_log_log_proto_init() {
	if File_log_log_proto != nil {
		return
	}
	file_log_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_log_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_log_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddLogReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_log_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddLogRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_log_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogsReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_log_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogsRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_log_log_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_log_log_proto_goTypes,
		DependencyIndexes: file_log_log_proto_depIdxs,
		EnumInfos:         file_log_log_proto_enumTypes,
		MessageInfos:      file_log_log_proto_msgTypes,
	}.Build()
	File_log_log_proto = out.File
	file_log_log_proto_rawDesc = nil
	file_log_log_proto_goTypes = nil
	file_log_log_proto_depIdxs = nil
}
