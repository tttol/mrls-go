package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tttol/mrls-go/model"
)

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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var mr []model.Mr
	jsonStr := string(body)
	err = json.Unmarshal([]byte(jsonStr), &mr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(mr)

	mrDetailsJson, err := json.Marshal(generateMrDetailsMap(mr))
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	// fmt.Println(string(mrDetailsJson))

	c.IndentedJSON(http.StatusOK, string(mrDetailsJson))
}

func generateMrDetailsMap(mr []model.Mr) map[int][]model.MrDetail {
	mrDetailsMap := map[int][]model.MrDetail{}
	for _, e := range mr {
		asigneeId := e.Assignee.ID
		mrDetailList := mrDetailsMap[asigneeId]
		mrDetailsMap[asigneeId] = append(mrDetailList, generateMrDetail(e))
	}

	return mrDetailsMap
}

func generateMrDetail(mr model.Mr) model.MrDetail {
	return model.MrDetail{
		Title:     mr.Title,
		WebUrl:    mr.WebUrl,
		Author:    mr.Author,
		Upvotes:   mr.Upvotes,
		Labels:    mr.Labels,
		CreatedAt: mr.CreatedAt,
		UpdatedAt: mr.UpdatedAt,
	}
}
