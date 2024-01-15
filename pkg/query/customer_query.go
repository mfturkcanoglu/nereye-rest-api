package query

func CustomerInsertQuery() string {
	return "INSERT INTO customer (company_name, customer_type, user_id) VALUES ($1, $2, $3)"
}

func CustomerUserSelectQuery() string {
	return `
		SELECT
		c.id,
		c.company_name,
		c.customer_type,
		u.username,
		u.phone_number,
		u.email,
		u.full_name,
		u.surname
		FROM customer c
		INNER JOIN users u
		ON c.user_id = u.id
		order by c.updated_at desc
	`
}


