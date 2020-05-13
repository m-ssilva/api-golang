package repositories

import (
	dbConn "github.com/m-ssilva/api-golang/database"
	models "github.com/m-ssilva/api-golang/models"
)

var db = dbConn.CreateDatabaseConnection()

// CreateUser inserts a new user into database
func CreateUser(user models.User) error {
	stmt, err := db.Prepare("INSERT INTO users(name, email, password) VALUES(?,?,?)")
	_, err = stmt.Exec(user.Name, user.Email, user.Password)
	return err
}

// GetUserByEmail returns a user from database using a email as parameter
func GetUserByEmail(email string) (models.User, error) {
	result, err := db.Query("SELECT * FROM users WHERE email = ?", email)
	var user models.User
	for result.Next() {
		err = result.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	}
	return user, err
}
