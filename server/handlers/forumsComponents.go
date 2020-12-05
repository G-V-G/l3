package handlers

import (
	"encoding/json"
	"net/http"
	
	"github.com/G-V-G/l3/server/tools"
	gs "github.com/G-V-G/l3/server/generalStore"
)

func addForum(db *gs.GeneralStore, rw http.ResponseWriter, req *http.Request) {
	var forum tools.Forum
	if err := json.NewDecoder(req.Body).Decode(&forum); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
		return
	}
	if err := db.CreateForum(forum.Name, forum.Topic); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
	} else {
		rw.WriteHeader(http.StatusCreated)
	}
}

func getForums(db *gs.GeneralStore, rw http.ResponseWriter, req *http.Request) {
	if res, err := db.ListForums(); err != nil {
		tools.WriteJsonInternalError(rw, err.Error())
	} else {
		tools.WriteJsonOk(rw, res)
	}
}
