package logic_test

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/sgerogia/hello-goa/logic"
	"log"
	"os"
	"testing"
	"time"
)

const (
	URL         = "https://localhost:8080"
	PRIVATE_KEY = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`
	PUBLIC_KEY = `-----BEGIN PUBLIC KEY-----
MIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEAsOCWb+2BnDiBHKPSz1J3
XhGRgOniCCf4Nlz00kcsfM3qazCI2q1HJ6h5khi4kNGRWwAVPVKqti0LOgfJMN+O
YeVPp7fR2yQJVkvGOXYzBpVRvUUCeLMR6Av9Eh4H7ssj9aVawPnotf4525htEyXh
IqMImFui44ddVmd8ugbIMGAHTRLfH7nPGrgiP09FmyqTo7mJUJgsDBB4I9uVseB8
uNKEXJhw0HOjTrr24tuOq6YdexVOx65ZaaTEAT5RZwXjcwKDJbTvzcAaUtni5NSp
hqWXNqjOymf7wKHqB6F1okKfK898KTLdHoXVigzu+aaWDGAJIEISlXAdSUVSevJi
OiXVdenIwrn92b373UyPU69UZuOz+uzcrHiEec7EGIx2mKt2aWSaSJAoL9PMggWM
w1+weoQEk4+lBMhIqnHiriwSVxSKqXrXy8AXh1R171jkkd/W45KpnsiJ1BvSk9RD
Bf+/8yBkAYlxCharVgbIVhF6TX18XQmKrMNaKDwnOKTDAgMBAAE=
-----END PUBLIC KEY-----
`
)

// --- Global variable for child tests ---
var testingCtx struct {
	mathsvc *logic.MathSrvc
	tokensvc *logic.TokenSrvc
}

func TestMain(m *testing.M) {

	// global setup
	testingCtx.mathsvc = newMathSvc()
	testingCtx.tokensvc = newTokenSvc()

	// Run test suites
	exitVal := m.Run()

	// clean up here

	// ...and exit test suite
	os.Exit(exitVal)
}

// Wrap normal initializer, explicitly cast to service struct to expose internal methods
func newMathSvc() *logic.MathSrvc {
	pKey, _ := jwtgo.ParseRSAPublicKeyFromPEM([]byte(PUBLIC_KEY))
	jwtExp := 60

	return logic.NewMath(
		log.New(os.Stderr, "[math-test] ", log.Ltime),
		URL,
		pKey,
		&jwtExp,
	).(*logic.MathSrvc)
}

// Wrap normal initializer, explicitly cast to service struct to expose internal methods
func newTokenSvc() *logic.TokenSrvc {
	key, _ := jwtgo.ParseRSAPrivateKeyFromPEM([]byte(PRIVATE_KEY))
	jwtExp := 60

	return logic.NewToken(
		log.New(os.Stderr, "[token-test] ", log.Ltime),
		URL,
		key,
		&jwtExp,
	).(*logic.TokenSrvc)
}

type testClaims struct {
	sub *string
	iat *int64
	exp *int64
	iss *string
	aud *string
}

func generateJWT(c *testClaims, s *logic.TokenSrvc) (res string, err error) {

	var iat int64
	if c.iat != nil {
		iat = *c.iat
	} else {
		iat = time.Now().Unix()
	}

	var exp int64
	if c.exp != nil {
		exp = *c.exp
	} else {
		exp = time.Now().Add(time.Duration(*s.JwtExpiryMins) * time.Minute).Unix()
	}

	var iss string
	if c.iss != nil {
		iss = *c.iss
	} else {
		iss = s.Url
	}

	var aud string
	if c.aud != nil {
		aud = *c.aud
	} else {
		aud = *c.sub
	}

	// Required claims: https://openid.net/specs/openid-connect-core-1_0.html#IDToken
	claims := &logic.Claims{
		StandardClaims: jwtgo.StandardClaims{
			Id:        uuid.NewString(),
			Subject:   *c.sub,
			IssuedAt:  iat,
			ExpiresAt: exp,
			Issuer:    iss,
			Audience:  aud,
		},
	}

	jwt := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, claims)
	return jwt.SignedString(s.PrivateKey)
}


