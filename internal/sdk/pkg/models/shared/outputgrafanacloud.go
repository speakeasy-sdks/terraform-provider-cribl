// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type OutputGrafanaCloud2ExtraHTTPHeaders struct {
	// Field name
	Name *string `json:"name,omitempty"`
	// Field value
	Value string `json:"value"`
}

// OutputGrafanaCloud2FailedRequestLoggingMode - Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
type OutputGrafanaCloud2FailedRequestLoggingMode string

const (
	OutputGrafanaCloud2FailedRequestLoggingModePayload           OutputGrafanaCloud2FailedRequestLoggingMode = "payload"
	OutputGrafanaCloud2FailedRequestLoggingModePayloadAndHeaders OutputGrafanaCloud2FailedRequestLoggingMode = "payloadAndHeaders"
	OutputGrafanaCloud2FailedRequestLoggingModeNone              OutputGrafanaCloud2FailedRequestLoggingMode = "none"
)

func (e OutputGrafanaCloud2FailedRequestLoggingMode) ToPointer() *OutputGrafanaCloud2FailedRequestLoggingMode {
	return &e
}

func (e *OutputGrafanaCloud2FailedRequestLoggingMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputGrafanaCloud2FailedRequestLoggingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGrafanaCloud2FailedRequestLoggingMode: %v", v)
	}
}

type OutputGrafanaCloud2Labels struct {
	// Name of the label.
	Name string `json:"name"`
	// Value of the label.
	Value string `json:"value"`
}

// OutputGrafanaCloud2LokiAuthAuthenticationType - The authentication method to use for the HTTP requests
type OutputGrafanaCloud2LokiAuthAuthenticationType string

const (
	OutputGrafanaCloud2LokiAuthAuthenticationTypeCredentialsSecret OutputGrafanaCloud2LokiAuthAuthenticationType = "credentialsSecret"
	OutputGrafanaCloud2LokiAuthAuthenticationTypeToken             OutputGrafanaCloud2LokiAuthAuthenticationType = "token"
	OutputGrafanaCloud2LokiAuthAuthenticationTypeTextSecret        OutputGrafanaCloud2LokiAuthAuthenticationType = "textSecret"
	OutputGrafanaCloud2LokiAuthAuthenticationTypeBasic             OutputGrafanaCloud2LokiAuthAuthenticationType = "basic"
)

func (e OutputGrafanaCloud2LokiAuthAuthenticationType) ToPointer() *OutputGrafanaCloud2LokiAuthAuthenticationType {
	return &e
}

func (e *OutputGrafanaCloud2LokiAuthAuthenticationType) UnmarshalJSON(data []byte) error {
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
		*e = OutputGrafanaCloud2LokiAuthAuthenticationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGrafanaCloud2LokiAuthAuthenticationType: %v", v)
	}
}

type OutputGrafanaCloud2LokiAuth struct {
	// The authentication method to use for the HTTP requests
	AuthType *OutputGrafanaCloud2LokiAuthAuthenticationType `json:"authType,omitempty"`
	// Select (or create) a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
	// Password (a.k.a API key in Grafana Cloud domain) for authentication
	Password *string `json:"password,omitempty"`
	// Select (or create) a stored text secret
	TextSecret *string `json:"textSecret,omitempty"`
	// Bearer token to include in the authorization header. In Grafana Cloud, this is generally built by concatenating the username and the API key, separated by a colon. E.g.: <your-username>:<your-api-key>.
	Token *string `json:"token,omitempty"`
	// Username for authentication
	Username *string `json:"username,omitempty"`
}

// OutputGrafanaCloud2MessageFormat - Which format to use when sending logs to Loki (Protobuf or JSON).  Defaults to Protobuf.
type OutputGrafanaCloud2MessageFormat string

const (
	OutputGrafanaCloud2MessageFormatProtobuf OutputGrafanaCloud2MessageFormat = "protobuf"
	OutputGrafanaCloud2MessageFormatJSON     OutputGrafanaCloud2MessageFormat = "json"
)

func (e OutputGrafanaCloud2MessageFormat) ToPointer() *OutputGrafanaCloud2MessageFormat {
	return &e
}

func (e *OutputGrafanaCloud2MessageFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "protobuf":
		fallthrough
	case "json":
		*e = OutputGrafanaCloud2MessageFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGrafanaCloud2MessageFormat: %v", v)
	}
}

