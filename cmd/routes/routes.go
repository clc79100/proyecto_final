package routes

import (
	"net/http"

	"proyecto_final/cmd/handlers"
)

func SetupRoutes(studentHandler *handlers.StudentHandler, subjectHandler *handlers.SubjectHandler, grade *handlers.GradeHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /students/", studentHandler.GetAllStudents)
	mux.HandleFunc("GET /students/{id}", studentHandler.GetStudentByID)
	mux.HandleFunc("POST /students/", studentHandler.CreateStudent)
	mux.HandleFunc("PUT /students/{id}", studentHandler.UpdateStudent)
	mux.HandleFunc("DELETE /students/{id}", studentHandler.DeleteStudent)

	mux.HandleFunc("GET /subjects/", subjectHandler.GetAllSubjects)
	mux.HandleFunc("GET /subjects/{id}", subjectHandler.GetSubjectByID)
	mux.HandleFunc("POST /subjects/", subjectHandler.CreateSubjects)
	mux.HandleFunc("PUT /subjects/{id}", subjectHandler.UpdateSubject)
	mux.HandleFunc("DELETE /subjects/{id}", subjectHandler.DeleteSubject)

	return mux
}
