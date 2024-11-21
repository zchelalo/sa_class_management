package memberError

import "errors"

var ErrRoleNotFound = errors.New("role not found")

var ErrMemberNotFound = errors.New("member not found")

var ErrIdInvalid = errors.New("id invalid")
var ErrIdRequired = errors.New("id required")

var ErrRoleIdInvalid = errors.New("role id invalid")
var ErrRoleIdRequired = errors.New("role id required")

var ErrUserIdInvalid = errors.New("user id invalid")
var ErrUserIdRequired = errors.New("user id required")
