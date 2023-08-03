// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetCertificateRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetCertificateResponse struct {
	// a list of Certificate objects
	Certificate *shared.Certificate
	ContentType string
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}