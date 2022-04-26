package dbrepo

import (
	"context"
	"jschool/models"
	"time"
)

func (dbRepo *PostgressDbRepo) CreateCourseImage(imageBody string, courseId int) (models.CourseImage,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := "insert into courses_images (image,course_id) values($1,$2) returning id"

	image := models.CourseImage{CourseId: courseId}
	row := db.QueryRowContext(ctx,query,imageBody,courseId)
	err := row.Scan(&image.Id)
	return image,err
}

func (dbRepo *PostgressDbRepo) GetCourseImages(courseId int) ([]models.CourseImage,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `select id,course_id,image from courses_images where course_id = $1`
	images := []models.CourseImage{}
	rows,err := db.QueryContext(ctx,query,courseId)
	if(err != nil){
		return images,err
	}

	for rows.Next() {
		image := models.CourseImage{}
		err = rows.Scan(&image.Id,&image.CourseId,&image.Image)
		if(err != nil){
			return images,err
		}
		images = append(images, image)
	}
	return images,nil
}

func (dbRepo *PostgressDbRepo) DeleteCourseImage(imageId int) (models.CourseImage,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `delete from courses_images where id = $1 returning course_id`

	row := db.QueryRowContext(ctx,query,imageId)
	image := models.CourseImage{Id: imageId}
	err := row.Scan(&image.CourseId)
	return image,err
}