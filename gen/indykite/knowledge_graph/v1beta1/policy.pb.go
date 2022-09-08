// Copyright (c) 2022 IndyKite AS.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: indykite/knowledge_graph/v1beta1/policy.proto

package knowledge_graphv1beta1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	v1beta1 "github.com/indykite/jarvis-sdk-go/gen/indykite/identity/v1beta1"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path    *Path    `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Actions []string `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
	Active  bool     `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_indykite_knowledge_graph_v1beta1_policy_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetPath() *Path {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *Policy) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *Policy) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type Path struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubjectId     string               `protobuf:"bytes,1,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	ResourceId    string               `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Entities      []*Path_Entity       `protobuf:"bytes,3,rep,name=entities,proto3" json:"entities,omitempty"`
	Relationships []*Path_Relationship `protobuf:"bytes,4,rep,name=relationships,proto3" json:"relationships,omitempty"`
}

func (x *Path) Reset() {
	*x = Path{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Path) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Path) ProtoMessage() {}

func (x *Path) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Path.ProtoReflect.Descriptor instead.
func (*Path) Descriptor() ([]byte, []int) {
	return file_indykite_knowledge_graph_v1beta1_policy_proto_rawDescGZIP(), []int{1}
}

func (x *Path) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *Path) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *Path) GetEntities() []*Path_Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *Path) GetRelationships() []*Path_Relationship {
	if x != nil {
		return x.Relationships
	}
	return nil
}

type Path_Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  string                           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Labels              []string                         `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	IdentityProperties  []*Path_Entity_IdentityProperty  `protobuf:"bytes,4,rep,name=identity_properties,json=identityProperties,proto3" json:"identity_properties,omitempty"`
	KnowledgeProperties []*Path_Entity_KnowledgeProperty `protobuf:"bytes,5,rep,name=knowledge_properties,json=knowledgeProperties,proto3" json:"knowledge_properties,omitempty"`
}

func (x *Path_Entity) Reset() {
	*x = Path_Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Path_Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Path_Entity) ProtoMessage() {}

func (x *Path_Entity) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Path_Entity.ProtoReflect.Descriptor instead.
func (*Path_Entity) Descriptor() ([]byte, []int) {
	return file_indykite_knowledge_graph_v1beta1_policy_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Path_Entity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Path_Entity) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Path_Entity) GetIdentityProperties() []*Path_Entity_IdentityProperty {
	if x != nil {
		return x.IdentityProperties
	}
	return nil
}

func (x *Path_Entity) GetKnowledgeProperties() []*Path_Entity_KnowledgeProperty {
	if x != nil {
		return x.KnowledgeProperties
	}
	return nil
}

type Path_Relationship struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source         string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Target         string   `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Types          []string `protobuf:"bytes,3,rep,name=types,proto3" json:"types,omitempty"`
	NonDirectional bool     `protobuf:"varint,4,opt,name=non_directional,json=nonDirectional,proto3" json:"non_directional,omitempty"`
}

func (x *Path_Relationship) Reset() {
	*x = Path_Relationship{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Path_Relationship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Path_Relationship) ProtoMessage() {}

func (x *Path_Relationship) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Path_Relationship.ProtoReflect.Descriptor instead.
func (*Path_Relationship) Descriptor() ([]byte, []int) {
	return file_indykite_knowledge_graph_v1beta1_policy_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Path_Relationship) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Path_Relationship) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Path_Relationship) GetTypes() []string {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *Path_Relationship) GetNonDirectional() bool {
	if x != nil {
		return x.NonDirectional
	}
	return false
}

