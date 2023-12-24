package query

func CreateAddressQueryReturnId() string {
	return "INSERT INTO address (country, city, county, district, full_address, latitude, longitude) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;"
}
