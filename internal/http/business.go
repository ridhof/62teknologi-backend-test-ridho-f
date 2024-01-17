package http

import (
	"fmt"
	"net/http"
	"strconv"

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
	business.DELETE("/:id", deleteBusinessByID)
}

func getBusinesses(c *gin.Context) {
	var request models.GetBusinessRequest
	var response models.GetBusinessResponse
	db := database.GetInstance()

	var err error

	var latitude float64
	latitude, err = strconv.ParseFloat(c.DefaultQuery("latitude", "-6.2016627"), 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	request.Latitude = float32(latitude)
	response.Coordinates.Latitude = request.Latitude

	var longitude float64
	longitude, err = strconv.ParseFloat(c.DefaultQuery("longitude", "106.7881607"), 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	request.Longitude = float32(longitude)
	response.Coordinates.Longitude = request.Longitude

	var radius float64
	radius, err = strconv.ParseFloat(c.DefaultQuery("radius", "10"), 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	request.Radius = float32(radius)
	response.Coordinates.Radius = request.Radius

	request.Limit, err = strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	response.Pagination.Limit = request.Limit

	request.Offset, err = strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	response.Pagination.Offset = request.Offset

	response.Data, err = queries.GetBusinessesWithPagination(db, request)
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest, 
			gin.H{"message": fmt.Sprintf("%v", err)},
		)
		return
	}

	var totalRow, totalPage int
	totalRow, totalPage, err = queries.GetBusinessTotalPages(db, request)
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": fmt.Sprintf("%v", err)},
		)
	}
	response.Pagination.Total = totalRow
	response.Pagination.TotalPage = totalPage

	c.IndentedJSON(http.StatusOK, response)
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

	c.IndentedJSON(http.StatusCreated, gin.H{"data": business})
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

	c.IndentedJSON(httpStatusCode, gin.H{"data": updated})
}

func deleteBusinessByID(c *gin.Context) {
	db := database.GetInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": fmt.Sprintf("%v", err)},
		)
		return
	}

	business, httpStatusCode, err := queries.DeleteBusinessByID(db, int64(id))
	if err != nil {
		c.IndentedJSON(
			httpStatusCode,
			gin.H{"message": fmt.Sprintf("%v", err)},
		)
		return
	}

	c.IndentedJSON(httpStatusCode, gin.H{"data": business})
}
