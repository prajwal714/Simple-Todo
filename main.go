package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"main.go/common"
	"main.go/config"
	"main.go/contracts"
	"main.go/delivery"
	"main.go/repository"

	"github.com/gorilla/mux"
)

var TodoList []contracts.Item

func seed() {

	var items = []contracts.Item{
		{
			Name:        "Code",
			ID:          1,
			Description: "Revive Golang Skill",
		},
		{
			Name:        "Hobby",
			ID:          2,
			Description: "Find one good hobby",
		},
		{
			Name:        "Hobby",
			ID:          3,
			Description: "Find one good hobby",
		},
	}

	for _, item := range items {
		repository.Insert(item)
	}

}

func main() {

	common.Initialize()

	dbConn := config.InitDB()
	defer dbConn.Close()

	dbRedis := config.InitRedis()

	// seed()
	r := mux.NewRouter()
	Load(r, dbConn, dbRedis)
	r.HandleFunc("/listItems", delivery.ListItems).Methods("GET")
	r.HandleFunc("/addItem", delivery.AddItem).Methods("POST")
	r.HandleFunc("/deleteItem/{id:[0-9]+}", delivery.DeleteItem).Methods("POST")
	r.HandleFunc("/signupUser", delivery.SignupUser).Methods("POST")
	fmt.Println(TodoList)

	log.Fatal(http.ListenAndServe(":8080", r))
}
