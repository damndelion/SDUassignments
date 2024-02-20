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

type department struct {
	router *mux.Router
	repo   *repository.Repo
}

func NewDepartmentRouter(router *mux.Router, repo *repository.Repo) {
	d := department{router: router, repo: repo}

	router.HandleFunc("/{id}", d.getByID).Methods("GET")
	router.HandleFunc("", d.getAll).Methods("GET")
	router.HandleFunc("", d.createDepartment).Methods("POST")
	router.HandleFunc("/{id}", d.updateDepartment).Methods("PUT")
	router.HandleFunc("/{id}", d.deleteDepartment).Methods("DELETE")
}

func (i department) getAll(w http.ResponseWriter, r *http.Request) {

	res, err := i.repo.GetDepartments()
	if res == nil {
		log.Println("department not found")
		http.Error(w, "department not found", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Println("Error retrieving departments")
		http.Error(w, "Error retrieving departments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (i department) getByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		log.Println("Missing or invalid department ID")
		http.Error(w, "Missing or invalid department ID", http.StatusBadRequest)
		return
	}

	res, err := i.repo.GetDepartmentByID(id)
	if res == nil {
		log.Println("department not found")
		http.Error(w, "department not found", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Println("Error retrieving department")
		http.Error(w, "Error retrieving department", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (i department) createDepartment(w http.ResponseWriter, r *http.Request) {
	var reqDepartment entity.Department
	_ = json.NewDecoder(r.Body).Decode(&reqDepartment)

	id, err := i.repo.CreateDepartment(&reqDepartment)
	if err != nil {
		log.Println("Error creating a department")
		http.Error(w, "Error creating a department", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, fmt.Sprintf("successfully created department with id: %v", id))

}

func (i department) updateDepartment(w http.ResponseWriter, r *http.Request) {
	var reqDepartment entity.Department
	_ = json.NewDecoder(r.Body).Decode(&reqDepartment)
	reqID := mux.Vars(r)["id"]
	if reqID == "" {
		log.Println("Missing or invalid department ID")
		http.Error(w, "Missing or invalid department ID", http.StatusBadRequest)
		return
	}

	temp, _ := strconv.Atoi(reqID)
	reqDepartment.ID = temp
	id, err := i.repo.UpdateDepartment(&reqDepartment)
	if err != nil {
		log.Println("Error updating a department")
		http.Error(w, "Error updating a department", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, fmt.Sprintf("successfully updated department with id: %v", id))

}

func (i department) deleteDepartment(w http.ResponseWriter, r *http.Request) {
	reqID := mux.Vars(r)["id"]
	if reqID == "" {
		log.Println("Missing or invalid department ID")
		http.Error(w, "Missing or invalid department ID", http.StatusBadRequest)
		return
	}

	err := i.repo.DeleteDepartment(reqID)
	if err != nil {
		log.Println("Error deleting a department")
		http.Error(w, "Error deleting a department", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, fmt.Sprintf("successfully deleted department with id: %v", reqID))

}
