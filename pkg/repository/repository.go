package repository

import (
	"gorm.io/gorm"

	"subscribe-test-service/models"
)

// Subscription — это интерфейс для работы с хранилищем подписок.
type Subscription interface {
	Create(sub models.Subscription) (uint, error)
	GetByID(id uint) (models.Subscription, error)
	GetAll(filter models.GetAllSubscriptionsFilter) ([]models.Subscription, error)
	Update(sub models.Subscription) error
	Delete(id uint) error
}
type Repository struct {
	Subscription
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Subscription: NewSubscriptionRepository(db),
	}
}
