package service

import (
	"rbs-feedbox/internal/models"
	"rbs-feedbox/internal/storage/postgres"
)

// Service - бизнес-логика и взаимодействие с хранилищем данных
type Service struct {
	storage *postgres.StoragePostgres
}

// New - конструктор для создания нового экземпляра Service
// принимает хранилище данных и возвращает новый объект Service
func New(storage *postgres.StoragePostgres) *Service {
	return &Service{storage: storage}
}

// CreateForm - создаёт новую форму
// принимает название и схему формы, вызывает соответствующий метод в хранилище
func (s *Service) CreateForm(name, schema, description string, projectID int) error {
	return s.storage.CreateForm(name, schema, description, projectID)
}

// GetForms - возвращает список всех форм
// получает формы из хранилища, преобразует их в модели и возвращает
func (s *Service) GetForms() ([]models.Form, error) {
	pgForms, err := s.storage.GetForms()
	if err != nil {
		return nil, err
	}

	var result []models.Form
	for _, f := range pgForms {
		result = append(result, models.Form{
			ID:          f.ID,
			ProjectID:   f.ProjectID,
			Title:       f.Title,
			Description: f.Description,
			Schema:      f.Schema,
			CreatedAt:   f.CreatedAt,
		})
	}

	return result, nil
}

// GetFormByID - возвращает форму по её ID
// получает данные из хранилища и преобразует в модель Form
func (s *Service) GetFormByID(id int) (models.Form, error) {
	f, err := s.storage.GetFormByID(id)
	if err != nil {
		return models.Form{}, err
	}

	return models.Form{
		ID:          f.ID,
		ProjectID:   f.ProjectID,
		Title:       f.Title,
		Description: f.Description,
		Schema:      f.Schema,
		CreatedAt:   f.CreatedAt,
	}, nil
}

// SaveResponse - сохраняет ответ пользователя для формы
// принимает ID формы и данные ответа в виде строки
func (s *Service) SaveResponse(formID int, data string) error {
	return s.storage.SaveResponse(formID, data)
}

// GetResponsesByFormID - получает все ответы по конкретной форме
// возвращает список структур Response
func (s *Service) GetResponsesByFormID(formID int) ([]models.Response, error) {
	pgResp, err := s.storage.GetResponsesByFormID(formID)
	if err != nil {
		return nil, err
	}

	var result []models.Response
	for _, r := range pgResp {
		result = append(result, models.Response{
			ID:        r.ID,
			FormID:    r.FormID,
			Data:      r.Data,
			CreatedAt: r.CreatedAt,
			Status:    r.Status,
		})
	}

	return result, nil
}

// UpdateResponseStatus - обновляет статус конкретного ответа по его ID
func (s *Service) UpdateResponseStatus(id int, status string) error {
	return s.storage.UpdateResponseStatus(id, status)
}

// CreateProject создаёт новый проект
func (s *Service) CreateProject(title, description string) error {
	return s.storage.CreateProject(title, description)
}

func (s *Service) GetProjects() ([]postgres.Project, error) {
	pgProjects, err := s.storage.GetProjects()
	if err != nil {
		return nil, err
	}

	var projects []postgres.Project
	for _, p := range pgProjects {
		projects = append(projects, postgres.Project{
			ID:          p.ID,
			Title:       p.Title,
			Description: p.Description,
			CreatedAt:   p.CreatedAt,
		})
	}
	return projects, nil

}
func (s *Service) GetFormsByProjectID(projectID int) ([]postgres.Form, error) {
	return s.storage.GetFormsByProjectID(projectID)
}
