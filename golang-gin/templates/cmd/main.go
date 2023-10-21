package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	// Version is the version of the application.
	VERSION string
	// COMMIT is the commit hash of the application build.
	COMMIT string
	// BUILDTIME is the time the application was built.
	BUILDTIME string
)

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version":   VERSION,
			"commit":    COMMIT,
			"buildtime": BUILDTIME,
		})
	})

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
