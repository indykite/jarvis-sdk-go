// Copyright (c) 2022 IndyKite
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
// source: indykite/ingest/v1beta1/model.proto

package ingestv1beta1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	v1beta1 "github.com/indykite/indykite-sdk-go/gen/indykite/objects/v1beta1"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the record, for client side reference
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Specifying which of the fields in the record data is the unique identifier for the entity
	ExternalId string                    `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	Data       map[string]*v1beta1.Value `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_ingest_v1beta1_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_ingest_v1beta1_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_indykite_ingest_v1beta1_model_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Record) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *Record) GetData() map[string]*v1beta1.Value {
	if x != nil {
		return x.Data
	}
	return nil
}

type PropertyError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []string `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *PropertyError) Reset() {
	*x = PropertyError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_ingest_v1beta1_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropertyError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropertyError) ProtoMessage() {}

func (x *PropertyError) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_ingest_v1beta1_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropertyError.ProtoReflect.Descriptor instead.
func (*PropertyError) Descriptor() ([]byte, []int) {
	return file_indykite_ingest_v1beta1_model_proto_rawDescGZIP(), []int{1}
}

func (x *PropertyError) GetMessages() []string {
	if x != nil {
		return x.Messages
	}
	return nil
}

type RecordError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// map from property type to property error, if any
	PropertyErrors map[string]*PropertyError `protobuf:"bytes,1,rep,name=property_errors,json=propertyErrors,proto3" json:"property_errors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RecordError) Reset() {
	*x = RecordError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_ingest_v1beta1_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordError) ProtoMessage() {}

func (x *RecordError) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_ingest_v1beta1_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordError.ProtoReflect.Descriptor instead.
func (*RecordError) Descriptor() ([]byte, []int) {
	return file_indykite_ingest_v1beta1_model_proto_rawDescGZIP(), []int{2}
}

func (x *RecordError) GetPropertyErrors() map[string]*PropertyError {
	if x != nil {
		return x.PropertyErrors
	}
	return nil
}

var File_indykite_ingest_v1beta1_model_proto protoreflect.FileDescriptor

var file_indykite_ingest_v1beta1_model_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e,
	0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x25,
	0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd,
	0x01, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80, 0x02, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10,
	0x01, 0x18, 0x80, 0x02, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64,
	0x12, 0x52, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x13, 0xfa, 0x42, 0x10, 0x9a, 0x01,
	0x0d, 0x08, 0x01, 0x10, 0x14, 0x22, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0x80, 0x02, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x58, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2b,
	0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0xdb, 0x01, 0x0a, 0x0b,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x61, 0x0a, 0x0f, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e,
	0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x69,
	0x0a, 0x13, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74,
	0x65, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xf6, 0x01, 0x0a, 0x1b, 0x63, 0x6f,
	0x6d, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x69, 0x6e, 0x64,
	0x79, 0x6b, 0x69, 0x74, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x49, 0x49, 0x58, 0xaa, 0x02, 0x17, 0x49,
	0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x17, 0x49, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74,
	0x65, 0x5c, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xe2, 0x02, 0x23, 0x49, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x5c, 0x49, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x49, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74,
	0x65, 0x3a, 0x3a, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_indykite_ingest_v1beta1_model_proto_rawDescOnce sync.Once
	file_indykite_ingest_v1beta1_model_proto_rawDescData = file_indykite_ingest_v1beta1_model_proto_rawDesc
)

func file_indykite_ingest_v1beta1_model_proto_rawDescGZIP() []byte {
	file_indykite_ingest_v1beta1_model_proto_rawDescOnce.Do(func() {
		file_indykite_ingest_v1beta1_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_indykite_ingest_v1beta1_model_proto_rawDescData)
	})
	return file_indykite_ingest_v1beta1_model_proto_rawDescData
}

var file_indykite_ingest_v1beta1_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_indykite_ingest_v1beta1_model_proto_goTypes = []interface{}{
	(*Record)(nil),        // 0: indykite.ingest.v1beta1.Record
	(*PropertyError)(nil), // 1: indykite.ingest.v1beta1.PropertyError
	(*RecordError)(nil),   // 2: indykite.ingest.v1beta1.RecordError
	nil,                   // 3: indykite.ingest.v1beta1.Record.DataEntry
	nil,                   // 4: indykite.ingest.v1beta1.RecordError.PropertyErrorsEntry
	(*v1beta1.Value)(nil), // 5: indykite.objects.v1beta1.Value
}
var file_indykite_ingest_v1beta1_model_proto_depIdxs = []int32{
	3, // 0: indykite.ingest.v1beta1.Record.data:type_name -> indykite.ingest.v1beta1.Record.DataEntry
	4, // 1: indykite.ingest.v1beta1.RecordError.property_errors:type_name -> indykite.ingest.v1beta1.RecordError.PropertyErrorsEntry
	5, // 2: indykite.ingest.v1beta1.Record.DataEntry.value:type_name -> indykite.objects.v1beta1.Value
	1, // 3: indykite.ingest.v1beta1.RecordError.PropertyErrorsEntry.value:type_name -> indykite.ingest.v1beta1.PropertyError
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_indykite_ingest_v1beta1_model_proto_init() }
func file_indykite_ingest_v1beta1_model_proto_init() {
	if File_indykite_ingest_v1beta1_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_indykite_ingest_v1beta1_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_indykite_ingest_v1beta1_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertyError); i {
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
		file_indykite_ingest_v1beta1_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordError); i {
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
			RawDescriptor: file_indykite_ingest_v1beta1_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_indykite_ingest_v1beta1_model_proto_goTypes,
		DependencyIndexes: file_indykite_ingest_v1beta1_model_proto_depIdxs,
		MessageInfos:      file_indykite_ingest_v1beta1_model_proto_msgTypes,
	}.Build()
	File_indykite_ingest_v1beta1_model_proto = out.File
	file_indykite_ingest_v1beta1_model_proto_rawDesc = nil
	file_indykite_ingest_v1beta1_model_proto_goTypes = nil
	file_indykite_ingest_v1beta1_model_proto_depIdxs = nil
}
