package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint
	Name      string
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type Address struct {
	ID          uint
	UserID      uint   `db:"user_id"`
	ZipCode     string `db:"zip_code"`
	Prefecture  string
	Address1    string       `db:"address_1"`
	Address2    string       `db:"address_2"`
	PhoneNumber string       `db:"phone_number"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}
