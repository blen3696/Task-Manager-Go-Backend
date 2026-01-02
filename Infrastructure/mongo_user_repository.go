package infrastructure

import (
	"context"
	"time"

	"github.com/blen/task_manager_api/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	col *mongo.Collection
}

func NewMongoUserRepository(col *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{col}
}

func (m *MongoUserRepository) Create(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.col.InsertOne(ctx, user)
	return err
}

func (m *MongoUserRepository) FindByEmail(email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user domain.User
	err := m.col.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return &user, err
}

func (m *MongoUserRepository) PromoteToAdmin(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.col.UpdateOne(
		ctx,
		bson.M{"id": userID},
		bson.M{"$set": bson.M{"role": "admin"}},
	)
	return err
}
