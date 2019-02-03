// Package errors contains errors for the domain "lis".
//
// This file is automatically generated by errawr-gen. Do not modify it.
package errors

import (
	errawrgo "github.com/puppetlabs/errawr-go"
	impl "github.com/puppetlabs/errawr-go/impl"
)

// Error is the type of all errors generated by this package.
type Error interface {
	errawrgo.Error
}

// External contains methods that can be used externally to help consume errors from this package.
type External struct{}

// API is a singleton instance of the External type.
var API External

// Domain is the general domain in which all errors in this package belong.
var Domain = &impl.ErrorDomain{
	Key:   "lis",
	Title: "Insights Scheduler Library",
}

// LifecycleSection defines a section of errors with the following scope:
// Lifecycle errors
var LifecycleSection = &impl.ErrorSection{
	Key:   "lifecycle",
	Title: "Lifecycle errors",
}

// LifecycleCloseErrorCode is the code for an instance of "close_error".
const LifecycleCloseErrorCode = "lis_lifecycle_close_error"

// IsLifecycleCloseError tests whether a given error is an instance of "close_error".
func IsLifecycleCloseError(err errawrgo.Error) bool {
	return err != nil && err.Is(LifecycleCloseErrorCode)
}

// IsLifecycleCloseError tests whether a given error is an instance of "close_error".
func (External) IsLifecycleCloseError(err errawrgo.Error) bool {
	return IsLifecycleCloseError(err)
}

// LifecycleCloseErrorBuilder is a builder for "close_error" errors.
type LifecycleCloseErrorBuilder struct {
	arguments impl.ErrorArguments
}

// Build creates the error for the code "close_error" from this builder.
func (b *LifecycleCloseErrorBuilder) Build() Error {
	description := &impl.ErrorDescription{
		Friendly:  "The scheduled tasks failed to terminate.",
		Technical: "The scheduled tasks failed to terminate.",
	}

	return &impl.Error{
		ErrorArguments:   b.arguments,
		ErrorCode:        "close_error",
		ErrorDescription: description,
		ErrorDomain:      Domain,
		ErrorMetadata:    &impl.ErrorMetadata{},
		ErrorSection:     LifecycleSection,
		ErrorSensitivity: errawrgo.ErrorSensitivityNone,
		ErrorTitle:       "Close error",
		Version:          1,
	}
}

// NewLifecycleCloseErrorBuilder creates a new error builder for the code "close_error".
func NewLifecycleCloseErrorBuilder() *LifecycleCloseErrorBuilder {
	return &LifecycleCloseErrorBuilder{arguments: impl.ErrorArguments{}}
}

// NewLifecycleCloseError creates a new error with the code "close_error".
func NewLifecycleCloseError() Error {
	return NewLifecycleCloseErrorBuilder().Build()
}

// LifecycleDescriptorErrorCode is the code for an instance of "descriptor_error".
const LifecycleDescriptorErrorCode = "lis_lifecycle_descriptor_error"

// IsLifecycleDescriptorError tests whether a given error is an instance of "descriptor_error".
func IsLifecycleDescriptorError(err errawrgo.Error) bool {
	return err != nil && err.Is(LifecycleDescriptorErrorCode)
}

// IsLifecycleDescriptorError tests whether a given error is an instance of "descriptor_error".
func (External) IsLifecycleDescriptorError(err errawrgo.Error) bool {
	return IsLifecycleDescriptorError(err)
}

// LifecycleDescriptorErrorBuilder is a builder for "descriptor_error" errors.
type LifecycleDescriptorErrorBuilder struct {
	arguments impl.ErrorArguments
}

// Build creates the error for the code "descriptor_error" from this builder.
func (b *LifecycleDescriptorErrorBuilder) Build() Error {
	description := &impl.ErrorDescription{
		Friendly:  "Descriptor \\#{{i}} of type {{pre type}} terminated unexpectedly.",
		Technical: "Descriptor \\#{{i}} of type {{pre type}} terminated unexpectedly.",
	}

	return &impl.Error{
		ErrorArguments:   b.arguments,
		ErrorCode:        "descriptor_error",
		ErrorDescription: description,
		ErrorDomain:      Domain,
		ErrorMetadata:    &impl.ErrorMetadata{},
		ErrorSection:     LifecycleSection,
		ErrorSensitivity: errawrgo.ErrorSensitivityNone,
		ErrorTitle:       "Descriptor error",
		Version:          1,
	}
}

