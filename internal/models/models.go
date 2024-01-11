package models

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
