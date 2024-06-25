// Copyright (c) 2024 IndyKite
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Trusted Data Access Service Description.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: indykite/tda/v1beta1/trusted_data_access_api.proto

package tdav1beta1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	v1beta1 "github.com/indykite/indykite-sdk-go/gen/indykite/knowledge/objects/v1beta1"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListConsentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the user.
	User *v1beta1.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// ApplicationId is the id of the application that the consent is for.
	ApplicationId string `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
}

func (x *ListConsentsRequest) Reset() {
	*x = ListConsentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConsentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConsentsRequest) ProtoMessage() {}

func (x *ListConsentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConsentsRequest.ProtoReflect.Descriptor instead.
func (*ListConsentsRequest) Descriptor() ([]byte, []int) {
	return file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDescGZIP(), []int{0}
}

func (x *ListConsentsRequest) GetUser() *v1beta1.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ListConsentsRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

type ListConsentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consents []*Consent `protobuf:"bytes,1,rep,name=consents,proto3" json:"consents,omitempty"`
}

func (x *ListConsentsResponse) Reset() {
	*x = ListConsentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConsentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConsentsResponse) ProtoMessage() {}

func (x *ListConsentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConsentsResponse.ProtoReflect.Descriptor instead.
func (*ListConsentsResponse) Descriptor() ([]byte, []int) {
	return file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDescGZIP(), []int{1}
}

func (x *ListConsentsResponse) GetConsents() []*Consent {
	if x != nil {
		return x.Consents
	}
	return nil
}

type DataAccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the consent
	ConsentId string `protobuf:"bytes,1,opt,name=consent_id,json=consentId,proto3" json:"consent_id,omitempty"`
	// ApplicationId is the id of the application that the consent is for.
	ApplicationId string `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	// Optional, unique identifier for the user.
	User *v1beta1.User `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *DataAccessRequest) Reset() {
	*x = DataAccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataAccessRequest) ProtoMessage() {}

func (x *DataAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataAccessRequest.ProtoReflect.Descriptor instead.
func (*DataAccessRequest) Descriptor() ([]byte, []int) {
	return file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDescGZIP(), []int{2}
}

func (x *DataAccessRequest) GetConsentId() string {
	if x != nil {
		return x.ConsentId
	}
	return ""
}

func (x *DataAccessRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *DataAccessRequest) GetUser() *v1beta1.User {
	if x != nil {
		return x.User
	}
	return nil
}

type DataAccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*v1beta1.Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *DataAccessResponse) Reset() {
	*x = DataAccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataAccessResponse) ProtoMessage() {}

func (x *DataAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataAccessResponse.ProtoReflect.Descriptor instead.
func (*DataAccessResponse) Descriptor() ([]byte, []int) {
	return file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDescGZIP(), []int{3}
}

func (x *DataAccessResponse) GetNodes() []*v1beta1.Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

// The request to grant consent
type GrantConsentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the user
	User *v1beta1.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Unique identifier for the consent
	ConsentId string `protobuf:"bytes,2,opt,name=consent_id,json=consentId,proto3" json:"consent_id,omitempty"`
	// Optional: Specifies the duration in second that the consent remains valid, ranging from 1 day to 2 years.
	// Should be lower or equal to the consent's configuration.
	ValidityPeriod uint64 `protobuf:"varint,4,opt,name=validity_period,json=validityPeriod,proto3" json:"validity_period,omitempty"`
}

func (x *GrantConsentRequest) Reset() {
	*x = GrantConsentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantConsentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantConsentRequest) ProtoMessage() {}

func (x *GrantConsentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantConsentRequest.ProtoReflect.Descriptor instead.
func (*GrantConsentRequest) Descriptor() ([]byte, []int) {
	return file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDescGZIP(), []int{4}
}

func (x *GrantConsentRequest) GetUser() *v1beta1.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *GrantConsentRequest) GetConsentId() string {
	if x != nil {
		return x.ConsentId
	}
	return ""
}

func (x *GrantConsentRequest) GetValidityPeriod() uint64 {
	if x != nil {
		return x.ValidityPeriod
	}
	return 0
}

type GrantConsentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertiesGrantedCount uint64 `protobuf:"varint,1,opt,name=properties_granted_count,json=propertiesGrantedCount,proto3" json:"properties_granted_count,omitempty"`
}

func (x *GrantConsentResponse) Reset() {
	*x = GrantConsentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantConsentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantConsentResponse) ProtoMessage() {}

func (x *GrantConsentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantConsentResponse.ProtoReflect.Descriptor instead.
func (*GrantConsentResponse) Descriptor() ([]byte, []int) {
	return file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDescGZIP(), []int{5}
}

func (x *GrantConsentResponse) GetPropertiesGrantedCount() uint64 {
	if x != nil {
		return x.PropertiesGrantedCount
	}
	return 0
}

type RevokeConsentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the user
	User *v1beta1.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Unique identifier for the consent
	ConsentId string `protobuf:"bytes,2,opt,name=consent_id,json=consentId,proto3" json:"consent_id,omitempty"`
}

func (x *RevokeConsentRequest) Reset() {
	*x = RevokeConsentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeConsentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeConsentRequest) ProtoMessage() {}

func (x *RevokeConsentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeConsentRequest.ProtoReflect.Descriptor instead.
func (*RevokeConsentRequest) Descriptor() ([]byte, []int) {
	return file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDescGZIP(), []int{6}
}

func (x *RevokeConsentRequest) GetUser() *v1beta1.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *RevokeConsentRequest) GetConsentId() string {
	if x != nil {
		return x.ConsentId
	}
	return ""
}

type RevokeConsentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RevokeConsentResponse) Reset() {
	*x = RevokeConsentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeConsentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeConsentResponse) ProtoMessage() {}

func (x *RevokeConsentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeConsentResponse.ProtoReflect.Descriptor instead.
func (*RevokeConsentResponse) Descriptor() ([]byte, []int) {
	return file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDescGZIP(), []int{7}
}

var File_indykite_tda_v1beta1_trusted_data_access_api_proto protoreflect.FileDescriptor

var file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDesc = []byte{
	0x0a, 0x32, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x74, 0x64, 0x61, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x74,
	0x64, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x2c, 0x69, 0x6e, 0x64, 0x79,
	0x6b, 0x69, 0x74, 0x65, 0x2f, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x69,
	0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69,
	0x74, 0x65, 0x2f, 0x74, 0x64, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e, 0x64, 0x79,
	0x6b, 0x69, 0x74, 0x65, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xfa, 0x42, 0x24,
	0x72, 0x22, 0x10, 0x16, 0x18, 0xfe, 0x01, 0x32, 0x18, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d,
	0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5f, 0x3a, 0x5d, 0x7b, 0x32, 0x32, 0x2c, 0x32, 0x35, 0x34, 0x7d,
	0x24, 0xd0, 0x01, 0x01, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x74, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xf0, 0x01, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x24, 0xfa, 0x42, 0x21, 0x72, 0x1f, 0x10, 0x16, 0x18, 0xfe, 0x01, 0x32, 0x18, 0x5e, 0x5b,
	0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5f, 0x3a, 0x5d, 0x7b, 0x32, 0x32,
	0x2c, 0x32, 0x35, 0x34, 0x7d, 0x24, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x4e, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xfa, 0x42, 0x24, 0x72, 0x22,
	0x10, 0x16, 0x18, 0xfe, 0x01, 0x32, 0x18, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30,
	0x2d, 0x39, 0x2d, 0x5f, 0x3a, 0x5d, 0x7b, 0x32, 0x32, 0x2c, 0x32, 0x35, 0x34, 0x7d, 0x24, 0xd0,
	0x01, 0x01, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x46, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x00, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x54, 0x0a, 0x12, 0x44, 0x61, 0x74,
	0x61, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22,
	0xcb, 0x01, 0x0a, 0x13, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65,
	0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x43, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x24, 0xfa, 0x42, 0x21, 0x72, 0x1f, 0x10, 0x16, 0x18, 0xfe, 0x01, 0x32,
	0x18, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5f, 0x3a, 0x5d,
	0x7b, 0x32, 0x32, 0x2c, 0x32, 0x35, 0x34, 0x7d, 0x24, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x50, 0x0a,
	0x14, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0xa3, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74,
	0x65, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x43, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0xfa, 0x42, 0x21, 0x72, 0x1f, 0x10, 0x16, 0x18, 0xfe, 0x01,
	0x32, 0x18, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5f, 0x3a,
	0x5d, 0x7b, 0x32, 0x32, 0x2c, 0x32, 0x35, 0x34, 0x7d, 0x24, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x43,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xaf,
	0x03, 0x0a, 0x14, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x41, 0x50, 0x49, 0x12, 0x5f, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x27, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65,
	0x2e, 0x74, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x74, 0x64, 0x61, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x0c, 0x47, 0x72, 0x61, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b,
	0x69, 0x74, 0x65, 0x2e, 0x74, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x47, 0x72, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x74,
	0x64, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x68, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74,
	0x12, 0x2a, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x74, 0x64, 0x61, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x69,
	0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x74, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x29, 0x2e, 0x69, 0x6e, 0x64, 0x79,
	0x6b, 0x69, 0x74, 0x65, 0x2e, 0x74, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e,
	0x74, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0xf0, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74,
	0x65, 0x2e, 0x74, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x19, 0x54,
	0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f,
	0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x74, 0x64, 0x61,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x74, 0x64, 0x61, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x49, 0x54, 0x58, 0xaa, 0x02, 0x14, 0x49, 0x6e, 0x64, 0x79,
	0x6b, 0x69, 0x74, 0x65, 0x2e, 0x54, 0x64, 0x61, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xca, 0x02, 0x14, 0x49, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x5c, 0x54, 0x64, 0x61, 0x5c,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x20, 0x49, 0x6e, 0x64, 0x79, 0x6b, 0x69,
	0x74, 0x65, 0x5c, 0x54, 0x64, 0x61, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x49, 0x6e, 0x64,
	0x79, 0x6b, 0x69, 0x74, 0x65, 0x3a, 0x3a, 0x54, 0x64, 0x61, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDescOnce sync.Once
	file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDescData = file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDesc
)

func file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDescGZIP() []byte {
	file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDescOnce.Do(func() {
		file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDescData)
	})
	return file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDescData
}

var file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_indykite_tda_v1beta1_trusted_data_access_api_proto_goTypes = []any{
	(*ListConsentsRequest)(nil),   // 0: indykite.tda.v1beta1.ListConsentsRequest
	(*ListConsentsResponse)(nil),  // 1: indykite.tda.v1beta1.ListConsentsResponse
	(*DataAccessRequest)(nil),     // 2: indykite.tda.v1beta1.DataAccessRequest
	(*DataAccessResponse)(nil),    // 3: indykite.tda.v1beta1.DataAccessResponse
	(*GrantConsentRequest)(nil),   // 4: indykite.tda.v1beta1.GrantConsentRequest
	(*GrantConsentResponse)(nil),  // 5: indykite.tda.v1beta1.GrantConsentResponse
	(*RevokeConsentRequest)(nil),  // 6: indykite.tda.v1beta1.RevokeConsentRequest
	(*RevokeConsentResponse)(nil), // 7: indykite.tda.v1beta1.RevokeConsentResponse
	(*v1beta1.User)(nil),          // 8: indykite.knowledge.objects.v1beta1.User
	(*Consent)(nil),               // 9: indykite.tda.v1beta1.Consent
	(*v1beta1.Node)(nil),          // 10: indykite.knowledge.objects.v1beta1.Node
}
var file_indykite_tda_v1beta1_trusted_data_access_api_proto_depIdxs = []int32{
	8,  // 0: indykite.tda.v1beta1.ListConsentsRequest.user:type_name -> indykite.knowledge.objects.v1beta1.User
	9,  // 1: indykite.tda.v1beta1.ListConsentsResponse.consents:type_name -> indykite.tda.v1beta1.Consent
	8,  // 2: indykite.tda.v1beta1.DataAccessRequest.user:type_name -> indykite.knowledge.objects.v1beta1.User
	10, // 3: indykite.tda.v1beta1.DataAccessResponse.nodes:type_name -> indykite.knowledge.objects.v1beta1.Node
	8,  // 4: indykite.tda.v1beta1.GrantConsentRequest.user:type_name -> indykite.knowledge.objects.v1beta1.User
	8,  // 5: indykite.tda.v1beta1.RevokeConsentRequest.user:type_name -> indykite.knowledge.objects.v1beta1.User
	2,  // 6: indykite.tda.v1beta1.TrustedDataAccessAPI.DataAccess:input_type -> indykite.tda.v1beta1.DataAccessRequest
	4,  // 7: indykite.tda.v1beta1.TrustedDataAccessAPI.GrantConsent:input_type -> indykite.tda.v1beta1.GrantConsentRequest
	6,  // 8: indykite.tda.v1beta1.TrustedDataAccessAPI.RevokeConsent:input_type -> indykite.tda.v1beta1.RevokeConsentRequest
	0,  // 9: indykite.tda.v1beta1.TrustedDataAccessAPI.ListConsents:input_type -> indykite.tda.v1beta1.ListConsentsRequest
	3,  // 10: indykite.tda.v1beta1.TrustedDataAccessAPI.DataAccess:output_type -> indykite.tda.v1beta1.DataAccessResponse
	5,  // 11: indykite.tda.v1beta1.TrustedDataAccessAPI.GrantConsent:output_type -> indykite.tda.v1beta1.GrantConsentResponse
	7,  // 12: indykite.tda.v1beta1.TrustedDataAccessAPI.RevokeConsent:output_type -> indykite.tda.v1beta1.RevokeConsentResponse
	1,  // 13: indykite.tda.v1beta1.TrustedDataAccessAPI.ListConsents:output_type -> indykite.tda.v1beta1.ListConsentsResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_indykite_tda_v1beta1_trusted_data_access_api_proto_init() }
func file_indykite_tda_v1beta1_trusted_data_access_api_proto_init() {
	if File_indykite_tda_v1beta1_trusted_data_access_api_proto != nil {
		return
	}
	file_indykite_tda_v1beta1_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListConsentsRequest); i {
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
		file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListConsentsResponse); i {
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
		file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DataAccessRequest); i {
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
		file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DataAccessResponse); i {
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
		file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GrantConsentRequest); i {
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
		file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GrantConsentResponse); i {
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
		file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*RevokeConsentRequest); i {
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
		file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*RevokeConsentResponse); i {
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
			RawDescriptor: file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_indykite_tda_v1beta1_trusted_data_access_api_proto_goTypes,
		DependencyIndexes: file_indykite_tda_v1beta1_trusted_data_access_api_proto_depIdxs,
		MessageInfos:      file_indykite_tda_v1beta1_trusted_data_access_api_proto_msgTypes,
	}.Build()
	File_indykite_tda_v1beta1_trusted_data_access_api_proto = out.File
	file_indykite_tda_v1beta1_trusted_data_access_api_proto_rawDesc = nil
	file_indykite_tda_v1beta1_trusted_data_access_api_proto_goTypes = nil
	file_indykite_tda_v1beta1_trusted_data_access_api_proto_depIdxs = nil
}
