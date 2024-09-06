package customer

import "context"

type Repository interface {
	InsertCustomer(ctx context.Context, request *Customer) (*Customer, error)
	GetCustomers(ctx context.Context, tenantID string) ([]*Customer, error)
	GetCustomer(ctx context.Context, tenantID, customerID string) (*Customer, error)
	UpdateCustomer(ctx context.Context, request *Customer) (*Customer, error)
	DeleteCustomer(ctx context.Context, tenantID, customerID string) error
}

type PhoneNumberRepository interface {
	InsertPhoneNumber(ctx context.Context, request *PhoneNumber) (*PhoneNumber, error)
	GetPhoneNumbers(ctx context.Context, customerID string) ([]*PhoneNumber, error)
	GetPhoneNumber(ctx context.Context, customerID, phoneNumberID string) (*PhoneNumber, error)
	UpdatePhoneNumber(ctx context.Context, request *PhoneNumber) (*PhoneNumber, error)
	DeletePhoneNumber(ctx context.Context, customerID, phoneNumberID string) error
}

type EmailAddressRepository interface {
	InsertEmailAddress(ctx context.Context, request *EmailAddress) (*EmailAddress, error)
	GetEmailAddresses(ctx context.Context, customerID string) ([]*EmailAddress, error)
	GetEmailAddress(ctx context.Context, customerID, emailAddressID string) (*EmailAddress, error)
	UpdateEmailAddress(ctx context.Context, request *EmailAddress) (*EmailAddress, error)
	DeleteEmailAddress(ctx context.Context, customerID, emailAddressID string) error
}

type AddressRepository interface {
	InsertAddress(ctx context.Context, request *Address) (*Address, error)
	GetAddresses(ctx context.Context, customerID string) ([]*Address, error)
	GetAddress(ctx context.Context, customerID, addressID string) (*Address, error)
	UpdateAddress(ctx context.Context, request *Address) (*Address, error)
	DeleteAddress(ctx context.Context, customerID, addressID string) error
}

type NoteRepository interface {
	InsertNote(ctx context.Context, request *Note) (*Note, error)
	GetNotes(ctx context.Context, customerID string) ([]*Note, error)
	GetNote(ctx context.Context, customerID, noteID string) (*Note, error)
	UpdateNote(ctx context.Context, request *Note) (*Note, error)
	DeleteNote(ctx context.Context, customerID, noteID string) error
}
