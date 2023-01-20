// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: neutral_diet/user/v1/api.proto

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

// Validate checks the field values on AddFoodItemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddFoodItemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddFoodItemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddFoodItemRequestMultiError, or nil if none found.
func (m *AddFoodItemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddFoodItemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetFoodLogItem() == nil {
		err := AddFoodItemRequestValidationError{
			field:  "FoodLogItem",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetFoodLogItem()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddFoodItemRequestValidationError{
					field:  "FoodLogItem",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddFoodItemRequestValidationError{
					field:  "FoodLogItem",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFoodLogItem()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddFoodItemRequestValidationError{
				field:  "FoodLogItem",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AddFoodItemRequestMultiError(errors)
	}

	return nil
}

// AddFoodItemRequestMultiError is an error wrapping multiple validation errors
// returned by AddFoodItemRequest.ValidateAll() if the designated constraints
// aren't met.
type AddFoodItemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddFoodItemRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddFoodItemRequestMultiError) AllErrors() []error { return m }

// AddFoodItemRequestValidationError is the validation error returned by
// AddFoodItemRequest.Validate if the designated constraints aren't met.
type AddFoodItemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddFoodItemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddFoodItemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddFoodItemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddFoodItemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddFoodItemRequestValidationError) ErrorName() string {
	return "AddFoodItemRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddFoodItemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddFoodItemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddFoodItemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddFoodItemRequestValidationError{}

// Validate checks the field values on AddFoodItemResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddFoodItemResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddFoodItemResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddFoodItemResponseMultiError, or nil if none found.
func (m *AddFoodItemResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AddFoodItemResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return AddFoodItemResponseMultiError(errors)
	}

	return nil
}

// AddFoodItemResponseMultiError is an error wrapping multiple validation
// errors returned by AddFoodItemResponse.ValidateAll() if the designated
// constraints aren't met.
type AddFoodItemResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddFoodItemResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddFoodItemResponseMultiError) AllErrors() []error { return m }

// AddFoodItemResponseValidationError is the validation error returned by
// AddFoodItemResponse.Validate if the designated constraints aren't met.
type AddFoodItemResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddFoodItemResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddFoodItemResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddFoodItemResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddFoodItemResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddFoodItemResponseValidationError) ErrorName() string {
	return "AddFoodItemResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AddFoodItemResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddFoodItemResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddFoodItemResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddFoodItemResponseValidationError{}

// Validate checks the field values on CreateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUserRequestMultiError, or nil if none found.
func (m *CreateUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetFirebaseUid()) < 1 {
		err := CreateUserRequestValidationError{
			field:  "FirebaseUid",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateUserRequestMultiError(errors)
	}

	return nil
}

// CreateUserRequestMultiError is an error wrapping multiple validation errors
// returned by CreateUserRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserRequestMultiError) AllErrors() []error { return m }

// CreateUserRequestValidationError is the validation error returned by
// CreateUserRequest.Validate if the designated constraints aren't met.
type CreateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserRequestValidationError) ErrorName() string {
	return "CreateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserRequestValidationError{}

// Validate checks the field values on CreateUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUserResponseMultiError, or nil if none found.
func (m *CreateUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateUserResponseMultiError(errors)
	}

	return nil
}

// CreateUserResponseMultiError is an error wrapping multiple validation errors
// returned by CreateUserResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserResponseMultiError) AllErrors() []error { return m }

// CreateUserResponseValidationError is the validation error returned by
// CreateUserResponse.Validate if the designated constraints aren't met.
type CreateUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserResponseValidationError) ErrorName() string {
	return "CreateUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserResponseValidationError{}

// Validate checks the field values on DeleteUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteUserRequestMultiError, or nil if none found.
func (m *DeleteUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteUserRequestMultiError(errors)
	}

	return nil
}

// DeleteUserRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteUserRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteUserRequestMultiError) AllErrors() []error { return m }

// DeleteUserRequestValidationError is the validation error returned by
// DeleteUserRequest.Validate if the designated constraints aren't met.
type DeleteUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUserRequestValidationError) ErrorName() string {
	return "DeleteUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUserRequestValidationError{}

// Validate checks the field values on DeleteUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteUserResponseMultiError, or nil if none found.
func (m *DeleteUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteUserResponseMultiError(errors)
	}

	return nil
}

