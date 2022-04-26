package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"jschool/helpers"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
)

func (ar *AppRepository) CreateCourseImage(w http.ResponseWriter, r *http.Request) {

	courseIdStr := chi.URLParam(r,"courseId")
	courseId,err := strconv.Atoi(courseIdStr)	
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	err = r.ParseMultipartForm(32 << 20)
	if(err != nil){
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	file,_,err := r.FormFile("files")
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	buffer := bytes.NewBuffer(nil)
	_,err = io.Copy(buffer,file)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	imageStr := base64.StdEncoding.EncodeToString(buffer.Bytes())

	image,err := ar.DB.CreateCourseImage(imageStr,courseId)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	out,err := json.Marshal(image)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	w.Write(out)

}

func (ar *AppRepository) GetCourseImages(w http.ResponseWriter, r *http.Request){
	courseIdStr := chi.URLParam(r,"courseId")
	courseId,err := strconv.Atoi(courseIdStr)	
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	images,err := ar.DB.GetCourseImages(courseId)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	out,err := json.Marshal(images)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	w.Write(out)
}

func (ar *AppRepository) DeleteCourseImage(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r,"id")
	id,err := strconv.Atoi(idStr)	
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	image,err := ar.DB.DeleteCourseImage(id)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	out,err := json.Marshal(image)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	w.Write(out)
}