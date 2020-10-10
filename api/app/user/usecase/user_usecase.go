package usecase

import (
	"errors"
	"time"

	"github.com/muhammadisa/go-service-taxi/api/app/user"
	"github.com/muhammadisa/go-service-taxi/api/auth"
	"github.com/muhammadisa/go-service-taxi/api/models"
	uuid "github.com/satori/go.uuid"
)

// userUsecase struct
type userUsecase struct {
	userRepository user.Repository
}

// NewUserUsecase function
func NewUserUsecase(userRepository user.Repository) user.Usecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (userUsecase userUsecase) Login(user *models.User) (*models.User, *auth.Authenticated, error) {
	userResult, authenticated, err := userUsecase.userRepository.Login(user)
	if err != nil {
		return nil, nil, err
	}
	err = auth.VerifyPassword(userResult.Password, user.Password)
	if err != nil {
		return nil, nil, errors.New("Email or Password is incorrect")
	}
	token, refresh, err := auth.GenerateToken(userResult.ID)
	if err != nil {
		return nil, nil, err
	}
	authenticated.UserID = userResult.ID
	authenticated.AccessToken = token
	authenticated.RefreshToken = refresh
	return userResult, authenticated, nil
}

func (userUsecase userUsecase) Register(user *models.User) error {
	err := userUsecase.userRepository.Register(user)
	if err != nil {
		return err
	}
	return nil
}

func (userUsecase userUsecase) Update(user *models.User) error {
	user.UpdatedAt = time.Now()
	err := userUsecase.userRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (userUsecase userUsecase) Delete(id uuid.UUID) error {
	return userUsecase.userRepository.Delete(id)
}
