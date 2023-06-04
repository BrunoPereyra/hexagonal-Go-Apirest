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

func CreateUser(user domain.User, FileHeader *multipart.FileHeader) (domain.User, error) {
	PostImageChanel := make(chan string)
	errChanel := make(chan error)
	go processimage.Processimage(FileHeader, PostImageChanel, errChanel)

	r := NewUserRepository()
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
	fmt.Println(userExist.NameUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "no exist", nil, nil
		} else {
			return "", nil, errors.New("internal Server error")
		}
	}

	return "exist", &userExist, nil
}
