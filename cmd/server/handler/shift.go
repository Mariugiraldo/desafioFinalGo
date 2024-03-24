package handler

import (
	"errors"
	"repositoryapi/internal/domain"
	"repositoryapi/internal/shift"
	"repositoryapi/pkg/web"
	"strconv"

	"github.com/gin-gonic/gin"
)

type shiftHandler struct {
    service shift.ShiftService
}

func NewShiftHandler(s shift.ShiftService) *shiftHandler {
    return &shiftHandler{service: s}
}

func (h *shiftHandler) GetByIDShift() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		shift, err := h.service.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("shift not found"))
			return
		}
		web.Success(c, 200, shift)

	}

}

func ( h *shiftHandler) CreateShift() gin.HandlerFunc{
	return func(c *gin.Context){
		var shift domain.Shift
		if err := c.BindJSON(&shift); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		shift, err := h.service.CreateShift(shift)
		if err != nil {
			c.JSON(500,gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, shift)
	}

	}

	func (handler *shiftHandler) PutShift() gin.HandlerFunc {
		return func(c *gin.Context) {
			var shift domain.Shift
			if err := c.BindJSON(&shift); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
	
			shift, err := handler.service.UpdateShift(shift)
			if err != nil {
				c.JSON(500,gin.H{"error": err.Error()})
				return
			}
	
			c.JSON(200, shift)
		}
	}

	func (h *shiftHandler) DeleteShift() gin.HandlerFunc {
		return func(c *gin.Context) {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				web.Failure(c, 400, errors.New("invalid id"))
				return
			}
			h.service.DeleteShift(id)
			web.Success(c, 200, true)
	
		}

	}

	func (handler *shiftHandler) Patch() gin.HandlerFunc {
		return func(c *gin.Context) {
			var shift domain.Shift
			if err := c.BindJSON(&shift); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
	
			shift, err := handler.service.PatchShift(shift)
			if err != nil {
				c.JSON(500,gin.H{"error": err.Error()})
				return
			}
	
			c.JSON(200, shift)
		}
	}
	