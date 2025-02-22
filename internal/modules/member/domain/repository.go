package memberDomain

import (
	"context"

	memberRoleDomain "github.com/zchelalo/sa_class_management/internal/modules/member_role/domain"
)

type MemberRepository interface {
	// Create(ctx context.Context, newMember *MemberEntity) (*MemberEntity, error)
	ListByClassID(ctx context.Context, classID string, offset, limit int32) ([]*MemberEntity, error)
	Count(ctx context.Context, classID string) (int32, error)
	GetByUserIDAndClassID(ctx context.Context, userID, classID string) (*MemberEntity, error)
	GetRoleIDByKey(ctx context.Context, key string) (*memberRoleDomain.MemberRoleEntity, error)
}
