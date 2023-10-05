package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string       `xorm:"not null pk autoincr VARCHAR(255) id"`
	CreatedAt time.Time    `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time    `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt sql.NullTime `xorm:"TIMESTAMP deleted_at"`
	Username  string       `xorm:"not null VARCHAR(50) UNIQUE username"`
	Email     string       `xorm:"not null VARCHAR(255) passwd"`
	Passwd    string       `xorm:"not null VARCHAR(100) UNIQUE email"`
}

func (User) TableName() string {
	return "user"
}
