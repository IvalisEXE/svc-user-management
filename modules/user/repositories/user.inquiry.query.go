package repositories

const (
	findUserByID = `SELECT
	id,
	phone_no,
	fullname,
	password,
	created_at
	FROM users WHERE id = $1
	`

	findCountUserByID = `SELECT
	count(id)
	FROM users WHERE id = $1
	`

	findUserByPhoneNo = `SELECT
	id,
	phone_no,
	fullname,
	password,
	created_at
	FROM users WHERE phone_no = $1
	`

	findCountUserByPhoneNo = `SELECT
	count(id)
	FROM users WHERE phone_no = $1`
)
