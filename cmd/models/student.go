package models

type Student struct {
	StudentID int    `json:"student_id"`
	Name      string `json:"name"`
	Group     string `json:"group_name"`
	Email     string `json:"email"`
}
