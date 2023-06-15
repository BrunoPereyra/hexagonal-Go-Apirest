package application

import (
	"errors"
	"hexagonal/internal/post/domain"
	"hexagonal/internal/post/infrastructure/repository"
	"mime/multipart"
)

func CreatePost(fileHeader *multipart.FileHeader, Post *domain.Post) (string, *domain.Post, error) {
	CreatePostRes, errCreatePostRes := repository.CreatePost(Post, fileHeader)
	if errCreatePostRes != nil {
		return "", Post, errors.New(errCreatePostRes.Error())
	}
	return CreatePostRes, Post, nil
}
