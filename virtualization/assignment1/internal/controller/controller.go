package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/damndelion/SDUassignments/virtualization/assignment1/internal/entity"
	"github.com/gorilla/mux"
	"net/http"
)

type studentRouter struct {
	mux *mux.Router
	db  *sql.DB
}

func NewRouter(mux *mux.Router, db *sql.DB) {
	router := &studentRouter{mux, db}
	mux.HandleFunc("/get", router.getAll).Methods("GET")
	mux.HandleFunc("/get/{id}", router.getByID).Methods("GET")
	mux.HandleFunc("/create", router.createStudent).Methods("POST")
	mux.HandleFunc("/update/{id}", router.update).Methods("PUT")
	mux.HandleFunc("/delete/{id}", router.delete).Methods("DELETE")

}

func (router studentRouter) getByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	stud := entity.Student{}
	row := router.db.QueryRow("SELECT * FROM students where id = ?", id)
	err := row.Scan(&stud.ID, &stud.Name, &stud.Age, &stud.Course, &stud.Faculty)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(w, stud)
}

func (router studentRouter) getAll(w http.ResponseWriter, r *http.Request) {

	rows, err := router.db.Query("SELECT * FROM students")
	if err != nil {
		http.Error(w, "Error retrieving students", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	students := []entity.Student{}

	for rows.Next() {
		stud := entity.Student{}
		err := rows.Scan(&stud.ID, &stud.Name, &stud.Age, &stud.Course, &stud.Faculty)
		if err != nil {
			http.Error(w, "Error reading student data", http.StatusInternalServerError)
			return
		}
		students = append(students, stud)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating student rows", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(students)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (router studentRouter) createStudent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newStudent entity.Student
	err := decoder.Decode(&newStudent)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	stmt, err := router.db.Prepare("INSERT INTO students (name, age, course, faculty) VALUES (?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Error preparing insert statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(newStudent.Name, newStudent.Age, newStudent.Course, newStudent.Faculty)
	if err != nil {
		http.Error(w, "Error saving student data", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Student created successfully!")
}

func (router studentRouter) update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	decoder := json.NewDecoder(r.Body)
	var updatedStudent entity.Student
	err := decoder.Decode(&updatedStudent)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	stmt, err := router.db.Prepare("UPDATE students SET name = ?, age = ?, course = ?, faculty = ? WHERE id = ?")
	if err != nil {
		http.Error(w, "Error preparing update statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(updatedStudent.Name, updatedStudent.Age, updatedStudent.Course, updatedStudent.Faculty, id)
	if err != nil {
		http.Error(w, "Error updating student data", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Student updated successfully!")
}

func (router studentRouter) delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	stmt, err := router.db.Prepare("DELETE FROM students WHERE id = ?")
	if err != nil {
		http.Error(w, "Error preparing delete statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		http.Error(w, "Error deleting student data", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error checking deleted rows", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Student with provided ID not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Student deleted successfully!")
}
