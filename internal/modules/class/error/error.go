package classError

import "errors"

var ErrNameRequired = errors.New("name is required")
var ErrNameTooShort = errors.New("name must be at least 3 characters")

var ErrSubjectRequired = errors.New("subject is required")
var ErrSubjectTooShort = errors.New("subject must be at least 2 characters")

var ErrGradeRequired = errors.New("grade is required")
var ErrGradeTooShort = errors.New("grade must be at least 1 character")
