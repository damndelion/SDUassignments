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

type instructor struct {
	router *mux.Router
	repo   *repository.Repo
}

func NewInstructorRouter(router *mux.Router, repo *repository.Repo) {
	i := instructor{router: router, repo: repo}

	router.HandleFunc("/{id}", i.getByID).Methods("GET")
	router.HandleFunc("", i.getAll).Methods("GET")
	router.HandleFunc("", i.createInstructor).Methods("POST")
	router.HandleFunc("/{id}", i.updateInstructor).Methods("PUT")
	router.HandleFunc("/{id}", i.deleteInstructor).Methods("DELETE")
	router.HandleFunc("/{id}/course", i.getCourses).Methods("GET")

}

func (i instructor) getAll(w http.ResponseWriter, r *http.Request) {

	res, err := i.repo.GetInstructors()
	if res == nil {
		log.Println("instructors not found")
		http.Error(w, "instructors not found", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Println("Error retrieving instructors")
		http.Error(w, "Error retrieving instructors", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (i instructor) getByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		log.Println("Missing or invalid instructor ID")
		http.Error(w, "Missing or invalid instructor ID", http.StatusBadRequest)
		return
	}

	res, err := i.repo.GetInstructorByID(id)
	if res == nil {
		log.Println("instructor not found")
		http.Error(w, "instructor not found", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Println("Error retrieving instructor")
		http.Error(w, "Error retrieving instructor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (i instructor) createInstructor(w http.ResponseWriter, r *http.Request) {
	var reqInstructor entity.Instructor
	_ = json.NewDecoder(r.Body).Decode(&reqInstructor)

	id, err := i.repo.CreateInstructor(&reqInstructor)
	if err != nil {
		log.Println("Error creating a instructor")
		http.Error(w, "Error creating a instructor", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, fmt.Sprintf("successfully created instructor with id: %v", id))

}

func (i instructor) updateInstructor(w http.ResponseWriter, r *http.Request) {
	var reqInstructor entity.Instructor
	_ = json.NewDecoder(r.Body).Decode(&reqInstructor)
	reqID := mux.Vars(r)["id"]
	if reqID == "" {
		log.Println("Missing or invalid instructor ID")
		http.Error(w, "Missing or invalid instructor ID", http.StatusBadRequest)
		return
	}
	temp, _ := strconv.Atoi(reqID)
	reqInstructor.ID = temp
	id, err := i.repo.UpdateInstructor(&reqInstructor)
	if err != nil {
		log.Println("Error updating a instructor")
		http.Error(w, "Error updating a instructor", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, fmt.Sprintf("successfully updated instructor with id: %v", id))

}

func (i instructor) deleteInstructor(w http.ResponseWriter, r *http.Request) {
	reqID := mux.Vars(r)["id"]
	if reqID == "" {
		log.Println("Missing or invalid instructor ID")
		http.Error(w, "Missing or invalid instructor ID", http.StatusBadRequest)
		return
	}

	err := i.repo.DeleteInstructor(reqID)
	if err != nil {
		log.Println("Error deleting a instructor")
		http.Error(w, "Error deleting a instructor", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, fmt.Sprintf("successfully deleted instructor with id: %v", reqID))
}

func (i instructor) getCourses(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	res, err := i.repo.GetCoursesByInstructor(id)

	if res == nil {
		log.Println("courses not found")
		http.Error(w, "courses not found", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Println("Error retrieving courses")
		http.Error(w, "Error retrieving courses", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
