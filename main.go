package main

import (
	"log"
	"net/http"

	"user-api/db"
	"user-api/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	db.InitDB()

	fs := http.FileServer(http.Dir("./static"))

	mux := router.SetupRouter()
	mux.Handle("/", fs)

	log.Println("Сервер запущен на localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
