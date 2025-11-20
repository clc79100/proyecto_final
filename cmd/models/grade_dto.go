package models

type GradeDTO struct {
	GradeID   int     `json:"grade_id"`
	SubjectID int     `json:"subject_id"`
	StudentID int     `json:"Student_id"`
	Grade     float32 `json:"grade"`
}
