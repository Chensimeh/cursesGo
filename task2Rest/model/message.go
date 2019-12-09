package model

import _ "encoding/json" // ...

type Message struct {
	Id int `json:"id"`
	Message string `json:"message"`
	User_id int `json:"user_id"`
}
