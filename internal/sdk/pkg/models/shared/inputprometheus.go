// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// InputPrometheusAuthenticationMethod - Enter credentials directly, or select a stored secret
type InputPrometheusAuthenticationMethod string

const (
	InputPrometheusAuthenticationMethodSecret InputPrometheusAuthenticationMethod = "secret"
	InputPrometheusAuthenticationMethodManual InputPrometheusAuthenticationMethod = "manual"
)

func (e InputPrometheusAuthenticationMethod) ToPointer() *InputPrometheusAuthenticationMethod {
	return &e
}

func (e *InputPrometheusAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "secret":
		fallthrough
	case "manual":
		*e = InputPrometheusAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusAuthenticationMethod: %v", v)
	}
}

// InputPrometheusAuthenticationMethod1 - AWS authentication method. Choose Auto to use IAM roles.
type InputPrometheusAuthenticationMethod1 string

const (
	InputPrometheusAuthenticationMethod1Auto   InputPrometheusAuthenticationMethod1 = "auto"
	InputPrometheusAuthenticationMethod1Manual InputPrometheusAuthenticationMethod1 = "manual"
	InputPrometheusAuthenticationMethod1Secret InputPrometheusAuthenticationMethod1 = "secret"
)

func (e InputPrometheusAuthenticationMethod1) ToPointer() *InputPrometheusAuthenticationMethod1 {
	return &e
}

func (e *InputPrometheusAuthenticationMethod1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auto":
		fallthrough
	case "manual":
		fallthrough
	case "secret":
		*e = InputPrometheusAuthenticationMethod1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusAuthenticationMethod1: %v", v)
	}
}

type InputPrometheusConnections struct {
	// Select a Destination.
	Output string `json:"output"`
	// Select Pipeline or Pack. Optional.
	Pipeline *string `json:"pipeline,omitempty"`
}

// InputPrometheusDiscoveryType - Target discovery mechanism. Use static to manually enter a list of targets.
type InputPrometheusDiscoveryType string

const (
	InputPrometheusDiscoveryTypeEc2 InputPrometheusDiscoveryType = "ec2"
	InputPrometheusDiscoveryTypeDNS InputPrometheusDiscoveryType = "dns"
)

func (e InputPrometheusDiscoveryType) ToPointer() *InputPrometheusDiscoveryType {
	return &e
}

func (e *InputPrometheusDiscoveryType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ec2":
		fallthrough
	case "dns":
		*e = InputPrometheusDiscoveryType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusDiscoveryType: %v", v)
	}
}

// InputPrometheusLogLevel - Collector runtime Log Level
type InputPrometheusLogLevel string

const (
	InputPrometheusLogLevelError InputPrometheusLogLevel = "error"
	InputPrometheusLogLevelWarn  InputPrometheusLogLevel = "warn"
	InputPrometheusLogLevelInfo  InputPrometheusLogLevel = "info"
	InputPrometheusLogLevelDebug InputPrometheusLogLevel = "debug"
)

func (e InputPrometheusLogLevel) ToPointer() *InputPrometheusLogLevel {
	return &e
}

func (e *InputPrometheusLogLevel) UnmarshalJSON(data []byte) error {
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
		*e = InputPrometheusLogLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusLogLevel: %v", v)
	}
}

type InputPrometheusMetadata struct {
	// Field name
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

// InputPrometheusDNSNames - List of DNS names to resolve
type InputPrometheusDNSNames struct {
}

type InputPrometheusOptionalFieldsInGeneralSection string

const (
	InputPrometheusOptionalFieldsInGeneralSectionDiscoveryType InputPrometheusOptionalFieldsInGeneralSection = "discoveryType"
)

func (e InputPrometheusOptionalFieldsInGeneralSection) ToPointer() *InputPrometheusOptionalFieldsInGeneralSection {
	return &e
}

func (e *InputPrometheusOptionalFieldsInGeneralSection) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "discoveryType":
		*e = InputPrometheusOptionalFieldsInGeneralSection(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusOptionalFieldsInGeneralSection: %v", v)
	}
}

// InputPrometheusPqCompression - Codec to use to compress the persisted data.
type InputPrometheusPqCompression string

const (
	InputPrometheusPqCompressionNone InputPrometheusPqCompression = "none"
	InputPrometheusPqCompressionGzip InputPrometheusPqCompression = "gzip"
)

