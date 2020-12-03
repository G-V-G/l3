package handlers

import (
	gs "github.com/G-V-G/l3/server/generalStore"
)

// NewHandler returns server methods
func NewHandler(gs *gs.GeneralStore) *Handlers {
	return &Handlers{gs}
}
