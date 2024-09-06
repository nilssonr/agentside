package v1

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nilssonr/agentside/customer"
)

// CreateCustomer implements ServerInterface.
func (ah AgentsideHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
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

	result, err := ah.CustomerService.CreateCustomer(r.Context(), &request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// GetCustomers implements ServerInterface.
func (ah AgentsideHandler) GetCustomers(w http.ResponseWriter, r *http.Request, params GetCustomersParams) {
	result, err := ah.CustomerService.GetCustomers(r.Context(), tenantID(r))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// GetCustomer implements ServerInterface.
func (ah AgentsideHandler) GetCustomer(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	result, err := ah.CustomerService.GetCustomer(r.Context(), tenantID(r), customerId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// UpdateCustomer implements ServerInterface.
func (ah AgentsideHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
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

	result, err := ah.CustomerService.UpdateCustomer(r.Context(), &request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// DeleteCustomer implements ServerInterface.
func (ah AgentsideHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	if err := ah.CustomerService.DeleteCustomer(r.Context(), tenantID(r), customerId.String()); err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}

// CreateCustomerNote implements ServerInterface.
func (ah AgentsideHandler) CreateCustomerNote(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
}

// GetCustomerNotes implements ServerInterface.
func (ah AgentsideHandler) GetCustomerNotes(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	panic("unimplemented")
}

// GetCustomerNote implements ServerInterface.
func (ah AgentsideHandler) GetCustomerNote(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, noteId uuid.UUID) {
	panic("unimplemented")
}

// UpdateCustomerNote implements ServerInterface.
func (ah AgentsideHandler) UpdateCustomerNote(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, noteId uuid.UUID) {
	panic("unimplemented")
}

// DeleteCustomerNote implements ServerInterface.
func (ah AgentsideHandler) DeleteCustomerNote(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, noteId uuid.UUID) {
	panic("unimplemented")
}

// CreateCustomerEmailAddress implements ServerInterface.
func (ah AgentsideHandler) CreateCustomerEmailAddress(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	panic("unimplemented")
}

// GetCustomerEmailAddresses implements ServerInterface.
func (ah AgentsideHandler) GetCustomerEmailAddresses(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	panic("unimplemented")
}

// GetCustomerEmailAddress implements ServerInterface.
func (ah AgentsideHandler) GetCustomerEmailAddress(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, emailAddressId uuid.UUID) {
	panic("unimplemented")
}

// UpdateCustomerEmailAddress implements ServerInterface.
func (ah AgentsideHandler) UpdateCustomerEmailAddress(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, emailAddressId uuid.UUID) {
	panic("unimplemented")
}

// DeleteCustomerEmailAddress implements ServerInterface.
func (ah AgentsideHandler) DeleteCustomerEmailAddress(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, emailAddressId uuid.UUID) {
	panic("unimplemented")
}

// CreateCustomerPhoneNumber implements ServerInterface.
func (ah AgentsideHandler) CreateCustomerPhoneNumber(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
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

	result, err := ah.CustomerPhoneNumberService.CreatePhoneNumber(r.Context(), &request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// GetCustomerPhoneNumbers implements ServerInterface.
func (ah AgentsideHandler) GetCustomerPhoneNumbers(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	result, err := ah.CustomerPhoneNumberService.GetPhoneNumbers(r.Context(), customerId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// GetCustomerPhoneNumber implements ServerInterface.
func (ah AgentsideHandler) GetCustomerPhoneNumber(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, phoneNumberId uuid.UUID) {
	result, err := ah.CustomerPhoneNumberService.GetPhoneNumber(r.Context(), customerId.String(), phoneNumberId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// UpdateCustomerPhoneNumber implements ServerInterface.
func (ah AgentsideHandler) UpdateCustomerPhoneNumber(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, phoneNumberId uuid.UUID) {
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

	result, err := ah.CustomerPhoneNumberService.UpdatePhoneNumber(r.Context(), &request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// DeleteCustomerPhoneNumber implements ServerInterface.
func (ah AgentsideHandler) DeleteCustomerPhoneNumber(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, phoneNumberId uuid.UUID) {
	if err := ah.CustomerPhoneNumberService.DeletePhoneNumber(r.Context(), customerId.String(), phoneNumberId.String()); err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}
