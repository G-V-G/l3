package generalStore

import (
	"database/sql"
	"fmt"
	"github.com/G-V-G/l3/server/tools"
	_ "github.com/lib/pq"
)

type ForumStore struct {
	Db *sql.DB
}

func NewForumStore(db *sql.DB) *ForumStore {
	return &ForumStore{Db: db}
}

func (s *ForumStore) ListForums() (*tools.Forums, error) {
	rows, err := s.Db.Query("SELECT * FROM forums")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*tools.Forum
	for rows.Next() {
		var f tools.Forum
		if err := rows.Scan(&f.Id, &f.Name, &f.Topic); err != nil {
			return nil, err
		}
		res = append(res, &f)
	}

	var fullForums []*tools.Forum
	if res == nil {
		fullForums = make([]*tools.Forum, 0)
	} else {
		for i := 0; i < len(res); i++ {
			users, err := s.GetForumUsersByID(res[i].Id)
			if err != nil {
				return nil, err
			}
			fullForum := tools.Forum{
				Id:    res[i].Id,
				Name:  res[i].Name,
				Topic: res[i].Topic,
				Users: users}
			fullForums = append(fullForums, &fullForum)
		}
	}

	result := &tools.Forums{fullForums}
	return result, err
}

func (s *ForumStore) FindForumByName(name string) (*tools.Forums, error) {
	var textError string
	var err error
	var fullForums []*tools.Forum

	if len(name) == 0 {
		textError = "Forum name is not provided"
		err = fmt.Errorf(textError)
		fullForums = make([]*tools.Forum, 0)
		return nil, err
	}
	rows, err := s.Db.Query(`SELECT * FROM forums where name = $1`, name)
	if err != nil {
		textError = "There is no such forum"
		err = fmt.Errorf(textError)
		return nil, err
	}

	defer rows.Close()

	var res []*tools.Forum
	for rows.Next() {
		var f tools.Forum
		if err = rows.Scan(&f.Id, &f.Name, &f.Topic); err != nil {
			return nil, err
		}
		res = append(res, &f)
	}
	if res == nil {
		textError = "No such forum"
		err = fmt.Errorf(textError)
		return nil, err
	}
	for i := 0; i < len(res); i++ {
		users, err := s.GetForumUsersByID(res[i].Id)
		if err != nil {
			return nil, err
		}
		fullForum := tools.Forum{
			Id:    res[i].Id,
			Name:  res[i].Name,
			Topic: res[i].Topic,
			Users: users}
		fullForums = append(fullForums, &fullForum)
	}

	result := &tools.Forums{fullForums}
	return result, nil
}

func (s *ForumStore) FindForumByTopic(name string) ([]*tools.Forum, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("Topic name is not provided")
	}
	rows, err := s.Db.Query(`SELECT * FROM forums where topicKeyword = $1`, name)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*tools.Forum
	for rows.Next() {
		var f tools.Forum
		if err := rows.Scan(&f.Id, &f.Name, &f.Topic); err != nil {
			return nil, err
		}
		res = append(res, &f)
	}
	if res == nil {
		res = make([]*tools.Forum, 0)
		return res, fmt.Errorf("no such forum")
	}
	return res, nil
}

func (s *ForumStore) CreateForum(name, topicKeyword string) error {
	if len(name) == 0 {
		return fmt.Errorf("Forum name is not provided")
	}
	if len(topicKeyword) == 0 {
		return fmt.Errorf("Topic keyword name is not provided")
	}
	_, err := s.Db.Exec(`INSERT INTO forums (name, topicKeyword) VALUES ($1, $2)`, name, topicKeyword)
	if err != nil {
		return fmt.Errorf("Forum with this name or topic already exists")
	}
	forums, err := s.FindForumByName(name)
	_, err = s.Db.Exec(`INSERT INTO usersList (forumsID) VALUES ($1)`, forums.ForumsArr[0].Id)
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
