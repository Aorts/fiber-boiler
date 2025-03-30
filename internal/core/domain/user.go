package domain

import "time"

type User struct {
	Id        string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(username, email, password string) *User {
	return &User{
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
