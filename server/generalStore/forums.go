package generalStore

import (
	"database/sql"
	"errors"
	"fmt"
)

type Forum struct {
	Id      int    `json:"-"`
	Name    string `json:"name"`
	Topic   string `json:"topic"`
	Users []string `json:"users"`
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

	var fullForums []*Forum
	if res == nil {
		fullForums = make([]*Forum, 0)
	} else {
		for i := 0; i < len(res); i++ {
			users, err := s.GetForumUsersByID(res[i].Id)
			if err != nil {
				return nil, err
			}
			fullForum := Forum{
				Id:    res[i].Id,
				Name:  res[i].Name,
				Topic: res[i].Topic,
				Users: users}
			fullForums = append(fullForums, &fullForum)
		}
	}
	return fullForums, err
}

func (s *ForumStore) FindForumByName(name string) ([]*Forum, error) {
	var textError string
	var err error
	var fullForums []*Forum

	if len(name) < 0 {
		textError = "Forum name is not provided"
		err = errors.New(textError)
		fullForums = make([]*Forum, 0)
		return nil, err
	}
	rows, err := s.Db.Query(`SELECT * FROM forums where name = $1`, name)
	if err != nil {
		textError = "There is no such forum"
		err = errors.New(textError)
		return nil, err
	}

	defer rows.Close()

	var res []*Forum
	for rows.Next() {
		var f Forum
		if err = rows.Scan(&f.Id, &f.Name, &f.Topic); err != nil {
			return nil, err
		}
		res = append(res, &f)
	}
	if res == nil {
		textError = "No such forum"
		err = errors.New(textError)
		return nil, err
	}
	for i := 0; i < len(res); i++ {
		users, err := s.GetForumUsersByID(res[i].Id)
		if err != nil {
			return nil, err
		}
		fullForum := Forum{
			Id:    res[i].Id,
			Name:  res[i].Name,
			Topic: res[i].Topic,
			Users: users}
		fullForums = append(fullForums, &fullForum)
	}
	return fullForums, nil
}

func (s *ForumStore) FindForumByTopic(name string) ([]*Forum, error) {
	if len(name) < 0 {
		return nil, fmt.Errorf("Topic name is not provided")
	}
	rows, err := s.Db.Query(`SELECT * FROM forums where topicKeyword = $1`, name)
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
		return res, errors.New("no such forum")
	}
	return res, nil
}

func (s *ForumStore) CreateForum(name, topicKeyword string) error {
	if len(name) < 0 {
		return fmt.Errorf("Forum name is not provided")
	}
	_, err := s.Db.Exec(`INSERT INTO forums (name, topicKeyword) VALUES ($1, $2)`, name, topicKeyword)
	forum, err := s.FindForumByName(name)
	_, err = s.Db.Exec(`INSERT INTO usersList (forumsID) VALUES ($1)`, forum[0].Id)
	return err
}

func (s *ForumStore) GetForumUsersByID(id int) ([]string, error) {
	if id < 1 {
		return nil, fmt.Errorf("ID is incorrect")
	}
	rows, err := s.Db.Query(`
	select
		users.name
	from
		forums
	left join
		usersList
	on
		usersList.forumsID = forums.id
	left join
		users
	on
		users.id = usersList.userID
	where
		forums.id = $1
	GROUP BY
		users.id
	HAVING users.id is not NULL
	`,
		id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []string
	for rows.Next() {
		var u string
		if err := rows.Scan(&u); err != nil {
			return nil, err
		}
		if u != "" {
			res = append(res, u)
		}
	}
	if res == nil {
		res = make([]string, 0)
	}

	return res, nil
}

func (s *ForumStore) AddUserToForum(idForum, idUser int) error {
	_, err := s.Db.Exec(`INSERT INTO usersList (forumsID, userID) VALUES ($1, $2)`, idForum, idUser)
	return err
}
