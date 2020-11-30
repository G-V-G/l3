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

func (gs *GeneralStore) ListForums() ([]*Forum, error) {
	return gs.FStore.ListForums()
}

func (gs *GeneralStore) FindForumByName(name string) []*Forum {
	return gs.FStore.FindForumByName(name)
}

func (gs *GeneralStore) CreateForum(name, topicKeyword string) error {
	return gs.FStore.CreateForum(name, topicKeyword)
}
