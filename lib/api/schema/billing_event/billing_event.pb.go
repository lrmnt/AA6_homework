// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.2
// source: api/schema/billing_event/billing_event.proto

package billing_event

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
	Event_EVENT_UNKNOWN                                   Event = 0
	Event_EVENT_PAYED_TO_USER                             Event = 1
	Event_EVENT_USER_BALANCE_BELOW_ZERO_AT_THE_END_OF_DAY Event = 2
)

// Enum value maps for Event.
var (
	Event_name = map[int32]string{
		0: "EVENT_UNKNOWN",
		1: "EVENT_PAYED_TO_USER",
		2: "EVENT_USER_BALANCE_BELOW_ZERO_AT_THE_END_OF_DAY",
	}
	Event_value = map[string]int32{
		"EVENT_UNKNOWN":       0,
		"EVENT_PAYED_TO_USER": 1,
		"EVENT_USER_BALANCE_BELOW_ZERO_AT_THE_END_OF_DAY": 2,
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
	return file_api_schema_billing_event_billing_event_proto_enumTypes[0].Descriptor()
}

func (Event) Type() protoreflect.EnumType {
	return &file_api_schema_billing_event_billing_event_proto_enumTypes[0]
}

func (x Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event.Descriptor instead.
func (Event) EnumDescriptor() ([]byte, []int) {
	return file_api_schema_billing_event_billing_event_proto_rawDescGZIP(), []int{0}
}

type BillingEventV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event     Event  `protobuf:"varint,1,opt,name=event,proto3,enum=billing_event.Event" json:"event,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	EventId   string `protobuf:"bytes,3,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	UserId    string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount    int64  `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *BillingEventV1) Reset() {
	*x = BillingEventV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_schema_billing_event_billing_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingEventV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingEventV1) ProtoMessage() {}

func (x *BillingEventV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_schema_billing_event_billing_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingEventV1.ProtoReflect.Descriptor instead.
func (*BillingEventV1) Descriptor() ([]byte, []int) {
	return file_api_schema_billing_event_billing_event_proto_rawDescGZIP(), []int{0}
}

func (x *BillingEventV1) GetEvent() Event {
	if x != nil {
		return x.Event
	}
	return Event_EVENT_UNKNOWN
}

func (x *BillingEventV1) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *BillingEventV1) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *BillingEventV1) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BillingEventV1) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_api_schema_billing_event_billing_event_proto protoreflect.FileDescriptor

var file_api_schema_billing_event_billing_event_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xa6, 0x01,
	0x0a, 0x0e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x31,
	0x12, 0x2a, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x68, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x11, 0x0a, 0x0d, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x41, 0x59, 0x45,
	0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12, 0x33, 0x0a, 0x2f, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x42, 0x41, 0x4c, 0x41, 0x4e, 0x43,
	0x45, 0x5f, 0x42, 0x45, 0x4c, 0x4f, 0x57, 0x5f, 0x5a, 0x45, 0x52, 0x4f, 0x5f, 0x41, 0x54, 0x5f,
	0x54, 0x48, 0x45, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x4f, 0x46, 0x5f, 0x44, 0x41, 0x59, 0x10, 0x02,
	0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x72, 0x6d, 0x6e, 0x74, 0x2f, 0x41, 0x41, 0x36, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_schema_billing_event_billing_event_proto_rawDescOnce sync.Once
	file_api_schema_billing_event_billing_event_proto_rawDescData = file_api_schema_billing_event_billing_event_proto_rawDesc
)

func file_api_schema_billing_event_billing_event_proto_rawDescGZIP() []byte {
	file_api_schema_billing_event_billing_event_proto_rawDescOnce.Do(func() {
		file_api_schema_billing_event_billing_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_schema_billing_event_billing_event_proto_rawDescData)
	})
	return file_api_schema_billing_event_billing_event_proto_rawDescData
}

var file_api_schema_billing_event_billing_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_schema_billing_event_billing_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_schema_billing_event_billing_event_proto_goTypes = []interface{}{
	(Event)(0),             // 0: billing_event.Event
	(*BillingEventV1)(nil), // 1: billing_event.BillingEventV1
}
var file_api_schema_billing_event_billing_event_proto_depIdxs = []int32{
	0, // 0: billing_event.BillingEventV1.event:type_name -> billing_event.Event
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_schema_billing_event_billing_event_proto_init() }
func file_api_schema_billing_event_billing_event_proto_init() {
	if File_api_schema_billing_event_billing_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_schema_billing_event_billing_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingEventV1); i {
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
			RawDescriptor: file_api_schema_billing_event_billing_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_schema_billing_event_billing_event_proto_goTypes,
		DependencyIndexes: file_api_schema_billing_event_billing_event_proto_depIdxs,
		EnumInfos:         file_api_schema_billing_event_billing_event_proto_enumTypes,
		MessageInfos:      file_api_schema_billing_event_billing_event_proto_msgTypes,
	}.Build()
	File_api_schema_billing_event_billing_event_proto = out.File
	file_api_schema_billing_event_billing_event_proto_rawDesc = nil
	file_api_schema_billing_event_billing_event_proto_goTypes = nil
	file_api_schema_billing_event_billing_event_proto_depIdxs = nil
}