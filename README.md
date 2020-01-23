# Go ORM Example

This repository is example of Go ORM.  
Use as example in Lightning Talk on [DeNA.go #4](https://dena.connpass.com/event/160018/).

- [gorm](https://github.com/jinzhu/gorm)
- [gorp](https://github.com/go-gorp/gorp)
- [sqlx](https://github.com/jmoiron/sqlx)

## Setup

```console
$ mysql -u root -h 127.0.0.1 < misc/queries/create_database.sql
$ mysql -u root -h 127.0.0.1 go_orm_example < misc/queries/ddl.sql
$ mysql -u root -h 127.0.0.1 go_orm_example < misc/queries/test_data.sql
```

## License

[MIT](LICENSE)

## Author

[178inaba](https://github.com/178inaba)
