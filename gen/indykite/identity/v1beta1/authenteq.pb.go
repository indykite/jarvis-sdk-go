// Copyright (c) 2021 IndyKite
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
// source: indykite/identity/v1beta1/authenteq.proto

package identityv1beta1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthenteqDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DocumentData *AuthenteqDocumentData `protobuf:"bytes,4,opt,name=document_data,json=documentData,proto3" json:"document_data,omitempty"`
}

func (x *AuthenteqDetails) Reset() {
	*x = AuthenteqDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_identity_v1beta1_authenteq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenteqDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenteqDetails) ProtoMessage() {}

func (x *AuthenteqDetails) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_identity_v1beta1_authenteq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenteqDetails.ProtoReflect.Descriptor instead.
func (*AuthenteqDetails) Descriptor() ([]byte, []int) {
	return file_indykite_identity_v1beta1_authenteq_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenteqDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuthenteqDetails) GetDocumentData() *AuthenteqDocumentData {
	if x != nil {
		return x.DocumentData
	}
	return nil
}

type AuthenteqDocumentData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentType         string                          `protobuf:"bytes,1,opt,name=document_type,json=documentType,proto3" json:"document_type,omitempty"`
	DocumentNumber       string                          `protobuf:"bytes,2,opt,name=document_number,json=documentNumber,proto3" json:"document_number,omitempty"`
	IssuingCountry       string                          `protobuf:"bytes,3,opt,name=issuing_country,json=issuingCountry,proto3" json:"issuing_country,omitempty"`
	Jurisdiction         string                          `protobuf:"bytes,4,opt,name=jurisdiction,proto3" json:"jurisdiction,omitempty"`
	Nationality          string                          `protobuf:"bytes,5,opt,name=nationality,proto3" json:"nationality,omitempty"`
	SurnameAndGivenNames string                          `protobuf:"bytes,6,opt,name=surname_and_given_names,json=surnameAndGivenNames,proto3" json:"surname_and_given_names,omitempty"`
	Surname              string                          `protobuf:"bytes,7,opt,name=surname,proto3" json:"surname,omitempty"`
	GivenNames           string                          `protobuf:"bytes,8,opt,name=given_names,json=givenNames,proto3" json:"given_names,omitempty"`
	NameSuffixes         string                          `protobuf:"bytes,9,opt,name=name_suffixes,json=nameSuffixes,proto3" json:"name_suffixes,omitempty"`
	NamePrefixes         string                          `protobuf:"bytes,10,opt,name=name_prefixes,json=namePrefixes,proto3" json:"name_prefixes,omitempty"`
	Sex                  string                          `protobuf:"bytes,11,opt,name=sex,proto3" json:"sex,omitempty"`
	DateOfBirth          string                          `protobuf:"bytes,12,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	DateOfExpiry         string                          `protobuf:"bytes,13,opt,name=date_of_expiry,json=dateOfExpiry,proto3" json:"date_of_expiry,omitempty"`
	DateOfIssue          string                          `protobuf:"bytes,14,opt,name=date_of_issue,json=dateOfIssue,proto3" json:"date_of_issue,omitempty"`
	LicenseClass         string                          `protobuf:"bytes,15,opt,name=license_class,json=licenseClass,proto3" json:"license_class,omitempty"`
	LicenseClassDetails  map[string]*LicenseClassDetails `protobuf:"bytes,16,rep,name=license_class_details,json=licenseClassDetails,proto3" json:"license_class_details,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CroppedFrontImage    *WebImage                       `protobuf:"bytes,17,opt,name=cropped_front_image,json=croppedFrontImage,proto3" json:"cropped_front_image,omitempty"`
	CroppedBackImage     *WebImage                       `protobuf:"bytes,18,opt,name=cropped_back_image,json=croppedBackImage,proto3" json:"cropped_back_image,omitempty"`
	FaceImage            *WebImage                       `protobuf:"bytes,19,opt,name=face_image,json=faceImage,proto3" json:"face_image,omitempty"`
}

func (x *AuthenteqDocumentData) Reset() {
	*x = AuthenteqDocumentData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_identity_v1beta1_authenteq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenteqDocumentData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenteqDocumentData) ProtoMessage() {}

func (x *AuthenteqDocumentData) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_identity_v1beta1_authenteq_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenteqDocumentData.ProtoReflect.Descriptor instead.
func (*AuthenteqDocumentData) Descriptor() ([]byte, []int) {
	return file_indykite_identity_v1beta1_authenteq_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenteqDocumentData) GetDocumentType() string {
	if x != nil {
		return x.DocumentType
	}
	return ""
}

