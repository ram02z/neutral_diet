// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: neutral_diet/food/v1/food_item.proto

package foodv1

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

// Validate checks the field values on FoodItem with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FoodItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FoodItem with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FoodItemMultiError, or nil
// if none found.
func (m *FoodItem) ValidateAll() error {
	return m.validate(true)
}

func (m *FoodItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for TypologyId

	if _, ok := FoodItem_CfType_name[int32(m.GetCfType())]; !ok {
		err := FoodItemValidationError{
			field:  "CfType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return FoodItemMultiError(errors)
	}

	return nil
}

// FoodItemMultiError is an error wrapping multiple validation errors returned
// by FoodItem.ValidateAll() if the designated constraints aren't met.
type FoodItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FoodItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FoodItemMultiError) AllErrors() []error { return m }

// FoodItemValidationError is the validation error returned by
// FoodItem.Validate if the designated constraints aren't met.
type FoodItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FoodItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FoodItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FoodItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FoodItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FoodItemValidationError) ErrorName() string { return "FoodItemValidationError" }

// Error satisfies the builtin error interface
func (e FoodItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFoodItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FoodItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FoodItemValidationError{}

// Validate checks the field values on FoodItemInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FoodItemInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FoodItemInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FoodItemInfoMultiError, or
// nil if none found.
func (m *FoodItemInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *FoodItemInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TypologyName

	// no validation rules for SubTypologyName

	// no validation rules for NonUniqueSources

	for idx, item := range m.GetSources() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FoodItemInfoValidationError{
						field:  fmt.Sprintf("Sources[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FoodItemInfoValidationError{
						field:  fmt.Sprintf("Sources[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FoodItemInfoValidationError{
					field:  fmt.Sprintf("Sources[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return FoodItemInfoMultiError(errors)
	}

	return nil
}

// FoodItemInfoMultiError is an error wrapping multiple validation errors
// returned by FoodItemInfo.ValidateAll() if the designated constraints aren't met.
type FoodItemInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FoodItemInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FoodItemInfoMultiError) AllErrors() []error { return m }

// FoodItemInfoValidationError is the validation error returned by
// FoodItemInfo.Validate if the designated constraints aren't met.
type FoodItemInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FoodItemInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FoodItemInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FoodItemInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FoodItemInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FoodItemInfoValidationError) ErrorName() string { return "FoodItemInfoValidationError" }

// Error satisfies the builtin error interface
func (e FoodItemInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFoodItemInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FoodItemInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FoodItemInfoValidationError{}

// Validate checks the field values on AggregateFoodItem with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AggregateFoodItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AggregateFoodItem with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AggregateFoodItemMultiError, or nil if none found.
func (m *AggregateFoodItem) ValidateAll() error {
	return m.validate(true)
}

func (m *AggregateFoodItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for FoodName

	// no validation rules for MedianCarbonFootprint

	// no validation rules for Region

	// no validation rules for TypologyName

	// no validation rules for SubTypologyName

	if len(errors) > 0 {
		return AggregateFoodItemMultiError(errors)
	}

	return nil
}

// AggregateFoodItemMultiError is an error wrapping multiple validation errors
// returned by AggregateFoodItem.ValidateAll() if the designated constraints
// aren't met.
type AggregateFoodItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AggregateFoodItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AggregateFoodItemMultiError) AllErrors() []error { return m }

// AggregateFoodItemValidationError is the validation error returned by
// AggregateFoodItem.Validate if the designated constraints aren't met.
type AggregateFoodItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AggregateFoodItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AggregateFoodItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AggregateFoodItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AggregateFoodItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AggregateFoodItemValidationError) ErrorName() string {
	return "AggregateFoodItemValidationError"
}

// Error satisfies the builtin error interface
func (e AggregateFoodItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAggregateFoodItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AggregateFoodItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AggregateFoodItemValidationError{}
