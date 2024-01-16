package query

func CategoryGetQuery() string {
	return `
		SELECT
		c.id,
		c.category,
		c.photo_url
		FROM category c
		order by c.updated_at desc
	`
}

func Category_InsertQuery() string {
	return `INSERT INTO category
		(category,
		photo_url)
		VALUES ($1, $2);`
}
