// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSavedQueriesResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of SavedQuery objects
	SavedQuery  *shared.SavedQuery
	StatusCode  int
	RawResponse *http.Response
}
