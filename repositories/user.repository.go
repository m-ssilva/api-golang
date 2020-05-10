package repositories

import (
	"fmt"

	dbConn "github.com/m-ssilva/api-golang/database"
	models "github.com/m-ssilva/api-golang/models"
)

// CreateUser inserts a new user into database
func CreateUser(user models.User) {
	fmt.Println("User received in repository layer")
	fmt.Println(user)
	db := dbConn.CreateDatabaseConnection()
	stmt, err := db.Prepare("INSERT INTO users(name, email, password) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Inserting into database...")
	_, err = stmt.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("User inserted into database")
}
