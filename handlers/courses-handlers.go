package handlers

import (
	"encoding/json"
	"io/ioutil"
	"jschool/helpers"
	"jschool/models"
	"jschool/validations"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)


func (ar *AppRepository) GetAllCourses(w http.ResponseWriter, r *http.Request){

	courses,err := ar.DB.GetAllCourses()
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	out,err := json.Marshal(courses)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	w.Write(out)

}

func (ar *AppRepository) GetCourse(w http.ResponseWriter, r *http.Request){
	idStr := chi.URLParam(r,"id")
	id,err := strconv.Atoi(idStr)
	if(err !=nil){
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	course,err := ar.DB.GetCourse(id)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	out, err := json.Marshal(course);
	if(err != nil){
		helpers.ServerError(w,err);
		return
	}
	w.Write(out)

}


func (ar *AppRepository) CreateCourse(w http.ResponseWriter, r *http.Request){

	data,err := ioutil.ReadAll(r.Body)
	if(err !=nil){
		helpers.ServerError(w,err)
		return
	}

	course := models.Course{}
	err = json.Unmarshal(data,&course)
	if(err != nil){
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	if !validations.CreateCourse(course) {
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	createdCourse,err := ar.DB.CreateCourse(course)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	out,err := json.Marshal(createdCourse)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	w.Write(out)

}

func (ar *AppRepository) UpdateCourse(w http.ResponseWriter, r *http.Request){
	data,err := ioutil.ReadAll(r.Body)
	if(err !=nil){
		helpers.ServerError(w,err)
		return
	}

	course := models.Course{}
	err = json.Unmarshal(data,&course)
	if(err != nil){
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	if(!validations.UpdateCourse(course)){
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	updatedCourse,err := ar.DB.UpdateCourse(course)
	if(err != nil){
		helpers.ServerError(w,err)
		return;
	}
	out,err := json.Marshal(updatedCourse)
	if(err != nil){
		helpers.ServerError(w,err)
		return;
	}
	w.Write(out)
}

func (ar *AppRepository) DeleteCourse(w http.ResponseWriter , r *http.Request){
	idStr := chi.URLParam(r,"id")
	id,err := strconv.Atoi(idStr)
	if(err !=nil){
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	course,err := ar.DB.DeleteCourse(id)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	out,err := json.Marshal(course)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	w.Write(out)
}