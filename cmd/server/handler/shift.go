package handler

import (
	
	"errors"
	"repositoryapi/pkg/web"
	"strconv"
	"github.com/gin-gonic/gin"
	"repositoryapi/internal/shift"

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