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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: indykite/entitymatching/v1beta1/entity_matching_api.proto

package entitymatchingv1beta1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RunEntityMatchingPipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id is the Globally unique identifier of the pipeline to run.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// CustomPropertyMappings contains the rules to match nodes properties.
	// If empty, the default rules will be used (stored as part of the pipeline configuration).
	CustomPropertyMappings []*PropertyMapping `protobuf:"bytes,2,rep,name=custom_property_mappings,json=customPropertyMappings,proto3" json:"custom_property_mappings,omitempty"`
	// SimilarityScoreCutoff defines the threshold (in range [0,1]), above which entities will be automatically matched.
	SimilarityScoreCutoff float32 `protobuf:"fixed32,3,opt,name=similarity_score_cutoff,json=similarityScoreCutoff,proto3" json:"similarity_score_cutoff,omitempty"`
}

func (x *RunEntityMatchingPipelineRequest) Reset() {
	*x = RunEntityMatchingPipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunEntityMatchingPipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunEntityMatchingPipelineRequest) ProtoMessage() {}

func (x *RunEntityMatchingPipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunEntityMatchingPipelineRequest.ProtoReflect.Descriptor instead.
func (*RunEntityMatchingPipelineRequest) Descriptor() ([]byte, []int) {
	return file_indykite_entitymatching_v1beta1_entity_matching_api_proto_rawDescGZIP(), []int{0}
}

func (x *RunEntityMatchingPipelineRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RunEntityMatchingPipelineRequest) GetCustomPropertyMappings() []*PropertyMapping {
	if x != nil {
		return x.CustomPropertyMappings
	}
	return nil
}

func (x *RunEntityMatchingPipelineRequest) GetSimilarityScoreCutoff() float32 {
	if x != nil {
		return x.SimilarityScoreCutoff
	}
	return 0
}

type RunEntityMatchingPipelineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id is the Globally unique identifier of the pipeline that was started.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. The time at which the pipeline was last run.
	LastRunTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_run_time,json=lastRunTime,proto3" json:"last_run_time,omitempty"`
	// Output only. Multiversion concurrency control version.
	Etag string `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (x *RunEntityMatchingPipelineResponse) Reset() {
	*x = RunEntityMatchingPipelineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunEntityMatchingPipelineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunEntityMatchingPipelineResponse) ProtoMessage() {}

func (x *RunEntityMatchingPipelineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunEntityMatchingPipelineResponse.ProtoReflect.Descriptor instead.
func (*RunEntityMatchingPipelineResponse) Descriptor() ([]byte, []int) {
	return file_indykite_entitymatching_v1beta1_entity_matching_api_proto_rawDescGZIP(), []int{1}
}

func (x *RunEntityMatchingPipelineResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RunEntityMatchingPipelineResponse) GetLastRunTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastRunTime
	}
	return nil
}

func (x *RunEntityMatchingPipelineResponse) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

type ReadSuggestedPropertyMappingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id contains the Globally Unique Identifier of the object with server generated id.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadSuggestedPropertyMappingRequest) Reset() {
	*x = ReadSuggestedPropertyMappingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadSuggestedPropertyMappingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadSuggestedPropertyMappingRequest) ProtoMessage() {}

func (x *ReadSuggestedPropertyMappingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadSuggestedPropertyMappingRequest.ProtoReflect.Descriptor instead.
func (*ReadSuggestedPropertyMappingRequest) Descriptor() ([]byte, []int) {
	return file_indykite_entitymatching_v1beta1_entity_matching_api_proto_rawDescGZIP(), []int{2}
}

func (x *ReadSuggestedPropertyMappingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReadSuggestedPropertyMappingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Globally unique identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// PropertyMappings contains the rules the pipeline will use to match source nodes with target nodes.
	PropertyMappings []*PropertyMapping `protobuf:"bytes,2,rep,name=property_mappings,json=propertyMappings,proto3" json:"property_mappings,omitempty"`
	// PropertyMappingStatus is the status assigned to the pipeline's step that maps node types' properties.
	// If the status is not SUCCESS, the PropertyMappings might be empty.
	PropertyMappingStatus PipelineStatus `protobuf:"varint,3,opt,name=property_mapping_status,json=propertyMappingStatus,proto3,enum=indykite.entitymatching.v1beta1.PipelineStatus" json:"property_mapping_status,omitempty"`
}

func (x *ReadSuggestedPropertyMappingResponse) Reset() {
	*x = ReadSuggestedPropertyMappingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadSuggestedPropertyMappingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadSuggestedPropertyMappingResponse) ProtoMessage() {}

func (x *ReadSuggestedPropertyMappingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadSuggestedPropertyMappingResponse.ProtoReflect.Descriptor instead.
func (*ReadSuggestedPropertyMappingResponse) Descriptor() ([]byte, []int) {
	return file_indykite_entitymatching_v1beta1_entity_matching_api_proto_rawDescGZIP(), []int{3}
}

func (x *ReadSuggestedPropertyMappingResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReadSuggestedPropertyMappingResponse) GetPropertyMappings() []*PropertyMapping {
	if x != nil {
		return x.PropertyMappings
	}
	return nil
}

func (x *ReadSuggestedPropertyMappingResponse) GetPropertyMappingStatus() PipelineStatus {
	if x != nil {
		return x.PropertyMappingStatus
	}
	return PipelineStatus_PIPELINE_STATUS_STATUS_INVALID
}

type ReadEntityMatchingReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id contains the Globally Unique Identifier of the object with server generated id.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadEntityMatchingReportRequest) Reset() {
	*x = ReadEntityMatchingReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadEntityMatchingReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadEntityMatchingReportRequest) ProtoMessage() {}

func (x *ReadEntityMatchingReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadEntityMatchingReportRequest.ProtoReflect.Descriptor instead.
func (*ReadEntityMatchingReportRequest) Descriptor() ([]byte, []int) {
	return file_indykite_entitymatching_v1beta1_entity_matching_api_proto_rawDescGZIP(), []int{4}
}

func (x *ReadEntityMatchingReportRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReadEntityMatchingReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Globally unique identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ReportURL contains a pre-signed URL to the report document.
	ReportUrl string `protobuf:"bytes,2,opt,name=report_url,json=reportUrl,proto3" json:"report_url,omitempty"`
	// URLExpireTime defines when the report will no longer be accessible at the given ReportURL.
	UrlExpireTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=url_expire_time,json=urlExpireTime,proto3" json:"url_expire_time,omitempty"`
	// EntityMappingStatus is the status assigned to the pipeline's step that matches node entities.
	// If the status is not SUCCESS, the report_url might be empty.
	EntityMatchingStatus PipelineStatus `protobuf:"varint,4,opt,name=entity_matching_status,json=entityMatchingStatus,proto3,enum=indykite.entitymatching.v1beta1.PipelineStatus" json:"entity_matching_status,omitempty"`
}

