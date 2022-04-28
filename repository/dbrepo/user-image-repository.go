package dbrepo

import (
	"context"
	"jschool/models"
	"time"
)

func (dbRepo *PostgressDbRepo) CreateUserImage(image models.UserImage) (models.UserImage,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `insert into users_images(user_id,image) values($1,$2) 
				returning id`
	row := db.QueryRowContext(ctx,query,image.UserId,image.Image)
	err := row.Scan(&image.Id)
	return image,err;

}

func (dbRepo *PostgressDbRepo) GetUserImages(userId int) ([]models.UserImage,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `select id,user_id,image from users_images 
				where user_id = $1`
	rows,err := db.QueryContext(ctx,query,userId)
	images := []models.UserImage{}
	if(err != nil){
		return images,err
	}
	for rows.Next() {
		image := models.UserImage{}
		err = rows.Scan(&image.Id,&image.UserId,&image.Image)
		if(err != nil){
			return images,err
		}
		images = append(images,image)
	}
	return images,nil
	
}

func (dbRepo *PostgressDbRepo) DeleteUserImage(id int) (models.UserImage,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `delete from users_images where id = $1 
				returning user_id`
	row := db.QueryRowContext(ctx,query,id)
	image := models.UserImage{Id: id}
	err := row.Scan(&image.UserId)
	return image,err
}