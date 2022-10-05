package service

import (
	"bytes"
	"fmt"
	"net/http"

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

func NewService() (*Service, error) {

	srv := &Service{}
	srv.Handler = handlers.NewHandler()
	u := srv.CreateUser(srv.Log, srv.Pass)
	var err error
	srv.MarshalledBody, err = u.MarshalBody(u.GetUserModel())
	if err != nil {
		return nil, fmt.Errorf("cannot marshal user date: %w", err)
	}
	srv.Response, err = u.UnMarshallBody(srv.MarshalledBody)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	srv.Pass = u.CreatePassword(srv.Pass)

	return srv, nil
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
	fmt.Println(s.Log, s.Pass)
	return reg
}

func (s *Service) HandlerHttp() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Service) NewAccountId() int {
	id := AccountId + 1
	return id
}
func (s *Service) CreateUser(log, pass string) usr.IUser {

	accId := s.NewAccountId()
	user := usr.CreateNewUser(log, pass, accId)

	return user
}
