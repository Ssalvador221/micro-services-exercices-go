package dto

import (
	"v1-monolith/internal/user/domain"
	"time"
)

type CreateUserInput struct {
	FullName  string
	Email     string
	Role      domain.Role
	CreatedAt time.Time
}
