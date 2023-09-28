// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: indykite/knowledge/v1beta1/model.proto

package knowledgev1beta1

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

// Validate checks the field values on Path with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Path) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Path with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PathMultiError, or nil if none found.
func (m *Path) ValidateAll() error {
	return m.validate(true)
}

func (m *Path) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetNodes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PathValidationError{
						field:  fmt.Sprintf("Nodes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PathValidationError{
						field:  fmt.Sprintf("Nodes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PathValidationError{
					field:  fmt.Sprintf("Nodes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetRelationships() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PathValidationError{
						field:  fmt.Sprintf("Relationships[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PathValidationError{
						field:  fmt.Sprintf("Relationships[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PathValidationError{
					field:  fmt.Sprintf("Relationships[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PathMultiError(errors)
	}

	return nil
}

// PathMultiError is an error wrapping multiple validation errors returned by
// Path.ValidateAll() if the designated constraints aren't met.
type PathMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PathMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PathMultiError) AllErrors() []error { return m }

// PathValidationError is the validation error returned by Path.Validate if the
// designated constraints aren't met.
type PathValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PathValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PathValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PathValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PathValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PathValidationError) ErrorName() string { return "PathValidationError" }

// Error satisfies the builtin error interface
func (e PathValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPath.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PathValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PathValidationError{}

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

	// no validation rules for Id

	// no validation rules for ExternalId

	// no validation rules for Type

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

	// no validation rules for Key

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

// Validate checks the field values on InputParam with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *InputParam) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InputParam with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in InputParamMultiError, or
// nil if none found.
func (m *InputParam) ValidateAll() error {
	return m.validate(true)
}

func (m *InputParam) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofValuePresent := false
	switch v := m.Value.(type) {
	case *InputParam_StringValue:
		if v == nil {
			err := InputParamValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofValuePresent = true

		if l := utf8.RuneCountInString(m.GetStringValue()); l < 1 || l > 50 {
			err := InputParamValidationError{
				field:  "StringValue",
				reason: "value length must be between 1 and 50 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *InputParam_BoolValue:
		if v == nil {
			err := InputParamValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofValuePresent = true
		// no validation rules for BoolValue
	case *InputParam_IntegerValue:
		if v == nil {
			err := InputParamValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofValuePresent = true
		// no validation rules for IntegerValue
	case *InputParam_DoubleValue:
		if v == nil {
			err := InputParamValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofValuePresent = true
		// no validation rules for DoubleValue
	default:
		_ = v // ensures v is used
	}
	if !oneofValuePresent {
		err := InputParamValidationError{
			field:  "Value",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return InputParamMultiError(errors)
	}

	return nil
}

// InputParamMultiError is an error wrapping multiple validation errors
// returned by InputParam.ValidateAll() if the designated constraints aren't met.
type InputParamMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InputParamMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InputParamMultiError) AllErrors() []error { return m }

// InputParamValidationError is the validation error returned by
// InputParam.Validate if the designated constraints aren't met.
type InputParamValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InputParamValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InputParamValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InputParamValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InputParamValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InputParamValidationError) ErrorName() string { return "InputParamValidationError" }

// Error satisfies the builtin error interface
func (e InputParamValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInputParam.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InputParamValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InputParamValidationError{}
