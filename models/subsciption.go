package model

import (
	"github.com/google/uuid"
	"time"
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

// TableName specifies the table name for the Subscription model
func (Subscription) TableName() string {
	return "subscriptions"
}
