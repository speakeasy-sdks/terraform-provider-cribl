// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteLicenseRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteLicenseResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of License objects
	License     *shared.License
	StatusCode  int
	RawResponse *http.Response
}
