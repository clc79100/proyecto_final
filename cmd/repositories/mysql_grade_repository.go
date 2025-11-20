package repositories

import (
	"database/sql"
	"errors"

	"proyecto_final/cmd/models"
)

type MySQLGradeRepository struct {
	db *sql.DB
}

func NewSQLGradeRepository(db *sql.DB) *MySQLGradeRepository {
	return &MySQLGradeRepository{db: db}
}

func (r *MySQLGradeRepository) GetAllByStudentID(studentID int) ([]models.Grade, error) {
	rows, err := r.db.Query("SELECT G.grade_id, Sb.name AS subject, G.grade FROM Student S INNER JOIN Grade G ON S.student_id = G.student_id INNER JOIN Subject Sb ON G.subject_id = Sb.subject_id WHERE S.student_id = ? ;", studentID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var grades []models.Grade
	for rows.Next() {
		var grade models.Grade
		err := rows.Scan(&grade.GradeID, &grade.Subject, &grade.Grade)
		if err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}
	return grades, nil
}

func (r *MySQLGradeRepository) GetByStudentIDAndSubjectID(studentID int, subjectID int) (models.Grade, error) {
	var grade models.Grade
	row := r.db.QueryRow("SELECT G.grade_id, Sb.name AS subject, G.grade FROM Student S INNER JOIN Grade G ON S.student_id = G.student_id INNER JOIN Subject Sb ON G.subject_id = Sb.subject_id WHERE S.student_id = ? AND Sb.subject_id = ?", studentID, subjectID)
	err := row.Scan(&grade.GradeID, &grade.Subject, &grade.Grade)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Grade{}, errors.New("grade not found")
		}
		return models.Grade{}, err
	}

	return grade, nil
}
