// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListPipelineObjectResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of Pipeline objects
	Pipelines   *shared.Pipelines
	StatusCode  int
	RawResponse *http.Response
}
