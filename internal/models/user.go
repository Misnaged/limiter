package models

type UserModel struct {
	Ip        string `json:"ip"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	AccountId int    `json:"account_id"`
}

func CreateNewUser(ip, nick, pass string, acc int) *UserModel {
	return &UserModel{
		Ip:        ip,
		Nickname:  nick,
		Password:  pass,
		AccountId: acc,
	}
}
