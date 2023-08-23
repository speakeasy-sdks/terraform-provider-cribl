// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// InputOffice365ServiceAuthenticationMethod - Enter client secret directly, or select a stored secret
type InputOffice365ServiceAuthenticationMethod string

const (
	InputOffice365ServiceAuthenticationMethodSecret InputOffice365ServiceAuthenticationMethod = "secret"
	InputOffice365ServiceAuthenticationMethodManual InputOffice365ServiceAuthenticationMethod = "manual"
)

func (e InputOffice365ServiceAuthenticationMethod) ToPointer() *InputOffice365ServiceAuthenticationMethod {
	return &e
}

func (e *InputOffice365ServiceAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "secret":
		fallthrough
	case "manual":
		*e = InputOffice365ServiceAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputOffice365ServiceAuthenticationMethod: %v", v)
	}
}

type InputOffice365ServiceConnections struct {
	// Select a Destination.
	Output string `json:"output"`
	// Select Pipeline or Pack. Optional.
	Pipeline *string `json:"pipeline,omitempty"`
}

// InputOffice365ServiceContentConfigLogLevel - Collector runtime Log Level
type InputOffice365ServiceContentConfigLogLevel string

const (
	InputOffice365ServiceContentConfigLogLevelError InputOffice365ServiceContentConfigLogLevel = "error"
	InputOffice365ServiceContentConfigLogLevelWarn  InputOffice365ServiceContentConfigLogLevel = "warn"
	InputOffice365ServiceContentConfigLogLevelInfo  InputOffice365ServiceContentConfigLogLevel = "info"
	InputOffice365ServiceContentConfigLogLevelDebug InputOffice365ServiceContentConfigLogLevel = "debug"
)

func (e InputOffice365ServiceContentConfigLogLevel) ToPointer() *InputOffice365ServiceContentConfigLogLevel {
	return &e
}

func (e *InputOffice365ServiceContentConfigLogLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "error":
		fallthrough
	case "warn":
		fallthrough
	case "info":
		fallthrough
	case "debug":
		*e = InputOffice365ServiceContentConfigLogLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputOffice365ServiceContentConfigLogLevel: %v", v)
	}
}

type InputOffice365ServiceContentConfig struct {
	// Office 365 Services API Content Type
	ContentType *string `json:"contentType,omitempty"`
	// If interval type is minutes the value entered must evenly divisible by 60 or save will fail
	Description *string `json:"description,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
	Interval    *int64  `json:"interval,omitempty"`
	// Collector runtime Log Level
	LogLevel *InputOffice365ServiceContentConfigLogLevel `json:"logLevel,omitempty"`
}

type InputOffice365ServiceMetadata struct {
	// Field name
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

// InputOffice365ServicePqCompression - Codec to use to compress the persisted data.
type InputOffice365ServicePqCompression string

const (
	InputOffice365ServicePqCompressionNone InputOffice365ServicePqCompression = "none"
	InputOffice365ServicePqCompressionGzip InputOffice365ServicePqCompression = "gzip"
)

func (e InputOffice365ServicePqCompression) ToPointer() *InputOffice365ServicePqCompression {
	return &e
}

func (e *InputOffice365ServicePqCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputOffice365ServicePqCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputOffice365ServicePqCompression: %v", v)
	}
}

// InputOffice365ServicePqMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputOffice365ServicePqMode string

const (
	InputOffice365ServicePqModeSmart  InputOffice365ServicePqMode = "smart"
	InputOffice365ServicePqModeAlways InputOffice365ServicePqMode = "always"
)

func (e InputOffice365ServicePqMode) ToPointer() *InputOffice365ServicePqMode {
	return &e
}

func (e *InputOffice365ServicePqMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputOffice365ServicePqMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputOffice365ServicePqMode: %v", v)
	}
}

type InputOffice365ServicePq struct {
	// The number of events to send downstream before committing that Stream has read them.
	CommitFrequency *int64 `json:"commitFrequency,omitempty"`
	// Codec to use to compress the persisted data.
	Compress *InputOffice365ServicePqCompression `json:"compress,omitempty"`
	// The maximum number of events to hold in memory before writing the events to disk.
	MaxBufferSize *int64 `json:"maxBufferSize,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	MaxFileSize *string `json:"maxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	MaxSize *string `json:"maxSize,omitempty"`
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputOffice365ServicePqMode `json:"mode,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/inputs/<input-id>.
	Path *string `json:"path,omitempty"`
}

type InputOffice365ServiceType string

const (
	InputOffice365ServiceTypeOffice365Service InputOffice365ServiceType = "office365_service"
)

func (e InputOffice365ServiceType) ToPointer() *InputOffice365ServiceType {
	return &e
}

func (e *InputOffice365ServiceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "office365_service":
		*e = InputOffice365ServiceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputOffice365ServiceType: %v", v)
	}
}

type InputOffice365Service struct {
	// Office 365 Azure Application ID
	AppID string `json:"appId"`
	// Enter client secret directly, or select a stored secret
	AuthType *InputOffice365ServiceAuthenticationMethod `json:"authType,omitempty"`
	// Office 365 Azure client secret
	ClientSecret *string `json:"clientSecret,omitempty"`
	// Direct connections to Destinations, optionally via a Pipeline or a Pack.
	Connections []InputOffice365ServiceConnections `json:"connections,omitempty"`
	// Enable Office 365 Service Communication API content types and polling intervals. Polling intervals are used to set up search date range and cron schedule, e.g.: */${interval} * * * *. Because of this, intervals entered for current and historical status must be evenly divisible by 60 to give a predictable schedule.
	ContentConfig []InputOffice365ServiceContentConfig `json:"contentConfig,omitempty"`
	// Enable/disable this input
	Disabled *bool `json:"disabled,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Unique ID for this input
	ID *string `json:"id,omitempty"`
	// Maximum time the job is allowed to run (e.g., 30, 45s or 15m). Units are seconds, if not specified. Enter 0 for unlimited time.
	JobTimeout *string `json:"jobTimeout,omitempty"`
	// How often workers should check in with the scheduler to keep job subscription alive
	KeepAliveTime *int64 `json:"keepAliveTime,omitempty"`
	// The number of Keep Alive Time periods before an inactive worker will have its job subscription revoked.
	MaxMissedKeepAlives *int64 `json:"maxMissedKeepAlives,omitempty"`
	// Fields to add to events from this input.
	Metadata []InputOffice365ServiceMetadata `json:"metadata,omitempty"`
	// Pipeline to process data from this Source before sending it through the Routes.
	Pipeline *string                  `json:"pipeline,omitempty"`
	Pq       *InputOffice365ServicePq `json:"pq,omitempty"`
	// For details on Persistent Queues, see: [https://docs.cribl.io/stream/persistent-queues](https://docs.cribl.io/stream/persistent-queues)
	PqEnabled *bool `json:"pqEnabled,omitempty"`
	// Select whether to send data to Routes, or directly to Destinations.
	SendToRoutes *bool `json:"sendToRoutes,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Office 365 Azure Tenant ID
	TenantID string `json:"tenantId"`
	// Select (or create) a stored text secret
	TextSecret *string `json:"textSecret,omitempty"`
	// HTTP request inactivity timeout, use 0 to disable
	Timeout *int64                     `json:"timeout,omitempty"`
	Type    *InputOffice365ServiceType `json:"type,omitempty"`
}