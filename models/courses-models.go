package models

import ("time")

type CourseImage struct {
	Id int			`json:"id"`
	Image string	`json:"image"`
	CourseId int	`json:"courseId"`
}

type CourseComment struct {
	Id int				`json:"id"`
	CourseId int		`json:"courseId"`
	UserId int			`json:"userId"`
	Comment string		`json:"comment"`
	Date time.Time  	`json:"date"`
}

type Course struct {
	Id int				`json:"id"`
	Name string			`json:"name"`
	Description string	`json:"description"`
	Comments []CourseComment	`json:"comments"`
	Images []CourseImage		`json:"images"`
}