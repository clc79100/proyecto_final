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
	rows, err := r.db.Query("SELECT G.grade_id, Sb.subject_id, Sb.name AS subject, G.grade FROM Student S INNER JOIN Grade G ON S.student_id = G.student_id INNER JOIN Subject Sb ON G.subject_id = Sb.subject_id WHERE S.student_id = ? ;", studentID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var grades []models.Grade
	for rows.Next() {
		var grade models.Grade
		err := rows.Scan(&grade.GradeID, &grade.SubjectID, &grade.Subject, &grade.Grade)
		if err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}
	return grades, nil
}

func (r *MySQLGradeRepository) GetByStudentIDAndSubjectID(gradeID int, studentID int) (models.Grade, error) {
	var grade models.Grade
	row := r.db.QueryRow("SELECT G.grade_id, Sb.subject_id, Sb.name AS subject, G.grade FROM Student S INNER JOIN Grade G ON S.student_id = G.student_id INNER JOIN Subject Sb ON G.subject_id = Sb.subject_id WHERE S.student_id = ? AND G.grade_id = ?", studentID, gradeID)
	err := row.Scan(&grade.GradeID, &grade.SubjectID, &grade.Subject, &grade.Grade)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Grade{}, errors.New("grade not found")
		}
		return models.Grade{}, err
	}

	return grade, nil
}

func (r *MySQLGradeRepository) Create(gradeDTO models.GradeDTO) (models.GradeDTO, error) {
	result, err := r.db.Exec("INSERT INTO Grade (student_id, subject_id, grade) VALUE (?, ?, ?)", gradeDTO.StudentID, gradeDTO.SubjectID, gradeDTO.Grade)
	if err != nil {
		return models.GradeDTO{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.GradeDTO{}, err
	}
	gradeDTO.GradeID = int(id)
	return gradeDTO, nil
}

func (r *MySQLGradeRepository) Update(id int, gradeDTO models.GradeDTO) (models.GradeDTO, error) {
	_, err := r.db.Exec("UPDATE Grade SET grade = ?", gradeDTO.Grade)
	if err != nil {
		return models.GradeDTO{}, err
	}
	gradeDTO.GradeID = int(id)

	return gradeDTO, nil
}

func (r *MySQLGradeRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM Grade WHERE grade_id = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("grade not found")
	}

	return nil
}
