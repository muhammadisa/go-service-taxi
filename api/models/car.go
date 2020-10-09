package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Car struct
type Car struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	UserID    string    `gorm:"column:user_id" json:"user_id,omitempty" validate:"required"`
	BrandID   string    `gorm:"column:brand_id" json:"brand_id,omitempty" validate:"required"`
	CarName   string    `gorm:"size:255;not null" json:"employee_name" validate:"required"`
	Condition string    `gorm:"size:5;not null" json:"condition" validate:"required"`
	Unit      int64     `gorm:"not null" json:"unit"`
	Price     int64     `gorm:"not null" json:"price"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// BeforeCreate generate uuid v4 and hashing password
func (car Car) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
