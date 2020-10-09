package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhammadisa/go-service-taxi/api/app/car"
	"github.com/muhammadisa/go-service-taxi/api/auth"
	"github.com/muhammadisa/go-service-taxi/api/models"
	"github.com/muhammadisa/go-service-taxi/api/response"
	"github.com/muhammadisa/go-service-taxi/api/utils/message"
	"github.com/muhammadisa/go-service-taxi/api/utils/paging"
	uuid "github.com/satori/go.uuid"
)

// CarHandler struct
type CarHandler struct {
	carUsecase car.Usecase
}

// NewCarHandler initialize enpoints
func NewCarHandler(e *echo.Group, carUsecase car.Usecase) {
	handler := &CarHandler{
		carUsecase: carUsecase,
	}
	e.GET("/cars/", handler.Fetch)
	e.GET("/car/:id", handler.GetByID)
	e.POST("/car/", handler.Store)
	e.PATCH("/car/update/:id", handler.Update)
	e.DELETE("/car/delete/:id", handler.Delete)
}

var (
	model = models.Car{}
)

// Fetch car data
func (carHandler *CarHandler) Fetch(c echo.Context) error {
	var err error

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	db, cars, err := carHandler.carUsecase.Fetch()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    message.GenerateMessage(uuid.Nil, "GET", model, true),
		Data:       paging.GetPaginator(db, page, limit, cars),
	})
}

// GetByID car data
func (carHandler *CarHandler) GetByID(c echo.Context) error {
	var err error

	uid, err := uuid.FromString(c.Param("id"))
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	car, err := carHandler.carUsecase.GetByID(uid)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    message.GenerateMessage(uid, "GET", model, true),
		Data:       car,
	})
}

// Store car data
func (carHandler *CarHandler) Store(c echo.Context) error {
	var err error
	var car models.Car

	userID, err := auth.JWTExtractData(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	userUUID, err := uuid.FromString(userID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err = c.Bind(&car)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = c.Validate(car)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = carHandler.carUsecase.Store(&car, userUUID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Message:    message.GenerateMessage(car.ID, "POST", model, true),
		Data:       car,
	})
}

// Update car data
func (carHandler *CarHandler) Update(c echo.Context) error {
	var err error
	var car models.Car

	userID, err := auth.JWTExtractData(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	userUUID, err := uuid.FromString(userID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err = c.Bind(&car)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = c.Validate(car)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    message.GenerateMessage(uuid.Nil, "POST", model, false),
			Data:       err,
		})
	}
	_, err = carHandler.carUsecase.GetByID(car.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = carHandler.carUsecase.Update(&car, userUUID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Message:    message.GenerateMessage(car.ID, "PATCH", model, true),
		Data:       car,
	})
}

// Delete car data
func (carHandler *CarHandler) Delete(c echo.Context) error {
	var err error

	uid, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	car, err := carHandler.carUsecase.GetByID(uid)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = carHandler.carUsecase.Delete(car.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    message.GenerateMessage(uid, "DELETE", model, true),
		Data:       car.ID,
	})
}