type Path_Entity_IdentityProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Property              string                 `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
	Value                 string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	MinimumAssuranceLevel v1beta1.AssuranceLevel `protobuf:"varint,3,opt,name=minimum_assurance_level,json=minimumAssuranceLevel,proto3,enum=indykite.identity.v1beta1.AssuranceLevel" json:"minimum_assurance_level,omitempty"`
	AllowedIssuers        []string               `protobuf:"bytes,4,rep,name=allowed_issuers,json=allowedIssuers,proto3" json:"allowed_issuers,omitempty"`
	VerificationTime      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=verification_time,json=verificationTime,proto3" json:"verification_time,omitempty"`
	AllowedVerifiers      []string               `protobuf:"bytes,6,rep,name=allowed_verifiers,json=allowedVerifiers,proto3" json:"allowed_verifiers,omitempty"`
	MustBePrimary         bool                   `protobuf:"varint,7,opt,name=must_be_primary,json=mustBePrimary,proto3" json:"must_be_primary,omitempty"`
}

func (x *Path_Entity_IdentityProperty) Reset() {
	*x = Path_Entity_IdentityProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Path_Entity_IdentityProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Path_Entity_IdentityProperty) ProtoMessage() {}

func (x *Path_Entity_IdentityProperty) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Path_Entity_IdentityProperty.ProtoReflect.Descriptor instead.
func (*Path_Entity_IdentityProperty) Descriptor() ([]byte, []int) {
	return file_indykite_knowledge_graph_v1beta1_policy_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *Path_Entity_IdentityProperty) GetProperty() string {
	if x != nil {
		return x.Property
	}
	return ""
}

func (x *Path_Entity_IdentityProperty) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Path_Entity_IdentityProperty) GetMinimumAssuranceLevel() v1beta1.AssuranceLevel {
	if x != nil {
		return x.MinimumAssuranceLevel
	}
	return v1beta1.AssuranceLevel(0)
}

func (x *Path_Entity_IdentityProperty) GetAllowedIssuers() []string {
	if x != nil {
		return x.AllowedIssuers
	}
	return nil
}

func (x *Path_Entity_IdentityProperty) GetVerificationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.VerificationTime
	}
	return nil
}

func (x *Path_Entity_IdentityProperty) GetAllowedVerifiers() []string {
	if x != nil {
		return x.AllowedVerifiers
	}
	return nil
}

func (x *Path_Entity_IdentityProperty) GetMustBePrimary() bool {
	if x != nil {
		return x.MustBePrimary
	}
	return false
}

