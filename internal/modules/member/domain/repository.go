package memberDomain

import "context"

type MemberRepository interface {
	// Create(ctx context.Context, newMember *MemberEntity) (*MemberEntity, error)
	GetByUserIDAndClassID(ctx context.Context, userID, classID string) (*MemberEntity, error)
	GetRoleIDByKey(ctx context.Context, key string) (string, error)
}
