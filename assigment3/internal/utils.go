package utils

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

type ApiResponse struct {
	Data interface{} `json:"data"`
}

func FetchData(url string, results chan ApiResponse) {
	resp, err := http.Get(url)
	if err != nil {
		logrus.Error("Error fetching data from", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("Error reading response body from", url, err)
		return
	}

	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		logrus.Error("Error unmarshaling JSON from", url, err)
		return
	}
	results <- apiResponse
}

func Worker(id int, jobs <-chan string, results chan<- string) {
	for url := range jobs {
		logrus.Println("worker", id, "started  job", url)
		resp, err := http.Get(url)
		if err != nil {
			logrus.Println("Error fetching data from", url, err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Println("Error reading response body from", url, err)
			return
		}
		logrus.Println("worker", id, "finished job", url)
		results <- string(body)
	}
}
func ReadUrls(urls []string) []string {
	jobs := make(chan string, len(urls))
	results := make(chan string, len(urls))

	for w := 1; w <= 3; w++ {
		go Worker(w, jobs, results)
	}
	for _, v := range urls {
		jobs <- v
	}
	close(jobs)
	dataFromUrls := make([]string, 0)
	for i := 1; i <= len(urls); i++ {
		dataFromUrls = append(dataFromUrls, <-results)
	}
	return dataFromUrls
}

func ComparePasswords(attemptedPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(attemptedPassword))
	return err == nil
}
