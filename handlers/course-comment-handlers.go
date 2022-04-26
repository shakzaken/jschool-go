package handlers

import (
	"encoding/json"
	"io/ioutil"
	"jschool/helpers"
	"jschool/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (ar *AppRepository) CreateCourseComment(w http.ResponseWriter, r *http.Request){
	body,err := ioutil.ReadAll(r.Body) 
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	comment := models.CourseComment{}
	err = json.Unmarshal(body,&comment)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	comment,err = ar.DB.CreateCourseComment(comment)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	out,err := json.Marshal(comment)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	w.Write(out)	
}

func (ar * AppRepository) GetCourseComments(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r,"id")
	id,err := strconv.Atoi(idStr)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	comments,err := ar.DB.GetCourseComments(id)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	out,err := json.Marshal(comments)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	w.Write(out)
}

func (ar *AppRepository) DeleteCourseComment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r,"id")
	id,err := strconv.Atoi(idStr)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	comment,err := ar.DB.DeleteCourseComment(id)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	out,err := json.Marshal(comment)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	w.Write(out)

}