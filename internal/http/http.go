package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"bizsearch/internal/database"
	"bizsearch/internal/models"
	"bizsearch/internal/queries"
)

func getRoot(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

func getBusinesses(c *gin.Context) {
	var businesses []models.Business
	db := database.GetInstance()

	var err error
	businesses, err = queries.GetBusinesses(db)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}

	c.IndentedJSON(http.StatusOK, businesses)
}

func Create() (*gin.Engine) {
	router := gin.Default()
	router.GET("/", getRoot)
	router.GET("/business", getBusinesses)

	return router
}
