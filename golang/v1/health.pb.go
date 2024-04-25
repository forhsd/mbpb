// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: health.proto

package mbpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{0}
}

func (x *HealthRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type HealthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Healthy bool `protobuf:"varint,1,opt,name=Healthy,proto3" json:"Healthy,omitempty"`
}

func (x *HealthReply) Reset() {
	*x = HealthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthReply) ProtoMessage() {}

func (x *HealthReply) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthReply.ProtoReflect.Descriptor instead.
func (*HealthReply) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{1}
}

func (x *HealthReply) GetHealthy() bool {
	if x != nil {
		return x.Healthy
	}
	return false
}

// 概览
type Overview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key          string    `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`                                   //
	Owner        string    `protobuf:"bytes,2,opt,name=Owner,proto3" json:"Owner,omitempty"`                               // 节点
	EnterpriseID string    `protobuf:"bytes,3,opt,name=EnterpriseID,proto3" json:"EnterpriseID,omitempty"`                 // 企业ID
	CardId       int64     `protobuf:"varint,4,opt,name=CardId,proto3" json:"CardId,omitempty"`                            // CardId
	SequenceID   string    `protobuf:"bytes,5,opt,name=SequenceID,proto3" json:"SequenceID,omitempty"`                     // 序列ID
	QueryUPtime  *int64    `protobuf:"varint,7,opt,name=QueryUPtime,proto3,oneof" json:"QueryUPtime,omitempty"`            // 查询耗时(毫秒)
	WriteUPtime  *int64    `protobuf:"varint,8,opt,name=WriteUPtime,proto3,oneof" json:"WriteUPtime,omitempty"`            // 写入耗时(毫秒)
	TakeUpTime   *int64    `protobuf:"varint,9,opt,name=TakeUpTime,proto3,oneof" json:"TakeUpTime,omitempty"`              // 运行耗时(毫秒)
	RunStatus    RunStatus `protobuf:"varint,10,opt,name=RunStatus,proto3,enum=mbpb.RunStatus" json:"RunStatus,omitempty"` // 运行状态
	Progress     int32     `protobuf:"varint,11,opt,name=Progress,proto3" json:"Progress,omitempty"`                       // 当前进度
	StartTime    int64     `protobuf:"varint,12,opt,name=StartTime,proto3" json:"StartTime,omitempty"`                     // 开始运行时间 精确到纳秒
	EndTime      int64     `protobuf:"varint,13,opt,name=EndTime,proto3" json:"EndTime,omitempty"`                         // 结束运行时间 精确到纳秒
	NextRunTime  *int64    `protobuf:"varint,14,opt,name=NextRunTime,proto3,oneof" json:"NextRunTime,omitempty"`           // 下次运行时间 精确到纳秒
	Detail       *Detail   `protobuf:"bytes,15,opt,name=Detail,proto3" json:"Detail,omitempty"`                            // 详情
	RunType      RunType   `protobuf:"varint,16,opt,name=RunType,proto3,enum=mbpb.RunType" json:"RunType,omitempty"`       // 运行类型
}

func (x *Overview) Reset() {
	*x = Overview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Overview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Overview) ProtoMessage() {}

func (x *Overview) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Overview.ProtoReflect.Descriptor instead.
func (*Overview) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{2}
}

func (x *Overview) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Overview) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Overview) GetEnterpriseID() string {
	if x != nil {
		return x.EnterpriseID
	}
	return ""
}

func (x *Overview) GetCardId() int64 {
	if x != nil {
		return x.CardId
	}
	return 0
}

func (x *Overview) GetSequenceID() string {
	if x != nil {
		return x.SequenceID
	}
	return ""
}

func (x *Overview) GetQueryUPtime() int64 {
	if x != nil && x.QueryUPtime != nil {
		return *x.QueryUPtime
	}
	return 0
}

func (x *Overview) GetWriteUPtime() int64 {
	if x != nil && x.WriteUPtime != nil {
		return *x.WriteUPtime
	}
	return 0
}

