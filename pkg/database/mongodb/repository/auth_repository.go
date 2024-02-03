package repository

import (
	"OrgManagementApp/pkg/database/mongodb/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository handles database operations related to users
type UserRepository struct {
	collection *mongo.Collection
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *mongo.Database) *UserRepository {
	collection := db.Collection("users")
	return &UserRepository{collection: collection}
}

// SaveUser saves a new user to the database
func (r *UserRepository) SaveUser(user models.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

// GetUserByEmail retrieves a user by email from the database
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	filter := bson.D{{Key: "email", Value: email}}

	err := r.collection.FindOne(context.Background(), filter).Decode(&user)

	// Check if there was no document found
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	// Handle other errors
	if err != nil {
		return nil, err
	}

	return &user, nil
}
