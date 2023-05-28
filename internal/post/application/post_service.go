package application

import (
	"errors"
	"hexagonal/internal/post/infrastructure/repository"
	"mime/multipart"
)

func a() {
}

func CreatePost(fileHeader *multipart.FileHeader, TextPost string) (string, repository.Post, error) {
	CreatePostRes, NewPost, errCreatePostRes := repository.CreatePost(TextPost, fileHeader)
	if errCreatePostRes != nil {
		return "", repository.Post{}, errors.New(errCreatePostRes.Error())
	}
	return CreatePostRes, NewPost, nil
}
