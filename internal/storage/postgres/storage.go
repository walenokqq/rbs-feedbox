package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type StoragePostgres struct {
	db *sql.DB
}

func NewStoragePostgres(dsn string) *StoragePostgres {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Недоступна база данных: %v", err)
	}
	log.Println("Подключение к базе данных успешно установлено")
	return &StoragePostgres{db: db}
}
