// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: neutral_diet/user/v1/food_item_log.proto

package userv1

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

// Validate checks the field values on FoodLogItem with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FoodLogItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FoodLogItem with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FoodLogItemMultiError, or
// nil if none found.
func (m *FoodLogItem) ValidateAll() error {
	return m.validate(true)
}

func (m *FoodLogItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FoodItemId

	// no validation rules for Weight

	// no validation rules for CarbonFootprint

	if len(errors) > 0 {
		return FoodLogItemMultiError(errors)
	}

	return nil
}

// FoodLogItemMultiError is an error wrapping multiple validation errors
// returned by FoodLogItem.ValidateAll() if the designated constraints aren't met.
type FoodLogItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FoodLogItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FoodLogItemMultiError) AllErrors() []error { return m }

// FoodLogItemValidationError is the validation error returned by
// FoodLogItem.Validate if the designated constraints aren't met.
type FoodLogItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FoodLogItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FoodLogItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FoodLogItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FoodLogItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FoodLogItemValidationError) ErrorName() string { return "FoodLogItemValidationError" }

// Error satisfies the builtin error interface
func (e FoodLogItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFoodLogItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FoodLogItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FoodLogItemValidationError{}