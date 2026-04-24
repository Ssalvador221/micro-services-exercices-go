package dto

import (
	"v1-monolith/internal/user/domain"
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID
	FullName  string
	Email     string
	Role      domain.Role
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
