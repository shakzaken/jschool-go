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

type Course struct {
	Id int				`json:"id"`
	Name string			`json:"name"`
	Description string	`json:"description"`
	Comments []CourseComment	`json:"comments"`
	Images []CourseImage		`json:"images"`
}