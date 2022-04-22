package repository

import "jschool/models"




type DatabaseRepo interface {
	 GetAllUsers() ([]models.User,error)
	 CreateUser(models.User) (int,error)
	 DeleteUser(id int) error
	 UpdateUser(models.User) error
}