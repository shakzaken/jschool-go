package validations

import "jschool/models"


func CreateDegree(degree models.Degree) bool {
	if len(degree.Name) < 4 || len(degree.Description)<4  {
		return false
	}
	return true
}

func UpdateDegree(degree models.Degree) bool {
	if len(degree.Name) < 4 || len(degree.Description)<4 || degree.Id <=0  {
		return false
	}
	return true
}