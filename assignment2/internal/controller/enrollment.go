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

type enrollment struct {
	router *mux.Router
	repo   *repository.Repo
}

func NewEnrollmentRouter(router *mux.Router, repo *repository.Repo) {
	i := enrollment{router: router, repo: repo}

	router.HandleFunc("/{id}", i.getByID).Methods("GET")
	router.HandleFunc("", i.getAll).Methods("GET")
	router.HandleFunc("", i.createEnrollment).Methods("POST")
	router.HandleFunc("/{id}", i.updateEnrollment).Methods("PUT")
	router.HandleFunc("/{id}", i.deleteEnrollment).Methods("DELETE")
}

func (e enrollment) getAll(w http.ResponseWriter, r *http.Request) {

	res, err := e.repo.GetEnrollments()
	if res == nil {
		log.Println("enrollments not found")
		http.Error(w, "enrollments not found", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Println("Error retrieving enrollments")
		http.Error(w, "Error retrieving enrollments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (e enrollment) getByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		log.Println("Missing or invalid enrollment ID")
		http.Error(w, "Missing or invalid enrollment ID", http.StatusBadRequest)
		return
	}

	res, err := e.repo.GetEnrollmentByID(id)
	if res == nil {
		log.Println("enrollment not found")
		http.Error(w, "enrollment not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("Error retrieving enrollment")
		http.Error(w, "Error retrieving enrollment", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (e enrollment) createEnrollment(w http.ResponseWriter, r *http.Request) {
	var reqEnrollment entity.Enrollment
	_ = json.NewDecoder(r.Body).Decode(&reqEnrollment)

	err := e.repo.EnrollStudent(&reqEnrollment)
	if err != nil {
		log.Println("Error creating a enrollment")
		http.Error(w, "Error creating a enrollment", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "successfully created enrollment ")

}

func (e enrollment) updateEnrollment(w http.ResponseWriter, r *http.Request) {
	var reqEnrollment entity.Enrollment
	_ = json.NewDecoder(r.Body).Decode(&reqEnrollment)
	reqID := mux.Vars(r)["id"]
	if reqID == "" {
		log.Println("Missing or invalid enrollment ID")
		http.Error(w, "Missing or invalid enrollment ID", http.StatusBadRequest)
		return
	}
	temp, _ := strconv.Atoi(reqID)
	reqEnrollment.ID = temp
	id, err := e.repo.UpdateEnrollment(&reqEnrollment)
	if err != nil {
		log.Println("Error updating a enrollment")
		http.Error(w, "Error updating a enrollment", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, fmt.Sprintf("successfully updated enrollment with id: %v", id))

}

func (e enrollment) deleteEnrollment(w http.ResponseWriter, r *http.Request) {
	reqID := mux.Vars(r)["id"]
	if reqID == "" {
		log.Println("Missing or invalid enrollment ID")
		http.Error(w, "Missing or invalid enrollment ID", http.StatusBadRequest)
		return
	}

	err := e.repo.DeleteEnrollment(reqID)
	if err != nil {
		log.Println("Error deleting a enrollment")
		http.Error(w, "Error deleting a enrollment", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, fmt.Sprintf("successfully deleted enrollment with id: %v", reqID))

}
