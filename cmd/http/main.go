package main

import (
	"fmt"
	"net/http"

	"rbs-feedbox/cmd/http/routes"
	"rbs-feedbox/internal/storage/postgres"
)

func main() {
	dsn := "host=localhost port=5432 user=newuser password=newpass dbname=feedbox sslmode=disable"
	storage := postgres.NewStoragepostgres(dsn)

	routes.Register(storage)

	fmt.Println("📡 Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
