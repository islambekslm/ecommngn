package user

import "time"

type user struct {
	ID        string
	FirstName string
	LastName  string
	FullName  string
	BirthDate time.Time
	Email     string
	Password  string
	IsActive  bool
	IsValid   bool
}

func NewUser(id, firstName, lastName, email, password string, birthDate time.Time) *user {
	return &user{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		FullName:  firstName + " " + lastName,
		BirthDate: birthDate,
		Email:     email,
		Password:  password,
		IsActive:  false,
		IsValid:   false,
	}
}
