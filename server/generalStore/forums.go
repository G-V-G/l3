package generalStore

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Forum struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Topic   string `json:"topic"`
	UsersId []int  `json:"usersId"`
}

type FullForum struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Topic string  `json:"topic"`
	Users []*User `json:"users"`
}

type ForumStore struct {
	Db *sql.DB
}

func NewForumStore(db *sql.DB) *ForumStore {
	return &ForumStore{Db: db}
}

func (s *ForumStore) ListForums() ([]*FullForum, error) {
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

	var fullForums []*FullForum
	if res == nil {
		fullForums = make([]*FullForum, 0)
	} else {
		for i := 0; i < len(res); i++ {
			users := s.GetForumUsersByID(res[i].Id)
			fullForum := FullForum{
				Id:    res[i].Id,
				Name:  res[i].Name,
				Topic: res[i].Topic,
				Users: users}
			fullForums = append(fullForums, &fullForum)
		}
	}
	return fullForums, nil
}

func (s *ForumStore) FindForumByName(name string) []*FullForum {
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
	var fullForums []*FullForum
	if res == nil {
		fullForums = make([]*FullForum, 0)
	} else {
		for i := 0; i < len(res); i++ {
			users := s.GetForumUsersByID(res[i].Id)
			fullForum := FullForum{
				Id:    res[i].Id,
				Name:  res[i].Name,
				Topic: res[i].Topic,
				Users: users}
			fullForums = append(fullForums, &fullForum)
		}
	}
	return fullForums
}

func (s *ForumStore) FindForumByTopic(name string) ([]*Forum, error) {
	if len(name) < 0 {
		log.Fatal("Topic name is not provided")
	}
	rows, err := s.Db.Query(`SELECT * FROM forums where "topicKeyword" = $1`, name)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	var res []*Forum
	for rows.Next() {
		var f Forum
		if err := rows.Scan(&f.Id, &f.Name, &f.Topic); err != nil {
			log.Println(err)
		}
		res = append(res, &f)
	}
	if res == nil {
		res = make([]*Forum, 0)
		return res, errors.New("no such forum")
	}
	return res, nil
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

func (s *ForumStore) GetForumUsersByID(id int) []*User {
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

	var res []*User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Id, &u.Username); err != nil {
			log.Println(err)
		}
		res = append(res, &u)
	}
	if res == nil {
		res = make([]*User, 0)
	}

	return res
}

func (s *ForumStore) AddUserToForum(idForum, idUser int) error {
	_, err := s.Db.Exec(`INSERT INTO "usersList" ("forumsID", "userID") VALUES ($1, $2)`, idForum, idUser)
	return err
}
