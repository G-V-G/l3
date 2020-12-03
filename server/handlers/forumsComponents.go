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
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if err := db.CreateForum(forum.Name, forum.Topic); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	} else {
		rw.WriteHeader(http.StatusCreated)
	}
}

func getForums(db *gs.GeneralStore, rw http.ResponseWriter, req *http.Request) {
	res, err := db.ListForums()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	tools.WriteJsonOk(rw, res)
}
