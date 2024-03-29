// Code generated by goa v3.0.6, DO NOT EDIT.
//
// stress HTTP client types
//
// Command:
// $ goa gen github.com/ikawaha/kagome-stress/design

package client

import (
	stress "github.com/ikawaha/kagome-stress/gen/stress"
	stressviews "github.com/ikawaha/kagome-stress/gen/stress/views"
	goa "goa.design/goa/v3/pkg"
)

// TokenizeRequestBody is the type of the "stress" service "tokenize" endpoint
// HTTP request body.
type TokenizeRequestBody struct {
	Sentence string `form:"sentence" json:"sentence" xml:"sentence"`
}

// TokenizeResponseBody is the type of the "stress" service "tokenize" endpoint
// HTTP response body.
type TokenizeResponseBody []*TokenResponse

// TokenResponse is used to define fields on response body types.
type TokenResponse struct {
	Surface *string `form:"surface,omitempty" json:"surface,omitempty" xml:"surface,omitempty"`
	Pos     *string `form:"pos,omitempty" json:"pos,omitempty" xml:"pos,omitempty"`
	Start   *int    `form:"start,omitempty" json:"start,omitempty" xml:"start,omitempty"`
	End     *int    `form:"end,omitempty" json:"end,omitempty" xml:"end,omitempty"`
	Type    *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// NewTokenizeRequestBody builds the HTTP request body from the payload of the
// "tokenize" endpoint of the "stress" service.
func NewTokenizeRequestBody(p *stress.TokenizePayload) *TokenizeRequestBody {
	body := &TokenizeRequestBody{
		Sentence: p.Sentence,
	}
	return body
}

// NewTokenizeTokenCollectionOK builds a "stress" service "tokenize" endpoint
// result from a HTTP "OK" response.
func NewTokenizeTokenCollectionOK(body TokenizeResponseBody) stressviews.TokenCollectionView {
	v := make([]*stressviews.TokenView, len(body))
	for i, val := range body {
		v[i] = &stressviews.TokenView{
			Surface: val.Surface,
			Pos:     val.Pos,
			Start:   val.Start,
			End:     val.End,
			Type:    val.Type,
		}
	}
	return v
}

// ValidateTokenResponse runs the validations defined on TokenResponse
func ValidateTokenResponse(body *TokenResponse) (err error) {
	if body.Surface == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("surface", "body"))
	}
	if body.Pos == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("pos", "body"))
	}
	if body.Start == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("start", "body"))
	}
	if body.End == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("end", "body"))
	}
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	return
}
