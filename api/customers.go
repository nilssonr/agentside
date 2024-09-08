package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nilssonr/agentside/customer"
)

type customerHandler struct {
	CustomerService             customer.Service
	CustomerAddressService      customer.AddressService
	CustomerEmailAddressService customer.EmailAddressService
	CustomerNoteService         customer.NoteService
	CustomerPhoneNumberService  customer.PhoneNumberService
}

func (h customerHandler) Register(r chi.Router) {
	r.Route("/customers", func(r chi.Router) {
		r.Post("/", h.createCustomer)
		r.Get("/", h.getCustomers)
		r.Route("/{customerID}", func(r chi.Router) {
			r.Get("/", h.getCustomer)
			r.Put("/", h.updateCustomer)
			r.Delete("/", h.deleteCustomer)
		})
	})
}

type createCustomerRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type updateCustomerRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (h customerHandler) createCustomer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h customerHandler) getCustomers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h customerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {

}

func (h customerHandler) updateCustomer(w http.ResponseWriter, r *http.Request) {

}

func (h customerHandler) deleteCustomer(w http.ResponseWriter, r *http.Request) {

}