// OutputGrafanaCloud2BackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputGrafanaCloud2BackpressureBehavior string

const (
	OutputGrafanaCloud2BackpressureBehaviorQueue OutputGrafanaCloud2BackpressureBehavior = "queue"
	OutputGrafanaCloud2BackpressureBehaviorDrop  OutputGrafanaCloud2BackpressureBehavior = "drop"
	OutputGrafanaCloud2BackpressureBehaviorBlock OutputGrafanaCloud2BackpressureBehavior = "block"
)

func (e OutputGrafanaCloud2BackpressureBehavior) ToPointer() *OutputGrafanaCloud2BackpressureBehavior {
	return &e
}

func (e *OutputGrafanaCloud2BackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputGrafanaCloud2BackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGrafanaCloud2BackpressureBehavior: %v", v)
	}
}

// OutputGrafanaCloud2Compression - Codec to use to compress the persisted data.
type OutputGrafanaCloud2Compression string

const (
	OutputGrafanaCloud2CompressionNone OutputGrafanaCloud2Compression = "none"
	OutputGrafanaCloud2CompressionGzip OutputGrafanaCloud2Compression = "gzip"
)

func (e OutputGrafanaCloud2Compression) ToPointer() *OutputGrafanaCloud2Compression {
	return &e
}

func (e *OutputGrafanaCloud2Compression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputGrafanaCloud2Compression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGrafanaCloud2Compression: %v", v)
	}
}

type OutputGrafanaCloud2PqControls struct {
}

// OutputGrafanaCloud2QueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputGrafanaCloud2QueueFullBehavior string

const (
	OutputGrafanaCloud2QueueFullBehaviorBlock OutputGrafanaCloud2QueueFullBehavior = "block"
	OutputGrafanaCloud2QueueFullBehaviorDrop  OutputGrafanaCloud2QueueFullBehavior = "drop"
)

func (e OutputGrafanaCloud2QueueFullBehavior) ToPointer() *OutputGrafanaCloud2QueueFullBehavior {
	return &e
}

func (e *OutputGrafanaCloud2QueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputGrafanaCloud2QueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGrafanaCloud2QueueFullBehavior: %v", v)
	}
}

// OutputGrafanaCloud2PrometheusAuthAuthenticationType - The authentication method to use for the HTTP requests
type OutputGrafanaCloud2PrometheusAuthAuthenticationType string

const (
	OutputGrafanaCloud2PrometheusAuthAuthenticationTypeCredentialsSecret OutputGrafanaCloud2PrometheusAuthAuthenticationType = "credentialsSecret"
	OutputGrafanaCloud2PrometheusAuthAuthenticationTypeToken             OutputGrafanaCloud2PrometheusAuthAuthenticationType = "token"
	OutputGrafanaCloud2PrometheusAuthAuthenticationTypeTextSecret        OutputGrafanaCloud2PrometheusAuthAuthenticationType = "textSecret"
	OutputGrafanaCloud2PrometheusAuthAuthenticationTypeBasic             OutputGrafanaCloud2PrometheusAuthAuthenticationType = "basic"
)

func (e OutputGrafanaCloud2PrometheusAuthAuthenticationType) ToPointer() *OutputGrafanaCloud2PrometheusAuthAuthenticationType {
	return &e
}

func (e *OutputGrafanaCloud2PrometheusAuthAuthenticationType) UnmarshalJSON(data []byte) error {
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
		*e = OutputGrafanaCloud2PrometheusAuthAuthenticationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGrafanaCloud2PrometheusAuthAuthenticationType: %v", v)
	}
}

type OutputGrafanaCloud2PrometheusAuth struct {
	// The authentication method to use for the HTTP requests
	AuthType *OutputGrafanaCloud2PrometheusAuthAuthenticationType `json:"authType,omitempty"`
	// Select (or create) a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
	// Password (a.k.a API key in Grafana Cloud domain) for authentication
	Password *string `json:"password,omitempty"`
	// Select (or create) a stored text secret
	TextSecret *string `json:"textSecret,omitempty"`
	// Bearer token to include in the authorization header. In Grafana Cloud, this is generally built by concatenating the username and the API key, separated by a colon. E.g.: <your-username>:<your-api-key>.
	Token *string `json:"token,omitempty"`
	// Username for authentication
	Username *string `json:"username,omitempty"`
}

