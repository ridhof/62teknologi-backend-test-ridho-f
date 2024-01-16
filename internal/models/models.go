package models

import (
	"fmt"
	"strings"
)

type User struct {
	ID int64 `json:"id"`

	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
	
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}

type Business struct {
	ID int64 `json:"id"`

	Alias string `json:"alias"`
	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
	Transactions []string `json:"transactions"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Price string `json:"price"`
	LocationAddressFirst string `json:"location_address_1"`
	LocationAddressSecond *string `json:"location_address_2"`
	LocationAddressThird *string `json:"location_address_3"`
	City string `json:"city"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
	State string `json:"state"`
	DisplayAddress []string `json:"display_address"`
	Phone string `json:"phone"`
	DisplayPhone string `json:"display_phone"`

	Distance float32 `json:"distance"`
	
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}

type Category struct {
	ID int64 `json:"id"`

	Title string `json:"title"`
	Alias string `json:"alias"`
	
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}

type BusinessCategory struct {
	ID int64 `json:"id"`

	BusinessId int64 `json:"business_id"`
	CategoryId int64 `json:"category_id"`
	
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}

type Review struct {
	ID int64 `json:"id"`

	Text string `json:"text"`
	Rating int `json:"rating"`
	UserId int64 `json:"user_id"`
	BusinessId int64 `json:"business_id"`
	
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}

type GetBusinessRequest struct {
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Radius float32 `json:"radius"`
	Limit int `json:"limit"`
	Offset int `json:"offset"`
}

func (business Business) Validate() (error) {
	var arrErr = []string{}

	if business.Price == "" {
		arrErr = append(
			arrErr,
			"empty price type is not valid",
		)
	} else {
		var priceType = map[string]bool {
			"$": true,
			"$$": true,
			"$$$": true,
			"$$$$": true,
			"$$$$$": true,
		}
		_, priceTypeOk := priceType[business.Price]
		if priceTypeOk != true {
			arrErr = append(
				arrErr,
				fmt.Sprintf(
					"%v price type is not valid",
					business.Price,
				),
			)
		}
	}

	var transactionType = map[string]bool {
		"pickup": true,
		"delivery": true,
	}
	var invalidTrxType = []string{}
	if len(business.Transactions) == 0 {
		invalidTrxType = append(invalidTrxType, "empty")
	} else {
		for _, trxType := range business.Transactions {
			_, transactionTypeOk := transactionType[trxType]
			if transactionTypeOk != true {
				invalidTrxType = append(invalidTrxType, trxType)
			}
		}
	}
	if len(invalidTrxType) > 0 {
		arrErr = append(
			arrErr,
			fmt.Sprintf(
				"%v transaction type is not valid",
				strings.Join(invalidTrxType, ", "),
			),
		)
	}

	var validDisplayAddress bool = true
	if len(business.DisplayAddress) == 0 {
		validDisplayAddress = false
	} else {
		for _, displayAddress := range business.DisplayAddress {
			if displayAddress == "" {
				validDisplayAddress = false
			}
		}
	}
	if validDisplayAddress != true {
		arrErr = append(
			arrErr,
			"display address is not valid",
		)
	}

	var invalidStringCols = []string{}
	if business.Alias == "" {
		invalidStringCols = append(invalidStringCols, "alias")
	}
	if business.Name == "" {
		invalidStringCols = append(invalidStringCols, "name")
	}
	if business.ImageUrl == "" {
		invalidStringCols = append(invalidStringCols, "image_url")
	}
	if business.LocationAddressFirst == "" {
		invalidStringCols = append(invalidStringCols, "location_address_1")
	}
	if business.City == "" {
		invalidStringCols = append(invalidStringCols, "city")
	}
	if business.ZipCode == "" {
		invalidStringCols = append(invalidStringCols, "zip_code")
	}
	if business.Country == "" {
		invalidStringCols = append(invalidStringCols, "country")
	}
	if business.State == "" {
		invalidStringCols = append(invalidStringCols, "state")
	}
	if business.Phone == "" {
		invalidStringCols = append(invalidStringCols, "phone")
	}
	if business.DisplayPhone == "" {
		invalidStringCols = append(invalidStringCols, "display_phone")
	}
	if business.Latitude == 0 {
		invalidStringCols = append(invalidStringCols, "latitude")
	}
	if business.Longitude == 0 {
		invalidStringCols = append(invalidStringCols, "longitude")
	}
	if len(invalidStringCols) > 0 {
		arrErr = append(
			arrErr,
			fmt.Sprintf(
				"columns should not have an empty value: %v",
				strings.Join(invalidStringCols, ", "),
			),
		)
	}

	invalidStringCols = []string{}
	if len(business.State) > 3 {
		invalidStringCols = append(invalidStringCols, "state")
	}
	if len(business.Country) > 3 {
		invalidStringCols = append(invalidStringCols, "country")
	}
	if len(invalidStringCols) > 0 {
		arrErr = append(
			arrErr,
			fmt.Sprintf(
				"columns should not have a value with more than 3 character : %v",
				strings.Join(invalidStringCols, ", "),
			),
		)
	}

	if len(arrErr) > 0 {
		return fmt.Errorf(strings.Join(arrErr, ". "))
	}
	return nil
}
