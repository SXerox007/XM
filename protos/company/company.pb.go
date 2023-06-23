// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: protos/company/company.proto

package company

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type CompanyType int32

const (
	CompanyType_CORPORATIONS        CompanyType = 0
	CompanyType_NONPROFIT           CompanyType = 1
	CompanyType_COOPERATIVE         CompanyType = 2
	CompanyType_SOLE_PROPRIETORSHIP CompanyType = 3
)

// Enum value maps for CompanyType.
var (
	CompanyType_name = map[int32]string{
		0: "CORPORATIONS",
		1: "NONPROFIT",
		2: "COOPERATIVE",
		3: "SOLE_PROPRIETORSHIP",
	}
	CompanyType_value = map[string]int32{
		"CORPORATIONS":        0,
		"NONPROFIT":           1,
		"COOPERATIVE":         2,
		"SOLE_PROPRIETORSHIP": 3,
	}
)

func (x CompanyType) Enum() *CompanyType {
	p := new(CompanyType)
	*p = x
	return p
}

func (x CompanyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompanyType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_company_company_proto_enumTypes[0].Descriptor()
}

func (CompanyType) Type() protoreflect.EnumType {
	return &file_protos_company_company_proto_enumTypes[0]
}

func (x CompanyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompanyType.Descriptor instead.
func (CompanyType) EnumDescriptor() ([]byte, []int) {
	return file_protos_company_company_proto_rawDescGZIP(), []int{0}
}

type Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Employees   int32                  `protobuf:"varint,4,opt,name=employees,proto3" json:"employees,omitempty"`
	Registered  bool                   `protobuf:"varint,5,opt,name=registered,proto3" json:"registered,omitempty"`
	Type        CompanyType            `protobuf:"varint,6,opt,name=type,proto3,enum=company.CompanyType" json:"type,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Company) Reset() {
	*x = Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_company_company_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company) ProtoMessage() {}

func (x *Company) ProtoReflect() protoreflect.Message {
	mi := &file_protos_company_company_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Company.ProtoReflect.Descriptor instead.
func (*Company) Descriptor() ([]byte, []int) {
	return file_protos_company_company_proto_rawDescGZIP(), []int{0}
}

func (x *Company) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Company) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Company) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Company) GetEmployees() int32 {
	if x != nil {
		return x.Employees
	}
	return 0
}

func (x *Company) GetRegistered() bool {
	if x != nil {
		return x.Registered
	}
	return false
}

func (x *Company) GetType() CompanyType {
	if x != nil {
		return x.Type
	}
	return CompanyType_CORPORATIONS
}

func (x *Company) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Company) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string      `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Employees   int32       `protobuf:"varint,3,opt,name=employees,proto3" json:"employees,omitempty"`
	Registered  bool        `protobuf:"varint,4,opt,name=registered,proto3" json:"registered,omitempty"`
	Type        CompanyType `protobuf:"varint,5,opt,name=type,proto3,enum=company.CompanyType" json:"type,omitempty"`
}

func (x *CreateCompanyRequest) Reset() {
	*x = CreateCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_company_company_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCompanyRequest) ProtoMessage() {}

