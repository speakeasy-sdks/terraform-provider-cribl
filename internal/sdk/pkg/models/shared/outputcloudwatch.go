// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OutputCloudwatchAuthenticationMethod - AWS authentication method. Choose Auto to use IAM roles.
type OutputCloudwatchAuthenticationMethod string

const (
	OutputCloudwatchAuthenticationMethodSecret OutputCloudwatchAuthenticationMethod = "secret"
	OutputCloudwatchAuthenticationMethodManual OutputCloudwatchAuthenticationMethod = "manual"
)

func (e OutputCloudwatchAuthenticationMethod) ToPointer() *OutputCloudwatchAuthenticationMethod {
	return &e
}

func (e *OutputCloudwatchAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "secret":
		fallthrough
	case "manual":
		*e = OutputCloudwatchAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputCloudwatchAuthenticationMethod: %v", v)
	}
}

// OutputCloudwatchBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputCloudwatchBackpressureBehavior string

const (
	OutputCloudwatchBackpressureBehaviorQueue OutputCloudwatchBackpressureBehavior = "queue"
	OutputCloudwatchBackpressureBehaviorDrop  OutputCloudwatchBackpressureBehavior = "drop"
	OutputCloudwatchBackpressureBehaviorBlock OutputCloudwatchBackpressureBehavior = "block"
)

func (e OutputCloudwatchBackpressureBehavior) ToPointer() *OutputCloudwatchBackpressureBehavior {
	return &e
}

func (e *OutputCloudwatchBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputCloudwatchBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputCloudwatchBackpressureBehavior: %v", v)
	}
}

// OutputCloudwatchCompression - Codec to use to compress the persisted data.
type OutputCloudwatchCompression string

const (
	OutputCloudwatchCompressionNone OutputCloudwatchCompression = "none"
	OutputCloudwatchCompressionGzip OutputCloudwatchCompression = "gzip"
)

func (e OutputCloudwatchCompression) ToPointer() *OutputCloudwatchCompression {
	return &e
}

func (e *OutputCloudwatchCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputCloudwatchCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputCloudwatchCompression: %v", v)
	}
}

type OutputCloudwatchPqControls struct {
}

// OutputCloudwatchQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputCloudwatchQueueFullBehavior string

const (
	OutputCloudwatchQueueFullBehaviorBlock OutputCloudwatchQueueFullBehavior = "block"
	OutputCloudwatchQueueFullBehaviorDrop  OutputCloudwatchQueueFullBehavior = "drop"
)

func (e OutputCloudwatchQueueFullBehavior) ToPointer() *OutputCloudwatchQueueFullBehavior {
	return &e
}

func (e *OutputCloudwatchQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputCloudwatchQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputCloudwatchQueueFullBehavior: %v", v)
	}
}

// OutputCloudwatchRegion - Region where the CloudWatchLogs is located
type OutputCloudwatchRegion string

const (
	OutputCloudwatchRegionUsEast1      OutputCloudwatchRegion = "us-east-1"
	OutputCloudwatchRegionUsEast2      OutputCloudwatchRegion = "us-east-2"
	OutputCloudwatchRegionUsWest1      OutputCloudwatchRegion = "us-west-1"
	OutputCloudwatchRegionUsWest2      OutputCloudwatchRegion = "us-west-2"
	OutputCloudwatchRegionAfSouth1     OutputCloudwatchRegion = "af-south-1"
	OutputCloudwatchRegionCaCentral1   OutputCloudwatchRegion = "ca-central-1"
	OutputCloudwatchRegionEuWest1      OutputCloudwatchRegion = "eu-west-1"
	OutputCloudwatchRegionEuCentral1   OutputCloudwatchRegion = "eu-central-1"
	OutputCloudwatchRegionEuWest2      OutputCloudwatchRegion = "eu-west-2"
	OutputCloudwatchRegionEuSouth1     OutputCloudwatchRegion = "eu-south-1"
	OutputCloudwatchRegionEuWest3      OutputCloudwatchRegion = "eu-west-3"
	OutputCloudwatchRegionEuNorth1     OutputCloudwatchRegion = "eu-north-1"
	OutputCloudwatchRegionApEast1      OutputCloudwatchRegion = "ap-east-1"
	OutputCloudwatchRegionApNortheast1 OutputCloudwatchRegion = "ap-northeast-1"
	OutputCloudwatchRegionApNortheast2 OutputCloudwatchRegion = "ap-northeast-2"
	OutputCloudwatchRegionApSoutheast1 OutputCloudwatchRegion = "ap-southeast-1"
	OutputCloudwatchRegionApSoutheast2 OutputCloudwatchRegion = "ap-southeast-2"
	OutputCloudwatchRegionApSouth1     OutputCloudwatchRegion = "ap-south-1"
	OutputCloudwatchRegionMeSouth1     OutputCloudwatchRegion = "me-south-1"
	OutputCloudwatchRegionSaEast1      OutputCloudwatchRegion = "sa-east-1"
	OutputCloudwatchRegionUsGovEast1   OutputCloudwatchRegion = "us-gov-east-1"
	OutputCloudwatchRegionUsGovWest1   OutputCloudwatchRegion = "us-gov-west-1"
)

