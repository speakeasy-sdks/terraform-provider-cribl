// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// SavedJobScheduledSearchScheduleRunSettingsLogLevel - Level at which to set task logging.
type SavedJobScheduledSearchScheduleRunSettingsLogLevel string

const (
	SavedJobScheduledSearchScheduleRunSettingsLogLevelError SavedJobScheduledSearchScheduleRunSettingsLogLevel = "error"
	SavedJobScheduledSearchScheduleRunSettingsLogLevelWarn  SavedJobScheduledSearchScheduleRunSettingsLogLevel = "warn"
	SavedJobScheduledSearchScheduleRunSettingsLogLevelInfo  SavedJobScheduledSearchScheduleRunSettingsLogLevel = "info"
	SavedJobScheduledSearchScheduleRunSettingsLogLevelDebug SavedJobScheduledSearchScheduleRunSettingsLogLevel = "debug"
	SavedJobScheduledSearchScheduleRunSettingsLogLevelSilly SavedJobScheduledSearchScheduleRunSettingsLogLevel = "silly"
)

func (e SavedJobScheduledSearchScheduleRunSettingsLogLevel) ToPointer() *SavedJobScheduledSearchScheduleRunSettingsLogLevel {
	return &e
}

func (e *SavedJobScheduledSearchScheduleRunSettingsLogLevel) UnmarshalJSON(data []byte) error {
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
		*e = SavedJobScheduledSearchScheduleRunSettingsLogLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SavedJobScheduledSearchScheduleRunSettingsLogLevel: %v", v)
	}
}

type SavedJobScheduledSearchScheduleRunSettings struct {
	// Earliest time, for the given Range Timezone.
	Earliest *int64 `json:"earliest,omitempty"`
	// A filter for tokens in the provided collect path and/or the events being collected
	Expression *string `json:"expression,omitempty"`
	// Maximum time the job is allowed to run (e.g., 30, 45s or 15m). Units are seconds, if not specified. Enter 0 for unlimited time.
	JobTimeout *string `json:"jobTimeout,omitempty"`
	// Latest time, for the given Range Timezone.
	Latest *int64 `json:"latest,omitempty"`
	// Level at which to set task logging.
	LogLevel *SavedJobScheduledSearchScheduleRunSettingsLogLevel `json:"logLevel,omitempty"`
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

// SavedJobScheduledSearchSchedule - Configuration for a scheduled job.
type SavedJobScheduledSearchSchedule struct {
	// A cron schedule on which to run this job.
	CronSchedule *string `json:"cronSchedule,omitempty"`
	// Determines whether or not this schedule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The maximum number of instances that may be running of this scheduled job at any given time.
	MaxConcurrentRuns *int64                                      `json:"maxConcurrentRuns,omitempty"`
	ResumeMissed      interface{}                                 `json:"resumeMissed,omitempty"`
	Run               *SavedJobScheduledSearchScheduleRunSettings `json:"run,omitempty"`
	// Skippable jobs can be delayed, up to their next run time, if the system is hitting concurrency limits.
	Skippable *bool `json:"skippable,omitempty"`
}

// SavedJobScheduledSearchJobType - Job type.
type SavedJobScheduledSearchJobType string

const (
	SavedJobScheduledSearchJobTypeCollection      SavedJobScheduledSearchJobType = "collection"
	SavedJobScheduledSearchJobTypeExecutor        SavedJobScheduledSearchJobType = "executor"
	SavedJobScheduledSearchJobTypeScheduledSearch SavedJobScheduledSearchJobType = "scheduledSearch"
)

func (e SavedJobScheduledSearchJobType) ToPointer() *SavedJobScheduledSearchJobType {
	return &e
}

func (e *SavedJobScheduledSearchJobType) UnmarshalJSON(data []byte) error {
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
		*e = SavedJobScheduledSearchJobType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SavedJobScheduledSearchJobType: %v", v)
	}
}

type SavedJobScheduledSearch struct {
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Unique ID for this Job.
	ID *string `json:"id,omitempty"`
	// List of fields to remove from Discover results. Wildcards (e.g.: aws*) are allowed. This is useful when discovery returns sensitive fields that should not be exposed in the Jobs user interface.
	RemoveFields []string `json:"removeFields,omitempty"`
	// Resumes the ad hoc job if a failure condition causes Stream to restart during job execution.
	ResumeOnBoot *bool `json:"resumeOnBoot,omitempty"`
	// Identifies which search query to run
	SavedQueryID string `json:"savedQueryId"`
	// Configuration for a scheduled job.
	Schedule *SavedJobScheduledSearchSchedule `json:"schedule,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Time to keep the job's artifacts on disk after job completion. This also affects how long a job is listed in the Job Inspector.
	TTL *string `json:"ttl,omitempty"`
	// Job type.
	Type SavedJobScheduledSearchJobType `json:"type"`
}
