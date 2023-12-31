// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OutputInfluxdbAuthenticationType - InfluxDB authentication type
type OutputInfluxdbAuthenticationType string

const (
	OutputInfluxdbAuthenticationTypeOauth             OutputInfluxdbAuthenticationType = "oauth"
	OutputInfluxdbAuthenticationTypeBasic             OutputInfluxdbAuthenticationType = "basic"
	OutputInfluxdbAuthenticationTypeCredentialsSecret OutputInfluxdbAuthenticationType = "credentialsSecret"
	OutputInfluxdbAuthenticationTypeToken             OutputInfluxdbAuthenticationType = "token"
	OutputInfluxdbAuthenticationTypeTextSecret        OutputInfluxdbAuthenticationType = "textSecret"
)

func (e OutputInfluxdbAuthenticationType) ToPointer() *OutputInfluxdbAuthenticationType {
	return &e
}

func (e *OutputInfluxdbAuthenticationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth":
		fallthrough
	case "basic":
		fallthrough
	case "credentialsSecret":
		fallthrough
	case "token":
		fallthrough
	case "textSecret":
		*e = OutputInfluxdbAuthenticationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputInfluxdbAuthenticationType: %v", v)
	}
}

type OutputInfluxdbExtraHTTPHeaders struct {
	// Field name
	Name *string `json:"name,omitempty"`
	// Field value
	Value string `json:"value"`
}

// OutputInfluxdbFailedRequestLoggingMode - Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
type OutputInfluxdbFailedRequestLoggingMode string

const (
	OutputInfluxdbFailedRequestLoggingModePayload           OutputInfluxdbFailedRequestLoggingMode = "payload"
	OutputInfluxdbFailedRequestLoggingModePayloadAndHeaders OutputInfluxdbFailedRequestLoggingMode = "payloadAndHeaders"
	OutputInfluxdbFailedRequestLoggingModeNone              OutputInfluxdbFailedRequestLoggingMode = "none"
)

func (e OutputInfluxdbFailedRequestLoggingMode) ToPointer() *OutputInfluxdbFailedRequestLoggingMode {
	return &e
}

func (e *OutputInfluxdbFailedRequestLoggingMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputInfluxdbFailedRequestLoggingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputInfluxdbFailedRequestLoggingMode: %v", v)
	}
}

type OutputInfluxdbOauthHeaders struct {
	// OAuth header name
	Name string `json:"name"`
	// OAuth header value
	Value string `json:"value"`
}

type OutputInfluxdbOauthParams struct {
	// OAuth parameter name
	Name string `json:"name"`
	// OAuth parameter value
	Value string `json:"value"`
}

// OutputInfluxdbBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputInfluxdbBackpressureBehavior string

const (
	OutputInfluxdbBackpressureBehaviorQueue OutputInfluxdbBackpressureBehavior = "queue"
	OutputInfluxdbBackpressureBehaviorDrop  OutputInfluxdbBackpressureBehavior = "drop"
	OutputInfluxdbBackpressureBehaviorBlock OutputInfluxdbBackpressureBehavior = "block"
)

func (e OutputInfluxdbBackpressureBehavior) ToPointer() *OutputInfluxdbBackpressureBehavior {
	return &e
}

func (e *OutputInfluxdbBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputInfluxdbBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputInfluxdbBackpressureBehavior: %v", v)
	}
}

type OutputInfluxdbOptionalFieldsInGeneralSection string

const (
	OutputInfluxdbOptionalFieldsInGeneralSectionUseV2API OutputInfluxdbOptionalFieldsInGeneralSection = "useV2API"
)

func (e OutputInfluxdbOptionalFieldsInGeneralSection) ToPointer() *OutputInfluxdbOptionalFieldsInGeneralSection {
	return &e
}

func (e *OutputInfluxdbOptionalFieldsInGeneralSection) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "useV2API":
		*e = OutputInfluxdbOptionalFieldsInGeneralSection(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputInfluxdbOptionalFieldsInGeneralSection: %v", v)
	}
}

// OutputInfluxdbCompression - Codec to use to compress the persisted data.
type OutputInfluxdbCompression string

