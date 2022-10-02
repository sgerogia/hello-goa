package logic_test

import (
	"context"
	"encoding/json"
	"github.com/sgerogia/hello-goa/gen/math"
	"github.com/sgerogia/hello-goa/logic"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProcessJson(t *testing.T) {

	var doc interface{}

	json.Unmarshal([]byte(`{"a":6,"b":4}`), &doc)
	assert.Equal(t, "10", logic.ProcessJson(doc))

	json.Unmarshal([]byte(`[1,2,3,4]`), &doc)
	assert.Equal(t, "10", logic.ProcessJson(doc))

	json.Unmarshal([]byte(`[[[2]]]`), &doc)
	assert.Equal(t, "2", logic.ProcessJson(doc))

	json.Unmarshal([]byte(`{"a":{"b":4},"c":-2}`), &doc)
	assert.Equal(t, "2", logic.ProcessJson(doc))

	json.Unmarshal([]byte(`{"a":[-1,1,"dark"]}`), &doc)
	assert.Equal(t, "0", logic.ProcessJson(doc))

	json.Unmarshal([]byte(`[-1,{"a":1, "b":"light"}]`), &doc)
	assert.Equal(t, "0", logic.ProcessJson(doc))

	json.Unmarshal([]byte(`[]`), &doc)
	assert.Equal(t, "0", logic.ProcessJson(doc))

	json.Unmarshal([]byte(`{}`), &doc)
	assert.Equal(t, "0", logic.ProcessJson(doc))

	json.Unmarshal([]byte(`[-1.2, {"a":1.5432, "b":"light", "c":2.3423}]`), &doc)
	assert.Equal(t, "2.68550", logic.ProcessJson(doc))
}

func TestSum_Valid(t *testing.T) {
	// arrange
	p := math.SumPayload{
		Token: "token",
		Doc: `{"a":6,"b":4}`,
	}

	// act
	res, err := testingCtx.mathsvc.Sum(context.TODO(), &p)
	require.NoError(t, err)

	// assert
	assert.Equal(t, "10", res)
}

func TestSum_NoToken(t *testing.T) {
	// arrange
	p := math.SumPayload{
		Token: "",
		Doc: `{"a":6,"b":4}`,
	}

	// act
	_, err := testingCtx.mathsvc.Sum(context.TODO(), &p)

	// assert
	require.EqualError(t, err, "Invalid token")
}

func TestSum_EmptyBody(t *testing.T) {
	// arrange
	p := math.SumPayload{
		Token: "token",
		Doc: ``,
	}

	// act
	_, err := testingCtx.mathsvc.Sum(context.TODO(), &p)

	// assert
	require.EqualError(t, err, "Empty body")
}

func TestSum_InvalidJSON(t *testing.T) {
	// arrange
	p := math.SumPayload{
		Token: "token",
		Doc: `{"a":6,"b":}`,
	}

	// act
	_, err := testingCtx.mathsvc.Sum(context.TODO(), &p)

	// assert
	require.Error(t, err)
	require.Contains(t, err.Error(), "Malformed JSON: ")
}