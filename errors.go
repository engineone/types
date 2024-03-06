package types

import "github.com/palantir/stacktrace"

const (
	// Task error codes
	ErrTaskNotFound = stacktrace.ErrorCode(iota)
	ErrInvalidTask
)