type OutputGrafanaCloud2Type string

const (
	OutputGrafanaCloud2TypeGrafanaCloud OutputGrafanaCloud2Type = "grafana_cloud"
)

func (e OutputGrafanaCloud2Type) ToPointer() *OutputGrafanaCloud2Type {
	return &e
}

func (e *OutputGrafanaCloud2Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "grafana_cloud":
		*e = OutputGrafanaCloud2Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGrafanaCloud2Type: %v", v)
	}
}

type OutputGrafanaCloud2 struct {
	// Whether to compress the payload body before sending. Applies only to Loki's JSON payloads, as both Prometheus' and Loki's Protobuf variant are snappy-compressed by default.
	Compress *bool `json:"compress,omitempty"`
	// Maximum number of ongoing requests before blocking. Warning: Setting this value > 1 can cause Loki and Prometheus to complain about entries being delivered out of order.
	Concurrency *int64 `json:"concurrency,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Headers to add to all events.
	ExtraHTTPHeaders []OutputGrafanaCloud2ExtraHTTPHeaders `json:"extraHttpHeaders,omitempty"`
	// Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
	FailedRequestLoggingMode *OutputGrafanaCloud2FailedRequestLoggingMode `json:"failedRequestLoggingMode,omitempty"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Maximum time between requests. Small values can reduce the payload size below the configured 'Max record size' and 'Max events per request'. Warning: Setting this too low can increase the number of ongoing requests (depending on the value of 'Request concurrency'); this can cause Loki and Prometheus to complain about entries being delivered out of order.
	FlushPeriodSec *int64 `json:"flushPeriodSec,omitempty"`
	// Unique ID for this output
	ID *string `json:"id,omitempty"`
	// List of labels to send with logs. Labels define Loki streams, so use static labels to avoid proliferating label value combinations and streams. Can be merged and/or overridden by the event's __labels field (e.g.: '__labels: {host: "cribl.io", level: "error"}').
	Labels   []OutputGrafanaCloud2Labels  `json:"labels,omitempty"`
	LokiAuth *OutputGrafanaCloud2LokiAuth `json:"lokiAuth,omitempty"`
	// The endpoint to send logs to, e.g.: https://logs-prod-us-central1.grafana.net
	LokiURL string `json:"lokiUrl"`
	// Maximum number of events to include in the request body. Default is 0 (unlimited). Warning: Setting this too low can increase the number of ongoing requests (depending on the value of 'Request concurrency'); this can cause Loki and Prometheus to complain about entries being delivered out of order.
	MaxPayloadEvents *int64 `json:"maxPayloadEvents,omitempty"`
	// Maximum size, in KB, of the request body. Warning: Setting this too low can increase the number of ongoing requests (depending on the value of 'Request concurrency'); this can cause Loki and Prometheus to complain about entries being delivered out of order.
	MaxPayloadSizeKB *int64 `json:"maxPayloadSizeKB,omitempty"`
	// Name of the event field that contains the message to send. If not specified, Stream sends a JSON representation of the whole event.
	Message *string `json:"message,omitempty"`
	// Which format to use when sending logs to Loki (Protobuf or JSON).  Defaults to Protobuf.
	MessageFormat *OutputGrafanaCloud2MessageFormat `json:"messageFormat,omitempty"`
	// A JS expression that can be used to rename metrics. E.g.: name.replace(/\./g, '_') will replace all '.' characters in a metric's name with the supported '_' character. Use the 'name' global variable to access the metric's name.  You can access event fields' values via __e.<fieldName>.
	MetricRenameExpr *string `json:"metricRenameExpr,omitempty"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputGrafanaCloud2BackpressureBehavior `json:"onBackpressure,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputGrafanaCloud2Compression `json:"pqCompress,omitempty"`
	PqControls *OutputGrafanaCloud2PqControls  `json:"pqControls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `json:"pqMaxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `json:"pqMaxSize,omitempty"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputGrafanaCloud2QueueFullBehavior `json:"pqOnBackpressure,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `json:"pqPath,omitempty"`
	// Toggle this off to forward new events to receiver(s) before queue is flushed. Otherwise, default drain behavior is FIFO (first in, first out).
	PqStrictOrdering *bool                              `json:"pqStrictOrdering,omitempty"`
	PrometheusAuth   *OutputGrafanaCloud2PrometheusAuth `json:"prometheusAuth,omitempty"`
	// The remote_write endpoint to send Prometheus metrics to, e.g.: https://prometheus-blocks-prod-us-central1.grafana.net/api/prom/push
	PrometheusURL string `json:"prometheusUrl"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another trusted CA (e.g., the system's CA). Defaults to Yes.
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// List of headers that are safe to log in plain text.
	SafeHeaders []string `json:"safeHeaders,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported. These fields are added as dimensions and labels to generated metrics and logs respectively.
	SystemFields []string `json:"systemFields,omitempty"`
	// Amount of time, in seconds, to wait for a request to complete before aborting it.
	TimeoutSec *int64                   `json:"timeoutSec,omitempty"`
	Type       *OutputGrafanaCloud2Type `json:"type,omitempty"`
	// Enable to use round-robin DNS lookup. When a DNS server returns multiple addresses, this will cause Stream to cycle through them in the order returned.
	UseRoundRobinDNS *bool `json:"useRoundRobinDns,omitempty"`
}

