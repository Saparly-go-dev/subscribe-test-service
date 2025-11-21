package repository

import (
	"subscribe-test-service/models"

	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) Create(sub models.Subscription) (uint, error) {
	result := r.db.Create(&sub)
	if result.Error != nil {
		return 0, result.Error
	}
	return sub.ID, nil
}

// GetByID находит подписку по ID.
func (r *SubscriptionRepository) GetByID(id uint) (models.Subscription, error) {
	var sub models.Subscription
	result := r.db.First(&sub, id)
	return sub, result.Error
}

func (r *SubscriptionRepository) GetAll(filter models.GetAllSubscriptionsFilter) ([]models.Subscription, error) {
	var subs []models.Subscription
	query := r.db.Model(&models.Subscription{})

	if filter.UserID != nil {
		query = query.Where("user_id = ?", *filter.UserID)
	}

	if filter.ServiceName != nil && *filter.ServiceName != "" {
		query = query.Where("service_name = ?", *filter.ServiceName)
	}

	if filter.StartDate != nil && filter.EndDate != nil {
		query = query.Where("start_date <= ? AND (end_date IS NULL OR end_date >= ?)", *filter.EndDate, *filter.StartDate)
	}

	result := query.Find(&subs)
	return subs, result.Error
}

func (r *SubscriptionRepository) Update(sub models.Subscription) error {
	result := r.db.Save(&sub)
	return result.Error
}

// Delete удаляет подписку по ID.
func (r *SubscriptionRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Subscription{}, id)
	return result.Error
}
