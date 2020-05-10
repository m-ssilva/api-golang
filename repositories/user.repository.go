package repositories

import (
	dbConn "github.com/m-ssilva/api-golang/database"
	models "github.com/m-ssilva/api-golang/models"
)

// CreateUser inserts a new user into database
func CreateUser(user models.User) bool {
	db := dbConn.CreateDatabaseConnection()
	stmt, err := db.Prepare("INSERT INTO users(name, email, password) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		panic(err.Error())
	}

	return true
}
