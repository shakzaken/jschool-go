package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"jschool/helpers"
	"jschool/models"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
)

func (ar *AppRepository) CreateUserImage(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r,"userId")
	userId,err := strconv.Atoi(userIdStr)	
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
	image := models.UserImage {
		UserId: userId,
		Image: imageStr,
	}
	image,err = ar.DB.CreateUserImage(image)
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

func (ar *AppRepository) GetUserImages(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r,"userId")
	userId,err := strconv.Atoi(userIdStr)	
	if(err != nil){
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	images,err := ar.DB.GetUserImages(userId)
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

func (ar *AppRepository) DeleteUserImage(w http.ResponseWriter,r *http.Request){
	idStr := chi.URLParam(r,"id")
	id,err := strconv.Atoi(idStr)	
	if(err != nil){
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	image,err := ar.DB.DeleteUserImage(id)
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