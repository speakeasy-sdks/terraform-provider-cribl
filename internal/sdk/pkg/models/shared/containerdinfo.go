// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ContainerdInfo struct {
	Container ContainerdContainer    `json:"container"`
	Image     map[string]interface{} `json:"image,omitempty"`
	Metrics   map[string]interface{} `json:"metrics,omitempty"`
	Namespace map[string]interface{} `json:"namespace,omitempty"`
	Task      *ContainerdTask        `json:"task,omitempty"`
}