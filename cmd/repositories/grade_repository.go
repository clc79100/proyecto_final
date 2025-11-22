package repositories

import (
	"proyecto_final/cmd/models"
)

type GradeRepository interface {
	GetAllByStudentID(studentID int) ([]models.Grade, error)
	GetByStudentIDAndSubjectID(gradeID int, studentID int) (models.Grade, error)
	Create(grade models.GradeDTO) (models.GradeDTO, error)
	Update(id int, grade models.GradeDTO) (models.GradeDTO, error)
	Delete(id int) error
}
