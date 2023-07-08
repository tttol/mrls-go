package gitlab

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tttol/mrls-go/model"
)

var user = model.User{
	ID:        1,
	Username:  "username",
	Name:      "name",
	State:     "state",
	AvatarUrl: "http://avatar.url",
	WebUrl:    "http://web.url",
}
var mr = model.MergeRequest{
	ID:                          1,
	Iid:                         1,
	ProjectID:                   1,
	Title:                       "title",
	Description:                 "description",
	State:                       "state",
	CreatedAt:                   time.Now(),
	UpdatedAt:                   time.Now(),
	MergedBy:                    "mergedBy",
	MergeUser:                   nil,
	MergedAt:                    time.Now(),
	ClosedBy:                    "closedBy",
	ClosedAt:                    time.Now(),
	TargetBranch:                "targetBranch",
	SourceBranch:                "sourceBranch",
	UserNotesCount:              1,
	Upvotes:                     1,
	Downvotes:                   1,
	Author:                      user,
	Assignees:                   []model.User{user},
	Assignee:                    user,
	Reviewers:                   []model.User{user},
	SourceProjectID:             1,
	TargetProjectID:             1,
	Labels:                      []string{"label1", "label2"},
	Draft:                       false,
	WorkInProgress:              false,
	Milestone:                   nil,
	MergeWhenPipelineSucceeds:   false,
	MergeStatus:                 "mergeStatus",
	DetailedMergeStatus:         "detailedMergeStatus",
	Sha:                         "sha",
	MergeCommitSha:              "mergeCommitSha",
	SquashCommitSha:             "squashCommitSha",
	DiscussionLocked:            "discussionLocked",
	ShouldRemoveSourceBranch:    false,
	ForceRemoveSourceBranch:     false,
	Reference:                   "reference",
	References:                  nil,
	WebUrl:                      "http://web.url",
	TimeStats:                   nil,
	Squash:                      false,
	TaskCompletionStatus:        nil,
	HasConflicts:                false,
	BlockingDiscussionsResolved: false,
	ApprovalsBeforeMerge:        "approvalsBeforeMerge",
}

func GetGitLabMr(c *gin.Context) {
	host := os.Getenv("GITLAB_HOST")
	project := os.Getenv("GITLAB_PROJECT_ID")
	token := os.Getenv("GITLAB_ACCESS_TOKEN")

	req, _ := http.NewRequest("GET", "https://"+host+"/api/v4/projects/"+project+"/merge_requests?state=opened", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
	c.IndentedJSON(http.StatusOK, string(body))
}