type Path_Entity_KnowledgeProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Property string `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Path_Entity_KnowledgeProperty) Reset() {
	*x = Path_Entity_KnowledgeProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Path_Entity_KnowledgeProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Path_Entity_KnowledgeProperty) ProtoMessage() {}

func (x *Path_Entity_KnowledgeProperty) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Path_Entity_KnowledgeProperty.ProtoReflect.Descriptor instead.
func (*Path_Entity_KnowledgeProperty) Descriptor() ([]byte, []int) {
	return file_indykite_knowledge_graph_v1beta1_policy_proto_rawDescGZIP(), []int{1, 0, 1}
}

func (x *Path_Entity_KnowledgeProperty) GetProperty() string {
	if x != nil {
		return x.Property
	}
	return ""
}

func (x *Path_Entity_KnowledgeProperty) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_indykite_knowledge_graph_v1beta1_policy_proto protoreflect.FileDescriptor

var file_indykite_knowledge_graph_v1beta1_policy_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x6b, 0x6e, 0x6f, 0x77, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x20, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x69, 0x6e, 0x64, 0x79,
	0x6b, 0x69, 0x74, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x44, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x48, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x2e, 0xfa, 0x42, 0x2b, 0x92, 0x01,
	0x28, 0x08, 0x01, 0x10, 0x20, 0x18, 0x01, 0x22, 0x20, 0x72, 0x1e, 0x10, 0x02, 0x18, 0x64, 0x32,
	0x18, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2e, 0x3a, 0x5f, 0x5c,
	0x2d, 0x5c, 0x2f, 0x5d, 0x7b, 0x32, 0x2c, 0x7d, 0x24, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0xbc, 0x0a, 0x0a, 0x04, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x28, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x02,
	0x18, 0x32, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x32, 0x52, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x53, 0x0a, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x69, 0x6e,
	0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50,
	0x61, 0x74, 0x68, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92,
	0x01, 0x02, 0x08, 0x02, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x63,
	0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65,
	0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x2e, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92,
	0x01, 0x02, 0x08, 0x01, 0x52, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68,
	0x69, 0x70, 0x73, 0x1a, 0xe1, 0x06, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x19,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72,
	0x04, 0x10, 0x02, 0x18, 0x32, 0x52, 0x02, 0x69, 0x64, 0x12, 0x40, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x28, 0xfa, 0x42, 0x25, 0x92, 0x01,
	0x22, 0x08, 0x01, 0x10, 0x20, 0x18, 0x01, 0x22, 0x1a, 0x72, 0x18, 0x10, 0x02, 0x18, 0x32, 0x32,
	0x12, 0x5e, 0x28, 0x3f, 0x3a, 0x5b, 0x41, 0x2d, 0x5a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x2b,
	0x29, 0x2b, 0x24, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x6f, 0x0a, 0x13, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b,
	0x69, 0x74, 0x65, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x68,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x12, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x72, 0x0a, 0x14,
	0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x69, 0x6e, 0x64,
	0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x5f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61,
	0x74, 0x68, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x13, 0x6b, 0x6e, 0x6f,
	0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x1a, 0x9e, 0x03, 0x0a, 0x10, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0xfa, 0x42, 0x21, 0x72, 0x1f, 0x10, 0x02,
	0x18, 0x80, 0x02, 0x32, 0x18, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5f, 0x5d, 0x5b,
	0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x2b, 0x24, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x6b, 0x0a,
	0x17, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x61, 0x73, 0x73, 0x75, 0x72, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29,
	0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x75, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x15, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x41, 0x73, 0x73, 0x75,
	0x72, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x72, 0x73, 0x12, 0x47, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x75, 0x73,
	0x74, 0x5f, 0x62, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x6d, 0x75, 0x73, 0x74, 0x42, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x1a, 0x74, 0x0a, 0x11, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0xfa, 0x42, 0x21, 0x72, 0x1f, 0x10,
	0x02, 0x18, 0x80, 0x02, 0x32, 0x18, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5f, 0x5d,
	0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x2b, 0x24, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0xbf, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x21, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10,
	0x02, 0x18, 0x32, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06,
	0x72, 0x04, 0x10, 0x02, 0x18, 0x32, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x40,
	0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x2a, 0xfa,
	0x42, 0x27, 0x92, 0x01, 0x24, 0x08, 0x00, 0x10, 0x20, 0x18, 0x01, 0x22, 0x1c, 0x72, 0x1a, 0x10,
	0x02, 0x18, 0x32, 0x32, 0x14, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x5d, 0x2b, 0x28, 0x3f, 0x3a, 0x5f,
	0x5b, 0x41, 0x2d, 0x5a, 0x5d, 0x2b, 0x29, 0x2a, 0x24, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6e, 0x6f, 0x6e, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x42, 0xb0, 0x02, 0x0a, 0x24, 0x63, 0x6f,
	0x6d, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x42, 0x0b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x2d, 0x73, 0x64,
	0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74,
	0x65, 0x2f, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xa2, 0x02, 0x03, 0x49, 0x4b, 0x58, 0xaa, 0x02, 0x1f, 0x49, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74,
	0x65, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1f, 0x49, 0x6e, 0x64, 0x79, 0x6b,
	0x69, 0x74, 0x65, 0x5c, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x2b, 0x49, 0x6e, 0x64,
	0x79, 0x6b, 0x69, 0x74, 0x65, 0x5c, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x49, 0x6e, 0x64, 0x79, 0x6b,
	0x69, 0x74, 0x65, 0x3a, 0x3a, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_indykite_knowledge_graph_v1beta1_policy_proto_rawDescOnce sync.Once
	file_indykite_knowledge_graph_v1beta1_policy_proto_rawDescData = file_indykite_knowledge_graph_v1beta1_policy_proto_rawDesc
)

func file_indykite_knowledge_graph_v1beta1_policy_proto_rawDescGZIP() []byte {
	file_indykite_knowledge_graph_v1beta1_policy_proto_rawDescOnce.Do(func() {
		file_indykite_knowledge_graph_v1beta1_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_indykite_knowledge_graph_v1beta1_policy_proto_rawDescData)
	})
	return file_indykite_knowledge_graph_v1beta1_policy_proto_rawDescData
}

var file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_indykite_knowledge_graph_v1beta1_policy_proto_goTypes = []interface{}{
	(*Policy)(nil),                        // 0: indykite.knowledge_graph.v1beta1.Policy
	(*Path)(nil),                          // 1: indykite.knowledge_graph.v1beta1.Path
	(*Path_Entity)(nil),                   // 2: indykite.knowledge_graph.v1beta1.Path.Entity
	(*Path_Relationship)(nil),             // 3: indykite.knowledge_graph.v1beta1.Path.Relationship
	(*Path_Entity_IdentityProperty)(nil),  // 4: indykite.knowledge_graph.v1beta1.Path.Entity.IdentityProperty
	(*Path_Entity_KnowledgeProperty)(nil), // 5: indykite.knowledge_graph.v1beta1.Path.Entity.KnowledgeProperty
	(v1beta1.AssuranceLevel)(0),           // 6: indykite.identity.v1beta1.AssuranceLevel
	(*timestamppb.Timestamp)(nil),         // 7: google.protobuf.Timestamp
}
var file_indykite_knowledge_graph_v1beta1_policy_proto_depIdxs = []int32{
	1, // 0: indykite.knowledge_graph.v1beta1.Policy.path:type_name -> indykite.knowledge_graph.v1beta1.Path
	2, // 1: indykite.knowledge_graph.v1beta1.Path.entities:type_name -> indykite.knowledge_graph.v1beta1.Path.Entity
	3, // 2: indykite.knowledge_graph.v1beta1.Path.relationships:type_name -> indykite.knowledge_graph.v1beta1.Path.Relationship
	4, // 3: indykite.knowledge_graph.v1beta1.Path.Entity.identity_properties:type_name -> indykite.knowledge_graph.v1beta1.Path.Entity.IdentityProperty
	5, // 4: indykite.knowledge_graph.v1beta1.Path.Entity.knowledge_properties:type_name -> indykite.knowledge_graph.v1beta1.Path.Entity.KnowledgeProperty
	6, // 5: indykite.knowledge_graph.v1beta1.Path.Entity.IdentityProperty.minimum_assurance_level:type_name -> indykite.identity.v1beta1.AssuranceLevel
	7, // 6: indykite.knowledge_graph.v1beta1.Path.Entity.IdentityProperty.verification_time:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_indykite_knowledge_graph_v1beta1_policy_proto_init() }
func file_indykite_knowledge_graph_v1beta1_policy_proto_init() {
	if File_indykite_knowledge_graph_v1beta1_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Path); i {
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
		file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Path_Entity); i {
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
		file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Path_Relationship); i {
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
		file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Path_Entity_IdentityProperty); i {
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
		file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Path_Entity_KnowledgeProperty); i {
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
			RawDescriptor: file_indykite_knowledge_graph_v1beta1_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_indykite_knowledge_graph_v1beta1_policy_proto_goTypes,
		DependencyIndexes: file_indykite_knowledge_graph_v1beta1_policy_proto_depIdxs,
		MessageInfos:      file_indykite_knowledge_graph_v1beta1_policy_proto_msgTypes,
	}.Build()
	File_indykite_knowledge_graph_v1beta1_policy_proto = out.File
	file_indykite_knowledge_graph_v1beta1_policy_proto_rawDesc = nil
	file_indykite_knowledge_graph_v1beta1_policy_proto_goTypes = nil
	file_indykite_knowledge_graph_v1beta1_policy_proto_depIdxs = nil
}
