package domain

import (
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Avatar       string             `json:"avatar"`
	FullName     string             `json:"FullName"`
	NameUser     string             `json:"NameUser"`
	PasswordHash string             `json:"passwordHash"`
	Pais         string             `json:"Pais"`
	Ciudad       string             `json:"Ciudad"`
	Email        string             `json:"Email"`
	Instagram    string             `json:"instagram,omitempty"`
	Twitter      string             `json:"twitter,omitempty"`
	Youtube      string             `json:"youtube,omitempty"`
}

// validators

// signup

type UserModelValidator struct {
	FullName  string `json:"fullName" validate:"required,max=70"`
	NameUser  string `json:"NameUser" validate:"required,max=70"`
	Password  string `json:"password" validate:"required,min=8"`
	Pais      string `json:"Pais" validate:"required"`
	Ciudad    string `json:"Ciudad" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Instagram string `json:"instagram" default:""`
	Twitter   string `json:"twitter" default:""`
	Youtube   string `json:"youtube" default:""`
}

func (u *UserModelValidator) ValidateUserFind() error {
	validate := validator.New()
	return validate.Struct(u)
}

// login
type LoginValidatorStruct struct {
	NameUser string `json:"NameUser" validate:"required,max=70"`
	Password string `json:"password" validate:"required,min=8"`
}

func (L *LoginValidatorStruct) LoginValidator() error {
	validate := validator.New()
	return validate.Struct(L)
}
