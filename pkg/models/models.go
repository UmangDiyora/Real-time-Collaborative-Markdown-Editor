package models

// This file contains common types and interfaces used across models

import (
	"errors"
)

// Common errors
var (
	ErrNotFound          = errors.New("resource not found")
	ErrUnauthorized      = errors.New("unauthorized access")
	ErrInvalidInput      = errors.New("invalid input")
	ErrConflict          = errors.New("resource conflict")
	ErrInternalError     = errors.New("internal server error")
	ErrInvalidPermission = errors.New("invalid permission level")
	ErrExpiredPermission = errors.New("permission has expired")
	ErrDocumentDeleted   = errors.New("document has been deleted")
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

// Error implements the error interface
func (v ValidationError) Error() string {
	return v.Field + ": " + v.Message
}

// Validator interface for models that can be validated
type Validator interface {
	Validate() error
}
