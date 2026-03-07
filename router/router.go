package router

import (
	"net/http"
	"user-api/handlers"
	"user-api/middleware"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", middleware.CORSMiddleware(handlers.Register))
	mux.HandleFunc("/login", middleware.CORSMiddleware(handlers.Login))

	mux.HandleFunc("/users", middleware.CORSMiddleware(middleware.AuthMiddleware(handlers.GetAllUsers)))
	return mux
}

// JWT (JSON Web Token)
// reswtjheryjetyrtyu865srgjhdetj.hsrtjuertyjtyjrty745675.abc123