const (
	OutputInfluxdbCompressionNone OutputInfluxdbCompression = "none"
	OutputInfluxdbCompressionGzip OutputInfluxdbCompression = "gzip"
)

func (e OutputInfluxdbCompression) ToPointer() *OutputInfluxdbCompression {
	return &e
}

func (e *OutputInfluxdbCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputInfluxdbCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputInfluxdbCompression: %v", v)
	}
}

type OutputInfluxdbPqControls struct {
}

// OutputInfluxdbQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputInfluxdbQueueFullBehavior string

const (
	OutputInfluxdbQueueFullBehaviorBlock OutputInfluxdbQueueFullBehavior = "block"
	OutputInfluxdbQueueFullBehaviorDrop  OutputInfluxdbQueueFullBehavior = "drop"
)

func (e OutputInfluxdbQueueFullBehavior) ToPointer() *OutputInfluxdbQueueFullBehavior {
	return &e
}

func (e *OutputInfluxdbQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputInfluxdbQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputInfluxdbQueueFullBehavior: %v", v)
	}
}

// OutputInfluxdbTimestampPrecision - Sets the precision for the supplied Unix time values. Defaults to milliseconds.
type OutputInfluxdbTimestampPrecision string

const (
	OutputInfluxdbTimestampPrecisionNs OutputInfluxdbTimestampPrecision = "ns"
	OutputInfluxdbTimestampPrecisionU  OutputInfluxdbTimestampPrecision = "u"
	OutputInfluxdbTimestampPrecisionMs OutputInfluxdbTimestampPrecision = "ms"
	OutputInfluxdbTimestampPrecisionS  OutputInfluxdbTimestampPrecision = "s"
	OutputInfluxdbTimestampPrecisionM  OutputInfluxdbTimestampPrecision = "m"
	OutputInfluxdbTimestampPrecisionH  OutputInfluxdbTimestampPrecision = "h"
)

func (e OutputInfluxdbTimestampPrecision) ToPointer() *OutputInfluxdbTimestampPrecision {
	return &e
}

func (e *OutputInfluxdbTimestampPrecision) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ns":
		fallthrough
	case "u":
		fallthrough
	case "ms":
		fallthrough
	case "s":
		fallthrough
	case "m":
		fallthrough
	case "h":
		*e = OutputInfluxdbTimestampPrecision(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputInfluxdbTimestampPrecision: %v", v)
	}
}

type OutputInfluxdbType string

const (
	OutputInfluxdbTypeInfluxdb OutputInfluxdbType = "influxdb"
)

func (e OutputInfluxdbType) ToPointer() *OutputInfluxdbType {
	return &e
}

func (e *OutputInfluxdbType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "influxdb":
		*e = OutputInfluxdbType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputInfluxdbType: %v", v)
	}
}

