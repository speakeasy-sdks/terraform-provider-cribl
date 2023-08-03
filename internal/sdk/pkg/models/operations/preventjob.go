// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type PreventJobRequest struct {
	// Job Instance id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type PreventJobResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of JobInfo objects
	JobInfos    *shared.JobInfos
	StatusCode  int
	RawResponse *http.Response
}