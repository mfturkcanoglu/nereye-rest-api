package query

func UserInsertQueryWithReturn() string {
	return UserInsertQuery() + "RETURNING id;"
}

func UserInsertQuery() string {
	return "INSERT INTO users (username, email, phone_number, full_name, surname) VALUES ($1, $2, $3, $4, $5)"
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
