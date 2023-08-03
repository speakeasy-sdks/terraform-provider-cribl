// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OutputStatsdBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputStatsdBackpressureBehavior string

const (
	OutputStatsdBackpressureBehaviorQueue OutputStatsdBackpressureBehavior = "queue"
	OutputStatsdBackpressureBehaviorDrop  OutputStatsdBackpressureBehavior = "drop"
	OutputStatsdBackpressureBehaviorBlock OutputStatsdBackpressureBehavior = "block"
)

func (e OutputStatsdBackpressureBehavior) ToPointer() *OutputStatsdBackpressureBehavior {
	return &e
}

func (e *OutputStatsdBackpressureBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "queue":
		fallthrough
	case "drop":
		fallthrough
	case "block":
		*e = OutputStatsdBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputStatsdBackpressureBehavior: %v", v)
	}
}

// OutputStatsdCompression - Codec to use to compress the persisted data.
type OutputStatsdCompression string

const (
	OutputStatsdCompressionNone OutputStatsdCompression = "none"
	OutputStatsdCompressionGzip OutputStatsdCompression = "gzip"
)

func (e OutputStatsdCompression) ToPointer() *OutputStatsdCompression {
	return &e
}

func (e *OutputStatsdCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputStatsdCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputStatsdCompression: %v", v)
	}
}

type OutputStatsdPqControls struct {
}

// OutputStatsdQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputStatsdQueueFullBehavior string

const (
	OutputStatsdQueueFullBehaviorBlock OutputStatsdQueueFullBehavior = "block"
	OutputStatsdQueueFullBehaviorDrop  OutputStatsdQueueFullBehavior = "drop"
)

func (e OutputStatsdQueueFullBehavior) ToPointer() *OutputStatsdQueueFullBehavior {
	return &e
}

func (e *OutputStatsdQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputStatsdQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputStatsdQueueFullBehavior: %v", v)
	}
}

// OutputStatsdDestinationProtocol - Protocol to use when communicating with the destination.
type OutputStatsdDestinationProtocol string

const (
	OutputStatsdDestinationProtocolTCP OutputStatsdDestinationProtocol = "tcp"
)

func (e OutputStatsdDestinationProtocol) ToPointer() *OutputStatsdDestinationProtocol {
	return &e
}

func (e *OutputStatsdDestinationProtocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tcp":
		*e = OutputStatsdDestinationProtocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputStatsdDestinationProtocol: %v", v)
	}
}

type OutputStatsdType string

const (
	OutputStatsdTypeStatsd OutputStatsdType = "statsd"
)

func (e OutputStatsdType) ToPointer() *OutputStatsdType {
	return &e
}

func (e *OutputStatsdType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "statsd":
		*e = OutputStatsdType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputStatsdType: %v", v)
	}
}

type OutputStatsd struct {
	// Amount of time (milliseconds) to wait for the connection to establish before retrying
	ConnectionTimeout *int64 `json:"connectionTimeout,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Used when Protocol is TCP, to specify how often buffers should be flushed resulting in records sent to the destination.
	FlushPeriodSec *int64 `json:"flushPeriodSec,omitempty"`
	// The hostname of the destination.
	Host string `json:"host"`
	// Unique ID for this output
	ID *string `json:"id,omitempty"`
	// Used when Protocol is UDP, to specify the maximum size of packets sent to the destination. Also known as the MTU for the network path to the destination system.
	Mtu *int64 `json:"mtu,omitempty"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputStatsdBackpressureBehavior `json:"onBackpressure,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Destination port.
	Port int64 `json:"port"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputStatsdCompression `json:"pqCompress,omitempty"`
	PqControls *OutputStatsdPqControls  `json:"pqControls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `json:"pqMaxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `json:"pqMaxSize,omitempty"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputStatsdQueueFullBehavior `json:"pqOnBackpressure,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `json:"pqPath,omitempty"`
	// Toggle this off to forward new events to receiver(s) before queue is flushed. Otherwise, default drain behavior is FIFO (first in, first out).
	PqStrictOrdering *bool `json:"pqStrictOrdering,omitempty"`
	// Protocol to use when communicating with the destination.
	Protocol OutputStatsdDestinationProtocol `json:"protocol"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string `json:"systemFields,omitempty"`
	// Rate (in bytes per second) to throttle while writing to an output. Also takes values with multiple-byte units, such as KB, MB, GB, etc. (E.g., 42 MB.) Default value of 0 specifies no throttling.
	ThrottleRatePerSec *string           `json:"throttleRatePerSec,omitempty"`
	Type               *OutputStatsdType `json:"type,omitempty"`
	// Amount of time (milliseconds) to wait for a write to complete before assuming connection is dead
	WriteTimeout *int64 `json:"writeTimeout,omitempty"`
}
