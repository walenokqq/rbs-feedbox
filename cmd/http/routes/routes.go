package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"rbs-feedbox/internal/storage/postgres"
)

func Register(storage *postgres.Storagepostgres) {
	// api/responses/{id}
	http.HandleFunc("/api/responses/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			http.Error(w, "не поддерживается метод", http.StatusMethodNotAllowed)
			return
		}

		idStr := strings.TrimPrefix(r.URL.Path, "/api/responses/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "неверный id", http.StatusBadRequest)
			return
		}

		var req struct {
			Status string `json:"status"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "неверный JSON", http.StatusBadRequest)
			return
		}

		err = storage.UpdateResponseStatus(id, req.Status)
		if err != nil {
			http.Error(w, "ошибка статуса: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "Статус обновлён"})
	})

	//  /api/forms и GET /api/forms
	http.HandleFunc("/api/forms", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var req struct {
				Name   string `json:"name"`
				Schema string `json:"schema"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "неверный JSON", http.StatusBadRequest)
				return
			}
			err := storage.CreateForm(req.Name, req.Schema)
			if err != nil {
				http.Error(w, "ошибка создания: "+err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(map[string]string{"status": "Форма создана"})
			return
		}

		if r.Method == http.MethodGet {
			forms, err := storage.GetForms()
			if err != nil {
				http.Error(w, "неверное получение формы", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(forms)
			return
		}

		http.Error(w, "не поддержив метод", http.StatusMethodNotAllowed)
	})

	// /api/forms/{id}, /submit, /responses
	http.HandleFunc("/api/forms/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// POST /api/forms/{id}/submit
		if strings.HasSuffix(path, "/submit") && r.Method == http.MethodPost {
			parts := strings.Split(strings.TrimSuffix(path, "/submit"), "/")
			idStr := parts[len(parts)-1]

			formID, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "неверная или неккоректная форма id", http.StatusBadRequest)
				return
			}

			var req struct {
				Data string `json:"data"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "неверный JSON", http.StatusBadRequest)
				return
			}

			err = storage.SaveResponse(formID, req.Data)
			if err != nil {
				http.Error(w, "ошибка ответа: "+err.Error(), http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(map[string]string{"status": "Ответ сохранён"})
			return
		}

		// GET /api/forms/{id}/responses
		if strings.HasSuffix(path, "/responses") && r.Method == http.MethodGet {
			parts := strings.Split(strings.TrimSuffix(path, "/responses"), "/")
			idStr := parts[len(parts)-1]

			formID, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "неверная форма id", http.StatusBadRequest)
				return
			}

			responses, err := storage.GetResponsesByFormID(formID)
			if err != nil {
				http.Error(w, "ошибка ответа: "+err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responses)
			return
		}

		// GET /api/forms/{id}
		if r.Method == http.MethodGet {
			idStr := strings.TrimPrefix(path, "/api/forms/")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "неверный или неккор id", http.StatusBadRequest)
				return
			}

			form, err := storage.GetFormByID(id)
			if err != nil {
				http.Error(w, "форма не найдена: "+err.Error(), http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(form)
			return
		}

		http.Error(w, "не поддерживается", http.StatusNotFound)
	})

	// проверка
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("робит"))
	})
}
