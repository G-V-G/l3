package generalStore

import "database/sql"

type GeneralStore struct {
	FStore *ForumStore
	UStore *UserStore
}

func NewGeneralStore(db *sql.DB) *GeneralStore {
	fstore := NewForumStore(db)
	ustore := NewUserStore(db)
	return &GeneralStore{FStore: fstore, UStore: ustore}
}

func (gs *GeneralStore) ListForums() ([]*FullForum, error) {
	return gs.FStore.ListForums()
}

func (gs *GeneralStore) FindForumByName(name string) []*FullForum {
	return gs.FStore.FindForumByName(name)
}

func (gs *GeneralStore) FindUserByName(name string) []*FullUser {
	return gs.UStore.FindUserByName(name)
}

func (gs *GeneralStore) CreateForum(name, topicKeyword string) error {
	return gs.FStore.CreateForum(name, topicKeyword)
}

func (gs *GeneralStore) ListUsers() ([]*FullUser, error) {
	return gs.UStore.ListUsers()
}

func (gs *GeneralStore) CreateUser(username string, interests []string) error {
	return gs.UStore.CreateUser(username, interests)
}
