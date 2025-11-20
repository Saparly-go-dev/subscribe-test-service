package service

import (
	"subscribe-test-service/pkg/repository"
)

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
