// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: neutral_diet/food/v1/region.proto

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

// Validate checks the field values on Region with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Region) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Region with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RegionMultiError, or nil if none found.
func (m *Region) ValidateAll() error {
	return m.validate(true)
}

func (m *Region) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return RegionMultiError(errors)
	}

	return nil
}

// RegionMultiError is an error wrapping multiple validation errors returned by
// Region.ValidateAll() if the designated constraints aren't met.
type RegionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegionMultiError) AllErrors() []error { return m }

// RegionValidationError is the validation error returned by Region.Validate if
// the designated constraints aren't met.
type RegionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegionValidationError) ErrorName() string { return "RegionValidationError" }

// Error satisfies the builtin error interface
func (e RegionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegion.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegionValidationError{}