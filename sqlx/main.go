package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/178inaba/go-orm-example/entity"
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

	var us []entity.User

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
