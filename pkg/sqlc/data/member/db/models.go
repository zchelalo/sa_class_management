// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package memberData

import (
	"database/sql"
	"time"
)

type Class struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Subject   sql.NullString `json:"subject"`
	Grade     sql.NullString `json:"grade"`
	Code      string         `json:"code"`
	DeletedAt sql.NullTime   `json:"deleted_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type Member struct {
	ID        string       `json:"id"`
	RoleID    string       `json:"role_id"`
	UserID    string       `json:"user_id"`
	ClassID   string       `json:"class_id"`
	DeletedAt sql.NullTime `json:"deleted_at"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type Role struct {
	ID        string       `json:"id"`
	Key       string       `json:"key"`
	DeletedAt sql.NullTime `json:"deleted_at"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type Unit struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	ClassID   string       `json:"class_id"`
	DeletedAt sql.NullTime `json:"deleted_at"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
