package application

import (
	"v1-monolith/internal/user/application/dto"
	"v1-monolith/internal/user/domain"
)

type CreateUserUseCase struct {
	repo domain.UserRepository
}

func (uc *CreateUserUseCase) Execute(input dto.CreateUserInput) error {
	user, err := domain.NewUser(input.FullName, input.Email, domain.Role(input.Role))
	if err != nil {
		return err
	}

	return uc.repo.Create(user)
}
