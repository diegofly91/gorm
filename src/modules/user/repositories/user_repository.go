package repositories

import (
	"fmt"
	db "gorm/src/config"
	"gorm/src/modules/user/models"
)

func FindAll() models.Users {
	users := models.Users{}
	db.Database.Find(&users)
	return users
}

func FindById(id int) models.User {
	user := models.User{}
	db.Database.First(&user, id)
	return user
}

func Create(user models.User) models.User {
	fmt.Println("create user  ", user)
	db.Database.Create(&user)
	return user
}

func Update(user models.User) models.User {
	db.Database.Save(&user)
	return user
}

func Delete(userId int) models.User {
	user := models.User{}
	db.Database.First(&user, userId)
	db.Database.Delete(&user)
	return user
}
