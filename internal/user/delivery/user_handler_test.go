package delivery_test

import (
	"encoding/json"
	"hexagonal/internal/user/delivery"
	"hexagonal/internal/user/domain"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func Test_service(t *testing.T) {
	app := fiber.New()
	app.Post("login", delivery.Login)

	testCases := []struct {
		Data   domain.LoginValidatorStruct `json:"data"`
		Expect int
	}{
		{
			Data: domain.LoginValidatorStruct{
				NameUser: "bruno",
				Password: "123456789",
			},
			Expect: 200,
		},
		// {
		// 	Data: domain.LoginValidatorStruct{
		// 		NameUser: "bruno78",
		// 		Password: "123456789",
		// 	},
		// 	Expect: 100,
		// },
		// {
		// 	Data: domain.LoginValidatorStruct{
		// 		NameUser: "bruno",
		// 		Password: "1234567",
		// 	},
		// 	Expect: 100,
		// },
	}
	for _, TC := range testCases {
		requestBody, _ := json.Marshal(TC.Data)
		req := httptest.NewRequest(
			"POST",
			"login",
			strings.NewReader(string(requestBody)),
		)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req, 2000)

		if err != nil {
			t.Fatalf("Error al realizar la solicitud: %s", err.Error())
		}
		defer resp.Body.Close()

		if TC.Expect == resp.StatusCode {
			t.Logf("Caso %s - Esperado: %d, Respuesta: %s - Correcto", TC.Data, TC.Expect, resp.Status)
		} else {
			t.Errorf("Caso %s - Esperado: %d, Respuesta: %s - Incorrecto", TC.Data, TC.Expect, resp.Status)
		}
	}

}
