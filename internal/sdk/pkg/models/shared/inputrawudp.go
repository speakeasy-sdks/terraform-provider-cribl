// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type InputRawUDPConnections struct {
	// Select a Destination.
	Output string `json:"output"`
	// Select Pipeline or Pack. Optional.
	Pipeline *string `json:"pipeline,omitempty"`
}

type InputRawUDPMetadata struct {
	// Field name
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

// InputRawUDPPqCompression - Codec to use to compress the persisted data.
type InputRawUDPPqCompression string

const (
	InputRawUDPPqCompressionNone InputRawUDPPqCompression = "none"
	InputRawUDPPqCompressionGzip InputRawUDPPqCompression = "gzip"
)

func (e InputRawUDPPqCompression) ToPointer() *InputRawUDPPqCompression {
	return &e
}

func (e *InputRawUDPPqCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputRawUDPPqCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputRawUDPPqCompression: %v", v)
	}
}

// InputRawUDPPqMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputRawUDPPqMode string

const (
	InputRawUDPPqModeSmart  InputRawUDPPqMode = "smart"
	InputRawUDPPqModeAlways InputRawUDPPqMode = "always"
)

func (e InputRawUDPPqMode) ToPointer() *InputRawUDPPqMode {
	return &e
}

func (e *InputRawUDPPqMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputRawUDPPqMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputRawUDPPqMode: %v", v)
	}
}

type InputRawUDPPq struct {
	// The number of events to send downstream before committing that Stream has read them.
	CommitFrequency *int64 `json:"commitFrequency,omitempty"`
	// Codec to use to compress the persisted data.
	Compress *InputRawUDPPqCompression `json:"compress,omitempty"`
	// The maximum number of events to hold in memory before writing the events to disk.
	MaxBufferSize *int64 `json:"maxBufferSize,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	MaxFileSize *string `json:"maxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	MaxSize *string `json:"maxSize,omitempty"`
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputRawUDPPqMode `json:"mode,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/inputs/<input-id>.
	Path *string `json:"path,omitempty"`
}

type InputRawUDPType string

const (
	InputRawUDPTypeRawUDP InputRawUDPType = "raw_udp"
)

func (e InputRawUDPType) ToPointer() *InputRawUDPType {
	return &e
}

func (e *InputRawUDPType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "raw_udp":
		*e = InputRawUDPType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputRawUDPType: %v", v)
	}
}

type InputRawUDP struct {
	// Direct connections to Destinations, optionally via a Pipeline or a Pack.
	Connections []InputRawUDPConnections `json:"connections,omitempty"`
	// Enable/disable this input
	Disabled *bool `json:"disabled,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Address to bind on. For IPv4 (all addresses), use the default '0.0.0.0'. For IPv6, enter '::' (all addresses) or specify an IP address.
	Host string `json:"host"`
	// Unique ID for this input
	ID *string `json:"id,omitempty"`
	// If true, a __rawBytes field will be added to each event containing the raw bytes of the datagram.
	IngestRawBytes *bool `json:"ingestRawBytes,omitempty"`
	// Regex matching IP addresses that are allowed to send data
	IPWhitelistRegex *string `json:"ipWhitelistRegex,omitempty"`
	// Maximum number of events to buffer when downstream is blocking.
	MaxBufferSize *int64 `json:"maxBufferSize,omitempty"`
	// Fields to add to events from this input.
	Metadata []InputRawUDPMetadata `json:"metadata,omitempty"`
	// Pipeline to process data from this Source before sending it through the Routes.
	Pipeline *string `json:"pipeline,omitempty"`
	// Port to listen on.
	Port int64          `json:"port"`
	Pq   *InputRawUDPPq `json:"pq,omitempty"`
	// For details on Persistent Queues, see: [https://docs.cribl.io/stream/persistent-queues](https://docs.cribl.io/stream/persistent-queues)
	PqEnabled *bool `json:"pqEnabled,omitempty"`
	// Select whether to send data to Routes, or directly to Destinations.
	SendToRoutes *bool `json:"sendToRoutes,omitempty"`
	// If true, each UDP packet is assumed to contain a single message. If false, each UDP packet is assumed to contain multiple messages, separated by newlines.
	SingleMsgUDPPackets *bool `json:"singleMsgUdpPackets,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string         `json:"streamtags,omitempty"`
	Type       *InputRawUDPType `json:"type,omitempty"`
}
