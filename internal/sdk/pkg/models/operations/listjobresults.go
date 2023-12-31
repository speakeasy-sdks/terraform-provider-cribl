// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListJobResultsRequest struct {
	// Job instance id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type ListJobResultsResponse struct {
	ContentType string
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
	// Get results for a discover job by instance id
	ListJobResults200ApplicationXNdjsonBinaryString []byte
}
