package router

import (
	"net/http"
	"user-api/handlers"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", handlers.CreateUser)
	return mux
}

// JWT (JSON Web Token)
// reswtjheryjetyrtyu865srgjhdetj.hsrtjuertyjtyjrty745675.abc123
