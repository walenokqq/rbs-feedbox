package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type StoragePG struct {
	db *sql.DB
}

//подключение

func NewStoragePG(dsn string) *StoragePG {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		// panic(fmt.Sprintf("Ошибка подкл" + err.Error()))
		panic("Ошибка подкл" + err.Error())
	}

	if err = db.Ping(); err != nil {
		// panic(fmt.Sprintf("Недоступн база" + err.Error()))
		panic("Недоступн база" + err.Error())
	}
	fmt.Println("Подключено")
	return &StoragePG{db: db}
}
