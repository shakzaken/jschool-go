package handlers

import (
	"encoding/json"
	"io/ioutil"
	"jschool/repository/models"
	"log"
	"net/http"
)




func (ar *AppRepository) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users,err := ar.DB.GetAllUsers()
	if(err != nil){
		log.Println(err)
		return;
	}

	out,err := json.Marshal(users);
	if(err != nil){
		log.Println("Error while parsing user");
		return;
	}

	w.Header().Set("Content-Type","application/json")
	w.Write(out)
}

func (ar *AppRepository) CreateUser(w http.ResponseWriter,r *http.Request){
	data,err := ioutil.ReadAll(r.Body);
	if(err!= nil){
		log.Println("error while creating user")
		log.Println(err)
		return;
	}

	r.Body.Close()
	
	var user models.User
	err = json.Unmarshal(data,&user)
	if(err != nil){
		log.Println("error parsing user")
		return;
	}
	
	_,err = ar.DB.CreateUser(user)
	if(err != nil){
		log.Println("error creating user");
		log.Println(err)
		return;
	}
	
	w.Write([]byte("User created succesfuly"))

	
}