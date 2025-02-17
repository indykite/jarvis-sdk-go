// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: indykite/objects/v1beta2/value.proto

package objectsv1beta2

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

// Validate checks the field values on Value with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Value) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Value with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ValueMultiError, or nil if none found.
func (m *Value) ValidateAll() error {
	return m.validate(true)
}

func (m *Value) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofTypePresent := false
	switch v := m.Type.(type) {
	case *Value_BoolValue:
		if v == nil {
			err := ValueValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true
		// no validation rules for BoolValue
	case *Value_IntegerValue:
		if v == nil {
			err := ValueValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true
		// no validation rules for IntegerValue
	case *Value_DoubleValue:
		if v == nil {
			err := ValueValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true
		// no validation rules for DoubleValue
	case *Value_TimeValue:
		if v == nil {
			err := ValueValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if all {
			switch v := interface{}(m.GetTimeValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ValueValidationError{
						field:  "TimeValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ValueValidationError{
						field:  "TimeValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTimeValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ValueValidationError{
					field:  "TimeValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Value_DurationValue:
		if v == nil {
			err := ValueValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if all {
			switch v := interface{}(m.GetDurationValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ValueValidationError{
						field:  "DurationValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ValueValidationError{
						field:  "DurationValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDurationValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ValueValidationError{
					field:  "DurationValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Value_StringValue:
		if v == nil {
			err := ValueValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true
		// no validation rules for StringValue
	case *Value_BytesValue:
		if v == nil {
			err := ValueValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true
		// no validation rules for BytesValue
	case *Value_ArrayValue:
		if v == nil {
			err := ValueValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if all {
			switch v := interface{}(m.GetArrayValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ValueValidationError{
						field:  "ArrayValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ValueValidationError{
						field:  "ArrayValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetArrayValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ValueValidationError{
					field:  "ArrayValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Value_MapValue:
		if v == nil {
			err := ValueValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if all {
			switch v := interface{}(m.GetMapValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ValueValidationError{
						field:  "MapValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ValueValidationError{
						field:  "MapValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetMapValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ValueValidationError{
					field:  "MapValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofTypePresent {
		err := ValueValidationError{
			field:  "Type",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ValueMultiError(errors)
	}

	return nil
}

// ValueMultiError is an error wrapping multiple validation errors returned by
// Value.ValidateAll() if the designated constraints aren't met.
type ValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ValueMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ValueMultiError) AllErrors() []error { return m }

// ValueValidationError is the validation error returned by Value.Validate if
// the designated constraints aren't met.
type ValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValueValidationError) ErrorName() string { return "ValueValidationError" }

// Error satisfies the builtin error interface
func (e ValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValueValidationError{}

// Validate checks the field values on Array with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Array) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Array with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ArrayMultiError, or nil if none found.
func (m *Array) ValidateAll() error {
	return m.validate(true)
}

func (m *Array) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetValues() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ArrayValidationError{
						field:  fmt.Sprintf("Values[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ArrayValidationError{
						field:  fmt.Sprintf("Values[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ArrayValidationError{
					field:  fmt.Sprintf("Values[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ArrayMultiError(errors)
	}

	return nil
}

// ArrayMultiError is an error wrapping multiple validation errors returned by
// Array.ValidateAll() if the designated constraints aren't met.
type ArrayMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ArrayMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ArrayMultiError) AllErrors() []error { return m }

// ArrayValidationError is the validation error returned by Array.Validate if
// the designated constraints aren't met.
type ArrayValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArrayValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArrayValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArrayValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArrayValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArrayValidationError) ErrorName() string { return "ArrayValidationError" }

// Error satisfies the builtin error interface
func (e ArrayValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArray.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArrayValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArrayValidationError{}

// Validate checks the field values on Map with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Map) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Map with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MapMultiError, or nil if none found.
func (m *Map) ValidateAll() error {
	return m.validate(true)
}

func (m *Map) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]string, len(m.GetFields()))
		i := 0
		for key := range m.GetFields() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetFields()[key]
			_ = val

			// no validation rules for Fields[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, MapValidationError{
							field:  fmt.Sprintf("Fields[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, MapValidationError{
							field:  fmt.Sprintf("Fields[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return MapValidationError{
						field:  fmt.Sprintf("Fields[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return MapMultiError(errors)
	}

	return nil
}

// MapMultiError is an error wrapping multiple validation errors returned by
// Map.ValidateAll() if the designated constraints aren't met.
type MapMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MapMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MapMultiError) AllErrors() []error { return m }

// MapValidationError is the validation error returned by Map.Validate if the
// designated constraints aren't met.
type MapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MapValidationError) ErrorName() string { return "MapValidationError" }

// Error satisfies the builtin error interface
func (e MapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MapValidationError{}
