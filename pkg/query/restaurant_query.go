package query

import (
	"fmt"
	"strings"
)

func Restaurant_All(customerId string) string {
	query := `SELECT
	r.photo_url,
	r.sign_name,
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
	return "INSERT INTO restaurant (photo_url, sign_name, address_id, customer_id) VALUES ($1, $2, $3, $4);"
}
