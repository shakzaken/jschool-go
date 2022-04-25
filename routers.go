package main

import (
	"jschool/handlers"
	"jschool/middlewares"
	"github.com/go-chi/chi/v5"
)

func Routers(repo *handlers.AppRepository) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middlewares.JsonMiddleware)


	r.Get("/users", repo.GetAllUsers)
	r.Post("/users",repo.CreateUser)
	r.Delete("/users/{id}",repo.DeleteUser)
	r.Put("/users",repo.UpdateUser)

	r.Get("/courses",repo.GetAllCourses)
	r.Get("/courses/{id}",repo.GetCourse)
	r.Post("/courses",repo.CreateCourse)
	r.Put("/courses",repo.UpdateCourse)
	r.Delete("/courses/{id}",repo.DeleteCourse)


	r.Get("/degrees",repo.GetAllDegrees)
	r.Get("/degrees/{id}",repo.GetDegree)
	r.Post("/degrees",repo.CreateDegree)
	r.Put("/degrees",repo.UpdateDegree)
	r.Delete("/degrees/{id}",repo.DeleteDegree)

	return r;
}