// NewLifecycleDescriptorErrorBuilder creates a new error builder for the code "descriptor_error".
func NewLifecycleDescriptorErrorBuilder(i int64, type_ string) *LifecycleDescriptorErrorBuilder {
	return &LifecycleDescriptorErrorBuilder{arguments: impl.ErrorArguments{
		"i":    impl.NewErrorArgument(i, "the offset into the descriptor list of this descriptor\n"),
		"type": impl.NewErrorArgument(type_, "the type of this descriptor"),
	}}
}

// NewLifecycleDescriptorError creates a new error with the code "descriptor_error".
func NewLifecycleDescriptorError(i int64, type_ string) Error {
	return NewLifecycleDescriptorErrorBuilder(i, type_).Build()
}

// LifecycleExecutionErrorCode is the code for an instance of "execution_error".
const LifecycleExecutionErrorCode = "lis_lifecycle_execution_error"

// IsLifecycleExecutionError tests whether a given error is an instance of "execution_error".
func IsLifecycleExecutionError(err errawrgo.Error) bool {
	return err != nil && err.Is(LifecycleExecutionErrorCode)
}

// IsLifecycleExecutionError tests whether a given error is an instance of "execution_error".
func (External) IsLifecycleExecutionError(err errawrgo.Error) bool {
	return IsLifecycleExecutionError(err)
}

// LifecycleExecutionErrorBuilder is a builder for "execution_error" errors.
type LifecycleExecutionErrorBuilder struct {
	arguments impl.ErrorArguments
}

// Build creates the error for the code "execution_error" from this builder.
func (b *LifecycleExecutionErrorBuilder) Build() Error {
	description := &impl.ErrorDescription{
		Friendly:  "The scheduled processes failed to execute.",
		Technical: "The scheduled processes failed to execute.",
	}

	return &impl.Error{
		ErrorArguments:   b.arguments,
		ErrorCode:        "execution_error",
		ErrorDescription: description,
		ErrorDomain:      Domain,
		ErrorMetadata:    &impl.ErrorMetadata{},
		ErrorSection:     LifecycleSection,
		ErrorSensitivity: errawrgo.ErrorSensitivityNone,
		ErrorTitle:       "Execution error",
		Version:          1,
	}
}

// NewLifecycleExecutionErrorBuilder creates a new error builder for the code "execution_error".
func NewLifecycleExecutionErrorBuilder() *LifecycleExecutionErrorBuilder {
	return &LifecycleExecutionErrorBuilder{arguments: impl.ErrorArguments{}}
}

// NewLifecycleExecutionError creates a new error with the code "execution_error".
func NewLifecycleExecutionError() Error {
	return NewLifecycleExecutionErrorBuilder().Build()
}

// LifecycleProcessErrorCode is the code for an instance of "process_error".
const LifecycleProcessErrorCode = "lis_lifecycle_process_error"

// IsLifecycleProcessError tests whether a given error is an instance of "process_error".
func IsLifecycleProcessError(err errawrgo.Error) bool {
	return err != nil && err.Is(LifecycleProcessErrorCode)
}

// IsLifecycleProcessError tests whether a given error is an instance of "process_error".
func (External) IsLifecycleProcessError(err errawrgo.Error) bool {
	return IsLifecycleProcessError(err)
}

// LifecycleProcessErrorBuilder is a builder for "process_error" errors.
type LifecycleProcessErrorBuilder struct {
	arguments impl.ErrorArguments
}

// Build creates the error for the code "process_error" from this builder.
func (b *LifecycleProcessErrorBuilder) Build() Error {
	description := &impl.ErrorDescription{
		Friendly:  "The process {{quote description}} running under request {{pre request_id}} encountered an unrecoverable error.",
		Technical: "The process {{quote description}} running under request {{pre request_id}} encountered an unrecoverable error.",
	}

	return &impl.Error{
		ErrorArguments:   b.arguments,
		ErrorCode:        "process_error",
		ErrorDescription: description,
		ErrorDomain:      Domain,
		ErrorMetadata:    &impl.ErrorMetadata{},
		ErrorSection:     LifecycleSection,
		ErrorSensitivity: errawrgo.ErrorSensitivityNone,
		ErrorTitle:       "Process error",
		Version:          1,
	}
}

// NewLifecycleProcessErrorBuilder creates a new error builder for the code "process_error".
func NewLifecycleProcessErrorBuilder(requestID string, description string) *LifecycleProcessErrorBuilder {
	return &LifecycleProcessErrorBuilder{arguments: impl.ErrorArguments{
		"description": impl.NewErrorArgument(description, "the process's description"),
		"request_id":  impl.NewErrorArgument(requestID, "the request identifier generated for the process"),
	}}
}

