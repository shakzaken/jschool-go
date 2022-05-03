package dbrepo

import (
	"context"
	"jschool/models"
	"time"

	"golang.org/x/crypto/bcrypt"
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



func (dbRepo *PostgressDbRepo) CreateUser(user models.User) (models.User,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()

	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(user.Password),14)
	user.Password = ""
	if(err != nil){
		return user,err
	}
	
	query := `insert into users(name,email,password)
				values($1,$2,$3) returning id`;
	db := dbRepo.DB
	row := db.QueryRowContext(ctx,query,user.Name,user.Email,hashedPassword);

	err = row.Scan(&user.Id)
	if(err != nil){
		return user ,err;
	}
	return user,nil;

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

func (dbRepo *PostgressDbRepo) GetUserByEmail(email string) (models.User,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db:= dbRepo.DB

	user := models.User{Email: email}
	query := `select id,name,password from users where email = $1`
	row := db.QueryRowContext(ctx,query,email)
	err := row.Scan(&user.Id,&user.Name,&user.Password)
	return user,err
}
 
