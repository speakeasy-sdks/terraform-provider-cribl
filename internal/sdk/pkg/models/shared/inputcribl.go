// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type InputCriblConnections struct {
	// Select a Destination.
	Output string `json:"output"`
	// Select Pipeline or Pack. Optional.
	Pipeline *string `json:"pipeline,omitempty"`
}

type InputCriblMetadata struct {
	// Field name
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

// InputCriblPqCompression - Codec to use to compress the persisted data.
type InputCriblPqCompression string

const (
	InputCriblPqCompressionNone InputCriblPqCompression = "none"
	InputCriblPqCompressionGzip InputCriblPqCompression = "gzip"
)

func (e InputCriblPqCompression) ToPointer() *InputCriblPqCompression {
	return &e
}

func (e *InputCriblPqCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputCriblPqCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputCriblPqCompression: %v", v)
	}
}

// InputCriblPqMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputCriblPqMode string

const (
	InputCriblPqModeSmart  InputCriblPqMode = "smart"
	InputCriblPqModeAlways InputCriblPqMode = "always"
)

func (e InputCriblPqMode) ToPointer() *InputCriblPqMode {
	return &e
}

func (e *InputCriblPqMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputCriblPqMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputCriblPqMode: %v", v)
	}
}

type InputCriblPq struct {
	// The number of events to send downstream before committing that Stream has read them.
	CommitFrequency *int64 `json:"commitFrequency,omitempty"`
	// Codec to use to compress the persisted data.
	Compress *InputCriblPqCompression `json:"compress,omitempty"`
	// The maximum number of events to hold in memory before writing the events to disk.
	MaxBufferSize *int64 `json:"maxBufferSize,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	MaxFileSize *string `json:"maxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	MaxSize *string `json:"maxSize,omitempty"`
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputCriblPqMode `json:"mode,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/inputs/<input-id>.
	Path *string `json:"path,omitempty"`
}

type InputCriblType string

const (
	InputCriblTypeCribl InputCriblType = "cribl"
)

func (e InputCriblType) ToPointer() *InputCriblType {
	return &e
}

func (e *InputCriblType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cribl":
		*e = InputCriblType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputCriblType: %v", v)
	}
}

type InputCribl struct {
	// Direct connections to Destinations, optionally via a Pipeline or a Pack.
	Connections []InputCriblConnections `json:"connections,omitempty"`
	// Enable/disable this input
	Disabled *bool `json:"disabled,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Unique ID for this input
	ID string `json:"id"`
	// Fields to add to events from this input.
	Metadata []InputCriblMetadata `json:"metadata,omitempty"`
	// Pipeline to process data from this Source before sending it through the Routes.
	Pipeline *string       `json:"pipeline,omitempty"`
	Pq       *InputCriblPq `json:"pq,omitempty"`
	// For details on Persistent Queues, see: [https://docs.cribl.io/stream/persistent-queues](https://docs.cribl.io/stream/persistent-queues)
	PqEnabled *bool `json:"pqEnabled,omitempty"`
	// Select whether to send data to Routes, or directly to Destinations.
	SendToRoutes *bool `json:"sendToRoutes,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string       `json:"streamtags,omitempty"`
	Type       InputCriblType `json:"type"`
}