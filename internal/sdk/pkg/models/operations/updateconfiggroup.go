// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateConfigGroupRequest struct {
	// ConfigGroup object to be updated
	ConfigGroup *shared.ConfigGroup `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateConfigGroupResponse struct {
	// a list of ConfigGroup objects
	ConfigGroup *shared.ConfigGroup
	ContentType string
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
