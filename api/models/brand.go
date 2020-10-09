package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Brand struct
type Brand struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	BrandName string    `gorm:"size:255;not null" json:"employee_name" validate:"required"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// BeforeCreate generate uuid v4 and hashing password
func (brand Brand) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
