package models

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid;not null;index"`
	ServiceName string    `gorm:"type:varchar(255);not null;index"`
	Price       int       `gorm:"not null"`
	StartDate   time.Time `gorm:"not null"`
	EndDate     *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GetAllSubscriptionsFilter struct {
	UserID      *uuid.UUID
	ServiceName *string
	StartDate   *time.Time
	EndDate     *time.Time
}

// TableName specifies the table name for the Subscription model
func (Subscription) TableName() string {
	return "subscriptions"
}
