package grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/muhammadisa/go-service-taxi/api/app/car"
	"github.com/muhammadisa/go-service-taxi/api/models"
	"github.com/muhammadisa/go-service-taxi/api/protobuf/car_grpc"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

type server struct {
	usecase car.Usecase
}

// NewCarServerGrpc intialize handler
func NewCarServerGrpc(gserver *grpc.Server, carUsecase car.Usecase) {
	carServer := &server{
		usecase: carUsecase,
	}
	car_grpc.RegisterCarHandlerServer(gserver, carServer)
}

func (s *server) transformCarRPC(car *models.Car) *car_grpc.Car {
	if car == nil {
		return nil
	}
	UpdatedAt := &google_protobuf.Timestamp{
		Seconds: car.UpdatedAt.Unix(),
	}
	CraetedAt := &google_protobuf.Timestamp{
		Seconds: car.CreatedAt.Unix(),
	}

	res := &car_grpc.Car{
		ID:            car.ID.String(),
		UserID:        car.UserID,
		CarName:       car.CarName,
		BrandID:       car.BrandID,
		Condition:     car.Condition,
		Description:   car.Description,
		Specification: car.Specification,
		ImageURL:      car.ImageURL,
		Unit:          uint64(car.Unit),
		Price:         uint64(car.Price),
		UpdatedAt:     UpdatedAt,
		CreatedAt:     CraetedAt,
	}
	return res
}

func (s *server) transformCarData(car *car_grpc.Car) *models.Car {
	var UpdatedAt = time.Unix(time.Now().Unix(), 0)
	var CreatedAt = time.Unix(time.Now().Unix(), 0)

	id, err := uuid.FromString(car.ID)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	if id != uuid.Nil {
		CreatedAt = time.Unix(car.GetCreatedAt().GetSeconds(), 0)
		UpdatedAt = time.Unix(time.Now().Unix(), 0)
	}

	res := &models.Car{
		ID:            id,
		CarName:       car.CarName,
		UserID:        car.UserID,
		BrandID:       car.BrandID,
		Condition:     car.Condition,
		Description:   car.Description,
		Specification: car.Specification,
		ImageURL:      car.ImageURL,
		Unit:          int64(car.Unit),
		Price:         int64(car.Price),
		UpdatedAt:     UpdatedAt,
		CreatedAt:     CreatedAt,
	}
	return res
}

func (s *server) GetCarByID(
	ctx context.Context,
	in *car_grpc.RetrieveByID,
) (*car_grpc.Car, error) {

	uuid, err := uuid.FromString(in.ID)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	car, err := s.usecase.GetByID(uuid)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if car == nil {
		return nil, fmt.Errorf("Car is Not Found")
	}

	res := s.transformCarRPC(car)
	return res, nil

}

func (s *server) GetCars(
	ctx context.Context,
	in *car_grpc.RetrieveCars,
) (*car_grpc.ListCars, error) {

	_, res, err := s.usecase.Fetch()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var arrayCars []*car_grpc.Car
	for _, r := range *res {
		arrayCars = append(arrayCars, s.transformCarRPC(&r))
	}

	result := &car_grpc.ListCars{
		Cars: arrayCars,
	}

	return result, nil

}

func (s *server) GetCarsByUserID(
	ctx context.Context,
	in *car_grpc.RetrieveByUserID,
) (*car_grpc.ListCars, error) {

	userID, err := uuid.FromString(in.ID)
	if err != nil {
		return nil, err
	}

	_, res, err := s.usecase.GetByUserID(userID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var arrayCars []*car_grpc.Car
	for _, r := range *res {
		arrayCars = append(arrayCars, s.transformCarRPC(&r))
	}

	result := &car_grpc.ListCars{
		Cars: arrayCars,
	}

	return result, nil

}

func (s *server) Store(
	ctx context.Context,
	createCarRequest *car_grpc.CreateCar,
) (*car_grpc.Car, error) {

	car := s.transformCarData(createCarRequest.Car)
	userID, err := uuid.FromString(createCarRequest.UserID)
	if err != nil {
		return nil, err
	}

	err = s.usecase.Store(car, userID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	res := s.transformCarRPC(car)

	return res, nil

}

func (s *server) Update(
	ctx context.Context,
	in *car_grpc.UpdateRequest,
) (*car_grpc.Car, error) {

	car := s.transformCarData(in.Car)
	userID, err := uuid.FromString(in.UserID)
	if err != nil {
		return nil, err
	}

	updatedCar, err := s.usecase.Update(car, userID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	updated := s.transformCarRPC(updatedCar)

	return updated, nil
}

func (s *server) Delete(
	ctx context.Context,
	in *car_grpc.DeleteByID,
) (*car_grpc.DeleteResponse, error) {

	ID, err := uuid.FromString(in.ID)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}

	userID, err := uuid.FromString(in.UserID)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}

	err = s.usecase.Delete(ID, userID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		return nil, err
	}

	res := &car_grpc.DeleteResponse{
		Ok: 1,
	}

	return res, nil
}
