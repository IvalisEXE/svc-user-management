package repositories

const (
	findUserLoginByUserID = `SELECT
	id,
	user_id,
	total_login,
	last_login,
	created_bu,
	created_at
	FROM user_login WHERE user_id = $1
	`

	findCountUserLoginByUserID = `SELECT
	count(1)
	FROM user_login WHERE user_id = $1
	`

	updateUserLogin = `UPDATE user_login SET  total_login = total_login+1, last_login = $1 WHERE user_id = $2`

	insertUserLogin = `INSERT INTO user_login (
		user_id,
		total_login,
		last_login,
		created_by,
		created_at
	) VALUES ($1,$2,$3,$4,$5)`

	upsertUserLogin = `
	INSERT INTO user_login (
		user_id,
		last_login,
		created_by,
		created_at
	) VALUES ($1,$2,$3,$4,$5,$6) 
	ON CONFLICT (user_id) 
	DO UPDATE SET 
		last_login = excluded.last_login
	`
)
