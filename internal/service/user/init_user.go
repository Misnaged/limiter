package user

import (
	"encoding/base64"
	"fmt"
	"github.com/goccy/go-json"
	"limiter/internal/models"
)

type IUser interface {
	CreatePassword(pass string) string
	MarshalBody(u *models.UserModel) ([]byte, error)
	UnMarshallBody(body []byte) (interface{}, error)
	GetUserModel() *models.UserModel
}

func CreateNewUser(nick, pass string, acc int) *User {
	userModel := &models.UserModel{
		Nickname:  nick,
		Password:  pass,
		AccountId: acc,
	}
	user := &User{ModelUser: userModel}
	return user
}

type User struct {
	ModelUser *models.UserModel
}

func (usr *User) GetUserModel() *models.UserModel {
	return usr.ModelUser
}
func (usr *User) CreatePassword(pass string) string {
	enc := base64.StdEncoding.EncodeToString([]byte(pass))
	return enc
}

func (usr *User) CreateNickName(name string) string {
	return name
}
func (usr *User) MarshalBody(u *models.UserModel) ([]byte, error) {
	body, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("marshalling failed: %w", err)
	}
	return body, nil
}

func (usr *User) UnMarshallBody(body []byte) (interface{}, error) {
	err := json.Unmarshal(body, &usr.ModelUser)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling failed: %w", err)
	}
	return &usr.ModelUser, nil
}
