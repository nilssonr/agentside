package customer

import "go.uber.org/fx"

var Module = fx.Module("customer",
	fx.Provide(NewService),
	fx.Provide(NewPhoneNumberService),
	fx.Provide(NewEmailAddressService),
	fx.Provide(NewAddressService),
	fx.Provide(NewNoteService),
)