type OutputGrafanaCloud1ExtraHTTPHeaders struct {
	// Field name
	Name *string `json:"name,omitempty"`
	// Field value
	Value string `json:"value"`
}

// OutputGrafanaCloud1FailedRequestLoggingMode - Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
type OutputGrafanaCloud1FailedRequestLoggingMode string

const (
	OutputGrafanaCloud1FailedRequestLoggingModePayload           OutputGrafanaCloud1FailedRequestLoggingMode = "payload"
	OutputGrafanaCloud1FailedRequestLoggingModePayloadAndHeaders OutputGrafanaCloud1FailedRequestLoggingMode = "payloadAndHeaders"
	OutputGrafanaCloud1FailedRequestLoggingModeNone              OutputGrafanaCloud1FailedRequestLoggingMode = "none"
)

func (e OutputGrafanaCloud1FailedRequestLoggingMode) ToPointer() *OutputGrafanaCloud1FailedRequestLoggingMode {
	return &e
}

func (e *OutputGrafanaCloud1FailedRequestLoggingMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputGrafanaCloud1FailedRequestLoggingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGrafanaCloud1FailedRequestLoggingMode: %v", v)
	}
}

type OutputGrafanaCloud1Labels struct {
	// Name of the label.
	Name string `json:"name"`
	// Value of the label.
	Value string `json:"value"`
}

// OutputGrafanaCloud1LokiAuthAuthenticationType - The authentication method to use for the HTTP requests
type OutputGrafanaCloud1LokiAuthAuthenticationType string

const (
	OutputGrafanaCloud1LokiAuthAuthenticationTypeCredentialsSecret OutputGrafanaCloud1LokiAuthAuthenticationType = "credentialsSecret"
	OutputGrafanaCloud1LokiAuthAuthenticationTypeToken             OutputGrafanaCloud1LokiAuthAuthenticationType = "token"
	OutputGrafanaCloud1LokiAuthAuthenticationTypeTextSecret        OutputGrafanaCloud1LokiAuthAuthenticationType = "textSecret"
	OutputGrafanaCloud1LokiAuthAuthenticationTypeBasic             OutputGrafanaCloud1LokiAuthAuthenticationType = "basic"
)

func (e OutputGrafanaCloud1LokiAuthAuthenticationType) ToPointer() *OutputGrafanaCloud1LokiAuthAuthenticationType {
	return &e
}

func (e *OutputGrafanaCloud1LokiAuthAuthenticationType) UnmarshalJSON(data []byte) error {
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
		*e = OutputGrafanaCloud1LokiAuthAuthenticationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGrafanaCloud1LokiAuthAuthenticationType: %v", v)
	}
}

type OutputGrafanaCloud1LokiAuth struct {
	// The authentication method to use for the HTTP requests
	AuthType *OutputGrafanaCloud1LokiAuthAuthenticationType `json:"authType,omitempty"`
	// Select (or create) a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
	// Password (a.k.a API key in Grafana Cloud domain) for authentication
	Password *string `json:"password,omitempty"`
	// Select (or create) a stored text secret
	TextSecret *string `json:"textSecret,omitempty"`
	// Bearer token to include in the authorization header. In Grafana Cloud, this is generally built by concatenating the username and the API key, separated by a colon. E.g.: <your-username>:<your-api-key>.
	Token *string `json:"token,omitempty"`
	// Username for authentication
	Username *string `json:"username,omitempty"`
}

