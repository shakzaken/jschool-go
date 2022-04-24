package validations

import "jschool/models"


func CreateCourse(course models.Course) bool {
	if len(course.Name) < 4 || len(course.Description)<4  {
		return false
	}
	return true
}

func UpdateCourse(course models.Course) bool {
	if len(course.Name) < 4 || len(course.Description)<4 || course.Id <=0  {
		return false
	}
	return true
}