package main

import (
	"context"
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
func (fs *ForumServer) Run() error {
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
	fs.server = &http.Server{Addr: runnable}
	fmt.Printf("Server is running on port: %d, host: %s\n", fs.Senv.Port, fs.Senv.Host)
	return fs.server.ListenAndServe()
}

// Close forums server
func (fs *ForumServer) Close() error {
	if fs.server == nil {
		return fmt.Errorf("Server was not started")
	}
	return fs.server.Shutdown(context.Background())
}