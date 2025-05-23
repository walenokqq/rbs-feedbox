package handlers

import (
	"encoding/json"
	"net/http"
	"rbs-feedbox/internal/service"
	"strconv"
	"strings"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
}

func Register(svc *service.Service) {
	http.HandleFunc("/api/responses/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
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
		err = svc.UpdateResponseStatus(id, req.Status)
		if err != nil {
			http.Error(w, "ошибка обновления статуса: "+err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"status": "Статус обновлён"})
	})

	http.HandleFunc("/api/forms", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if r.Method == http.MethodPost {
			var req struct {
				Title       string `json:"title"`
				Schema      string `json:"schema"`
				Description string `json:"description"`
				ProjectID   int    `json:"project_id"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "неверный JSON", http.StatusBadRequest)
				return
			}
			err := svc.CreateForm(req.Title, req.Schema, req.Description, req.ProjectID)
			if err != nil {
				http.Error(w, "ошибка создания формы: "+err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(map[string]string{"status": "Форма создана"})
			return
		}
		if r.Method == http.MethodGet {
			forms, err := svc.GetForms()
			if err != nil {
				http.Error(w, "ошибка получения форм: "+err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(forms)
			return
		}
		http.Error(w, "неподдерживаемый метод", http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/api/form/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		path := r.URL.Path

		if strings.HasSuffix(path, "/submit") && r.Method == http.MethodPost {
			parts := strings.Split(strings.TrimSuffix(path, "/submit"), "/")
			idStr := parts[len(parts)-1]
			formID, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "неверный ID формы", http.StatusBadRequest)
				return
			}
			var req struct {
				Data string `json:"data"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "неверный JSON", http.StatusBadRequest)
				return
			}
			err = svc.SaveResponse(formID, req.Data)
			if err != nil {
				http.Error(w, "ошибка сохранения ответа: "+err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(map[string]string{"status": "Ответ сохранён"})
			return
		}

		if strings.HasSuffix(path, "/responses") && r.Method == http.MethodGet {
			parts := strings.Split(strings.TrimSuffix(path, "/responses"), "/")
			idStr := parts[len(parts)-1]
			formID, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "неверный ID формы", http.StatusBadRequest)
				return
			}
			responses, err := svc.GetResponsesByFormID(formID)
			if err != nil {
				http.Error(w, "ошибка получения ответов: "+err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responses)
			return
		}

		if r.Method == http.MethodGet {
			idStr := strings.TrimPrefix(path, "/api/form/")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "неверный ID", http.StatusBadRequest)
				return
			}
			form, err := svc.GetFormByID(id)
			if err != nil {
				http.Error(w, "форма не найдена: "+err.Error(), http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(form)
			return
		}

		http.Error(w, "неподдерживаемый маршрут", http.StatusNotFound)
	})

	http.HandleFunc("/api/projects", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		switch r.Method {
		case http.MethodPost:
			var req struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			}

			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "неверный JSON: "+err.Error(), http.StatusBadRequest)
				return
			}

			if err := svc.CreateProject(req.Title, req.Description); err != nil {
				http.Error(w, "ошибка создания проекта: "+err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"status": "Проект успешно создан",
			})
			return

		case http.MethodGet:
			projects, err := svc.GetProjects()
			if err != nil {
				http.Error(w, "Ошибка получения проектов: "+err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(projects)

		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/api/projects/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if r.Method == http.MethodGet {
			parts := strings.Split(r.URL.Path, "/")
			if len(parts) < 4 {
				http.Error(w, "Неверный URL", http.StatusBadRequest)
				return
			}

			// Проверяем, если запрос /api/projects/{id}/forms
			if len(parts) >= 5 && parts[4] == "forms" {
				projectID, err := strconv.Atoi(parts[3])
				if err != nil {
					http.Error(w, "Неверный ID проекта", http.StatusBadRequest)
					return
				}

				forms, err := svc.GetFormsByProjectID(projectID)
				if err != nil {
					http.Error(w, "Ошибка получения форм: "+err.Error(), http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(forms)
				return
			}

			http.Error(w, "Не найден", http.StatusNotFound)
		}
	})

	// Отдаем статику по маршруту http://localhost:8080/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("frontend/dist"))
	})
}
