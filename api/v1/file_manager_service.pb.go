// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: api/v1/file_manager_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Bytes     int64  `protobuf:"varint,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Filename  string `protobuf:"bytes,4,opt,name=filename,proto3" json:"filename,omitempty"`
	Object    string `protobuf:"bytes,5,opt,name=object,proto3" json:"object,omitempty"`
	Purpose   string `protobuf:"bytes,6,opt,name=purpose,proto3" json:"purpose,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_file_manager_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_file_manager_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_api_v1_file_manager_service_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *File) GetBytes() int64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

func (x *File) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *File) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *File) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *File) GetPurpose() string {
	if x != nil {
		return x.Purpose
	}
	return ""
}

type ListFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Purpose string `protobuf:"bytes,1,opt,name=purpose,proto3" json:"purpose,omitempty"`
}

func (x *ListFilesRequest) Reset() {
	*x = ListFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_file_manager_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFilesRequest) ProtoMessage() {}

func (x *ListFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_file_manager_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFilesRequest.ProtoReflect.Descriptor instead.
func (*ListFilesRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_file_manager_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListFilesRequest) GetPurpose() string {
	if x != nil {
		return x.Purpose
	}
	return ""
}

type ListFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object string  `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Data   []*File `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListFilesResponse) Reset() {
	*x = ListFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_file_manager_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFilesResponse) ProtoMessage() {}

func (x *ListFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_file_manager_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFilesResponse.ProtoReflect.Descriptor instead.
func (*ListFilesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_file_manager_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListFilesResponse) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *ListFilesResponse) GetData() []*File {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFileRequest) Reset() {
	*x = GetFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_file_manager_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileRequest) ProtoMessage() {}

func (x *GetFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_file_manager_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileRequest.ProtoReflect.Descriptor instead.
func (*GetFileRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_file_manager_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetFileRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteFileRequest) Reset() {
	*x = DeleteFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_file_manager_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileRequest) ProtoMessage() {}

func (x *DeleteFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_file_manager_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileRequest.ProtoReflect.Descriptor instead.
func (*DeleteFileRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_file_manager_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteFileRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Object  string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	Deleted bool   `protobuf:"varint,3,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeleteFileResponse) Reset() {
	*x = DeleteFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_file_manager_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileResponse) ProtoMessage() {}

func (x *DeleteFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_file_manager_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileResponse.ProtoReflect.Descriptor instead.
func (*DeleteFileResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_file_manager_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteFileResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteFileResponse) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *DeleteFileResponse) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type CreateFileFromObjectPathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The object path is the path to the object in the object storage.
	ObjectPath string `protobuf:"bytes,1,opt,name=object_path,json=objectPath,proto3" json:"object_path,omitempty"`
	Purpose    string `protobuf:"bytes,2,opt,name=purpose,proto3" json:"purpose,omitempty"`
}

func (x *CreateFileFromObjectPathRequest) Reset() {
	*x = CreateFileFromObjectPathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_file_manager_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileFromObjectPathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileFromObjectPathRequest) ProtoMessage() {}

func (x *CreateFileFromObjectPathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_file_manager_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileFromObjectPathRequest.ProtoReflect.Descriptor instead.
func (*CreateFileFromObjectPathRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_file_manager_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreateFileFromObjectPathRequest) GetObjectPath() string {
	if x != nil {
		return x.ObjectPath
	}
	return ""
}

func (x *CreateFileFromObjectPathRequest) GetPurpose() string {
	if x != nil {
		return x.Purpose
	}
	return ""
}

type GetFilePathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFilePathRequest) Reset() {
	*x = GetFilePathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_file_manager_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFilePathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilePathRequest) ProtoMessage() {}

