package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var mr []model.MergeRequest
	jsonStr := string(body)
	err = json.Unmarshal([]byte(jsonStr), &mr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(mr)

	c.IndentedJSON(http.StatusOK, string(body))
}