func (x *ReadEntityMatchingReportResponse) Reset() {
	*x = ReadEntityMatchingReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadEntityMatchingReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadEntityMatchingReportResponse) ProtoMessage() {}

func (x *ReadEntityMatchingReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadEntityMatchingReportResponse.ProtoReflect.Descriptor instead.
func (*ReadEntityMatchingReportResponse) Descriptor() ([]byte, []int) {
	return file_indykite_entitymatching_v1beta1_entity_matching_api_proto_rawDescGZIP(), []int{5}
}

func (x *ReadEntityMatchingReportResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReadEntityMatchingReportResponse) GetReportUrl() string {
	if x != nil {
		return x.ReportUrl
	}
	return ""
}

func (x *ReadEntityMatchingReportResponse) GetUrlExpireTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UrlExpireTime
	}
	return nil
}

func (x *ReadEntityMatchingReportResponse) GetEntityMatchingStatus() PipelineStatus {
	if x != nil {
		return x.EntityMatchingStatus
	}
	return PipelineStatus_PIPELINE_STATUS_STATUS_INVALID
}

var File_indykite_entitymatching_v1beta1_entity_matching_api_proto protoreflect.FileDescriptor

var file_indykite_entitymatching_v1beta1_entity_matching_api_proto_rawDesc = []byte{
	0x0a, 0x39, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x69, 0x6e, 0x64,
	0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x69,
	0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x02, 0x0a, 0x20, 0x52, 0x75, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0xfa, 0x42, 0x21, 0x72, 0x1f, 0x10, 0x16, 0x18, 0xfe, 0x01,
	0x32, 0x18, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5f, 0x3a,
	0x5d, 0x7b, 0x32, 0x32, 0x2c, 0x32, 0x35, 0x34, 0x7d, 0x24, 0x52, 0x02, 0x69, 0x64, 0x12, 0x6a,
	0x0a, 0x18, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x52, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x47, 0x0a, 0x17, 0x73, 0x69,
	0x6d, 0x69, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x63,
	0x75, 0x74, 0x6f, 0x66, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0f, 0xfa, 0x42, 0x0c,
	0x0a, 0x0a, 0x1d, 0x00, 0x00, 0x80, 0x3f, 0x25, 0x00, 0x00, 0x00, 0x00, 0x52, 0x15, 0x73, 0x69,
	0x6d, 0x69, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x75, 0x74,
	0x6f, 0x66, 0x66, 0x22, 0xad, 0x01, 0x0a, 0x21, 0x52, 0x75, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0xfa, 0x42, 0x21, 0x72, 0x1f, 0x10, 0x16, 0x18, 0xfe,
	0x01, 0x32, 0x18, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5f,
	0x3a, 0x5d, 0x7b, 0x32, 0x32, 0x2c, 0x32, 0x35, 0x34, 0x7d, 0x24, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x3e, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65,
	0x74, 0x61, 0x67, 0x22, 0x5b, 0x0a, 0x23, 0x52, 0x65, 0x61, 0x64, 0x53, 0x75, 0x67, 0x67, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0xfa, 0x42, 0x21, 0x72, 0x1f, 0x10, 0x16, 0x18,
	0xfe, 0x01, 0x32, 0x18, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d,
	0x5f, 0x3a, 0x5d, 0x7b, 0x32, 0x32, 0x2c, 0x32, 0x35, 0x34, 0x7d, 0x24, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xae, 0x02, 0x0a, 0x24, 0x52, 0x65, 0x61, 0x64, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0xfa, 0x42, 0x21, 0x72, 0x1f, 0x10, 0x16, 0x18, 0xfe,
	0x01, 0x32, 0x18, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5f,
	0x3a, 0x5d, 0x7b, 0x32, 0x32, 0x2c, 0x32, 0x35, 0x34, 0x7d, 0x24, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x5d, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x69, 0x6e, 0x64,
	0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x71,
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2f, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x15, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x57, 0x0a, 0x1f, 0x52, 0x65, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x24, 0xfa, 0x42, 0x21, 0x72, 0x1f, 0x10, 0x16, 0x18, 0xfe, 0x01, 0x32, 0x18, 0x5e, 0x5b,
	0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5f, 0x3a, 0x5d, 0x7b, 0x32, 0x32,
	0x2c, 0x32, 0x35, 0x34, 0x7d, 0x24, 0x52, 0x02, 0x69, 0x64, 0x22, 0xac, 0x02, 0x0a, 0x20, 0x52,
	0x65, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0xfa, 0x42, 0x21,
	0x72, 0x1f, 0x10, 0x16, 0x18, 0xfe, 0x01, 0x32, 0x18, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d,
	0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5f, 0x3a, 0x5d, 0x7b, 0x32, 0x32, 0x2c, 0x32, 0x35, 0x34, 0x7d,
	0x24, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x12, 0x42, 0x0a, 0x0f, 0x75, 0x72, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x75, 0x72, 0x6c, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x6f, 0x0a, 0x16, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b,
	0x69, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x14, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x88, 0x04, 0x0a, 0x11, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x41, 0x50, 0x49, 0x12,
	0xa2, 0x01, 0x0a, 0x19, 0x52, 0x75, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x41, 0x2e,
	0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x52, 0x75, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x42, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0xab, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x61, 0x64, 0x53, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x44, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x53, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x69, 0x6e,
	0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x9f, 0x01, 0x0a, 0x18, 0x52, 0x65, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x40, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x41, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0xba, 0x02, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x64,
	0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x16, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x41, 0x70, 0x69, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x69, 0x6e, 0x64, 0x79,
	0x6b, 0x69, 0x74, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x49, 0x45, 0x58, 0xaa, 0x02, 0x1f, 0x49, 0x6e,
	0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1f,
	0x49, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2,
	0x02, 0x2b, 0x49, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21,
	0x49, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_indykite_entitymatching_v1beta1_entity_matching_api_proto_rawDescOnce sync.Once
	file_indykite_entitymatching_v1beta1_entity_matching_api_proto_rawDescData = file_indykite_entitymatching_v1beta1_entity_matching_api_proto_rawDesc
)

