package application

import (
	"errors"
	"hexagonal/internal/user/domain"
	"hexagonal/internal/user/infrastructure/repository"
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

	findUser := repository.FindUser(user)

	if findUser != "ok" {
		return nil, errors.New("error find user or user exist")
	}

	repository.Save(user)

	return user, nil
}

// Otros métodos y funcionalidad relacionados con la lógica de usuario.
