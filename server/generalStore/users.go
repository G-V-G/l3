package generalStore

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type UserStore struct {
	Db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{Db: db}
}

func (s *UserStore) ListUsers() ([]*User, error) {
	rows, err := s.Db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Id, &u.Username); err != nil {
			return nil, err
		}
		res = append(res, &u)
	}
	if res == nil {
		res = make([]*User, 0)
	}
	return res, nil
}

func (s *UserStore) FindUserByName(name string) []User {
	if len(name) < 0 {
		log.Fatal("Forum name is not provided")
	}
	rows, err := s.Db.Query(`SELECT * FROM users where "name" = $1`, name)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var res []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Id, &u.Username); err != nil {
			log.Fatal(err)
		}
		res = append(res, u)
	}
	if res == nil {
		res = make([]User, 0)
	}

	return res
}

func (s *UserStore) CreateUser(username string, interests []string) error {
	if len(username) < 0 {
		return fmt.Errorf("Username is not provided")
	}
	_, err := s.Db.Exec(`INSERT INTO users (name) VALUES ($1)`, username)
	user := s.FindUserByName(username)
	log.Println(interests)
	for i := 0; i < len(interests); i++ {
		_, err = s.Db.Exec(`INSERT INTO "interestList" ("interest", "userID") VALUES ($1, $2)`,
			interests[i], user[0].Id)
		// TODO: search through topics in forums to add users to lists
	}
	return err
}

func (s *UserStore) GetUsersInterestByID(id int) []string {
	if id < 1 {
		log.Fatal("ID is incorrect")
	}
	rows, err := s.Db.Query(`
	select
		"interestList"."interest"
	from
		users, "interestList"
	where
		"interestList"."userID" = users.id
	and
		users.id = $1`,
		id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var res []string
	for rows.Next() {
		var i string
		if err := rows.Scan(&i); err != nil {
			log.Fatal(err)
		}
		res = append(res, i)
	}
	if res == nil {
		res = make([]string, 0)
	}

	return res
}
