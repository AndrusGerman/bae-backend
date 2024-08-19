package service

import (
	"bae-backend/internal/core/domain"
	"bae-backend/internal/core/port"
)

type UserService struct {
	userRepo port.UserRepository
}

func NewUserService(userRepo port.UserRepository) port.UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) GetUser(id domain.Id) (*domain.User, error) {
	return us.userRepo.GetByID(id)
}

func (us *UserService) Register(user *domain.User) (*domain.User, error) {
	// validate data //ensureXIsValid
	return us.userRepo.Create(user)
}

func (us *UserService) GetAllUsers() ([]domain.User, error) {
	// validate data
	return us.userRepo.GetAll()
}
