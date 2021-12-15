package model

import (
	"awesomeProject/services"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func New(id int64, user string, password string) User {
	return User{
		Id:       id,
		Username: user,
		Password: services.GetPasswordHash(password),
	}
}

func (s *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(s.Password), []byte(password))
	return err == nil
}
