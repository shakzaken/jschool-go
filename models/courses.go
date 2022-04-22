package models

import ("time")

type CourseImage struct {
	Id int
	Image string
	CourseId int
}

type CourseComment struct {
	Id int
	CourseId int
	UserId int
	Comment string
	Date time.Time  
}

type Couese struct {
	Id int
	Name string
	Description string
	Comments []CourseComment
	Images []CourseImage
}