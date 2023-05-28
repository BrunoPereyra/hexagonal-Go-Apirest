package repository

import (
	"context"
	"errors"
	"fmt"
	"hexagonal/internal/post/infrastructure/database"
	"hexagonal/pkg/cloudinary/processimage"
	"mime/multipart"

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

type Post struct {
	TextPost   string `json:"textPost" bson:"textPost"`
	ImgPostUrl string `json:"imgPostUrl" bson:"imgPostUrl"`
}

func CreatePost(TextPost string, fileHeader *multipart.FileHeader) (string, Post, error) {

	PostImageChanel := make(chan string)
	errChanel := make(chan error)
	go processimage.Processimage(fileHeader, PostImageChanel, errChanel)

	db := NewUserRepository()
	defer db.client.Disconnect(context.TODO())

	var NewPost Post
	NewPost.TextPost = TextPost

	select {
	case PostImage := <-PostImageChanel:
		NewPost.ImgPostUrl = PostImage
	case err := <-errChanel:
		return "", Post{}, err
	}

	_, err := db.collection.InsertOne(context.TODO(), NewPost)
	if err != nil {
		return "", Post{}, errors.New("error InsertOne")
	}
	return "ok", NewPost, nil
}
