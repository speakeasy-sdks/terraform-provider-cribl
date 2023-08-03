// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OutputLokiAuthenticationType - The authentication method to use for the HTTP requests
type OutputLokiAuthenticationType string

const (
	OutputLokiAuthenticationTypeCredentialsSecret OutputLokiAuthenticationType = "credentialsSecret"
	OutputLokiAuthenticationTypeToken             OutputLokiAuthenticationType = "token"
	OutputLokiAuthenticationTypeTextSecret        OutputLokiAuthenticationType = "textSecret"
	OutputLokiAuthenticationTypeBasic             OutputLokiAuthenticationType = "basic"
	OutputLokiAuthenticationTypeNone              OutputLokiAuthenticationType = "none"
)

func (e OutputLokiAuthenticationType) ToPointer() *OutputLokiAuthenticationType {
	return &e
}

func (e *OutputLokiAuthenticationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "credentialsSecret":
		fallthrough
	case "token":
		fallthrough
	case "textSecret":
		fallthrough
	case "basic":
		fallthrough
	case "none":
		*e = OutputLokiAuthenticationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputLokiAuthenticationType: %v", v)
	}
}

type OutputLokiExtraHTTPHeaders struct {
	// Field name
	Name *string `json:"name,omitempty"`
	// Field value
	Value string `json:"value"`
}

// OutputLokiFailedRequestLoggingMode - Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
type OutputLokiFailedRequestLoggingMode string

const (
	OutputLokiFailedRequestLoggingModePayload           OutputLokiFailedRequestLoggingMode = "payload"
	OutputLokiFailedRequestLoggingModePayloadAndHeaders OutputLokiFailedRequestLoggingMode = "payloadAndHeaders"
	OutputLokiFailedRequestLoggingModeNone              OutputLokiFailedRequestLoggingMode = "none"
)

func (e OutputLokiFailedRequestLoggingMode) ToPointer() *OutputLokiFailedRequestLoggingMode {
	return &e
}

func (e *OutputLokiFailedRequestLoggingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "payload":
		fallthrough
	case "payloadAndHeaders":
		fallthrough
	case "none":
		*e = OutputLokiFailedRequestLoggingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputLokiFailedRequestLoggingMode: %v", v)
	}
}

type OutputLokiLabels struct {
	// Name of the label.
	Name string `json:"name"`
	// Value of the label.
	Value string `json:"value"`
}

// OutputLokiMessageFormat - Which format to use when sending logs to Loki (Protobuf or JSON).  Defaults to Protobuf.
type OutputLokiMessageFormat string

const (
	OutputLokiMessageFormatProtobuf OutputLokiMessageFormat = "protobuf"
	OutputLokiMessageFormatJSON     OutputLokiMessageFormat = "json"
)

func (e OutputLokiMessageFormat) ToPointer() *OutputLokiMessageFormat {
	return &e
}

func (e *OutputLokiMessageFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "protobuf":
		fallthrough
	case "json":
		*e = OutputLokiMessageFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputLokiMessageFormat: %v", v)
	}
}

// OutputLokiBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputLokiBackpressureBehavior string

const (
	OutputLokiBackpressureBehaviorQueue OutputLokiBackpressureBehavior = "queue"
	OutputLokiBackpressureBehaviorDrop  OutputLokiBackpressureBehavior = "drop"
	OutputLokiBackpressureBehaviorBlock OutputLokiBackpressureBehavior = "block"
)

func (e OutputLokiBackpressureBehavior) ToPointer() *OutputLokiBackpressureBehavior {
	return &e
}

func (e *OutputLokiBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputLokiBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputLokiBackpressureBehavior: %v", v)
	}
}

// OutputLokiCompression - Codec to use to compress the persisted data.
type OutputLokiCompression string

const (
	OutputLokiCompressionNone OutputLokiCompression = "none"
	OutputLokiCompressionGzip OutputLokiCompression = "gzip"
)

func (e OutputLokiCompression) ToPointer() *OutputLokiCompression {
	return &e
}

func (e *OutputLokiCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputLokiCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputLokiCompression: %v", v)
	}
}

type OutputLokiPqControls struct {
}

// OutputLokiQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputLokiQueueFullBehavior string

const (
	OutputLokiQueueFullBehaviorBlock OutputLokiQueueFullBehavior = "block"
	OutputLokiQueueFullBehaviorDrop  OutputLokiQueueFullBehavior = "drop"
)

func (e OutputLokiQueueFullBehavior) ToPointer() *OutputLokiQueueFullBehavior {
	return &e
}

func (e *OutputLokiQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputLokiQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputLokiQueueFullBehavior: %v", v)
	}
}

