package models

type User struct {
	Id int 			`json:"id"`
	Name string		`json:"name"`
	Email string 	`json:"email"`
	Password string `json:"password"`
}

type UserImage struct {
	Id int  	`json:"id"`
	Image string `json:"image"`
	UserId int  `json:"userId"`
}