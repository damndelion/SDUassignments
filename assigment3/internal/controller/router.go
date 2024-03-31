package controller

import (
	"encoding/json"
	"fmt"
	"github.com/damndelion/SDUassignments/assigment3/internal"
	"github.com/damndelion/SDUassignments/assigment3/internal/entity"
	"github.com/damndelion/SDUassignments/assigment3/internal/middleware"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strings"
	"time"
)

type router struct {
	mux    *mux.Router
	db     *gorm.DB
	jwtKey string
}

func NewRouter(mux *mux.Router, db *gorm.DB, jwtKey string) {
	r := &router{mux: mux, db: db, jwtKey: jwtKey}

	usersRouter := mux.PathPrefix("/users").Subrouter()
	usersRouter.Use(middleware.JwtVerify(jwtKey))

	logrus.SetFormatter(&logrus.JSONFormatter{})
	NewUserRouter(usersRouter, db)
	mux.HandleFunc("/register", r.register).Methods("POST")
	mux.HandleFunc("/login", r.login).Methods("POST")
	mux.HandleFunc("/fetch", r.fetchAPIs).Methods("GET")
	mux.HandleFunc("/fetchFile", r.fetchFromFile).Methods("GET")

}

func (r router) register(w http.ResponseWriter, req *http.Request) {
	var newUser entity.User
	err := json.NewDecoder(req.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	generatedHash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(generatedHash)
	err = r.db.Create(&newUser).Error
	if err != nil {
		logrus.Error("[USER-REGISTER] Error creating user:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logrus.Info("[USER-REGISTER] User created with id: ", newUser.ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func (r router) login(w http.ResponseWriter, req *http.Request) {
	var reqUser entity.User
	err := json.NewDecoder(req.Body).Decode(&reqUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var dbUser entity.User
	err = r.db.Where("email = ?", reqUser.Email).First(&dbUser).Error
	if err != nil {
		//logrus.Error("[USER-LOGIN] Error logging user: ", err)
		logrus.WithFields(logrus.Fields{
			"email": reqUser.Email,
		}).Error("[USER-LOGIN] Error logging user: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !utils.ComparePasswords(reqUser.Password, dbUser.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	token, err := r.generateTokens(&dbUser)
	if err != nil {
		logrus.Error("[USER-LOGIN] Error creating token: ", err)
		http.Error(w, "Error creating token", http.StatusUnauthorized)
		return
	}
	logrus.Info("[USER-LOGIN] User logged in")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Login successful", "token": token})
}

func (r router) generateTokens(user *entity.User) (string, error) {
	tokenClaims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Duration(60) * time.Minute).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	accessTokenString, err := accessToken.SignedString([]byte(r.jwtKey))
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

func (r router) fetchAPIs(w http.ResponseWriter, req *http.Request) {
	urls := []string{
		"https://api.coincap.io/v2/assets/bitcoin",
		"https://api.coincap.io/v2/assets/ethereum",
		"https://api.coincap.io/v2/assets/tether",
	}
	apiResutls := make([]utils.ApiResponse, 0)
	results := make(chan utils.ApiResponse)

	for _, url := range urls {
		go utils.FetchData(url, results)
	}

	for i := 0; i < len(urls); i++ {
		select {
		case result := <-results:
			apiResutls = append(apiResutls, result)
		case <-time.After(time.Second * 3):
			fmt.Println("Timeout reached, no more data received")
			close(results)
			break
		}
	}
	logrus.Info("[FETCH-API] APIs are fetched")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiResutls)
}

func (r router) fetchFromFile(w http.ResponseWriter, req *http.Request) {
	urlData, err := os.ReadFile("urls.txt")
	if err != nil {
		logrus.Error("[FETCH-FROM-FILE] Error fetching from file: ", err)
		return
	}
	urls := strings.Split(string(urlData), "\n")

	data := utils.ReadUrls(urls)
	logrus.Info("[FETCH-API] APIs are fetched")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
