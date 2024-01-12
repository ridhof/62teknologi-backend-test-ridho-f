package queries

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"

	"bizsearch/internal/models"
)

func GetBusinesses(db *sql.DB) ([]models.Business, error) {
	var businesses []models.Business

	rows, err := db.Query(
		"SELECT * FROM businesses ORDER BY created_date DESC",
	)
	if err != nil {
		return businesses, fmt.Errorf("Unable to get businesses: %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var business models.Business
		if err := rows.Scan(
			&business.ID,

			&business.Alias,
			&business.Name,
			&business.ImageUrl,
			pq.Array(&business.Transactions),
			&business.Latitude,
			&business.Longitude,
			&business.Price,
			&business.LocationAddressFirst,
			&business.LocationAddressSecond,
			&business.LocationAddressThird,
			&business.City,
			&business.ZipCode,
			&business.Country,
			&business.State,
			pq.Array(&business.DisplayAddress),
			&business.Phone,
			&business.DisplayPhone,

			&business.CreatedDate,
			&business.UpdatedDate,
		); err != nil {
			return businesses, fmt.Errorf("Unable to get businesses: %v", err)
		}
		businesses = append(businesses, business)
	}

	if err := rows.Err(); err != nil {
		return businesses, fmt.Errorf("Unable to get businesses: %v", err)
	}

	return businesses, nil
}

func CreateBusiness(db *sql.DB, newBusiness models.Business) (models.Business, error) {
	var business models.Business

	err := db.QueryRow(
		`INSERT INTO businesses
			(
				alias,
				name,
				image_url,
				transactions,
				latitude,
				longitude,
				price,
				location_address_1,
				location_address_2,
				location_address_3,
				city,
				zip_code,
				country,
				state,
				display_address,
				phone,
				display_phone
			)
			VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7,
				$8,
				$9,
				$10,
				$11,
				$12,
				$13,
				$14,
				$15,
				$16,
				$17
			) RETURNING *;`,
		newBusiness.Alias,
		newBusiness.Name,
		newBusiness.ImageUrl,
		pq.Array(newBusiness.Transactions),
		newBusiness.Latitude,
		newBusiness.Longitude,
		newBusiness.Price,
		newBusiness.LocationAddressFirst,
		newBusiness.LocationAddressSecond,
		newBusiness.LocationAddressThird,
		newBusiness.City,
		newBusiness.ZipCode,
		newBusiness.Country,
		newBusiness.State,
		pq.Array(newBusiness.DisplayAddress),
		newBusiness.Phone,
		newBusiness.DisplayPhone,
	).Scan(
		&business.ID,

		&business.Alias,
		&business.Name,
		&business.ImageUrl,
		pq.Array(&business.Transactions),
		&business.Latitude,
		&business.Longitude,
		&business.Price,
		&business.LocationAddressFirst,
		&business.LocationAddressSecond,
		&business.LocationAddressThird,
		&business.City,
		&business.ZipCode,
		&business.Country,
		&business.State,
		pq.Array(&business.DisplayAddress),
		&business.Phone,
		&business.DisplayPhone,

		&business.CreatedDate,
		&business.UpdatedDate,
	)
	if err != nil {
		return business, fmt.Errorf("Unable to create a new business: %v", err)
	}

	return business, nil
}
