// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: event_service.proto

package eventpb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId             string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title              string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description        string                 `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	DateStart          *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=date_start,json=dateStart,proto3" json:"date_start,omitempty"`
	Duration           *durationpb.Duration   `protobuf:"bytes,6,opt,name=duration,proto3" json:"duration,omitempty"`
	NotificationPeriod *durationpb.Duration   `protobuf:"bytes,7,opt,name=notification_period,json=notificationPeriod,proto3" json:"notification_period,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Event) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event) GetDateStart() *timestamppb.Timestamp {
	if x != nil {
		return x.DateStart
	}
	return nil
}

func (x *Event) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *Event) GetNotificationPeriod() *durationpb.Duration {
	if x != nil {
		return x.NotificationPeriod
	}
	return nil
}

type Events struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *Events) Reset() {
	*x = Events{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Events) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Events) ProtoMessage() {}

func (x *Events) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Events.ProtoReflect.Descriptor instead.
func (*Events) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{1}
}

func (x *Events) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type AddEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId             string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title              string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description        string                 `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	DateStart          *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date_start,json=dateStart,proto3" json:"date_start,omitempty"`
	Duration           *durationpb.Duration   `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	NotificationPeriod *durationpb.Duration   `protobuf:"bytes,6,opt,name=notification_period,json=notificationPeriod,proto3" json:"notification_period,omitempty"`
}

func (x *AddEventRequest) Reset() {
	*x = AddEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEventRequest) ProtoMessage() {}

func (x *AddEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEventRequest.ProtoReflect.Descriptor instead.
func (*AddEventRequest) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddEventRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddEventRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddEventRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddEventRequest) GetDateStart() *timestamppb.Timestamp {
	if x != nil {
		return x.DateStart
	}
	return nil
}

func (x *AddEventRequest) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *AddEventRequest) GetNotificationPeriod() *durationpb.Duration {
	if x != nil {
		return x.NotificationPeriod
	}
	return nil
}

type AddEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddEventResponse) Reset() {
	*x = AddEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEventResponse) ProtoMessage() {}

func (x *AddEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEventResponse.ProtoReflect.Descriptor instead.
func (*AddEventResponse) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddEventResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteEventByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteEventByIDRequest) Reset() {
	*x = DeleteEventByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventByIDRequest) ProtoMessage() {}

func (x *DeleteEventByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventByIDRequest.ProtoReflect.Descriptor instead.
func (*DeleteEventByIDRequest) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteEventByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateEventByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateEventByIDResponse) Reset() {
	*x = UpdateEventByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventByIDResponse) ProtoMessage() {}

func (x *UpdateEventByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventByIDResponse.ProtoReflect.Descriptor instead.
func (*UpdateEventByIDResponse) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{5}
}

type DeleteEventByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteEventByIDResponse) Reset() {
	*x = DeleteEventByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventByIDResponse) ProtoMessage() {}

func (x *DeleteEventByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventByIDResponse.ProtoReflect.Descriptor instead.
func (*DeleteEventByIDResponse) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{6}
}

type GetEventByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetEventByIDRequest) Reset() {
	*x = GetEventByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventByIDRequest) ProtoMessage() {}

func (x *GetEventByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventByIDRequest.ProtoReflect.Descriptor instead.
func (*GetEventByIDRequest) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetEventByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeriodStart *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=periodStart,proto3,oneof" json:"periodStart,omitempty"`
	PeriodEnd   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=periodEnd,proto3,oneof" json:"periodEnd,omitempty"`
}

func (x *GetEventsRequest) Reset() {
	*x = GetEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsRequest) ProtoMessage() {}

func (x *GetEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsRequest.ProtoReflect.Descriptor instead.
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetEventsRequest) GetPeriodStart() *timestamppb.Timestamp {
	if x != nil {
		return x.PeriodStart
	}
	return nil
}

func (x *GetEventsRequest) GetPeriodEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.PeriodEnd
	}
	return nil
}

var File_event_service_proto protoreflect.FileDescriptor

