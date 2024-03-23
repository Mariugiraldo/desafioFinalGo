package handler

import (
	"errors"
	"repositoryapi/internal/domain"
	"repositoryapi/internal/product"
	"repositoryapi/pkg/web"
	"strconv"

	"github.com/gin-gonic/gin"
)

type shiftHandler struct {
	service product.ShiftService
}

func NewProductShiftHandler(service product.ShiftRepository)*shiftHandler{
	return &shiftHandler{
		service: service,
	}
}

func (h *shiftHandler)AddShift()gin.HandlerFunc{
	return func (c *gin.Context){
		var appoinment domain.Shift
		if err := c.BindJSON(&appoinment); err != nil{
			web.Failure(c, 400, errors.New("invalid request body"))
			return
		}

		err := h.service.AddShift(appoinment)
		if err != nil{
			web.Failure(c, 500, errors.New("failed to add shift"))
			return
		}
	}

}


func (h *shiftHandler) GetByIDShift() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		shift, err := h.service.GetByIDShift(id)
		if err != nil {
			web.Failure(c, 404, errors.New("shift not found"))
			return
		}
		web.Success(c, 200, shift)

	}

}