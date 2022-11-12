package user

import (
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID
	Name string
	Data string
}
type UserStore interface {
	Create(u User) (*uuid.UUID, error)
}
type Users struct {
	ustore UserStore
}

func (us *Users) Create(u User) (*User, error) {
	id, err := us.ustore.Create(u)
	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}
	u.ID = *id
	return &u, nil
}
