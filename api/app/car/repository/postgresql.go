package repository

import (
	"fmt"

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

func (carRepository *postgreCarRepo) GetByUserID(userID uuid.UUID) (*gorm.DB, *[]models.Car, error) {
	var err error
	var cars *[]models.Car = &[]models.Car{}

	db := carRepository.DB.Model(
		&models.Car{},
	).Set(
		"gorm:auto_preload",
		true,
	).Where(
		"user_id = ?",
		userID.String(),
	).Find(
		&cars,
	)
	err = db.Error
	if err != nil {
		return nil, nil, err
	}
	return db, cars, nil
}

func (carRepository *postgreCarRepo) Store(car *models.Car) error {
	var err error
	var selectedBrand *models.Brand = &models.Brand{}

	db := carRepository.DB.Model(
		&models.Brand{},
	).Where(
		"id = ?",
		car.BrandID,
	).Find(
		&selectedBrand,
	)
	if db.Error != nil {
		return fmt.Errorf("Brand ID not found")
	}

	car.BrandID = selectedBrand.ID.String()
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

func (carRepository *postgreCarRepo) Update(car *models.Car) (*models.Car, error) {
	var updatedCar *models.Car = &models.Car{}
	var selectedBrand *models.Brand = &models.Brand{}

	db := carRepository.DB.Model(
		&models.Brand{},
	).Where(
		"id = ?",
		car.BrandID,
	).Find(
		&selectedBrand,
	)
	if db.Error != nil {
		return nil, fmt.Errorf("Brand ID not found")
	}

	db = carRepository.DB.Model(
		&models.Car{},
	).Where(
		"id = ? AND user_id = ?",
		car.ID,
		car.UserID,
	).Update(
		models.Car{
			CarName:       car.CarName,
			BrandID:       selectedBrand.ID.String(),
			Condition:     car.Condition,
			Description:   car.Description,
			Specification: car.Specification,
			ImageURL:      car.ImageURL,
			Unit:          car.Unit,
			Price:         car.Price,
			UpdatedAt:     car.UpdatedAt,
		},
	).Where(
		"id = ? AND user_id = ?",
		car.ID.String(),
		car.UserID,
	).Find(
		&updatedCar,
	)
	if db.Error != nil {
		return nil, db.Error
	}
	return updatedCar, nil
}

func (carRepository *postgreCarRepo) Delete(id uuid.UUID, userID uuid.UUID) error {
	var err error
	var wantsToDeleted *models.Car = &models.Car{}

	err = carRepository.DB.Model(
		&models.Car{},
	).Where(
		"id = ? AND user_id = ?",
		id, userID,
	).Find(
		&wantsToDeleted,
	).Error
	if err != nil {
		return err
	}

	err = carRepository.DB.Model(
		&models.Car{},
	).Delete(&wantsToDeleted).Error
	if err != nil {
		return err
	}

	return nil
}
