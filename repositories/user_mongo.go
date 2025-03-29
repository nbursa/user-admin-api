package repositories

import (
	"context"

	"github.com/nbursa/user-admin-api/config"
	"github.com/nbursa/user-admin-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userMongoRepository struct {
	collection *mongo.Collection
}

func NewUserMongoRepository() UserRepository {
	return &userMongoRepository{
		collection: config.MongoDatabase.Collection("users"),
	}
}

func (r *userMongoRepository) Insert(ctx context.Context, user *models.User) error {
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *userMongoRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	count, err := r.collection.CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}