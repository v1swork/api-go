package auth

import (
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("Wreza3qoSp_uN8Py_4At")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
}
