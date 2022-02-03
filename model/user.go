package model

type User struct {
	UserId int64    `json:"user_id"`
	Name   string   `json:"name"`
	Pwd    string   `json:"pwd"`
	Roles  []string `json:"roles"`
}
