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

var (
	instance *SingletonDB
	once     sync.Once
	initErr  error
)

func initInstance(driver, source string) {
	once.Do(func() {
		conn, err := connection.New(driver, source)
		if err != nil {
			initErr = err
			return
		}
		instance = &SingletonDB{
			conn: conn,
		}
	})
}

func GetInstance(driver, source string) (*sql.DB, error) {
	if instance == nil && initErr == nil {
		initInstance(driver, source)
	}
	if initErr != nil {
		return nil, initErr
	}
	return instance.conn, nil
}

func Close() error {
	if instance == nil || instance.conn == nil {
		return fmt.Errorf("no active database connection")
	}
	return instance.conn.Close()
}
