package repositories

import (
	"proyecto_final/cmd/models"
)

type SubjectRespository interface {
	GetByID(id int) (models.Subject, error)
	Create(subject models.Subject) (models.Subject, error)
	Update(id int, subject models.Subject) (models.Subject, error)
	Delete(id int) error
}