// OutputGrafanaCloud1MessageFormat - Which format to use when sending logs to Loki (Protobuf or JSON).  Defaults to Protobuf.
type OutputGrafanaCloud1MessageFormat string

const (
	OutputGrafanaCloud1MessageFormatProtobuf OutputGrafanaCloud1MessageFormat = "protobuf"
	OutputGrafanaCloud1MessageFormatJSON     OutputGrafanaCloud1MessageFormat = "json"
)

func (e OutputGrafanaCloud1MessageFormat) ToPointer() *OutputGrafanaCloud1MessageFormat {
	return &e
}

func (e *OutputGrafanaCloud1MessageFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "protobuf":
		fallthrough
	case "json":
		*e = OutputGrafanaCloud1MessageFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGrafanaCloud1MessageFormat: %v", v)
	}
}

// OutputGrafanaCloud1BackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputGrafanaCloud1BackpressureBehavior string

const (
	OutputGrafanaCloud1BackpressureBehaviorQueue OutputGrafanaCloud1BackpressureBehavior = "queue"
	OutputGrafanaCloud1BackpressureBehaviorDrop  OutputGrafanaCloud1BackpressureBehavior = "drop"
	OutputGrafanaCloud1BackpressureBehaviorBlock OutputGrafanaCloud1BackpressureBehavior = "block"
)

func (e OutputGrafanaCloud1BackpressureBehavior) ToPointer() *OutputGrafanaCloud1BackpressureBehavior {
	return &e
}

func (e *OutputGrafanaCloud1BackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputGrafanaCloud1BackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGrafanaCloud1BackpressureBehavior: %v", v)
	}
}

// OutputGrafanaCloud1Compression - Codec to use to compress the persisted data.
type OutputGrafanaCloud1Compression string

const (
	OutputGrafanaCloud1CompressionNone OutputGrafanaCloud1Compression = "none"
	OutputGrafanaCloud1CompressionGzip OutputGrafanaCloud1Compression = "gzip"
)

func (e OutputGrafanaCloud1Compression) ToPointer() *OutputGrafanaCloud1Compression {
	return &e
}

func (e *OutputGrafanaCloud1Compression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputGrafanaCloud1Compression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGrafanaCloud1Compression: %v", v)
	}
}

type OutputGrafanaCloud1PqControls struct {
}

// OutputGrafanaCloud1QueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputGrafanaCloud1QueueFullBehavior string

const (
	OutputGrafanaCloud1QueueFullBehaviorBlock OutputGrafanaCloud1QueueFullBehavior = "block"
	OutputGrafanaCloud1QueueFullBehaviorDrop  OutputGrafanaCloud1QueueFullBehavior = "drop"
)

func (e OutputGrafanaCloud1QueueFullBehavior) ToPointer() *OutputGrafanaCloud1QueueFullBehavior {
	return &e
}

func (e *OutputGrafanaCloud1QueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputGrafanaCloud1QueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGrafanaCloud1QueueFullBehavior: %v", v)
	}
}

// OutputGrafanaCloud1PrometheusAuthAuthenticationType - The authentication method to use for the HTTP requests
type OutputGrafanaCloud1PrometheusAuthAuthenticationType string

const (
	OutputGrafanaCloud1PrometheusAuthAuthenticationTypeCredentialsSecret OutputGrafanaCloud1PrometheusAuthAuthenticationType = "credentialsSecret"
	OutputGrafanaCloud1PrometheusAuthAuthenticationTypeToken             OutputGrafanaCloud1PrometheusAuthAuthenticationType = "token"
	OutputGrafanaCloud1PrometheusAuthAuthenticationTypeTextSecret        OutputGrafanaCloud1PrometheusAuthAuthenticationType = "textSecret"
	OutputGrafanaCloud1PrometheusAuthAuthenticationTypeBasic             OutputGrafanaCloud1PrometheusAuthAuthenticationType = "basic"
)

func (e OutputGrafanaCloud1PrometheusAuthAuthenticationType) ToPointer() *OutputGrafanaCloud1PrometheusAuthAuthenticationType {
	return &e
}

