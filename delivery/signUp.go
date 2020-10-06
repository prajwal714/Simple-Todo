package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"main.go/contracts"
	"main.go/repository"
)

func SignupUser(w http.ResponseWriter, r *http.Request) {
	var newUser contracts.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Error("Error Parsing User Info", err)
		fmt.Fprintf(w, "Error Parsing Request")
		return
	}

	log.Println(newUser)

	userID, err := repository.CreateNewUser(newUser)
	if err != nil {
		log.Error(err)
		fmt.Fprintf(w, err.Error())
		return
	}

	log.Info("User successfully created")
	fmt.Fprintf(w, "Successfully created new user", userID)
	return

}
