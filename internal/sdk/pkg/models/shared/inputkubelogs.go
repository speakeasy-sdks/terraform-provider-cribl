// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type InputKubeLogsConnections struct {
	// Select a Destination.
	Output string `json:"output"`
	// Select Pipeline or Pack. Optional.
	Pipeline *string `json:"pipeline,omitempty"`
}

type InputKubeLogsMetadata struct {
	// Field name
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

// InputKubeLogsPqCompression - Codec to use to compress the persisted data.
type InputKubeLogsPqCompression string

const (
	InputKubeLogsPqCompressionNone InputKubeLogsPqCompression = "none"
	InputKubeLogsPqCompressionGzip InputKubeLogsPqCompression = "gzip"
)

func (e InputKubeLogsPqCompression) ToPointer() *InputKubeLogsPqCompression {
	return &e
}

func (e *InputKubeLogsPqCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputKubeLogsPqCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputKubeLogsPqCompression: %v", v)
	}
}

// InputKubeLogsPqMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputKubeLogsPqMode string

const (
	InputKubeLogsPqModeSmart  InputKubeLogsPqMode = "smart"
	InputKubeLogsPqModeAlways InputKubeLogsPqMode = "always"
)

func (e InputKubeLogsPqMode) ToPointer() *InputKubeLogsPqMode {
	return &e
}

func (e *InputKubeLogsPqMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputKubeLogsPqMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputKubeLogsPqMode: %v", v)
	}
}

