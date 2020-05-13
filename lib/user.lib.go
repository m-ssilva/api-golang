package lib

import (
	"encoding/json"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"

	models "github.com/m-ssilva/api-golang/models"
	repositories "github.com/m-ssilva/api-golang/repositories"
)

// CreateUser parse body information into user model and call repository layer
func CreateUser(body []byte) error {
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	user := models.User{
		Name:     keyVal["name"],
		Email:    keyVal["email"],
		Password: keyVal["password"],
	}

	err := repositories.CreateUser(user)

	if err != nil {
		return errors.New("EMAIL_ALREADY_IN_USE")
	}

	return err
}

// GetUserByEmail returns user from database using email as parameter
func GetUserByEmail(email string) (models.User, error) {
	return repositories.GetUserByEmail(email)
}

func signToken(password string, userDB models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":             userDB.ID,
		"expirationTime": time.Now().Local().Add(time.Hour * 24),
	})

	tokenString, err := token.SignedString([]byte(`TESTING123`))

	return tokenString, err
}

// AuthenticateUser parse body information and returns a token if information is valid
func AuthenticateUser(body []byte) (string, error) {
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	email := keyVal["email"]
	password := keyVal["password"]

	userDB, err := GetUserByEmail(email)
	if userDB.ID == 0 {
		err = errors.New("EMAIL_NOT_REGISTERED")
		return "", err
	}

	tokenString, err := signToken(password, userDB)
	if err != nil {
		panic(errors.New("ERROR_SIGNING_TOKEN").Error())
	}

	return tokenString, err
}
