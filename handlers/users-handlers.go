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




func (ar *AppRepository) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users,err := ar.DB.GetAllUsers()
	if(err != nil){
		helpers.ServerError(w,err);
		return;
	}

	out,err := json.Marshal(users);
	if(err != nil){
		helpers.ServerError(w,err);
		return;
	}

	w.Write(out)
}


func (ar *AppRepository) CreateUser(w http.ResponseWriter,r *http.Request){
	data,err := ioutil.ReadAll(r.Body);
	if(err!= nil){
		helpers.ServerError(w,err);
		return;
	}
	r.Body.Close()
	
	var user models.User
	err = json.Unmarshal(data,&user)
	if(err != nil){
		helpers.ServerError(w,err)
		return;
	}
	if(!validations.CreateUser(user)){
		helpers.ClientError(w,http.StatusBadRequest);
		return;
	}

	user,err = ar.DB.CreateUser(user)
	if(err != nil){
		helpers.ServerError(w,err);
		return;
	}
	
	out,err := json.Marshal(user)
	if(err != nil){
		helpers.ServerError(w,err);
		return;
	}
	w.Write(out)

}

func (ar *AppRepository) DeleteUser(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id");
	userId,err := strconv.Atoi(id);
	if(err != nil){
		helpers.ClientError(w,http.StatusBadRequest);
		return;
	}
	err = ar.DB.DeleteUser(userId);
	if(err != nil){
		helpers.ServerError(w,err)
		return;
	}
	w.Write([]byte("User deleted succesfuly"))
}

func (ar *AppRepository) UpdateUser(w http.ResponseWriter, r *http.Request){
	data,err := ioutil.ReadAll(r.Body);
	if(err!= nil){
		helpers.ServerError(w,err);
		return;
	}
	r.Body.Close()
	
	var user models.User
	err = json.Unmarshal(data,&user)
	if(err != nil){
		helpers.ServerError(w,err)
		return;
	}
	if(!validations.UpdateUser(user)){
		helpers.ClientError(w,http.StatusBadRequest);
		return;
	}

	err = ar.DB.UpdateUser(user)
	if(err !=nil) {
		helpers.ServerError(w,err);
		return;
	}

	w.Write([]byte("User updated succesfuly"))

}