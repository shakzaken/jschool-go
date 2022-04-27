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

func (ar *AppRepository) CreateDegreeImage(w http.ResponseWriter, r *http.Request) {
	degreeIdStr := chi.URLParam(r,"degreeId")
	degreeId,err := strconv.Atoi(degreeIdStr)	
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
	image := models.DegreeImage{
		Image: imageStr,
		DegreeId: degreeId,
	}
	image,err = ar.DB.CreateDegreeImage(image)
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

func (ar *AppRepository) GetDegreeImages(w http.ResponseWriter, r *http.Request) {
	degreeIdStr := chi.URLParam(r,"degreeId")
	degreeId,err := strconv.Atoi(degreeIdStr)	
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	images,err := ar.DB.GetDegreeImages(degreeId)
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

func (ar *AppRepository) DeleteDegreeImage(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r,"id")
	id,err := strconv.Atoi(idStr)	
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	image,err := ar.DB.DeleteDegreeImage(id)
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