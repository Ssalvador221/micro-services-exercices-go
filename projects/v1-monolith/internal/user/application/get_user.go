package application

import (
	"v1-monolith/internal/user/domain"
)

type GetUserUseCase struct {
	repo domain.UserRepository
}

func (uc *GetUserUseCase) Execute() ([]domain.User, error) {
	return uc.repo.FindMany()
}
