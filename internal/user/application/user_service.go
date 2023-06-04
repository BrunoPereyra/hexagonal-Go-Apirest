package application

import (
	"errors"
	"hexagonal/internal/user/domain"
	"hexagonal/internal/user/infrastructure/repository"
	"hexagonal/pkg/auth"
	"hexagonal/pkg/bcrypt"
	"mime/multipart"
)

// CreateUser crea un nuevo usuario.
func CreateUser(userData *domain.UserModelValidator, fileHader *multipart.FileHeader) (domain.User, error) {
	var modelNewUser domain.User

	passwordhashChan := make(chan string)
	go bcrypt.HashPassword(userData.Password, passwordhashChan)
	modelNewUser.NameUser = userData.NameUser
	exist, _, err := repository.FindUser(&modelNewUser)
	if err != nil {
		return domain.User{}, err
	}
	if exist != "no exist" {
		return domain.User{}, errors.New("user exist")
	}

	passwordHash := <-passwordhashChan
	if passwordHash == "errHash" {
		return domain.User{}, errors.New("errHash password")
	}

	userData.Password = passwordHash

	modelNewUser.FullName = userData.FullName
	modelNewUser.NameUser = userData.NameUser
	modelNewUser.PasswordHash = userData.Password
	modelNewUser.Pais = userData.Pais
	modelNewUser.Ciudad = userData.Ciudad
	modelNewUser.Email = userData.Email
	modelNewUser.Instagram = userData.Instagram
	modelNewUser.Twitter = userData.Twitter
	modelNewUser.Youtube = userData.Youtube

	Newuser, errorNewuser := repository.CreateUser(modelNewUser, fileHader)
	if errorNewuser != nil {
		return Newuser, errorNewuser
	}

	return Newuser, nil
}

// Otros métodos y funcionalidad relacionados con la lógica de usuario.

func Login(UserLogin *domain.LoginValidatorStruct) (string, error) {
	user := &domain.User{
		NameUser: UserLogin.NameUser,
	}
	exist, user, err := repository.FindUser(user)
	if err != nil {
		return "", err
	}
	if exist == "no exist" {
		return "", errors.New("nameUser incorrect")
	}
	errDecodePassword := bcrypt.DecodePassword(user.PasswordHash, UserLogin.Password)
	if errDecodePassword != nil {
		return "", errors.New("password error")
	}
	token, errTokenSigned := auth.CreateToken(user)
	if errTokenSigned != nil {
		return "", errTokenSigned
	}
	return token, nil
}
