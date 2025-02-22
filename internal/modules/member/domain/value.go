package memberDomain

import (
	"github.com/google/uuid"
	memberError "github.com/zchelalo/sa_class_management/internal/modules/member/error"
)

func New(userID, roleID string) (*MemberEntity, error) {
	id := uuid.NewString()

	if err := IsUserIdValid(userID); err != nil {
		return nil, err
	}

	if err := IsRoleIDValid(roleID); err != nil {
		return nil, err
	}

	return &MemberEntity{
		ID:     id,
		RoleID: roleID,
		UserID: userID,
	}, nil
}

func IsIdValid(id string) error {
	if id == "" {
		return memberError.ErrIdRequired
	}

	if _, err := uuid.Parse(id); err != nil {
		return memberError.ErrIdInvalid
	}

	return nil
}

func IsUserIdValid(userID string) error {
	if userID == "" {
		return memberError.ErrUserIdRequired
	}

	if _, err := uuid.Parse(userID); err != nil {
		return memberError.ErrUserIdInvalid
	}

	return nil
}

func IsRoleIDValid(roleID string) error {
	if roleID == "" {
		return memberError.ErrRoleIdRequired
	}

	if _, err := uuid.Parse(roleID); err != nil {
		return memberError.ErrRoleIdInvalid
	}

	return nil
}

func IsPageValid(page int32) error {
	if page < 1 {
		return memberError.ErrPageInvalid
	}

	return nil
}

func IsLimitValid(limit int32) error {
	if limit < 1 {
		return memberError.ErrLimitInvalid
	}

	return nil
}
