package constanta

const (
	EmailRegex    = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	UsernameRegex = `^[a-zA-Z0-9_]{3,20}$`
	PasswordRegex = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*()_+\[\]{}|;:'",.<>?/])[A-Za-z\d!@#$%^&*()_+\[\]{}|;:'",.<>?/]{8,20}$`
)
