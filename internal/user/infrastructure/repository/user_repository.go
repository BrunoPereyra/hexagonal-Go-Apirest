package repository

import (
	"context"
	"errors"
	"fmt"
	"hexagonal/internal/user/domain"
	"hexagonal/internal/user/infrastructure/database"
	"hexagonal/pkg/cloudinary/processimage"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	client, errdb := database.Connect()
	if errdb != nil {
		fmt.Println("NewUserRepository DB error")
	}
	collection := client.Database("goMoongodb").Collection("post")

	return &UserRepository{
		client:     client,
		collection: collection,
	}
}

func CreateUser(user domain.User, FileHeader *multipart.FileHeader) (domain.User, error) {
	PostImageChanel := make(chan string)
	errChanel := make(chan error)
	go processimage.Processimage(FileHeader, PostImageChanel, errChanel)

	r := NewUserRepository()
	defer r.client.Disconnect(context.Background())
	select {
	case avatar := <-PostImageChanel:
		user.Avatar = avatar
	case err := <-errChanel:
		fmt.Println(err, "{}{}{}{}{}{}{}")
		return domain.User{}, err
	}
	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil

}

func FindUser(user *domain.User) (string, *domain.User, error) {
	r := NewUserRepository()
	findUser := bson.D{
		{Key: "nameuser", Value: user.NameUser},
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
