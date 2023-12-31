// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// InputSplunkSearchAuthenticationType - Splunk Search authentication type
type InputSplunkSearchAuthenticationType string

const (
	InputSplunkSearchAuthenticationTypeOauth             InputSplunkSearchAuthenticationType = "oauth"
	InputSplunkSearchAuthenticationTypeBasic             InputSplunkSearchAuthenticationType = "basic"
	InputSplunkSearchAuthenticationTypeCredentialsSecret InputSplunkSearchAuthenticationType = "credentialsSecret"
	InputSplunkSearchAuthenticationTypeToken             InputSplunkSearchAuthenticationType = "token"
	InputSplunkSearchAuthenticationTypeTextSecret        InputSplunkSearchAuthenticationType = "textSecret"
)

func (e InputSplunkSearchAuthenticationType) ToPointer() *InputSplunkSearchAuthenticationType {
	return &e
}

func (e *InputSplunkSearchAuthenticationType) UnmarshalJSON(data []byte) error {
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
		*e = InputSplunkSearchAuthenticationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSplunkSearchAuthenticationType: %v", v)
	}
}

type InputSplunkSearchConnections struct {
	// Select a Destination.
	Output string `json:"output"`
	// Select Pipeline or Pack. Optional.
	Pipeline *string `json:"pipeline,omitempty"`
}

type InputSplunkSearchEndpointHeaders struct {
	// Header Name
	Name string `json:"name"`
	// JavaScript expression to compute the header's value, normally enclosed in backticks (e.g., `${earliest}`). If a constant, use single quotes (e.g., 'earliest'). Values without delimiters (e.g., earliest) are evaluated as strings.
	Value string `json:"value"`
}

type InputSplunkSearchEndpointParams struct {
	// Parameter name
	Name string `json:"name"`
	// JavaScript expression to compute the parameter's value, normally enclosed in backticks (e.g., `${earliest}`). If a constant, use single quotes (e.g., 'earliest'). Values without delimiters (e.g., earliest) are evaluated as strings.
	Value string `json:"value"`
}

// InputSplunkSearchLogLevel - Collector runtime log Level (verbosity).
type InputSplunkSearchLogLevel string

const (
	InputSplunkSearchLogLevelError InputSplunkSearchLogLevel = "error"
	InputSplunkSearchLogLevelWarn  InputSplunkSearchLogLevel = "warn"
	InputSplunkSearchLogLevelInfo  InputSplunkSearchLogLevel = "info"
	InputSplunkSearchLogLevelDebug InputSplunkSearchLogLevel = "debug"
)

func (e InputSplunkSearchLogLevel) ToPointer() *InputSplunkSearchLogLevel {
	return &e
}

func (e *InputSplunkSearchLogLevel) UnmarshalJSON(data []byte) error {
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
		*e = InputSplunkSearchLogLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSplunkSearchLogLevel: %v", v)
	}
}

type InputSplunkSearchMetadata struct {
	// Field name
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

type InputSplunkSearchOauthHeaders struct {
	// OAuth header name
	Name string `json:"name"`
	// OAuth header value
	Value string `json:"value"`
}

type InputSplunkSearchOauthParams struct {
	// OAuth parameter name
	Name string `json:"name"`
	// OAuth parameter value
	Value string `json:"value"`
}

// InputSplunkSearchOutputMode - Format of the returned output
type InputSplunkSearchOutputMode string

const (
	InputSplunkSearchOutputModeCsv  InputSplunkSearchOutputMode = "csv"
	InputSplunkSearchOutputModeJSON InputSplunkSearchOutputMode = "json"
)

func (e InputSplunkSearchOutputMode) ToPointer() *InputSplunkSearchOutputMode {
	return &e
}

func (e *InputSplunkSearchOutputMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "csv":
		fallthrough
	case "json":
		*e = InputSplunkSearchOutputMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSplunkSearchOutputMode: %v", v)
	}
}

// InputSplunkSearchPqCompression - Codec to use to compress the persisted data.
type InputSplunkSearchPqCompression string

const (
	InputSplunkSearchPqCompressionNone InputSplunkSearchPqCompression = "none"
	InputSplunkSearchPqCompressionGzip InputSplunkSearchPqCompression = "gzip"
)

func (e InputSplunkSearchPqCompression) ToPointer() *InputSplunkSearchPqCompression {
	return &e
}

func (e *InputSplunkSearchPqCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputSplunkSearchPqCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSplunkSearchPqCompression: %v", v)
	}
}

// InputSplunkSearchPqMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputSplunkSearchPqMode string

const (
	InputSplunkSearchPqModeSmart  InputSplunkSearchPqMode = "smart"
	InputSplunkSearchPqModeAlways InputSplunkSearchPqMode = "always"
)

