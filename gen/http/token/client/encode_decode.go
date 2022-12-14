// Code generated by goa v3.8.5, DO NOT EDIT.
//
// token HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/sgerogia/hello-goa/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	token "github.com/sgerogia/hello-goa/gen/token"
	goahttp "goa.design/goa/v3/http"
)

// BuildAuthRequest instantiates a HTTP request object with method and path set
// to call the "token" service "auth" endpoint
func (c *Client) BuildAuthRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AuthTokenPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("token", "auth", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAuthRequest returns an encoder for requests sent to the token auth
// server.
func EncodeAuthRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*token.User)
		if !ok {
			return goahttp.ErrInvalidType("token", "auth", "*token.User", v)
		}
		body := NewAuthRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("token", "auth", err)
		}
		return nil
	}
}

// DecodeAuthResponse returns a decoder for responses returned by the token
// auth endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeAuthResponse may return the following errors:
//   - "MalformedPayload" (type *goa.ServiceError): http.StatusBadRequest
//   - error: internal error
func DecodeAuthResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("token", "auth", err)
			}
			return body, nil
		case http.StatusBadRequest:
			var (
				body AuthMalformedPayloadResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("token", "auth", err)
			}
			err = ValidateAuthMalformedPayloadResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("token", "auth", err)
			}
			return nil, NewAuthMalformedPayload(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("token", "auth", resp.StatusCode, string(body))
		}
	}
}
