// Code generated by goa v3.8.5, DO NOT EDIT.
//
// token client
//
// Command:
// $ goa gen github.com/sgerogia/hello-goa/design

package token

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "token" service client.
type Client struct {
	AuthEndpoint goa.Endpoint
}

// NewClient initializes a "token" service client given the endpoints.
func NewClient(auth goa.Endpoint) *Client {
	return &Client{
		AuthEndpoint: auth,
	}
}

// Auth calls the "auth" endpoint of the "token" service.
// Auth may return the following errors:
//   - "MalformedPayload" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) Auth(ctx context.Context, p *User) (res string, err error) {
	var ires interface{}
	ires, err = c.AuthEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}