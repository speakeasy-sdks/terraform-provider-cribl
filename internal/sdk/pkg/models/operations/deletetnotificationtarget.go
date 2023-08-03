// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeletetNotificationTargetRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeletetNotificationTargetResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of NotificationTarget objects
	NotificationTargets *shared.NotificationTargets
	StatusCode          int
	RawResponse         *http.Response
}