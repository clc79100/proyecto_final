package repositories

import (
	"database/sql"
	"errors"

	"proyecto_final/cmd/models"
)

type MYSQLSubjectRepository struct {
	db *sql.DB
}

func NewMySQLSubjectRepository(db *sql.DB) *MYSQLSubjectRepository {
	return &MYSQLSubjectRepository{db: db}
}

func (r *MYSQLSubjectRepository) GetAll() ([]models.Subject, error) {
	rows, err := r.db.Query("SELECT subject_id, name FROM Subject;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var subjects []models.Subject
	for rows.Next() {
		var subject models.Subject
		err := rows.Scan(&subject.SubjectID, &subject.Name)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, subject)
	}
	return subjects, nil
}

func (r *MYSQLSubjectRepository) GetByID(id int) (models.Subject, error) {
	var subject models.Subject
	row := r.db.QueryRow("SELECT * subject_id, name FROM Subject WHERE subject_id = ?;", subject.SubjectID)
	err := row.Scan(&subject.SubjectID, &subject.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Subject{}, errors.New("subject not found")
		}
		return models.Subject{}, err
	}
	return subject, nil
}

func (r *MYSQLSubjectRepository) Create(subject models.Subject) (models.Subject, error) {
	result, err := r.db.Exec("INSERT INTO Subject(name) VALUE (?);", subject.Name)
	if err != nil {
		return models.Subject{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Subject{}, err
	}
	subject.SubjectID = int(id)
	return subject, nil
}

func (r *MYSQLSubjectRepository) Update(id int, subject models.Subject) (models.Subject, error) {
	_, err := r.db.Exec("UPDATE Subject SET name = ? WHERE subject_id = ?", subject.Name, id)
	if err != nil {
		return models.Subject{}, err
	}

	subject.SubjectID = int(id)

	return subject, nil
}

func (r *MYSQLSubjectRepository) Delete(id int) error {
	resut, err := r.db.Exec("DELETE FROM Subject WHERE subject_id = ?", id)
	if err != nil {
		return err
	}

	rows, err := resut.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("subject not found")
	}

	return nil
}
