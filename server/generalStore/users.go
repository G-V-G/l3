package generalStore

import (
	"database/sql"
	"fmt"
	"github.com/G-V-G/l3/server/tools"
	_ "github.com/lib/pq"
)

type UserStore struct {
	Db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{Db: db}
}

func (s *UserStore) ListUsers() (*tools.Users, error) {
	rows, err := s.Db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*tools.User
	for rows.Next() {
		var u tools.User
		if err := rows.Scan(&u.Id, &u.Name); err != nil {
			return nil, err
		}
		res = append(res, &u)
	}

	var fullUsers []*tools.User
	if res == nil {
		fullUsers = make([]*tools.User, 0)
	} else {
		for i := 0; i < len(res); i++ {
			interests, err := s.GetUsersInterestByID(res[i].Id)
			if err != nil {
				return nil, err
			}
			fullUser := tools.User{Id: res[i].Id, Name: res[i].Name, Interests: interests}
			fullUsers = append(fullUsers, &fullUser)
		}
	}

	result := &tools.Users{fullUsers}
	return result, nil
}

func (s *UserStore) FindUserByName(name string) (*tools.Users, error) {
	var textError string
	var err error
	var fullUsers []*tools.User

	if len(name) == 0 {
		textError = "User name is not provided"
		err = fmt.Errorf(textError)
		return nil, err
	}
	rows, err := s.Db.Query(`SELECT * FROM users where name = $1`, name)
	if err != nil {
		textError = "There is no such user"
		err = fmt.Errorf(textError)
		return nil, err
	}

	defer rows.Close()

	var res []*tools.User
	for rows.Next() {
		var u tools.User
		if err := rows.Scan(&u.Id, &u.Name); err != nil {
			return nil, err
		}
		res = append(res, &u)
	}

	if res == nil {
		textError = "No such user"
		err = fmt.Errorf(textError)
		fullUsers = make([]*tools.User, 0)
		return nil, err
	} 
	for i := 0; i < len(res); i++ {
		interests, err := s.GetUsersInterestByID(res[i].Id)
		if err != nil {
			return nil, err
		}
		fullUser := tools.User{Id: res[i].Id, Name: res[i].Name, Interests: interests}
		fullUsers = append(fullUsers, &fullUser)
	}
	err = nil
	result := &tools.Users{fullUsers}
	return result, err
}

func (s *UserStore) CreateUser(username string, interests []string) error {
	store := NewForumStore(s.Db)
	if len(username) == 0 {
		return fmt.Errorf("Username is not provided")
	}
	if len(interests) == 0 {
		return fmt.Errorf("Interests are not provided")
	}
	for _, interest := range interests {
		if len(interest) == 0 {
			return fmt.Errorf("Interest cannot be empty")
		}
	}
	_, err := s.Db.Exec(`INSERT INTO users (name) VALUES ($1)`, username)
	if err != nil {
		return fmt.Errorf("User with this name already exists")
	}
	users, err := s.FindUserByName(username)
	for i := 0; i < len(interests); i++ {
		_, err = s.Db.Exec(`INSERT INTO interestList (interest, userID) VALUES ($1, $2)`,
			interests[i], users.UsersArr[0].Id)
		forum, indicate := store.FindForumByTopic(interests[i])
		if indicate == nil {
			err = store.AddUserToForum(forum[0].Id, users.UsersArr[0].Id)
		}
	}
	return err
}

func (s *UserStore) GetUsersInterestByID(id int) ([]string, error) {
	if id < 1 {
		return nil, fmt.Errorf("ID is incorrect")
	}
	rows, err := s.Db.Query(`
	select
		interestList.interest
	from
		users, interestList
	where
		interestList.userID = users.id
	and
		users.id = $1`,
		id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []string
	for rows.Next() {
		var i string
		if err := rows.Scan(&i); err != nil {
			return nil, err
		}
		res = append(res, i)
	}
	if res == nil {
		res = make([]string, 0)
	}

	return res, nil
}
