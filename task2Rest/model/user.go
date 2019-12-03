package model

import _ "encoding/json" // ...
// User ...
type User struct {
	Id  int `json:"id"`
	FistName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
}
