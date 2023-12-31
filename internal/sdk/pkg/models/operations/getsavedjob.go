// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSavedJobRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetSavedJobResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of SavedJob objects
	SavedJob    *shared.SavedJob
	StatusCode  int
	RawResponse *http.Response
}
