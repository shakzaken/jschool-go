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
	Id int
	Name string
	Description string
}