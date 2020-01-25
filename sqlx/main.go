package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/go_orm_example?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	dbx := sqlx.NewDb(db, "mysql")

	ctx := context.Background()

	var us []User

	if err := dbx.SelectContext(ctx, &us, `
SELECT
  id
  , name
  , created_at
  , updated_at
  , deleted_at
FROM
  users`); err != nil {
		log.Fatal(err)
	}

	pp.Println(us)
}

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
