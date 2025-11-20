package models

type Subject struct {
	SubjectID int     `json:"subject_id"`
	Name      string  `json:"name"`
	Grade     float32 `json:"grade"`
}
