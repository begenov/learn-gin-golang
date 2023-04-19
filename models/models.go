package models

type Credentials struct {
	ID       int
	Username string `json:"username"`
	Password string `json:"password"`
}
