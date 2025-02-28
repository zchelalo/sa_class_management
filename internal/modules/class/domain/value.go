package classDomain

import (
	"crypto/rand"
	"math/big"

	"github.com/google/uuid"
	classError "github.com/zchelalo/sa_class_management/internal/modules/class/error"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const codeLength = 6

func New(name, subject, grade string) (*ClassEntity, error) {
	id := uuid.NewString()

	if err := IsNameValid(name); err != nil {
		return nil, err
	}

	if err := IsSubjectValid(subject); err != nil {
		return nil, err
	}

	if err := IsGradeValid(grade); err != nil {
		return nil, err
	}

	code, err := generateUniqueCode(codeLength)
	if err != nil {
		return nil, err
	}

	return &ClassEntity{
		ID:      id,
		Name:    name,
		Subject: subject,
		Grade:   grade,
		Code:    code,
	}, nil
}

func IsIdValid(id string) error {
	if id == "" {
		return classError.ErrIdInvalid
	}

	if _, err := uuid.Parse(id); err != nil {
		return classError.ErrIdInvalid
	}

	return nil
}

func IsNameValid(name string) error {
	if name == "" {
		return classError.ErrNameRequired
	}

	if len(name) < 3 {
		return classError.ErrNameTooShort
	}

	return nil
}

func IsSubjectValid(subject string) error {
	if subject == "" {
		return classError.ErrSubjectRequired
	}

	if len(subject) < 2 {
		return classError.ErrSubjectTooShort
	}

	return nil
}

func IsGradeValid(grade string) error {
	if grade == "" {
		return classError.ErrGradeRequired
	}

	if len(grade) < 1 {
		return classError.ErrGradeTooShort
	}

	return nil
}

func IsPageValid(page int32) error {
	if page < 1 {
		return classError.ErrPageInvalid
	}

	return nil
}

func IsLimitValid(limit int32) error {
	if limit < 1 {
		return classError.ErrLimitInvalid
	}

	return nil
}

func IsCodeValid(code string) error {
	if code == "" {
		return classError.ErrCodeRequired
	}

	if len(code) != codeLength {
		return classError.ErrCodeInvalid
	}

	return nil
}

func generateUniqueCode(length int) (string, error) {
	b := make([]byte, length)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[num.Int64()]
	}
	return string(b), nil
}
