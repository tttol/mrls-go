package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tttol/mrls-go/gitlab"
)

func main() {
	router := gin.Default()
	router.GET("/v1/mr", gitlab.GetGitLabMr)

	router.Run("localhost:8080")
}
