package dbrepo

import (
	"context"
	"jschool/repository/models"
	"time"
)

func (dbRepo *PostgressDbRepo) GetAllUsers() ([]models.User,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()


	users := []models.User{};
	db := dbRepo.DB
	query := `select name,email,phone from users`;
	rows,err := db.QueryContext(ctx,query)
	if(err != nil){
		return users,err
	}

	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.Name,&user.Email,&user.Phone)
		users = append(users, user)
	}
	
	if(err != nil){
		return users,err
	}
	return users,nil;

}

func (dbRepo *PostgressDbRepo) CreateUser(user models.User) (int,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()

	
	query := `insert into users(name,email,phone)
				values($1,$2,$3) returning id`;
	db := dbRepo.DB
	row := db.QueryRowContext(ctx,query,user.Name,user.Email,user.Phone);

	var userId int;
	err := row.Scan(&userId)
	if(err != nil){
		return userId ,err;
	}
	return userId,nil;

}

