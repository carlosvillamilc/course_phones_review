package models

// Smartphone model structure for smartphones
type User struct {
	Id       int64
	Username string
	Password string
}

// CreateUserCMD
type CreateUserCMD struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
