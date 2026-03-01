package repository

import (
	"user-api/db"
	"user-api/models"
)

// func CreateUser(u models.User) (int64, error) {
// 	res, err := db.DB.Exec(`INSERT INTO users (name, age)
// 	VALUES (?,?)`, u.Name, u.Age)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return res.LastInsertId()
// }

func RegisterUser(u models.User) (int64, error) {
	res, err := db.DB.Exec(`INSERT INTO users (name, age, login, password)
	VALUES (?,?,?,?)`, u.Name, u.Age, u.Login, u.Password)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetUserByLogin(login string) (models.User, error) {
	var u models.User
	err := db.DB.QueryRow(
		"SELECT id, name, age, login, password FROM users WHERE login = ?", login,
	).Scan(&u.ID, &u.Name, &u.Age, &u.Login, &u.Password)
	return u, err
}