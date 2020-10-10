package car

import (
	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/go-service-taxi/api/models"
	uuid "github.com/satori/go.uuid"
)

// Repository interface
type Repository interface {
	Fetch() (*gorm.DB, *[]models.Car, error)
	GetByID(id uuid.UUID) (*models.Car, error)
	GetByUserID(userID uuid.UUID) (*gorm.DB, *[]models.Car, error)
	Update(car *models.Car) error
	Store(car *models.Car) error
	Delete(id uuid.UUID) error
}
