// Code generated by goa v3.8.5, DO NOT EDIT.
//
// math HTTP client CLI support package
//
// Command:
// $ goa gen github.com/sgerogia/hello-goa/design

package client

import (
	"encoding/json"
	"fmt"

	math "github.com/sgerogia/hello-goa/gen/math"
)

// BuildSumPayload builds the payload for the math sum endpoint from CLI flags.
func BuildSumPayload(mathSumBody string, mathSumToken string) (*math.SumPayload, error) {
	var body string
	{
		body = mathSumBody
	}
	var token string
	{
		token = mathSumToken
	}
	v := body
	res := &math.SumPayload{
		Doc: v,
	}
	res.Token = token

	return res, nil
}

// BuildMulPayload builds the payload for the math mul endpoint from CLI flags.
func BuildMulPayload(mathMulNumbers string, mathMulToken string) (*math.MulPayload, error) {
	var err error
	var numbers []string
	{
		err = json.Unmarshal([]byte(mathMulNumbers), &numbers)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for numbers, \nerror: %s, \nexample of valid JSON:\n%s", err, "'[\n      \"4\",\n      \"3.543\",\n      \"-2\"\n   ]'")
		}
	}
	var token string
	{
		token = mathMulToken
	}
	v := &math.MulPayload{}
	v.Numbers = numbers
	v.Token = token

	return v, nil
}
