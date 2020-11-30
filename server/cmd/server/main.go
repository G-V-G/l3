package main

import (
	"database/sql"
	"log"
	"net/http"

	db "github.com/G-V-G/l3/server/db"
	generalStore "github.com/G-V-G/l3/server/generalStore"

	// forums "github.com/G-V-G/l3/server/forums"
	tools "github.com/G-V-G/l3/server/tools"
	// users "github.com/G-V-G/l3/server/users"
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
	http.HandleFunc("/forums", func(rw http.ResponseWriter, r *http.Request) {
		store := generalStore.NewGeneralStore(db)
		if r.Method == "GET" {
			res, err := store.ListForums()
			if err != nil {
				log.Printf("Error making query to the db: %s", err)
				tools.WriteJsonInternalError(rw)
				return
			}
			tools.WriteJsonOk(rw, res)
		}
		// else if r.Method == "POST" {
		// 		var f forums.Forum
		// 		if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		// 			log.Printf("Error decoding channel input: %s", err)
		// 			tools.WriteJsonBadRequest(rw, "bad JSON payload")
		// 			return
		// 		}
		// 		err := store.CreateForum(f.Name, f.Topic)
		// 		if err == nil {
		// 			tools.WriteJsonOk(rw, &f)
		// 		} else {
		// 			log.Printf("Error inserting record: %s", err)
		// 			tools.WriteJsonInternalError(rw)
		// 		}
		// 	}
	})

	// http.HandleFunc("/users", func(rw http.ResponseWriter, r *http.Request) {
	// 	store := users.NewUserStore(db)
	// 	if r.Method == "GET" {
	// 		res, err := store.ListUsers()
	// 		if err != nil {
	// 			log.Printf("Error making query to the db: %s", err)
	// 			tools.WriteJsonInternalError(rw)
	// 			return
	// 		}
	// 		tools.WriteJsonOk(rw, res)
	// 	} else if r.Method == "POST" {
	// 		var fu tools.FullUser
	// 		body, err := ioutil.ReadAll(r.Body)
	// 		if err != nil {
	// 			log.Printf("Error decoding channel input: %s", err)
	// 			tools.WriteJsonBadRequest(rw, "bad JSON payload")
	// 			return
	// 		}
	// 		if err := json.Unmarshal(body, &fu); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		err = store.CreateUser(fu.Username, fu.Interests)
	// 		if err == nil {
	// 			tools.WriteJsonOk(rw, &fu)
	// 		} else {
	// 			log.Printf("Error inserting record: %s", err)
	// 			tools.WriteJsonInternalError(rw)
	// 		}
	// 	}
	// })

	// http.HandleFunc("/GetUsersInterestByID", func(rw http.ResponseWriter, r *http.Request) {
	// 	store := users.NewUserStore(db)
	// 	if r.Method == "POST" {
	// 		var i []string
	// 		var id tools.ID
	// 		body, err := ioutil.ReadAll(r.Body)
	// 		if err != nil {
	// 			log.Printf("Error decoding channel input: %s", err)
	// 			tools.WriteJsonBadRequest(rw, "bad JSON payload")
	// 			return
	// 		}
	// 		if err := json.Unmarshal(body, &id); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		i = store.GetUsersInterestByID(id.Id)
	// 		if err == nil {
	// 			tools.WriteJsonOk(rw, &i)
	// 		} else {
	// 			log.Printf("Error inserting record: %s", err)
	// 			tools.WriteJsonInternalError(rw)
	// 		}
	// 	}
	// })

	// http.HandleFunc("/GetForumUsersByID", func(rw http.ResponseWriter, r *http.Request) {
	// 	store := forums.NewForumStore(db)
	// 	if r.Method == "POST" {
	// 		var i []*users.User
	// 		var id tools.ID
	// 		body, err := ioutil.ReadAll(r.Body)
	// 		if err != nil {
	// 			log.Printf("Error decoding channel input: %s", err)
	// 			tools.WriteJsonBadRequest(rw, "bad JSON payload")
	// 			return
	// 		}
	// 		if err := json.Unmarshal(body, &id); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		i = store.GetForumUsersByID(id.Id)
	// 		if err == nil {
	// 			tools.WriteJsonOk(rw, &i)
	// 		} else {
	// 			log.Printf("Error inserting record: %s", err)
	// 			tools.WriteJsonInternalError(rw)
	// 		}
	// 	}
	// })
	log.Fatal(http.ListenAndServe(":8080", nil))
}
