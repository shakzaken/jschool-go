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


func (ar *AppRepository) GetAllDegrees(w http.ResponseWriter, r *http.Request) {
	degrees,err := ar.DB.GetAllDegrees()
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	out,err := json.Marshal(degrees)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	w.Write(out)
}

func (ar *AppRepository) GetDegree(w http.ResponseWriter, r * http.Request) {
	idStr := chi.URLParam(r,"id")
	id,err := strconv.Atoi(idStr)
	if(err != nil){
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	degree,err := ar.DB.GetDegree(id)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	out,err := json.Marshal(degree)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	w.Write(out)
}

func (ar *AppRepository) CreateDegree(w http.ResponseWriter, r *http.Request) {

	body,err := ioutil.ReadAll(r.Body)
	if(err != nil){
		helpers.ServerError(w,err)
		return;
	}
	degree := models.Degree{}
	err = json.Unmarshal(body,&degree)
	if(err != nil){
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	if !validations.CreateDegree(degree) {
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	degree,err = ar.DB.CreateDegree(degree)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	out,err := json.Marshal(degree)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	w.Write(out)
}

func (ar *AppRepository) UpdateDegree(w http.ResponseWriter, r *http.Request){
	body,err := ioutil.ReadAll(r.Body)
	if(err != nil){
		helpers.ServerError(w,err)
		return;
	}
	degree := models.Degree{}
	err = json.Unmarshal(body,&degree)
	if(err != nil){
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	if !validations.UpdateDegree(degree) {
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	degree,err = ar.DB.UpdateDegree(degree)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	out,err := json.Marshal(degree)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	w.Write(out)
}

func (ar *AppRepository) DeleteDegree(w http.ResponseWriter, r *http.Request){
	idStr := chi.URLParam(r,"id")
	id,err := strconv.Atoi(idStr)
	if(err != nil){
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	degree,err := ar.DB.DeleteDegree(id)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	out,err := json.Marshal(degree)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	w.Write(out)
	
}