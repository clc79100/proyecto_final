package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"proyecto_final/cmd/models"
	"proyecto_final/cmd/repositories"
)

type GradeHandler struct {
	repo repositories.GradeRepository
}

func NewGradeHandler(repo repositories.GradeRepository) *GradeHandler {
	return &GradeHandler{repo: repo}
}

func (h *GradeHandler) GetAllGradesByStudentID(w http.ResponseWriter, r *http.Request) {
	studentID, err := strconv.Atoi(r.PathValue("student_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	grades, err := h.repo.GetAllByStudentID(studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(grades)
}

func (h *GradeHandler) GetGradeByStudentIDAndSubjectID(w http.ResponseWriter, r *http.Request) {
	gradeID, err := strconv.Atoi(r.PathValue("grade_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	studentID, err := strconv.Atoi(r.PathValue("student_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grade, err := h.repo.GetByStudentIDAndSubjectID(gradeID, studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(grade)
}

func (h *GradeHandler) CreateGrade(w http.ResponseWriter, r *http.Request) {
	var grade models.GradeDTO
	err := json.NewDecoder(r.Body).Decode(&grade)
	if err != nil {
		http.Error(w, "Invalid Request payload", http.StatusBadRequest)
		return
	}

	createdGrade, err := h.repo.Create(grade)
	if err != nil {
		http.Error(w, "Failed to create Grade", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdGrade)
}

func (h *GradeHandler) UpdateGrade(w http.ResponseWriter, r *http.Request) {
	var grade models.GradeDTO
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&grade)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedGrade, err := h.repo.Update(id, grade)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedGrade)
}

func (h *GradeHandler) DeleteGrade(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.repo.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
