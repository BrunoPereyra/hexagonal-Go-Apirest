package repository

import (
	"context"
	"fmt"
	"hexagonal/internal/user/domain"
	"hexagonal/internal/user/infrastructure/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	db, errdb := database.Connect()
	if errdb != nil {
		fmt.Println("NewUserRepository DB error")
	}
	collection := db.Collection("users")
	return &UserRepository{
		collection: collection,
	}
}

func Save(user *domain.User) error {
	r := NewUserRepository()
	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func FindUser(user *domain.User) string {
	r := NewUserRepository()
	findUser := bson.D{
		{Key: "username", Value: user.Username},
	}
	var userExist domain.User
	if err := r.collection.FindOne(context.TODO(), findUser).Decode(&userExist); err != nil {
		if err == mongo.ErrNoDocuments {
			return "ok"
		} else {
			return "exist user"
		}
	}

	return "?"
}
