package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func registerUserRouter(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", createUser)
		r.Get("/", getUsers)
		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", getUser)
			r.Put("/", updateUser)
			r.Delete("/", deleteUser)
		})
	})
}

type createUserRequest struct {
	Name string `json:"name"`
}

type updateUserRequest struct {
	Name string `json:"name"`
}

func createUser(w http.ResponseWriter, r *http.Request) {

}

func getUsers(w http.ResponseWriter, r *http.Request) {

}

func getUser(w http.ResponseWriter, r *http.Request) {

}

func updateUser(w http.ResponseWriter, r *http.Request) {

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

}
