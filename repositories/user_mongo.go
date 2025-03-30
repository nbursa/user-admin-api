package repositories

import (
	"context"

	"github.com/nbursa/user-admin-api/config"
	"github.com/nbursa/user-admin-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *userMongoRepository) FindByNameOrEmailPaginated(ctx context.Context, search string, page, limit int) ([]*models.User, int, error) {
	filter := bson.M{}
	if search != "" {
		regex := primitive.Regex{Pattern: search, Options: "i"}
		filter = bson.M{
			"$or": []bson.M{
				{"name": bson.M{"$regex": regex}},
				{"email": bson.M{"$regex": regex}},
			},
		}
	}

	skip := (page - 1) * limit

	opts := options.Find()
	opts.SetSkip(int64(skip))
	opts.SetLimit(int64(limit))

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var users []*models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, 0, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, 0, err
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return users, int(count), nil
}
