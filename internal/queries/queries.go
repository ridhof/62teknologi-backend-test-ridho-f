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
