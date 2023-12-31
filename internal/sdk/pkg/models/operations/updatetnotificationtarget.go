// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdatetNotificationTargetRequest struct {
	// NotificationTarget object to be updated
	NotificationTarget *shared.NotificationTarget `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdatetNotificationTargetResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of NotificationTarget objects
	NotificationTarget *shared.NotificationTarget
	StatusCode         int
	RawResponse        *http.Response
}
