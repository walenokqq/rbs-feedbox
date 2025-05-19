package models

import "time"

// Form - структура формы
type Form struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Schema    string    `json:"schema"`
	CreatedAt time.Time `json:"created_at"`
}

// Response - структура ответа на форму
type Response struct {
	ID        int       `json:"id"`
	FormID    int       `json:"form_id"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"status"`
}
