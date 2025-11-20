package repositories

import (
	"proyecto_final/cmd/models"
)

type StudentRepository interface {
	GetAll() ([]models.Student, error)
	GetByID(id int) (models.Student, error)
	Create(student models.Student) (models.Student, error)
	Update(id int, student models.Student) (models.Student, error)
	Delete(id int) error
}
