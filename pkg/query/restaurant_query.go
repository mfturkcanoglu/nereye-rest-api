package query

import (
	"fmt"
	"strings"
)

func Restaurant_All(customerId string) string {
	query := `SELECT
	r.photo_url,
	r.sign_name,
	r.about_us,
	r.extra_info,
	r.phone_number,
	r.workplace_phone_number,
	r.available_at_start,
	r.available_at_end,
	r.weekend_available_at_start,
	r.weekend_available_at_end,
	a.country,
	a.city,
	a.county,
	a.district,
	a.full_address,
	a.latitude,
	a.longitude
	FROM restaurant r
	INNER JOIN address a
	ON r.address_id = a.id`

	if strings.TrimSpace(customerId) != "" {
		return fmt.Sprintf("%s WHERE r.customer_id = '%s' order by r.updated_at desc", query, customerId)
	}

	return query + " order by r.updated_at desc"
}

func Restaurant_InsertQuery() string {
	return `INSERT INTO restaurant 
		(photo_url,
		sign_name,
		about_us,
		extra_info,
		phone_number,
		workplace_phone_number,
		available_at_start,
		available_at_end,
		weekend_available_at_start,
		weekend_available_at_end,
		address_id,
		customer_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);`
}
