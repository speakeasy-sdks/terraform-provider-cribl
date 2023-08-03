// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteSavedJobRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteSavedJobResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of SavedJob objects
	SavedJobs   *shared.SavedJobs
	StatusCode  int
	RawResponse *http.Response
}
