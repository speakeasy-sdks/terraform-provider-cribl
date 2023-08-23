// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type InputHTTPRawConnections struct {
	// Select a Destination.
	Output string `json:"output"`
	// Select Pipeline or Pack. Optional.
	Pipeline *string `json:"pipeline,omitempty"`
}

type InputHTTPRawMetadata struct {
	// Field name
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

// InputHTTPRawPqCompression - Codec to use to compress the persisted data.
type InputHTTPRawPqCompression string

const (
	InputHTTPRawPqCompressionNone InputHTTPRawPqCompression = "none"
	InputHTTPRawPqCompressionGzip InputHTTPRawPqCompression = "gzip"
)

func (e InputHTTPRawPqCompression) ToPointer() *InputHTTPRawPqCompression {
	return &e
}

func (e *InputHTTPRawPqCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputHTTPRawPqCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputHTTPRawPqCompression: %v", v)
	}
}

// InputHTTPRawPqMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputHTTPRawPqMode string

const (
	InputHTTPRawPqModeSmart  InputHTTPRawPqMode = "smart"
	InputHTTPRawPqModeAlways InputHTTPRawPqMode = "always"
)

func (e InputHTTPRawPqMode) ToPointer() *InputHTTPRawPqMode {
	return &e
}

func (e *InputHTTPRawPqMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputHTTPRawPqMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputHTTPRawPqMode: %v", v)
	}
}

