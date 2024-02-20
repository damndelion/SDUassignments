package controller

import (
	"encoding/json"
	"fmt"
	"github.com/damndelion/SDUassignments/assignment2/internal/entity"
	"github.com/damndelion/SDUassignments/assignment2/internal/repository"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type student struct {
	router *mux.Router
	repo   *repository.Repo
}

func NewStudentRouter(router *mux.Router, repo *repository.Repo) {
	s := student{router: router, repo: repo}

	router.HandleFunc("/{id}", s.getByID).Methods("GET")
	router.HandleFunc("", s.getAll).Methods("GET")
	router.HandleFunc("", s.createStudent).Methods("POST")
	router.HandleFunc("/{id}", s.updateStudent).Methods("PUT")
	router.HandleFunc("/{id}", s.deleteStudent).Methods("DELETE")
	router.HandleFunc("/department/{id}", s.getByDepartmentID).Methods("GET")
	router.HandleFunc("/{id}/enrollment", s.getStudentEnrollments).Methods("GET")
}

func (s student) getAll(w http.ResponseWriter, r *http.Request) {

	res, err := s.repo.GetStudents()
	if res == nil {
		log.Println("Students not found")
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("Error retrieving students")
		http.Error(w, "Error retrieving student", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (s student) getByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		log.Println("Missing or invalid student ID")
		http.Error(w, "Missing or invalid student ID", http.StatusBadRequest)
		return
	}

	res, err := s.repo.GetStudentByID(id)
	if res == nil {
		log.Println("Student not found")
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("Error retrieving student")
		http.Error(w, "Error retrieving student", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (s student) createStudent(w http.ResponseWriter, r *http.Request) {
	var reqStudent entity.Student
	_ = json.NewDecoder(r.Body).Decode(&reqStudent)

	id, err := s.repo.CreateStudent(&reqStudent)
	if err != nil {
		log.Println("Error creating a student")
		http.Error(w, "Error creating a student", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, fmt.Sprintf("successfully created student with id: %v", id))

}

func (s student) updateStudent(w http.ResponseWriter, r *http.Request) {
	var reqStudent entity.Student
	_ = json.NewDecoder(r.Body).Decode(&reqStudent)
	reqID := mux.Vars(r)["id"]
	if reqID == "" {
		log.Println("Missing or invalid student ID")
		http.Error(w, "Missing or invalid student ID", http.StatusBadRequest)
		return
	}
	temp, _ := strconv.Atoi(reqID)
	reqStudent.ID = temp
	id, err := s.repo.UpdateStudent(&reqStudent)
	if err != nil {
		log.Println("Error updating a student")
		http.Error(w, "Error updating a student", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, fmt.Sprintf("successfully updated student with id: %v", id))

}

func (s student) deleteStudent(w http.ResponseWriter, r *http.Request) {
	reqID := mux.Vars(r)["id"]
	if reqID == "" {
		log.Println("Missing or invalid student ID")
		http.Error(w, "Missing or invalid student ID", http.StatusBadRequest)
		return
	}

	err := s.repo.DeleteStudent(reqID)
	if err != nil {
		log.Println("Error deleting a student")
		http.Error(w, "Error deleting a student", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, fmt.Sprintf("successfully deleted student with id: %v", reqID))

}

func (s student) getByDepartmentID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	res, err := s.repo.GetStudentByDepID(id)
	if err != nil {
		log.Println("Error retrieving student")
		http.Error(w, "Error retrieving student", http.StatusInternalServerError)
		return
	}

	if res == nil {
		log.Println("Students not found")
		http.Error(w, "Students not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (s student) getStudentEnrollments(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	res, err := s.repo.GetStudentEnrollments(id)
	if err != nil {
		log.Println("Error retrieving enrollments")
		http.Error(w, "Error retrieving enrollments", http.StatusInternalServerError)
		return
	}

	if res == nil {
		log.Println("enrollments not found")
		http.Error(w, "enrollments not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
