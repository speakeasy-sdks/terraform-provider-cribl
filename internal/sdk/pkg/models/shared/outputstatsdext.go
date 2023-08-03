// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OutputStatsdExtBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputStatsdExtBackpressureBehavior string

const (
	OutputStatsdExtBackpressureBehaviorQueue OutputStatsdExtBackpressureBehavior = "queue"
	OutputStatsdExtBackpressureBehaviorDrop  OutputStatsdExtBackpressureBehavior = "drop"
	OutputStatsdExtBackpressureBehaviorBlock OutputStatsdExtBackpressureBehavior = "block"
)

func (e OutputStatsdExtBackpressureBehavior) ToPointer() *OutputStatsdExtBackpressureBehavior {
	return &e
}

func (e *OutputStatsdExtBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputStatsdExtBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputStatsdExtBackpressureBehavior: %v", v)
	}
}

// OutputStatsdExtCompression - Codec to use to compress the persisted data.
type OutputStatsdExtCompression string

const (
	OutputStatsdExtCompressionNone OutputStatsdExtCompression = "none"
	OutputStatsdExtCompressionGzip OutputStatsdExtCompression = "gzip"
)

func (e OutputStatsdExtCompression) ToPointer() *OutputStatsdExtCompression {
	return &e
}

func (e *OutputStatsdExtCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputStatsdExtCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputStatsdExtCompression: %v", v)
	}
}

type OutputStatsdExtPqControls struct {
}

// OutputStatsdExtQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputStatsdExtQueueFullBehavior string

const (
	OutputStatsdExtQueueFullBehaviorBlock OutputStatsdExtQueueFullBehavior = "block"
	OutputStatsdExtQueueFullBehaviorDrop  OutputStatsdExtQueueFullBehavior = "drop"
)

func (e OutputStatsdExtQueueFullBehavior) ToPointer() *OutputStatsdExtQueueFullBehavior {
	return &e
}

func (e *OutputStatsdExtQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputStatsdExtQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputStatsdExtQueueFullBehavior: %v", v)
	}
}

// OutputStatsdExtDestinationProtocol - Protocol to use when communicating with the destination.
type OutputStatsdExtDestinationProtocol string

const (
	OutputStatsdExtDestinationProtocolTCP OutputStatsdExtDestinationProtocol = "tcp"
)

func (e OutputStatsdExtDestinationProtocol) ToPointer() *OutputStatsdExtDestinationProtocol {
	return &e
}

func (e *OutputStatsdExtDestinationProtocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tcp":
		*e = OutputStatsdExtDestinationProtocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputStatsdExtDestinationProtocol: %v", v)
	}
}

type OutputStatsdExtType string

const (
	OutputStatsdExtTypeStatsdExt OutputStatsdExtType = "statsd_ext"
)

func (e OutputStatsdExtType) ToPointer() *OutputStatsdExtType {
	return &e
}

func (e *OutputStatsdExtType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "statsd_ext":
		*e = OutputStatsdExtType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputStatsdExtType: %v", v)
	}
}

type OutputStatsdExt struct {
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
	OnBackpressure *OutputStatsdExtBackpressureBehavior `json:"onBackpressure,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Destination port.
	Port int64 `json:"port"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputStatsdExtCompression `json:"pqCompress,omitempty"`
	PqControls *OutputStatsdExtPqControls  `json:"pqControls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `json:"pqMaxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `json:"pqMaxSize,omitempty"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputStatsdExtQueueFullBehavior `json:"pqOnBackpressure,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `json:"pqPath,omitempty"`
	// Toggle this off to forward new events to receiver(s) before queue is flushed. Otherwise, default drain behavior is FIFO (first in, first out).
	PqStrictOrdering *bool `json:"pqStrictOrdering,omitempty"`
	// Protocol to use when communicating with the destination.
	Protocol OutputStatsdExtDestinationProtocol `json:"protocol"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string `json:"systemFields,omitempty"`
	// Rate (in bytes per second) to throttle while writing to an output. Also takes values with multiple-byte units, such as KB, MB, GB, etc. (E.g., 42 MB.) Default value of 0 specifies no throttling.
	ThrottleRatePerSec *string              `json:"throttleRatePerSec,omitempty"`
	Type               *OutputStatsdExtType `json:"type,omitempty"`
	// Amount of time (milliseconds) to wait for a write to complete before assuming connection is dead
	WriteTimeout *int64 `json:"writeTimeout,omitempty"`
}