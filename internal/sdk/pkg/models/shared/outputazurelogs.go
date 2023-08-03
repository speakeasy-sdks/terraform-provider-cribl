// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OutputAzureLogsAuthenticationMethod - Enter workspace ID and workspace key directly, or select a stored secret
type OutputAzureLogsAuthenticationMethod string

const (
	OutputAzureLogsAuthenticationMethodSecret OutputAzureLogsAuthenticationMethod = "secret"
	OutputAzureLogsAuthenticationMethodManual OutputAzureLogsAuthenticationMethod = "manual"
)

func (e OutputAzureLogsAuthenticationMethod) ToPointer() *OutputAzureLogsAuthenticationMethod {
	return &e
}

func (e *OutputAzureLogsAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "secret":
		fallthrough
	case "manual":
		*e = OutputAzureLogsAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputAzureLogsAuthenticationMethod: %v", v)
	}
}

type OutputAzureLogsExtraHTTPHeaders struct {
	// Field name
	Name *string `json:"name,omitempty"`
	// Field value
	Value string `json:"value"`
}

// OutputAzureLogsFailedRequestLoggingMode - Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
type OutputAzureLogsFailedRequestLoggingMode string

const (
	OutputAzureLogsFailedRequestLoggingModePayload           OutputAzureLogsFailedRequestLoggingMode = "payload"
	OutputAzureLogsFailedRequestLoggingModePayloadAndHeaders OutputAzureLogsFailedRequestLoggingMode = "payloadAndHeaders"
	OutputAzureLogsFailedRequestLoggingModeNone              OutputAzureLogsFailedRequestLoggingMode = "none"
)

func (e OutputAzureLogsFailedRequestLoggingMode) ToPointer() *OutputAzureLogsFailedRequestLoggingMode {
	return &e
}

func (e *OutputAzureLogsFailedRequestLoggingMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputAzureLogsFailedRequestLoggingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputAzureLogsFailedRequestLoggingMode: %v", v)
	}
}

// OutputAzureLogsBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputAzureLogsBackpressureBehavior string

const (
	OutputAzureLogsBackpressureBehaviorQueue OutputAzureLogsBackpressureBehavior = "queue"
	OutputAzureLogsBackpressureBehaviorDrop  OutputAzureLogsBackpressureBehavior = "drop"
	OutputAzureLogsBackpressureBehaviorBlock OutputAzureLogsBackpressureBehavior = "block"
)

func (e OutputAzureLogsBackpressureBehavior) ToPointer() *OutputAzureLogsBackpressureBehavior {
	return &e
}

func (e *OutputAzureLogsBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputAzureLogsBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputAzureLogsBackpressureBehavior: %v", v)
	}
}

// OutputAzureLogsCompression - Codec to use to compress the persisted data.
type OutputAzureLogsCompression string

const (
	OutputAzureLogsCompressionNone OutputAzureLogsCompression = "none"
	OutputAzureLogsCompressionGzip OutputAzureLogsCompression = "gzip"
)

func (e OutputAzureLogsCompression) ToPointer() *OutputAzureLogsCompression {
	return &e
}

func (e *OutputAzureLogsCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputAzureLogsCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputAzureLogsCompression: %v", v)
	}
}

type OutputAzureLogsPqControls struct {
}

// OutputAzureLogsQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputAzureLogsQueueFullBehavior string

const (
	OutputAzureLogsQueueFullBehaviorBlock OutputAzureLogsQueueFullBehavior = "block"
	OutputAzureLogsQueueFullBehaviorDrop  OutputAzureLogsQueueFullBehavior = "drop"
)

func (e OutputAzureLogsQueueFullBehavior) ToPointer() *OutputAzureLogsQueueFullBehavior {
	return &e
}

func (e *OutputAzureLogsQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputAzureLogsQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputAzureLogsQueueFullBehavior: %v", v)
	}
}

type OutputAzureLogsType string

const (
	OutputAzureLogsTypeAzureLogs OutputAzureLogsType = "azure_logs"
)

func (e OutputAzureLogsType) ToPointer() *OutputAzureLogsType {
	return &e
}

