package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username     string
	PasswordHash string
}

// Métodos y funcionalidad relacionados con el dominio de usuario.
