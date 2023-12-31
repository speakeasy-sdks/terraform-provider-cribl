// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateDataSampleIDRequest struct {
	// DataSample object to be updated
	DataSample *shared.DataSample `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateDataSampleIDResponse struct {
	ContentType string
	// a list of DataSample objects
	DataSample *shared.DataSample
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
