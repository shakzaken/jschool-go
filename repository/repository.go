package repository

import "jschool/models"




type DatabaseRepo interface {
	 GetAllUsers() ([]models.User,error)
	 CreateUser(models.User) (models.User,error)
	 DeleteUser(id int) error
	 UpdateUser(models.User) error
	 GetUserByEmail(email string) (models.User,error)

	 CreateUserImage(image models.UserImage) (models.UserImage,error)
	 GetUserImages(userId int) ([]models.UserImage,error)
	 DeleteUserImage(id int) (models.UserImage,error)

	 GetAllCourses() ([]models.Course,error)
	 GetCourse(id int) (models.Course,error)
	 CreateCourse(course models.Course) (models.Course,error)
	 UpdateCourse(course models.Course) (models.Course,error)
	 DeleteCourse(id int) (models.Course,error)

	 CreateCourseComment(comment models.CourseComment) (models.CourseComment,error)
	 GetCourseComments(courseId int) ([]models.CourseComment,error)
	 DeleteCourseComment(commentId int) (models.CourseComment,error)

	 CreateCourseImage(imageBody string, courseId int) (models.CourseImage,error)
	 GetCourseImages(courseId int) ([]models.CourseImage,error)
	 DeleteCourseImage(imageId int) (models.CourseImage,error)

	 GetAllDegrees() ([]models.Degree,error)
	 GetDegree(id int) (models.Degree,error)
	 CreateDegree(degree models.Degree) (models.Degree,error)
	 UpdateDegree(degree models.Degree) (models.Degree,error)
	 DeleteDegree(id int) (models.Degree,error)

	 CreateDegreeComment(comment models.DegreeComment) (models.DegreeComment, error)
	 GetDegreeComments(degreeId int) ([]models.DegreeComment,error)
	 DeleteDegreeComment(commentId int) (models.DegreeComment,error)

	 CreateDegreeImage(image models.DegreeImage) (models.DegreeImage,error)
	 GetDegreeImages(degreeId int) ([]models.DegreeImage,error)
	 DeleteDegreeImage(imageId int) (models.DegreeImage,error)
}