// DeleteUserResponseMultiError is an error wrapping multiple validation errors
// returned by DeleteUserResponse.ValidateAll() if the designated constraints
// aren't met.
type DeleteUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteUserResponseMultiError) AllErrors() []error { return m }

// DeleteUserResponseValidationError is the validation error returned by
// DeleteUserResponse.Validate if the designated constraints aren't met.
type DeleteUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUserResponseValidationError) ErrorName() string {
	return "DeleteUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUserResponseValidationError{}

// Validate checks the field values on UpdateUserRegionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateUserRegionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUserRegionRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateUserRegionRequestMultiError, or nil if none found.
func (m *UpdateUserRegionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserRegionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRegion()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateUserRegionRequestValidationError{
					field:  "Region",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateUserRegionRequestValidationError{
					field:  "Region",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRegion()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateUserRegionRequestValidationError{
				field:  "Region",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateUserRegionRequestMultiError(errors)
	}

	return nil
}

// UpdateUserRegionRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateUserRegionRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateUserRegionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserRegionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserRegionRequestMultiError) AllErrors() []error { return m }

// UpdateUserRegionRequestValidationError is the validation error returned by
// UpdateUserRegionRequest.Validate if the designated constraints aren't met.
type UpdateUserRegionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserRegionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserRegionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserRegionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserRegionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserRegionRequestValidationError) ErrorName() string {
	return "UpdateUserRegionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserRegionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserRegionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserRegionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserRegionRequestValidationError{}

// Validate checks the field values on UpdateUserRegionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateUserRegionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUserRegionResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateUserRegionResponseMultiError, or nil if none found.
func (m *UpdateUserRegionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserRegionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateUserRegionResponseMultiError(errors)
	}

	return nil
}

// UpdateUserRegionResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateUserRegionResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateUserRegionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserRegionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserRegionResponseMultiError) AllErrors() []error { return m }

// UpdateUserRegionResponseValidationError is the validation error returned by
// UpdateUserRegionResponse.Validate if the designated constraints aren't met.
type UpdateUserRegionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserRegionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserRegionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserRegionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserRegionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserRegionResponseValidationError) ErrorName() string {
	return "UpdateUserRegionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserRegionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserRegionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserRegionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserRegionResponseValidationError{}

// Validate checks the field values on GetUserSettingsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserSettingsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserSettingsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserSettingsRequestMultiError, or nil if none found.
func (m *GetUserSettingsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserSettingsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetUserSettingsRequestMultiError(errors)
	}

	return nil
}

// GetUserSettingsRequestMultiError is an error wrapping multiple validation
// errors returned by GetUserSettingsRequest.ValidateAll() if the designated
// constraints aren't met.
type GetUserSettingsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserSettingsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserSettingsRequestMultiError) AllErrors() []error { return m }

// GetUserSettingsRequestValidationError is the validation error returned by
// GetUserSettingsRequest.Validate if the designated constraints aren't met.
type GetUserSettingsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserSettingsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserSettingsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserSettingsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserSettingsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserSettingsRequestValidationError) ErrorName() string {
	return "GetUserSettingsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserSettingsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserSettingsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserSettingsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserSettingsRequestValidationError{}

// Validate checks the field values on GetUserSettingsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserSettingsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserSettingsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserSettingsResponseMultiError, or nil if none found.
func (m *GetUserSettingsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserSettingsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUserSettings()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetUserSettingsResponseValidationError{
					field:  "UserSettings",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetUserSettingsResponseValidationError{
					field:  "UserSettings",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserSettings()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserSettingsResponseValidationError{
				field:  "UserSettings",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetUserSettingsResponseMultiError(errors)
	}

	return nil
}

// GetUserSettingsResponseMultiError is an error wrapping multiple validation
// errors returned by GetUserSettingsResponse.ValidateAll() if the designated
// constraints aren't met.
type GetUserSettingsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserSettingsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserSettingsResponseMultiError) AllErrors() []error { return m }

// GetUserSettingsResponseValidationError is the validation error returned by
// GetUserSettingsResponse.Validate if the designated constraints aren't met.
type GetUserSettingsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserSettingsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserSettingsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserSettingsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserSettingsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserSettingsResponseValidationError) ErrorName() string {
	return "GetUserSettingsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserSettingsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserSettingsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserSettingsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserSettingsResponseValidationError{}
