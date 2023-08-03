// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetJobResultRequest struct {
	// Group the job belongs to
	Group string `pathParam:"style=simple,explode=false,name=group"`
	// Instance id of the job to get results for
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Maximum files to get job results
	MaxFiles *int64 `queryParam:"style=form,explode=true,name=maxFiles"`
}

type GetJobResultResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of any objects
	JobResult   *shared.JobResult
	StatusCode  int
	RawResponse *http.Response
}