func (e InputPrometheusPqCompression) ToPointer() *InputPrometheusPqCompression {
	return &e
}

func (e *InputPrometheusPqCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputPrometheusPqCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusPqCompression: %v", v)
	}
}

// InputPrometheusPqMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputPrometheusPqMode string

const (
	InputPrometheusPqModeSmart  InputPrometheusPqMode = "smart"
	InputPrometheusPqModeAlways InputPrometheusPqMode = "always"
)

func (e InputPrometheusPqMode) ToPointer() *InputPrometheusPqMode {
	return &e
}

func (e *InputPrometheusPqMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputPrometheusPqMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusPqMode: %v", v)
	}
}

type InputPrometheusPq struct {
	// The number of events to send downstream before committing that Stream has read them.
	CommitFrequency *int64 `json:"commitFrequency,omitempty"`
	// Codec to use to compress the persisted data.
	Compress *InputPrometheusPqCompression `json:"compress,omitempty"`
	// The maximum number of events to hold in memory before writing the events to disk.
	MaxBufferSize *int64 `json:"maxBufferSize,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	MaxFileSize *string `json:"maxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	MaxSize *string `json:"maxSize,omitempty"`
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputPrometheusPqMode `json:"mode,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/inputs/<input-id>.
	Path *string `json:"path,omitempty"`
}

// InputPrometheusRecordType - DNS Record type to resolve
type InputPrometheusRecordType string

const (
	InputPrometheusRecordTypeSrv  InputPrometheusRecordType = "SRV"
	InputPrometheusRecordTypeA    InputPrometheusRecordType = "A"
	InputPrometheusRecordTypeAaaa InputPrometheusRecordType = "AAAA"
)

func (e InputPrometheusRecordType) ToPointer() *InputPrometheusRecordType {
	return &e
}

func (e *InputPrometheusRecordType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SRV":
		fallthrough
	case "A":
		fallthrough
	case "AAAA":
		*e = InputPrometheusRecordType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusRecordType: %v", v)
	}
}

// InputPrometheusRegion - Region where the EC2 is located
type InputPrometheusRegion string

const (
	InputPrometheusRegionUsEast1      InputPrometheusRegion = "us-east-1"
	InputPrometheusRegionUsEast2      InputPrometheusRegion = "us-east-2"
	InputPrometheusRegionUsWest1      InputPrometheusRegion = "us-west-1"
	InputPrometheusRegionUsWest2      InputPrometheusRegion = "us-west-2"
	InputPrometheusRegionAfSouth1     InputPrometheusRegion = "af-south-1"
	InputPrometheusRegionCaCentral1   InputPrometheusRegion = "ca-central-1"
	InputPrometheusRegionEuWest1      InputPrometheusRegion = "eu-west-1"
	InputPrometheusRegionEuCentral1   InputPrometheusRegion = "eu-central-1"
	InputPrometheusRegionEuWest2      InputPrometheusRegion = "eu-west-2"
	InputPrometheusRegionEuSouth1     InputPrometheusRegion = "eu-south-1"
	InputPrometheusRegionEuWest3      InputPrometheusRegion = "eu-west-3"
	InputPrometheusRegionEuNorth1     InputPrometheusRegion = "eu-north-1"
	InputPrometheusRegionApEast1      InputPrometheusRegion = "ap-east-1"
	InputPrometheusRegionApNortheast1 InputPrometheusRegion = "ap-northeast-1"
	InputPrometheusRegionApNortheast2 InputPrometheusRegion = "ap-northeast-2"
	InputPrometheusRegionApSoutheast1 InputPrometheusRegion = "ap-southeast-1"
	InputPrometheusRegionApSoutheast2 InputPrometheusRegion = "ap-southeast-2"
	InputPrometheusRegionApSouth1     InputPrometheusRegion = "ap-south-1"
	InputPrometheusRegionMeSouth1     InputPrometheusRegion = "me-south-1"
	InputPrometheusRegionSaEast1      InputPrometheusRegion = "sa-east-1"
	InputPrometheusRegionUsGovEast1   InputPrometheusRegion = "us-gov-east-1"
	InputPrometheusRegionUsGovWest1   InputPrometheusRegion = "us-gov-west-1"
)

func (e InputPrometheusRegion) ToPointer() *InputPrometheusRegion {
	return &e
}

func (e *InputPrometheusRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "us-east-1":
		fallthrough
	case "us-east-2":
		fallthrough
	case "us-west-1":
		fallthrough
	case "us-west-2":
		fallthrough
	case "af-south-1":
		fallthrough
	case "ca-central-1":
		fallthrough
	case "eu-west-1":
		fallthrough
	case "eu-central-1":
		fallthrough
	case "eu-west-2":
		fallthrough
	case "eu-south-1":
		fallthrough
	case "eu-west-3":
		fallthrough
	case "eu-north-1":
		fallthrough
	case "ap-east-1":
		fallthrough
	case "ap-northeast-1":
		fallthrough
	case "ap-northeast-2":
		fallthrough
	case "ap-southeast-1":
		fallthrough
	case "ap-southeast-2":
		fallthrough
	case "ap-south-1":
		fallthrough
	case "me-south-1":
		fallthrough
	case "sa-east-1":
		fallthrough
	case "us-gov-east-1":
		fallthrough
	case "us-gov-west-1":
		*e = InputPrometheusRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusRegion: %v", v)
	}
}

// InputPrometheusMetricsProtocol - Protocol to use when collecting metrics
type InputPrometheusMetricsProtocol string

const (
	InputPrometheusMetricsProtocolHTTP  InputPrometheusMetricsProtocol = "http"
	InputPrometheusMetricsProtocolHTTPS InputPrometheusMetricsProtocol = "https"
)

func (e InputPrometheusMetricsProtocol) ToPointer() *InputPrometheusMetricsProtocol {
	return &e
}

func (e *InputPrometheusMetricsProtocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "http":
		fallthrough
	case "https":
		*e = InputPrometheusMetricsProtocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusMetricsProtocol: %v", v)
	}
}

type InputPrometheusSearchFilter struct {
	// Search filter attribute name, see: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstances.html for more information. Attributes can be manually entered if not present in the drop down list
	Name string `json:"Name"`
	// Search Filter Values, if empty only "running" EC2 instances will be returned
	Values []string `json:"Values"`
}

// InputPrometheusSignatureVersion - Signature version to use for signing EC2 requests.
type InputPrometheusSignatureVersion string

const (
	InputPrometheusSignatureVersionV2 InputPrometheusSignatureVersion = "v2"
	InputPrometheusSignatureVersionV4 InputPrometheusSignatureVersion = "v4"
)

func (e InputPrometheusSignatureVersion) ToPointer() *InputPrometheusSignatureVersion {
	return &e
}

func (e *InputPrometheusSignatureVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "v2":
		fallthrough
	case "v4":
		*e = InputPrometheusSignatureVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusSignatureVersion: %v", v)
	}
}

type InputPrometheusInputType string

const (
	InputPrometheusInputTypeSplunk            InputPrometheusInputType = "splunk"
	InputPrometheusInputTypeSplunkHec         InputPrometheusInputType = "splunk_hec"
	InputPrometheusInputTypeSyslog            InputPrometheusInputType = "syslog"
	InputPrometheusInputTypeTcpjson           InputPrometheusInputType = "tcpjson"
	InputPrometheusInputTypeGrafana           InputPrometheusInputType = "grafana"
	InputPrometheusInputTypeLoki              InputPrometheusInputType = "loki"
	InputPrometheusInputTypeHTTP              InputPrometheusInputType = "http"
	InputPrometheusInputTypeHTTPRaw           InputPrometheusInputType = "http_raw"
	InputPrometheusInputTypeFirehose          InputPrometheusInputType = "firehose"
	InputPrometheusInputTypeElastic           InputPrometheusInputType = "elastic"
	InputPrometheusInputTypeKafka             InputPrometheusInputType = "kafka"
	InputPrometheusInputTypeConfluentCloud    InputPrometheusInputType = "confluent_cloud"
	InputPrometheusInputTypeMsk               InputPrometheusInputType = "msk"
	InputPrometheusInputTypeKinesis           InputPrometheusInputType = "kinesis"
	InputPrometheusInputTypeEventhub          InputPrometheusInputType = "eventhub"
	InputPrometheusInputTypeAzureBlob         InputPrometheusInputType = "azure_blob"
	InputPrometheusInputTypeMetrics           InputPrometheusInputType = "metrics"
	InputPrometheusInputTypeSqs               InputPrometheusInputType = "sqs"
	InputPrometheusInputTypeS3                InputPrometheusInputType = "s3"
	InputPrometheusInputTypeSnmp              InputPrometheusInputType = "snmp"
	InputPrometheusInputTypeCrowdstrike       InputPrometheusInputType = "crowdstrike"
	InputPrometheusInputTypeTCP               InputPrometheusInputType = "tcp"
	InputPrometheusInputTypeRawUDP            InputPrometheusInputType = "raw_udp"
	InputPrometheusInputTypeOffice365Service  InputPrometheusInputType = "office365_service"
	InputPrometheusInputTypeOffice365Mgmt     InputPrometheusInputType = "office365_mgmt"
	InputPrometheusInputTypeOffice365MsgTrace InputPrometheusInputType = "office365_msg_trace"
	InputPrometheusInputTypePrometheus        InputPrometheusInputType = "prometheus"
	InputPrometheusInputTypeEdgePrometheus    InputPrometheusInputType = "edge_prometheus"
	InputPrometheusInputTypePrometheusRw      InputPrometheusInputType = "prometheus_rw"
	InputPrometheusInputTypeAppscope          InputPrometheusInputType = "appscope"
	InputPrometheusInputTypeGooglePubsub      InputPrometheusInputType = "google_pubsub"
	InputPrometheusInputTypeOpenTelemetry     InputPrometheusInputType = "open_telemetry"
	InputPrometheusInputTypeDatadogAgent      InputPrometheusInputType = "datadog_agent"
	InputPrometheusInputTypeWef               InputPrometheusInputType = "wef"
	InputPrometheusInputTypeDatagen           InputPrometheusInputType = "datagen"
	InputPrometheusInputTypeCribl             InputPrometheusInputType = "cribl"
	InputPrometheusInputTypeCriblmetrics      InputPrometheusInputType = "criblmetrics"
	InputPrometheusInputTypeCriblHTTP         InputPrometheusInputType = "cribl_http"
	InputPrometheusInputTypeCriblTCP          InputPrometheusInputType = "cribl_tcp"
	InputPrometheusInputTypeWinEventLogs      InputPrometheusInputType = "win_event_logs"
	InputPrometheusInputTypeSystemMetrics     InputPrometheusInputType = "system_metrics"
	InputPrometheusInputTypeWindowsMetrics    InputPrometheusInputType = "windows_metrics"
	InputPrometheusInputTypeSystemState       InputPrometheusInputType = "system_state"
	InputPrometheusInputTypeKubeMetrics       InputPrometheusInputType = "kube_metrics"
	InputPrometheusInputTypeKubeLogs          InputPrometheusInputType = "kube_logs"
	InputPrometheusInputTypeKubeEvents        InputPrometheusInputType = "kube_events"
	InputPrometheusInputTypeExec              InputPrometheusInputType = "exec"
	InputPrometheusInputTypeSplunkSearch      InputPrometheusInputType = "splunk_search"
	InputPrometheusInputTypeFile              InputPrometheusInputType = "file"
	InputPrometheusInputTypeJournalFiles      InputPrometheusInputType = "journal_files"
)

func (e InputPrometheusInputType) ToPointer() *InputPrometheusInputType {
	return &e
}

func (e *InputPrometheusInputType) UnmarshalJSON(data []byte) error {
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
		*e = InputPrometheusInputType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusInputType: %v", v)
	}
}

type InputPrometheus struct {
	// Amazon Resource Name (ARN) of the role to assume
	AssumeRoleArn *string `json:"assumeRoleArn,omitempty"`
	// External ID to use when assuming role
	AssumeRoleExternalID *string `json:"assumeRoleExternalId,omitempty"`
	// Enter credentials directly, or select a stored secret
	AuthType *InputPrometheusAuthenticationMethod `json:"authType,omitempty"`
	// AWS authentication method. Choose Auto to use IAM roles.
	AwsAuthenticationMethod *InputPrometheusAuthenticationMethod1 `json:"awsAuthenticationMethod,omitempty"`
	// Secret key
	AwsSecretKey *string `json:"awsSecretKey,omitempty"`
	// Direct connections to Destinations, optionally via a Pipeline or a Pack.
	Connections []InputPrometheusConnections `json:"connections,omitempty"`
	// Select (or create) a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
	// Other dimensions to include in events
	DimensionList []string `json:"dimensionList,omitempty"`
	// Enable/disable this input
	Disabled *bool `json:"disabled,omitempty"`
	// Target discovery mechanism. Use static to manually enter a list of targets.
	DiscoveryType *InputPrometheusDiscoveryType `json:"discoveryType,omitempty"`
	// Use Assume Role credentials to access EC2
	EnableAssumeRole *bool `json:"enableAssumeRole,omitempty"`
	// EC2 service endpoint. If empty, defaults to AWS' Region-specific endpoint. Otherwise, it must point to EC2-compatible endpoint.
	Endpoint *string `json:"endpoint,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Unique ID for this input
	ID *string `json:"id,omitempty"`
	// How often in minutes to scrape targets for metrics, 60 must be evenly divisible by the value or save will fail.
	Interval int64 `json:"interval"`
	// Maximum time the job is allowed to run (e.g., 30, 45s or 15m). Units are seconds, if not specified. Enter 0 for unlimited time.
	JobTimeout *string `json:"jobTimeout,omitempty"`
	// How often workers should check in with the scheduler to keep job subscription alive
	KeepAliveTime *int64 `json:"keepAliveTime,omitempty"`
	// Collector runtime Log Level
	LogLevel InputPrometheusLogLevel `json:"logLevel"`
	// The number of Keep Alive Time periods before an inactive worker will have its job subscription revoked.
	MaxMissedKeepAlives *int64 `json:"maxMissedKeepAlives,omitempty"`
	// Fields to add to events from this input.
	Metadata []InputPrometheusMetadata `json:"metadata,omitempty"`
	// List of DNS names to resolve
	NameList                       *InputPrometheusDNSNames                       `json:"nameList,omitempty"`
	Names                          *string                                        `json:"names,omitempty"`
	OptionalFieldsInGeneralSection *InputPrometheusOptionalFieldsInGeneralSection `json:"optionalFieldsInGeneralSection,omitempty"`
	// Password for Prometheus Basic authentication
	Password *string `json:"password,omitempty"`
	// Pipeline to process data from this Source before sending it through the Routes.
	Pipeline *string            `json:"pipeline,omitempty"`
	Pq       *InputPrometheusPq `json:"pq,omitempty"`
	// For details on Persistent Queues, see: [https://docs.cribl.io/stream/persistent-queues](https://docs.cribl.io/stream/persistent-queues)
	PqEnabled *bool `json:"pqEnabled,omitempty"`
	// DNS Record type to resolve
	RecordType *InputPrometheusRecordType `json:"recordType,omitempty"`
	// Region where the EC2 is located
	Region *InputPrometheusRegion `json:"region,omitempty"`
	// Whether to reject certificates that cannot be verified against a valid CA (e.g., self-signed certificates).
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// Whether to reuse connections between requests, which can improve performance.
	ReuseConnections *bool `json:"reuseConnections,omitempty"`
	// Path to use when collecting metrics from discovered targets
	ScrapePath *string `json:"scrapePath,omitempty"`
	// The port number in the metrics URL for discovered targets.
	ScrapePort *int64 `json:"scrapePort,omitempty"`
	// Protocol to use when collecting metrics
	ScrapeProtocol *InputPrometheusMetricsProtocol `json:"scrapeProtocol,omitempty"`
	// EC2 Instance Search Filter
	SearchFilter []InputPrometheusSearchFilter `json:"searchFilter,omitempty"`
	// Select whether to send data to Routes, or directly to Destinations.
	SendToRoutes *bool `json:"sendToRoutes,omitempty"`
	// Signature version to use for signing EC2 requests.
	SignatureVersion *InputPrometheusSignatureVersion `json:"signatureVersion,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// List of Prometheus targets to pull metrics from. Values can be in URL or host[:port] format. For example: http://localhost:9090/metrics, localhost:9090, or localhost. In cases where just host[:port] is specified, the endpoint will resolve to 'http://host[:port]/metrics'.
	Tarlist []string                  `json:"tarlist,omitempty"`
	Type    *InputPrometheusInputType `json:"type,omitempty"`
	// Use public IP address for discovered targets. Set to false if the private IP address should be used.
	UsePublicIP *bool `json:"usePublicIp,omitempty"`
	// Username for Prometheus Basic authentication
	Username *string `json:"username,omitempty"`
}