func (e *OutputGrafanaCloud1PrometheusAuthAuthenticationType) UnmarshalJSON(data []byte) error {
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
		*e = OutputGrafanaCloud1PrometheusAuthAuthenticationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGrafanaCloud1PrometheusAuthAuthenticationType: %v", v)
	}
}

type OutputGrafanaCloud1PrometheusAuth struct {
	// The authentication method to use for the HTTP requests
	AuthType *OutputGrafanaCloud1PrometheusAuthAuthenticationType `json:"authType,omitempty"`
	// Select (or create) a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
	// Password (a.k.a API key in Grafana Cloud domain) for authentication
	Password *string `json:"password,omitempty"`
	// Select (or create) a stored text secret
	TextSecret *string `json:"textSecret,omitempty"`
	// Bearer token to include in the authorization header. In Grafana Cloud, this is generally built by concatenating the username and the API key, separated by a colon. E.g.: <your-username>:<your-api-key>.
	Token *string `json:"token,omitempty"`
	// Username for authentication
	Username *string `json:"username,omitempty"`
}

type OutputGrafanaCloud1Type string

const (
	OutputGrafanaCloud1TypeGrafanaCloud OutputGrafanaCloud1Type = "grafana_cloud"
)

func (e OutputGrafanaCloud1Type) ToPointer() *OutputGrafanaCloud1Type {
	return &e
}

func (e *OutputGrafanaCloud1Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "grafana_cloud":
		*e = OutputGrafanaCloud1Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGrafanaCloud1Type: %v", v)
	}
}

type OutputGrafanaCloud1 struct {
	// Whether to compress the payload body before sending. Applies only to Loki's JSON payloads, as both Prometheus' and Loki's Protobuf variant are snappy-compressed by default.
	Compress *bool `json:"compress,omitempty"`
	// Maximum number of ongoing requests before blocking. Warning: Setting this value > 1 can cause Loki and Prometheus to complain about entries being delivered out of order.
	Concurrency *int64 `json:"concurrency,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Headers to add to all events.
	ExtraHTTPHeaders []OutputGrafanaCloud1ExtraHTTPHeaders `json:"extraHttpHeaders,omitempty"`
	// Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
	FailedRequestLoggingMode *OutputGrafanaCloud1FailedRequestLoggingMode `json:"failedRequestLoggingMode,omitempty"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Maximum time between requests. Small values can reduce the payload size below the configured 'Max record size' and 'Max events per request'. Warning: Setting this too low can increase the number of ongoing requests (depending on the value of 'Request concurrency'); this can cause Loki and Prometheus to complain about entries being delivered out of order.
	FlushPeriodSec *int64 `json:"flushPeriodSec,omitempty"`
	// Unique ID for this output
	ID *string `json:"id,omitempty"`
	// List of labels to send with logs. Labels define Loki streams, so use static labels to avoid proliferating label value combinations and streams. Can be merged and/or overridden by the event's __labels field (e.g.: '__labels: {host: "cribl.io", level: "error"}').
	Labels   []OutputGrafanaCloud1Labels  `json:"labels,omitempty"`
	LokiAuth *OutputGrafanaCloud1LokiAuth `json:"lokiAuth,omitempty"`
	// The endpoint to send logs to, e.g.: https://logs-prod-us-central1.grafana.net
	LokiURL string `json:"lokiUrl"`
	// Maximum number of events to include in the request body. Default is 0 (unlimited). Warning: Setting this too low can increase the number of ongoing requests (depending on the value of 'Request concurrency'); this can cause Loki and Prometheus to complain about entries being delivered out of order.
	MaxPayloadEvents *int64 `json:"maxPayloadEvents,omitempty"`
	// Maximum size, in KB, of the request body. Warning: Setting this too low can increase the number of ongoing requests (depending on the value of 'Request concurrency'); this can cause Loki and Prometheus to complain about entries being delivered out of order.
	MaxPayloadSizeKB *int64 `json:"maxPayloadSizeKB,omitempty"`
	// Name of the event field that contains the message to send. If not specified, Stream sends a JSON representation of the whole event.
	Message *string `json:"message,omitempty"`
	// Which format to use when sending logs to Loki (Protobuf or JSON).  Defaults to Protobuf.
	MessageFormat *OutputGrafanaCloud1MessageFormat `json:"messageFormat,omitempty"`
	// A JS expression that can be used to rename metrics. E.g.: name.replace(/\./g, '_') will replace all '.' characters in a metric's name with the supported '_' character. Use the 'name' global variable to access the metric's name.  You can access event fields' values via __e.<fieldName>.
	MetricRenameExpr *string `json:"metricRenameExpr,omitempty"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputGrafanaCloud1BackpressureBehavior `json:"onBackpressure,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputGrafanaCloud1Compression `json:"pqCompress,omitempty"`
	PqControls *OutputGrafanaCloud1PqControls  `json:"pqControls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `json:"pqMaxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `json:"pqMaxSize,omitempty"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputGrafanaCloud1QueueFullBehavior `json:"pqOnBackpressure,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `json:"pqPath,omitempty"`
	// Toggle this off to forward new events to receiver(s) before queue is flushed. Otherwise, default drain behavior is FIFO (first in, first out).
	PqStrictOrdering *bool                              `json:"pqStrictOrdering,omitempty"`
	PrometheusAuth   *OutputGrafanaCloud1PrometheusAuth `json:"prometheusAuth,omitempty"`
	// The remote_write endpoint to send Prometheus metrics to, e.g.: https://prometheus-blocks-prod-us-central1.grafana.net/api/prom/push
	PrometheusURL string `json:"prometheusUrl"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another trusted CA (e.g., the system's CA). Defaults to Yes.
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// List of headers that are safe to log in plain text.
	SafeHeaders []string `json:"safeHeaders,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported. These fields are added as dimensions and labels to generated metrics and logs respectively.
	SystemFields []string `json:"systemFields,omitempty"`
	// Amount of time, in seconds, to wait for a request to complete before aborting it.
	TimeoutSec *int64                   `json:"timeoutSec,omitempty"`
	Type       *OutputGrafanaCloud1Type `json:"type,omitempty"`
	// Enable to use round-robin DNS lookup. When a DNS server returns multiple addresses, this will cause Stream to cycle through them in the order returned.
	UseRoundRobinDNS *bool `json:"useRoundRobinDns,omitempty"`
}

