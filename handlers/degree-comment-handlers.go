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

func (ar *AppRepository) CreateDegreeComment(w http.ResponseWriter,r *http.Request){
	body,err := ioutil.ReadAll(r.Body)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	comment := models.DegreeComment{}
	err = json.Unmarshal(body,&comment)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	comment,err = ar.DB.CreateDegreeComment(comment)
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

func (ar *AppRepository) GetDegreeComments(w http.ResponseWriter, r *http.Request){
	degreeIdStr := chi.URLParam(r,"degreeId")
	degreeId,err := strconv.Atoi(degreeIdStr)
	if(err != nil){
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	comments,err := ar.DB.GetDegreeComments(degreeId)
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

func (ar *AppRepository) DeleteDegreeComment(w http.ResponseWriter,r *http.Request) {
	idStr := chi.URLParam(r,"id")
	id,err := strconv.Atoi(idStr)
	if(err != nil){
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	comment,err := ar.DB.DeleteDegreeComment(id)
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