package logic_test

import (
	"context"
	"fmt"
	"github.com/sgerogia/hello-goa/gen/token"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJWTAuth_Valid(t *testing.T) {
	// arrange
	un := "user"
	pwd := "pwd"
	u := token.User{
		Username: &un,
		Password: &pwd,
	}
	jwt, err := testingCtx.tokensvc.GenerateJWT(&u, testingCtx.tokensvc.PrivateKey)
	require.NoError(t, err)

	// act
	_, err = testingCtx.mathsvc.JWTAuth(context.TODO(), jwt, nil)

	// assert
	require.NoError(t, err)
}

func TestJWTAuth_Invalid(t *testing.T) {
	// act
	_, err := testingCtx.mathsvc.JWTAuth(context.TODO(), "foobar", nil)

	// assert
	require.EqualError(t, err, "Invalid token")
}

func TestJWTAuth_WrongIssuer(t *testing.T) {
	// arrange
	iss := "foo"
	sub := "sub"
	c := testClaims{
		iss: &iss,
		sub: &sub,
	}
	jwt, err := generateJWT(&c, testingCtx.tokensvc)
	require.NoError(t, err)

	// act
	_, err = testingCtx.mathsvc.JWTAuth(context.TODO(), jwt, nil)

	// assert
	require.EqualError(t, err, "Unexpected token issuer: foo")
}

func TestJWTAuth_WrongSub(t *testing.T) {
	// arrange
	sub := ""
	c := testClaims{
		sub: &sub,
	}
	jwt, err := generateJWT(&c, testingCtx.tokensvc)
	require.NoError(t, err)

	// act
	_, err = testingCtx.mathsvc.JWTAuth(context.TODO(), jwt, nil)

	// assert
	require.EqualError(t, err, "Incorrect token subject: ")
}

func TestJWTAuth_DiffAud(t *testing.T) {
	// arrange
	sub := "sub"
	aud := "aud"
	c := testClaims{
		sub: &sub,
		aud: &aud,
	}
	jwt, err := generateJWT(&c, testingCtx.tokensvc)
	require.NoError(t, err)

	// act
	_, err = testingCtx.mathsvc.JWTAuth(context.TODO(), jwt, nil)

	// assert
	require.EqualError(t, err, "Incorrect token audience: aud")
}

func TestJWTAuth_Expired(t *testing.T) {
	// arrange
	sub := "sub"
	iat := time.Now().Add(-70 * time.Minute).Unix()
	exp := time.Now().Add(-10 * time.Minute).Unix()
	c := testClaims{
		sub: &sub,
		iat: &iat,
		exp: &exp,
	}
	jwt, err := generateJWT(&c, testingCtx.tokensvc)
	require.NoError(t, err)

	// act
	_, err = testingCtx.mathsvc.JWTAuth(context.TODO(), jwt, nil)

	// assert
	require.EqualError(t, err, "Invalid token")
}

func TestJWTAuth_WrongLifespan(t *testing.T) {
	// arrange
	sub := "sub"
	iat := time.Now().Add(-10 * time.Minute).Unix()
	exp := time.Now().Add(100 * time.Minute).Unix()
	c := testClaims{
		sub: &sub,
		iat: &iat,
		exp: &exp,
	}
	jwt, err := generateJWT(&c, testingCtx.tokensvc)
	require.NoError(t, err)

	// act
	_, err = testingCtx.mathsvc.JWTAuth(context.TODO(), jwt, nil)

	// assert
	require.EqualError(t, err, fmt.Sprintf("Unexpected token expiry: %d", exp))
}

