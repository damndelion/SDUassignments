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

type course struct {
	router *mux.Router
	repo   *repository.Repo
}

func NewCourseRouter(router *mux.Router, repo *repository.Repo) {
	c := course{router: router, repo: repo}

	router.HandleFunc("/{id}", c.getByID).Methods("GET")
	router.HandleFunc("", c.getAll).Methods("GET")
	router.HandleFunc("", c.createCourse).Methods("POST")
	router.HandleFunc("/{id}", c.updateCourse).Methods("PUT")
	router.HandleFunc("/{id}", c.deleteCourse).Methods("DELETE")
}

func (c course) getAll(w http.ResponseWriter, r *http.Request) {
	res, err := c.repo.GetCourses()

	if res == nil {
		log.Println("Courses not found")
		http.Error(w, "Courses not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("Error retrieving courses")
		http.Error(w, "Error retrieving courses", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (c course) getByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		log.Println("Missing or invalid course ID")
		http.Error(w, "Missing or invalid course ID", http.StatusBadRequest)
		return
	}

	res, err := c.repo.GetCourseByID(id)
	if res == nil {
		log.Println("course not found")
		http.Error(w, "course not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("Error retrieving course")
		http.Error(w, "Error retrieving course", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (c course) createCourse(w http.ResponseWriter, r *http.Request) {
	var reqCourse entity.Course
	_ = json.NewDecoder(r.Body).Decode(&reqCourse)

	id, err := c.repo.CreateCourse(&reqCourse)
	if err != nil {
		log.Println("Error creating a course")
		http.Error(w, "Error creating a course", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, fmt.Sprintf("successfully created course with id: %v", id))

}

func (c course) updateCourse(w http.ResponseWriter, r *http.Request) {
	var reqCourse entity.Course
	_ = json.NewDecoder(r.Body).Decode(&reqCourse)
	reqID := mux.Vars(r)["id"]
	if reqID == "" {
		log.Println("Missing or invalid course ID")
		http.Error(w, "Missing or invalid course ID", http.StatusBadRequest)
		return
	}
	temp, _ := strconv.Atoi(reqID)
	reqCourse.ID = temp
	id, err := c.repo.UpdateCourse(&reqCourse)
	if err != nil {
		log.Println("Error updating a course")
		http.Error(w, "Error updating a course", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, fmt.Sprintf("successfully updated course with id: %v", id))

}

func (c course) deleteCourse(w http.ResponseWriter, r *http.Request) {
	reqID := mux.Vars(r)["id"]
	if reqID == "" {
		log.Println("Missing or invalid course ID")
		http.Error(w, "Missing or invalid course ID", http.StatusBadRequest)
		return
	}

	err := c.repo.DeleteCourse(reqID)
	if err != nil {
		log.Println("Error deleting a course")
		http.Error(w, "Error deleting a course", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, fmt.Sprintf("successfully deleted course with id: %v", reqID))

}
