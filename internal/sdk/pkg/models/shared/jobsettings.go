// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobSettings struct {
	ConcurrentJobLimit          int64  `json:"concurrentJobLimit"`
	ConcurrentScheduledJobLimit int64  `json:"concurrentScheduledJobLimit"`
	ConcurrentSystemJobLimit    int64  `json:"concurrentSystemJobLimit"`
	ConcurrentSystemTaskLimit   int64  `json:"concurrentSystemTaskLimit"`
	ConcurrentTaskLimit         int64  `json:"concurrentTaskLimit"`
	FinishedJobArtifactsLimit   int64  `json:"finishedJobArtifactsLimit"`
	FinishedTaskArtifactsLimit  int64  `json:"finishedTaskArtifactsLimit"`
	JobArtifactsReaperPeriod    string `json:"jobArtifactsReaperPeriod"`
	JobTimeout                  string `json:"jobTimeout"`
	MaxTaskPerc                 int64  `json:"maxTaskPerc"`
	SchedulingPolicy            string `json:"schedulingPolicy"`
	TaskHeartbeatPeriod         int64  `json:"taskHeartbeatPeriod"`
	TaskManifestFlushPeriodMs   int64  `json:"taskManifestFlushPeriodMs"`
	TaskManifestMaxBufferSize   int64  `json:"taskManifestMaxBufferSize"`
	TaskManifestReadBufferSize  string `json:"taskManifestReadBufferSize"`
	TaskPollTimeoutMs           int64  `json:"taskPollTimeoutMs"`
}
