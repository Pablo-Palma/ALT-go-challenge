package auth

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db	*mongo.Collection
}

func NewUserRepository (db *mongo.Database) *UserRepository {
	return &UserRepository {
		db: db.Collection("users"),
	}
}

func (repo *UserRepository) Create (user User) (primitive.ObjectID, error) {
	user.ID = primitive.NewObjectID()
	_, err := repo.db.InsertOne(context.Background(), user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return user.ID, nil
}

func (repo *UserRepository) GetByUsername(username string) (User, error) {
	var user User
	err := repo.db.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}
