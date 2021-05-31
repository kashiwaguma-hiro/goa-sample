package adminapi

import (
	"context"
	"log"

	admin "github.com/kashiwaguma-hiro/goa-sample/gen/admin"
)

// admin service example implementation.
// The example methods log the requests and return zero values.
type adminsrvc struct {
	logger *log.Logger
}

// NewAdmin returns the admin service implementation.
func NewAdmin(logger *log.Logger) admin.Service {
	return &adminsrvc{logger}
}

// Users implements users.
func (s *adminsrvc) Users(ctx context.Context, p *admin.UsersPayload) (res *admin.UsersResult, err error) {
	res = &admin.UsersResult{}
	s.logger.Print("admin.users")
	return
}
