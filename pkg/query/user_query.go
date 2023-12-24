package query

func UserInsertQuery() string {
	return "INSERT INTO users (username, email, phone_number, full_name, surname) VALUES ($1, $2, $3, $4, $5) RETURNING id"
}

func UserSelectQuery() string {
	return `
		SELECT
		u.username,
		u.phone_number,
		u.email,
		u.full_name,
		u.surname
		from users u
		order by u.updated_at desc
	`
}
