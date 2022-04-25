package dbrepo

import (
	"jschool/models"
	"context"
	"time"
)


func (dbRepo *PostgressDbRepo)  GetAllDegrees() ([]models.Degree,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `select id,name,description from degrees`

	degrees := []models.Degree{}
	rows,err := db.QueryContext(ctx,query)
	if(err != nil){
		return degrees,nil
	}

	for rows.Next() {
		degree := models.Degree{}
		err = rows.Scan(&degree.Id,&degree.Name,&degree.Description)
		if(err != nil){
			return degrees,err
		}
		degrees = append(degrees, degree)
	}
	return degrees,nil
}

func (dbRepo *PostgressDbRepo) GetDegree(id int) (models.Degree,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `select id,name,description from degrees where id = $1`;
	degree := models.Degree{}
	row := db.QueryRowContext(ctx,query,id)
	err := row.Scan(&degree.Id,&degree.Name,&degree.Description)
	return degree,err;

}

func (dbRepo *PostgressDbRepo) CreateDegree(degree models.Degree) (models.Degree,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `insert into degrees (name,description) values($1,$2) returning id`
	row := db.QueryRowContext(ctx,query,degree.Name,degree.Description)
	err := row.Scan(&degree.Id)
	return degree,err;
}

func (dbRepo *PostgressDbRepo) UpdateDegree(degree models.Degree) (models.Degree,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `update degrees set name = $1, description = $2 where id = $3`;
	_,err := db.ExecContext(ctx,query,degree.Name,degree.Description,degree.Id)
	return degree,err;
}

func (dbRepo *PostgressDbRepo) DeleteDegree(id int) (models.Degree,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	degree := models.Degree{ Id: id }
	query := `delete from degrees where id = $1 returning name,description`;
	row := db.QueryRowContext(ctx,query,id)
	err := row.Scan(&degree.Name,&degree.Description)
	return degree,err

}