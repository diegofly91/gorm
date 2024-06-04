package repository

import (
	"gorm/src/modules/user/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	FindAll() models.Users
	FindById(id int) (models.User, error)
	Update(user models.User) models.User
	Deleted(id int) models.User
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

func (r *userRepository) FindAll() models.Users {
	users := models.Users{}
	r.db.Find(&users)
	return users
}

func (r *userRepository) FindById(id int) (models.User, error) {
	user := models.User{}
	if err := r.db.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) Update(user models.User) models.User {
	r.db.Save(&user)
	return user
}

func (r *userRepository) Deleted(id int) models.User {
	user := models.User{}
	r.db.First(&user, id)
	r.db.Delete(&user)
	return user
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
