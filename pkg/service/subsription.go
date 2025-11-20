package service

import (
	"subscribe-test-service/models"
	"subscribe-test-service/pkg/repository"
)

type SubscriptionService struct {
	repo repository.Subscription
}

func NewSubscriptionService(repo repository.Subscription) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) Create(sub models.Subscription) (uint, error) {
	return s.repo.Create(sub)
}

func (s *SubscriptionService) GetByID(id uint) (models.Subscription, error) {
	return s.repo.GetByID(id)
}

func (s *SubscriptionService) GetAll(filter models.GetAllSubscriptionsFilter) ([]models.Subscription, error) {
	return s.repo.GetAll(filter)
}

func (s *SubscriptionService) Update(sub models.Subscription) error {
	return s.repo.Update(sub)
}

func (s *SubscriptionService) Delete(id uint) error {
	return s.repo.Delete(id)
}
