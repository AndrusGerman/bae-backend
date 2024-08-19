package port

import (
	"bae-backend/internal/core/domain"
)

//go:generate mockgen -source=user.go -destination=mock/user.go -package=mock

type UserRepository interface {
	// CreateUser inserts a new user into the database
	Create(user *domain.User) (*domain.User, error)
	// GetUserByID selects a user by id
	GetByID(id domain.Id) (*domain.User, error)
	// GetUserByEmail selects a user by email
	//GetUserByEmail(email string) (*domain.User, error)
	// ListUsers selects a list of users with pagination
	//ListUsers(skip, limit uint64) ([]domain.User, error)
	// UpdateUser updates a user
	//UpdateUser(user *domain.User) (*domain.User, error)
	// DeleteUser deletes a user
	//DeleteUser(id uint64) error

	// GetAll get all users
	GetAll() ([]domain.User, error)
}

// UserService is an interface for interacting with user-related business logic
type UserService interface {
	// Register registers a new user
	Register(user *domain.User) (*domain.User, error)
	// GetUser returns a user by id
	GetUser(id domain.Id) (*domain.User, error)
	// UpdateUser updates a user
	//UpdateUser(user *domain.User) (*domain.User, error)
	// DeleteUser deletes a user
	//DeleteUser(id uint64) error

	// GetAll get all users
	GetAllUsers() ([]domain.User, error)
}
