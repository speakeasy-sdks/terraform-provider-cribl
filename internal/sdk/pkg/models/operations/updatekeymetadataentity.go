// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateKeyMetadataEntityRequest struct {
	// KeyMetadataEntity object to be updated
	KeyMetadataEntity *shared.KeyMetadataEntity `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateKeyMetadataEntityResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of KeyMetadataEntity objects
	KeyMetadataEntity *shared.KeyMetadataEntity
	StatusCode        int
	RawResponse       *http.Response
}
