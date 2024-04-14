package model

import "errors"

var (
	ErrUserAlreadyExists     = errors.New("user already exist")
	ErrAuthenticationFailure = errors.New("authentication failed")
)
