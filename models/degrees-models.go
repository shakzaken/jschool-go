package models

import "time"

type DegreeImage struct {
	Id int			`json:"id"`
	Image string	`json:"image"`
	DegreeId int	`json:"degreeId"`
}

type DegreeComment struct {
	Id int				`json:"id"`
	Comment string		`json:"comment"`
	DegreeId int		`json:"degreeId"`
	UserId int			`json:"userId"`
	User User			`json:"user"`
	Date time.Time		`json:"date"`

}

type Degree struct {
	Id int						`json:"id"`
	Name string					`json:"name"`
	Description string			`json:"description"`
	Comments []CourseComment	`json:"comments"`
	Images []CourseImage		`json:"images"`
}