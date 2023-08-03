// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CaptureParams - CaptureParams object
type CaptureParams struct {
	Filter          string  `json:"filter"`
	Level           int64   `json:"level"`
	MaxEvents       int64   `json:"maxEvents"`
	StepDuration    *int64  `json:"stepDuration,omitempty"`
	WorkerID        *string `json:"workerId,omitempty"`
	WorkerThreshold *int64  `json:"workerThreshold,omitempty"`
}