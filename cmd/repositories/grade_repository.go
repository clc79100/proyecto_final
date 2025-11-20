package repositories

import (
	"proyecto_final/cmd/models"
)

type GradeRepository interface {
	GetAllByStudentID(studentID int) ([]models.Grade, error)
	GetByStudentIDAndSubjectID(studentID int, subjectID int) (models.Grade, error)
	// Create(grade models.Grade) (models.Grade, error)
	// Update(id int, grade models.Grade) (models.Grade, error)
	// Delete(id int) error
}
