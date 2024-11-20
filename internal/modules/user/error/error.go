package userError

import "errors"

var ErrNameRequired = errors.New("name is required")
var ErrNameTooShort = errors.New("name must be at least 3 characters")

var ErrEmailRequired = errors.New("email is required")
var ErrEmailInvalid = errors.New("email is invalid")

var ErrIdRequired = errors.New("id is required")
var ErrIdInvalid = errors.New("id is invalid")

var ErrUserNotFound = errors.New("user not found")
