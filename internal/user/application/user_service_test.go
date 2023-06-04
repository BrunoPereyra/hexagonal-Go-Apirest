package application_test

import (
	"hexagonal/internal/user/application"
	"hexagonal/internal/user/domain"
	"mime/multipart"
	"os"
	"testing"
)

func TestSignup(t *testing.T) {
	// Crear un archivo temporal para la prueba
	// Resto de tu código de prueba
	TestCases := []struct {
		Data   *domain.UserModelValidator
		Expect string
	}{
		{
			Data: &domain.UserModelValidator{
				FullName:  "John Doe202",
				NameUser:  "johndoe",
				Password:  "password",
				Pais:      "Country",
				Ciudad:    "City",
				Email:     "johndoe@example.com",
				Instagram: "johndoe_insta",
				Twitter:   "johndoe_twitter",
				Youtube:   "johndoe_youtube",
			},
			Expect: "ok",
		},
	}

	filePath := "./1366_2000.jpeg"
	fileHeader, err := ReadImage(filePath)
	if err != nil {
		// Manejar el error
		return
	}

	for _, TC := range TestCases {
		user, err := application.CreateUser(TC.Data, fileHeader)
		if user.NameUser == "" {
			t.Logf("Error: %s", err)
		} else {
			t.Logf("Esperado: %s, Respuesta: %s", TC.Data.NameUser, user.NameUser)
		}
	}
}

func ReadImage(filePath string) (*multipart.FileHeader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Obtener información del archivo
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// Crear un FileHeader
	fileHeader := &multipart.FileHeader{
		Filename: fileInfo.Name(),
		Size:     fileInfo.Size(),
	}

	return fileHeader, nil
}
