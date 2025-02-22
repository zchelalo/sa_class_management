package memberDomain

import (
	"github.com/google/uuid"
	memberError "github.com/zchelalo/sa_class_management/internal/modules/member/error"
	memberRoleDomain "github.com/zchelalo/sa_class_management/internal/modules/member_role/domain"
	userDomain "github.com/zchelalo/sa_class_management/internal/modules/user/domain"
)

func New(user userDomain.UserEntity, role memberRoleDomain.MemberRoleEntity) (*MemberEntity, error) {
	id := uuid.NewString()

	if err := userDomain.IsIdValid(user.ID); err != nil {
		return nil, err
	}

	if err := memberRoleDomain.IsIdValid(role.ID); err != nil {
		return nil, err
	}

	return &MemberEntity{
		ID:   id,
		User: user,
		Role: role,
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
