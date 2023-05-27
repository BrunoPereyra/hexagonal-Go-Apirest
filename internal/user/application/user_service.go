package application

import (
	"errors"
	"hexagonal/internal/user/domain"
	"hexagonal/internal/user/infrastructure/repository"
	"hexagonal/pkg/auth"
	"hexagonal/pkg/bcrypt"
)

// CreateUser crea un nuevo usuario.
func CreateUser(username, password string) (*domain.User, error) {
	// Lógica para crear un nuevo usuario, como validar datos, generar ID, encriptar contraseña, etc.
	// Utiliza el repositorio de usuarios para persistir los datos.
	if len(username) <= 2 {
		return nil, errors.New("userName > 2a")
	}
	passwordhash, errPasswordHash := bcrypt.HashPassword(password)
	if errPasswordHash != nil {
		return nil, errors.New("passwordHashError")
	}
	user := &domain.User{
		Username:     username,
		PasswordHash: passwordhash,
		// Otros campos...
	}

	exist, _, err := repository.FindUser(user)
	if err != nil {
		return nil, err
	}
	if exist != "no exist" {
		return nil, errors.New("user exist")
	}

	repository.CreateUser(user)

	return user, nil
}

// Otros métodos y funcionalidad relacionados con la lógica de usuario.

func Login(username, password string) (string, error) {
	user := &domain.User{
		Username:     username,
		PasswordHash: password,
	}
	if len(user.PasswordHash) < 2 || len(user.Username) < 2 {
		return "", errors.New("len(user.PasswordHash) < 2 || len(user.Username) < 2")
	}
	exist, userLogin, err := repository.FindUser(user)
	if err != nil {
		return "", err
	}
	if exist == "no exist" {
		return "", errors.New("nameUser incorrect")
	}
	errDecodePassword := bcrypt.DecodePassword(userLogin.PasswordHash, user.PasswordHash)
	if errDecodePassword != nil {
		return "", errors.New("password error")
	}
	token, errTokenSigned := auth.CreateToken(user)
	if errTokenSigned != nil {
		return "", errTokenSigned
	}
	return token, nil
}
