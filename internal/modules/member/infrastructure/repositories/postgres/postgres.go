package memberPostgresRepo

import (
	"context"
	"database/sql"

	memberDomain "github.com/zchelalo/sa_class_management/internal/modules/member/domain"
	memberError "github.com/zchelalo/sa_class_management/internal/modules/member/error"
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
		ID:     memberObtained.ID,
		UserID: memberObtained.UserID,
		RoleID: memberObtained.RoleID,
	}, nil
}

func (r *PostgresRepository) GetRoleIDByKey(ctx context.Context, key string) (string, error) {
	roleObtained, err := r.store.RoleQueries.GetRoleByKey(ctx, key)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", memberError.ErrRoleNotFound
		}

		return "", err
	}

	return roleObtained.ID, nil
}
