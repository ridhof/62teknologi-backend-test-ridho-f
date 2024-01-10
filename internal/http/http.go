package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRoot(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

func Create() (*gin.Engine) {
	router := gin.Default()
	router.GET("/", getRoot)

	return router
}
