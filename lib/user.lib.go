package lib

import (
	"encoding/json"

	models "github.com/m-ssilva/api-golang/models"
	repositories "github.com/m-ssilva/api-golang/repositories"
)

// CreateUser parse body information into user model and call repository layer
func CreateUser(body []byte) {
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	user := models.User{
		Name:     keyVal["name"],
		Email:    keyVal["email"],
		Password: keyVal["password"],
	}

	repositories.CreateUser(user)
}
