// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListLicensesResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of License objects
	Licenses    *shared.Licenses
	StatusCode  int
	RawResponse *http.Response
}
