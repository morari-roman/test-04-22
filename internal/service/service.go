package service

import (
	"fmt"

	"github.com/morari-roman/test-04-22/asana/internal/asana"
	"github.com/morari-roman/test-04-22/asana/internal/request"
)

const (
	users    = "users"
	projects = "projects"
)

type Request interface {
	MakeRequest() (*asana.AsanaData, error)
}

type Service struct {
	user    *request.Request
	project *request.Request
}

func New() *Service {
	return &Service{
		user:    request.New(users),
		project: request.New(projects),
	}
}

func (s *Service) GetData() error {
	u, err := s.user.MakeRequest()
	if err != nil {
		return err
	}

	fmt.Println(u)

	p, err := s.project.MakeRequest()
	if err != nil {
		return err
	}

	fmt.Println(p)

	return nil
}
