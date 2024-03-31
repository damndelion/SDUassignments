package app

import (
	"github.com/damndelion/SDUassignments/assigment3/internal/controller"
	"github.com/damndelion/SDUassignments/assignment2/db"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/gorilla/mux"
	"net/http"
)

func Run() {
	db, err := db.ConnectMySql()
	if err != nil {
		logrus.Fatalf("[ERROR] Error connecting to database %v", err)
	}
	logrus.Info("[DB] Connected to database")
	err = godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("[ERROR] Error loading .env file: %v", err)
		return
	}

	jwtKey := os.Getenv("JWT_KEY")

	router := mux.NewRouter()
	controller.NewRouter(router, db, jwtKey)

	logrus.Info("[HTTP] HTTP server listening on locahost:8080 ")
	http.ListenAndServe(":8080", router)
}
