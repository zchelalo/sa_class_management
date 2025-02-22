package memberPostgresRepo

import (
	"context"
	"database/sql"

	memberDomain "github.com/zchelalo/sa_class_management/internal/modules/member/domain"
	memberError "github.com/zchelalo/sa_class_management/internal/modules/member/error"
	memberRoleDomain "github.com/zchelalo/sa_class_management/internal/modules/member_role/domain"
	userDomain "github.com/zchelalo/sa_class_management/internal/modules/user/domain"
	memberData "github.com/zchelalo/sa_class_management/pkg/sqlc/data/member/db"
	"github.com/zchelalo/sa_class_management/pkg/sqlc/db"
)

type PostgresRepository struct {
	store *db.SQLStore
}

func New(store *db.SQLStore) memberDomain.MemberRepository {
	return &PostgresRepository{
		store: store,
	}
}

func (r *PostgresRepository) ListByClassID(ctx context.Context, classID string, offset, limit int32) ([]*memberDomain.MemberEntity, error) {
	members, err := r.store.MemberQueries.ListMembers(ctx, memberData.ListMembersParams{
		ClassID: classID,
		Offset:  offset,
		Limit:   limit,
	})
	if err != nil {
		return nil, err
	}

	if len(members) == 0 {
		return nil, memberError.ErrMembersNotFound
	}

	var memberEntities []*memberDomain.MemberEntity
	for _, member := range members {
		memberEntities = append(memberEntities, &memberDomain.MemberEntity{
			ID: member.ID,
			User: userDomain.UserEntity{
				ID: member.UserID,
			},
			Role: memberRoleDomain.MemberRoleEntity{
				ID:  member.RoleID,
				Key: member.RoleKey,
			},
		})
	}

	return memberEntities, nil
}

func (r *PostgresRepository) Count(ctx context.Context, classID string) (int32, error) {
	count, err := r.store.MemberQueries.CountMembers(ctx, classID)
	if err != nil {
		return 0, err
	}

	return int32(count), nil
}

func (r *PostgresRepository) GetByUserIDAndClassID(ctx context.Context, userID, classID string) (*memberDomain.MemberEntity, error) {
	memberObtained, err := r.store.MemberQueries.GetMember(ctx, memberData.GetMemberParams{
		UserID:  userID,
		ClassID: classID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, memberError.ErrMemberNotFound
		}

		return nil, err
	}

	return &memberDomain.MemberEntity{
		ID: memberObtained.ID,
		User: userDomain.UserEntity{
			ID: memberObtained.UserID,
		},
		Role: memberRoleDomain.MemberRoleEntity{
			ID:  memberObtained.RoleID,
			Key: memberObtained.RoleKey,
		},
	}, nil
}

func (r *PostgresRepository) GetRoleIDByKey(ctx context.Context, key string) (*memberRoleDomain.MemberRoleEntity, error) {
	roleObtained, err := r.store.RoleQueries.GetRoleByKey(ctx, key)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, memberError.ErrRoleNotFound
		}

		return nil, err
	}

	return &memberRoleDomain.MemberRoleEntity{
		ID:  roleObtained.ID,
		Key: roleObtained.Key,
	}, nil
}
