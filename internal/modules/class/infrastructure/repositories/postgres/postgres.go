package classPostgresRepo

import (
	"context"
	"database/sql"

	classDomain "github.com/zchelalo/sa_class_management/internal/modules/class/domain"
	classError "github.com/zchelalo/sa_class_management/internal/modules/class/error"
	memberDomain "github.com/zchelalo/sa_class_management/internal/modules/member/domain"
	classData "github.com/zchelalo/sa_class_management/pkg/sqlc/data/class/db"
	memberData "github.com/zchelalo/sa_class_management/pkg/sqlc/data/member/db"
	"github.com/zchelalo/sa_class_management/pkg/sqlc/db"
)

type PostgresRepository struct {
	store *db.SQLStore
}

func New(store *db.SQLStore) classDomain.ClassRepository {
	return &PostgresRepository{
		store: store,
	}
}

func (r *PostgresRepository) Create(ctx context.Context, newMember *memberDomain.MemberEntity, newClass *classDomain.ClassEntity) (*classDomain.ClassEntity, error) {
	var class *classDomain.ClassEntity

	err := r.store.ExecTx(ctx, func(store *db.SQLStore) error {
		classCreated, err := store.ClassQueries.CreateClass(ctx, classData.CreateClassParams{
			ID:      newClass.ID,
			Name:    newClass.Name,
			Subject: sql.NullString{String: newClass.Subject, Valid: true},
			Grade:   sql.NullString{String: newClass.Grade, Valid: true},
			Code:    newClass.Code,
		})
		if err != nil {
			return err
		}

		_, err = store.MemberQueries.CreateMember(ctx, memberData.CreateMemberParams{
			ID:      newMember.ID,
			RoleID:  newMember.RoleID,
			UserID:  newMember.UserID,
			ClassID: classCreated.ID,
		})
		if err != nil {
			return err
		}

		class = &classDomain.ClassEntity{
			ID:      classCreated.ID,
			Name:    classCreated.Name,
			Subject: classCreated.Subject.String,
			Grade:   classCreated.Grade.String,
			Code:    classCreated.Code,
		}

		return nil
	})

	return class, err
}

func (r *PostgresRepository) Join(ctx context.Context, newMember *memberDomain.MemberEntity, classID string) error {
	_, err := r.store.MemberQueries.CreateMember(ctx, memberData.CreateMemberParams{
		ID:      newMember.ID,
		RoleID:  newMember.RoleID,
		UserID:  newMember.UserID,
		ClassID: classID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return classError.ErrClassNotFound
		}

		return err
	}

	return nil
}

func (r *PostgresRepository) List(ctx context.Context, userID string, offset, limit int32) ([]*classDomain.ClassEntity, error) {
	classes, err := r.store.ClassQueries.ListClasses(ctx, classData.ListClassesParams{
		UserID: userID,
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}

	var classEntities []*classDomain.ClassEntity
	for _, class := range classes {
		classEntities = append(classEntities, &classDomain.ClassEntity{
			ID:      class.ID,
			Name:    class.Name,
			Subject: class.Subject.String,
			Grade:   class.Grade.String,
		})
	}

	return classEntities, nil
}

func (r *PostgresRepository) Count(ctx context.Context, userID string) (int32, error) {
	count, err := r.store.ClassQueries.CountClasses(ctx, userID)
	if err != nil {
		return 0, err
	}

	return int32(count), nil
}

func (r *PostgresRepository) GetClassByCode(ctx context.Context, code string) (*classDomain.ClassEntity, error) {
	classObtained, err := r.store.ClassQueries.GetClassByCode(ctx, code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, classError.ErrClassNotFound
		}

		return nil, err
	}

	return &classDomain.ClassEntity{
		ID:      classObtained.ID,
		Name:    classObtained.Name,
		Subject: classObtained.Subject.String,
		Grade:   classObtained.Grade.String,
	}, nil
}

func (r *PostgresRepository) GetMemberRoleByClassIDAndUserID(ctx context.Context, classID, userID string) (string, error) {
	roleObtained, err := r.store.ClassQueries.GetMemberRoleByClassIDAndUserID(ctx, classData.GetMemberRoleByClassIDAndUserIDParams{
		ClassID: classID,
		UserID:  userID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return "", classError.ErrClassNotFound
		}

		return "", err
	}

	return roleObtained, nil
}

func (r *PostgresRepository) GetClassCodeByClassID(ctx context.Context, classID string) (string, error) {
	code, err := r.store.ClassQueries.GetClassCodeByClassID(ctx, classID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", classError.ErrClassNotFound
		}

		return "", err
	}

	return code, nil
}
