package repositories

const (
	createUser = "INSERT INTO users (id,phone_no,fullname,password,created_at,created_by) VALUES ($1,$2,$3,$4,$5,$6)"
)
