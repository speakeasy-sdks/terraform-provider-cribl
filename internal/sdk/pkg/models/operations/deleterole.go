// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteRoleRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteRoleResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of Role objects
	Roles       *shared.Roles
	StatusCode  int
	RawResponse *http.Response
}
