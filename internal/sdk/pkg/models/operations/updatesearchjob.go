// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateSearchJobRequest struct {
	// SearchJob object to be updated
	SearchJob *shared.SearchJob `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateSearchJobResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of SearchJob objects
	SearchJob   *shared.SearchJob
	StatusCode  int
	RawResponse *http.Response
}
