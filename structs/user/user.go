package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDay  string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	// user     User
	User
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDay)
}

func (a *Admin) OutputAdminDetails() {
	fmt.Println(a.email, a.password, a.firstName, a.lastName, a.birthDay)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthday string) (*User, error) {
	if firstName == "" || lastName == "" || birthday == "" {
		errorMessage := "first name, last name and birthday are required"
		return nil, errors.New(errorMessage)
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDay:  birthday,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password, firstName, lastName, birthday string) (*Admin, error) {
	if email == "" || password == "" {
		errorMessage := "email and password are required"
		return nil, errors.New(errorMessage)
	}

	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: firstName,
			lastName:  lastName,
			birthDay:  birthday,
		},
	}, nil
}
