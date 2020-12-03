package handlers

import (
	"encoding/json"
	"net/http"
	
	"github.com/G-V-G/l3/server/tools"
	gs "github.com/G-V-G/l3/server/generalStore"
)

func addUser(db *gs.GeneralStore, rw http.ResponseWriter, req *http.Request) {
	var user tools.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if err := db.CreateUser(user.Name, user.Interests); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	} else {
		rw.WriteHeader(http.StatusCreated)
	}
}

func getUsers(db *gs.GeneralStore, rw http.ResponseWriter, req *http.Request) {
	res, err := db.ListUsers()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	tools.WriteJsonOk(rw, res)
}