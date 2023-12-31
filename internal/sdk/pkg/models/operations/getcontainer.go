// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetContainerRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetContainerResponse struct {
	// a list of Container objects
	Containers  *shared.Containers
	ContentType string
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
