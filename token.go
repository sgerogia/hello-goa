package mathapi

import (
	"context"
	"log"

	token "github.com/sgerogia/hello-goa/gen/token"
)

// token service example implementation.
// The example methods log the requests and return zero values.
type tokensrvc struct {
	logger *log.Logger
}

// NewToken returns the token service implementation.
func NewToken(logger *log.Logger) token.Service {
	return &tokensrvc{logger}
}

// Accepts username and password in the body and returns JWT OAUTH2/OIDC token
// with the username as a subject, expiring in 1 hour.
// The username and password are not verified, but cannot be empty strings.
func (s *tokensrvc) Auth(ctx context.Context, p *token.User) (res string, err error) {
	s.logger.Print("token.auth")
	return
}