func (e *OutputAzureLogsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "azure_logs":
		*e = OutputAzureLogsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputAzureLogsType: %v", v)
	}
}

type OutputAzureLogs struct {
	// Enter the DNS name of the Log API endpoint that sends log data to a Log Analytics workspace in Azure Monitor. Defaults to .ods.opinsights.azure.com. @{product} will add a prefix and suffix around this DNS name to construct a URI in this format: <https://<Workspace_ID><your_DNS_name>/api/logs?api-version=<API version>.
	APIURL *string `json:"apiUrl,omitempty"`
	// Enter workspace ID and workspace key directly, or select a stored secret
	AuthType *OutputAzureLogsAuthenticationMethod `json:"authType,omitempty"`
	// Whether to compress the payload body before sending.
	Compress *bool `json:"compress,omitempty"`
	// Maximum number of ongoing requests before blocking.
	Concurrency *int64 `json:"concurrency,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Headers to add to all events.
	ExtraHTTPHeaders []OutputAzureLogsExtraHTTPHeaders `json:"extraHttpHeaders,omitempty"`
	// Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
	FailedRequestLoggingMode *OutputAzureLogsFailedRequestLoggingMode `json:"failedRequestLoggingMode,omitempty"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Max body size.
	FlushPeriodSec *int64 `json:"flushPeriodSec,omitempty"`
	// Unique ID for this output
	ID *string `json:"id,omitempty"`
	// Select (or create) a stored secret that references your access key and secret key.
	KeypairSecret *string `json:"keypairSecret,omitempty"`
	// The Log Type of events sent to this LogAnalytics workspace. Defaults to `Cribl`. Use only letters, numbers, and `_` characters, and can't exceed 100 characters. Can be overwritten by event field __logType.
	LogType string `json:"logType"`
	// Max number of events to include in the request body. Default is 0 (unlimited).
	MaxPayloadEvents *int64 `json:"maxPayloadEvents,omitempty"`
	// Maximum size, in KB, of the request body.
	MaxPayloadSizeKB *int64 `json:"maxPayloadSizeKB,omitempty"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputAzureLogsBackpressureBehavior `json:"onBackpressure,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputAzureLogsCompression `json:"pqCompress,omitempty"`
	PqControls *OutputAzureLogsPqControls  `json:"pqControls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `json:"pqMaxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `json:"pqMaxSize,omitempty"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputAzureLogsQueueFullBehavior `json:"pqOnBackpressure,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `json:"pqPath,omitempty"`
	// Toggle this off to forward new events to receiver(s) before queue is flushed. Otherwise, default drain behavior is FIFO (first in, first out).
	PqStrictOrdering *bool `json:"pqStrictOrdering,omitempty"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another trusted CA (e.g., the system's CA). Defaults to Yes.
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// Optional Resource ID of the Azure resource to associate the data with. Can be overridden by the __resourceId event field. This ID populates the _ResourceId property, allowing the data to be included in resource-centric queries. If the ID is neither specified nor overridden, resource-centric queries will omit the data.
	ResourceID *string `json:"resourceId,omitempty"`
	// List of headers that are safe to log in plain text.
	SafeHeaders []string `json:"safeHeaders,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string `json:"systemFields,omitempty"`
	// Amount of time, in seconds, to wait for a request to complete before aborting it.
	TimeoutSec *int64              `json:"timeoutSec,omitempty"`
	Type       OutputAzureLogsType `json:"type"`
	// Enable to use round-robin DNS lookup. When a DNS server returns multiple addresses, this will cause Stream to cycle through them in the order returned.
	UseRoundRobinDNS *bool `json:"useRoundRobinDns,omitempty"`
	// Azure Log Analytics Workspace ID. See Azure Dashboard Workspace > Advanced settings.
	WorkspaceID *string `json:"workspaceId,omitempty"`
	// Azure Log Analytics Workspace Primary or Secondary Shared Key. See Azure Dashboard Workspace > Advanced settings.
	WorkspaceKey *string `json:"workspaceKey,omitempty"`
}