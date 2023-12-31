// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// InputEdgePrometheusAuthenticationMethod - Enter credentials directly, or select a stored secret
type InputEdgePrometheusAuthenticationMethod string

const (
	InputEdgePrometheusAuthenticationMethodKubernetes InputEdgePrometheusAuthenticationMethod = "kubernetes"
	InputEdgePrometheusAuthenticationMethodSecret     InputEdgePrometheusAuthenticationMethod = "secret"
)

func (e InputEdgePrometheusAuthenticationMethod) ToPointer() *InputEdgePrometheusAuthenticationMethod {
	return &e
}

func (e *InputEdgePrometheusAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "kubernetes":
		fallthrough
	case "secret":
		*e = InputEdgePrometheusAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputEdgePrometheusAuthenticationMethod: %v", v)
	}
}

// InputEdgePrometheusAuthenticationMethod1 - AWS authentication method. Choose Auto to use IAM roles.
type InputEdgePrometheusAuthenticationMethod1 string

const (
	InputEdgePrometheusAuthenticationMethod1Auto   InputEdgePrometheusAuthenticationMethod1 = "auto"
	InputEdgePrometheusAuthenticationMethod1Manual InputEdgePrometheusAuthenticationMethod1 = "manual"
	InputEdgePrometheusAuthenticationMethod1Secret InputEdgePrometheusAuthenticationMethod1 = "secret"
)

func (e InputEdgePrometheusAuthenticationMethod1) ToPointer() *InputEdgePrometheusAuthenticationMethod1 {
	return &e
}

func (e *InputEdgePrometheusAuthenticationMethod1) UnmarshalJSON(data []byte) error {
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
		*e = InputEdgePrometheusAuthenticationMethod1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputEdgePrometheusAuthenticationMethod1: %v", v)
	}
}

type InputEdgePrometheusConnections struct {
	// Select a Destination.
	Output string `json:"output"`
	// Select Pipeline or Pack. Optional.
	Pipeline *string `json:"pipeline,omitempty"`
}

// InputEdgePrometheusDiscoveryType - Target discovery mechanism. Use static to manually enter a list of targets.
type InputEdgePrometheusDiscoveryType string

const (
	InputEdgePrometheusDiscoveryTypeK8sPods InputEdgePrometheusDiscoveryType = "k8s-pods"
	InputEdgePrometheusDiscoveryTypeDNS     InputEdgePrometheusDiscoveryType = "dns"
	InputEdgePrometheusDiscoveryTypeEc2     InputEdgePrometheusDiscoveryType = "ec2"
	InputEdgePrometheusDiscoveryTypeK8sNode InputEdgePrometheusDiscoveryType = "k8s-node"
)

func (e InputEdgePrometheusDiscoveryType) ToPointer() *InputEdgePrometheusDiscoveryType {
	return &e
}

func (e *InputEdgePrometheusDiscoveryType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "k8s-pods":
		fallthrough
	case "dns":
		fallthrough
	case "ec2":
		fallthrough
	case "k8s-node":
		*e = InputEdgePrometheusDiscoveryType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputEdgePrometheusDiscoveryType: %v", v)
	}
}

type InputEdgePrometheusMetadata struct {
	// Field name
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

// InputEdgePrometheusDNSNames - List of DNS names to resolve
type InputEdgePrometheusDNSNames struct {
}

type InputEdgePrometheusOptionalFieldsInGeneralSection string

const (
	InputEdgePrometheusOptionalFieldsInGeneralSectionRegion       InputEdgePrometheusOptionalFieldsInGeneralSection = "region"
	InputEdgePrometheusOptionalFieldsInGeneralSectionUsePublicIP  InputEdgePrometheusOptionalFieldsInGeneralSection = "usePublicIp"
	InputEdgePrometheusOptionalFieldsInGeneralSectionSearchFilter InputEdgePrometheusOptionalFieldsInGeneralSection = "searchFilter"
)

func (e InputEdgePrometheusOptionalFieldsInGeneralSection) ToPointer() *InputEdgePrometheusOptionalFieldsInGeneralSection {
	return &e
}

func (e *InputEdgePrometheusOptionalFieldsInGeneralSection) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "region":
		fallthrough
	case "usePublicIp":
		fallthrough
	case "searchFilter":
		*e = InputEdgePrometheusOptionalFieldsInGeneralSection(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputEdgePrometheusOptionalFieldsInGeneralSection: %v", v)
	}
}

