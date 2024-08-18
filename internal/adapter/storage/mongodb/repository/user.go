package repository

import (
	"bae-backend/internal/adapter/storage/mongodb"
	"bae-backend/internal/core/domain"
	"bae-backend/internal/core/port"

	"go.mongodb.org/mongo-driver/bson"
)

/**
 * UserRepository implements port.UserRepository interface
 * and provides an access to the postgres database
 */
type UserRepository struct {
	collection *mongodb.Collection
}

// NewUserRepository creates a new user repository instance
func NewUserRepository(db *mongodb.DB) port.UserRepository {
	return &UserRepository{
		db.NewCollection("users"),
	}
}

// CreateUser creates a new user in the database
func (ur *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	var err = ur.collection.InsertOne(user)
	return user, err
}

// GetUserByID gets a user by ID from the database
func (ur *UserRepository) GetUserByID(id domain.Id) (*domain.User, error) {
	var user = new(domain.User)
	var err error
	if err = ur.collection.FindOneById(id, user); err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByEmailAndPassword gets a user by email from the database
func (ur *UserRepository) GetUserByEmail(email string) (*domain.User, error) {

	return nil, nil
}

// UpdateUser updates a user by ID in the database
func (ur *UserRepository) UpdateUser(user *domain.User) (*domain.User, error) {

	return nil, nil
}

// DeleteUser deletes a user by ID from the database
func (ur *UserRepository) DeleteUser(id domain.Id) error {

	return nil
}

func (ur *UserRepository) GetAll() ([]domain.User, error) {
	var users = new([]domain.User)
	var err = ur.collection.FindMany(bson.M{}, users)
	if err != nil {
		return nil, err
	}
	return *users, nil
}
