package main

import (
	"fmt"
	"net/http"
	"os"

	"rbs-feedbox/internal/service"
	"rbs-feedbox/internal/storage/postgres"
	httproutes "rbs-feedbox/internal/transport"
)

// main - инициализирует хранилище, сервисный слой и регистрирует HTTP-маршруты.
func main() {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	storage := postgres.NewStoragePostgres(dsn)
	svc := service.New(storage)

	httproutes.Register(svc)

	// fs := http.FileServer(http.Dir("./public"))
	// http.Handle("/", fs)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
