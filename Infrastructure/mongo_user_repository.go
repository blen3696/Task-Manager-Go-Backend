package infrastructure

import (
	"context"
	"time"

	"github.com/blen/task_manager_api/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	col *mongo.Collection
}

func NewMongoUserRepository(col *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{col}
}

// Mongo-specific struct
type mongoUser struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
}

func (m *MongoUserRepository) Create(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mUser := mongoUser{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}

	result, err := m.col.InsertOne(ctx, mUser)
	if err != nil {
		return err
	}

	// Save the generated ObjectID as hex string in domain.User
	user.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (m *MongoUserRepository) FindByEmail(email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var mUser mongoUser
	err := m.col.FindOne(ctx, bson.M{"email": email}).Decode(&mUser)
	if err != nil {
		return nil, err
	}

	// Map mongoUser → domain.User
	user := &domain.User{
		ID:       mUser.ID.Hex(),
		Username: mUser.Username,
		Email:    mUser.Email,
		Password: mUser.Password,
		Role:     mUser.Role,
	}

	return user, nil
}

func (m *MongoUserRepository) PromoteToAdmin(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	_, err = m.col.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"role": "admin"}},
	)
	return err
}