func (x *GetFilePathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_file_manager_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilePathRequest.ProtoReflect.Descriptor instead.
func (*GetFilePathRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_file_manager_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetFilePathRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetFilePathResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *GetFilePathResponse) Reset() {
	*x = GetFilePathResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_file_manager_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFilePathResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilePathResponse) ProtoMessage() {}

func (x *GetFilePathResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_file_manager_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilePathResponse.ProtoReflect.Descriptor instead.
func (*GetFilePathResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_file_manager_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetFilePathResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GetFilePathResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

var File_api_v1_file_manager_service_proto protoreflect.FileDescriptor

var file_api_v1_file_manager_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6c, 0x6c, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a,
	0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x22, 0x60, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x6c, 0x6c, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x56, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x5c, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50,
	0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x75,
	0x72, 0x70, 0x6f, 0x73, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x32, 0xa1, 0x04, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x79, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x12, 0x2b, 0x2e, 0x6c, 0x6c, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x6c, 0x6c, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x6d,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x29, 0x2e, 0x6c, 0x6c, 0x6d, 0x61,
	0x72, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6c, 0x6c, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x72,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x81, 0x01,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2c, 0x2e, 0x6c,
	0x6c, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6c, 0x6c, 0x6d,
	0x61, 0x72, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x2a, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0xa2, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x3a,
	0x2e, 0x6c, 0x6c, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50,
	0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6c, 0x6c, 0x6d,
	0x61, 0x72, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x3a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x50, 0x61, 0x74, 0x68, 0x32, 0x84, 0x01, 0x0a, 0x12, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2d, 0x2e, 0x6c,
	0x6c, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6c, 0x6c,
	0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x86, 0x01,
	0x0a, 0x14, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2d, 0x2e, 0x6c, 0x6c, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65,
	0x72, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6c, 0x6c, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x72,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6c, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x72, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_file_manager_service_proto_rawDescOnce sync.Once
	file_api_v1_file_manager_service_proto_rawDescData = file_api_v1_file_manager_service_proto_rawDesc
)

func file_api_v1_file_manager_service_proto_rawDescGZIP() []byte {
	file_api_v1_file_manager_service_proto_rawDescOnce.Do(func() {
		file_api_v1_file_manager_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_file_manager_service_proto_rawDescData)
	})
	return file_api_v1_file_manager_service_proto_rawDescData
}

var file_api_v1_file_manager_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_v1_file_manager_service_proto_goTypes = []interface{}{
	(*File)(nil),                            // 0: llmariner.files.server.v1.File
	(*ListFilesRequest)(nil),                // 1: llmariner.files.server.v1.ListFilesRequest
	(*ListFilesResponse)(nil),               // 2: llmariner.files.server.v1.ListFilesResponse
	(*GetFileRequest)(nil),                  // 3: llmariner.files.server.v1.GetFileRequest
	(*DeleteFileRequest)(nil),               // 4: llmariner.files.server.v1.DeleteFileRequest
	(*DeleteFileResponse)(nil),              // 5: llmariner.files.server.v1.DeleteFileResponse
	(*CreateFileFromObjectPathRequest)(nil), // 6: llmariner.files.server.v1.CreateFileFromObjectPathRequest
	(*GetFilePathRequest)(nil),              // 7: llmariner.files.server.v1.GetFilePathRequest
	(*GetFilePathResponse)(nil),             // 8: llmariner.files.server.v1.GetFilePathResponse
}
var file_api_v1_file_manager_service_proto_depIdxs = []int32{
	0, // 0: llmariner.files.server.v1.ListFilesResponse.data:type_name -> llmariner.files.server.v1.File
	1, // 1: llmariner.files.server.v1.FilesService.ListFiles:input_type -> llmariner.files.server.v1.ListFilesRequest
	3, // 2: llmariner.files.server.v1.FilesService.GetFile:input_type -> llmariner.files.server.v1.GetFileRequest
	4, // 3: llmariner.files.server.v1.FilesService.DeleteFile:input_type -> llmariner.files.server.v1.DeleteFileRequest
	6, // 4: llmariner.files.server.v1.FilesService.CreateFileFromObjectPath:input_type -> llmariner.files.server.v1.CreateFileFromObjectPathRequest
	7, // 5: llmariner.files.server.v1.FilesWorkerService.GetFilePath:input_type -> llmariner.files.server.v1.GetFilePathRequest
	7, // 6: llmariner.files.server.v1.FilesInternalService.GetFilePath:input_type -> llmariner.files.server.v1.GetFilePathRequest
	2, // 7: llmariner.files.server.v1.FilesService.ListFiles:output_type -> llmariner.files.server.v1.ListFilesResponse
	0, // 8: llmariner.files.server.v1.FilesService.GetFile:output_type -> llmariner.files.server.v1.File
	5, // 9: llmariner.files.server.v1.FilesService.DeleteFile:output_type -> llmariner.files.server.v1.DeleteFileResponse
	0, // 10: llmariner.files.server.v1.FilesService.CreateFileFromObjectPath:output_type -> llmariner.files.server.v1.File
	8, // 11: llmariner.files.server.v1.FilesWorkerService.GetFilePath:output_type -> llmariner.files.server.v1.GetFilePathResponse
	8, // 12: llmariner.files.server.v1.FilesInternalService.GetFilePath:output_type -> llmariner.files.server.v1.GetFilePathResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_file_manager_service_proto_init() }
func file_api_v1_file_manager_service_proto_init() {
	if File_api_v1_file_manager_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_file_manager_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_api_v1_file_manager_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFilesRequest); i {
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
		file_api_v1_file_manager_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFilesResponse); i {
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
		file_api_v1_file_manager_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileRequest); i {
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
		file_api_v1_file_manager_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileRequest); i {
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
		file_api_v1_file_manager_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileResponse); i {
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
		file_api_v1_file_manager_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileFromObjectPathRequest); i {
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
		file_api_v1_file_manager_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFilePathRequest); i {
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
		file_api_v1_file_manager_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFilePathResponse); i {
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
			RawDescriptor: file_api_v1_file_manager_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_api_v1_file_manager_service_proto_goTypes,
		DependencyIndexes: file_api_v1_file_manager_service_proto_depIdxs,
		MessageInfos:      file_api_v1_file_manager_service_proto_msgTypes,
	}.Build()
	File_api_v1_file_manager_service_proto = out.File
	file_api_v1_file_manager_service_proto_rawDesc = nil
	file_api_v1_file_manager_service_proto_goTypes = nil
	file_api_v1_file_manager_service_proto_depIdxs = nil
}
