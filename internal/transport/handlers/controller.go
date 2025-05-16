package http

import (
	"encoding/json"
	"net/http"
	"rbs-feedbox/internal/service"
)

func RegisterRoutes(svc *service.Service) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Работает"))
	})

	http.HandleFunc("/api/forms", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			forms, err := svc.GetForms()
			if err != nil {
				http.Error(w, "Ошибка получения форм", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(forms)
			return
		}
	})
}
