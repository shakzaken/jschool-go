package models

import "time"

type DegreeImage struct {
	Id int
	Image string
	DegreeId int
}

type DegreeComment struct {
	Id int
	Comment string
	DegreeId string
	User User
	Date time.Time

}

type Degree struct {
	Id int			`json:"id"`
	Name string		`json:"name"`
	Description string	`json:"description"`
	Comments []CourseComment	`json:"comments"`
	Images []CourseImage		`json:"images"`
}