func (e OutputCloudwatchRegion) ToPointer() *OutputCloudwatchRegion {
	return &e
}

func (e *OutputCloudwatchRegion) UnmarshalJSON(data []byte) error {
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
		*e = OutputCloudwatchRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputCloudwatchRegion: %v", v)
	}
}

// OutputCloudwatchSignatureVersion - Signature version to use for signing CloudWatchLogs requests.
type OutputCloudwatchSignatureVersion string

const (
	OutputCloudwatchSignatureVersionV2 OutputCloudwatchSignatureVersion = "v2"
	OutputCloudwatchSignatureVersionV4 OutputCloudwatchSignatureVersion = "v4"
)

func (e OutputCloudwatchSignatureVersion) ToPointer() *OutputCloudwatchSignatureVersion {
	return &e
}

func (e *OutputCloudwatchSignatureVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "v2":
		fallthrough
	case "v4":
		*e = OutputCloudwatchSignatureVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputCloudwatchSignatureVersion: %v", v)
	}
}

type OutputCloudwatchType string

const (
	OutputCloudwatchTypeCloudwatch OutputCloudwatchType = "cloudwatch"
)

func (e OutputCloudwatchType) ToPointer() *OutputCloudwatchType {
	return &e
}

func (e *OutputCloudwatchType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cloudwatch":
		*e = OutputCloudwatchType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputCloudwatchType: %v", v)
	}
}

type OutputCloudwatch struct {
	// Amazon Resource Name (ARN) of the role to assume
	AssumeRoleArn *string `json:"assumeRoleArn,omitempty"`
	// External ID to use when assuming role
	AssumeRoleExternalID *string `json:"assumeRoleExternalId,omitempty"`
	// Access key
	AwsAPIKey *string `json:"awsApiKey,omitempty"`
	// AWS authentication method. Choose Auto to use IAM roles.
	AwsAuthenticationMethod *OutputCloudwatchAuthenticationMethod `json:"awsAuthenticationMethod,omitempty"`
	// Select (or create) a stored secret that references your access key and secret key.
	AwsSecret *string `json:"awsSecret,omitempty"`
	// Secret key
	AwsSecretKey *string `json:"awsSecretKey,omitempty"`
	// Use Assume Role credentials to access CloudWatchLogs
	EnableAssumeRole *bool `json:"enableAssumeRole,omitempty"`
	// CloudWatchLogs service endpoint. If empty, defaults to AWS' Region-specific endpoint. Otherwise, it must point to CloudWatchLogs-compatible endpoint.
	Endpoint *string `json:"endpoint,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Max record size.
	FlushPeriodSec *int64 `json:"flushPeriodSec,omitempty"`
	// Unique ID for this output
	ID *string `json:"id,omitempty"`
	// CloudWatch log group to associate events with
	LogGroupName string `json:"logGroupName"`
	// Prefix for CloudWatch log stream name. This prefix will be used to generate a unique log stream name per cribl instance, for example: myStream_myHost_myOutputId
	LogStreamName string `json:"logStreamName"`
	// Maximum number of queued batches before blocking
	MaxQueueSize *int64 `json:"maxQueueSize,omitempty"`
	// Maximum size (KB) of each individual record before compression. For non compressible data 1MB is the max recommended size
	MaxRecordSizeKB *int64 `json:"maxRecordSizeKB,omitempty"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputCloudwatchBackpressureBehavior `json:"onBackpressure,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputCloudwatchCompression `json:"pqCompress,omitempty"`
	PqControls *OutputCloudwatchPqControls  `json:"pqControls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `json:"pqMaxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `json:"pqMaxSize,omitempty"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputCloudwatchQueueFullBehavior `json:"pqOnBackpressure,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `json:"pqPath,omitempty"`
	// Toggle this off to forward new events to receiver(s) before queue is flushed. Otherwise, default drain behavior is FIFO (first in, first out).
	PqStrictOrdering *bool `json:"pqStrictOrdering,omitempty"`
	// Region where the CloudWatchLogs is located
	Region OutputCloudwatchRegion `json:"region"`
	// Whether to reject certificates that cannot be verified against a valid CA (e.g., self-signed certificates).
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// Whether to reuse connections between requests, which can improve performance.
	ReuseConnections *bool `json:"reuseConnections,omitempty"`
	// Signature version to use for signing CloudWatchLogs requests.
	SignatureVersion *OutputCloudwatchSignatureVersion `json:"signatureVersion,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string              `json:"systemFields,omitempty"`
	Type         *OutputCloudwatchType `json:"type,omitempty"`
}
