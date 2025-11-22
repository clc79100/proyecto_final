package models

type Grade struct {
	GradeID   int     `json:"grade_id"`
	SubjectID int     `json:"subject_id"`
	Subject   string  `json:"subject"`
	Grade     float32 `json:"grade"`
}
