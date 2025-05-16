package service

import "rbs-feedbox/internal/models"

type Storage interface {
	CreateForm(name, schema string) error
	GetForms() ([]models.Form, error)
	GetFormByID(id int) (models.Form, error)
	SaveResponse(formID int, data string) error
	GetResponsesByFormID(formID int) ([]models.Response, error)
	UpdateResponseStatus(id int, status string) error
}
