package service

import (
	"gorm/src/modules/user/models"
	"gorm/src/modules/user/repository"
)

type UserService interface {
	CreateUser(user models.User) (models.User, error)
	FindAll() models.Users
	FindById(id int) (models.User, error)
	Update(user models.User) models.User
	Delete(userId int) models.User
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
	return s.repo.CreateUser(user)
}

func (s *userService) FindAll() models.Users {
	return s.repo.FindAll()
}

func (s *userService) FindById(id int) (models.User, error) {
	return s.repo.FindById(id)
}

func (s *userService) Update(user models.User) models.User {
	return s.repo.Update(user)
}

func (s *userService) Delete(userId int) models.User {
	return s.repo.Deleted(userId)
}

/*
func FindAll() models.Users {
	return repositories.FindAll()
}

func FindById(id int) models.User {
	return repositories.FindById(id)
}

func Update(user models.User) models.User {
	return repositories.Update(user)
}

func Delete(userId int) models.User {
	return repositories.Delete(userId)
}
*/
