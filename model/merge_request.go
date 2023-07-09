package model

import "time"

type MergeRequest struct {
	ID                          int         `json:"id"`
	Iid                         int         `json:"iid"`
	ProjectID                   int         `json:"project_id"`
	Title                       string      `json:"title"`
	Description                 string      `json:"description"`
	State                       string      `json:"state"`
	CreatedAt                   time.Time   `json:"created_at"`
	UpdatedAt                   time.Time   `json:"updated_at"`
	MergedBy                    string      `json:"merged_by"`
	MergeUser                   interface{} `json:"merge_user"`
	MergedAt                    time.Time   `json:"merged_at"`
	ClosedBy                    string      `json:"closed_by"`
	ClosedAt                    time.Time   `json:"closed_at"`
	TargetBranch                string      `json:"target_branch"`
	SourceBranch                string      `json:"source_branch"`
	UserNotesCount              int         `json:"user_notes_count"`
	Upvotes                     int         `json:"upvotes"`
	Downvotes                   int         `json:"downvotes"`
	Author                      User        `json:"author"`
	Assignees                   []User      `json:"assignees"`
	Assignee                    User        `json:"assignee"`
	Reviewers                   []User      `json:"reviewers"`
	SourceProjectID             int         `json:"source_project_id"`
	TargetProjectID             int         `json:"target_project_id"`
	Labels                      []string    `json:"labels"`
	Draft                       bool        `json:"draft"`
	WorkInProgress              bool        `json:"work_in_progress"`
	Milestone                   interface{} `json:"milestone"`
	MergeWhenPipelineSucceeds   bool        `json:"merge_when_pipeline_succeeds"`
	MergeStatus                 string      `json:"merge_status"`
	DetailedMergeStatus         string      `json:"detailed_merge_status"`
	Sha                         string      `json:"sha"`
	MergeCommitSha              string      `json:"merge_commit_sha"`
	SquashCommitSha             string      `json:"squash_commit_sha"`
	DiscussionLocked            string      `json:"discussion_locked"`
	ShouldRemoveSourceBranch    bool        `json:"should_remove_source_branch"`
	ForceRemoveSourceBranch     bool        `json:"force_remove_source_branch"`
	Reference                   string      `json:"reference"`
	References                  interface{} `json:"references"`
	WebUrl                      string      `json:"web_url"`
	TimeStats                   interface{} `json:"time_stats"`
	Squash                      bool        `json:"squash"`
	TaskCompletionStatus        interface{} `json:"task_completion_status"`
	HasConflicts                bool        `json:"has_conflicts"`
	BlockingDiscussionsResolved bool        `json:"blocking_discussions_resolved"`
	ApprovalsBeforeMerge        string      `json:"approvals_before_merge"`
}
