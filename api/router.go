package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/nilssonr/agentside/api/middleware"
)

type handler interface {
	Register(r chi.Router)
}

func NewRouter(opts ...func(*Options)) *chi.Mux {
	o := &Options{}
	for _, v := range opts {
		v(o)
	}

	r := chi.NewRouter()

	// Register middlewares
	r.Use(middleware.RequestLogger(o.Logger))

	// Register handlers
	handlers := []handler{
		customerHandler{
			CustomerService:             o.CustomerService,
			CustomerAddressService:      o.CustomerAddressService,
			CustomerEmailAddressService: o.CustomerEmailAddressService,
			CustomerNoteService:         o.CustomerNoteService,
			CustomerPhoneNumberService:  o.CustomerPhoneNumberService,
		},
		interactionHandler{
			InteractionService: o.InteractionService,
		},
	}

	r.Route("/api", func(r chi.Router) {
		for _, v := range handlers {
			v.Register(r)
		}
	})

	registerQueueRouter(r)
	registerSkillRouter(r)
	registerTenantRouter(r)
	registerUserRouter(r)

	return r
}
