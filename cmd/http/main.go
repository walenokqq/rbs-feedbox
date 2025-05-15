package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rbs-feedbox/internal/storage/postgres"
)

func main() {
	dsn := "host=localhost port=5432 user=newuser password=newpass dbname=feedbox sslmode=disable"
	storage := postgres.NewStoragePG(dsn)
	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			projects, err := storage.GetProjects()
			if err != nil {
				http.Error(w, "Ошибка получения проектов", 500)
				return
			}
			json.NewEncoder(w).Encode(projects)
		}
	})
	fmt.Println("http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