func (e InputSplunkSearchPqMode) ToPointer() *InputSplunkSearchPqMode {
	return &e
}

func (e *InputSplunkSearchPqMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputSplunkSearchPqMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSplunkSearchPqMode: %v", v)
	}
}

type InputSplunkSearchPq struct {
	// The number of events to send downstream before committing that Stream has read them.
	CommitFrequency *int64 `json:"commitFrequency,omitempty"`
	// Codec to use to compress the persisted data.
	Compress *InputSplunkSearchPqCompression `json:"compress,omitempty"`
	// The maximum number of events to hold in memory before writing the events to disk.
	MaxBufferSize *int64 `json:"maxBufferSize,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	MaxFileSize *string `json:"maxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	MaxSize *string `json:"maxSize,omitempty"`
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputSplunkSearchPqMode `json:"mode,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/inputs/<input-id>.
	Path *string `json:"path,omitempty"`
}

type InputSplunkSearchInputType string

const (
	InputSplunkSearchInputTypeSplunk            InputSplunkSearchInputType = "splunk"
	InputSplunkSearchInputTypeSplunkHec         InputSplunkSearchInputType = "splunk_hec"
	InputSplunkSearchInputTypeSyslog            InputSplunkSearchInputType = "syslog"
	InputSplunkSearchInputTypeTcpjson           InputSplunkSearchInputType = "tcpjson"
	InputSplunkSearchInputTypeGrafana           InputSplunkSearchInputType = "grafana"
	InputSplunkSearchInputTypeLoki              InputSplunkSearchInputType = "loki"
	InputSplunkSearchInputTypeHTTP              InputSplunkSearchInputType = "http"
	InputSplunkSearchInputTypeHTTPRaw           InputSplunkSearchInputType = "http_raw"
	InputSplunkSearchInputTypeFirehose          InputSplunkSearchInputType = "firehose"
	InputSplunkSearchInputTypeElastic           InputSplunkSearchInputType = "elastic"
	InputSplunkSearchInputTypeKafka             InputSplunkSearchInputType = "kafka"
	InputSplunkSearchInputTypeConfluentCloud    InputSplunkSearchInputType = "confluent_cloud"
	InputSplunkSearchInputTypeMsk               InputSplunkSearchInputType = "msk"
	InputSplunkSearchInputTypeKinesis           InputSplunkSearchInputType = "kinesis"
	InputSplunkSearchInputTypeEventhub          InputSplunkSearchInputType = "eventhub"
	InputSplunkSearchInputTypeAzureBlob         InputSplunkSearchInputType = "azure_blob"
	InputSplunkSearchInputTypeMetrics           InputSplunkSearchInputType = "metrics"
	InputSplunkSearchInputTypeSqs               InputSplunkSearchInputType = "sqs"
	InputSplunkSearchInputTypeS3                InputSplunkSearchInputType = "s3"
	InputSplunkSearchInputTypeSnmp              InputSplunkSearchInputType = "snmp"
	InputSplunkSearchInputTypeCrowdstrike       InputSplunkSearchInputType = "crowdstrike"
	InputSplunkSearchInputTypeTCP               InputSplunkSearchInputType = "tcp"
	InputSplunkSearchInputTypeRawUDP            InputSplunkSearchInputType = "raw_udp"
	InputSplunkSearchInputTypeOffice365Service  InputSplunkSearchInputType = "office365_service"
	InputSplunkSearchInputTypeOffice365Mgmt     InputSplunkSearchInputType = "office365_mgmt"
	InputSplunkSearchInputTypeOffice365MsgTrace InputSplunkSearchInputType = "office365_msg_trace"
	InputSplunkSearchInputTypePrometheus        InputSplunkSearchInputType = "prometheus"
	InputSplunkSearchInputTypeEdgePrometheus    InputSplunkSearchInputType = "edge_prometheus"
	InputSplunkSearchInputTypePrometheusRw      InputSplunkSearchInputType = "prometheus_rw"
	InputSplunkSearchInputTypeAppscope          InputSplunkSearchInputType = "appscope"
	InputSplunkSearchInputTypeGooglePubsub      InputSplunkSearchInputType = "google_pubsub"
	InputSplunkSearchInputTypeOpenTelemetry     InputSplunkSearchInputType = "open_telemetry"
	InputSplunkSearchInputTypeDatadogAgent      InputSplunkSearchInputType = "datadog_agent"
	InputSplunkSearchInputTypeWef               InputSplunkSearchInputType = "wef"
	InputSplunkSearchInputTypeDatagen           InputSplunkSearchInputType = "datagen"
	InputSplunkSearchInputTypeCribl             InputSplunkSearchInputType = "cribl"
	InputSplunkSearchInputTypeCriblmetrics      InputSplunkSearchInputType = "criblmetrics"
	InputSplunkSearchInputTypeCriblHTTP         InputSplunkSearchInputType = "cribl_http"
	InputSplunkSearchInputTypeCriblTCP          InputSplunkSearchInputType = "cribl_tcp"
	InputSplunkSearchInputTypeWinEventLogs      InputSplunkSearchInputType = "win_event_logs"
	InputSplunkSearchInputTypeSystemMetrics     InputSplunkSearchInputType = "system_metrics"
	InputSplunkSearchInputTypeWindowsMetrics    InputSplunkSearchInputType = "windows_metrics"
	InputSplunkSearchInputTypeSystemState       InputSplunkSearchInputType = "system_state"
	InputSplunkSearchInputTypeKubeMetrics       InputSplunkSearchInputType = "kube_metrics"
	InputSplunkSearchInputTypeKubeLogs          InputSplunkSearchInputType = "kube_logs"
	InputSplunkSearchInputTypeKubeEvents        InputSplunkSearchInputType = "kube_events"
	InputSplunkSearchInputTypeExec              InputSplunkSearchInputType = "exec"
	InputSplunkSearchInputTypeSplunkSearch      InputSplunkSearchInputType = "splunk_search"
	InputSplunkSearchInputTypeFile              InputSplunkSearchInputType = "file"
	InputSplunkSearchInputTypeJournalFiles      InputSplunkSearchInputType = "journal_files"
)

func (e InputSplunkSearchInputType) ToPointer() *InputSplunkSearchInputType {
	return &e
}

func (e *InputSplunkSearchInputType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "splunk":
		fallthrough
	case "splunk_hec":
		fallthrough
	case "syslog":
		fallthrough
	case "tcpjson":
		fallthrough
	case "grafana":
		fallthrough
	case "loki":
		fallthrough
	case "http":
		fallthrough
	case "http_raw":
		fallthrough
	case "firehose":
		fallthrough
	case "elastic":
		fallthrough
	case "kafka":
		fallthrough
	case "confluent_cloud":
		fallthrough
	case "msk":
		fallthrough
	case "kinesis":
		fallthrough
	case "eventhub":
		fallthrough
	case "azure_blob":
		fallthrough
	case "metrics":
		fallthrough
	case "sqs":
		fallthrough
	case "s3":
		fallthrough
	case "snmp":
		fallthrough
	case "crowdstrike":
		fallthrough
	case "tcp":
		fallthrough
	case "raw_udp":
		fallthrough
	case "office365_service":
		fallthrough
	case "office365_mgmt":
		fallthrough
	case "office365_msg_trace":
		fallthrough
	case "prometheus":
		fallthrough
	case "edge_prometheus":
		fallthrough
	case "prometheus_rw":
		fallthrough
	case "appscope":
		fallthrough
	case "google_pubsub":
		fallthrough
	case "open_telemetry":
		fallthrough
	case "datadog_agent":
		fallthrough
	case "wef":
		fallthrough
	case "datagen":
		fallthrough
	case "cribl":
		fallthrough
	case "criblmetrics":
		fallthrough
	case "cribl_http":
		fallthrough
	case "cribl_tcp":
		fallthrough
	case "win_event_logs":
		fallthrough
	case "system_metrics":
		fallthrough
	case "windows_metrics":
		fallthrough
	case "system_state":
		fallthrough
	case "kube_metrics":
		fallthrough
	case "kube_logs":
		fallthrough
	case "kube_events":
		fallthrough
	case "exec":
		fallthrough
	case "splunk_search":
		fallthrough
	case "file":
		fallthrough
	case "journal_files":
		*e = InputSplunkSearchInputType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSplunkSearchInputType: %v", v)
	}
}

type InputSplunkSearch struct {
	// JavaScript expression to compute the Authorization header value to pass in requests. The value `${token}` is used to reference the token obtained from authentication, e.g.: `Bearer ${token}`.
	AuthHeaderExpr *string `json:"authHeaderExpr,omitempty"`
	// Splunk Search authentication type
	AuthType *InputSplunkSearchAuthenticationType `json:"authType,omitempty"`
	// A list of event breaking rulesets that will be applied, in order, to the input data stream.
	BreakerRulesets []string `json:"breakerRulesets,omitempty"`
	// Direct connections to Destinations, optionally via a Pipeline or a Pack.
	Connections []InputSplunkSearchConnections `json:"connections,omitempty"`
	// Select (or create) a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
	// A cron schedule on which to run this job.
	CronSchedule string `json:"cronSchedule"`
	// Enable/disable this input
	Disabled *bool `json:"disabled,omitempty"`
	// The earliest time boundary for the search. Can be an exact or relative time. For example: '2022-01-14T12:00:00Z' or '-16m@m'
	Earliest *string `json:"earliest,omitempty"`
	// REST API used to create a search.
	Endpoint string `json:"endpoint"`
	// Optional request headers to send to the endpoint.
	EndpointHeaders []InputSplunkSearchEndpointHeaders `json:"endpointHeaders,omitempty"`
	// Optional request parameters to send to the endpoint.
	EndpointParams []InputSplunkSearchEndpointParams `json:"endpointParams,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Unique ID for this input
	ID *string `json:"id,omitempty"`
	// Maximum time the job is allowed to run (e.g., 30, 45s or 15m). Units are seconds, if not specified. Enter 0 for unlimited time.
	JobTimeout *string `json:"jobTimeout,omitempty"`
	// How often workers should check in with the scheduler to keep job subscription alive
	KeepAliveTime *int64 `json:"keepAliveTime,omitempty"`
	// The latest time boundary for the search. Can be an exact or relative time. For example: '2022-01-14T12:00:00Z' or '-1m@m'
	Latest *string `json:"latest,omitempty"`
	// Collector runtime log Level (verbosity).
	LogLevel *InputSplunkSearchLogLevel `json:"logLevel,omitempty"`
	// URL for OAuth
	LoginURL *string `json:"loginUrl,omitempty"`
	// The number of Keep Alive Time periods before an inactive worker will have its job subscription revoked.
	MaxMissedKeepAlives *int64 `json:"maxMissedKeepAlives,omitempty"`
	// Fields to add to events from this input.
	Metadata []InputSplunkSearchMetadata `json:"metadata,omitempty"`
	// Additional headers to send in the OAuth login request. @{product} will automatically add the content-type header 'application/x-www-form-urlencoded' when sending this request.
	OauthHeaders []InputSplunkSearchOauthHeaders `json:"oauthHeaders,omitempty"`
	// Additional parameters to send in the OAuth login request. @{product} will combine the secret with these parameters, and will send the URL-encoded result in a POST request to the endpoint specified in the 'Login URL'. We'll automatically add the content-type header 'application/x-www-form-urlencoded' when sending this request.
	OauthParams []InputSplunkSearchOauthParams `json:"oauthParams,omitempty"`
	// Format of the returned output
	OutputMode InputSplunkSearchOutputMode `json:"outputMode"`
	// Password for Basic authentication
	Password *string `json:"password,omitempty"`
	// Pipeline to process data from this Source before sending it through the Routes.
	Pipeline *string              `json:"pipeline,omitempty"`
	Pq       *InputSplunkSearchPq `json:"pq,omitempty"`
	// For details on Persistent Queues, see: [https://docs.cribl.io/stream/persistent-queues](https://docs.cribl.io/stream/persistent-queues)
	PqEnabled *bool `json:"pqEnabled,omitempty"`
	// HTTP request inactivity timeout, use 0 to disable
	RequestTimeout *int64 `json:"requestTimeout,omitempty"`
	// Enter Splunk search here. For example: 'index=myAppLogs level=error channel=myApp' OR '| mstats avg(myStat) as myStat WHERE index=myStatsIndex.'
	Search string `json:"search"`
	// Search head base URL, can be expression, default is https://localhost:8089.
	SearchHead string `json:"searchHead"`
	// Secret parameter value to pass in request body
	Secret *string `json:"secret,omitempty"`
	// Secret parameter name to pass in request body
	SecretParamName *string `json:"secretParamName,omitempty"`
	// Select whether to send data to Routes, or directly to Destinations.
	SendToRoutes *bool   `json:"sendToRoutes,omitempty"`
	Spacer       *string `json:"spacer,omitempty"`
	// The amount of time (in milliseconds) the Event Breaker will wait for new data to be sent to a specific channel, before flushing the data stream out, as-is, to the Pipelines.
	StaleChannelFlushMs *int64 `json:"staleChannelFlushMs,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Select (or create) a stored text secret
	TextSecret *string `json:"textSecret,omitempty"`
	// Bearer token to include in the authorization header
	Token *string `json:"token,omitempty"`
	// Name of the auth token attribute in the OAuth response. Can be top-level (e.g., 'token'); or nested, using a period (e.g., 'data.token').
	TokenAttributeName *string `json:"tokenAttributeName,omitempty"`
	// How often the OAuth token should be refreshed.
	TokenTimeoutSecs *int64                      `json:"tokenTimeoutSecs,omitempty"`
	Type             *InputSplunkSearchInputType `json:"type,omitempty"`
	// Enable to use round-robin DNS lookup. When a DNS server returns multiple addresses, this will cause Stream to cycle through them in the order returned.
	UseRoundRobinDNS *bool `json:"useRoundRobinDns,omitempty"`
	// Username for Basic authentication
	Username *string `json:"username,omitempty"`
}