func (x *CreateCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_company_company_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCompanyRequest.ProtoReflect.Descriptor instead.
func (*CreateCompanyRequest) Descriptor() ([]byte, []int) {
	return file_protos_company_company_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCompanyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCompanyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateCompanyRequest) GetEmployees() int32 {
	if x != nil {
		return x.Employees
	}
	return 0
}

func (x *CreateCompanyRequest) GetRegistered() bool {
	if x != nil {
		return x.Registered
	}
	return false
}

func (x *CreateCompanyRequest) GetType() CompanyType {
	if x != nil {
		return x.Type
	}
	return CompanyType_CORPORATIONS
}

type UpdateCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string      `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Employees   int32       `protobuf:"varint,4,opt,name=employees,proto3" json:"employees,omitempty"`
	Registered  bool        `protobuf:"varint,5,opt,name=registered,proto3" json:"registered,omitempty"`
	Type        CompanyType `protobuf:"varint,6,opt,name=type,proto3,enum=company.CompanyType" json:"type,omitempty"`
}

func (x *UpdateCompanyRequest) Reset() {
	*x = UpdateCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_company_company_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCompanyRequest) ProtoMessage() {}

func (x *UpdateCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_company_company_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCompanyRequest.ProtoReflect.Descriptor instead.
func (*UpdateCompanyRequest) Descriptor() ([]byte, []int) {
	return file_protos_company_company_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCompanyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCompanyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCompanyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateCompanyRequest) GetEmployees() int32 {
	if x != nil {
		return x.Employees
	}
	return 0
}

func (x *UpdateCompanyRequest) GetRegistered() bool {
	if x != nil {
		return x.Registered
	}
	return false
}

func (x *UpdateCompanyRequest) GetType() CompanyType {
	if x != nil {
		return x.Type
	}
	return CompanyType_CORPORATIONS
}

type DeleteCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCompanyRequest) Reset() {
	*x = DeleteCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_company_company_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCompanyRequest) ProtoMessage() {}

func (x *DeleteCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_company_company_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCompanyRequest.ProtoReflect.Descriptor instead.
func (*DeleteCompanyRequest) Descriptor() ([]byte, []int) {
	return file_protos_company_company_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteCompanyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCompanyRequest) Reset() {
	*x = GetCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_company_company_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyRequest) ProtoMessage() {}

func (x *GetCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_company_company_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyRequest.ProtoReflect.Descriptor instead.
func (*GetCompanyRequest) Descriptor() ([]byte, []int) {
	return file_protos_company_company_proto_rawDescGZIP(), []int{4}
}

func (x *GetCompanyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_protos_company_company_proto protoreflect.FileDescriptor

var file_protos_company_company_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xad, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xb4, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64,
	0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xc4, 0x01, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x58,
	0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x0c, 0x43, 0x4f, 0x52, 0x50, 0x4f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x4e, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x54, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x43, 0x4f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12,
	0x17, 0x0a, 0x13, 0x53, 0x4f, 0x4c, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x52, 0x49, 0x45, 0x54,
	0x4f, 0x52, 0x53, 0x48, 0x49, 0x50, 0x10, 0x03, 0x32, 0xf0, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1d, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x58, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a,
	0x01, 0x2a, 0x32, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12,
	0x5b, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a,
	0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x4d, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b,
	0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x58, 0x65, 0x72, 0x6f, 0x78,
	0x30, 0x30, 0x37, 0x2f, 0x58, 0x4d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_company_company_proto_rawDescOnce sync.Once
	file_protos_company_company_proto_rawDescData = file_protos_company_company_proto_rawDesc
)

func file_protos_company_company_proto_rawDescGZIP() []byte {
	file_protos_company_company_proto_rawDescOnce.Do(func() {
		file_protos_company_company_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_company_company_proto_rawDescData)
	})
	return file_protos_company_company_proto_rawDescData
}

var file_protos_company_company_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_company_company_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_company_company_proto_goTypes = []interface{}{
	(CompanyType)(0),              // 0: company.CompanyType
	(*Company)(nil),               // 1: company.Company
	(*CreateCompanyRequest)(nil),  // 2: company.CreateCompanyRequest
	(*UpdateCompanyRequest)(nil),  // 3: company.UpdateCompanyRequest
	(*DeleteCompanyRequest)(nil),  // 4: company.DeleteCompanyRequest
	(*GetCompanyRequest)(nil),     // 5: company.GetCompanyRequest
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_protos_company_company_proto_depIdxs = []int32{
	0, // 0: company.Company.type:type_name -> company.CompanyType
	6, // 1: company.Company.created_at:type_name -> google.protobuf.Timestamp
	6, // 2: company.Company.updated_at:type_name -> google.protobuf.Timestamp
	0, // 3: company.CreateCompanyRequest.type:type_name -> company.CompanyType
	0, // 4: company.UpdateCompanyRequest.type:type_name -> company.CompanyType
	2, // 5: company.CompanyService.CreateCompany:input_type -> company.CreateCompanyRequest
	3, // 6: company.CompanyService.UpdateCompany:input_type -> company.UpdateCompanyRequest
	4, // 7: company.CompanyService.DeleteCompany:input_type -> company.DeleteCompanyRequest
	5, // 8: company.CompanyService.GetCompany:input_type -> company.GetCompanyRequest
	1, // 9: company.CompanyService.CreateCompany:output_type -> company.Company
	1, // 10: company.CompanyService.UpdateCompany:output_type -> company.Company
	7, // 11: company.CompanyService.DeleteCompany:output_type -> google.protobuf.Empty
	1, // 12: company.CompanyService.GetCompany:output_type -> company.Company
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protos_company_company_proto_init() }
func file_protos_company_company_proto_init() {
	if File_protos_company_company_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_company_company_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Company); i {
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
		file_protos_company_company_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCompanyRequest); i {
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
		file_protos_company_company_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCompanyRequest); i {
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
		file_protos_company_company_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCompanyRequest); i {
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
		file_protos_company_company_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompanyRequest); i {
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
			RawDescriptor: file_protos_company_company_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_company_company_proto_goTypes,
		DependencyIndexes: file_protos_company_company_proto_depIdxs,
		EnumInfos:         file_protos_company_company_proto_enumTypes,
		MessageInfos:      file_protos_company_company_proto_msgTypes,
	}.Build()
	File_protos_company_company_proto = out.File
	file_protos_company_company_proto_rawDesc = nil
	file_protos_company_company_proto_goTypes = nil
	file_protos_company_company_proto_depIdxs = nil
}