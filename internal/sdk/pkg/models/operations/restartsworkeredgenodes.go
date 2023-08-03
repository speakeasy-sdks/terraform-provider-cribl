// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type RestartsWorkerEdgeNodesResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of RestartResponse objects
	RestartResponses *shared.RestartResponses
	StatusCode       int
	RawResponse      *http.Response
}