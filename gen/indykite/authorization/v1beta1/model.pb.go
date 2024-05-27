// Copyright (c) 2023 IndyKite
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

// Authorization Service Description.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: indykite/authorization/v1beta1/model.proto

package authorizationv1beta1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	v1beta2 "github.com/indykite/indykite-sdk-go/gen/indykite/identity/v1beta2"
	v1beta1 "github.com/indykite/indykite-sdk-go/gen/indykite/objects/v1beta1"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/google/gnostic/openapiv3"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Subject to check if is authorized.
type Subject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Subject:
	//
	//	*Subject_DigitalTwinIdentifier
	//	*Subject_DigitalTwinId
	//	*Subject_DigitalTwinProperty
	//	*Subject_IndykiteAccessToken
	//	*Subject_ExternalId
	Subject isSubject_Subject `protobuf_oneof:"subject"`
}

func (x *Subject) Reset() {
	*x = Subject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_authorization_v1beta1_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subject) ProtoMessage() {}

func (x *Subject) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_authorization_v1beta1_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subject.ProtoReflect.Descriptor instead.
func (*Subject) Descriptor() ([]byte, []int) {
	return file_indykite_authorization_v1beta1_model_proto_rawDescGZIP(), []int{0}
}

func (m *Subject) GetSubject() isSubject_Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

// Deprecated: Marked as deprecated in indykite/authorization/v1beta1/model.proto.
func (x *Subject) GetDigitalTwinIdentifier() *v1beta2.DigitalTwinIdentifier {
	if x, ok := x.GetSubject().(*Subject_DigitalTwinIdentifier); ok {
		return x.DigitalTwinIdentifier
	}
	return nil
}

func (x *Subject) GetDigitalTwinId() *DigitalTwin {
	if x, ok := x.GetSubject().(*Subject_DigitalTwinId); ok {
		return x.DigitalTwinId
	}
	return nil
}

func (x *Subject) GetDigitalTwinProperty() *Property {
	if x, ok := x.GetSubject().(*Subject_DigitalTwinProperty); ok {
		return x.DigitalTwinProperty
	}
	return nil
}

func (x *Subject) GetIndykiteAccessToken() string {
	if x, ok := x.GetSubject().(*Subject_IndykiteAccessToken); ok {
		return x.IndykiteAccessToken
	}
	return ""
}

func (x *Subject) GetExternalId() *ExternalID {
	if x, ok := x.GetSubject().(*Subject_ExternalId); ok {
		return x.ExternalId
	}
	return nil
}

type isSubject_Subject interface {
	isSubject_Subject()
}

type Subject_DigitalTwinIdentifier struct {
	// Deprecated: Marked as deprecated in indykite/authorization/v1beta1/model.proto.
	DigitalTwinIdentifier *v1beta2.DigitalTwinIdentifier `protobuf:"bytes,1,opt,name=digital_twin_identifier,json=digitalTwinIdentifier,proto3,oneof"`
}

type Subject_DigitalTwinId struct {
	DigitalTwinId *DigitalTwin `protobuf:"bytes,2,opt,name=digital_twin_id,json=digitalTwinId,proto3,oneof"`
}

type Subject_DigitalTwinProperty struct {
	DigitalTwinProperty *Property `protobuf:"bytes,3,opt,name=digital_twin_property,json=digitalTwinProperty,proto3,oneof"`
}

type Subject_IndykiteAccessToken struct {
	IndykiteAccessToken string `protobuf:"bytes,4,opt,name=indykite_access_token,json=indykiteAccessToken,proto3,oneof"`
}

type Subject_ExternalId struct {
	ExternalId *ExternalID `protobuf:"bytes,5,opt,name=external_id,json=externalId,proto3,oneof"`
}

func (*Subject_DigitalTwinIdentifier) isSubject_Subject() {}

func (*Subject_DigitalTwinId) isSubject_Subject() {}

func (*Subject_DigitalTwinProperty) isSubject_Subject() {}

func (*Subject_IndykiteAccessToken) isSubject_Subject() {}

func (*Subject_ExternalId) isSubject_Subject() {}

// DigitalTwin represents a digital entity.
type DigitalTwin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id the unique credential identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DigitalTwin) Reset() {
	*x = DigitalTwin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_authorization_v1beta1_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DigitalTwin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DigitalTwin) ProtoMessage() {}

func (x *DigitalTwin) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_authorization_v1beta1_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DigitalTwin.ProtoReflect.Descriptor instead.
func (*DigitalTwin) Descriptor() ([]byte, []int) {
	return file_indykite_authorization_v1beta1_model_proto_rawDescGZIP(), []int{1}
}

func (x *DigitalTwin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string         `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value *v1beta1.Value `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Property) Reset() {
	*x = Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_authorization_v1beta1_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Property) ProtoMessage() {}

