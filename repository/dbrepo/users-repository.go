package dbrepo

import (
	"context"
	"jschool/models"
	"time"
)

func (dbRepo *PostgressDbRepo) GetAllUsers() ([]models.User,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()


	users := []models.User{};
	db := dbRepo.DB
	query := `select id,name,email from users`;
	rows,err := db.QueryContext(ctx,query)
	if(err != nil){
		return users,err
	}

	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.Id,&user.Name,&user.Email)
		if(err != nil){
			return users,err
		}
		users = append(users, user)
	}
	
	
	return users,nil;

}

func (dbRepo *PostgressDbRepo) CreateUser(user models.User) (int,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()

	
	query := `insert into users(name,email,password)
				values($1,$2,$3) returning id`;
	db := dbRepo.DB
	row := db.QueryRowContext(ctx,query,user.Name,user.Email,user.Password);

	var userId int;
	err := row.Scan(&userId)
	if(err != nil){
		return userId ,err;
	}
	return userId,nil;

}

func (dbRepo * PostgressDbRepo) DeleteUser(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()

	query:= `delete from users
				where id = $1`;
	_,err := dbRepo.DB.ExecContext(ctx,query,id);
	return err;
}

func (dbRepo *PostgressDbRepo) UpdateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()

	query:= `update users 
			 set name=$2,email=$3
			 where id = $1`;
	_,err := dbRepo.DB.ExecContext(ctx,query,user.Id,user.Name,user.Email);
	return err;
}
 
