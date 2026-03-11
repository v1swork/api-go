package main

import (
	"log"
	"net/http"

	"user-api/db"
	"user-api/router"
)

func main() {
	db.InitDB()
	log.Println("Сервер запущен на localhost:8080")
	fs := http.FileServer(http.Dir("./static"))
	mux := router.SetupRouter()
	mux.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