// NewLifecycleProcessError creates a new error with the code "process_error".
func NewLifecycleProcessError(requestID string, description string) Error {
	return NewLifecycleProcessErrorBuilder(requestID, description).Build()
}

// LifecycleTimeoutErrorCode is the code for an instance of "timeout_error".
const LifecycleTimeoutErrorCode = "lis_lifecycle_timeout_error"

// IsLifecycleTimeoutError tests whether a given error is an instance of "timeout_error".
func IsLifecycleTimeoutError(err errawrgo.Error) bool {
	return err != nil && err.Is(LifecycleTimeoutErrorCode)
}

// IsLifecycleTimeoutError tests whether a given error is an instance of "timeout_error".
func (External) IsLifecycleTimeoutError(err errawrgo.Error) bool {
	return IsLifecycleTimeoutError(err)
}

// LifecycleTimeoutErrorBuilder is a builder for "timeout_error" errors.
type LifecycleTimeoutErrorBuilder struct {
	arguments impl.ErrorArguments
}

// Build creates the error for the code "timeout_error" from this builder.
func (b *LifecycleTimeoutErrorBuilder) Build() Error {
	description := &impl.ErrorDescription{
		Friendly:  "The operation timed out while waiting for tasks to complete.",
		Technical: "The operation timed out while waiting for tasks to complete.",
	}

	return &impl.Error{
		ErrorArguments:   b.arguments,
		ErrorCode:        "timeout_error",
		ErrorDescription: description,
		ErrorDomain:      Domain,
		ErrorMetadata:    &impl.ErrorMetadata{},
		ErrorSection:     LifecycleSection,
		ErrorSensitivity: errawrgo.ErrorSensitivityNone,
		ErrorTitle:       "Timeout",
		Version:          1,
	}
}

// NewLifecycleTimeoutErrorBuilder creates a new error builder for the code "timeout_error".
func NewLifecycleTimeoutErrorBuilder() *LifecycleTimeoutErrorBuilder {
	return &LifecycleTimeoutErrorBuilder{arguments: impl.ErrorArguments{}}
}

// NewLifecycleTimeoutError creates a new error with the code "timeout_error".
func NewLifecycleTimeoutError() Error {
	return NewLifecycleTimeoutErrorBuilder().Build()
}

// ProcessSection defines a section of errors with the following scope:
// Process errors
var ProcessSection = &impl.ErrorSection{
	Key:   "process",
	Title: "Process errors",
}

// ProcessPanicErrorCode is the code for an instance of "panic_error".
const ProcessPanicErrorCode = "lis_process_panic_error"

// IsProcessPanicError tests whether a given error is an instance of "panic_error".
func IsProcessPanicError(err errawrgo.Error) bool {
	return err != nil && err.Is(ProcessPanicErrorCode)
}

// IsProcessPanicError tests whether a given error is an instance of "panic_error".
func (External) IsProcessPanicError(err errawrgo.Error) bool {
	return IsProcessPanicError(err)
}

// ProcessPanicErrorBuilder is a builder for "panic_error" errors.
type ProcessPanicErrorBuilder struct {
	arguments impl.ErrorArguments
}

// Build creates the error for the code "panic_error" from this builder.
func (b *ProcessPanicErrorBuilder) Build() Error {
	description := &impl.ErrorDescription{
		Friendly:  "The process panic()ed.",
		Technical: "The process panic()ed.",
	}

	return &impl.Error{
		ErrorArguments:   b.arguments,
		ErrorCode:        "panic_error",
		ErrorDescription: description,
		ErrorDomain:      Domain,
		ErrorMetadata:    &impl.ErrorMetadata{},
		ErrorSection:     ProcessSection,
		ErrorSensitivity: errawrgo.ErrorSensitivityNone,
		ErrorTitle:       "Panic",
		Version:          1,
	}
}

// NewProcessPanicErrorBuilder creates a new error builder for the code "panic_error".
func NewProcessPanicErrorBuilder() *ProcessPanicErrorBuilder {
	return &ProcessPanicErrorBuilder{arguments: impl.ErrorArguments{}}
}

// NewProcessPanicError creates a new error with the code "panic_error".
func NewProcessPanicError() Error {
	return NewProcessPanicErrorBuilder().Build()
}