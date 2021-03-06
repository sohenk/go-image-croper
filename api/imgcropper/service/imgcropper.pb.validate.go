// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: imgcropper.proto

package service

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

// Validate checks the field values on CropImgRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CropImgRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CropImgRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CropImgRequestMultiError,
// or nil if none found.
func (m *CropImgRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CropImgRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUrl()) < 1 {
		err := CropImgRequestValidationError{
			field:  "Url",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Width

	// no validation rules for Refresh

	if len(errors) > 0 {
		return CropImgRequestMultiError(errors)
	}

	return nil
}

// CropImgRequestMultiError is an error wrapping multiple validation errors
// returned by CropImgRequest.ValidateAll() if the designated constraints
// aren't met.
type CropImgRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CropImgRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CropImgRequestMultiError) AllErrors() []error { return m }

// CropImgRequestValidationError is the validation error returned by
// CropImgRequest.Validate if the designated constraints aren't met.
type CropImgRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CropImgRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CropImgRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CropImgRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CropImgRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CropImgRequestValidationError) ErrorName() string { return "CropImgRequestValidationError" }

// Error satisfies the builtin error interface
func (e CropImgRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCropImgRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CropImgRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CropImgRequestValidationError{}

// Validate checks the field values on CropImgReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CropImgReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CropImgReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CropImgReplyMultiError, or
// nil if none found.
func (m *CropImgReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CropImgReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Imgname

	// no validation rules for Imagetype

	// no validation rules for Imgdata

	if len(errors) > 0 {
		return CropImgReplyMultiError(errors)
	}

	return nil
}

// CropImgReplyMultiError is an error wrapping multiple validation errors
// returned by CropImgReply.ValidateAll() if the designated constraints aren't met.
type CropImgReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CropImgReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CropImgReplyMultiError) AllErrors() []error { return m }

// CropImgReplyValidationError is the validation error returned by
// CropImgReply.Validate if the designated constraints aren't met.
type CropImgReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CropImgReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CropImgReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CropImgReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CropImgReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CropImgReplyValidationError) ErrorName() string { return "CropImgReplyValidationError" }

// Error satisfies the builtin error interface
func (e CropImgReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCropImgReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CropImgReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CropImgReplyValidationError{}
