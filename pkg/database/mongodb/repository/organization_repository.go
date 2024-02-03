// repository/organization_repository.go

package repository

import (
	"OrgManagementApp/pkg/database/mongodb/models"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// OrganizationRepository is an interface for organization database interactions
type OrganizationRepository interface {
	CreateOrganization(org *models.Organization) error
	GetOrganizationByID(orgID string) (*models.Organization, error)
	UpdateOrganization(org *models.Organization) error
	DeleteOrganization(orgID string) error
}

// MongoDBOrganizationRepository is a MongoDB implementation of OrganizationRepository
type MongoDBOrganizationRepository struct {
	collection *mongo.Collection
}

// NewMongoDBOrganizationRepository creates a new instance of MongoDBOrganizationRepository
func NewOrganizationRepository(db *mongo.Database) *MongoDBOrganizationRepository {
	collection := db.Collection("organization")
	return &MongoDBOrganizationRepository{collection: collection}
}

// CreateOrganization creates a new organization in the MongoDB database
func (repo *MongoDBOrganizationRepository) CreateOrganization(org *models.Organization) error {
	collection := repo.collection

	_, err := collection.InsertOne(context.Background(), org)
	if err != nil {
		return fmt.Errorf("error inserting organization into MongoDB: %v", err)
	}

	return nil
}

// GetOrganizationByID retrieves an organization by ID from the MongoDB database
func (repo *MongoDBOrganizationRepository) GetOrganizationByID(orgID string) (*models.Organization, error) {
	collection := repo.collection
	var org models.Organization
	filter := bson.M{"id": orgID}

	err := collection.FindOne(context.Background(), filter).Decode(&org)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("organization not found")
		}
		return nil, fmt.Errorf("error retrieving organization from MongoDB: %v", err)
	}

	return &org, nil
}

// GetAllOrganizations retrieves all organizations from MongoDB
func (r *MongoDBOrganizationRepository) GetAllOrganizations() ([]models.Organization, error) {
	var organizations []models.Organization

	cursor, err := r.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &organizations); err != nil {
		return nil, err
	}

	return organizations, nil
}

// UpdateOrganization updates an organization in the MongoDB database
func (repo *MongoDBOrganizationRepository) UpdateOrganization(org *models.Organization) error {
	collection := repo.collection
	filter := bson.M{"id": org.ID}
	update := bson.M{"$set": org}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("organization not found")
		}
		return fmt.Errorf("error updating organization in MongoDB: %v", err)
	}

	return nil
}

// DeleteOrganization deletes an organization from the MongoDB database
func (repo *MongoDBOrganizationRepository) DeleteOrganization(orgID string) error {
	collection := repo.collection
	filter := bson.M{"id": orgID}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("organization not found")
		}
		return fmt.Errorf("error deleting organization from MongoDB: %v", err)
	}

	return nil
}
