// Code generated by goa v3.8.5, DO NOT EDIT.
//
// math client HTTP transport
//
// Command:
// $ goa gen github.com/sgerogia/hello-goa/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the math service endpoint HTTP clients.
type Client struct {
	// Sum Doer is the HTTP client used to make requests to the sum endpoint.
	SumDoer goahttp.Doer

	// Mul Doer is the HTTP client used to make requests to the mul endpoint.
	MulDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the math service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		SumDoer:             doer,
		MulDoer:             doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Sum returns an endpoint that makes HTTP requests to the math service sum
// server.
func (c *Client) Sum() goa.Endpoint {
	var (
		encodeRequest  = EncodeSumRequest(c.encoder)
		decodeResponse = DecodeSumResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSumRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SumDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("math", "sum", err)
		}
		return decodeResponse(resp)
	}
}

// Mul returns an endpoint that makes HTTP requests to the math service mul
// server.
func (c *Client) Mul() goa.Endpoint {
	var (
		encodeRequest  = EncodeMulRequest(c.encoder)
		decodeResponse = DecodeMulResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildMulRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.MulDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("math", "mul", err)
		}
		return decodeResponse(resp)
	}
}
