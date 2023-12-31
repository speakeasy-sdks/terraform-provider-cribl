// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteJobRequest struct {
	// Job instance id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteJobResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of any objects
	JobDelete   *shared.JobDelete
	StatusCode  int
	RawResponse *http.Response
}
