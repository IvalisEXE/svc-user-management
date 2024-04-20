package repositories

const (
	updateUserProfileByID  = `UPDATE users SET phone_no = $1, fullname = $2 WHERE id = $3 `
	updateUserFullnameByID = `UPDATE users SET fullname = $1 WHERE id = $2`
)
