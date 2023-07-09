package gitlab

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tttol/mrls-go/model"
)

func TestGenerateMrDetailsMap(t *testing.T) {
	// 1. Prepare the data
	author := model.User{
		ID:        1,
		Username:  "tarosan",
		Name:      "太郎",
		State:     "opened",
		AvatarUrl: "http://avatar.example.com",
		WebUrl:    "http://weburl.example.com",
	}
	labels := []string{"label1", "label2"}
	createdAt := time.Now().Add(-time.Hour)
	updatedAt := time.Now()
	assignee := model.User{ID: 1} // Assuming User has an ID field
	mr1 := model.Mr{
		Title:     "Test MR 1",
		WebUrl:    "https://example.com/1",
		Author:    author,
		Upvotes:   5,
		Labels:    labels,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Assignee:  assignee,
	}
	mr2 := model.Mr{
		Title:     "Test MR 2",
		WebUrl:    "https://example.com/2",
		Author:    author,
		Upvotes:   10,
		Labels:    labels,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Assignee:  assignee,
	}
	mr := []model.Mr{mr1, mr2}

	// 2. Call the function
	actual := generateMrDetailsMap(mr)

	// 3. Check the result
	mrDetail1 := generateMrDetail(mr1)
	mrDetail2 := generateMrDetail(mr2)
	expected := map[int][]model.MrDetail{
		1: {mrDetail1, mrDetail2}, // Assuming the ID of the assignee is 1
	}
	assert.Equal(t, expected, actual, "The MrDetail map should match the expected result.")
}

func TestGenerateMrDetail(t *testing.T) {
	// 1. Prepare the data
	author := model.User{
		ID:        1,
		Username:  "tarosan",
		Name:      "太郎",
		State:     "opened",
		AvatarUrl: "http://avatar.example.com",
		WebUrl:    "http://weburl.example.com",
	}
	labels := []string{"label1", "label2"}
	createdAt := time.Now().Add(-time.Hour)
	updatedAt := time.Now()
	mr := model.Mr{
		Title:     "Test MR",
		WebUrl:    "https://example.com",
		Author:    author,
		Upvotes:   5,
		Labels:    labels,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	// 2. Call the function
	actual := generateMrDetail(mr)

	// 3. Check the result
	expected := model.MrDetail{
		Title:     "Test MR",
		WebUrl:    "https://example.com",
		Author:    author,
		Upvotes:   5,
		Labels:    labels,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	assert.Equal(t, expected, actual, "The actual MrDetail object should match the expected result.")
}
