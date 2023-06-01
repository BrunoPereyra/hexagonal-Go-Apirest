package application_test

import (
	"fmt"
	"hexagonal/internal/user/application"
	"hexagonal/internal/user/domain"
	"io/ioutil"
	"mime/multipart"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {

	file, err := ioutil.TempFile("", "test-image-*.jpeg")
	if err != nil {
		t.Fatalf("Error al crear el archivo temporal: %v", err)
	}
	defer os.Remove(file.Name())

	// Configurar los datos de prueba
	userData := &domain.UserModelValidator{
		FullName:  "John Doe202",
		NameUser:  "johndoe",
		Password:  "password",
		Pais:      "Country",
		Ciudad:    "City",
		Email:     "johndoe@example.com",
		Instagram: "johndoe_insta",
		Twitter:   "johndoe_twitter",
		Youtube:   "johndoe_youtube",
	}

	fileHeader := &multipart.FileHeader{
		Filename: file.Name(),
	}

	user, err := application.CreateUser(userData, fileHeader)
	fmt.Println(err, "+++++++++++++")
	assert.Equal(t, userData.FullName, user.FullName)
}
