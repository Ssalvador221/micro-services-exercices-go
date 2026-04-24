package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	ID        uuid.UUID
	FullName  string
	Email     string
	Role      Role
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func NewUser(fullName, email string, role Role) (*User, error) {
	if !IsValidRole(role) {
		return nil, errors.New("Please insert a correct role!")
	}

	now := time.Now()

	return &User{
		ID:        uuid.New(),
		FullName:  fullName,
		Email:     email,
		Role:      role,
		CreatedAt: now,
	}, nil
}

func (u User) GetUser() *User {
	return &User{
		ID:        u.ID,
		FullName:  u.FullName,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}

func IsValidRole(r Role) bool {
	switch r {
	case RoleAdmin, RoleUser:
		return true
	default:
		return false
	}
}