var file_event_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x02, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x13, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x12, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x22, 0x2e, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x24,
	0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0xa0, 0x02, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x13, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x12, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x22, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xb2, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x09, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x09, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x32, 0xfa, 0x05, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x08, 0x41, 0x64, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64,
	0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x92, 0x41,
	0x33, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0c, 0x41, 0x64, 0x64, 0x20, 0x61,
	0x6e, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x41, 0x64, 0x64, 0x20, 0x61, 0x6e, 0x20,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x12, 0xab, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x59, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a,
	0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x3b, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x1a, 0x20, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x12, 0x98, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x57, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x1a, 0x10, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01,
	0x2a, 0x92, 0x41, 0x39, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x20, 0x6f,
	0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x12, 0x96, 0x01,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1a,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x5c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x92, 0x41, 0x41, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x47,
	0x65, 0x74, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x20, 0x62, 0x79, 0x20, 0x69,
	0x64, 0x1a, 0x23, 0x47, 0x65, 0x74, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x20,
	0x62, 0x79, 0x20, 0x69, 0x64, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x12, 0x7c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x47, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x92, 0x41, 0x31, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0a, 0x47, 0x65, 0x74,
	0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1b, 0x47, 0x65, 0x74, 0x20, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x42, 0x96, 0x01, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x92, 0x41, 0x86, 0x01, 0x12, 0x05, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01,
	0x01, 0x72, 0x7a, 0x0a, 0x27, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x20, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x4f, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x5a, 0x68, 0x75, 0x72, 0x37, 0x37, 0x31, 0x2f, 0x6f, 0x74, 0x75, 0x73, 0x2d, 0x67,
	0x6f, 0x2d, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x77, 0x31, 0x32, 0x5f, 0x31, 0x33, 0x5f, 0x31,
	0x34, 0x5f, 0x31, 0x35, 0x5f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_service_proto_rawDescOnce sync.Once
	file_event_service_proto_rawDescData = file_event_service_proto_rawDesc
)

func file_event_service_proto_rawDescGZIP() []byte {
	file_event_service_proto_rawDescOnce.Do(func() {
		file_event_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_service_proto_rawDescData)
	})
	return file_event_service_proto_rawDescData
}

var file_event_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_event_service_proto_goTypes = []interface{}{
	(*Event)(nil),                   // 0: event.Event
	(*Events)(nil),                  // 1: event.Events
	(*AddEventRequest)(nil),         // 2: event.AddEventRequest
	(*AddEventResponse)(nil),        // 3: event.AddEventResponse
	(*DeleteEventByIDRequest)(nil),  // 4: event.DeleteEventByIDRequest
	(*UpdateEventByIDResponse)(nil), // 5: event.UpdateEventByIDResponse
	(*DeleteEventByIDResponse)(nil), // 6: event.DeleteEventByIDResponse
	(*GetEventByIDRequest)(nil),     // 7: event.GetEventByIDRequest
	(*GetEventsRequest)(nil),        // 8: event.GetEventsRequest
	(*timestamppb.Timestamp)(nil),   // 9: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),     // 10: google.protobuf.Duration
}
var file_event_service_proto_depIdxs = []int32{
	9,  // 0: event.Event.date_start:type_name -> google.protobuf.Timestamp
	10, // 1: event.Event.duration:type_name -> google.protobuf.Duration
	10, // 2: event.Event.notification_period:type_name -> google.protobuf.Duration
	0,  // 3: event.Events.events:type_name -> event.Event
	9,  // 4: event.AddEventRequest.date_start:type_name -> google.protobuf.Timestamp
	10, // 5: event.AddEventRequest.duration:type_name -> google.protobuf.Duration
	10, // 6: event.AddEventRequest.notification_period:type_name -> google.protobuf.Duration
	9,  // 7: event.GetEventsRequest.periodStart:type_name -> google.protobuf.Timestamp
	9,  // 8: event.GetEventsRequest.periodEnd:type_name -> google.protobuf.Timestamp
	2,  // 9: event.EventService.AddEvent:input_type -> event.AddEventRequest
	4,  // 10: event.EventService.DeleteEventByID:input_type -> event.DeleteEventByIDRequest
	0,  // 11: event.EventService.UpdateEventByID:input_type -> event.Event
	7,  // 12: event.EventService.GetEventByID:input_type -> event.GetEventByIDRequest
	8,  // 13: event.EventService.GetEvents:input_type -> event.GetEventsRequest
	3,  // 14: event.EventService.AddEvent:output_type -> event.AddEventResponse
	6,  // 15: event.EventService.DeleteEventByID:output_type -> event.DeleteEventByIDResponse
	5,  // 16: event.EventService.UpdateEventByID:output_type -> event.UpdateEventByIDResponse
	0,  // 17: event.EventService.GetEventByID:output_type -> event.Event
	1,  // 18: event.EventService.GetEvents:output_type -> event.Events
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_event_service_proto_init() }
func file_event_service_proto_init() {
	if File_event_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_event_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Events); i {
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
		file_event_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEventRequest); i {
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
		file_event_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEventResponse); i {
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
		file_event_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventByIDRequest); i {
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
		file_event_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventByIDResponse); i {
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
		file_event_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventByIDResponse); i {
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
		file_event_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventByIDRequest); i {
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
		file_event_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsRequest); i {
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
	file_event_service_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_event_service_proto_goTypes,
		DependencyIndexes: file_event_service_proto_depIdxs,
		MessageInfos:      file_event_service_proto_msgTypes,
	}.Build()
	File_event_service_proto = out.File
	file_event_service_proto_rawDesc = nil
	file_event_service_proto_goTypes = nil
	file_event_service_proto_depIdxs = nil
}
