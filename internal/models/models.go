package models

import "time"

// Form — структура для работы с таблицей forms в БД
type Form struct {
	ID          int       `json:"id"`
	ProjectID   int       `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Schema      string    `json:"schema"`
	CreatedAt   time.Time `json:"created_at"`
}

// Response - структура ответа на форму
type Response struct {
	ID        int       `json:"id"`
	FormID    int       `json:"form_id"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"status"`
}
