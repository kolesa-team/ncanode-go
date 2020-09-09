package ncanode

import "errors"

var (
	// ErrInvalidRequestBody is returned if input parameters of client method has
	// not enough or empty values.
	ErrInvalidRequestBody = errors.New("ncanode: invalid request body")
)
