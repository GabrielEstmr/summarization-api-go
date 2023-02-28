package main_gateways

import (
	"context"
	main_domains "mpindicator/main/domains"
)

type UserDatabaseGateway interface {
	Save(context.Context, main_domains.User) (string, error)
}