type InputEdgePrometheusPodFilter struct {
	// Optional description of this rule's purpose
	Description *string `json:"description,omitempty"`
	// JavaScript expression applied to pods objects. Return 'true' to include it.
	Filter string `json:"filter"`
}

// InputEdgePrometheusPqCompression - Codec to use to compress the persisted data.
type InputEdgePrometheusPqCompression string

const (
	InputEdgePrometheusPqCompressionNone InputEdgePrometheusPqCompression = "none"
	InputEdgePrometheusPqCompressionGzip InputEdgePrometheusPqCompression = "gzip"
)

func (e InputEdgePrometheusPqCompression) ToPointer() *InputEdgePrometheusPqCompression {
	return &e
}

func (e *InputEdgePrometheusPqCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputEdgePrometheusPqCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputEdgePrometheusPqCompression: %v", v)
	}
}

// InputEdgePrometheusPqMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputEdgePrometheusPqMode string

const (
	InputEdgePrometheusPqModeSmart  InputEdgePrometheusPqMode = "smart"
	InputEdgePrometheusPqModeAlways InputEdgePrometheusPqMode = "always"
)

func (e InputEdgePrometheusPqMode) ToPointer() *InputEdgePrometheusPqMode {
	return &e
}

func (e *InputEdgePrometheusPqMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputEdgePrometheusPqMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputEdgePrometheusPqMode: %v", v)
	}
}

