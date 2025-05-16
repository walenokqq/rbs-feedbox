package main

import (
	"fmt"
	"net/http"
	"os"

	"rbs-feedbox/internal/service"
	"rbs-feedbox/internal/storage/postgres"
	httproutes "rbs-feedbox/internal/transport"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	storage := postgres.NewStoragepostgres(dsn)
	svc := service.New(storage)

	httproutes.Register(svc)
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
