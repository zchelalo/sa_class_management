package classDomain

import (
	"context"

	memberDomain "github.com/zchelalo/sa_class_management/internal/modules/member/domain"
)

type ClassRepository interface {
	Create(ctx context.Context, newMember *memberDomain.MemberEntity, newClass *ClassEntity) (*ClassEntity, error)
	Join(ctx context.Context, newMember *memberDomain.MemberEntity, classID string) error
	List(ctx context.Context, userID string, offset, limit int32) ([]*ClassEntity, error)
}