func (x *Overview) GetTakeUpTime() int64 {
	if x != nil && x.TakeUpTime != nil {
		return *x.TakeUpTime
	}
	return 0
}

func (x *Overview) GetRunStatus() RunStatus {
	if x != nil {
		return x.RunStatus
	}
	return RunStatus_Unknown
}

func (x *Overview) GetProgress() int32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *Overview) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Overview) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *Overview) GetNextRunTime() int64 {
	if x != nil && x.NextRunTime != nil {
		return *x.NextRunTime
	}
	return 0
}

func (x *Overview) GetDetail() *Detail {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *Overview) GetRunType() RunType {
	if x != nil {
		return x.RunType
	}
	return RunType_Cycle
}

// 详情
type Detail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowsCount *int64    `protobuf:"varint,1,opt,name=RowsCount,proto3,oneof" json:"RowsCount,omitempty"` // 行数
	Columns   []*Column `protobuf:"bytes,2,rep,name=Columns,proto3" json:"Columns,omitempty"`            // 字段类
	OutTable  *Table    `protobuf:"bytes,3,opt,name=OutTable,proto3,oneof" json:"OutTable,omitempty"`    // 实体表
	Error     *Error    `protobuf:"bytes,4,opt,name=Error,proto3,oneof" json:"Error,omitempty"`          // 错误明细
}

func (x *Detail) Reset() {
	*x = Detail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Detail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Detail) ProtoMessage() {}

func (x *Detail) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Detail.ProtoReflect.Descriptor instead.
func (*Detail) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{3}
}

func (x *Detail) GetRowsCount() int64 {
	if x != nil && x.RowsCount != nil {
		return *x.RowsCount
	}
	return 0
}

func (x *Detail) GetColumns() []*Column {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *Detail) GetOutTable() *Table {
	if x != nil {
		return x.OutTable
	}
	return nil
}

