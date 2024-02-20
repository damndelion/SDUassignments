package applicator

import (
	"fmt"
	"github.com/damndelion/SDUassignments/virtualization/assignment1/db"
	"github.com/damndelion/SDUassignments/virtualization/assignment1/internal/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func Run() {
	conn, err := db.Connect()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("connected")
	router := mux.NewRouter()
	controller.NewRouter(router, conn)

	http.ListenAndServe(":8080", router)
}
