package userDomain

import (
	"regexp"

	"github.com/google/uuid"
	userError "github.com/zchelalo/sa_class_management/internal/modules/user/error"
)

const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func IsIdValid(id string) error {
	if id == "" {
		return userError.ErrIdRequired
	}
	if _, err := uuid.Parse(id); err != nil {
		return userError.ErrIdInvalid
	}
	return nil
}

func IsNameValid(name string) error {
	if name == "" {
		return userError.ErrNameRequired
	}

	if len(name) < 3 {
		return userError.ErrNameTooShort
	}

	return nil
}

func IsEmailValid(email string) error {
	if email == "" {
		return userError.ErrEmailRequired
	}
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return userError.ErrEmailInvalid
	}
	return nil
}
