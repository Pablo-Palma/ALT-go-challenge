package auth

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID			primitive.ObjectID	`json:"id" json:"id"`
	Username	string				`json:"username"`
	Password	string				`json:"password"`
}

/*
	Recibe la password y devuelve un hash usando bcrypt.
*/
func	HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)	
	return string(bytes), err
}

/*
	Compara la contraseña con el hash.
*/
func	CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
