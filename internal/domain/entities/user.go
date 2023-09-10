package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashedPassword"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func NewUser(firstName, lastName, email, hashedPassword string) *User {
	return &User{
		ID:             uuid.New(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		FirstName:      firstName,
		LastName:       lastName,
		Email:          email,
		HashedPassword: hashedPassword,
	}
}

func (u User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u User) GetEmail() string {
	return u.Email
}
