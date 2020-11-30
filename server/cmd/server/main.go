package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	db "github.com/G-V-G/l3/server/db"
	generalStore "github.com/G-V-G/l3/server/generalStore"
	tools "github.com/G-V-G/l3/server/tools"
)

func NewDbConnection() (*sql.DB, error) {
	conn := &db.Connection{
		DbName:     "Forums",
		User:       "postgres",
		Password:   "password",
		Host:       "localhost",
		DisableSSL: true,
	}
	return conn.Open()
}

func main() {
	db, err := NewDbConnection()
	if err != nil {
		log.Fatal(err)
	}
	store := generalStore.NewGeneralStore(db)
	http.HandleFunc("/forums", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			res, err := store.ListForums()
			if err != nil {
				log.Printf("Error making query to the db: %s", err)
				tools.WriteJsonInternalError(rw)
				return
			}
			tools.WriteJsonOk(rw, res)
		} else if r.Method == "POST" {
			var f generalStore.Forum
			if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
				log.Printf("Error decoding channel input: %s", err)
				tools.WriteJsonBadRequest(rw, "bad JSON payload")
				return
			}
			err := store.CreateForum(f.Name, f.Topic)
			if err == nil {
				tools.WriteJsonOk(rw, &f)
			} else {
				log.Printf("Error inserting record: %s", err)
				tools.WriteJsonInternalError(rw)
			}
		}
	})

	http.HandleFunc("/users", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			res, err := store.ListUsers()
			if err != nil {
				log.Printf("Error making query to the db: %s", err)
				tools.WriteJsonInternalError(rw)
				return
			}
			tools.WriteJsonOk(rw, res)
		} else if r.Method == "POST" {
			var fu tools.User
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Printf("Error decoding channel input: %s", err)
				tools.WriteJsonBadRequest(rw, "bad JSON payload")
				return
			}
			if err := json.Unmarshal(body, &fu); err != nil {
				log.Fatal(err)
			}
			err = store.CreateUser(fu.Username, fu.Interests)
			if err == nil {
				tools.WriteJsonOk(rw, &fu)
			} else {
				log.Printf("Error inserting record: %s", err)
				tools.WriteJsonInternalError(rw)
			}
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
