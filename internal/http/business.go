package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"bizsearch/internal/database"
	"bizsearch/internal/models"
	"bizsearch/internal/queries"
)

func createBusiness(router *gin.Engine) {
	business := router.Group("/business")
	business.GET("", getBusinesses)
	business.POST("", postBusiness)
	business.PUT("", updateBusiness)
	// business.DELETE("/:id", deleteBusinessByID)
}

func getBusinesses(c *gin.Context) {
	var businesses []models.Business
	db := database.GetInstance()

	var err error
	businesses, err = queries.GetBusinesses(db)
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest, 
			gin.H{"message": fmt.Sprintf("%v", err)},
		)
		return
	}

	c.IndentedJSON(http.StatusOK, businesses)
}

func postBusiness(c *gin.Context) {
	var businessRequest models.Business
	db := database.GetInstance()

	if err := c.BindJSON(&businessRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v",err)})
		return
	}
	
	if err := businessRequest.Validate(); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": fmt.Sprintf("%v", err)},
		)
		return
	}

	business, err := queries.CreateBusiness(db, businessRequest)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}

	c.IndentedJSON(http.StatusCreated, business)
}

func updateBusiness(c *gin.Context) {
	var request models.Business
	db := database.GetInstance()

	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": fmt.Sprintf("%v", err)},
		)
		return
	}

	if err := request.Validate(); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": fmt.Sprintf("%v", err)},
		)
		return
	}

	updated, httpStatusCode, err := queries.UpdateBusiness(db, request)
	if err != nil {
		c.IndentedJSON(httpStatusCode, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}

	c.IndentedJSON(httpStatusCode, updated)
}
