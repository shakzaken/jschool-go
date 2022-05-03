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

	r.Post("/login",repo.Login)
	r.Post("/auth",repo.CheckToken)

	r.Get("/users/images/{userId}",repo.GetUserImages)
	r.Post("/users/images/{userId}",repo.CreateUserImage)
	r.Delete("/users/images/{id}",repo.DeleteUserImage)

	r.Get("/courses",repo.GetAllCourses)
	r.Get("/courses/{id}",repo.GetCourse)
	r.Post("/courses",repo.CreateCourse)
	r.Put("/courses",repo.UpdateCourse)
	r.Delete("/courses/{id}",repo.DeleteCourse)

	r.Get("/courses/comments/{id}",repo.GetCourseComments)
	r.Post("/courses/comments",repo.CreateCourseComment)
	r.Delete("/courses/comments/{id}",repo.DeleteCourseComment)

	r.Get("/courses/images/{courseId}",repo.GetCourseImages)
	r.Post("/courses/images/{courseId}",repo.CreateCourseImage)
	r.Delete("/courses/images/{id}",repo.DeleteCourseImage)

	r.Post("/courses/images",repo.CreateCourseImage)

	r.Get("/degrees",repo.GetAllDegrees)
	r.Get("/degrees/{id}",repo.GetDegree)
	r.Post("/degrees",repo.CreateDegree)
	r.Put("/degrees",repo.UpdateDegree)
	r.Delete("/degrees/{id}",repo.DeleteDegree)

	r.Get("/degrees/comments/{degreeId}",repo.GetDegreeComments)
	r.Post("/degrees/comments",repo.CreateDegreeComment)
	r.Delete("/degrees/comments/{id}",repo.DeleteDegreeComment)

	r.Get("/degrees/images/{degreeId}",repo.GetDegreeImages)
	r.Post("/degrees/images/{degreeId}",repo.CreateDegreeImage)
	r.Delete("/degrees/images/{id}",repo.DeleteDegreeImage)



	return r;
}