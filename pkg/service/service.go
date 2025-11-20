package service

import (
	"subscribe-test-service/models"
	"subscribe-test-service/pkg/repository"
)

type Subscription interface {
	Create(sub models.Subscription) (uint, error)
	GetByID(id uint) (models.Subscription, error)
	GetAll(filter models.GetAllSubscriptionsFilter) ([]models.Subscription, error)
	Update(sub models.Subscription) error
	Delete(id uint) error
}

type Service struct {
	Subscription
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Subscription: NewSubscriptionService(repos.Subscription),
	}
}
