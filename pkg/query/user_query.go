package query

import "fmt"

func UserInsertQueryWithReturn() string {
	return UserInsertQuery() + "RETURNING id;"
}

func UserInsertQuery() string {
	return "INSERT INTO users (username, email, phone_number, full_name, surname, password_hash) VALUES ($1, $2, $3, $4, $5, $6)"
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

func UserByUsernameQuery(username string) string {
	return fmt.Sprintf(`
		SELECT
		u.username,
		u.phone_number,
		u.email,
		u.full_name,
		u.surname
		from users u
		WHERE u.username = '%s'
	`, username)
}

func UserGetByUsernameAndPasswordQuery(username string, password string) string {
	return fmt.Sprintf(
		`
		SELECT
		u.username,
		u.phone_number,
		u.email,
		u.full_name,
		u.surname
		FROM users u
		WHERE u.username = '%s' AND u.password = '%s'
	`, username, password)
}
