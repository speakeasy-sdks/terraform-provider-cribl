// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SavedJobCollectionCollectorCollectorSpecificSettings struct {
}

type SavedJobCollectionCollector struct {
	Conf SavedJobCollectionCollectorCollectorSpecificSettings `json:"conf"`
	// If set to Yes, the collector will delete any files that it collects (where applicable).
	Destructive *bool `json:"destructive,omitempty"`
	// The type of collector to run.
	Type string `json:"type"`
}

type SavedJobCollectionInputMetadata struct {
	// Field name
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

type SavedJobCollectionInputPreprocess struct {
	// Arguments
	Args []string `json:"args,omitempty"`
	// Command to feed the data through (via stdin) and process its output (stdout)
	Command *string `json:"command,omitempty"`
	// Enable Custom Command
	Disabled bool `json:"disabled"`
}

type SavedJobCollectionInputType string

const (
	SavedJobCollectionInputTypeCollection SavedJobCollectionInputType = "collection"
)

func (e SavedJobCollectionInputType) ToPointer() *SavedJobCollectionInputType {
	return &e
}

func (e *SavedJobCollectionInputType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "collection":
		*e = SavedJobCollectionInputType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SavedJobCollectionInputType: %v", v)
	}
}

type SavedJobCollectionInput struct {
	// A list of event breaking rulesets that will be applied, in order, to the input data stream.
	BreakerRulesets []string `json:"breakerRulesets,omitempty"`
	// Fields to add to events from this input.
	Metadata []SavedJobCollectionInputMetadata `json:"metadata,omitempty"`
	// Destination to send results to.
	Output *string `json:"output,omitempty"`
	// Pipeline to process results.
	Pipeline   *string                            `json:"pipeline,omitempty"`
	Preprocess *SavedJobCollectionInputPreprocess `json:"preprocess,omitempty"`
	// If set to Yes, events will be sent to normal routing and event processing. Set to No if you want to select a specific Pipeline/Destination combination.
	SendToRoutes *bool `json:"sendToRoutes,omitempty"`
	// The amount of time (in milliseconds) the Event Breaker will wait for new data to be sent to a specific channel, before flushing the data stream out, as-is, to the Pipelines.
	StaleChannelFlushMs *int64 `json:"staleChannelFlushMs,omitempty"`
	// Rate (in bytes per second) to throttle while writing to an output. Also takes values with multiple-byte units, such as KB, MB, GB, etc. (E.g., 42 MB.) Default value of 0 specifies no throttling.
	ThrottleRatePerSec *string                      `json:"throttleRatePerSec,omitempty"`
	Type               *SavedJobCollectionInputType `json:"type,omitempty"`
}

// SavedJobCollectionScheduleRunSettingsLogLevel - Level at which to set task logging.
type SavedJobCollectionScheduleRunSettingsLogLevel string

const (
	SavedJobCollectionScheduleRunSettingsLogLevelError SavedJobCollectionScheduleRunSettingsLogLevel = "error"
	SavedJobCollectionScheduleRunSettingsLogLevelWarn  SavedJobCollectionScheduleRunSettingsLogLevel = "warn"
	SavedJobCollectionScheduleRunSettingsLogLevelInfo  SavedJobCollectionScheduleRunSettingsLogLevel = "info"
	SavedJobCollectionScheduleRunSettingsLogLevelDebug SavedJobCollectionScheduleRunSettingsLogLevel = "debug"
	SavedJobCollectionScheduleRunSettingsLogLevelSilly SavedJobCollectionScheduleRunSettingsLogLevel = "silly"
)

func (e SavedJobCollectionScheduleRunSettingsLogLevel) ToPointer() *SavedJobCollectionScheduleRunSettingsLogLevel {
	return &e
}

func (e *SavedJobCollectionScheduleRunSettingsLogLevel) UnmarshalJSON(data []byte) error {
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
		fallthrough
	case "silly":
		*e = SavedJobCollectionScheduleRunSettingsLogLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SavedJobCollectionScheduleRunSettingsLogLevel: %v", v)
	}
}

