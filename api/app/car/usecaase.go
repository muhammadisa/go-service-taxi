package car

import (
	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/go-service-taxi/api/models"
	uuid "github.com/satori/go.uuid"
)

// Usecase interface
type Usecase interface {
	Fetch(query string) (*gorm.DB, *[]models.Car, error)
	GetByID(id uuid.UUID) (*models.Car, error)
	GetByUserID(userID uuid.UUID) (*gorm.DB, *[]models.Car, error)
	Update(car *models.Car, userID uuid.UUID) (*models.Car, error)
	Store(car *models.Car, userID uuid.UUID) error
	Delete(id uuid.UUID, userID uuid.UUID) error
}
