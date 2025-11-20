package repository

import (
	"github.com/google/uuid"
	"subscribe-test-service/models"
	"time"
)

// GetAllSubscriptionsFilter определяет параметры фильтрации для получения списка подписок.
type GetAllSubscriptionsFilter struct {
	UserID      *uuid.UUID
	ServiceName *string
	// StartDate и EndDate используются для поиска подписок, которые были активны в указанном периоде.
	StartDate *time.Time
	EndDate   *time.Time
}

// Subscription — это интерфейс для работы с хранилищем подписок.
type Subscription interface {
	Create(sub models.Subscription) (uint, error)
	GetByID(id uint) (models.Subscription, error)
	GetAll(filter GetAllSubscriptionsFilter) ([]models.Subscription, error)
	Update(sub models.Subscription) error
	Delete(id uint) error
}
