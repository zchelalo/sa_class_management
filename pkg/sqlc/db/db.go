package db

import (
	"context"
	"database/sql"
	"fmt"

	classDb "github.com/zchelalo/sa_class_management/pkg/sqlc/class/data"
	unitDb "github.com/zchelalo/sa_class_management/pkg/sqlc/unit/data"
)

type Querier interface {
	classDb.Querier
	unitDb.Querier
}

type SQLStore struct {
	db           *sql.DB
	classQueries *classDb.Queries
	unitQueries  *unitDb.Queries
}

func NewStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db:           db,
		classQueries: classDb.New(db),
		unitQueries:  unitDb.New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*SQLStore) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	classQueries := classDb.New(tx)
	unitQueries := unitDb.New(tx)

	transactionStore := &SQLStore{
		db:           store.db,
		classQueries: classQueries,
		unitQueries:  unitQueries,
	}

	err = fn(transactionStore)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
