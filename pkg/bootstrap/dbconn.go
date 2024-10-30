package bootstrap

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/zchelalo/sa_class_management/pkg/sqlc/connection"
)

type SingletonDB struct {
	conn *sql.DB
}

var instance *SingletonDB
var onceDb sync.Once

func InitInstance(driver, source string) error {
	var err error
	onceDb.Do(func() {
		conn, initErr := connection.New(driver, source)
		if initErr != nil {
			err = initErr
			return
		}
		instance = &SingletonDB{
			conn: conn,
		}
	})
	return err
}

func GetInstance() (*sql.DB, error) {
	if instance == nil {
		return nil, fmt.Errorf("instance not initialized")
	}
	return instance.conn, nil
}
