package logic_test

import (
	"context"
	"github.com/sgerogia/hello-goa/gen/math"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)


func TestMul_Valid(t *testing.T) {
	// arrange
	p := math.MulPayload{
		Token: "token",
		Numbers: []string{"6", "4"},
	}

	// act
	res, err := testingCtx.mathsvc.Mul(context.TODO(), &p)
	require.NoError(t, err)

	// assert
	assert.Equal(t, "24", res)
}

func TestMul_NoToken(t *testing.T) {
	// arrange
	p := math.MulPayload{
		Token: "",
		Numbers: []string{"6", "4"},
	}

	// act
	_, err := testingCtx.mathsvc.Mul(context.TODO(), &p)

	// assert
	require.EqualError(t, err, "Invalid token")
}

func TestMul_EmptyArray(t *testing.T) {
	// arrange
	p := math.MulPayload{
		Token: "token",
		Numbers: []string{},
	}

	// act
	res, err := testingCtx.mathsvc.Mul(context.TODO(), &p)
	require.NoError(t, err)

	// assert
	assert.Equal(t, "0", res)
}

func TestMul_InvalidNumber(t *testing.T) {
	// arrange
	p := math.MulPayload{
		Token: "token",
		Numbers: []string{"10", "a"},
	}

	// act
	_, err := testingCtx.mathsvc.Mul(context.TODO(), &p)

	// assert
	require.Error(t, err)
	require.Contains(t, err.Error(), "Malformed number: a")
}