func file_indykite_entitymatching_v1beta1_entity_matching_api_proto_rawDescGZIP() []byte {
	file_indykite_entitymatching_v1beta1_entity_matching_api_proto_rawDescOnce.Do(func() {
		file_indykite_entitymatching_v1beta1_entity_matching_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_indykite_entitymatching_v1beta1_entity_matching_api_proto_rawDescData)
	})
	return file_indykite_entitymatching_v1beta1_entity_matching_api_proto_rawDescData
}

var file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_indykite_entitymatching_v1beta1_entity_matching_api_proto_goTypes = []any{
	(*RunEntityMatchingPipelineRequest)(nil),     // 0: indykite.entitymatching.v1beta1.RunEntityMatchingPipelineRequest
	(*RunEntityMatchingPipelineResponse)(nil),    // 1: indykite.entitymatching.v1beta1.RunEntityMatchingPipelineResponse
	(*ReadSuggestedPropertyMappingRequest)(nil),  // 2: indykite.entitymatching.v1beta1.ReadSuggestedPropertyMappingRequest
	(*ReadSuggestedPropertyMappingResponse)(nil), // 3: indykite.entitymatching.v1beta1.ReadSuggestedPropertyMappingResponse
	(*ReadEntityMatchingReportRequest)(nil),      // 4: indykite.entitymatching.v1beta1.ReadEntityMatchingReportRequest
	(*ReadEntityMatchingReportResponse)(nil),     // 5: indykite.entitymatching.v1beta1.ReadEntityMatchingReportResponse
	(*PropertyMapping)(nil),                      // 6: indykite.entitymatching.v1beta1.PropertyMapping
	(*timestamppb.Timestamp)(nil),                // 7: google.protobuf.Timestamp
	(PipelineStatus)(0),                          // 8: indykite.entitymatching.v1beta1.PipelineStatus
}
var file_indykite_entitymatching_v1beta1_entity_matching_api_proto_depIdxs = []int32{
	6, // 0: indykite.entitymatching.v1beta1.RunEntityMatchingPipelineRequest.custom_property_mappings:type_name -> indykite.entitymatching.v1beta1.PropertyMapping
	7, // 1: indykite.entitymatching.v1beta1.RunEntityMatchingPipelineResponse.last_run_time:type_name -> google.protobuf.Timestamp
	6, // 2: indykite.entitymatching.v1beta1.ReadSuggestedPropertyMappingResponse.property_mappings:type_name -> indykite.entitymatching.v1beta1.PropertyMapping
	8, // 3: indykite.entitymatching.v1beta1.ReadSuggestedPropertyMappingResponse.property_mapping_status:type_name -> indykite.entitymatching.v1beta1.PipelineStatus
	7, // 4: indykite.entitymatching.v1beta1.ReadEntityMatchingReportResponse.url_expire_time:type_name -> google.protobuf.Timestamp
	8, // 5: indykite.entitymatching.v1beta1.ReadEntityMatchingReportResponse.entity_matching_status:type_name -> indykite.entitymatching.v1beta1.PipelineStatus
	0, // 6: indykite.entitymatching.v1beta1.EntityMatchingAPI.RunEntityMatchingPipeline:input_type -> indykite.entitymatching.v1beta1.RunEntityMatchingPipelineRequest
	2, // 7: indykite.entitymatching.v1beta1.EntityMatchingAPI.ReadSuggestedPropertyMapping:input_type -> indykite.entitymatching.v1beta1.ReadSuggestedPropertyMappingRequest
	4, // 8: indykite.entitymatching.v1beta1.EntityMatchingAPI.ReadEntityMatchingReport:input_type -> indykite.entitymatching.v1beta1.ReadEntityMatchingReportRequest
	1, // 9: indykite.entitymatching.v1beta1.EntityMatchingAPI.RunEntityMatchingPipeline:output_type -> indykite.entitymatching.v1beta1.RunEntityMatchingPipelineResponse
	3, // 10: indykite.entitymatching.v1beta1.EntityMatchingAPI.ReadSuggestedPropertyMapping:output_type -> indykite.entitymatching.v1beta1.ReadSuggestedPropertyMappingResponse
	5, // 11: indykite.entitymatching.v1beta1.EntityMatchingAPI.ReadEntityMatchingReport:output_type -> indykite.entitymatching.v1beta1.ReadEntityMatchingReportResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_indykite_entitymatching_v1beta1_entity_matching_api_proto_init() }
func file_indykite_entitymatching_v1beta1_entity_matching_api_proto_init() {
	if File_indykite_entitymatching_v1beta1_entity_matching_api_proto != nil {
		return
	}
	file_indykite_entitymatching_v1beta1_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RunEntityMatchingPipelineRequest); i {
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
		file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RunEntityMatchingPipelineResponse); i {
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
		file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ReadSuggestedPropertyMappingRequest); i {
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
		file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ReadSuggestedPropertyMappingResponse); i {
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
		file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ReadEntityMatchingReportRequest); i {
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
		file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ReadEntityMatchingReportResponse); i {
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
			RawDescriptor: file_indykite_entitymatching_v1beta1_entity_matching_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_indykite_entitymatching_v1beta1_entity_matching_api_proto_goTypes,
		DependencyIndexes: file_indykite_entitymatching_v1beta1_entity_matching_api_proto_depIdxs,
		MessageInfos:      file_indykite_entitymatching_v1beta1_entity_matching_api_proto_msgTypes,
	}.Build()
	File_indykite_entitymatching_v1beta1_entity_matching_api_proto = out.File
	file_indykite_entitymatching_v1beta1_entity_matching_api_proto_rawDesc = nil
	file_indykite_entitymatching_v1beta1_entity_matching_api_proto_goTypes = nil
	file_indykite_entitymatching_v1beta1_entity_matching_api_proto_depIdxs = nil
}
