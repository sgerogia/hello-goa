package logic_test

import (
	"context"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/sgerogia/hello-goa/gen/token"
	"github.com/sgerogia/hello-goa/logic"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGenerateJWT(t *testing.T) {
	t.Parallel()

	// arrange
	un := "user"
	pwd := "pwd"
	u := token.User{
		Username: &un,
		Password: &pwd,
	}
	claims := &logic.Claims{}

	// act
	jwt, err := testingCtx.tokensvc.GenerateJWT(&u, testingCtx.tokensvc.PrivateKey)
	require.NoError(t, err)

	// assert
	token, err := jwtgo.ParseWithClaims(jwt, claims, func(token *jwtgo.Token) (interface{}, error) {
		return testingCtx.mathsvc.PublicKey, nil
	})
	require.NoError(t, err)

	require.NoError(t, token.Claims.Valid())

	assert.NotEmpty(t, claims.StandardClaims.Id)
	assert.Equal(t, testingCtx.tokensvc.Url, claims.StandardClaims.Issuer)
	assert.Equal(t, un, claims.StandardClaims.Subject)
	assert.Equal(t, un, claims.StandardClaims.Audience)
	exp := time.Unix(claims.ExpiresAt, 0)
	iat := time.Unix(claims.IssuedAt, 0)
	assert.Equal(t, time.Duration(time.Hour), exp.Sub(iat))
}

func TestAuth_ValidUser(t *testing.T) {
	t.Parallel()

	// arrange
	un := "user"
	pwd := "pwd"
	u := token.User{
		Username: &un,
		Password: &pwd,
	}

	// act
	jwt, err := testingCtx.tokensvc.Auth(context.TODO(), &u)

	// assert
	require.NoError(t, err)
	require.NotEmpty(t, jwt)
}

func TestAuth_InvalidUser_EmptyField(t *testing.T) {
	t.Parallel()

	// arrange
	un := "user"
	pwd := ""
	u := token.User{
		Username: &un,
		Password: &pwd,
	}

	// act
	jwt, err := testingCtx.tokensvc.Auth(context.TODO(), &u)

	// assert
	require.Error(t, err)
	require.Empty(t, jwt)
	require.Equal(t, err.Error(), "MalformedPayload")
}

func TestAuth_InvalidUser_NilField(t *testing.T) {
	t.Parallel()

	// arrange
	pwd := "pwd"
	u := token.User{
		Username: nil,
		Password: &pwd,
	}

	// act
	jwt, err := testingCtx.tokensvc.Auth(context.TODO(), &u)

	// assert
	require.Error(t, err)
	require.Empty(t, jwt)
	require.Equal(t, err.Error(), "MalformedPayload")
}