package main

import (
	"net/http"
	"fmt"
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
	runnable := fs.Senv.Host + ":" + fmt.Sprint(fs.Senv.Port)
	fmt.Printf("Server is running on port: %d, host: %s\n", fs.Senv.Port, fs.Senv.Host)
	http.ListenAndServe(runnable, nil)
}