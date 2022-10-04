package service

import (
	"bytes"
	"fmt"
	"limiter/pkg/logger"
	"net/http"

	"limiter/internal/models"
	"limiter/internal/service/handlers"
	usr "limiter/internal/service/user"
)

const AccountId = 0

type Service struct {
	MarshalledBody []byte
	User           usr.IUser
	Handler        handlers.IHandler
	Response       interface{}
	Log, Pass      string
}

// SendMsg is a simple http Request sending from user
func (s *Service) SendMsg(msg []byte, url string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(msg))
	if err != nil {
		return nil, fmt.Errorf("failed to create message")
	}
	return req, nil
}

func (s *Service) NewCredentials() http.HandlerFunc {
	reg := s.Handler.Registration(s.Log, s.Pass)
	s.Pass = s.User.CreatePassword(s.Pass)
	return reg
}

func (s *Service) HandlerHttp() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		usr, err := s.CreateUser(s.Log, s.Pass, r.RemoteAddr)
		if err != nil {
			logger.Log().Errorf("cannot create user: %v", err)
			return
		}
		s.Response, err = usr.UnMarshallBody(s.MarshalledBody)
		if err != nil {
			logger.Log().Errorf("cannot unmarshal user date: %v", err)
			return
		}
	}
}

func (s *Service) NewAccountId() int {
	id := AccountId + 1
	return id
}
func (s *Service) CreateUser(login, pass, ip string) (usr.IUser, error) {

	var err error
	accId := s.NewAccountId()
	user := models.CreateNewUser(ip, login, pass, accId)
	s.MarshalledBody, err = s.User.MarshalBody(user)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal user date: %w", err)
	}

	return s.User, nil
}
