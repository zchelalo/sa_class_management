// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: class/service.proto

package classProto

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

type ClassData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Subject string `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Grade   string `protobuf:"bytes,4,opt,name=grade,proto3" json:"grade,omitempty"`
	Code    string `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *ClassData) Reset() {
	*x = ClassData{}
	mi := &file_class_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClassData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassData) ProtoMessage() {}

func (x *ClassData) ProtoReflect() protoreflect.Message {
	mi := &file_class_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassData.ProtoReflect.Descriptor instead.
func (*ClassData) Descriptor() ([]byte, []int) {
	return file_class_service_proto_rawDescGZIP(), []int{0}
}

func (x *ClassData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClassData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClassData) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *ClassData) GetGrade() string {
	if x != nil {
		return x.Grade
	}
	return ""
}

func (x *ClassData) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type ClassMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage    int32 `protobuf:"varint,2,opt,name=perPage,proto3" json:"perPage,omitempty"`
	Count      int32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	TotalCount int32 `protobuf:"varint,4,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
}

func (x *ClassMeta) Reset() {
	*x = ClassMeta{}
	mi := &file_class_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClassMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassMeta) ProtoMessage() {}

func (x *ClassMeta) ProtoReflect() protoreflect.Message {
	mi := &file_class_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassMeta.ProtoReflect.Descriptor instead.
func (*ClassMeta) Descriptor() ([]byte, []int) {
	return file_class_service_proto_rawDescGZIP(), []int{1}
}

func (x *ClassMeta) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ClassMeta) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *ClassMeta) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ClassMeta) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type ClassesWithMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Classes []*ClassData `protobuf:"bytes,1,rep,name=classes,proto3" json:"classes,omitempty"`
	Meta    *ClassMeta   `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *ClassesWithMeta) Reset() {
	*x = ClassesWithMeta{}
	mi := &file_class_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClassesWithMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassesWithMeta) ProtoMessage() {}

func (x *ClassesWithMeta) ProtoReflect() protoreflect.Message {
	mi := &file_class_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassesWithMeta.ProtoReflect.Descriptor instead.
func (*ClassesWithMeta) Descriptor() ([]byte, []int) {
	return file_class_service_proto_rawDescGZIP(), []int{2}
}

func (x *ClassesWithMeta) GetClasses() []*ClassData {
	if x != nil {
		return x.Classes
	}
	return nil
}

func (x *ClassesWithMeta) GetMeta() *ClassMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type ClassError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ClassError) Reset() {
	*x = ClassError{}
	mi := &file_class_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClassError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassError) ProtoMessage() {}

func (x *ClassError) ProtoReflect() protoreflect.Message {
	mi := &file_class_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassError.ProtoReflect.Descriptor instead.
func (*ClassError) Descriptor() ([]byte, []int) {
	return file_class_service_proto_rawDescGZIP(), []int{3}
}

func (x *ClassError) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ClassError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Subject string `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Grade   string `protobuf:"bytes,4,opt,name=grade,proto3" json:"grade,omitempty"`
}

func (x *CreateClassRequest) Reset() {
	*x = CreateClassRequest{}
	mi := &file_class_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClassRequest) ProtoMessage() {}

func (x *CreateClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_class_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClassRequest.ProtoReflect.Descriptor instead.
func (*CreateClassRequest) Descriptor() ([]byte, []int) {
	return file_class_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateClassRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateClassRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateClassRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *CreateClassRequest) GetGrade() string {
	if x != nil {
		return x.Grade
	}
	return ""
}

type CreateClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*CreateClassResponse_Class
	//	*CreateClassResponse_Error
	Result isCreateClassResponse_Result `protobuf_oneof:"result"`
}

func (x *CreateClassResponse) Reset() {
	*x = CreateClassResponse{}
	mi := &file_class_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClassResponse) ProtoMessage() {}

func (x *CreateClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_class_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClassResponse.ProtoReflect.Descriptor instead.
func (*CreateClassResponse) Descriptor() ([]byte, []int) {
	return file_class_service_proto_rawDescGZIP(), []int{5}
}

func (m *CreateClassResponse) GetResult() isCreateClassResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *CreateClassResponse) GetClass() *ClassData {
	if x, ok := x.GetResult().(*CreateClassResponse_Class); ok {
		return x.Class
	}
	return nil
}

func (x *CreateClassResponse) GetError() *ClassError {
	if x, ok := x.GetResult().(*CreateClassResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isCreateClassResponse_Result interface {
	isCreateClassResponse_Result()
}

type CreateClassResponse_Class struct {
	Class *ClassData `protobuf:"bytes,1,opt,name=class,proto3,oneof"`
}

type CreateClassResponse_Error struct {
	Error *ClassError `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*CreateClassResponse_Class) isCreateClassResponse_Result() {}

func (*CreateClassResponse_Error) isCreateClassResponse_Result() {}

type JoinClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Code   string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *JoinClassRequest) Reset() {
	*x = JoinClassRequest{}
	mi := &file_class_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinClassRequest) ProtoMessage() {}

