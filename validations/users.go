package validations

import (
	"jschool/models"
)

func CreateUser(user models.User) bool {

	if len(user.Name) < 4 {
		return false;
	}
	if len(user.Email)<4 {
		return false
	}
	if len(user.Password) < 4 {
		return false;
	}
	return true;
}

func UpdateUser(user models.User) bool {
	
	if user.Id <=0 {
		return false;
	}
	if len(user.Name) < 4 {
		return false;
	}
	if len(user.Email)< 4 {
		return false
	}
	
	return true;
}