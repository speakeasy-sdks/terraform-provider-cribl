// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type LogoutIDPUserAuthRequest struct {
	RelayState *string `queryParam:"style=form,explode=true,name=RelayState"`
	// Logout request object
	SAMLResponse *string `queryParam:"style=form,explode=true,name=SAMLResponse"`
}

type LogoutIDPUserAuthResponse struct {
	ContentType string
	// Bad Request
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
	// Logout success
	Success *shared.Success
}
