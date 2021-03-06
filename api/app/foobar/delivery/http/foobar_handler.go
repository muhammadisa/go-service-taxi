package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhammadisa/go-service-taxi/api/app/foobar"
	"github.com/muhammadisa/go-service-taxi/api/models"
	"github.com/muhammadisa/go-service-taxi/api/response"
	"github.com/muhammadisa/go-service-taxi/api/utils/message"
	"github.com/muhammadisa/go-service-taxi/api/utils/paging"
	uuid "github.com/satori/go.uuid"
)

// FoobarHandler struct
type FoobarHandler struct {
	fbUsecase foobar.Usecase
}

// NewFoobarHandler initialize enpoints
func NewFoobarHandler(e *echo.Group, fBu foobar.Usecase) {
	handler := &FoobarHandler{
		fbUsecase: fBu,
	}
	e.GET("/foobars/", handler.Fetch)
	e.GET("/foobar/:id", handler.GetByID)
	e.POST("/foobar/", handler.Store)
	e.PATCH("/foobar/update/:id", handler.Update)
	e.DELETE("/foobar/delete/:id", handler.Delete)
}

var (
	model = models.Foobar{}
)

// Fetch foobar data
func (fbHandler *FoobarHandler) Fetch(c echo.Context) error {
	var err error

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	db, fBars, err := fbHandler.fbUsecase.Fetch()
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
		Data:       paging.GetPaginator(db, page, limit, fBars),
	})
}

// GetByID foobar data
func (fbHandler *FoobarHandler) GetByID(c echo.Context) error {
	var err error

	uid, err := uuid.FromString(c.Param("id"))
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	fBar, err := fbHandler.fbUsecase.GetByID(uid)
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
		Data:       fBar,
	})
}

// Store foobar data
func (fbHandler *FoobarHandler) Store(c echo.Context) error {
	var err error
	var fooBar models.Foobar

	err = c.Bind(&fooBar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = c.Validate(fooBar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = fbHandler.fbUsecase.Store(&fooBar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Message:    message.GenerateMessage(fooBar.ID, "POST", model, true),
		Data:       fooBar,
	})
}

// Update foobar data
func (fbHandler *FoobarHandler) Update(c echo.Context) error {
	var err error
	var fooBar models.Foobar

	err = c.Bind(&fooBar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = c.Validate(fooBar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    message.GenerateMessage(uuid.Nil, "POST", model, false),
			Data:       err,
		})
	}
	_, err = fbHandler.fbUsecase.GetByID(fooBar.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = fbHandler.fbUsecase.Update(&fooBar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Message:    message.GenerateMessage(fooBar.ID, "PATCH", model, true),
		Data:       fooBar,
	})
}

// Delete foobar data
func (fbHandler *FoobarHandler) Delete(c echo.Context) error {
	var err error

	uid, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	fBar, err := fbHandler.fbUsecase.GetByID(uid)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err = fbHandler.fbUsecase.Delete(fBar.ID)
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
		Data:       fBar.ID,
	})
}
