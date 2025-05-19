package service

import "rbs-feedbox/internal/models"

// Storage - интерфейс для абстрагирования хранилища данных
// определяет контракт, который должны реализовывать конкретные реализации хранилищ (в нашем случае - Postgres).
type Storage interface {
	CreateForm(name, schema string) error
	GetForms() ([]models.Form, error)
	GetFormByID(id int) (models.Form, error)
	SaveResponse(formID int, data string) error
	GetResponsesByFormID(formID int) ([]models.Response, error)
	UpdateResponseStatus(id int, status string) error
}
