package repositories

import (
	"database/sql"
	"errors"

	"proyecto_final/cmd/models"
)

type MySQLStudentRepository struct {
	db *sql.DB
}

func NewMySQLStudentRepository(db *sql.DB) *MySQLStudentRepository {
	return &MySQLStudentRepository{db: db}
}

func (r *MySQLStudentRepository) GetAll() ([]models.Student, error) {
	rows, err := r.db.Query("SELECT student_id, name, group_name, email FROM Student;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.StudentID, &student.Name, &student.Email, &student.Group)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (r *MySQLStudentRepository) GetByID(id int) (models.Student, error) {
	var student models.Student
	row := r.db.QueryRow("SELECT student_id, name, group_name, email FROM Student WHERE student_id = ?;", id)
	err := row.Scan(&student.StudentID, &student.Name, &student.Email, &student.Group)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Student{}, errors.New("student not found")
		}
		return models.Student{}, err
	}
	return student, nil
}

func (r *MySQLStudentRepository) Create(student models.Student) (models.Student, error) {
	resutl, err := r.db.Exec("INSERT INTO Student (name, group_name, email) VALUE (?, ?, ?);", student.Name, student.Group, student.Email)
	if err != nil {
		return models.Student{}, err
	}

	id, err := resutl.LastInsertId()
	if err != nil {
		return models.Student{}, err
	}
	student.StudentID = int(id)
	return student, nil
}

func (r *MySQLStudentRepository) Update(id int, student models.Student) (models.Student, error) {
	_, err := r.db.Exec("UPDATE Student SET  name = ?, group_name = ?, email = ? WHERE student_id = ?;", student.Name, student.Group, student.Email, id)
	if err != nil {
		return models.Student{}, err
	}

	student.StudentID = int(id)

	return student, nil
}

func (r *MySQLStudentRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM Student WHERE student_id = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("student not found")
	}

	return nil
}
