package v0alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// The repository name and type are stored as labels
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Job struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JobSpec   `json:"spec,omitempty"`
	Status JobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type JobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Job `json:"items,omitempty"`
}

// HistoricJob is a history entry of Job. It is used to store Jobs that have been processed.
//
// The repository name and type are stored as labels.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type HistoricJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JobSpec   `json:"spec,omitempty"`
	Status JobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type HistoricJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []HistoricJob `json:"items,omitempty"`
}

// +enum
type JobAction string

const (
	// JobActionPull replicates the remote branch in the local copy of the repository.
	JobActionPull JobAction = "pull"

	// JobActionPush replicates the local copy of the repository in the remote branch.
	JobActionPush JobAction = "push"

	// JobActionPullRequest adds additional useful information to a PR, such as comments with preview links and rendered images.
	JobActionPullRequest JobAction = "pr"

	// JobActionMigrate acts like JobActionExport, then JobActionPull. It also tries to preserve the history.
	JobActionMigrate JobAction = "migrate"
)

// +enum
type JobState string

const (
	// Job has been submitted, but not processed yet
	JobStatePending JobState = "pending"

	// The job is running
	JobStateWorking JobState = "working"

	// Finished with success
	JobStateSuccess JobState = "success"

	// Finished with errors
	JobStateError JobState = "error"
)

func (j JobState) Finished() bool {
	return j == JobStateSuccess || j == JobStateError
}

type JobSpec struct {
	Action JobAction `json:"action,omitempty"`

	// The the repository reference (for now also in labels)
	// This value is required, but will be popuplated from the job making the request
	Repository string `json:"repository,omitempty"`

	// Pull request options
	PullRequest *PullRequestJobOptions `json:"pr,omitempty"`

	// Required when the action is `push`
	Push *ExportJobOptions `json:"push,omitempty"`

	// Required when the action is `pull`
	Pull *SyncJobOptions `json:"pull,omitempty"`

	// Required when the action is `migrate`
	Migrate *MigrateJobOptions `json:"migrate,omitempty"`
}

type PullRequestJobOptions struct {
	// The branch of commit hash
	Ref string `json:"ref,omitempty"`

	// Pull request number (when appropriate)
	PR   int    `json:"pr,omitempty"`
	Hash string `json:"hash,omitempty"` // used in PR code... not sure it is necessary

	// URL to the originator (eg, PR URL)
	URL string `json:"url,omitempty"`
}

type SyncJobOptions struct {
	// Incremental synchronization for versioned repositories
	Incremental bool `json:"incremental"`
}

type ExportJobOptions struct {
	// The source folder (or empty) to export
	Folder string `json:"folder,omitempty"`

	// Target branch for export (only git)
	Branch string `json:"branch,omitempty"`

	// Prefix in target file system
	Path string `json:"path,omitempty"`
}

type MigrateJobOptions struct {
	// Preserve history (if possible)
	History bool `json:"history,omitempty"`
}

// The job status
type JobStatus struct {
	State    JobState `json:"state,omitempty"`
	Started  int64    `json:"started,omitempty"`
	Finished int64    `json:"finished,omitempty"`
	Message  string   `json:"message,omitempty"`
	Errors   []string `json:"errors,omitempty"`

	// Optional value 0-100 that can be set while running
	Progress float64 `json:"progress,omitempty"`

	// Summary of processed actions
	Summary []*JobResourceSummary `json:"summary,omitempty"`
}

// Convert a JOB to a
func (in JobStatus) ToSyncStatus(jobId string) SyncStatus {
	return SyncStatus{
		JobID:    jobId,
		State:    in.State,
		Started:  in.Started,
		Finished: in.Finished,
		Message:  in.Errors,
	}
}

type JobResourceSummary struct {
	Group    string `json:"group,omitempty"`
	Resource string `json:"resource,omitempty"`
	Total    int64  `json:"total,omitempty"` // the count (if known)

	Create int64 `json:"create,omitempty"`
	Update int64 `json:"update,omitempty"`
	Delete int64 `json:"delete,omitempty"`
	Write  int64 `json:"write,omitempty"` // Create or update (export)
	Error  int64 `json:"error,omitempty"` // The error count

	// No action required (useful for sync)
	Noop int64 `json:"noop,omitempty"`

	// Report errors for this resource type
	// This may not be an exhaustive list and recommend looking at the logs for more info
	Errors []string `json:"errors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type WebhookResponse struct {
	metav1.TypeMeta `json:",inline"`

	// HTTP Status code
	// 200 implies that the payload was understood but nothing is required
	// 202 implies that an async job has been scheduled to handle the request
	Code int `json:"code,omitempty"`

	// Optional message
	Message string `json:"added,omitempty"`

	// Jobs to be processed
	// When the response is 202 (Accepted) the queued jobs will be returned
	Job *JobSpec `json:"job,omitempty"`
}
