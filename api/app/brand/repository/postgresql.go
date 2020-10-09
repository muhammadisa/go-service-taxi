package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/go-service-taxi/api/app/brand"
	"github.com/muhammadisa/go-service-taxi/api/cache"
	"github.com/muhammadisa/go-service-taxi/api/models"
	uuid "github.com/satori/go.uuid"
)

type postgreBrandRepo struct {
	DB    *gorm.DB
	Cache cache.Redis
}

// NewPostgresBrandRepo function
func NewPostgresBrandRepo(db *gorm.DB, cacheClient cache.Redis) brand.Repository {
	return &postgreBrandRepo{
		DB:    db,
		Cache: cacheClient,
	}
}

func (brandRepository *postgreBrandRepo) Fetch() (*gorm.DB, *[]models.Brand, error) {
	var err error
	var brands *[]models.Brand = &[]models.Brand{}

	db := brandRepository.DB.Model(
		&models.Brand{},
	).Order(
		"created_at asc",
	).Find(
		&brands,
	)
	err = db.Error
	if err != nil {
		return nil, nil, err
	}
	return db, brands, nil
}

func (brandRepository *postgreBrandRepo) GetByID(id uuid.UUID) (*models.Brand, error) {
	var err error
	var brand *models.Brand = &models.Brand{}

	// cache := brandRepository.Cache.Get(cache.Key(models.Brand{}, id))
	// if cache != "nil" && cache != "" {
	// 	json.Unmarshal([]byte(cache), &brand)
	// 	return brand, nil
	// }
	err = brandRepository.DB.Model(
		&models.Brand{},
	).Where(
		"id = ?",
		id,
	).First(
		&brand,
	).Error
	if err != nil {
		return nil, err
	}
	// brandRepository.Cache.Set(*brand)
	return brand, nil
}

func (brandRepository *postgreBrandRepo) Store(brand *models.Brand) error {
	var err error

	err = brandRepository.DB.Model(
		&models.Brand{},
	).Create(
		brand,
	).Error
	if err != nil {
		return err
	}

	return nil
}

func (brandRepository *postgreBrandRepo) Update(brand *models.Brand) error {
	var err error

	err = brandRepository.DB.Model(
		&models.Brand{},
	).Where(
		"id = ?",
		brand.ID.String(),
	).Update(models.Brand{
		BrandName: brand.BrandName,
		UpdatedAt: brand.UpdatedAt,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (brandRepository *postgreBrandRepo) Delete(id uuid.UUID) error {
	var err error

	err = brandRepository.DB.Model(
		&models.Brand{},
	).Where(
		"id = ?",
		id,
	).Delete(
		&models.Brand{},
	).Error
	if err != nil {
		return err
	}
	return nil
}