type InputKubeLogsPq struct {
	// The number of events to send downstream before committing that Stream has read them.
	CommitFrequency *int64 `json:"commitFrequency,omitempty"`
	// Codec to use to compress the persisted data.
	Compress *InputKubeLogsPqCompression `json:"compress,omitempty"`
	// The maximum number of events to hold in memory before writing the events to disk.
	MaxBufferSize *int64 `json:"maxBufferSize,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	MaxFileSize *string `json:"maxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	MaxSize *string `json:"maxSize,omitempty"`
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputKubeLogsPqMode `json:"mode,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/inputs/<input-id>.
	Path *string `json:"path,omitempty"`
}

type InputKubeLogsRules struct {
	// Optional description of this rule's purpose
	Description *string `json:"description,omitempty"`
	// JavaScript expression applied to Pod objects. Return 'true' to include it.
	Filter string `json:"filter"`
}

type InputKubeLogsInputType string

const (
	InputKubeLogsInputTypeSplunk            InputKubeLogsInputType = "splunk"
	InputKubeLogsInputTypeSplunkHec         InputKubeLogsInputType = "splunk_hec"
	InputKubeLogsInputTypeSyslog            InputKubeLogsInputType = "syslog"
	InputKubeLogsInputTypeTcpjson           InputKubeLogsInputType = "tcpjson"
	InputKubeLogsInputTypeGrafana           InputKubeLogsInputType = "grafana"
	InputKubeLogsInputTypeLoki              InputKubeLogsInputType = "loki"
	InputKubeLogsInputTypeHTTP              InputKubeLogsInputType = "http"
	InputKubeLogsInputTypeHTTPRaw           InputKubeLogsInputType = "http_raw"
	InputKubeLogsInputTypeFirehose          InputKubeLogsInputType = "firehose"
	InputKubeLogsInputTypeElastic           InputKubeLogsInputType = "elastic"
	InputKubeLogsInputTypeKafka             InputKubeLogsInputType = "kafka"
	InputKubeLogsInputTypeConfluentCloud    InputKubeLogsInputType = "confluent_cloud"
	InputKubeLogsInputTypeMsk               InputKubeLogsInputType = "msk"
	InputKubeLogsInputTypeKinesis           InputKubeLogsInputType = "kinesis"
	InputKubeLogsInputTypeEventhub          InputKubeLogsInputType = "eventhub"
	InputKubeLogsInputTypeAzureBlob         InputKubeLogsInputType = "azure_blob"
	InputKubeLogsInputTypeMetrics           InputKubeLogsInputType = "metrics"
	InputKubeLogsInputTypeSqs               InputKubeLogsInputType = "sqs"
	InputKubeLogsInputTypeS3                InputKubeLogsInputType = "s3"
	InputKubeLogsInputTypeSnmp              InputKubeLogsInputType = "snmp"
	InputKubeLogsInputTypeCrowdstrike       InputKubeLogsInputType = "crowdstrike"
	InputKubeLogsInputTypeTCP               InputKubeLogsInputType = "tcp"
	InputKubeLogsInputTypeRawUDP            InputKubeLogsInputType = "raw_udp"
	InputKubeLogsInputTypeOffice365Service  InputKubeLogsInputType = "office365_service"
	InputKubeLogsInputTypeOffice365Mgmt     InputKubeLogsInputType = "office365_mgmt"
	InputKubeLogsInputTypeOffice365MsgTrace InputKubeLogsInputType = "office365_msg_trace"
	InputKubeLogsInputTypePrometheus        InputKubeLogsInputType = "prometheus"
	InputKubeLogsInputTypeEdgePrometheus    InputKubeLogsInputType = "edge_prometheus"
	InputKubeLogsInputTypePrometheusRw      InputKubeLogsInputType = "prometheus_rw"
	InputKubeLogsInputTypeAppscope          InputKubeLogsInputType = "appscope"
	InputKubeLogsInputTypeGooglePubsub      InputKubeLogsInputType = "google_pubsub"
	InputKubeLogsInputTypeOpenTelemetry     InputKubeLogsInputType = "open_telemetry"
	InputKubeLogsInputTypeDatadogAgent      InputKubeLogsInputType = "datadog_agent"
	InputKubeLogsInputTypeWef               InputKubeLogsInputType = "wef"
	InputKubeLogsInputTypeDatagen           InputKubeLogsInputType = "datagen"
	InputKubeLogsInputTypeCribl             InputKubeLogsInputType = "cribl"
	InputKubeLogsInputTypeCriblmetrics      InputKubeLogsInputType = "criblmetrics"
	InputKubeLogsInputTypeCriblHTTP         InputKubeLogsInputType = "cribl_http"
	InputKubeLogsInputTypeCriblTCP          InputKubeLogsInputType = "cribl_tcp"
	InputKubeLogsInputTypeWinEventLogs      InputKubeLogsInputType = "win_event_logs"
	InputKubeLogsInputTypeSystemMetrics     InputKubeLogsInputType = "system_metrics"
	InputKubeLogsInputTypeWindowsMetrics    InputKubeLogsInputType = "windows_metrics"
	InputKubeLogsInputTypeSystemState       InputKubeLogsInputType = "system_state"
	InputKubeLogsInputTypeKubeMetrics       InputKubeLogsInputType = "kube_metrics"
	InputKubeLogsInputTypeKubeLogs          InputKubeLogsInputType = "kube_logs"
	InputKubeLogsInputTypeKubeEvents        InputKubeLogsInputType = "kube_events"
	InputKubeLogsInputTypeExec              InputKubeLogsInputType = "exec"
	InputKubeLogsInputTypeSplunkSearch      InputKubeLogsInputType = "splunk_search"
	InputKubeLogsInputTypeFile              InputKubeLogsInputType = "file"
	InputKubeLogsInputTypeJournalFiles      InputKubeLogsInputType = "journal_files"
)

func (e InputKubeLogsInputType) ToPointer() *InputKubeLogsInputType {
	return &e
}

func (e *InputKubeLogsInputType) UnmarshalJSON(data []byte) error {
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
		*e = InputKubeLogsInputType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputKubeLogsInputType: %v", v)
	}
}

type InputKubeLogs struct {
	// A list of event breaking rulesets that will be applied, in order, to the input data stream.
	BreakerRulesets []string `json:"breakerRulesets,omitempty"`
	// Direct connections to Destinations, optionally via a Pipeline or a Pack.
	Connections []InputKubeLogsConnections `json:"connections,omitempty"`
	// Enable/disable this input
	Disabled *bool `json:"disabled,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Unique ID for this input
	ID string `json:"id"`
	// Time, in seconds, between checks for new containers. Default is 15 secs.
	Interval *int64 `json:"interval,omitempty"`
	// Fields to add to events from this input.
	Metadata []InputKubeLogsMetadata `json:"metadata,omitempty"`
	// Pipeline to process data from this Source before sending it through the Routes.
	Pipeline *string          `json:"pipeline,omitempty"`
	Pq       *InputKubeLogsPq `json:"pq,omitempty"`
	// For details on Persistent Queues, see: [https://docs.cribl.io/stream/persistent-queues](https://docs.cribl.io/stream/persistent-queues)
	PqEnabled *bool `json:"pqEnabled,omitempty"`
	// Add rules to decide which Pods to collect logs for. Logs are collected if no rules are given or of all the rules' expressions evaluate to true.
	Rules []InputKubeLogsRules `json:"rules,omitempty"`
	// Select whether to send data to Routes, or directly to Destinations.
	SendToRoutes *bool `json:"sendToRoutes,omitempty"`
	// The amount of time (in milliseconds) the Event Breaker will wait for new data to be sent to a specific channel, before flushing the data stream out, as-is, to the Pipelines.
	StaleChannelFlushMs *int64 `json:"staleChannelFlushMs,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// For use when containers do not emit a timestamp, prefix each line of output with a timestamp. If you enable this setting, you can use the Kubernetes Logs Event Breaker and the kubernetes_logs Pre-processing Pipeline to remove them from the events after the timestamps are extracted.
	Timestamps *bool                  `json:"timestamps,omitempty"`
	Type       InputKubeLogsInputType `json:"type"`
}