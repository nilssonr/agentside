package v1

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nilssonr/agentside/customer"
)

type CustomerHandler struct {
	CustomerService             customer.Service
	CustomerPhoneNumberService  customer.PhoneNumberService
	CustomerEmailAddressService customer.EmailAddressService
	CustomerAddressService      customer.AddressService
	CustomerNoteService         customer.NoteService
}

func NewCustomerHandler(
	cs customer.Service,
	cpns customer.PhoneNumberService,
	ceas customer.EmailAddressService,
	cas customer.AddressService,
	cns customer.NoteService) CustomerHandler {
	return CustomerHandler{
		CustomerService:             cs,
		CustomerPhoneNumberService:  cpns,
		CustomerEmailAddressService: ceas,
		CustomerAddressService:      cas,
		CustomerNoteService:         cns,
	}
}

// CreateCustomer implements ServerInterface.
func (h CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var body CreateCustomerRequest
	if err := render.DecodeJSON(r.Body, &body); err != nil {
		handleError(w, r, err)
		return
	}
	defer r.Body.Close()

	var request customer.Customer
	request.FirstName = body.FirstName
	request.LastName = body.LastName
	request.TenantID = tenantID(r)
	request.LastModifiedAt = time.Now()
	request.LastModifiedBy = userID(r)

	result, err := h.CustomerService.CreateCustomer(r.Context(), &request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// GetCustomers implements ServerInterface.
func (h CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request, params GetCustomersParams) {
	result, err := h.CustomerService.GetCustomers(r.Context(), tenantID(r))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// GetCustomer implements ServerInterface.
func (h CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	result, err := h.CustomerService.GetCustomer(r.Context(), tenantID(r), customerId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// UpdateCustomer implements ServerInterface.
func (h CustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	var body UpdateCustomerRequest
	if err := render.DecodeJSON(r.Body, &body); err != nil {
		handleError(w, r, err)
		return
	}
	defer r.Body.Close()

	var request customer.Customer
	request.ID = customerId.String()
	request.FirstName = body.FirstName
	request.LastName = body.LastName
	request.TenantID = tenantID(r)
	request.LastModifiedBy = userID(r)
	request.LastModifiedAt = time.Now()

	result, err := h.CustomerService.UpdateCustomer(r.Context(), &request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// DeleteCustomer implements ServerInterface.
func (h CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	if err := h.CustomerService.DeleteCustomer(r.Context(), tenantID(r), customerId.String()); err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}

// CreateCustomerNote implements ServerInterface.
func (h CustomerHandler) CreateCustomerNote(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
}

// GetCustomerNotes implements ServerInterface.
func (h CustomerHandler) GetCustomerNotes(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	panic("unimplemented")
}

// GetCustomerNote implements ServerInterface.
func (h CustomerHandler) GetCustomerNote(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, noteId uuid.UUID) {
	panic("unimplemented")
}

// UpdateCustomerNote implements ServerInterface.
func (h CustomerHandler) UpdateCustomerNote(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, noteId uuid.UUID) {
	panic("unimplemented")
}

// DeleteCustomerNote implements ServerInterface.
func (h CustomerHandler) DeleteCustomerNote(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, noteId uuid.UUID) {
	panic("unimplemented")
}

// CreateCustomerEmailAddress implements ServerInterface.
func (h CustomerHandler) CreateCustomerEmailAddress(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	panic("unimplemented")
}

// GetCustomerEmailAddresses implements ServerInterface.
func (h CustomerHandler) GetCustomerEmailAddresses(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	panic("unimplemented")
}

// GetCustomerEmailAddress implements ServerInterface.
func (h CustomerHandler) GetCustomerEmailAddress(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, emailAddressId uuid.UUID) {
	panic("unimplemented")
}

// UpdateCustomerEmailAddress implements ServerInterface.
func (h CustomerHandler) UpdateCustomerEmailAddress(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, emailAddressId uuid.UUID) {
	panic("unimplemented")
}

// DeleteCustomerEmailAddress implements ServerInterface.
func (h CustomerHandler) DeleteCustomerEmailAddress(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, emailAddressId uuid.UUID) {
	panic("unimplemented")
}

// CreateCustomerPhoneNumber implements ServerInterface.
func (h CustomerHandler) CreateCustomerPhoneNumber(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	var body CreateCustomerPhoneNumberRequest
	if err := render.DecodeJSON(r.Body, &body); err != nil {
		handleError(w, r, err)
		return
	}
	defer r.Body.Close()

	var request customer.PhoneNumber
	request.CustomerID = customerId.String()
	request.PhoneNumber = body.PhoneNumber
	request.Type = body.Type
	request.LastModifiedAt = time.Now()
	request.LastModifiedBy = userID(r)

	result, err := h.CustomerPhoneNumberService.CreatePhoneNumber(r.Context(), &request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// GetCustomerPhoneNumbers implements ServerInterface.
func (h CustomerHandler) GetCustomerPhoneNumbers(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	result, err := h.CustomerPhoneNumberService.GetPhoneNumbers(r.Context(), customerId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// GetCustomerPhoneNumber implements ServerInterface.
func (h CustomerHandler) GetCustomerPhoneNumber(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, phoneNumberId uuid.UUID) {
	result, err := h.CustomerPhoneNumberService.GetPhoneNumber(r.Context(), customerId.String(), phoneNumberId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// UpdateCustomerPhoneNumber implements ServerInterface.
func (h CustomerHandler) UpdateCustomerPhoneNumber(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, phoneNumberId uuid.UUID) {
	var body UpdateCustomerPhoneNumberRequest
	if err := render.DecodeJSON(r.Body, &body); err != nil {
		handleError(w, r, err)
		return
	}
	defer r.Body.Close()

	var request customer.PhoneNumber
	request.ID = phoneNumberId.String()
	request.PhoneNumber = body.PhoneNumber
	request.Type = body.Type
	request.CustomerID = customerId.String()
	request.LastModifiedAt = time.Now()
	request.LastModifiedBy = userID(r)

	result, err := h.CustomerPhoneNumberService.UpdatePhoneNumber(r.Context(), &request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// DeleteCustomerPhoneNumber implements ServerInterface.
func (h CustomerHandler) DeleteCustomerPhoneNumber(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, phoneNumberId uuid.UUID) {
	if err := h.CustomerPhoneNumberService.DeletePhoneNumber(r.Context(), customerId.String(), phoneNumberId.String()); err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}

func (h CustomerHandler) CreateCustomerAddress(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	var body CreateCustomerAddressRequest
	if err := render.DecodeJSON(r.Body, &body); err != nil {
		handleError(w, r, err)
		return
	}

	request := &customer.Address{
		Type:           body.Type,
		StreetAddress:  body.StreetAddress,
		CustomerID:     customerId.String(),
		LastModifiedAt: time.Now(),
		LastModifiedBy: userID(r),
	}

	if body.Country != nil {
		request.Country = *body.Country
	}

	if body.State != nil {
		request.State = *body.State
	}

	if body.ZipCode != nil {
		request.ZipCode = *body.ZipCode
	}

	result, err := h.CustomerAddressService.CreateAddress(r.Context(), request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h CustomerHandler) GetCustomerAddresses(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {

}

func (h CustomerHandler) GetCustomerAddress(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, addressId uuid.UUID) {

}

func (h CustomerHandler) UpdateCustomerAddress(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, addressId uuid.UUID) {

}

func (h CustomerHandler) DeleteCustomerAddress(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, addressId uuid.UUID) {

}
