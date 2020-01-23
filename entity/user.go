package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

type Address struct {
	ID          uint
	UserID      uint
	ZipCode     string
	Prefecture  string
	Address1    string
	Address2    string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}
