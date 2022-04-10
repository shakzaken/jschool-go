package repository

import "jschool/repository/models"




type DatabaseRepo interface {
	 GetAllUsers() ([]models.User,error)
	 CreateUser(models.User) (int,error)
}