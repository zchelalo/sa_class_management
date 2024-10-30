package db

import (
	"context"
	"database/sql"
	"fmt"

	classData "github.com/zchelalo/sa_class_management/pkg/sqlc/data/class/db"
	unitData "github.com/zchelalo/sa_class_management/pkg/sqlc/data/unit/db"
)

type Querier interface {
	unitData.Querier
	classData.Querier
}

type SQLStore struct {
	DB           *sql.DB
	UnitQueries  *unitData.Queries
	ClassQueries *classData.Queries
}

func New(db *sql.DB) *SQLStore {
	return &SQLStore{
		DB:           db,
		UnitQueries:  unitData.New(db),
		ClassQueries: classData.New(db),
	}
}

func (store *SQLStore) ExecTx(ctx context.Context, fn func(*SQLStore) error) error {
	tx, err := store.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	unitQueries := unitData.New(tx)
	classQueries := classData.New(tx)

	transactionStore := &SQLStore{
		DB:           store.DB,
		UnitQueries:  unitQueries,
		ClassQueries: classQueries,
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
