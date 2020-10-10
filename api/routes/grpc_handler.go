package routes

import (
	"fmt"
	"net"

	_grpcFoobar "github.com/muhammadisa/go-service-taxi/api/app/foobar/delivery/grpc"
	_foobarRepo "github.com/muhammadisa/go-service-taxi/api/app/foobar/repository"
	_foobarUsecase "github.com/muhammadisa/go-service-taxi/api/app/foobar/usecase"

	_grpcCar "github.com/muhammadisa/go-service-taxi/api/app/car/delivery/grpc"
	_carRepo "github.com/muhammadisa/go-service-taxi/api/app/car/repository"
	_carUsecase "github.com/muhammadisa/go-service-taxi/api/app/car/usecase"

	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/go-service-taxi/api/cache"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GRPCConfigs struct
type GRPCConfigs struct {
	Port     string
	Protocol string
	DB       *gorm.DB
	Cache    cache.Redis
}

// IGRPCConfigs interface
type IGRPCConfigs interface {
	NewGRPC()
}

// NewGRPC grpc service initialization
func (gc GRPCConfigs) NewGRPC() {
	// Initialize grpc server
	listener, err := net.Listen(gc.Protocol, gc.Port)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while listening on %s", gc.Port))
	}
	fmt.Println(fmt.Sprintf("gRPC Server is Listening on %s", gc.Port))
	server := grpc.NewServer()

	// Init grpc services
	gc.initFoobarService(server)
	gc.initCarService(server)

	// Serve grpc
	reflection.Register(server)
	err = server.Serve(listener)
	if err != nil {
		fmt.Println("Unexpected Error", err)
	}
}

func (gc GRPCConfigs) initFoobarService(server *grpc.Server) {
	foobarRepo := _foobarRepo.NewPostgresFoobarRepo(gc.DB, gc.Cache)
	foobarUsecase := _foobarUsecase.NewFoobarUsecase(foobarRepo)
	_grpcFoobar.NewFoobarServerGrpc(server, foobarUsecase)
}

func (gc GRPCConfigs) initCarService(server *grpc.Server) {
	carRepo := _carRepo.NewPostgresCarRepo(gc.DB, gc.Cache)
	carUsecase := _carUsecase.NewCarUsecase(carRepo)
	_grpcCar.NewCarServerGrpc(server, carUsecase)
}