func (x *Property) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_authorization_v1beta1_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Property.ProtoReflect.Descriptor instead.
func (*Property) Descriptor() ([]byte, []int) {
	return file_indykite_authorization_v1beta1_model_proto_rawDescGZIP(), []int{2}
}

func (x *Property) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Property) GetValue() *v1beta1.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type ExternalID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	ExternalId string `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
}

func (x *ExternalID) Reset() {
	*x = ExternalID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_authorization_v1beta1_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalID) ProtoMessage() {}

func (x *ExternalID) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_authorization_v1beta1_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalID.ProtoReflect.Descriptor instead.
func (*ExternalID) Descriptor() ([]byte, []int) {
	return file_indykite_authorization_v1beta1_model_proto_rawDescGZIP(), []int{3}
}

func (x *ExternalID) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ExternalID) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

type InputParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*InputParam_StringValue
	//	*InputParam_BoolValue
	//	*InputParam_IntegerValue
	//	*InputParam_DoubleValue
	Value isInputParam_Value `protobuf_oneof:"value"`
}

func (x *InputParam) Reset() {
	*x = InputParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_authorization_v1beta1_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputParam) ProtoMessage() {}

func (x *InputParam) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_authorization_v1beta1_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputParam.ProtoReflect.Descriptor instead.
func (*InputParam) Descriptor() ([]byte, []int) {
	return file_indykite_authorization_v1beta1_model_proto_rawDescGZIP(), []int{4}
}

func (m *InputParam) GetValue() isInputParam_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *InputParam) GetStringValue() string {
	if x, ok := x.GetValue().(*InputParam_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *InputParam) GetBoolValue() bool {
	if x, ok := x.GetValue().(*InputParam_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *InputParam) GetIntegerValue() int64 {
	if x, ok := x.GetValue().(*InputParam_IntegerValue); ok {
		return x.IntegerValue
	}
	return 0
}

func (x *InputParam) GetDoubleValue() float64 {
	if x, ok := x.GetValue().(*InputParam_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

type isInputParam_Value interface {
	isInputParam_Value()
}

type InputParam_StringValue struct {
	// A string value.
	StringValue string `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type InputParam_BoolValue struct {
	// A boolean value.
	BoolValue bool `protobuf:"varint,2,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type InputParam_IntegerValue struct {
	// An integer value.
	IntegerValue int64 `protobuf:"varint,3,opt,name=integer_value,json=integerValue,proto3,oneof"`
}

type InputParam_DoubleValue struct {
	// A double value.
	DoubleValue float64 `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

func (*InputParam_StringValue) isInputParam_Value() {}

func (*InputParam_BoolValue) isInputParam_Value() {}

func (*InputParam_IntegerValue) isInputParam_Value() {}

func (*InputParam_DoubleValue) isInputParam_Value() {}

var File_indykite_authorization_v1beta1_model_proto protoreflect.FileDescriptor

var file_indykite_authorization_v1beta1_model_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x69, 0x6e,
	0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x24, 0x67, 0x6e,
	0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x69, 0x6e, 0x64,
	0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x25, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x03, 0x0a, 0x07, 0x53, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x7b, 0x0a, 0x17, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x5f, 0x74, 0x77, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74,
	0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x32, 0x2e, 0x44, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x54, 0x77, 0x69, 0x6e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x0f, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0xba, 0x47, 0x02, 0x40, 0x01, 0x18, 0x01, 0x48, 0x00, 0x52, 0x15, 0x64, 0x69, 0x67,
	0x69, 0x74, 0x61, 0x6c, 0x54, 0x77, 0x69, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x5f, 0x0a, 0x0f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x77,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x69, 0x6e,
	0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x69, 0x67,
	0x69, 0x74, 0x61, 0x6c, 0x54, 0x77, 0x69, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x48, 0x00, 0x52, 0x0d, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x54, 0x77, 0x69,
	0x6e, 0x49, 0x64, 0x12, 0x68, 0x0a, 0x15, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x5f, 0x74,
	0x77, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x13, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61,
	0x6c, 0x54, 0x77, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x3d, 0x0a,
	0x15, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x14, 0x48, 0x00, 0x52, 0x13, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x57, 0x0a, 0x0b,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x49, 0x64, 0x42, 0x0e, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x45, 0x0a, 0x0b, 0x44, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x54, 0x77, 0x69, 0x6e, 0x12, 0x36, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x26, 0xfa, 0x42, 0x23, 0x72, 0x21, 0x10, 0x1b, 0x18, 0x64, 0x32, 0x1b, 0x5e, 0x67, 0x69,
	0x64, 0x3a, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5f, 0x5d, 0x7b,
	0x32, 0x37, 0x2c, 0x31, 0x30, 0x30, 0x7d, 0x24, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6a, 0x0a, 0x08,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18,
	0x14, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74,
	0x65, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x60, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xfa, 0x42, 0x11, 0x72, 0x0f, 0x18, 0x40, 0x32, 0x0b, 0x5e,
	0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5d, 0x2a, 0x24, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x28, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x22, 0xb7, 0x01, 0x0a, 0x0a, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x2e, 0x0a, 0x0c, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x32, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6f, 0x6f,
	0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52,
	0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x03, 0xf8, 0x42, 0x01, 0x42, 0xa7, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x64,
	0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0a, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x69,
	0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x49, 0x41, 0x58, 0xaa, 0x02, 0x1e, 0x49,
	0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1e,
	0x49, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02,
	0x2a, 0x49, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x49, 0x6e,
	0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_indykite_authorization_v1beta1_model_proto_rawDescOnce sync.Once
	file_indykite_authorization_v1beta1_model_proto_rawDescData = file_indykite_authorization_v1beta1_model_proto_rawDesc
)

