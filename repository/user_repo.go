package repository

import (
	"user-api/db"
	"user-api/models"
)

func CreateUser(u models.User) (int64, error) {
	res, err := db.DB.Exec(`INSERT INTO users (name, age)
	VALUES (?,?)`, u.Name, u.Age)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
