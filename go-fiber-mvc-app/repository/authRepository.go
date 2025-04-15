package repository

import(
	"go-fiber-mvc-app/model"
	"go-fiber-mvc-app/config"
)

func GetUserByUsername(username string)(model.User, error) {
	var user model.User
	result := config.DB.Where("username = ?",username).First(&user)
	return user, result.Error
}