type InputHTTPRawPq struct {
	// The number of events to send downstream before committing that Stream has read them.
	CommitFrequency *int64 `json:"commitFrequency,omitempty"`
	// Codec to use to compress the persisted data.
	Compress *InputHTTPRawPqCompression `json:"compress,omitempty"`
	// The maximum number of events to hold in memory before writing the events to disk.
	MaxBufferSize *int64 `json:"maxBufferSize,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	MaxFileSize *string `json:"maxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	MaxSize *string `json:"maxSize,omitempty"`
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputHTTPRawPqMode `json:"mode,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/inputs/<input-id>.
	Path *string `json:"path,omitempty"`
}

// InputHTTPRawTLSSettingsServerSideMaximumTLSVersion - Maximum TLS version to accept from connections.
type InputHTTPRawTLSSettingsServerSideMaximumTLSVersion string

const (
	InputHTTPRawTLSSettingsServerSideMaximumTLSVersionTlSv1  InputHTTPRawTLSSettingsServerSideMaximumTLSVersion = "TLSv1"
	InputHTTPRawTLSSettingsServerSideMaximumTLSVersionTlSv11 InputHTTPRawTLSSettingsServerSideMaximumTLSVersion = "TLSv1.1"
	InputHTTPRawTLSSettingsServerSideMaximumTLSVersionTlSv12 InputHTTPRawTLSSettingsServerSideMaximumTLSVersion = "TLSv1.2"
	InputHTTPRawTLSSettingsServerSideMaximumTLSVersionTlSv13 InputHTTPRawTLSSettingsServerSideMaximumTLSVersion = "TLSv1.3"
)

func (e InputHTTPRawTLSSettingsServerSideMaximumTLSVersion) ToPointer() *InputHTTPRawTLSSettingsServerSideMaximumTLSVersion {
	return &e
}

func (e *InputHTTPRawTLSSettingsServerSideMaximumTLSVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TLSv1":
		fallthrough
	case "TLSv1.1":
		fallthrough
	case "TLSv1.2":
		fallthrough
	case "TLSv1.3":
		*e = InputHTTPRawTLSSettingsServerSideMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputHTTPRawTLSSettingsServerSideMaximumTLSVersion: %v", v)
	}
}

// InputHTTPRawTLSSettingsServerSideMinimumTLSVersion - Minimum TLS version to accept from connections.
type InputHTTPRawTLSSettingsServerSideMinimumTLSVersion string

const (
	InputHTTPRawTLSSettingsServerSideMinimumTLSVersionTlSv1  InputHTTPRawTLSSettingsServerSideMinimumTLSVersion = "TLSv1"
	InputHTTPRawTLSSettingsServerSideMinimumTLSVersionTlSv11 InputHTTPRawTLSSettingsServerSideMinimumTLSVersion = "TLSv1.1"
	InputHTTPRawTLSSettingsServerSideMinimumTLSVersionTlSv12 InputHTTPRawTLSSettingsServerSideMinimumTLSVersion = "TLSv1.2"
	InputHTTPRawTLSSettingsServerSideMinimumTLSVersionTlSv13 InputHTTPRawTLSSettingsServerSideMinimumTLSVersion = "TLSv1.3"
)

func (e InputHTTPRawTLSSettingsServerSideMinimumTLSVersion) ToPointer() *InputHTTPRawTLSSettingsServerSideMinimumTLSVersion {
	return &e
}

func (e *InputHTTPRawTLSSettingsServerSideMinimumTLSVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TLSv1":
		fallthrough
	case "TLSv1.1":
		fallthrough
	case "TLSv1.2":
		fallthrough
	case "TLSv1.3":
		*e = InputHTTPRawTLSSettingsServerSideMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputHTTPRawTLSSettingsServerSideMinimumTLSVersion: %v", v)
	}
}

type InputHTTPRawTLSSettingsServerSide struct {
	// Path on server containing CA certificates to use. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Path on server containing certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// The name of the predefined certificate.
	CertificateName *string     `json:"certificateName,omitempty"`
	CommonNameRegex interface{} `json:"commonNameRegex,omitempty"`
	Disabled        *bool       `json:"disabled,omitempty"`
	// Maximum TLS version to accept from connections.
	MaxVersion *InputHTTPRawTLSSettingsServerSideMaximumTLSVersion `json:"maxVersion,omitempty"`
	// Minimum TLS version to accept from connections.
	MinVersion *InputHTTPRawTLSSettingsServerSideMinimumTLSVersion `json:"minVersion,omitempty"`
	// Passphrase to use to decrypt private key.
	Passphrase *string `json:"passphrase,omitempty"`
	// Path on server containing the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath        *string     `json:"privKeyPath,omitempty"`
	RejectUnauthorized interface{} `json:"rejectUnauthorized,omitempty"`
	// Whether to require clients to present their certificates. Used to perform client authentication using SSL certs.
	RequestCert *bool `json:"requestCert,omitempty"`
}

type InputHTTPRawType string

const (
	InputHTTPRawTypeHTTPRaw InputHTTPRawType = "http_raw"
)

func (e InputHTTPRawType) ToPointer() *InputHTTPRawType {
	return &e
}

func (e *InputHTTPRawType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "http_raw":
		*e = InputHTTPRawType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputHTTPRawType: %v", v)
	}
}

type InputHTTPRaw struct {
	// How often request activity is logged at the `info` level. A value of 1 would log every request, 10 every 10th request, etc.
	ActivityLogSampleRate *int64 `json:"activityLogSampleRate,omitempty"`
	// List of HTTP methods accepted by this input, wildcards are supported, e.g. P*, GET. Defaults to allow all.
	AllowedMethods []string `json:"allowedMethods,omitempty"`
	// List of URI paths accepted by this input, wildcards are supported, e.g /api/v*/hook. Defaults to allow all.
	AllowedPaths []string `json:"allowedPaths,omitempty"`
	// Shared secrets to be provided by any client (Authorization: <token>). If empty, unauthed access is permitted.
	AuthTokens []string `json:"authTokens,omitempty"`
	// A list of event breaking rulesets that will be applied, in order, to the input data stream.
	BreakerRulesets []string `json:"breakerRulesets,omitempty"`
	// Toggle this to Yes to add request headers to events, in the __headers field.
	CaptureHeaders *bool `json:"captureHeaders,omitempty"`
	// Direct connections to Destinations, optionally via a Pipeline or a Pack.
	Connections []InputHTTPRawConnections `json:"connections,omitempty"`
	// Enable/disable this input
	Disabled *bool `json:"disabled,omitempty"`
	// Enable if the connection is proxied by a device that supports Proxy Protocol V1 or V2.
	EnableProxyHeader *bool `json:"enableProxyHeader,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Address to bind on. Defaults to 0.0.0.0 (all addresses).
	Host string `json:"host"`
	// Unique ID for this input
	ID *string `json:"id,omitempty"`
	// After the last response is sent, @{product} will wait this long for additional data before closing the socket connection. Minimum 1 sec.; maximum 600 sec. (10 min.).
	KeepAliveTimeout *int64 `json:"keepAliveTimeout,omitempty"`
	// Maximum number of active requests per Worker Process. Use 0 for unlimited.
	MaxActiveReq *int64 `json:"maxActiveReq,omitempty"`
	// Fields to add to events from this input.
	Metadata []InputHTTPRawMetadata `json:"metadata,omitempty"`
	// Pipeline to process data from this Source before sending it through the Routes.
	Pipeline *string `json:"pipeline,omitempty"`
	// Port to listen on.
	Port int64           `json:"port"`
	Pq   *InputHTTPRawPq `json:"pq,omitempty"`
	// For details on Persistent Queues, see: [https://docs.cribl.io/stream/persistent-queues](https://docs.cribl.io/stream/persistent-queues)
	PqEnabled *bool `json:"pqEnabled,omitempty"`
	// How long to wait for an incoming request to complete before aborting it. Use 0 to disable.
	RequestTimeout *int64 `json:"requestTimeout,omitempty"`
	// Select whether to send data to Routes, or directly to Destinations.
	SendToRoutes *bool `json:"sendToRoutes,omitempty"`
	// How long @{product} should wait before assuming that an inactive socket has timed out. To wait forever, set to 0.
	SocketTimeout *int64 `json:"socketTimeout,omitempty"`
	// The amount of time (in milliseconds) the Event Breaker will wait for new data to be sent to a specific channel, before flushing the data stream out, as-is, to the Pipelines.
	StaleChannelFlushMs *int64 `json:"staleChannelFlushMs,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string                           `json:"streamtags,omitempty"`
	TLS        *InputHTTPRawTLSSettingsServerSide `json:"tls,omitempty"`
	Type       *InputHTTPRawType                  `json:"type,omitempty"`
}