type OutputInfluxdb struct {
	// JavaScript expression to compute the Authorization header value to pass in requests. The value `${token}` is used to reference the token obtained from authentication, e.g.: `Bearer ${token}`.
	AuthHeaderExpr *string `json:"authHeaderExpr,omitempty"`
	// InfluxDB authentication type
	AuthType *OutputInfluxdbAuthenticationType `json:"authType,omitempty"`
	// Bucket to write to.
	Bucket *string `json:"bucket,omitempty"`
	// Whether to compress the payload body before sending.
	Compress *bool `json:"compress,omitempty"`
	// Maximum number of ongoing requests before blocking.
	Concurrency *int64 `json:"concurrency,omitempty"`
	// Select (or create) a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
	// Database to write to.
	Database *string `json:"database,omitempty"`
	// Enabling this will pull the value field from the metric name. E,g, 'db.query.user' will use 'db.query' as the measurement and 'user' as the value field.
	DynamicValueFieldName *bool `json:"dynamicValueFieldName,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Headers to add to all events.
	ExtraHTTPHeaders []OutputInfluxdbExtraHTTPHeaders `json:"extraHttpHeaders,omitempty"`
	// Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
	FailedRequestLoggingMode *OutputInfluxdbFailedRequestLoggingMode `json:"failedRequestLoggingMode,omitempty"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Max body size.
	FlushPeriodSec *int64 `json:"flushPeriodSec,omitempty"`
	// Unique ID for this output
	ID *string `json:"id,omitempty"`
	// URL for OAuth
	LoginURL *string `json:"loginUrl,omitempty"`
	// Max number of events to include in the request body. Default is 0 (unlimited).
	MaxPayloadEvents *int64 `json:"maxPayloadEvents,omitempty"`
	// Maximum size, in KB, of the request body.
	MaxPayloadSizeKB *int64 `json:"maxPayloadSizeKB,omitempty"`
	// Additional headers to send in the OAuth login request. @{product} will automatically add the content-type header 'application/x-www-form-urlencoded' when sending this request.
	OauthHeaders []OutputInfluxdbOauthHeaders `json:"oauthHeaders,omitempty"`
	// Additional parameters to send in the OAuth login request. @{product} will combine the secret with these parameters, and will send the URL-encoded result in a POST request to the endpoint specified in the 'Login URL'. We'll automatically add the content-type header 'application/x-www-form-urlencoded' when sending this request.
	OauthParams []OutputInfluxdbOauthParams `json:"oauthParams,omitempty"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure                 *OutputInfluxdbBackpressureBehavior           `json:"onBackpressure,omitempty"`
	OptionalFieldsInGeneralSection *OutputInfluxdbOptionalFieldsInGeneralSection `json:"optionalFieldsInGeneralSection,omitempty"`
	// Organization ID for this bucket.
	Org *string `json:"org,omitempty"`
	// Password for Basic authentication
	Password *string `json:"password,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputInfluxdbCompression `json:"pqCompress,omitempty"`
	PqControls *OutputInfluxdbPqControls  `json:"pqControls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `json:"pqMaxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `json:"pqMaxSize,omitempty"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputInfluxdbQueueFullBehavior `json:"pqOnBackpressure,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `json:"pqPath,omitempty"`
	// Toggle this off to forward new events to receiver(s) before queue is flushed. Otherwise, default drain behavior is FIFO (first in, first out).
	PqStrictOrdering *bool `json:"pqStrictOrdering,omitempty"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another trusted CA (e.g., the system's CA). Defaults to No.
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// List of headers that are safe to log in plain text.
	SafeHeaders []string `json:"safeHeaders,omitempty"`
	// Secret parameter value to pass in request body
	Secret *string `json:"secret,omitempty"`
	// Secret parameter name to pass in request body
	SecretParamName *string `json:"secretParamName,omitempty"`
	Spacer          *string `json:"spacer,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string `json:"systemFields,omitempty"`
	// Select (or create) a stored text secret
	TextSecret *string `json:"textSecret,omitempty"`
	// Amount of time, in seconds, to wait for a request to complete before aborting it.
	TimeoutSec *int64 `json:"timeoutSec,omitempty"`
	// Sets the precision for the supplied Unix time values. Defaults to milliseconds.
	TimestampPrecision *OutputInfluxdbTimestampPrecision `json:"timestampPrecision,omitempty"`
	// Bearer token to include in the authorization header
	Token *string `json:"token,omitempty"`
	// Name of the auth token attribute in the OAuth response. Can be top-level (e.g., 'token'); or nested, using a period (e.g., 'data.token').
	TokenAttributeName *string `json:"tokenAttributeName,omitempty"`
	// How often the OAuth token should be refreshed.
	TokenTimeoutSecs *int64             `json:"tokenTimeoutSecs,omitempty"`
	Type             OutputInfluxdbType `json:"type"`
	// URL of an InfluxDB cluster to send events to, e.g., http://localhost:8086/write
	URL string `json:"url"`
	// Enable to use round-robin DNS lookup. When a DNS server returns multiple addresses, this will cause Stream to cycle through them in the order returned.
	UseRoundRobinDNS *bool `json:"useRoundRobinDns,omitempty"`
	// The v2 API can be enabled with InfluxDB versions 1.8 and later.
	UseV2API *bool `json:"useV2API,omitempty"`
	// Username for Basic authentication
	Username *string `json:"username,omitempty"`
	// Name of the field in which to store the metric when sending to InfluxDB. If dynamic generation is enabled and fails, this will be used as a fallback.
	ValueFieldName *string `json:"valueFieldName,omitempty"`
}
