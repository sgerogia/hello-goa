package logic

import (
	"context"
	"crypto/rsa"
	"errors"
	"github.com/google/uuid"
	"log"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	token "github.com/sgerogia/hello-goa/gen/token"
)

// token service implementation.
type TokenSrvc struct {
	logger *log.Logger
	Url           string
	PrivateKey    *rsa.PrivateKey
	JwtExpiryMins *int
}

type Claims struct {
	jwtgo.StandardClaims
}

// NewToken returns the token service implementation.
func NewToken(logger *log.Logger, url string, privateKey *rsa.PrivateKey, jwtExpiry *int) token.Service {
	return &TokenSrvc{logger, url, privateKey, jwtExpiry}
}

// Accepts username and password in the body and returns JWT OAUTH2/OIDC token
// with the username as a subject, expiring in 1 hour.
// The username and password are not verified, but cannot be empty strings.
func (s *TokenSrvc) Auth(ctx context.Context, p *token.User) (res string, err error) {
	if !IsUserValid(p) {
		return "", errors.New("MalformedPayload")
	}

	jwt, err := s.GenerateJWT(p, s.PrivateKey)

	if err != nil {
		return "", err
	}

	return jwt, nil
}

// Generate an RS256 JWT for the user.
// Argument is assumed to be valid
func (s *TokenSrvc) GenerateJWT(u *token.User, key *rsa.PrivateKey) (res string, err error) {

	t := time.Now().Unix()
	exp := time.Now().Add(60 * time.Minute).Unix()

	// Required claims: https://openid.net/specs/openid-connect-core-1_0.html#IDToken
	claims := &Claims{
		StandardClaims: jwtgo.StandardClaims{
			Id:        uuid.NewString(),
			Subject:   *u.Username,
			IssuedAt:  t,
			ExpiresAt: exp,
			Issuer:    s.Url,
			Audience:  *u.Username,
		},
	}

	jwt := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, claims)
	return jwt.SignedString(key)
}

func IsUserValid(u *token.User) (res bool) {
	if u == nil || u.Username == nil || u.Password == nil {
		return false
	}
	if len(*u.Username) == 0 || len(*u.Password) == 0 {
		return false
	}
	return true
}
