package repository

import (
	"bae-backend/internal/adapter/storage/mongodb"
	"bae-backend/internal/core/domain"
	"bae-backend/internal/core/port"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	var collection = db.NewCollection("users")
	return &UserRepository{
		collection,
	}
}

// CreateUser creates a new user in the database
func (ur *UserRepository) Create(user *domain.User) (*domain.User, error) {
	var _, err = ur.GetByPhone(user.Phone)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}
	if err == nil {
		return nil, domain.ErrThisElementIsAlredyExist
	}

	user.Id = domain.NewId()
	err = ur.collection.InsertOne(user)
	return user, err
}

// GetUserByID gets a user by ID from the database
func (ur *UserRepository) GetByID(id domain.Id) (*domain.User, error) {
	var user = new(domain.User)
	var err error
	if err = ur.collection.FindOneById(id, user); err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates a user by ID in the database
func (ur *UserRepository) Save(user *domain.User) (*domain.User, error) {

	return nil, nil
}

func (ur *UserRepository) GetByPhone(phone domain.Phone) (*domain.User, error) {
	var user = new(domain.User)
	var err = ur.collection.FindOne(bson.D{
		{"phone.number", phone.Number},
		{"phone.callCode", phone.CallCode.Int64()},
	}, user)
	return user, err
}

// DeleteUser deletes a user by ID from the database
func (ur *UserRepository) Delete(id domain.Id) error {

	return nil
}

func (ur *UserRepository) GetAll() ([]domain.User, error) {
	var users = make([]domain.User, 0)
	var err = ur.collection.FindMany(bson.M{}, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) GetByFullPhone(fullPhone string) (*domain.User, error) {
	var user = new(domain.User)
	var err = ur.collection.FindOne(
		bson.D{
			{"$expr",
				bson.D{
					{"$eq",
						bson.A{
							bson.D{
								{"$concat",
									bson.A{
										bson.D{{"$toString", "$phone.callCode"}},
										bson.D{{"$toString", "$phone.number"}},
									},
								},
							},
							fullPhone,
						},
					},
				},
			},
		}, user)
	return user, err
}
