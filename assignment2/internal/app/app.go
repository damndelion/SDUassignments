package app

import (
	"github.com/damndelion/SDUassignments/assignment2/db"
	"github.com/damndelion/SDUassignments/assignment2/internal/controller"
	"github.com/damndelion/SDUassignments/assignment2/internal/entity"
	"github.com/damndelion/SDUassignments/assignment2/internal/repository"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Run() {
	db, err := db.ConnectMySql()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	db.AutoMigrate(&entity.Course{}, &entity.Department{}, &entity.Enrollment{}, &entity.Instructor{}, &entity.Student{})

	router := mux.NewRouter()
	controller.NewRouter(router, repository.NewRepo(db))

	http.ListenAndServe(":8080", router)
}
