// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListDiagBundlesResponse struct {
	ContentType string
	// Invalid path
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
	// Get list of existing diag bundles
	ListDiagBundles200ApplicationTarPlusGzipBinaryString []byte
}
