package http

import (
	"fmt"
	"net/http"
	"strings"

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

	var priceType = map[string]bool {
		"$": true,
		"$$": true,
		"$$$": true,
		"$$$$": true,
		"$$$$$": true,
	}
	_, priceTypeOk := priceType[businessRequest.Price]
	if priceTypeOk != true {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Unable to create business: %v price type is not valid", businessRequest.Price)})
		return
	}

	var transactionType = map[string]bool {
		"pickup": true,
		"delivery": true,
	}
	var invalidTrxType = []string{}
	for _, trxType := range businessRequest.Transactions {
		_, transactionTypeOk := transactionType[trxType]
		if transactionTypeOk != true {
			invalidTrxType = append(invalidTrxType, trxType)
		}
	}
	if len(invalidTrxType) > 0 {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"message": fmt.Sprintf(
					"Unable to create business: %v transaction type is not valid",
					strings.Join(invalidTrxType, ", "),
				),
			},
		)
		return
	}

	var invalidStringCols = []string{}
	if businessRequest.Alias == "" {
		invalidStringCols = append(invalidStringCols, "alias")
	}
	if businessRequest.Name == "" {
		invalidStringCols = append(invalidStringCols, "name")
	}
	if businessRequest.ImageUrl == "" {
		invalidStringCols = append(invalidStringCols, "image_url")
	}
	if businessRequest.LocationAddressFirst == "" {
		invalidStringCols = append(invalidStringCols, "location_address_1")
	}
	if businessRequest.City == "" {
		invalidStringCols = append(invalidStringCols, "city")
	}
	if businessRequest.ZipCode == "" {
		invalidStringCols = append(invalidStringCols, "zip_code")
	}
	if businessRequest.Country == "" {
		invalidStringCols = append(invalidStringCols, "country")
	}
	if businessRequest.State == "" {
		invalidStringCols = append(invalidStringCols, "state")
	}
	if businessRequest.Phone == "" {
		invalidStringCols = append(invalidStringCols, "phone")
	}
	if businessRequest.DisplayPhone == "" {
		invalidStringCols = append(invalidStringCols, "display_phone")
	}
	if len(invalidStringCols) > 0 {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"message": fmt.Sprintf(
					"Unable to create business: Value of these columns should not be empty. Columns with empty value: %v", 
					strings.Join(invalidStringCols, ", "),
				),
			},
		)
		return
	}

	invalidStringCols = []string{}
	if len(businessRequest.State) > 3 {
		invalidStringCols = append(invalidStringCols, "state")
	}
	if len(businessRequest.Country) > 3 {
		invalidStringCols = append(invalidStringCols, "country")
	}
	if len(invalidStringCols) > 0 {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"message": fmt.Sprintf(
					"Unable to create business: Value of these columns should not be more than 3 characters. Columns that has value with more than 3 character : %v",
					strings.Join(invalidStringCols, ", "),
				),
			},
		)
		return
	}

	businessRequest.State = strings.ToLower(businessRequest.State)
	businessRequest.Country = strings.ToLower(businessRequest.Country)

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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}

	var priceType = map[string]bool {
		"$": true,
		"$$": true,
		"$$$": true,
		"$$$$": true,
		"$$$$$": true,
	}
	_, priceTypeOk := priceType[request.Price]
	if priceTypeOk != true {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Unable to update business: %q price type is not valid", request.Price)})
		return
	}

	var transactionType = map[string]bool {
		"pickup": true,
		"delivery": true,
	}
	var invalidTrxType = []string{}
	if len(request.Transactions) == 0 {
		invalidTrxType = append(invalidTrxType, "empty")
	} else {
		for _, trxType := range request.Transactions {
			_, transactionTypeOk := transactionType[trxType]
			if transactionTypeOk != true {
				invalidTrxType = append(invalidTrxType, trxType)
			}
		}
	}
	if len(invalidTrxType) > 0 {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"message": fmt.Sprintf(
					"Unable to update business: %v transaction type is not valid",
					strings.Join(invalidTrxType, ", "),
				),
			},
		)
		return
	}

	var validDisplayAddress bool = true
	if len(request.DisplayAddress) == 0 {
		validDisplayAddress = false
	} else {
		for _, displayAddress := range request.DisplayAddress {
			if displayAddress == "" {
				validDisplayAddress = false
			}
		}
	}
	if validDisplayAddress == false {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Unable to update business: display address is not valid",
			},
		)
		return
	}

	var invalidStringCols = []string{}
	if request.Alias == "" {
		invalidStringCols = append(invalidStringCols, "alias")
	}
	if request.Name == "" {
		invalidStringCols = append(invalidStringCols, "name")
	}
	if request.ImageUrl == "" {
		invalidStringCols = append(invalidStringCols, "image_url")
	}
	if request.LocationAddressFirst == "" {
		invalidStringCols = append(invalidStringCols, "location_address_1")
	}
	if request.City == "" {
		invalidStringCols = append(invalidStringCols, "city")
	}
	if request.ZipCode == "" {
		invalidStringCols = append(invalidStringCols, "zip_code")
	}
	if request.Country == "" {
		invalidStringCols = append(invalidStringCols, "country")
	}
	if request.State == "" {
		invalidStringCols = append(invalidStringCols, "state")
	}
	if request.Phone == "" {
		invalidStringCols = append(invalidStringCols, "phone")
	}
	if request.DisplayPhone == "" {
		invalidStringCols = append(invalidStringCols, "display_phone")
	}
	if request.Latitude == 0 {
		invalidStringCols = append(invalidStringCols, "latitude")
	}
	if request.Longitude == 0 {
		invalidStringCols = append(invalidStringCols, "longitude")
	}
	if len(invalidStringCols) > 0 {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"message": fmt.Sprintf(
					"Unable to update business: Value of these columns should not be empty. Columns with empty value: %v", 
					strings.Join(invalidStringCols, ", "),
				),
			},
		)
		return
	}

	invalidStringCols = []string{}
	if len(request.State) > 3 {
		invalidStringCols = append(invalidStringCols, "state")
	}
	if len(request.Country) > 3 {
		invalidStringCols = append(invalidStringCols, "country")
	}
	if len(invalidStringCols) > 0 {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"message": fmt.Sprintf(
					"Unable to update business: Value of these columns should not be more than 3 characters. Columns that has value with more than 3 character : %v",
					strings.Join(invalidStringCols, ", "),
				),
			},
		)
		return
	}

	request.State = strings.ToLower(request.State)
	request.Country = strings.ToLower(request.Country)

	updated, httpStatusCode, err := queries.UpdateBusiness(db, request)
	if err != nil {
		c.IndentedJSON(httpStatusCode, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}

	c.IndentedJSON(httpStatusCode, updated)
}
