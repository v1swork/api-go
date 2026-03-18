package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "host=127.0.0.1 port=5433 user=postgres password=postgres dbname=userapi sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("БД не отвечает:", err)
	}

	log.Println("Подключение к PostgreSQL успешно!")

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS users
	(
		id SERIAL PRIMARY KEY,
		name TEXT,
		age INTEGER,
		login TEXT UNIQUE,
		password TEXT
	);`)

	if err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}

}
