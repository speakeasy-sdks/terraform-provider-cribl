// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AppscopeLibEntry - New AppscopeLibEntry object
type AppscopeLibEntry struct {
	Config      AppscopeConfigWithCustom `json:"config"`
	Description string                   `json:"description"`
	ID          string                   `json:"id"`
	Lib         CriblLib                 `json:"lib"`
	Tags        *string                  `json:"tags,omitempty"`
}