type InputEdgePrometheusPq struct {
	// The number of events to send downstream before committing that Stream has read them.
	CommitFrequency *int64 `json:"commitFrequency,omitempty"`
	// Codec to use to compress the persisted data.
	Compress *InputEdgePrometheusPqCompression `json:"compress,omitempty"`
	// The maximum number of events to hold in memory before writing the events to disk.
	MaxBufferSize *int64 `json:"maxBufferSize,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	MaxFileSize *string `json:"maxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	MaxSize *string `json:"maxSize,omitempty"`
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputEdgePrometheusPqMode `json:"mode,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/inputs/<input-id>.
	Path *string `json:"path,omitempty"`
}

// InputEdgePrometheusRecordType - DNS Record type to resolve
type InputEdgePrometheusRecordType string

const (
	InputEdgePrometheusRecordTypeSrv  InputEdgePrometheusRecordType = "SRV"
	InputEdgePrometheusRecordTypeA    InputEdgePrometheusRecordType = "A"
	InputEdgePrometheusRecordTypeAaaa InputEdgePrometheusRecordType = "AAAA"
)

func (e InputEdgePrometheusRecordType) ToPointer() *InputEdgePrometheusRecordType {
	return &e
}

func (e *InputEdgePrometheusRecordType) UnmarshalJSON(data []byte) error {
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
		*e = InputEdgePrometheusRecordType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputEdgePrometheusRecordType: %v", v)
	}
}

// InputEdgePrometheusRegion - Region where the EC2 is located
type InputEdgePrometheusRegion string

const (
	InputEdgePrometheusRegionUsEast1      InputEdgePrometheusRegion = "us-east-1"
	InputEdgePrometheusRegionUsEast2      InputEdgePrometheusRegion = "us-east-2"
	InputEdgePrometheusRegionUsWest1      InputEdgePrometheusRegion = "us-west-1"
	InputEdgePrometheusRegionUsWest2      InputEdgePrometheusRegion = "us-west-2"
	InputEdgePrometheusRegionAfSouth1     InputEdgePrometheusRegion = "af-south-1"
	InputEdgePrometheusRegionCaCentral1   InputEdgePrometheusRegion = "ca-central-1"
	InputEdgePrometheusRegionEuWest1      InputEdgePrometheusRegion = "eu-west-1"
	InputEdgePrometheusRegionEuCentral1   InputEdgePrometheusRegion = "eu-central-1"
	InputEdgePrometheusRegionEuWest2      InputEdgePrometheusRegion = "eu-west-2"
	InputEdgePrometheusRegionEuSouth1     InputEdgePrometheusRegion = "eu-south-1"
	InputEdgePrometheusRegionEuWest3      InputEdgePrometheusRegion = "eu-west-3"
	InputEdgePrometheusRegionEuNorth1     InputEdgePrometheusRegion = "eu-north-1"
	InputEdgePrometheusRegionApEast1      InputEdgePrometheusRegion = "ap-east-1"
	InputEdgePrometheusRegionApNortheast1 InputEdgePrometheusRegion = "ap-northeast-1"
	InputEdgePrometheusRegionApNortheast2 InputEdgePrometheusRegion = "ap-northeast-2"
	InputEdgePrometheusRegionApSoutheast1 InputEdgePrometheusRegion = "ap-southeast-1"
	InputEdgePrometheusRegionApSoutheast2 InputEdgePrometheusRegion = "ap-southeast-2"
	InputEdgePrometheusRegionApSouth1     InputEdgePrometheusRegion = "ap-south-1"
	InputEdgePrometheusRegionMeSouth1     InputEdgePrometheusRegion = "me-south-1"
	InputEdgePrometheusRegionSaEast1      InputEdgePrometheusRegion = "sa-east-1"
	InputEdgePrometheusRegionUsGovEast1   InputEdgePrometheusRegion = "us-gov-east-1"
	InputEdgePrometheusRegionUsGovWest1   InputEdgePrometheusRegion = "us-gov-west-1"
)

func (e InputEdgePrometheusRegion) ToPointer() *InputEdgePrometheusRegion {
	return &e
}

func (e *InputEdgePrometheusRegion) UnmarshalJSON(data []byte) error {
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
		*e = InputEdgePrometheusRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputEdgePrometheusRegion: %v", v)
	}
}

// InputEdgePrometheusProtocol - Protocol to use when collecting metrics
type InputEdgePrometheusProtocol string

const (
	InputEdgePrometheusProtocolHTTP  InputEdgePrometheusProtocol = "http"
	InputEdgePrometheusProtocolHTTPS InputEdgePrometheusProtocol = "https"
)

func (e InputEdgePrometheusProtocol) ToPointer() *InputEdgePrometheusProtocol {
	return &e
}

func (e *InputEdgePrometheusProtocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "http":
		fallthrough
	case "https":
		*e = InputEdgePrometheusProtocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputEdgePrometheusProtocol: %v", v)
	}
}

type InputEdgePrometheusSearchFilter struct {
	// Search filter attribute name, see: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstances.html for more information. Attributes can be manually entered if not present in the drop down list
	Name string `json:"Name"`
	// Search Filter Values, if empty only "running" EC2 instances will be returned
	Values []string `json:"Values"`
}

// InputEdgePrometheusSignatureVersion - Signature version to use for signing EC2 requests.
type InputEdgePrometheusSignatureVersion string

const (
	InputEdgePrometheusSignatureVersionV2 InputEdgePrometheusSignatureVersion = "v2"
	InputEdgePrometheusSignatureVersionV4 InputEdgePrometheusSignatureVersion = "v4"
)

func (e InputEdgePrometheusSignatureVersion) ToPointer() *InputEdgePrometheusSignatureVersion {
	return &e
}

func (e *InputEdgePrometheusSignatureVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "v2":
		fallthrough
	case "v4":
		*e = InputEdgePrometheusSignatureVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputEdgePrometheusSignatureVersion: %v", v)
	}
}

// InputEdgePrometheusTargetsProtocol - Protocol to use when collecting metrics
type InputEdgePrometheusTargetsProtocol string

const (
	InputEdgePrometheusTargetsProtocolHTTP  InputEdgePrometheusTargetsProtocol = "http"
	InputEdgePrometheusTargetsProtocolHTTPS InputEdgePrometheusTargetsProtocol = "https"
)

func (e InputEdgePrometheusTargetsProtocol) ToPointer() *InputEdgePrometheusTargetsProtocol {
	return &e
}

func (e *InputEdgePrometheusTargetsProtocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "http":
		fallthrough
	case "https":
		*e = InputEdgePrometheusTargetsProtocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputEdgePrometheusTargetsProtocol: %v", v)
	}
}

type InputEdgePrometheusTargets struct {
	// Name of host from which to pull metrics.
	Host string `json:"host"`
	// Path to use when collecting metrics from discovered targets
	Path *string `json:"path,omitempty"`
	// The port number in the metrics URL for discovered targets.
	Port *int64 `json:"port,omitempty"`
	// Protocol to use when collecting metrics
	Protocol *InputEdgePrometheusTargetsProtocol `json:"protocol,omitempty"`
}

type InputEdgePrometheusInputType string

const (
	InputEdgePrometheusInputTypeSplunk            InputEdgePrometheusInputType = "splunk"
	InputEdgePrometheusInputTypeSplunkHec         InputEdgePrometheusInputType = "splunk_hec"
	InputEdgePrometheusInputTypeSyslog            InputEdgePrometheusInputType = "syslog"
	InputEdgePrometheusInputTypeTcpjson           InputEdgePrometheusInputType = "tcpjson"
	InputEdgePrometheusInputTypeGrafana           InputEdgePrometheusInputType = "grafana"
	InputEdgePrometheusInputTypeLoki              InputEdgePrometheusInputType = "loki"
	InputEdgePrometheusInputTypeHTTP              InputEdgePrometheusInputType = "http"
	InputEdgePrometheusInputTypeHTTPRaw           InputEdgePrometheusInputType = "http_raw"
	InputEdgePrometheusInputTypeFirehose          InputEdgePrometheusInputType = "firehose"
	InputEdgePrometheusInputTypeElastic           InputEdgePrometheusInputType = "elastic"
	InputEdgePrometheusInputTypeKafka             InputEdgePrometheusInputType = "kafka"
	InputEdgePrometheusInputTypeConfluentCloud    InputEdgePrometheusInputType = "confluent_cloud"
	InputEdgePrometheusInputTypeMsk               InputEdgePrometheusInputType = "msk"
	InputEdgePrometheusInputTypeKinesis           InputEdgePrometheusInputType = "kinesis"
	InputEdgePrometheusInputTypeEventhub          InputEdgePrometheusInputType = "eventhub"
	InputEdgePrometheusInputTypeAzureBlob         InputEdgePrometheusInputType = "azure_blob"
	InputEdgePrometheusInputTypeMetrics           InputEdgePrometheusInputType = "metrics"
	InputEdgePrometheusInputTypeSqs               InputEdgePrometheusInputType = "sqs"
	InputEdgePrometheusInputTypeS3                InputEdgePrometheusInputType = "s3"
	InputEdgePrometheusInputTypeSnmp              InputEdgePrometheusInputType = "snmp"
	InputEdgePrometheusInputTypeCrowdstrike       InputEdgePrometheusInputType = "crowdstrike"
	InputEdgePrometheusInputTypeTCP               InputEdgePrometheusInputType = "tcp"
	InputEdgePrometheusInputTypeRawUDP            InputEdgePrometheusInputType = "raw_udp"
	InputEdgePrometheusInputTypeOffice365Service  InputEdgePrometheusInputType = "office365_service"
	InputEdgePrometheusInputTypeOffice365Mgmt     InputEdgePrometheusInputType = "office365_mgmt"
	InputEdgePrometheusInputTypeOffice365MsgTrace InputEdgePrometheusInputType = "office365_msg_trace"
	InputEdgePrometheusInputTypePrometheus        InputEdgePrometheusInputType = "prometheus"
	InputEdgePrometheusInputTypeEdgePrometheus    InputEdgePrometheusInputType = "edge_prometheus"
	InputEdgePrometheusInputTypePrometheusRw      InputEdgePrometheusInputType = "prometheus_rw"
	InputEdgePrometheusInputTypeAppscope          InputEdgePrometheusInputType = "appscope"
	InputEdgePrometheusInputTypeGooglePubsub      InputEdgePrometheusInputType = "google_pubsub"
	InputEdgePrometheusInputTypeOpenTelemetry     InputEdgePrometheusInputType = "open_telemetry"
	InputEdgePrometheusInputTypeDatadogAgent      InputEdgePrometheusInputType = "datadog_agent"
	InputEdgePrometheusInputTypeWef               InputEdgePrometheusInputType = "wef"
	InputEdgePrometheusInputTypeDatagen           InputEdgePrometheusInputType = "datagen"
	InputEdgePrometheusInputTypeCribl             InputEdgePrometheusInputType = "cribl"
	InputEdgePrometheusInputTypeCriblmetrics      InputEdgePrometheusInputType = "criblmetrics"
	InputEdgePrometheusInputTypeCriblHTTP         InputEdgePrometheusInputType = "cribl_http"
	InputEdgePrometheusInputTypeCriblTCP          InputEdgePrometheusInputType = "cribl_tcp"
	InputEdgePrometheusInputTypeWinEventLogs      InputEdgePrometheusInputType = "win_event_logs"
	InputEdgePrometheusInputTypeSystemMetrics     InputEdgePrometheusInputType = "system_metrics"
	InputEdgePrometheusInputTypeWindowsMetrics    InputEdgePrometheusInputType = "windows_metrics"
	InputEdgePrometheusInputTypeSystemState       InputEdgePrometheusInputType = "system_state"
	InputEdgePrometheusInputTypeKubeMetrics       InputEdgePrometheusInputType = "kube_metrics"
	InputEdgePrometheusInputTypeKubeLogs          InputEdgePrometheusInputType = "kube_logs"
	InputEdgePrometheusInputTypeKubeEvents        InputEdgePrometheusInputType = "kube_events"
	InputEdgePrometheusInputTypeExec              InputEdgePrometheusInputType = "exec"
	InputEdgePrometheusInputTypeSplunkSearch      InputEdgePrometheusInputType = "splunk_search"
	InputEdgePrometheusInputTypeFile              InputEdgePrometheusInputType = "file"
	InputEdgePrometheusInputTypeJournalFiles      InputEdgePrometheusInputType = "journal_files"
)

func (e InputEdgePrometheusInputType) ToPointer() *InputEdgePrometheusInputType {
	return &e
}

func (e *InputEdgePrometheusInputType) UnmarshalJSON(data []byte) error {
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
		*e = InputEdgePrometheusInputType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputEdgePrometheusInputType: %v", v)
	}
}

type InputEdgePrometheus struct {
	// Amazon Resource Name (ARN) of the role to assume
	AssumeRoleArn *string `json:"assumeRoleArn,omitempty"`
	// External ID to use when assuming role
	AssumeRoleExternalID *string `json:"assumeRoleExternalId,omitempty"`
	// Enter credentials directly, or select a stored secret
	AuthType *InputEdgePrometheusAuthenticationMethod `json:"authType,omitempty"`
	// AWS authentication method. Choose Auto to use IAM roles.
	AwsAuthenticationMethod *InputEdgePrometheusAuthenticationMethod1 `json:"awsAuthenticationMethod,omitempty"`
	// Secret key
	AwsSecretKey *string `json:"awsSecretKey,omitempty"`
	// Direct connections to Destinations, optionally via a Pipeline or a Pack.
	Connections []InputEdgePrometheusConnections `json:"connections,omitempty"`
	// Select (or create) a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
	// Other dimensions to include in events
	DimensionList []string `json:"dimensionList,omitempty"`
	// Enable/disable this input
	Disabled *bool `json:"disabled,omitempty"`
	// Target discovery mechanism. Use static to manually enter a list of targets.
	DiscoveryType InputEdgePrometheusDiscoveryType `json:"discoveryType"`
	// Use Assume Role credentials to access EC2
	EnableAssumeRole *bool `json:"enableAssumeRole,omitempty"`
	// EC2 service endpoint. If empty, defaults to AWS' Region-specific endpoint. Otherwise, it must point to EC2-compatible endpoint.
	Endpoint *string `json:"endpoint,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Unique ID for this input
	ID *string `json:"id,omitempty"`
	// How often in seconds to scrape targets for metrics.
	Interval int64 `json:"interval"`
	// Fields to add to events from this input.
	Metadata []InputEdgePrometheusMetadata `json:"metadata,omitempty"`
	// List of DNS names to resolve
	NameList                       *InputEdgePrometheusDNSNames                       `json:"nameList,omitempty"`
	Names                          *string                                            `json:"names,omitempty"`
	OptionalFieldsInGeneralSection *InputEdgePrometheusOptionalFieldsInGeneralSection `json:"optionalFieldsInGeneralSection,omitempty"`
	// Password for Prometheus Basic authentication
	Password *string `json:"password,omitempty"`
	// Pipeline to process data from this Source before sending it through the Routes.
	Pipeline *string `json:"pipeline,omitempty"`
	//   Add rules to decide which pods to discover for metrics.
	//   Pods are searched if no rules are given or of all the rules'
	//   expressions evaluate to true.
	//
	PodFilter []InputEdgePrometheusPodFilter `json:"podFilter,omitempty"`
	Pq        *InputEdgePrometheusPq         `json:"pq,omitempty"`
	// For details on Persistent Queues, see: [https://docs.cribl.io/stream/persistent-queues](https://docs.cribl.io/stream/persistent-queues)
	PqEnabled *bool `json:"pqEnabled,omitempty"`
	// DNS Record type to resolve
	RecordType *InputEdgePrometheusRecordType `json:"recordType,omitempty"`
	// Region where the EC2 is located
	Region *InputEdgePrometheusRegion `json:"region,omitempty"`
	// Whether to reject certificates that cannot be verified against a valid CA (e.g., self-signed certificates).
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// Whether to reuse connections between requests, which can improve performance.
	ReuseConnections *bool `json:"reuseConnections,omitempty"`
	// Path to use when collecting metrics from discovered targets
	ScrapePath *string `json:"scrapePath,omitempty"`
	// Path to use when collecting metrics from discovered targets
	ScrapePathExpr *string `json:"scrapePathExpr,omitempty"`
	// The port number in the metrics URL for discovered targets.
	ScrapePort *int64 `json:"scrapePort,omitempty"`
	// The port number in the metrics URL for discovered targets.
	ScrapePortExpr *string `json:"scrapePortExpr,omitempty"`
	// Protocol to use when collecting metrics
	ScrapeProtocol *InputEdgePrometheusProtocol `json:"scrapeProtocol,omitempty"`
	// Protocol to use when collecting metrics
	ScrapeProtocolExpr *string `json:"scrapeProtocolExpr,omitempty"`
	// EC2 Instance Search Filter
	SearchFilter []InputEdgePrometheusSearchFilter `json:"searchFilter,omitempty"`
	// Select whether to send data to Routes, or directly to Destinations.
	SendToRoutes *bool `json:"sendToRoutes,omitempty"`
	// Signature version to use for signing EC2 requests.
	SignatureVersion *InputEdgePrometheusSignatureVersion `json:"signatureVersion,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string                     `json:"streamtags,omitempty"`
	Targets    []InputEdgePrometheusTargets `json:"targets,omitempty"`
	// Timeout, in milliseconds, before aborting HTTP connection attempts; 1-60000 or 0 to disable
	Timeout *int64                        `json:"timeout,omitempty"`
	Type    *InputEdgePrometheusInputType `json:"type,omitempty"`
	// Use public IP address for discovered targets. Set to false if the private IP address should be used.
	UsePublicIP *bool `json:"usePublicIp,omitempty"`
	// Username for Prometheus Basic authentication
	Username *string `json:"username,omitempty"`
}
