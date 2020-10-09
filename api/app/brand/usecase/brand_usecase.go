package usecase

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/go-service-taxi/api/app/brand"
	"github.com/muhammadisa/go-service-taxi/api/models"
	uuid "github.com/satori/go.uuid"
)

// brandUsecase struct
type brandUsecase struct {
	brandRepository brand.Repository
}

// NewBrandUsecase function
func NewBrandUsecase(brandRepository brand.Repository) brand.Usecase {
	return &brandUsecase{
		brandRepository: brandRepository,
	}
}

func (brandUsecases brandUsecase) Fetch() (*gorm.DB, *[]models.Brand, error) {
	db, res, err := brandUsecases.brandRepository.Fetch()
	if err != nil {
		return nil, nil, err
	}
	return db, res, nil
}

func (brandUsecases brandUsecase) GetByID(id uuid.UUID) (*models.Brand, error) {
	res, err := brandUsecases.brandRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (brandUsecases brandUsecase) Store(brand *models.Brand) error {
	err := brandUsecases.brandRepository.Store(brand)
	if err != nil {
		return err
	}
	return nil
}

func (brandUsecases brandUsecase) Update(brand *models.Brand) error {
	brand.UpdatedAt = time.Now()
	return brandUsecases.brandRepository.Update(brand)
}

func (brandUsecases brandUsecase) Delete(id uuid.UUID) error {
	return brandUsecases.brandRepository.Delete(id)
}
