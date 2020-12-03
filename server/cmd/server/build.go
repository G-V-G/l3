//+build wireinject

package main

import (
	"github.com/google/wire"
	gs "github.com/G-V-G/l3/server/generalStore"
	h "github.com/G-V-G/l3/server/handlers"
)

var providers = wire.NewSet(gs.NewGeneralStore, h.NewHandler)

// NewServer generates main forum server
func NewServer(senv *ServerEnv) (*ForumServer, error) {
	wire.Build(
		NewDbConnection,
		providers,
		wire.Struct(new(ForumServer), "Handlers", "Senv"),
	)

	return nil, nil
}