package services

import (
	"gorm/src/modules/user/models"
	"gorm/src/modules/user/repositories"
)

func FindAll() models.Users {
	return repositories.FindAll()
}

func FindById(id int) models.User {
	return repositories.FindById(id)
}

func Create(user models.User) models.User {
	return repositories.Create(user)
}

func Update(user models.User) models.User {
	return repositories.Update(user)
}

func Delete(userId int) models.User {
	return repositories.Delete(userId)
}
