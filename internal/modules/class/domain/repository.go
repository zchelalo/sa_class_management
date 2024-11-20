package classDomain

import "context"

type ClassRepository interface {
	Create(ctx context.Context, userID string, newClass *ClassEntity) (*ClassEntity, error)
	Join(ctx context.Context, userID, code string) (*ClassEntity, error)
	List(ctx context.Context, userID string, page, limit int32) ([]*ClassEntity, error)
}
