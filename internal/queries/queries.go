package queries

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/lib/pq"

	"bizsearch/internal/models"
)

func GetBusinesses(db *sql.DB, request models.GetBusinessRequest) ([]models.Business, error) {
	var businesses []models.Business

	// limit, offset

	rows, err := db.Query(
		`SELECT * FROM (
			SELECT b.*, (
				6371 * acos(
					cos( radians($1) ) * cos( radians(latitude) )
					* cos(
						radians(longitude) - radians($2)
					) + sin( radians($1) ) * sin( radians(latitude) )
				)
			) as distance
			FROM businesses as b
			WHERE EXISTS (SELECT 1 FROM reviews as r WHERE r.business_id = b.id)
		)
		WHERE distance < $3
		ORDER BY distance ASC
		OFFSET $4 LIMIT $5`,
		request.Latitude,
		request.Longitude,
		request.Radius,
		request.Offset,
		request.Limit,
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

			&business.Distance,
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
			) RETURNING *`,
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
		strings.ToLower(newBusiness.Country),
		strings.ToLower(newBusiness.State),
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

func UpdateBusiness(db *sql.DB, business models.Business) (models.Business, int, error) {
	var updated models.Business

	err := db.QueryRow(
		`UPDATE businesses
			SET alias = $1,
			name = $2,
			image_url = $3,
			transactions = $4,
			latitude = $5,
			longitude = $6,
			price = $7,
			location_address_1 = $8,
			location_address_2 = $9,
			location_address_3 = $10,
			city = $11,
			zip_code = $12,
			country = $13,
			state = $14,
			display_address = $15,
			phone = $16,
			display_phone = $17,
			updated_date = now()
		WHERE id = $18 RETURNING *`,
		business.Alias,
		business.Name,
		business.ImageUrl,
		pq.Array(business.Transactions),
		business.Latitude,
		business.Longitude,
		business.Price,
		business.LocationAddressFirst,
		business.LocationAddressSecond,
		business.LocationAddressThird,
		business.City,
		business.ZipCode,
		strings.ToLower(business.Country),
		strings.ToLower(business.State),
		pq.Array(business.DisplayAddress),
		business.Phone,
		business.DisplayPhone,
		business.ID,
	).Scan(
		&updated.ID,

		&updated.Alias,
		&updated.Name,
		&updated.ImageUrl,
		pq.Array(&updated.Transactions),
		&updated.Latitude,
		&updated.Longitude,
		&updated.Price,
		&updated.LocationAddressFirst,
		&updated.LocationAddressSecond,
		&updated.LocationAddressThird,
		&updated.City,
		&updated.ZipCode,
		&updated.Country,
		&updated.State,
		pq.Array(&updated.DisplayAddress),
		&updated.Phone,
		&updated.DisplayPhone,

		&updated.CreatedDate,
		&updated.UpdatedDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return business, 404, fmt.Errorf(
				"Could not find business with ID %v: %v",
				business.ID,
				err,
			)
		}
		return business, 401, fmt.Errorf(
			"Could not update business with ID %v: %v",
			business.ID,
			err,
		)
	}

	return updated, 200, nil
}

func DeleteBusinessByID(db *sql.DB, id int64) (models.Business, int, error) {
	var business models.Business

	result := db.QueryRow(
		"DELETE FROM businesses WHERE id = $1 RETURNING *",
		id,
	)
	if err := result.Scan(
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
		if err == sql.ErrNoRows {
			return business, 404, fmt.Errorf(
				"Could not find business with ID %v: %v",
				id,
				err,
			)
		}
		return business, 400, fmt.Errorf(
			"Could not delete business with ID %v: %v",
			id,
			err,
		)
	}

	return business, 200, nil
}
