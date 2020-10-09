package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhammadisa/go-service-taxi/api/app/brand"
	"github.com/muhammadisa/go-service-taxi/api/models"
	"github.com/muhammadisa/go-service-taxi/api/response"
	"github.com/muhammadisa/go-service-taxi/api/utils/message"
	"github.com/muhammadisa/go-service-taxi/api/utils/paging"
	uuid "github.com/satori/go.uuid"
)

// BrandHandler struct
type BrandHandler struct {
	brandUsecase brand.Usecase
}

// NewBrandHandler initialize enpoints
func NewBrandHandler(e *echo.Group, brandUsecase brand.Usecase) {
	handler := &BrandHandler{
		brandUsecase: brandUsecase,
	}
	e.GET("/brands/", handler.Fetch)
	e.GET("/brand/:id", handler.GetByID)
	e.POST("/brand/", handler.Store)
	e.PATCH("/brand/update/:id", handler.Update)
	e.DELETE("/brand/delete/:id", handler.Delete)
}

var (
	model = models.Brand{}
)

// Fetch brand data
func (brandHandler *BrandHandler) Fetch(c echo.Context) error {
	var err error

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	db, brands, err := brandHandler.brandUsecase.Fetch()
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
		Data:       paging.GetPaginator(db, page, limit, brands),
	})
}

// GetByID brand data
func (brandHandler *BrandHandler) GetByID(c echo.Context) error {
	var err error

	uid, err := uuid.FromString(c.Param("id"))
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	brand, err := brandHandler.brandUsecase.GetByID(uid)
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
		Data:       brand,
	})
}

// Store brand data
func (brandHandler *BrandHandler) Store(c echo.Context) error {
	var err error
	var brand models.Brand

	err = c.Bind(&brand)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = c.Validate(brand)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = brandHandler.brandUsecase.Store(&brand)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Message:    message.GenerateMessage(brand.ID, "POST", model, true),
		Data:       brand,
	})
}

// Update brand data
func (brandHandler *BrandHandler) Update(c echo.Context) error {
	var err error
	var brand models.Brand

	err = c.Bind(&brand)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = c.Validate(brand)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    message.GenerateMessage(uuid.Nil, "POST", model, false),
			Data:       err,
		})
	}
	_, err = brandHandler.brandUsecase.GetByID(brand.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = brandHandler.brandUsecase.Update(&brand)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Message:    message.GenerateMessage(brand.ID, "PATCH", model, true),
		Data:       brand,
	})
}

// Delete brand data
func (brandHandler *BrandHandler) Delete(c echo.Context) error {
	var err error

	uid, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	brand, err := brandHandler.brandUsecase.GetByID(uid)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = brandHandler.brandUsecase.Delete(brand.ID)
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
		Data:       brand.ID,
	})
}
