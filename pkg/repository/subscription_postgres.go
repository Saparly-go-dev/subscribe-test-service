package repository

import (
	"gorm.io/gorm"
	"subscribe-test-service/models"
)

// SubscriptionRepository реализует интерфейс Subscription для работы с PostgreSQL.
type SubscriptionRepository struct {
	db *gorm.DB
}

// NewSubscriptionRepository создает новый экземпляр SubscriptionRepository.
func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

// Create создает новую подписку в базе данных.
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

func (r *SubscriptionRepository) GetAll(filter GetAllSubscriptionsFilter) ([]models.Subscription, error) {
	var subs []models.Subscription
	query := r.db.Model(&models.Subscription{})

	if filter.UserID != nil {
		query = query.Where("user_id = ?", *filter.UserID)
	}

	if filter.ServiceName != nil && *filter.ServiceName != "" {
		query = query.Where("service_name = ?", *filter.ServiceName)
	}

	// Логика для фильтрации по дате:
	// Нам нужны все подписки, которые пересекаются с заданным диапазоном [filter.StartDate, filter.EndDate].
	// Условие: sub.StartDate <= filter.EndDate AND (sub.EndDate IS NULL OR sub.EndDate >= filter.StartDate)
	if filter.StartDate != nil && filter.EndDate != nil {
		query = query.Where("start_date <= ? AND (end_date IS NULL OR end_date >= ?)", *filter.EndDate, *filter.StartDate)
	}

	result := query.Find(&subs)
	return subs, result.Error
}

// Update обновляет информацию о подписке.
// GORM's Save() обновляет все поля, если запись существует, или создает новую.
func (r *SubscriptionRepository) Update(sub models.Subscription) error {
	result := r.db.Save(&sub)
	return result.Error
}

// Delete удаляет подписку по ID.
func (r *SubscriptionRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Subscription{}, id)
	return result.Error
}
