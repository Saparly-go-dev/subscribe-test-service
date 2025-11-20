package service

import ()

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
