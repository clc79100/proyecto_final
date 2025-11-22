package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"proyecto_final/cmd/models"
	"proyecto_final/cmd/repositories"
)

type SubjectHandler struct {
	repo repositories.SubjectRespository
}

func NewSubjectHandler(repo repositories.SubjectRespository) *SubjectHandler {
	return &SubjectHandler{repo: repo}
}

func (h *SubjectHandler) GetSubjectByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	subject, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subject)
}

func (h *SubjectHandler) CreateSubjects(w http.ResponseWriter, r *http.Request) {
	var subject models.Subject
	err := json.NewDecoder(r.Body).Decode(&subject)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdSubject, err := h.repo.Create(subject)
	if err != nil {
		http.Error(w, "Failed to create subject", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdSubject)
}

func (h *SubjectHandler) UpdateSubject(w http.ResponseWriter, r *http.Request) {
	var subject models.Subject
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&subject)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedSubject, err := h.repo.Update(id, subject)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedSubject)
}

func (h *SubjectHandler) DeleteSubject(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = h.repo.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
