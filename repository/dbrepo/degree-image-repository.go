package dbrepo

import (
	"context"
	"jschool/models"
	"time"
)

func (dbRepo *PostgressDbRepo) CreateDegreeImage(image models.DegreeImage) (models.DegreeImage,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `insert into degrees_images(degree_id,image) values($1,$2)
				returning id`

	row := db.QueryRowContext(ctx,query,image.DegreeId,image.Image)
	err := row.Scan(&image.Id)
	return image,err
}

func (dbRepo *PostgressDbRepo) GetDegreeImages(degreeId int) ([]models.DegreeImage,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `select id,degree_id,image from degrees_images
				where degree_id = $1`
	images := []models.DegreeImage{}
	rows,err := db.QueryContext(ctx,query,degreeId)
	if(err != nil){
		return images,err
	}
	for rows.Next() {
		image := models.DegreeImage{}
		err = rows.Scan(&image.Id,&image.DegreeId,&image.Image)
		if(err != nil){
			return images,err
		}
		images = append(images, image)
	}
	return images,err
}

func (dbRepo *PostgressDbRepo) DeleteDegreeImage(imageId int) (models.DegreeImage,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `delete from degrees_images where id = $1
				returning degree_id`
	row := db.QueryRowContext(ctx,query,imageId)
	image := models.DegreeImage{Id: imageId}
	err := row.Scan(&image.DegreeId)
	return image,err

}