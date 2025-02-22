package classError

import "errors"

var ErrNameRequired = errors.New("name is required")
var ErrNameTooShort = errors.New("name must be at least 3 characters")

var ErrSubjectRequired = errors.New("subject is required")
var ErrSubjectTooShort = errors.New("subject must be at least 2 characters")

var ErrGradeRequired = errors.New("grade is required")
var ErrGradeTooShort = errors.New("grade must be at least 1 character")

var ErrClassNotFound = errors.New("class not found")
var ErrClassesNotFound = errors.New("classes not found")

var ErrPageInvalid = errors.New("page is invalid")
var ErrLimitInvalid = errors.New("limit is invalid")

var ErrCodeRequired = errors.New("code is required")
var ErrCodeInvalid = errors.New("code is invalid")

var ErrIDRequired = errors.New("id is required")
var ErrIDInvalid = errors.New("id is invalid")

var ErrUnauthorized = errors.New("unauthorized access")

var ErrAlreadyJoined = errors.New("already joined")
