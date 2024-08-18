package service

import (
	"bae-backend/internal/core/domain"
	"bae-backend/internal/core/port"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) port.UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) GetUser(id domain.Id) (*domain.User, error) {
	return us.repo.GetUserByID(id)
}

func (us *UserService) Register(user *domain.User) (*domain.User, error) {
	// validate data
	return us.repo.CreateUser(user)
}

func (us *UserService) GetAll() ([]domain.User, error) {
	// validate data
	return us.repo.GetAll()
}
