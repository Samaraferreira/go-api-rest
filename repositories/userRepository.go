package repositories

import (
	"errors"

	"github.com/Samaraferreira/go-api-rest/database"
	"github.com/Samaraferreira/go-api-rest/models"
)

func FindAll() []models.User {
	var users []models.User
	database.DB.Find(&users)
	return users
}

func FindUserById(id int) (models.User, error) {
	var user models.User

	database.DB.First(&user, id)

	var err error
	if user.Id == 0 {
		err = errors.New("User not found")
	}

	if err != nil {
		return user, err
	}

	return user, nil
}

func AddNewUser(newUser models.User) models.User {
	user := newUser
	database.DB.Create(&user)
	return user
}

func DeleteUser(id int) models.User {
	var user models.User
	database.DB.Delete(&user, id)

	return user
}
