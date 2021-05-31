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

// RegisterUser implements RegisterUser.
func (s *adminsrvc) RegisterUser(ctx context.Context, p *admin.RegisterUserPayload) (res *admin.RegisterUserResult, err error) {
	res = &admin.RegisterUserResult{}
	s.logger.Print("admin.RegisterUser")
	return
}

// GetUser implements GetUser.
func (s *adminsrvc) GetUser(ctx context.Context, p *admin.GetUserPayload) (res *admin.GetUserResult, err error) {
	res = &admin.GetUserResult{}
	s.logger.Print("admin.GetUser")
	return
}
