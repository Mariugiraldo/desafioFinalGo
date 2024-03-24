package handler

import (
	"errors"
	"net/http"
	"repositoryapi/internal/domain"
	"repositoryapi/internal/shift"
	"repositoryapi/pkg/web"
	"strconv"

	"github.com/gin-gonic/gin"
)

type shiftHandler struct {
	s shift.ShiftServiceInterface
}

func NewShiftHandler(s shift.ShiftServiceInterface) *shiftHandler {
	return &shiftHandler{s}
}

func (h *shiftHandler) FindAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		shifts, _ := h.s.ReadAllShift()
		c.JSON(200, shifts)
	}
}

func (h *shiftHandler) FindShiftById() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		shift, err := h.s.FindShiftById(id)
		if err != nil {
			web.Failure(c, 404, errors.New("shift not found"))
			return
		}
		web.Success(c, http.StatusOK, shift)

	}
}

func (handler *shiftHandler) CreateShift() gin.HandlerFunc {
	return func(c *gin.Context) {
		var shift domain.Shift
		if err := c.BindJSON(&shift); err != nil {
			web.Failure(c, 400, err)
			return
		}

		shift, err := handler.s.CreateShift(shift)
		if err != nil {
			web.Failure(c, 500, err)
			return
		}
		web.Success(c, http.StatusCreated, shift)

	}
}

func (handler *shiftHandler) UpdateShift() gin.HandlerFunc {
	return func(c *gin.Context) {
		var shift domain.Shift
		if err := c.BindJSON(&shift); err != nil {
			web.Failure(c, 400, err)
			return
		}

		shift, err := handler.s.UpdateShift(shift)
		if err != nil {
			web.Failure(c, 500, err)
			return
		}
		web.Success(c, http.StatusOK, shift)
	}
}

func (handler *shiftHandler) PatchShift() gin.HandlerFunc {
	return func(c *gin.Context) {
		var shift domain.Shift
		if err := c.BindJSON(&shift); err != nil {
			web.Failure(c, 400, err)
			return
		}

		shift, err := handler.s.PatchShift(shift)
		if err != nil {
			web.Failure(c, 500, err)
			return
		}
		web.Success(c, http.StatusOK, shift)
	}
}

func (handler *shiftHandler) DeleteShift() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = handler.s.DeleteShift(id)
		if err != nil {
			web.Failure(c, 500, err)
			return
		}
		web.Success(c, http.StatusOK, nil)
	}
}
