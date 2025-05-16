package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storagepostgres struct {
	db *sql.DB
}

//подключение

func NewStoragepostgres(dsn string) *Storagepostgres {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("Ошибка подкл" + err.Error())
	}
	if err = db.Ping(); err != nil {
		panic("Недоступн база" + err.Error())
	}
	fmt.Println("Подключено")
	return &Storagepostgres{db: db}
}
