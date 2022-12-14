// Code generated by goa v3.8.5, DO NOT EDIT.
//
// token service
//
// Command:
// $ goa gen github.com/sgerogia/hello-goa/design

package token

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// OAuth2 Authorisation service.
type Service interface {
	// Accepts username and password in the body and returns JWT OAUTH2/OIDC token
	// with the username as a subject, expiring in 1 hour.
	// The username and password are not verified, but cannot be empty strings.
	Auth(context.Context, *User) (res string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "token"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"auth"}

// User is the payload type of the token service auth method.
type User struct {
	// Username to access the service.
	Username *string
	// Password to access the service.
	Password *string
}

// MakeMalformedPayload builds a goa.ServiceError from an error.
func MakeMalformedPayload(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "MalformedPayload", false, false, false)
}
