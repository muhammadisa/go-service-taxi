package brand

import (
	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/go-service-taxi/api/models"
	uuid "github.com/satori/go.uuid"
)

// Repository interface
type Repository interface {
	Fetch() (*gorm.DB, *[]models.Brand, error)
	GetByID(id uuid.UUID) (*models.Brand, error)
	Update(brand *models.Brand) error
	Store(brand *models.Brand) error
	Delete(id uuid.UUID) error
}
