package models

// User ...
type User struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

// ...USERPREFIX
const (
	USERPREFIX = "user"
)
