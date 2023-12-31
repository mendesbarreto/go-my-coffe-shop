// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: barista.proto

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

// Validate checks the field values on GetMenuRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetMenuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetMenuRequestMultiError,
// or nil if none found.
func (m *GetMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetMenuRequestMultiError(errors)
	}

	return nil
}

// GetMenuRequestMultiError is an error wrapping multiple validation errors
// returned by GetMenuRequest.ValidateAll() if the designated constraints
// aren't met.
type GetMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetMenuRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetMenuRequestMultiError) AllErrors() []error { return m }

// GetMenuRequestValidationError is the validation error returned by
// GetMenuRequest.Validate if the designated constraints aren't met.
type GetMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMenuRequestValidationError) ErrorName() string { return "GetMenuRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMenuRequestValidationError{}

// Validate checks the field values on GetMenuResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetMenuResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetMenuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetMenuResponseMultiError, or nil if none found.
func (m *GetMenuResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetMenuResponse) validate(all bool) error {
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
					errors = append(errors, GetMenuResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetMenuResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetMenuResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetMenuResponseMultiError(errors)
	}

	return nil
}

// GetMenuResponseMultiError is an error wrapping multiple validation errors
// returned by GetMenuResponse.ValidateAll() if the designated constraints
// aren't met.
type GetMenuResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetMenuResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetMenuResponseMultiError) AllErrors() []error { return m }

// GetMenuResponseValidationError is the validation error returned by
// GetMenuResponse.Validate if the designated constraints aren't met.
type GetMenuResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMenuResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMenuResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMenuResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMenuResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMenuResponseValidationError) ErrorName() string { return "GetMenuResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetMenuResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMenuResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMenuResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMenuResponseValidationError{}

// Validate checks the field values on OrderRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderRequestMultiError, or
// nil if none found.
func (m *OrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProductId

	if len(errors) > 0 {
		return OrderRequestMultiError(errors)
	}

	return nil
}

// OrderRequestMultiError is an error wrapping multiple validation errors
// returned by OrderRequest.ValidateAll() if the designated constraints aren't met.
type OrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderRequestMultiError) AllErrors() []error { return m }

// OrderRequestValidationError is the validation error returned by
// OrderRequest.Validate if the designated constraints aren't met.
type OrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderRequestValidationError) ErrorName() string { return "OrderRequestValidationError" }

// Error satisfies the builtin error interface
func (e OrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderRequestValidationError{}

// Validate checks the field values on OrderResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderResponseMultiError, or
// nil if none found.
func (m *OrderResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OrderResponseMultiError(errors)
	}

	return nil
}

// OrderResponseMultiError is an error wrapping multiple validation errors
// returned by OrderResponse.ValidateAll() if the designated constraints
// aren't met.
type OrderResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderResponseMultiError) AllErrors() []error { return m }

// OrderResponseValidationError is the validation error returned by
// OrderResponse.Validate if the designated constraints aren't met.
type OrderResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderResponseValidationError) ErrorName() string { return "OrderResponseValidationError" }

// Error satisfies the builtin error interface
func (e OrderResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderResponseValidationError{}
