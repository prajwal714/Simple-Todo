package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"main.go/contracts"
	"main.go/repository"
)

func AddItem(w http.ResponseWriter, r *http.Request) {

	var newItem contracts.Item

	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		log.Error("Error Parsing Request", err)
		fmt.Fprintf(w, "Error Parsing Request")
		return
	}

	log.Println(newItem)

	itemID, err := repository.Insert(newItem)
	if err != nil {
		log.Error(err)
		return
	}
	// TodoList = append(TodoList, newItem)

	log.Info("Item Successfully Added")
	fmt.Fprintf(w, "Successfully Added Item %s", string(itemID))
	return
}

func ListItems(w http.ResponseWriter, r *http.Request) {

	items, err := repository.List()
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("Successfully fetched list of items")
	log.Println(items)
	json.NewEncoder(w).Encode(items)
	return

}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	itemIDString := vars["id"]
	itemID, err := strconv.ParseInt(itemIDString, 10, 64)

	if err != nil {
		log.Error(err)
		fmt.Fprintf(w, err.Error())
	}

	if itemID <= 0 {
		log.Warn("Invalid Item ID !")
		fmt.Fprintf(w, "Invalid Item ID !")
		return
	}

	err = repository.Delete(itemID)
	if err != nil {
		log.Error(err)
		fmt.Fprintf(w, err.Error())
		return
	}

	log.Println("Item Successfully Deleted")
	fmt.Fprintf(w, "Successfully Deleted Item")

}
