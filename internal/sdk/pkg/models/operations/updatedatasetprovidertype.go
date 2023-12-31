// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateDatasetProviderTypeRequest struct {
	// DatasetProviderType object to be updated
	DatasetProviderType *shared.DatasetProviderType `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateDatasetProviderTypeResponse struct {
	ContentType string
	// a list of DatasetProviderType objects
	DatasetProviderType *shared.DatasetProviderType
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
