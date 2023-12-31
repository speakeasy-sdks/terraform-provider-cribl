// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListDatasetObjectsResponse struct {
	ContentType string
	// a list of Dataset objects
	Datasets *shared.Datasets
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
