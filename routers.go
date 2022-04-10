package main

import (
	"jschool/handlers"
	"github.com/go-chi/chi/v5"
)

func Routers(repo *handlers.AppRepository) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/users", repo.GetAllUsers)
	r.Post("/users",repo.CreateUser)
	
	

	return r;
}