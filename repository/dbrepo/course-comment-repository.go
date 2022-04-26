package dbrepo

import (
	"context"
	"jschool/models"
	"time"
)

func (dbRepo *PostgressDbRepo) CreateCourseComment(comment models.CourseComment) (models.CourseComment,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	query := `insert into courses_comments (course_id,user_id,comment,date)
				values($1,$2,$3,now()) returning id,date`
	row := db.QueryRowContext(ctx,query,comment.CourseId,comment.UserId,comment.Comment)
	err := row.Scan(&comment.Id,&comment.Date)
	return comment,err

}

func (dbRepo *PostgressDbRepo) GetCourseComments(courseId int) ([]models.CourseComment,error){
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	comments := []models.CourseComment{}
	query := `select id,course_id,user_id,comment,date from courses_comments
				where course_id = $1`
	rows,err := db.QueryContext(ctx,query,courseId)
	if(err != nil){
		return comments,err
	}
	for rows.Next() {
		comment := models.CourseComment{}
		err = rows.Scan(&comment.Id,&comment.CourseId,&comment.UserId,&comment.Comment,&comment.Date)
		if(err != nil){
			return comments,err
		}
		comments = append(comments, comment)
	}
	return comments,nil
}

func (dbRepo *PostgressDbRepo) DeleteCourseComment(commentId int) (models.CourseComment,error) {
	ctx, cancel := context.WithTimeout(context.Background(),3 *time.Second)
	defer cancel()
	db := dbRepo.DB

	comment := models.CourseComment{Id : commentId}
	query := `delete from courses_comments where id = $1
				returning course_id,user_id,comment,date`;
	row:= db.QueryRowContext(ctx,query,commentId)
	err := row.Scan(&comment.CourseId,&comment.UserId,&comment.Comment,&comment.Date)
	return comment,err	
}