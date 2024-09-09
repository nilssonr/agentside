package api

import (
	"github.com/nilssonr/agentside/customer"
	"github.com/nilssonr/agentside/interaction"
	"github.com/nilssonr/agentside/queue"
	"github.com/nilssonr/agentside/skill"
	"github.com/nilssonr/agentside/tenant"
	"github.com/nilssonr/agentside/user"
	"go.uber.org/zap"
)

type Options struct {
	Logger                      *zap.Logger
	AuthDomain                  string
	AuthAudience                string
	CustomerService             customer.Service
	CustomerAddressService      customer.AddressService
	CustomerEmailAddressService customer.EmailAddressService
	CustomerNoteService         customer.NoteService
	CustomerPhoneNumberService  customer.PhoneNumberService
	InteractionService          interaction.Service
	QueueService                queue.Service
	QueueSkillService           queue.SkillService
	SkillService                skill.Service
	TenantService               tenant.Service
	UserService                 user.Service
	UserSkillService            user.SkillService
}

func WithLogger(l *zap.Logger) func(*Options) {
	return func(o *Options) {
		o.Logger = l
	}
}

func WithAuth0Domain(d string) func(*Options) {
	return func(o *Options) {
		o.AuthDomain = d
	}
}

func WithAuth0Audience(a string) func(*Options) {
	return func(o *Options) {
		o.AuthAudience = a
	}
}

func WithCustomerService(cs customer.Service) func(*Options) {
	return func(o *Options) {
		o.CustomerService = cs
	}
}

func WithCustomerAddressService(cas customer.AddressService) func(*Options) {
	return func(o *Options) {
		o.CustomerAddressService = cas
	}
}

func WithCustomerEmailAddressService(ceas customer.EmailAddressService) func(*Options) {
	return func(o *Options) {
		o.CustomerEmailAddressService = ceas
	}
}

func WithCustomerNoteService(cns customer.NoteService) func(*Options) {
	return func(o *Options) {
		o.CustomerNoteService = cns
	}
}

func WithCustomerPhoneNumberService(cpns customer.PhoneNumberService) func(*Options) {
	return func(o *Options) {
		o.CustomerPhoneNumberService = cpns
	}
}

func WithInteractionService(is interaction.Service) func(*Options) {
	return func(o *Options) {
		o.InteractionService = is
	}
}

func WithQueueService(qs queue.Service) func(*Options) {
	return func(o *Options) {
		o.QueueService = qs
	}
}

func WithQueueSkillService(qss queue.SkillService) func(*Options) {
	return func(o *Options) {
		o.QueueSkillService = qss
	}
}

func WithSkillService(ss skill.Service) func(*Options) {
	return func(o *Options) {
		o.SkillService = ss
	}
}

func WithTenantService(ts tenant.Service) func(*Options) {
	return func(o *Options) {
		o.TenantService = ts
	}
}

func WithUserService(us user.Service) func(*Options) {
	return func(o *Options) {
		o.UserService = us
	}
}

func WithUserSkillservice(uss user.SkillService) func(*Options) {
	return func(o *Options) {
		o.UserSkillService = uss
	}
}
