package service

import (
	"rbs-feedbox/internal/models"
	"rbs-feedbox/internal/storage/postgres"
)

type Service struct {
	storage *postgres.Storagepostgres
}

func New(storage *postgres.Storagepostgres) *Service {
	return &Service{storage: storage}
}

func (s *Service) CreateForm(name, schema string) error {
	return s.storage.CreateForm(name, schema)
}

func (s *Service) GetForms() ([]models.Form, error) {
	pgForms, err := s.storage.GetForms()
	if err != nil {
		return nil, err
	}

	var result []models.Form
	for _, f := range pgForms {
		result = append(result, models.Form{
			ID:        f.ID,
			Name:      f.Name,
			Schema:    f.Schema,
			CreatedAt: f.CreatedAt,
		})
	}

	return result, nil
}

func (s *Service) GetFormByID(id int) (models.Form, error) {
	f, err := s.storage.GetFormByID(id)
	if err != nil {
		return models.Form{}, err
	}

	return models.Form{
		ID:        f.ID,
		Name:      f.Name,
		Schema:    f.Schema,
		CreatedAt: f.CreatedAt,
	}, nil
}

func (s *Service) SaveResponse(formID int, data string) error {
	return s.storage.SaveResponse(formID, data)
}

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

func (s *Service) UpdateResponseStatus(id int, status string) error {
	return s.storage.UpdateResponseStatus(id, status)
}
