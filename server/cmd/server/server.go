package main

import (
	"net/http"
	h "github.com/G-V-G/l3/server/handlers"
)

// ForumServer runs handlers
type ForumServer struct {
	server *http.Server
	Senv *ServerEnv
	Handlers *h.Handlers
}

// Run forums server
func (fs *ForumServer) Run() {
	handlersCollection := map[string] http.HandlerFunc {
		"/users": fs.Handlers.HandleUsers,
		"/forums": fs.Handlers.HandleForums,
		"/user": fs.Handlers.GetUser,
		"/forum": fs.Handlers.GetForum,
	}
	for route, handler := range handlersCollection {
		http.Handle(route, handler)
	}
	runnable := fs.Senv.Host + ":" + string(fs.Senv.Port)
	http.ListenAndServe(runnable, nil)
}