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
// GetById godoc
// swagger:parameters id query
// @Summary get a shift
// @Tags Shift
// @Description get shifts by id
// @Accept json
// @Produce json
// @Success 200 {shifts} domain.Dentist
// @Router /shifts/get/{id} [get]
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
// CreateShift godoc
// swagger:parameters id query
// @Summary create a shift
// @Tags Shift
// @Description get shift by id
// @Accept json
// @Produce json
// @Success 200 {shift} domain.Dentist
// @Router /shifts/get/{id} [post]
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
	
	// Put godoc
	// swagger:parameters id query
	// @Summary update a shift
	// @Tags Shift
	// @Description update shift
	// @Accept json
	// @Produce json
	// @Success 200 {shift} domain.Shift
	// @Router /shifts [put]
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

	// DeleteShift godoc
	// swagger:parameters id query
	// @Summary deletes a shift
	// @Tags Shift
	// @Description deletes shift by id
	// @Accept json
	// @Produce json
	// @Success 200 {shift} domain.Shift
	// @Router /shifts/delete/{id} [delete]
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

	// swagger:parameters id query
	// @Summary update a field shift
	// @Tags Shift
	// @Description update a field shift
	// @Accept json
	// @Produce json
	// @Success 200 {shift} domain.Shift
	// @Router /shifts [patch]
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
	