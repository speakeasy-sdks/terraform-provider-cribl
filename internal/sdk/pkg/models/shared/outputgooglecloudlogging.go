// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OutputGoogleCloudLoggingAuthenticationMethod - Google authentication method. Choose Auto to use environment variable GOOGLE_APPLICATION_CREDENTIALS..
type OutputGoogleCloudLoggingAuthenticationMethod string

const (
	OutputGoogleCloudLoggingAuthenticationMethodSecret OutputGoogleCloudLoggingAuthenticationMethod = "secret"
	OutputGoogleCloudLoggingAuthenticationMethodManual OutputGoogleCloudLoggingAuthenticationMethod = "manual"
)

func (e OutputGoogleCloudLoggingAuthenticationMethod) ToPointer() *OutputGoogleCloudLoggingAuthenticationMethod {
	return &e
}

func (e *OutputGoogleCloudLoggingAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "secret":
		fallthrough
	case "manual":
		*e = OutputGoogleCloudLoggingAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGoogleCloudLoggingAuthenticationMethod: %v", v)
	}
}

type OutputGoogleCloudLoggingLogLabels struct {
	// Label name
	Label string `json:"label"`
	// JavaScript expression to compute the label's value.
	ValueExpression string `json:"valueExpression"`
}

type OutputGoogleCloudLoggingLogLocationType string

const (
	OutputGoogleCloudLoggingLogLocationTypeFolder         OutputGoogleCloudLoggingLogLocationType = "folder"
	OutputGoogleCloudLoggingLogLocationTypeOrganization   OutputGoogleCloudLoggingLogLocationType = "organization"
	OutputGoogleCloudLoggingLogLocationTypeBillingAccount OutputGoogleCloudLoggingLogLocationType = "billingAccount"
)

func (e OutputGoogleCloudLoggingLogLocationType) ToPointer() *OutputGoogleCloudLoggingLogLocationType {
	return &e
}

func (e *OutputGoogleCloudLoggingLogLocationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "folder":
		fallthrough
	case "organization":
		fallthrough
	case "billingAccount":
		*e = OutputGoogleCloudLoggingLogLocationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGoogleCloudLoggingLogLocationType: %v", v)
	}
}

// OutputGoogleCloudLoggingBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputGoogleCloudLoggingBackpressureBehavior string

const (
	OutputGoogleCloudLoggingBackpressureBehaviorQueue OutputGoogleCloudLoggingBackpressureBehavior = "queue"
	OutputGoogleCloudLoggingBackpressureBehaviorDrop  OutputGoogleCloudLoggingBackpressureBehavior = "drop"
	OutputGoogleCloudLoggingBackpressureBehaviorBlock OutputGoogleCloudLoggingBackpressureBehavior = "block"
)

func (e OutputGoogleCloudLoggingBackpressureBehavior) ToPointer() *OutputGoogleCloudLoggingBackpressureBehavior {
	return &e
}

func (e *OutputGoogleCloudLoggingBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputGoogleCloudLoggingBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGoogleCloudLoggingBackpressureBehavior: %v", v)
	}
}

// OutputGoogleCloudLoggingPayloadFormat - Format to use when sending payload. Defaults to Text.
type OutputGoogleCloudLoggingPayloadFormat string

const (
	OutputGoogleCloudLoggingPayloadFormatJSON OutputGoogleCloudLoggingPayloadFormat = "json"
	OutputGoogleCloudLoggingPayloadFormatText OutputGoogleCloudLoggingPayloadFormat = "text"
)

func (e OutputGoogleCloudLoggingPayloadFormat) ToPointer() *OutputGoogleCloudLoggingPayloadFormat {
	return &e
}

func (e *OutputGoogleCloudLoggingPayloadFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "text":
		*e = OutputGoogleCloudLoggingPayloadFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGoogleCloudLoggingPayloadFormat: %v", v)
	}
}

// OutputGoogleCloudLoggingCompression - Codec to use to compress the persisted data.
type OutputGoogleCloudLoggingCompression string

const (
	OutputGoogleCloudLoggingCompressionNone OutputGoogleCloudLoggingCompression = "none"
	OutputGoogleCloudLoggingCompressionGzip OutputGoogleCloudLoggingCompression = "gzip"
)

func (e OutputGoogleCloudLoggingCompression) ToPointer() *OutputGoogleCloudLoggingCompression {
	return &e
}

func (e *OutputGoogleCloudLoggingCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputGoogleCloudLoggingCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGoogleCloudLoggingCompression: %v", v)
	}
}

type OutputGoogleCloudLoggingPqControls struct {
}

// OutputGoogleCloudLoggingQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputGoogleCloudLoggingQueueFullBehavior string

const (
	OutputGoogleCloudLoggingQueueFullBehaviorBlock OutputGoogleCloudLoggingQueueFullBehavior = "block"
	OutputGoogleCloudLoggingQueueFullBehaviorDrop  OutputGoogleCloudLoggingQueueFullBehavior = "drop"
)

func (e OutputGoogleCloudLoggingQueueFullBehavior) ToPointer() *OutputGoogleCloudLoggingQueueFullBehavior {
	return &e
}

func (e *OutputGoogleCloudLoggingQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputGoogleCloudLoggingQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGoogleCloudLoggingQueueFullBehavior: %v", v)
	}
}

