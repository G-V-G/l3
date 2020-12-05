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
		tools.WriteJsonBadRequest(rw, err.Error())
		return
	}
	if err := db.CreateUser(user.Name, user.Interests); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
	} else {
		rw.WriteHeader(http.StatusCreated)
	}
}

func getUsers(db *gs.GeneralStore, rw http.ResponseWriter, req *http.Request) {
	if res, err := db.ListUsers(); err != nil {
		tools.WriteJsonInternalError(rw, err.Error())
	} else {
		tools.WriteJsonOk(rw, res)
	}
}