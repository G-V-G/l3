package generalStore

import (
	"database/sql"
	"github.com/G-V-G/l3/server/tools"
	_ "github.com/lib/pq"
)

type GeneralStore struct {
	FStore *ForumStore
	UStore *UserStore
}

func NewGeneralStore(db *sql.DB) *GeneralStore {
	fstore := NewForumStore(db)
	ustore := NewUserStore(db)
	return &GeneralStore{FStore: fstore, UStore: ustore}
}

func (gs *GeneralStore) ListForums() ([]*tools.Forum, error) {
	return gs.FStore.ListForums()
}

func (gs *GeneralStore) FindForumByName(name string) ([]*tools.Forum, error) {
	return gs.FStore.FindForumByName(name)
}

func (gs *GeneralStore) FindUserByName(name string) ([]*tools.User, error) {
	return gs.UStore.FindUserByName(name)
}

func (gs *GeneralStore) CreateForum(name, topicKeyword string) error {
	return gs.FStore.CreateForum(name, topicKeyword)
}

func (gs *GeneralStore) ListUsers() ([]*tools.User, error) {
	return gs.UStore.ListUsers()
}

func (gs *GeneralStore) CreateUser(username string, interests []string) error {
	return gs.UStore.CreateUser(username, interests)
}