func file_indykite_authorization_v1beta1_model_proto_rawDescGZIP() []byte {
	file_indykite_authorization_v1beta1_model_proto_rawDescOnce.Do(func() {
		file_indykite_authorization_v1beta1_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_indykite_authorization_v1beta1_model_proto_rawDescData)
	})
	return file_indykite_authorization_v1beta1_model_proto_rawDescData
}

var file_indykite_authorization_v1beta1_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_indykite_authorization_v1beta1_model_proto_goTypes = []interface{}{
	(*Subject)(nil),                       // 0: indykite.authorization.v1beta1.Subject
	(*DigitalTwin)(nil),                   // 1: indykite.authorization.v1beta1.DigitalTwin
	(*Property)(nil),                      // 2: indykite.authorization.v1beta1.Property
	(*ExternalID)(nil),                    // 3: indykite.authorization.v1beta1.ExternalID
	(*InputParam)(nil),                    // 4: indykite.authorization.v1beta1.InputParam
	(*v1beta2.DigitalTwinIdentifier)(nil), // 5: indykite.identity.v1beta2.DigitalTwinIdentifier
	(*v1beta1.Value)(nil),                 // 6: indykite.objects.v1beta1.Value
}
var file_indykite_authorization_v1beta1_model_proto_depIdxs = []int32{
	5, // 0: indykite.authorization.v1beta1.Subject.digital_twin_identifier:type_name -> indykite.identity.v1beta2.DigitalTwinIdentifier
	1, // 1: indykite.authorization.v1beta1.Subject.digital_twin_id:type_name -> indykite.authorization.v1beta1.DigitalTwin
	2, // 2: indykite.authorization.v1beta1.Subject.digital_twin_property:type_name -> indykite.authorization.v1beta1.Property
	3, // 3: indykite.authorization.v1beta1.Subject.external_id:type_name -> indykite.authorization.v1beta1.ExternalID
	6, // 4: indykite.authorization.v1beta1.Property.value:type_name -> indykite.objects.v1beta1.Value
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_indykite_authorization_v1beta1_model_proto_init() }
func file_indykite_authorization_v1beta1_model_proto_init() {
	if File_indykite_authorization_v1beta1_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_indykite_authorization_v1beta1_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subject); i {
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
		file_indykite_authorization_v1beta1_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DigitalTwin); i {
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
		file_indykite_authorization_v1beta1_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Property); i {
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
		file_indykite_authorization_v1beta1_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalID); i {
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
		file_indykite_authorization_v1beta1_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputParam); i {
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
	file_indykite_authorization_v1beta1_model_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Subject_DigitalTwinIdentifier)(nil),
		(*Subject_DigitalTwinId)(nil),
		(*Subject_DigitalTwinProperty)(nil),
		(*Subject_IndykiteAccessToken)(nil),
		(*Subject_ExternalId)(nil),
	}
	file_indykite_authorization_v1beta1_model_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*InputParam_StringValue)(nil),
		(*InputParam_BoolValue)(nil),
		(*InputParam_IntegerValue)(nil),
		(*InputParam_DoubleValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_indykite_authorization_v1beta1_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_indykite_authorization_v1beta1_model_proto_goTypes,
		DependencyIndexes: file_indykite_authorization_v1beta1_model_proto_depIdxs,
		MessageInfos:      file_indykite_authorization_v1beta1_model_proto_msgTypes,
	}.Build()
	File_indykite_authorization_v1beta1_model_proto = out.File
	file_indykite_authorization_v1beta1_model_proto_rawDesc = nil
	file_indykite_authorization_v1beta1_model_proto_goTypes = nil
	file_indykite_authorization_v1beta1_model_proto_depIdxs = nil
}
