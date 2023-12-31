// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteDatasetObjectRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteDatasetObjectResponse struct {
	ContentType string
	// a list of Dataset objects
	Dataset interface{}
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
