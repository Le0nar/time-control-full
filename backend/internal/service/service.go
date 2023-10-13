package service

import "github.com/le0nar/time-control/internal/repository"

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