func (x *JoinClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_class_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinClassRequest.ProtoReflect.Descriptor instead.
func (*JoinClassRequest) Descriptor() ([]byte, []int) {
	return file_class_service_proto_rawDescGZIP(), []int{6}
}

func (x *JoinClassRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *JoinClassRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type JoinClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*JoinClassResponse_Class
	//	*JoinClassResponse_Error
	Result isJoinClassResponse_Result `protobuf_oneof:"result"`
}

func (x *JoinClassResponse) Reset() {
	*x = JoinClassResponse{}
	mi := &file_class_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinClassResponse) ProtoMessage() {}

func (x *JoinClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_class_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinClassResponse.ProtoReflect.Descriptor instead.
func (*JoinClassResponse) Descriptor() ([]byte, []int) {
	return file_class_service_proto_rawDescGZIP(), []int{7}
}

func (m *JoinClassResponse) GetResult() isJoinClassResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *JoinClassResponse) GetClass() *ClassData {
	if x, ok := x.GetResult().(*JoinClassResponse_Class); ok {
		return x.Class
	}
	return nil
}

func (x *JoinClassResponse) GetError() *ClassError {
	if x, ok := x.GetResult().(*JoinClassResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isJoinClassResponse_Result interface {
	isJoinClassResponse_Result()
}

type JoinClassResponse_Class struct {
	Class *ClassData `protobuf:"bytes,1,opt,name=class,proto3,oneof"`
}

type JoinClassResponse_Error struct {
	Error *ClassError `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*JoinClassResponse_Class) isJoinClassResponse_Result() {}

func (*JoinClassResponse_Error) isJoinClassResponse_Result() {}

type ListClassesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Page   int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit  int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListClassesRequest) Reset() {
	*x = ListClassesRequest{}
	mi := &file_class_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListClassesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClassesRequest) ProtoMessage() {}

func (x *ListClassesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_class_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClassesRequest.ProtoReflect.Descriptor instead.
func (*ListClassesRequest) Descriptor() ([]byte, []int) {
	return file_class_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListClassesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListClassesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListClassesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListClassesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*ListClassesResponse_Data
	//	*ListClassesResponse_Error
	Result isListClassesResponse_Result `protobuf_oneof:"result"`
}

func (x *ListClassesResponse) Reset() {
	*x = ListClassesResponse{}
	mi := &file_class_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListClassesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClassesResponse) ProtoMessage() {}

func (x *ListClassesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_class_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClassesResponse.ProtoReflect.Descriptor instead.
func (*ListClassesResponse) Descriptor() ([]byte, []int) {
	return file_class_service_proto_rawDescGZIP(), []int{9}
}

func (m *ListClassesResponse) GetResult() isListClassesResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *ListClassesResponse) GetData() *ClassesWithMeta {
	if x, ok := x.GetResult().(*ListClassesResponse_Data); ok {
		return x.Data
	}
	return nil
}

func (x *ListClassesResponse) GetError() *ClassError {
	if x, ok := x.GetResult().(*ListClassesResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isListClassesResponse_Result interface {
	isListClassesResponse_Result()
}

type ListClassesResponse_Data struct {
	Data *ClassesWithMeta `protobuf:"bytes,1,opt,name=data,proto3,oneof"`
}

type ListClassesResponse_Error struct {
	Error *ClassError `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*ListClassesResponse_Data) isListClassesResponse_Result() {}

func (*ListClassesResponse_Error) isListClassesResponse_Result() {}

type GetClassCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ClassId string `protobuf:"bytes,2,opt,name=classId,proto3" json:"classId,omitempty"`
}

func (x *GetClassCodeRequest) Reset() {
	*x = GetClassCodeRequest{}
	mi := &file_class_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetClassCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassCodeRequest) ProtoMessage() {}

func (x *GetClassCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_class_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassCodeRequest.ProtoReflect.Descriptor instead.
func (*GetClassCodeRequest) Descriptor() ([]byte, []int) {
	return file_class_service_proto_rawDescGZIP(), []int{10}
}

func (x *GetClassCodeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetClassCodeRequest) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

type GetClassCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*GetClassCodeResponse_Code
	//	*GetClassCodeResponse_Error
	Result isGetClassCodeResponse_Result `protobuf_oneof:"result"`
}

func (x *GetClassCodeResponse) Reset() {
	*x = GetClassCodeResponse{}
	mi := &file_class_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetClassCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassCodeResponse) ProtoMessage() {}

func (x *GetClassCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_class_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassCodeResponse.ProtoReflect.Descriptor instead.
func (*GetClassCodeResponse) Descriptor() ([]byte, []int) {
	return file_class_service_proto_rawDescGZIP(), []int{11}
}

func (m *GetClassCodeResponse) GetResult() isGetClassCodeResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *GetClassCodeResponse) GetCode() string {
	if x, ok := x.GetResult().(*GetClassCodeResponse_Code); ok {
		return x.Code
	}
	return ""
}

func (x *GetClassCodeResponse) GetError() *ClassError {
	if x, ok := x.GetResult().(*GetClassCodeResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isGetClassCodeResponse_Result interface {
	isGetClassCodeResponse_Result()
}

type GetClassCodeResponse_Code struct {
	Code string `protobuf:"bytes,1,opt,name=code,proto3,oneof"`
}

type GetClassCodeResponse_Error struct {
	Error *ClassError `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*GetClassCodeResponse_Code) isGetClassCodeResponse_Result() {}

func (*GetClassCodeResponse_Error) isGetClassCodeResponse_Result() {}

var File_class_service_proto protoreflect.FileDescriptor

var file_class_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x6f, 0x0a, 0x09, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65,
	0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x57, 0x0a, 0x0f, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x24,
	0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x0a, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x70, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x22, 0x68, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x23, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3e, 0x0a, 0x10,
	0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x66, 0x0a, 0x11,
	0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x05,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x56, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x6c, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x4d,
	0x65, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x47, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0xf3, 0x01, 0x0a, 0x0c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x38, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x12, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x6a,
	0x6f, 0x69, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x11, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x4a, 0x6f,
	0x69, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x38, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x12, 0x13,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x67, 0x65, 0x74,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_class_service_proto_rawDescOnce sync.Once
	file_class_service_proto_rawDescData = file_class_service_proto_rawDesc
)

func file_class_service_proto_rawDescGZIP() []byte {
	file_class_service_proto_rawDescOnce.Do(func() {
		file_class_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_class_service_proto_rawDescData)
	})
	return file_class_service_proto_rawDescData
}

var file_class_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_class_service_proto_goTypes = []any{
	(*ClassData)(nil),            // 0: ClassData
	(*ClassMeta)(nil),            // 1: ClassMeta
	(*ClassesWithMeta)(nil),      // 2: ClassesWithMeta
	(*ClassError)(nil),           // 3: ClassError
	(*CreateClassRequest)(nil),   // 4: CreateClassRequest
	(*CreateClassResponse)(nil),  // 5: CreateClassResponse
	(*JoinClassRequest)(nil),     // 6: JoinClassRequest
	(*JoinClassResponse)(nil),    // 7: JoinClassResponse
	(*ListClassesRequest)(nil),   // 8: ListClassesRequest
	(*ListClassesResponse)(nil),  // 9: ListClassesResponse
	(*GetClassCodeRequest)(nil),  // 10: GetClassCodeRequest
	(*GetClassCodeResponse)(nil), // 11: GetClassCodeResponse
}
var file_class_service_proto_depIdxs = []int32{
	0,  // 0: ClassesWithMeta.classes:type_name -> ClassData
	1,  // 1: ClassesWithMeta.meta:type_name -> ClassMeta
	0,  // 2: CreateClassResponse.class:type_name -> ClassData
	3,  // 3: CreateClassResponse.error:type_name -> ClassError
	0,  // 4: JoinClassResponse.class:type_name -> ClassData
	3,  // 5: JoinClassResponse.error:type_name -> ClassError
	2,  // 6: ListClassesResponse.data:type_name -> ClassesWithMeta
	3,  // 7: ListClassesResponse.error:type_name -> ClassError
	3,  // 8: GetClassCodeResponse.error:type_name -> ClassError
	4,  // 9: ClassService.createClass:input_type -> CreateClassRequest
	6,  // 10: ClassService.joinClass:input_type -> JoinClassRequest
	8,  // 11: ClassService.listClasses:input_type -> ListClassesRequest
	10, // 12: ClassService.getClassCode:input_type -> GetClassCodeRequest
	5,  // 13: ClassService.createClass:output_type -> CreateClassResponse
	7,  // 14: ClassService.joinClass:output_type -> JoinClassResponse
	9,  // 15: ClassService.listClasses:output_type -> ListClassesResponse
	11, // 16: ClassService.getClassCode:output_type -> GetClassCodeResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_class_service_proto_init() }
func file_class_service_proto_init() {
	if File_class_service_proto != nil {
		return
	}
	file_class_service_proto_msgTypes[5].OneofWrappers = []any{
		(*CreateClassResponse_Class)(nil),
		(*CreateClassResponse_Error)(nil),
	}
	file_class_service_proto_msgTypes[7].OneofWrappers = []any{
		(*JoinClassResponse_Class)(nil),
		(*JoinClassResponse_Error)(nil),
	}
	file_class_service_proto_msgTypes[9].OneofWrappers = []any{
		(*ListClassesResponse_Data)(nil),
		(*ListClassesResponse_Error)(nil),
	}
	file_class_service_proto_msgTypes[11].OneofWrappers = []any{
		(*GetClassCodeResponse_Code)(nil),
		(*GetClassCodeResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_class_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_class_service_proto_goTypes,
		DependencyIndexes: file_class_service_proto_depIdxs,
		MessageInfos:      file_class_service_proto_msgTypes,
	}.Build()
	File_class_service_proto = out.File
	file_class_service_proto_rawDesc = nil
	file_class_service_proto_goTypes = nil
	file_class_service_proto_depIdxs = nil
}
