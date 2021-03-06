package user

import (
	"github.com/muhammadisa/go-service-taxi/api/auth"
	"github.com/muhammadisa/go-service-taxi/api/models"
	uuid "github.com/satori/go.uuid"
)

// Usecase interface
type Usecase interface {
	Login(usr *models.User) (*models.User, *auth.Authenticated, error)
	Register(usr *models.User) error
	Update(usr *models.User) error
	Delete(id uuid.UUID) error
}