type OutputLokiType string

const (
	OutputLokiTypeLoki OutputLokiType = "loki"
)

func (e OutputLokiType) ToPointer() *OutputLokiType {
	return &e
}

func (e *OutputLokiType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "loki":
		*e = OutputLokiType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputLokiType: %v", v)
	}
}

type OutputLoki struct {
	// The authentication method to use for the HTTP requests
	AuthType *OutputLokiAuthenticationType `json:"authType,omitempty"`
	// Whether to compress the payload body before sending.
	Compress *bool `json:"compress,omitempty"`
	// Maximum number of ongoing requests before blocking. Warning: Setting this value > 1 can cause Loki to complain about entries being delivered out of order.
	Concurrency *int64 `json:"concurrency,omitempty"`
	// Select (or create) a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Headers to add to all events.
	ExtraHTTPHeaders []OutputLokiExtraHTTPHeaders `json:"extraHttpHeaders,omitempty"`
	// Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
	FailedRequestLoggingMode *OutputLokiFailedRequestLoggingMode `json:"failedRequestLoggingMode,omitempty"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Maximum time between requests. Small values can reduce the payload size below the configured 'Max record size' and 'Max events per request'. Warning: Setting this too low can increase the number of ongoing requests (depending on the value of 'Request concurrency'); this can cause Loki to complain about entries being delivered out of order.
	FlushPeriodSec *int64 `json:"flushPeriodSec,omitempty"`
	// Unique ID for this output
	ID *string `json:"id,omitempty"`
	// List of labels to send with logs. Labels define Loki streams, so use static labels to avoid proliferating label value combinations and streams. Can be merged and/or overridden by the event's __labels field (e.g.: '__labels: {host: "cribl.io", level: "error"}').
	Labels []OutputLokiLabels `json:"labels,omitempty"`
	// Maximum number of events to include in the request body. Defaults to 0 (unlimited). Warning: Setting this too low can increase the number of ongoing requests (depending on the value of 'Request concurrency'); this can cause Loki to complain about entries being delivered out of order.
	MaxPayloadEvents *int64 `json:"maxPayloadEvents,omitempty"`
	// Maximum size, in KB, of the request body. Warning: Setting this too low can increase the number of ongoing requests (depending on the value of 'Request concurrency'); this can cause Loki to complain about entries being delivered out of order.
	MaxPayloadSizeKB *int64 `json:"maxPayloadSizeKB,omitempty"`
	// Name of the event field that contains the message to send. If not specified, Stream sends a JSON representation of the whole event.
	Message *string `json:"message,omitempty"`
	// Which format to use when sending logs to Loki (Protobuf or JSON).  Defaults to Protobuf.
	MessageFormat *OutputLokiMessageFormat `json:"messageFormat,omitempty"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputLokiBackpressureBehavior `json:"onBackpressure,omitempty"`
	// Password (a.k.a API key in Grafana Cloud domain) for authentication
	Password *string `json:"password,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputLokiCompression `json:"pqCompress,omitempty"`
	PqControls *OutputLokiPqControls  `json:"pqControls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `json:"pqMaxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `json:"pqMaxSize,omitempty"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputLokiQueueFullBehavior `json:"pqOnBackpressure,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `json:"pqPath,omitempty"`
	// Toggle this off to forward new events to receiver(s) before queue is flushed. Otherwise, default drain behavior is FIFO (first in, first out).
	PqStrictOrdering *bool `json:"pqStrictOrdering,omitempty"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another trusted CA (e.g., the system's CA). Defaults to No.
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// List of headers that are safe to log in plain text.
	SafeHeaders []string `json:"safeHeaders,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported. These fields are added as labels to generated logs.
	SystemFields []string `json:"systemFields,omitempty"`
	// Select (or create) a stored text secret
	TextSecret *string `json:"textSecret,omitempty"`
	// Amount of time, in seconds, to wait for a request to complete before aborting it.
	TimeoutSec *int64 `json:"timeoutSec,omitempty"`
	// Bearer token to include in the authorization header. In Grafana Cloud, this is generally built by concatenating the username and the API key, separated by a colon. E.g.: <your-username>:<your-api-key>.
	Token *string        `json:"token,omitempty"`
	Type  OutputLokiType `json:"type"`
	// The endpoint to send logs to.
	URL string `json:"url"`
	// Enable to use round-robin DNS lookup. When a DNS server returns multiple addresses, this will cause Stream to cycle through them in the order returned.
	UseRoundRobinDNS *bool `json:"useRoundRobinDns,omitempty"`
	// Username for authentication
	Username *string `json:"username,omitempty"`
}