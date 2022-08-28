package models

type User struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
}
