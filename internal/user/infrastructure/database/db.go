package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Database, error) {
	// Configura los detalles de conexión a MongoDB
	URI := "mongodb+srv://bruno:Ej622XsRFaiuh9kg@cluster0.0leyfts.mongodb.net/TDD?retryWrites=true&w=majority"

	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Comprueba si la conexión es exitosa
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Obtiene una instancia de la base de datos
	db := client.Database("goMoongodb")

	return db, nil
}