type SavedJobCollectionScheduleRunSettings struct {
	// Earliest time, for the given Range Timezone.
	Earliest *int64 `json:"earliest,omitempty"`
	// A filter for tokens in the provided collect path and/or the events being collected
	Expression *string `json:"expression,omitempty"`
	// Maximum time the job is allowed to run (e.g., 30, 45s or 15m). Units are seconds, if not specified. Enter 0 for unlimited time.
	JobTimeout *string `json:"jobTimeout,omitempty"`
	// Latest time, for the given Range Timezone.
	Latest *int64 `json:"latest,omitempty"`
	// Level at which to set task logging.
	LogLevel *SavedJobCollectionScheduleRunSettingsLogLevel `json:"logLevel,omitempty"`
	// Max number of times a task can be rescheduled.
	MaxTaskReschedule *int64 `json:"maxTaskReschedule,omitempty"`
	// Limits the bundle size for files above the Lower task bundle size. E.g., bundle five 2MB files into one 10MB task bundle. Files greater than this size will be assigned to individual tasks.
	MaxTaskSize *string `json:"maxTaskSize,omitempty"`
	// Limits the bundle size for small tasks. E.g., bundle five 200KB files into one 1M task.
	MinTaskSize *string `json:"minTaskSize,omitempty"`
	// Job run mode. Preview will either return up to N matching results, or will run until capture time T is reached. Discovery will gather the list of files to turn into streaming tasks, without running the data collection job. Full Run will run the collection job.
	Mode *string `json:"mode,omitempty"`
	// Reschedule tasks that failed with non-fatal errors.
	RescheduleDroppedTasks *bool `json:"rescheduleDroppedTasks,omitempty"`
	// Time range for scheduled job.
	TimeRangeType *string `json:"timeRangeType,omitempty"`
	// Timezone to use for Earliest and Latest times (defaults to UTC).
	TimestampTimezone *string `json:"timestampTimezone,omitempty"`
}

// SavedJobCollectionSchedule - Configuration for a scheduled job.
type SavedJobCollectionSchedule struct {
	// A cron schedule on which to run this job.
	CronSchedule *string `json:"cronSchedule,omitempty"`
	// Determines whether or not this schedule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The maximum number of instances that may be running of this scheduled job at any given time.
	MaxConcurrentRuns *int64                                 `json:"maxConcurrentRuns,omitempty"`
	ResumeMissed      interface{}                            `json:"resumeMissed,omitempty"`
	Run               *SavedJobCollectionScheduleRunSettings `json:"run,omitempty"`
	// Skippable jobs can be delayed, up to their next run time, if the system is hitting concurrency limits.
	Skippable *bool `json:"skippable,omitempty"`
}

// SavedJobCollectionJobType - Job type.
type SavedJobCollectionJobType string

const (
	SavedJobCollectionJobTypeCollection      SavedJobCollectionJobType = "collection"
	SavedJobCollectionJobTypeExecutor        SavedJobCollectionJobType = "executor"
	SavedJobCollectionJobTypeScheduledSearch SavedJobCollectionJobType = "scheduledSearch"
)

func (e SavedJobCollectionJobType) ToPointer() *SavedJobCollectionJobType {
	return &e
}

func (e *SavedJobCollectionJobType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "collection":
		fallthrough
	case "executor":
		fallthrough
	case "scheduledSearch":
		*e = SavedJobCollectionJobType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SavedJobCollectionJobType: %v", v)
	}
}

type SavedJobCollection struct {
	Collector SavedJobCollectionCollector `json:"collector"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Unique ID for this Job.
	ID    *string                  `json:"id,omitempty"`
	Input *SavedJobCollectionInput `json:"input,omitempty"`
	// List of fields to remove from Discover results. Wildcards (e.g.: aws*) are allowed. This is useful when discovery returns sensitive fields that should not be exposed in the Jobs user interface.
	RemoveFields []string `json:"removeFields,omitempty"`
	// Resumes the ad hoc job if a failure condition causes Stream to restart during job execution.
	ResumeOnBoot *bool `json:"resumeOnBoot,omitempty"`
	// Configuration for a scheduled job.
	Schedule *SavedJobCollectionSchedule `json:"schedule,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Time to keep the job's artifacts on disk after job completion. This also affects how long a job is listed in the Job Inspector.
	TTL *string `json:"ttl,omitempty"`
	// Job type.
	Type SavedJobCollectionJobType `json:"type"`
	// If enabled tasks are created and run by the same worker node.
	WorkerAffinity *bool `json:"workerAffinity,omitempty"`
}