type OutputGrafanaCloudType string

const (
	OutputGrafanaCloudTypeOutputGrafanaCloud1 OutputGrafanaCloudType = "OutputGrafanaCloud_1"
	OutputGrafanaCloudTypeOutputGrafanaCloud2 OutputGrafanaCloudType = "OutputGrafanaCloud_2"
)

type OutputGrafanaCloud struct {
	OutputGrafanaCloud1 *OutputGrafanaCloud1
	OutputGrafanaCloud2 *OutputGrafanaCloud2

	Type OutputGrafanaCloudType
}

func CreateOutputGrafanaCloudOutputGrafanaCloud1(outputGrafanaCloud1 OutputGrafanaCloud1) OutputGrafanaCloud {
	typ := OutputGrafanaCloudTypeOutputGrafanaCloud1

	return OutputGrafanaCloud{
		OutputGrafanaCloud1: &outputGrafanaCloud1,
		Type:                typ,
	}
}

func CreateOutputGrafanaCloudOutputGrafanaCloud2(outputGrafanaCloud2 OutputGrafanaCloud2) OutputGrafanaCloud {
	typ := OutputGrafanaCloudTypeOutputGrafanaCloud2

	return OutputGrafanaCloud{
		OutputGrafanaCloud2: &outputGrafanaCloud2,
		Type:                typ,
	}
}

func (u *OutputGrafanaCloud) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	outputGrafanaCloud1 := new(OutputGrafanaCloud1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&outputGrafanaCloud1); err == nil {
		u.OutputGrafanaCloud1 = outputGrafanaCloud1
		u.Type = OutputGrafanaCloudTypeOutputGrafanaCloud1
		return nil
	}

	outputGrafanaCloud2 := new(OutputGrafanaCloud2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&outputGrafanaCloud2); err == nil {
		u.OutputGrafanaCloud2 = outputGrafanaCloud2
		u.Type = OutputGrafanaCloudTypeOutputGrafanaCloud2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u OutputGrafanaCloud) MarshalJSON() ([]byte, error) {
	if u.OutputGrafanaCloud1 != nil {
		return json.Marshal(u.OutputGrafanaCloud1)
	}

	if u.OutputGrafanaCloud2 != nil {
		return json.Marshal(u.OutputGrafanaCloud2)
	}

	return nil, nil
}
