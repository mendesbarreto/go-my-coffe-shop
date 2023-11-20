// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: product.proto

package gen

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

// Validate checks the field values on ProductDetails with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ProductDetails) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProductDetails with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ProductDetailsMultiError,
// or nil if none found.
func (m *ProductDetails) ValidateAll() error {
	return m.validate(true)
}

func (m *ProductDetails) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetId()) > 12 {
		err := ProductDetailsValidationError{
			field:  "Id",
			reason: "value length must be at most 12 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Name

	// no validation rules for Price

	if len(errors) > 0 {
		return ProductDetailsMultiError(errors)
	}

	return nil
}

// ProductDetailsMultiError is an error wrapping multiple validation errors
// returned by ProductDetails.ValidateAll() if the designated constraints
// aren't met.
type ProductDetailsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductDetailsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductDetailsMultiError) AllErrors() []error { return m }

// ProductDetailsValidationError is the validation error returned by
// ProductDetails.Validate if the designated constraints aren't met.
type ProductDetailsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductDetailsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductDetailsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductDetailsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductDetailsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductDetailsValidationError) ErrorName() string { return "ProductDetailsValidationError" }

// Error satisfies the builtin error interface
func (e ProductDetailsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductDetails.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductDetailsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductDetailsValidationError{}

// Validate checks the field values on ListProductsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListProductsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListProductsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListProductsRequestMultiError, or nil if none found.
func (m *ListProductsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListProductsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListProductsRequestMultiError(errors)
	}

	return nil
}

// ListProductsRequestMultiError is an error wrapping multiple validation
// errors returned by ListProductsRequest.ValidateAll() if the designated
// constraints aren't met.
type ListProductsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListProductsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListProductsRequestMultiError) AllErrors() []error { return m }

// ListProductsRequestValidationError is the validation error returned by
// ListProductsRequest.Validate if the designated constraints aren't met.
type ListProductsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProductsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProductsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProductsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProductsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProductsRequestValidationError) ErrorName() string {
	return "ListProductsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListProductsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProductsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProductsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProductsRequestValidationError{}

// Validate checks the field values on ListProductsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListProductsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListProductsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListProductsResponseMultiError, or nil if none found.
func (m *ListProductsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListProductsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListProductsResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListProductsResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListProductsResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListProductsResponseMultiError(errors)
	}

	return nil
}

// ListProductsResponseMultiError is an error wrapping multiple validation
// errors returned by ListProductsResponse.ValidateAll() if the designated
// constraints aren't met.
type ListProductsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListProductsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListProductsResponseMultiError) AllErrors() []error { return m }

// ListProductsResponseValidationError is the validation error returned by
// ListProductsResponse.Validate if the designated constraints aren't met.
type ListProductsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProductsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProductsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProductsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProductsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProductsResponseValidationError) ErrorName() string {
	return "ListProductsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListProductsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProductsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProductsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProductsResponseValidationError{}
