package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/nilssonr/agentside/api/rest/middleware"
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
	r.Use(middleware.Auth0(o.AuthDomain, o.AuthAudience))

	// Register handlers
	handlers := []handler{
		authClientHandler{
			AuthClientService: o.AuthClientService,
		},
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
		queueHandler{
			QueueService:      o.QueueService,
			QueueSkillService: o.QueueSkillService,
		},
		skillHandler{
			SkillService: o.SkillService,
		},
		tenantHandler{
			TenantService: o.TenantService,
		},
		userHandler{
			UserService:      o.UserService,
			UserSkillService: o.UserSkillService,
		},
	}

	r.Route("/api", func(r chi.Router) {
		for _, v := range handlers {
			v.Register(r)
		}
	})

	return r
}
