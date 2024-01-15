package query

import "fmt"

func RestaurantPhotoSelectQuery() string {
	return `
		SELECT
		p.name,
		p.url,
		p.parent_id,
		p.photo_order
		FROM restaurant_photo
		order by p.photo_order desc
	`
}

func RestaurantPhotoSelectQueryByRestaurantId(restaurantId string) string {
	return fmt.Sprintf(`
		SELECT
		p.name,
		p.url,
		p.parent_id,
		p.photo_order
		FROM restaurant_photo
		ORDER BY p.photo_order DESC
		WHERE p.parent_id = '%s'`,
		restaurantId)
}
