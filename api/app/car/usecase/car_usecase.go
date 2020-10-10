package usecase

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/go-service-taxi/api/app/car"
	"github.com/muhammadisa/go-service-taxi/api/models"
	uuid "github.com/satori/go.uuid"
)

// carUsecase struct
type carUsecase struct {
	carRepository car.Repository
}

// NewCarUsecase function
func NewCarUsecase(carRepository car.Repository) car.Usecase {
	return &carUsecase{
		carRepository: carRepository,
	}
}

func (carUsecases carUsecase) Fetch(query string) (*gorm.DB, *[]models.Car, error) {
	db, res, err := carUsecases.carRepository.Fetch(query)
	if err != nil {
		return nil, nil, err
	}
	return db, res, nil
}

func (carUsecases carUsecase) GetByID(id uuid.UUID) (*models.Car, error) {
	res, err := carUsecases.carRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (carUsecases carUsecase) GetByUserID(userID uuid.UUID) (*gorm.DB, *[]models.Car, error) {
	db, res, err := carUsecases.carRepository.GetByUserID(userID)
	if err != nil {
		return nil, nil, err
	}
	return db, res, nil
}

func (carUsecases carUsecase) Store(car *models.Car, userID uuid.UUID) error {
	car.UserID = userID.String()
	err := carUsecases.carRepository.Store(car)
	if err != nil {
		return err
	}
	return nil
}

func (carUsecases carUsecase) Update(car *models.Car, userID uuid.UUID) (*models.Car, error) {
	car.UserID = userID.String()
	car.UpdatedAt = time.Now()
	return carUsecases.carRepository.Update(car)
}

func (carUsecases carUsecase) Delete(id uuid.UUID, userID uuid.UUID) error {
	return carUsecases.carRepository.Delete(id, userID)
}
