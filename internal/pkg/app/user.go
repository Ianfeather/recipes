package app

import (
	"encoding/json"
	"log"
	"net/http"
	"recipes/internal/pkg/common"
	"recipes/internal/pkg/service"
)

func (a *App) addUser(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)
	user := common.User{}
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, "Error decoding json body", http.StatusBadRequest)
		return
	}

	user.ID = req.Context().Value(contextKey("userID")).(string)

	err := service.AddUser(a.db, user)
	if err != nil {
		log.Println("Error: could not add new user")
		http.Error(w, "could not add new user", http.StatusInternalServerError)
		return
	}
	err = service.CreateAccount(a.db, user)
	if err != nil {
		log.Println("Error creating account for user")
		http.Error(w, "Error creating account for user", http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(user)
	if err != nil {
		http.Error(w, "Error encoding json", http.StatusInternalServerError)
	}
}