func (x *AuthenteqDocumentData) GetDocumentNumber() string {
	if x != nil {
		return x.DocumentNumber
	}
	return ""
}

func (x *AuthenteqDocumentData) GetIssuingCountry() string {
	if x != nil {
		return x.IssuingCountry
	}
	return ""
}

func (x *AuthenteqDocumentData) GetJurisdiction() string {
	if x != nil {
		return x.Jurisdiction
	}
	return ""
}

func (x *AuthenteqDocumentData) GetNationality() string {
	if x != nil {
		return x.Nationality
	}
	return ""
}

func (x *AuthenteqDocumentData) GetSurnameAndGivenNames() string {
	if x != nil {
		return x.SurnameAndGivenNames
	}
	return ""
}

func (x *AuthenteqDocumentData) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *AuthenteqDocumentData) GetGivenNames() string {
	if x != nil {
		return x.GivenNames
	}
	return ""
}

func (x *AuthenteqDocumentData) GetNameSuffixes() string {
	if x != nil {
		return x.NameSuffixes
	}
	return ""
}

func (x *AuthenteqDocumentData) GetNamePrefixes() string {
	if x != nil {
		return x.NamePrefixes
	}
	return ""
}

func (x *AuthenteqDocumentData) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *AuthenteqDocumentData) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *AuthenteqDocumentData) GetDateOfExpiry() string {
	if x != nil {
		return x.DateOfExpiry
	}
	return ""
}

func (x *AuthenteqDocumentData) GetDateOfIssue() string {
	if x != nil {
		return x.DateOfIssue
	}
	return ""
}

func (x *AuthenteqDocumentData) GetLicenseClass() string {
	if x != nil {
		return x.LicenseClass
	}
	return ""
}

func (x *AuthenteqDocumentData) GetLicenseClassDetails() map[string]*LicenseClassDetails {
	if x != nil {
		return x.LicenseClassDetails
	}
	return nil
}

func (x *AuthenteqDocumentData) GetCroppedFrontImage() *WebImage {
	if x != nil {
		return x.CroppedFrontImage
	}
	return nil
}

func (x *AuthenteqDocumentData) GetCroppedBackImage() *WebImage {
	if x != nil {
		return x.CroppedBackImage
	}
	return nil
}

func (x *AuthenteqDocumentData) GetFaceImage() *WebImage {
	if x != nil {
		return x.FaceImage
	}
	return nil
}

type LicenseClassDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To    string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Notes string `protobuf:"bytes,3,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *LicenseClassDetails) Reset() {
	*x = LicenseClassDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_identity_v1beta1_authenteq_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseClassDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseClassDetails) ProtoMessage() {}

func (x *LicenseClassDetails) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_identity_v1beta1_authenteq_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseClassDetails.ProtoReflect.Descriptor instead.
func (*LicenseClassDetails) Descriptor() ([]byte, []int) {
	return file_indykite_identity_v1beta1_authenteq_proto_rawDescGZIP(), []int{2}
}

func (x *LicenseClassDetails) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *LicenseClassDetails) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *LicenseClassDetails) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

type WebImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Content     string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *WebImage) Reset() {
	*x = WebImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indykite_identity_v1beta1_authenteq_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebImage) ProtoMessage() {}

func (x *WebImage) ProtoReflect() protoreflect.Message {
	mi := &file_indykite_identity_v1beta1_authenteq_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebImage.ProtoReflect.Descriptor instead.
func (*WebImage) Descriptor() ([]byte, []int) {
	return file_indykite_identity_v1beta1_authenteq_proto_rawDescGZIP(), []int{3}
}

func (x *WebImage) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *WebImage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_indykite_identity_v1beta1_authenteq_proto protoreflect.FileDescriptor

var file_indykite_identity_v1beta1_authenteq_proto_rawDesc = []byte{
	0x0a, 0x29, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69, 0x6e, 0x64,
	0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x22, 0x79, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x65, 0x71, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x55, 0x0a, 0x0d, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x65, 0x71, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x98, 0x08, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x65, 0x71, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x73, 0x73,
	0x75, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6a, 0x75, 0x72, 0x69, 0x73, 0x64, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6a, 0x75, 0x72, 0x69, 0x73, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x17, 0x73, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x41, 0x6e, 0x64, 0x47, 0x69, 0x76, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x69, 0x76,
	0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x67, 0x69, 0x76, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x61,
	0x6d, 0x65, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x65, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f,
	0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79,
	0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x7d, 0x0a, 0x15, 0x6c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b,
	0x69, 0x74, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x65, 0x71, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x13, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x53, 0x0a, 0x13, 0x63, 0x72, 0x6f, 0x70,
	0x70, 0x65, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x57, 0x65, 0x62, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x11, 0x63, 0x72, 0x6f, 0x70,
	0x70, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x51, 0x0a,
	0x12, 0x63, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x64, 0x79,
	0x6b, 0x69, 0x74, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x10,
	0x63, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x42, 0x0a, 0x0a, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x57, 0x65, 0x62, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x09, 0x66, 0x61, 0x63, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x1a, 0x76, 0x0a, 0x18, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x44, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4f, 0x0a, 0x13,
	0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x47, 0x0a,
	0x08, 0x57, 0x65, 0x62, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x88, 0x02, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x65, 0x71, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f,
	0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x69, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03,
	0x49, 0x49, 0x58, 0xaa, 0x02, 0x19, 0x49, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca,
	0x02, 0x19, 0x49, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x25, 0x49, 0x6e,
	0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x49, 0x6e, 0x64, 0x79, 0x6b, 0x69, 0x74, 0x65, 0x3a, 0x3a,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_indykite_identity_v1beta1_authenteq_proto_rawDescOnce sync.Once
	file_indykite_identity_v1beta1_authenteq_proto_rawDescData = file_indykite_identity_v1beta1_authenteq_proto_rawDesc
)

func file_indykite_identity_v1beta1_authenteq_proto_rawDescGZIP() []byte {
	file_indykite_identity_v1beta1_authenteq_proto_rawDescOnce.Do(func() {
		file_indykite_identity_v1beta1_authenteq_proto_rawDescData = protoimpl.X.CompressGZIP(file_indykite_identity_v1beta1_authenteq_proto_rawDescData)
	})
	return file_indykite_identity_v1beta1_authenteq_proto_rawDescData
}

var file_indykite_identity_v1beta1_authenteq_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_indykite_identity_v1beta1_authenteq_proto_goTypes = []interface{}{
	(*AuthenteqDetails)(nil),      // 0: indykite.identity.v1beta1.AuthenteqDetails
	(*AuthenteqDocumentData)(nil), // 1: indykite.identity.v1beta1.AuthenteqDocumentData
	(*LicenseClassDetails)(nil),   // 2: indykite.identity.v1beta1.LicenseClassDetails
	(*WebImage)(nil),              // 3: indykite.identity.v1beta1.WebImage
	nil,                           // 4: indykite.identity.v1beta1.AuthenteqDocumentData.LicenseClassDetailsEntry
}
var file_indykite_identity_v1beta1_authenteq_proto_depIdxs = []int32{
	1, // 0: indykite.identity.v1beta1.AuthenteqDetails.document_data:type_name -> indykite.identity.v1beta1.AuthenteqDocumentData
	4, // 1: indykite.identity.v1beta1.AuthenteqDocumentData.license_class_details:type_name -> indykite.identity.v1beta1.AuthenteqDocumentData.LicenseClassDetailsEntry
	3, // 2: indykite.identity.v1beta1.AuthenteqDocumentData.cropped_front_image:type_name -> indykite.identity.v1beta1.WebImage
	3, // 3: indykite.identity.v1beta1.AuthenteqDocumentData.cropped_back_image:type_name -> indykite.identity.v1beta1.WebImage
	3, // 4: indykite.identity.v1beta1.AuthenteqDocumentData.face_image:type_name -> indykite.identity.v1beta1.WebImage
	2, // 5: indykite.identity.v1beta1.AuthenteqDocumentData.LicenseClassDetailsEntry.value:type_name -> indykite.identity.v1beta1.LicenseClassDetails
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_indykite_identity_v1beta1_authenteq_proto_init() }
func file_indykite_identity_v1beta1_authenteq_proto_init() {
	if File_indykite_identity_v1beta1_authenteq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_indykite_identity_v1beta1_authenteq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenteqDetails); i {
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
		file_indykite_identity_v1beta1_authenteq_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenteqDocumentData); i {
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
		file_indykite_identity_v1beta1_authenteq_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseClassDetails); i {
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
		file_indykite_identity_v1beta1_authenteq_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebImage); i {
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
			RawDescriptor: file_indykite_identity_v1beta1_authenteq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_indykite_identity_v1beta1_authenteq_proto_goTypes,
		DependencyIndexes: file_indykite_identity_v1beta1_authenteq_proto_depIdxs,
		MessageInfos:      file_indykite_identity_v1beta1_authenteq_proto_msgTypes,
	}.Build()
	File_indykite_identity_v1beta1_authenteq_proto = out.File
	file_indykite_identity_v1beta1_authenteq_proto_rawDesc = nil
	file_indykite_identity_v1beta1_authenteq_proto_goTypes = nil
	file_indykite_identity_v1beta1_authenteq_proto_depIdxs = nil
}
