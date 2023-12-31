// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListExecutorObjectResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of Executor objects
	Executors   *shared.Executors
	StatusCode  int
	RawResponse *http.Response
}
