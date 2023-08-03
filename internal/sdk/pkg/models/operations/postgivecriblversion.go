// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type PostGiveCriblVersionRequest struct {
	// Version number
	Version string `pathParam:"style=simple,explode=false,name=version"`
}

type PostGiveCriblVersionResponse struct {
	ContentType string
	// a list of string objects
	Cribl *shared.Cribl
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}