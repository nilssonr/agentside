package rest

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
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
	var body createCustomerRequest
	if err := render.Decode(r, &body); err != nil {
		handleError(w, r, err)
		return
	}

	request := &customer.Customer{
		FirstName:      body.FirstName,
		LastName:       body.LastName,
		TenantID:       tenantID(r),
		LastModifiedAt: time.Now(),
		LastModifiedBy: userID(r),
	}

	result, err := h.CustomerService.CreateCustomer(r.Context(), request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h customerHandler) getCustomers(w http.ResponseWriter, r *http.Request) {
	result, err := h.CustomerService.GetCustomers(r.Context(), tenantID(r))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h customerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	result, err := h.CustomerService.GetCustomer(r.Context(), tenantID(r), chi.URLParam(r, "customerID"))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h customerHandler) updateCustomer(w http.ResponseWriter, r *http.Request) {
	var body updateCustomerRequest
	if err := render.Decode(r, &body); err != nil {
		handleError(w, r, err)
		return
	}

	request := &customer.Customer{
		ID:             chi.URLParam(r, "customerID"),
		FirstName:      body.FirstName,
		LastName:       body.LastName,
		TenantID:       tenantID(r),
		LastModifiedAt: time.Now(),
		LastModifiedBy: userID(r),
	}

	result, err := h.CustomerService.UpdateCustomer(r.Context(), request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h customerHandler) deleteCustomer(w http.ResponseWriter, r *http.Request) {
	if err := h.CustomerService.DeleteCustomer(r.Context(), tenantID(r), chi.URLParam(r, "customerID")); err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}
