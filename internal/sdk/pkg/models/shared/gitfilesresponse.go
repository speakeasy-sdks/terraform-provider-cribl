// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GitFilesResponse struct {
	CommitMessage map[string]interface{} `json:"commitMessage"`
	Count         int64                  `json:"count"`
	Items         []GitFile              `json:"items"`
}
