package repository

import (
	"context"
	"errors"
	"fmt"
	"hexagonal/internal/post/domain"
	"hexagonal/internal/post/infrastructure/database"
	"hexagonal/pkg/cloudinary/processimage"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewUserRepository() *PostRepository {
	client, errdb := database.Connect()

	if errdb != nil {
		fmt.Println("NewUserRepository DB error")
	}
	collection := client.Database("goMoongodb").Collection("post")
	return &PostRepository{
		client:     client,
		collection: collection,
	}
}

func CreatePost(Post *domain.Post, fileHeader *multipart.FileHeader) (string, error) {

	PostImageChanel := make(chan string)
	errChanel := make(chan error)
	go processimage.Processimage(fileHeader, PostImageChanel, errChanel)

	db := NewUserRepository()
	defer db.client.Disconnect(context.TODO())

	select {
	case PostImage := <-PostImageChanel:
		Post.ImgPostUrl = PostImage
		Post.Likes = []primitive.ObjectID{}
	case err := <-errChanel:
		return "", err
	}

	_, err := db.collection.InsertOne(context.TODO(), Post)
	if err != nil {
		return "", errors.New("error InsertOne")
	}
	return "ok", nil
}
