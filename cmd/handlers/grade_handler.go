package handlers

import (
	"net/http"

	"proyecto_final/cmd/repositories"
)

type GradeHandler struct {
	repo repositories.GradeRepository
}

func NewGradeHandler(repo repositories.GradeRepository) *GradeHandler {
	return &GradeHandler{repo: repo}
}

func (h *GradeHandler) GetAllGradesByStudentID(w http.ResponseWriter, r http.Request) {}

func (h *GradeHandler) GetGradeByStudentIDAndSubjectID(w http.ResponseWriter, r http.Request) {}

func (h *GradeHandler) CreateGrade(w http.ResponseWriter, r http.Request) {}

func (h *GradeHandler) UpdateGrade(w http.ResponseWriter, r http.Request) {}

func (h *GradeHandler) DeleteGrade(w http.ResponseWriter, r http.Request) {}
