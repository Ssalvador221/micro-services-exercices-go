package repository

import (
	"context"
	"fmt"
	"v1-monolith/internal/infraestructure/db/sqlc"
	"v1-monolith/internal/user/domain"

	"github.com/google/uuid"
)

type UserRepository struct {
	ctx     context.Context
	queries *sqlc.Queries
}

func NewUserRepository(ctx context.Context, q *sqlc.Queries) *UserRepository {
	return &UserRepository{
		ctx:     ctx,
		queries: q,
	}
}

// Create a new user.
func (ur *UserRepository) Create(ctx context.Context, user *domain.User) error {
	err := ur.queries.CreateUser(ctx, sqlc.CreateUserParams{
		ID:       user.ID.String(),
		FullName: user.FullName,
		Email:    user.Email,
		Role:     string(user.Role),
	})

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

// List all users avaibles in users table.
func (ur *UserRepository) FindMany(ctx context.Context) ([]domain.User, error) {
	users, err := ur.queries.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	response := make([]domain.User, 0, len(users))
	for _, u := range users {
		response = append(response, domain.User{
			ID:       uuid.MustParse(u.ID),
			FullName: u.FullName,
			Email:    u.Email,
			Role:     domain.Role(u.Role),
		})
	}

	return response, nil
}
