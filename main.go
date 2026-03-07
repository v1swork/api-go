package main

import (
	"log"
	"net/http"
	"user-api/db"
	"user-api/router"
)

func main() {
	db.InitDB()

	mux := router.SetupRouter()
	fs := http.FileServer(http.Dir("/static"))
	mux.Handle("/", fs)

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
