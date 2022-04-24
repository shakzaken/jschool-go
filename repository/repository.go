package repository

import "jschool/models"




type DatabaseRepo interface {
	 GetAllUsers() ([]models.User,error)
	 CreateUser(models.User) (int,error)
	 DeleteUser(id int) error
	 UpdateUser(models.User) error

	 GetAllCourses() ([]models.Course,error)
	 GetCourse(id int) (models.Course,error)
	 CreateCourse(course models.Course) (models.Course,error)
	 UpdateCourse(course models.Course) (models.Course,error)
	 DeleteCourse(id int) (models.Course,error)
}