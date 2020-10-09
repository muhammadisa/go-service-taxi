package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/go-service-taxi/api/app/car"
	"github.com/muhammadisa/go-service-taxi/api/cache"
	"github.com/muhammadisa/go-service-taxi/api/models"
	uuid "github.com/satori/go.uuid"
)

type postgreCarRepo struct {
	DB    *gorm.DB
	Cache cache.Redis
}

// NewPostgresCarRepo function
func NewPostgresCarRepo(db *gorm.DB, cacheClient cache.Redis) car.Repository {
	return &postgreCarRepo{
		DB:    db,
		Cache: cacheClient,
	}
}

func (carRepository *postgreCarRepo) Fetch() (*gorm.DB, *[]models.Car, error) {
	var err error
	var cars *[]models.Car = &[]models.Car{}

	db := carRepository.DB.Model(
		&models.Car{},
	).Set(
		"gorm:auto_preload",
		true,
	).Order(
		"created_at desc",
	).Find(
		&cars,
	)
	err = db.Error
	if err != nil {
		return nil, nil, err
	}
	return db, cars, nil
}

func (carRepository *postgreCarRepo) GetByID(id uuid.UUID) (*models.Car, error) {
	var err error
	var car *models.Car = &models.Car{}

	// cache := carRepository.Cache.Get(cache.Key(models.Car{}, id))
	// if cache != "nil" && cache != "" {
	// 	json.Unmarshal([]byte(cache), &car)
	// 	return car, nil
	// }
	err = carRepository.DB.Model(
		&models.Car{},
	).Set(
		"gorm:auto_preload",
		true,
	).Where(
		"id = ?",
		id,
	).First(
		&car,
	).Error
	if err != nil {
		return nil, err
	}
	// carRepository.Cache.Set(*car)
	return car, nil
}

func (carRepository *postgreCarRepo) Store(car *models.Car) error {
	var err error

	err = carRepository.DB.Model(
		&models.Car{},
	).Create(
		car,
	).Error
	if err != nil {
		return err
	}

	return nil
}

func (carRepository *postgreCarRepo) Update(car *models.Car) error {
	var err error

	err = carRepository.DB.Model(
		&models.Car{},
	).Where(
		"id = ?",
		car.ID.String(),
	).Update(models.Car{
		CarName:       car.CarName,
		BrandID:       car.BrandID,
		Condition:     car.Condition,
		Description:   car.Description,
		Specification: car.Specification,
		Unit:          car.Unit,
		Price:         car.Price,
		UpdatedAt:     car.UpdatedAt,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (carRepository *postgreCarRepo) Delete(id uuid.UUID) error {
	var err error

	err = carRepository.DB.Model(
		&models.Car{},
	).Where(
		"id = ?",
		id,
	).Delete(
		&models.Car{},
	).Error
	if err != nil {
		return err
	}
	return nil
}
