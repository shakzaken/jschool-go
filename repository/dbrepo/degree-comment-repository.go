package dbrepo

import (
	"context"
	"jschool/models"
	"time"
)



func (dbRepo *PostgressDbRepo) CreateDegreeComment(comment models.DegreeComment) (models.DegreeComment, error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `insert into degrees_comments (degree_id,user_id,comment,date)
				values($1,$2,$3,now()) returning id`
	
	row := db.QueryRowContext(ctx,query,comment.DegreeId,comment.UserId,comment.Comment)
	err := row.Scan(&comment.Id)
	return comment,err
}

func (dbRepo *PostgressDbRepo) GetDegreeComments(degreeId int) ([]models.DegreeComment,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `select id,degree_id,user_id,comment,date from degrees_comments 
				where degree_id = $1`
	comments := []models.DegreeComment{}
	rows,err := db.QueryContext(ctx,query,degreeId)
	if(err != nil){
		return comments,nil
	}
	for rows.Next() {
		comment := models.DegreeComment{}
		err = rows.Scan(&comment.Id,&comment.DegreeId,&comment.UserId,&comment.Comment,&comment.Date)
		if(err != nil){
			return comments,err
		}
		comments = append(comments, comment)
	}
	return comments,nil
}

func (dbRepo *PostgressDbRepo) DeleteDegreeComment(commentId int) (models.DegreeComment,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `delete from degrees_comments where id = $1 
				returning degree_id,user_id,comment,date`
	row := db.QueryRowContext(ctx,query,commentId)
	comment := models.DegreeComment{Id: commentId}
	err := row.Scan(&comment.DegreeId,&comment.UserId,&comment.Comment,&comment.Date)
	return comment,err
}