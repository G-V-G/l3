package forums

import (
	"database/sql"
	"fmt"
	"log"

	users "github.com/G-V-G/l3/server/users"
)

type Forum struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Topic   string `json:"topic"`
	UsersId []int  `json:"usersId"`
}

type ForumStore struct {
	Db *sql.DB
}

func NewForumStore(db *sql.DB) *ForumStore {
	return &ForumStore{Db: db}
}

func (s *ForumStore) ListForums() ([]*Forum, error) {
	rows, err := s.Db.Query("SELECT * FROM forums")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Forum
	for rows.Next() {
		var f Forum
		if err := rows.Scan(&f.Id, &f.Name, &f.Topic); err != nil {
			return nil, err
		}
		res = append(res, &f)
	}
	if res == nil {
		res = make([]*Forum, 0)
	}
	return res, nil
}

func (s *ForumStore) FindForumByName(name string) []*Forum {
	if len(name) < 0 {
		log.Fatal("Forum name is not provided")
	}
	rows, err := s.Db.Query(`SELECT * FROM forums where "name" = $1`, name)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var res []*Forum
	for rows.Next() {
		var f Forum
		if err := rows.Scan(&f.Id, &f.Name, &f.Topic); err != nil {
			log.Fatal(err)
		}
		res = append(res, &f)
	}
	if res == nil {
		res = make([]*Forum, 0)
	}

	return res
}

func (s *ForumStore) CreateForum(name, topicKeyword string) error {
	if len(name) < 0 {
		return fmt.Errorf("Forum name is not provided")
	}
	_, err := s.Db.Exec(`INSERT INTO forums (name, "topicKeyword") VALUES ($1, $2)`, name, topicKeyword)
	forum := s.FindForumByName(name)
	_, err = s.Db.Exec(`INSERT INTO "usersList" ("forumsID") VALUES ($1)`, forum[0].Id)
	return err
}

func (s *ForumStore) GetForumUsersByID(id int) []*users.User {
	if id < 1 {
		log.Fatal("ID is incorrect")
	}
	rows, err := s.Db.Query(`
	select
		users.id, users.name
	from
		forums
	left join
		"usersList"
	on
		"usersList"."forumsID" = "forums".id
	left join
		users
	on
		users.id = "usersList"."userID"
	where
		forums.id = $1
	GROUP BY
		users.id`,
		id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var res []*users.User
	for rows.Next() {
		var u users.User
		if err := rows.Scan(&u.Id, &u.Username); err != nil {
			log.Fatal(err)
		}
		res = append(res, &u)
	}
	if res == nil {
		res = make([]*users.User, 0)
	}

	return res
}
