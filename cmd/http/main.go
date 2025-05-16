package main

import (
	"fmt"
	"net/http"

	"rbs-feedbox/internal/service"
	"rbs-feedbox/internal/storage/postgres"
	httproutes "rbs-feedbox/internal/transport"
)

func main() {
	dsn := "host=localhost port=5432 user=newuser password=newpass dbname=feedbox sslmode=disable"
	storage := postgres.NewStoragepostgres(dsn)
	svc := service.New(storage)

	httproutes.Register(svc)
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
