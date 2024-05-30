package repository

import (
	"gorm/src/modules/user/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

/*
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
*/
