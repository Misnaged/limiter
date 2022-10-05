package models

type UserModel struct {
	Ip        string `json:"ip"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	AccountId int    `json:"account_id"`
}
