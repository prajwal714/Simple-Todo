package main

import (
	"encoding/json"
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

	// seed()
	r := mux.NewRouter()

	r.HandleFunc("/listItems", delivery.ListItems).Methods("GET")
	r.HandleFunc("/addItem", delivery.AddItem).Methods("POST")
	r.HandleFunc("/deleteItem/{id:[0-9]+}", delivery.DeleteItem).Methods("POST")
	// DeleteItem(2)
	fmt.Println(TodoList)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func GetItems(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(TodoList)
	return

}

// func AddItem(w http.ResponseWriter, r *http.Request) {

// 	var newItem contracts.Item

// 	err := json.NewDecoder(r.Body).Decode(&newItem)
// 	if err != nil {
// 		log.Error("Error Parsing Request", err)
// 		fmt.Fprintf(w, "Error Parsing Request")
// 	}

// 	if newItem.ID <= 0 {
// 		log.Error("Item ID Missing")
// 		fmt.Fprintf(w, "Item ID Missing")
// 		return
// 	}

// 	for _, item := range TodoList {
// 		if item.ID == newItem.ID {
// 			log.Warn("Item with ID already exists !!")
// 			fmt.Fprintf(w, "Item with ID Already exists")
// 			return
// 		}
// 	}

// 	TodoList = append(TodoList, newItem)
// 	log.Info("Item Successfully Added")
// 	fmt.Fprintf(w, "Successfully Added Item")
// 	return
// }

func DeleteItem(w http.ResponseWriter, r *http.Request) {

}
