// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UserProfile - UserProfile object
type UserProfile struct {
	Disabled bool     `json:"disabled"`
	Email    string   `json:"email"`
	First    string   `json:"first"`
	ID       string   `json:"id"`
	Last     string   `json:"last"`
	Password *string  `json:"password,omitempty"`
	Roles    []string `json:"roles,omitempty"`
	Username string   `json:"username"`
}
