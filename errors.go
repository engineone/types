package types

import "github.com/palantir/stacktrace"

const (
	// Task error codes
	ErrTaskNotFound = stacktrace.ErrorCode(iota)
	ErrInvalidTask

	// Executor error codes
	ErrExecutorFailed

	// Validation error codes
	ErrInvalidInput
)
