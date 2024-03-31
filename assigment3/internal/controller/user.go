package controller

import (
	"encoding/json"
	"github.com/damndelion/SDUassignments/assigment3/internal/entity"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type user struct {
	router *mux.Router
	db     *gorm.DB
}

func NewUserRouter(router *mux.Router, db *gorm.DB) {
	u := user{router: router, db: db}

	router.HandleFunc("", u.getAll).Methods("GET")

}

func (u user) getAll(w http.ResponseWriter, r *http.Request) {

	var users []*entity.User
	err := u.db.Find(&users).Error
	if err != nil {
		log.Println("Error retrieving students")
		http.Error(w, "Error retrieving student", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
