package classPostgres

// import (
// 	"context"

// 	classDomain "github.com/zchelalo/sa_class_management/internal/modules/class/domain"
// 	"github.com/zchelalo/sa_class_management/pkg/sqlc/db"
// )

// type PostgresRepository struct {
// 	store *db.SQLStore
// }

// func New(store *db.SQLStore) classDomain.ClassRepository {
// 	return &PostgresRepository{
// 		store: store,
// 	}
// }

// func (r *PostgresRepository) Create(ctx context.Context, userID string, newClass *classDomain.ClassEntity) (*classDomain.ClassEntity, error) {
// 	var class *classDomain.ClassEntity

// 	err := r.store.ExecTx(ctx, func(store *db.SQLStore) error {

// 	})

// 	return class, err
// }
