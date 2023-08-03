// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSearchJobResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of SearchJob objects
	SearchJob   *shared.SearchJob
	StatusCode  int
	RawResponse *http.Response
}
