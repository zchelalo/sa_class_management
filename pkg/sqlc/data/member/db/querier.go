// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package memberData

import (
	"context"
)

type Querier interface {
	CountMembers(ctx context.Context, classID string) (int64, error)
	CreateMember(ctx context.Context, arg CreateMemberParams) (Member, error)
	DeleteMember(ctx context.Context, arg DeleteMemberParams) error
	GetMember(ctx context.Context, arg GetMemberParams) (GetMemberRow, error)
	ListMembers(ctx context.Context, arg ListMembersParams) ([]ListMembersRow, error)
}

var _ Querier = (*Queries)(nil)