func (x *Detail) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        string                 `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	SequenceID int64                  `protobuf:"varint,2,opt,name=SequenceID,proto3" json:"SequenceID,omitempty"`
	Error      *Error                 `protobuf:"bytes,3,opt,name=Error,proto3,oneof" json:"Error,omitempty"`
	Now        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=Now,proto3,oneof" json:"Now,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{4}
}

func (x *Ack) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Ack) GetSequenceID() int64 {
	if x != nil {
		return x.SequenceID
	}
	return 0
}

func (x *Ack) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *Ack) GetNow() *timestamppb.Timestamp {
	if x != nil {
		return x.Now
	}
	return nil
}

var File_health_proto protoreflect.FileDescriptor

var file_health_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x6d, 0x62, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x27, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x79, 0x22, 0xb9, 0x04, 0x0a, 0x08, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12, 0x10,
	0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x73, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x61,
	0x72, 0x64, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x43, 0x61, 0x72, 0x64,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x44,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x49, 0x44, 0x12, 0x25, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x50, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x55, 0x50, 0x74, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x55, 0x50, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01,
	0x52, 0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x55, 0x50, 0x74, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x23, 0x0a, 0x0a, 0x54, 0x61, 0x6b, 0x65, 0x55, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x0a, 0x54, 0x61, 0x6b, 0x65, 0x55, 0x70, 0x54, 0x69,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6d, 0x62, 0x70, 0x62, 0x2e,
	0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x52, 0x75, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x4e, 0x65, 0x78, 0x74,
	0x52, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x48, 0x03, 0x52,
	0x0b, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x24, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x6d, 0x62, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x27, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x6d, 0x62, 0x70, 0x62, 0x2e, 0x52, 0x75,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x52, 0x75, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x50, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x57, 0x72, 0x69, 0x74, 0x65, 0x55, 0x50, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x54, 0x61, 0x6b, 0x65, 0x55, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xce, 0x01,
	0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x09, 0x52, 0x6f, 0x77, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x09, 0x52,
	0x6f, 0x77, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x07, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d,
	0x62, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x07, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x62, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x48, 0x01, 0x52, 0x08, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x26, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x6d, 0x62, 0x70, 0x62, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x02, 0x52,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x52, 0x6f,
	0x77, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x4f, 0x75, 0x74, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xa4,
	0x01, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x53, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x62, 0x70, 0x62, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01,
	0x12, 0x31, 0x0a, 0x03, 0x4e, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x03, 0x4e, 0x6f, 0x77,
	0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x4e, 0x6f, 0x77, 0x32, 0x6b, 0x0a, 0x06, 0x4d, 0x42, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x36, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x13, 0x2e, 0x6d, 0x62, 0x70, 0x62,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x6d, 0x62, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x29, 0x0a, 0x06, 0x52, 0x65, 0x56, 0x69, 0x65,
	0x77, 0x12, 0x0e, 0x2e, 0x6d, 0x62, 0x70, 0x62, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65,
	0x77, 0x1a, 0x09, 0x2e, 0x6d, 0x62, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x4e, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0c, 0x4d, 0x42, 0x45,
	0x74, 0x6c, 0x70, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x25, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6f, 0x72, 0x68, 0x73, 0x64, 0x2f, 0x6d,
	0x62, 0x70, 0x62, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x62,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_health_proto_rawDescOnce sync.Once
	file_health_proto_rawDescData = file_health_proto_rawDesc
)

func file_health_proto_rawDescGZIP() []byte {
	file_health_proto_rawDescOnce.Do(func() {
		file_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_health_proto_rawDescData)
	})
	return file_health_proto_rawDescData
}

var file_health_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_health_proto_goTypes = []interface{}{
	(*HealthRequest)(nil),         // 0: mbpb.HealthRequest
	(*HealthReply)(nil),           // 1: mbpb.HealthReply
	(*Overview)(nil),              // 2: mbpb.Overview
	(*Detail)(nil),                // 3: mbpb.Detail
	(*Ack)(nil),                   // 4: mbpb.Ack
	(RunStatus)(0),                // 5: mbpb.RunStatus
	(RunType)(0),                  // 6: mbpb.RunType
	(*Column)(nil),                // 7: mbpb.Column
	(*Table)(nil),                 // 8: mbpb.Table
	(*Error)(nil),                 // 9: mbpb.Error
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_health_proto_depIdxs = []int32{
	5,  // 0: mbpb.Overview.RunStatus:type_name -> mbpb.RunStatus
	3,  // 1: mbpb.Overview.Detail:type_name -> mbpb.Detail
	6,  // 2: mbpb.Overview.RunType:type_name -> mbpb.RunType
	7,  // 3: mbpb.Detail.Columns:type_name -> mbpb.Column
	8,  // 4: mbpb.Detail.OutTable:type_name -> mbpb.Table
	9,  // 5: mbpb.Detail.Error:type_name -> mbpb.Error
	9,  // 6: mbpb.Ack.Error:type_name -> mbpb.Error
	10, // 7: mbpb.Ack.Now:type_name -> google.protobuf.Timestamp
	0,  // 8: mbpb.MBLink.Health:input_type -> mbpb.HealthRequest
	2,  // 9: mbpb.MBLink.ReView:input_type -> mbpb.Overview
	1,  // 10: mbpb.MBLink.Health:output_type -> mbpb.HealthReply
	4,  // 11: mbpb.MBLink.ReView:output_type -> mbpb.Ack
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_health_proto_init() }
func file_health_proto_init() {
	if File_health_proto != nil {
		return
	}
	file_entrypoint_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthRequest); i {
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
		file_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthReply); i {
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
		file_health_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Overview); i {
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
		file_health_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Detail); i {
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
		file_health_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
	file_health_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_health_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_health_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_health_proto_goTypes,
		DependencyIndexes: file_health_proto_depIdxs,
		MessageInfos:      file_health_proto_msgTypes,
	}.Build()
	File_health_proto = out.File
	file_health_proto_rawDesc = nil
	file_health_proto_goTypes = nil
	file_health_proto_depIdxs = nil
}
