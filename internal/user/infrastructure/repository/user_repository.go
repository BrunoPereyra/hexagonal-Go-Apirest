package repository

import (
	"context"
	"errors"
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

func CreateUser(user *domain.User) error {
	r := NewUserRepository()
	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func FindUser(user *domain.User) (string, *domain.User, error) {
	r := NewUserRepository()
	findUser := bson.D{
		{Key: "username", Value: user.Username},
	}
	var userExist domain.User
	err := r.collection.FindOne(context.TODO(), findUser).Decode(&userExist)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "no exist", nil, nil
		} else {
			return "", nil, errors.New("internal Server error")
		}
	}

	return "exist", &userExist, nil
}
