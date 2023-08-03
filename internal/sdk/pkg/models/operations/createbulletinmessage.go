// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateBulletinMessageResponse struct {
	// a list of BulletinMessage objects
	BulletinMessage *shared.BulletinMessage
	ContentType     string
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}