package memberDomain

import (
	memberRoleDomain "github.com/zchelalo/sa_class_management/internal/modules/member_role/domain"
	userDomain "github.com/zchelalo/sa_class_management/internal/modules/user/domain"
)

type MemberEntity struct {
	ID   string                            `json:"id"`
	User userDomain.UserEntity             `json:"user"`
	Role memberRoleDomain.MemberRoleEntity `json:"role"`
}
