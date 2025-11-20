package handlers

import (
	"net/http"

	"proyecto_final/cmd/repositories"
)

type SubjectHandler struct {
	repo repositories.SubjectRespository
}

func NewSubjectHandler(repo repositories.SubjectRespository) *SubjectHandler {
	return &SubjectHandler{repo: repo}
}

func (h *SubjectHandler) GetAllSubjects(w http.ResponseWriter, r *http.Request) {
}

func (h *SubjectHandler) GetSubjectByID(w http.ResponseWriter, r *http.Request) {
}

func (h *SubjectHandler) CreateSubjects(w http.ResponseWriter, r *http.Request) {
}

func (h *SubjectHandler) UpdateSubject(w http.ResponseWriter, r *http.Request) {
}

func (h *SubjectHandler) DeleteSubject(w http.ResponseWriter, r *http.Request) {
}
