package model

import "errors"

var ErrUserAlreadyExists = errors.New("user already exist")
var ErrAuthenticationFailure = errors.New("authentication failed")