type OutputGoogleCloudLoggingResourceTypeLabels struct {
	// Label name
	Label string `json:"label"`
	// JavaScript expression to compute the label's value.
	ValueExpression string `json:"valueExpression"`
}

type OutputGoogleCloudLoggingType string

const (
	OutputGoogleCloudLoggingTypeGoogleCloudLogging OutputGoogleCloudLoggingType = "google_cloud_logging"
)

func (e OutputGoogleCloudLoggingType) ToPointer() *OutputGoogleCloudLoggingType {
	return &e
}

func (e *OutputGoogleCloudLoggingType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "google_cloud_logging":
		*e = OutputGoogleCloudLoggingType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGoogleCloudLoggingType: %v", v)
	}
}

type OutputGoogleCloudLogging struct {
	// Maximum number of ongoing requests before blocking.
	Concurrency *int64 `json:"concurrency,omitempty"`
	// Amount of time (milliseconds) to wait for the connection to establish before retrying
	ConnectionTimeout *int64 `json:"connectionTimeout,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Max record size.
	FlushPeriodSec *int64 `json:"flushPeriodSec,omitempty"`
	// Google authentication method. Choose Auto to use environment variable GOOGLE_APPLICATION_CREDENTIALS..
	GoogleAuthMethod *OutputGoogleCloudLoggingAuthenticationMethod `json:"googleAuthMethod,omitempty"`
	// Unique ID for this output
	ID *string `json:"id,omitempty"`
	// JavaScript expression to compute the value of the insert ID field.
	InsertIDExpression *string `json:"insertIdExpression,omitempty"`
	// Labels to apply to the log entry
	LogLabels []OutputGoogleCloudLoggingLogLabels `json:"logLabels,omitempty"`
	// JavaScript expression to compute the value of the folder ID with which log entries should be associated.
	LogLocationExpression string                                  `json:"logLocationExpression"`
	LogLocationType       OutputGoogleCloudLoggingLogLocationType `json:"logLocationType"`
	// JavaScript expression to compute the value of the log name.
	LogNameExpression string `json:"logNameExpression"`
	// Max number of events to include in the request body. Default is 0 (unlimited).
	MaxPayloadEvents *int64 `json:"maxPayloadEvents,omitempty"`
	// Maximum size, in KB, of the request body.
	MaxPayloadSizeKB *int64 `json:"maxPayloadSizeKB,omitempty"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputGoogleCloudLoggingBackpressureBehavior `json:"onBackpressure,omitempty"`
	// JavaScript expression to compute the value of the payload. Must evaluate to a JavaScript object value. If an invalid value is encountered it will result in the default value instead. Defaults to the entire event.
	PayloadExpression *string `json:"payloadExpression,omitempty"`
	// Format to use when sending payload. Defaults to Text.
	PayloadFormat *OutputGoogleCloudLoggingPayloadFormat `json:"payloadFormat,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputGoogleCloudLoggingCompression `json:"pqCompress,omitempty"`
	PqControls *OutputGoogleCloudLoggingPqControls  `json:"pqControls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `json:"pqMaxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `json:"pqMaxSize,omitempty"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputGoogleCloudLoggingQueueFullBehavior `json:"pqOnBackpressure,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `json:"pqPath,omitempty"`
	// Toggle this off to forward new events to receiver(s) before queue is flushed. Otherwise, default drain behavior is FIFO (first in, first out).
	PqStrictOrdering *bool `json:"pqStrictOrdering,omitempty"`
	// JavaScript expression to compute the value of the managed resource type field. Must evaluate to one of the valid values [here](https://cloud.google.com/logging/docs/api/v2/resource-list#resource-types). Defaults to "global".
	ResourceTypeExpression *string `json:"resourceTypeExpression,omitempty"`
	// Labels to apply to the managed resource. These must correspond to the valid labels for the specified resource type (see [here](https://cloud.google.com/logging/docs/api/v2/resource-list#resource-types)). Otherwise, they will be dropped by Google Cloud Logging.
	ResourceTypeLabels []OutputGoogleCloudLoggingResourceTypeLabels `json:"resourceTypeLabels,omitempty"`
	// Select (or create) a stored text secret
	Secret *string `json:"secret,omitempty"`
	// Contents of service account credentials (JSON keys) file downloaded from Google Cloud. To upload a file, click the upload button at this field's upper right. As an alternative, you can use environment variables (see [here](https://cloud.google.com/docs/authentication/provide-credentials-adc)).
	ServiceAccountCredentials *string `json:"serviceAccountCredentials,omitempty"`
	// JavaScript expression to compute the value of the severity field. Must evaluate to one of the severity values supported by Google Cloud Logging [here](https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#logseverity) (case insensitive). Defaults to "DEFAULT".
	SeverityExpression *string `json:"severityExpression,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string `json:"systemFields,omitempty"`
	// Maximum number of requests to limit to per second.
	ThrottleRateReqPerSec *int64 `json:"throttleRateReqPerSec,omitempty"`
	// Amount of time, in seconds, to wait for a request to complete before aborting it.
	TimeoutSec *int64                        `json:"timeoutSec,omitempty"`
	Type       *OutputGoogleCloudLoggingType `json:"type,omitempty"`
}
