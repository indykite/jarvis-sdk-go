// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: indykite/knowledge/objects/v1beta1/ikg.proto

package objectsv1beta1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Node with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Node) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Node with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in NodeMultiError, or nil if none found.
func (m *Node) ValidateAll() error {
	return m.validate(true)
}

func (m *Node) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() != "" {

		if l := utf8.RuneCountInString(m.GetId()); l < 22 || l > 256 {
			err := NodeValidationError{
				field:  "Id",
				reason: "value length must be between 22 and 256 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if !strings.HasPrefix(m.GetId(), "gid:") {
			err := NodeValidationError{
				field:  "Id",
				reason: "value does not have prefix \"gid:\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if l := utf8.RuneCountInString(m.GetExternalId()); l < 1 || l > 256 {
		err := NodeValidationError{
			field:  "ExternalId",
			reason: "value length must be between 1 and 256 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetType()) > 64 {
		err := NodeValidationError{
			field:  "Type",
			reason: "value length must be at most 64 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_Node_Type_Pattern.MatchString(m.GetType()) {
		err := NodeValidationError{
			field:  "Type",
			reason: "value does not match regex pattern \"^([A-Z][a-z]+)+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetTags()) > 32 {
		err := NodeValidationError{
			field:  "Tags",
			reason: "value must contain no more than 32 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	_Node_Tags_Unique := make(map[string]struct{}, len(m.GetTags()))

	for idx, item := range m.GetTags() {
		_, _ = idx, item

		if _, exists := _Node_Tags_Unique[item]; exists {
			err := NodeValidationError{
				field:  fmt.Sprintf("Tags[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_Node_Tags_Unique[item] = struct{}{}
		}

		if utf8.RuneCountInString(item) > 64 {
			err := NodeValidationError{
				field:  fmt.Sprintf("Tags[%v]", idx),
				reason: "value length must be at most 64 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if !_Node_Tags_Pattern.MatchString(item) {
			err := NodeValidationError{
				field:  fmt.Sprintf("Tags[%v]", idx),
				reason: "value does not match regex pattern \"^([A-Z][a-z]+)+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if all {
		switch v := interface{}(m.GetCreateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NodeValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NodeValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NodeValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NodeValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeValidationError{
				field:  "UpdateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetProperties()) > 50 {
		err := NodeValidationError{
			field:  "Properties",
			reason: "value must contain no more than 50 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetProperties() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NodeValidationError{
						field:  fmt.Sprintf("Properties[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NodeValidationError{
						field:  fmt.Sprintf("Properties[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeValidationError{
					field:  fmt.Sprintf("Properties[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for IsIdentity

	if len(errors) > 0 {
		return NodeMultiError(errors)
	}

	return nil
}

// NodeMultiError is an error wrapping multiple validation errors returned by
// Node.ValidateAll() if the designated constraints aren't met.
type NodeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NodeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NodeMultiError) AllErrors() []error { return m }

// NodeValidationError is the validation error returned by Node.Validate if the
// designated constraints aren't met.
type NodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeValidationError) ErrorName() string { return "NodeValidationError" }

// Error satisfies the builtin error interface
func (e NodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeValidationError{}

var _Node_Type_Pattern = regexp.MustCompile("^([A-Z][a-z]+)+$")

var _Node_Tags_Pattern = regexp.MustCompile("^([A-Z][a-z]+)+$")

// Validate checks the field values on Relationship with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Relationship) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Relationship with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RelationshipMultiError, or
// nil if none found.
func (m *Relationship) ValidateAll() error {
	return m.validate(true)
}

func (m *Relationship) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Type

	// no validation rules for Source

	// no validation rules for Target

	if all {
		switch v := interface{}(m.GetCreateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RelationshipValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RelationshipValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RelationshipValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RelationshipValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RelationshipValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RelationshipValidationError{
				field:  "UpdateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	{
		sorted_keys := make([]string, len(m.GetProperties()))
		i := 0
		for key := range m.GetProperties() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetProperties()[key]
			_ = val

			// no validation rules for Properties[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, RelationshipValidationError{
							field:  fmt.Sprintf("Properties[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, RelationshipValidationError{
							field:  fmt.Sprintf("Properties[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return RelationshipValidationError{
						field:  fmt.Sprintf("Properties[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return RelationshipMultiError(errors)
	}

	return nil
}

// RelationshipMultiError is an error wrapping multiple validation errors
// returned by Relationship.ValidateAll() if the designated constraints aren't met.
type RelationshipMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RelationshipMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RelationshipMultiError) AllErrors() []error { return m }

// RelationshipValidationError is the validation error returned by
// Relationship.Validate if the designated constraints aren't met.
type RelationshipValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RelationshipValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RelationshipValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RelationshipValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RelationshipValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RelationshipValidationError) ErrorName() string { return "RelationshipValidationError" }

// Error satisfies the builtin error interface
func (e RelationshipValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRelationship.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RelationshipValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RelationshipValidationError{}

// Validate checks the field values on Property with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Property) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Property with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PropertyMultiError, or nil
// if none found.
func (m *Property) ValidateAll() error {
	return m.validate(true)
}

func (m *Property) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetType()) > 256 {
		err := PropertyValidationError{
			field:  "Type",
			reason: "value length must be at most 256 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_Property_Type_Pattern.MatchString(m.GetType()) {
		err := PropertyValidationError{
			field:  "Type",
			reason: "value does not match regex pattern \"^[a-zA-Z_][a-zA-Z0-9_]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetValue()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PropertyValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PropertyValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PropertyValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PropertyValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PropertyValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PropertyValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PropertyMultiError(errors)
	}

	return nil
}

// PropertyMultiError is an error wrapping multiple validation errors returned
// by Property.ValidateAll() if the designated constraints aren't met.
type PropertyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PropertyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PropertyMultiError) AllErrors() []error { return m }

// PropertyValidationError is the validation error returned by
// Property.Validate if the designated constraints aren't met.
type PropertyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PropertyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PropertyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PropertyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PropertyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PropertyValidationError) ErrorName() string { return "PropertyValidationError" }

// Error satisfies the builtin error interface
func (e PropertyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProperty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PropertyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PropertyValidationError{}

var _Property_Type_Pattern = regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]+$")

// Validate checks the field values on Metadata with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Metadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Metadata with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MetadataMultiError, or nil
// if none found.
func (m *Metadata) ValidateAll() error {
	return m.validate(true)
}

func (m *Metadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetAssuranceLevel() != 0 {

		if _, ok := _Metadata_AssuranceLevel_InLookup[m.GetAssuranceLevel()]; !ok {
			err := MetadataValidationError{
				field:  "AssuranceLevel",
				reason: "value must be in list [1 2 3]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if all {
		switch v := interface{}(m.GetVerificationTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MetadataValidationError{
					field:  "VerificationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MetadataValidationError{
					field:  "VerificationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetVerificationTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetadataValidationError{
				field:  "VerificationTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetSource()) > 128 {
		err := MetadataValidationError{
			field:  "Source",
			reason: "value length must be at most 128 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	{
		sorted_keys := make([]string, len(m.GetCustomMetadata()))
		i := 0
		for key := range m.GetCustomMetadata() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetCustomMetadata()[key]
			_ = val

			// no validation rules for CustomMetadata[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, MetadataValidationError{
							field:  fmt.Sprintf("CustomMetadata[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, MetadataValidationError{
							field:  fmt.Sprintf("CustomMetadata[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return MetadataValidationError{
						field:  fmt.Sprintf("CustomMetadata[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return MetadataMultiError(errors)
	}

	return nil
}

// MetadataMultiError is an error wrapping multiple validation errors returned
// by Metadata.ValidateAll() if the designated constraints aren't met.
type MetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetadataMultiError) AllErrors() []error { return m }

// MetadataValidationError is the validation error returned by
// Metadata.Validate if the designated constraints aren't met.
type MetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetadataValidationError) ErrorName() string { return "MetadataValidationError" }

// Error satisfies the builtin error interface
func (e MetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetadataValidationError{}

var _Metadata_AssuranceLevel_InLookup = map[int32]struct{}{
	1: {},
	2: {},
	3: {},
}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofUserPresent := false
	switch v := m.User.(type) {
	case *User_UserId:
		if v == nil {
			err := UserValidationError{
				field:  "User",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofUserPresent = true

		if l := utf8.RuneCountInString(m.GetUserId()); l < 22 || l > 254 {
			err := UserValidationError{
				field:  "UserId",
				reason: "value length must be between 22 and 254 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if !_User_UserId_Pattern.MatchString(m.GetUserId()) {
			err := UserValidationError{
				field:  "UserId",
				reason: "value does not match regex pattern \"^[A-Za-z0-9-_:]{22,254}$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *User_Property_:
		if v == nil {
			err := UserValidationError{
				field:  "User",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofUserPresent = true

		if m.GetProperty() == nil {
			err := UserValidationError{
				field:  "Property",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetProperty()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserValidationError{
						field:  "Property",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserValidationError{
						field:  "Property",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetProperty()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserValidationError{
					field:  "Property",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *User_ExternalId:
		if v == nil {
			err := UserValidationError{
				field:  "User",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofUserPresent = true

		if m.GetExternalId() == nil {
			err := UserValidationError{
				field:  "ExternalId",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetExternalId()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserValidationError{
						field:  "ExternalId",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserValidationError{
						field:  "ExternalId",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetExternalId()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserValidationError{
					field:  "ExternalId",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofUserPresent {
		err := UserValidationError{
			field:  "User",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

var _User_UserId_Pattern = regexp.MustCompile("^[A-Za-z0-9-_:]{22,254}$")

// Validate checks the field values on User_ExternalID with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *User_ExternalID) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User_ExternalID with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// User_ExternalIDMultiError, or nil if none found.
func (m *User_ExternalID) ValidateAll() error {
	return m.validate(true)
}

func (m *User_ExternalID) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetType()) > 64 {
		err := User_ExternalIDValidationError{
			field:  "Type",
			reason: "value length must be at most 64 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_User_ExternalID_Type_Pattern.MatchString(m.GetType()) {
		err := User_ExternalIDValidationError{
			field:  "Type",
			reason: "value does not match regex pattern \"^[a-zA-Z]*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetExternalId()) < 1 {
		err := User_ExternalIDValidationError{
			field:  "ExternalId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return User_ExternalIDMultiError(errors)
	}

	return nil
}

// User_ExternalIDMultiError is an error wrapping multiple validation errors
// returned by User_ExternalID.ValidateAll() if the designated constraints
// aren't met.
type User_ExternalIDMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m User_ExternalIDMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m User_ExternalIDMultiError) AllErrors() []error { return m }

// User_ExternalIDValidationError is the validation error returned by
// User_ExternalID.Validate if the designated constraints aren't met.
type User_ExternalIDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e User_ExternalIDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e User_ExternalIDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e User_ExternalIDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e User_ExternalIDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e User_ExternalIDValidationError) ErrorName() string { return "User_ExternalIDValidationError" }

// Error satisfies the builtin error interface
func (e User_ExternalIDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser_ExternalID.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = User_ExternalIDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = User_ExternalIDValidationError{}

var _User_ExternalID_Type_Pattern = regexp.MustCompile("^[a-zA-Z]*$")

// Validate checks the field values on User_Property with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *User_Property) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User_Property with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in User_PropertyMultiError, or
// nil if none found.
func (m *User_Property) ValidateAll() error {
	return m.validate(true)
}

func (m *User_Property) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetType()); l < 2 || l > 20 {
		err := User_PropertyValidationError{
			field:  "Type",
			reason: "value length must be between 2 and 20 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetValue() == nil {
		err := User_PropertyValidationError{
			field:  "Value",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetValue()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, User_PropertyValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, User_PropertyValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return User_PropertyValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return User_PropertyMultiError(errors)
	}

	return nil
}

// User_PropertyMultiError is an error wrapping multiple validation errors
// returned by User_Property.ValidateAll() if the designated constraints
// aren't met.
type User_PropertyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m User_PropertyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m User_PropertyMultiError) AllErrors() []error { return m }

// User_PropertyValidationError is the validation error returned by
// User_Property.Validate if the designated constraints aren't met.
type User_PropertyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e User_PropertyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e User_PropertyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e User_PropertyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e User_PropertyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e User_PropertyValidationError) ErrorName() string { return "User_PropertyValidationError" }

// Error satisfies the builtin error interface
func (e User_PropertyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser_Property.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = User_PropertyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = User_PropertyValidationError{}
