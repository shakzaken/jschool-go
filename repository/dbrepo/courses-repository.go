package dbrepo

import (
	"jschool/models"
	"context"
	"time"
)

func (dbRepo *PostgressDbRepo) GetAllCourses() ([]models.Course,error){

	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `select id,name,description from courses`;


	courses := []models.Course{}
	rows,err := db.QueryContext(ctx,query)
	if(err != nil){
		return courses,err
	}

	for rows.Next() {
		course := models.Course{}
		err = rows.Scan(&course.Id,&course.Name,&course.Description)
		if(err != nil){
			return courses,err
		}
		courses = append(courses, course)
	}

	return courses,nil;
	
}

func (dbRepo *PostgressDbRepo) GetCourse(id int) (models.Course,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `select id,name,description from courses where id = $1`;

	row := db.QueryRowContext(ctx,query,id)
	course := models.Course{}
	err := row.Scan(&course.Id,&course.Name,&course.Description)
	return course, err
}

func (dbRepo *PostgressDbRepo) CreateCourse(course models.Course) (models.Course,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `insert into courses(name,description)
			 values ($1,$2) returning id`;
	row := db.QueryRowContext(ctx,query,course.Name,course.Description)
	err := row.Scan(&course.Id)
	return course,err
	
}

func (dbRepo *PostgressDbRepo) UpdateCourse(course models.Course) (models.Course,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `update courses 
			  set name = $1,description = $2
			  where id = $3`;
	_,err := db.ExecContext(ctx,query,course.Name,course.Description,course.Id)
	return course,err;
}

func (dbRepo *PostgressDbRepo) DeleteCourse(id int) (models.Course,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	course := models.Course{ Id: id }
	query := `delete from courses where id = $1 returning name,description`;
	row := db.QueryRowContext(ctx,query,id)
	err := row.Scan(&course.Name,&course.Description)
	return course,err;
	
}