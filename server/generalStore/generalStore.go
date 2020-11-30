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
