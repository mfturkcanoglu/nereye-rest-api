package query

func ProductGetQuery() string {
	return `
		SELECT
		p.id,
		p.product_name,
		p.photo_url,
		p.available_at_start,
		p.available_at_end,
		p.restaurant_id,
		p.category_id
		FROM product p
		order by p.updated_at desc
	`
}

func Product_InsertQuery() string {
	return `INSERT INTO product
		(product_name,
		photo_url,
		available_at_start,
		available_at_end,
		restaurant_id,
		category_id)
		VALUES ($1, $2, $3, $4, $5, $6);`
}
