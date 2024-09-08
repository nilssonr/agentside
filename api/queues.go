package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func registerQueueRouter(r chi.Router) {
	r.Route("/queues", func(r chi.Router) {
		r.Post("/", createQueue)
		r.Get("/", getQueues)
		r.Route("/{queueID}", func(r chi.Router) {
			r.Get("/", getQueue)
			r.Put("/", updateQueue)
			r.Delete("/", deleteQueue)
		})
	})
}

type createQueueRequest struct {
	Name string `json:"name"`
}

type updateQueueRequest struct {
	Name string `json:"name"`
}

func createQueue(w http.ResponseWriter, r *http.Request) {

}

func getQueues(w http.ResponseWriter, r *http.Request) {

}

func getQueue(w http.ResponseWriter, r *http.Request) {

}

func updateQueue(w http.ResponseWriter, r *http.Request) {

}

func deleteQueue(w http.ResponseWriter, r *http.Request) {

}
