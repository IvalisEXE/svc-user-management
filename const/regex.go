package constant

const (
	REGEX_USER_PHONE_NO_VALIDATION        = `^\+62\d{8,15}$`
	REGEX_USER_PASSWORD_CAPITAL           = `[A-Z]`
	REGEX_USER_PASSWORD_ALPHA_NUMERIC     = `[a-zA-Z0-9]`
	REGEX_USER_PASSWORD_NON_ALPHA_NUMERIC = `[^a-zA-Z\d\s:]`
)
