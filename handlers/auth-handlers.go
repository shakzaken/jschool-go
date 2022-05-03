package handlers

import (
	"encoding/json"
	"io/ioutil"
	"jschool/config"
	"jschool/helpers"
	"jschool/models"
	"net/http"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserClaims struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}


func (ar *AppRepository) Login(w http.ResponseWriter,r *http.Request) {
	body,err := ioutil.ReadAll(r.Body)
	if(err !=nil){
		helpers.ServerError(w,err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body,&user)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}
	dbUser,err := ar.DB.GetUserByEmail(user.Email)
	if(err != nil){
		helpers.ServerError(w,err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password),[]byte(user.Password))
	if(err == bcrypt.ErrMismatchedHashAndPassword){
		if(err != nil){
			http.Error(w,"Password is Invalid",http.StatusBadRequest)
			return
		}
	}else if (err != nil){
		helpers.ServerError(w,err)
		return
	}

	token,err := generateToken(dbUser)
	if (err != nil){
		helpers.ServerError(w,err)
		return
	}
	w.Header().Add("token",token)
	w.Write([]byte("login successfuly"))
}

func (ar *AppRepository) CheckToken(w http.ResponseWriter,r *http.Request) {

	tokenStr := r.Header.Get("token")
	token,err := jwt.Parse(tokenStr,func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Secret),nil
	})
	if err != nil {
		helpers.ClientError(w,http.StatusBadRequest)
		return
	}
	if token.Valid {
		w.Write([]byte("token is valid"))
		return;
	}
	helpers.ClientError(w,http.StatusUnauthorized)
}

func generateToken(user models.User) (string,error) {
	
	key :=[]byte(config.Secret)
		
	
	claims := jwt.MapClaims{
		"id" : user.Id,
		"name": user.Name,
		"email" : user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenStr,err := token.SignedString(key)
	return tokenStr,err
}

