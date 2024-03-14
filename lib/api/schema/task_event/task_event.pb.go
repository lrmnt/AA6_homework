// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.2
// source: api/schema/task_event/task_event.proto

package task_event

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

type Event int32

const (
	Event_EVENT_UNKNOWN    Event = 0
	Event_EVENT_REASSIGNED Event = 1
	Event_EVENT_DONE       Event = 2
)

// Enum value maps for Event.
var (
	Event_name = map[int32]string{
		0: "EVENT_UNKNOWN",
		1: "EVENT_REASSIGNED",
		2: "EVENT_DONE",
	}
	Event_value = map[string]int32{
		"EVENT_UNKNOWN":    0,
		"EVENT_REASSIGNED": 1,
		"EVENT_DONE":       2,
	}
)

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}

func (x Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event) Descriptor() protoreflect.EnumDescriptor {
	return file_api_schema_task_event_task_event_proto_enumTypes[0].Descriptor()
}

func (Event) Type() protoreflect.EnumType {
	return &file_api_schema_task_event_task_event_proto_enumTypes[0]
}

func (x Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event.Descriptor instead.
func (Event) EnumDescriptor() ([]byte, []int) {
	return file_api_schema_task_event_task_event_proto_rawDescGZIP(), []int{0}
}

type TaskEventV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event          Event  `protobuf:"varint,1,opt,name=event,proto3,enum=task_event.Event" json:"event,omitempty"`
	Timestamp      int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	EventId        string `protobuf:"bytes,3,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	TaskId         string `protobuf:"bytes,4,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	AssigneeUserId string `protobuf:"bytes,5,opt,name=assignee_user_id,json=assigneeUserId,proto3" json:"assignee_user_id,omitempty"`
}

func (x *TaskEventV1) Reset() {
	*x = TaskEventV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_schema_task_event_task_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskEventV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskEventV1) ProtoMessage() {}

func (x *TaskEventV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_schema_task_event_task_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskEventV1.ProtoReflect.Descriptor instead.
func (*TaskEventV1) Descriptor() ([]byte, []int) {
	return file_api_schema_task_event_task_event_proto_rawDescGZIP(), []int{0}
}

func (x *TaskEventV1) GetEvent() Event {
	if x != nil {
		return x.Event
	}
	return Event_EVENT_UNKNOWN
}

func (x *TaskEventV1) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *TaskEventV1) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *TaskEventV1) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *TaskEventV1) GetAssigneeUserId() string {
	if x != nil {
		return x.AssigneeUserId
	}
	return ""
}

var File_api_schema_task_event_task_event_proto protoreflect.FileDescriptor

var file_api_schema_task_event_task_event_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x22, 0xb2, 0x01, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x56, 0x31, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12,
	0x28, 0x0a, 0x10, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x40, 0x0a, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x52,
	0x45, 0x41, 0x53, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x42, 0x32, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x72, 0x6d, 0x6e, 0x74, 0x2f,
	0x41, 0x41, 0x36, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x69, 0x62,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_schema_task_event_task_event_proto_rawDescOnce sync.Once
	file_api_schema_task_event_task_event_proto_rawDescData = file_api_schema_task_event_task_event_proto_rawDesc
)

func file_api_schema_task_event_task_event_proto_rawDescGZIP() []byte {
	file_api_schema_task_event_task_event_proto_rawDescOnce.Do(func() {
		file_api_schema_task_event_task_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_schema_task_event_task_event_proto_rawDescData)
	})
	return file_api_schema_task_event_task_event_proto_rawDescData
}

var file_api_schema_task_event_task_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_schema_task_event_task_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_schema_task_event_task_event_proto_goTypes = []interface{}{
	(Event)(0),          // 0: task_event.Event
	(*TaskEventV1)(nil), // 1: task_event.TaskEventV1
}
var file_api_schema_task_event_task_event_proto_depIdxs = []int32{
	0, // 0: task_event.TaskEventV1.event:type_name -> task_event.Event
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_schema_task_event_task_event_proto_init() }
func file_api_schema_task_event_task_event_proto_init() {
	if File_api_schema_task_event_task_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_schema_task_event_task_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskEventV1); i {
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
			RawDescriptor: file_api_schema_task_event_task_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_schema_task_event_task_event_proto_goTypes,
		DependencyIndexes: file_api_schema_task_event_task_event_proto_depIdxs,
		EnumInfos:         file_api_schema_task_event_task_event_proto_enumTypes,
		MessageInfos:      file_api_schema_task_event_task_event_proto_msgTypes,
	}.Build()
	File_api_schema_task_event_task_event_proto = out.File
	file_api_schema_task_event_task_event_proto_rawDesc = nil
	file_api_schema_task_event_task_event_proto_goTypes = nil
	file_api_schema_task_event_task_event_proto_depIdxs = nil
}