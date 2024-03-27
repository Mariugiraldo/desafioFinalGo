package handler

import (
	"bytes"
	"errors"
	"fmt"
	"io"
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
//
//	@Param			id	path	int	true	"Shift id"
//	@Summary		get a shift
//	@Tags			Shift
//	@Description	get shifts by id
//	@Accept			json
//	@Produce		json
//	@Success		200	{shifts}	domain.Shift
//	@Router			/shifts/{id} [get]
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
//
//	@Param			Authorization	header	string			true	"token"
//	@Param			shift			body	domain.Shift	true	"Shift"
//	@Summary		create a shift
//	@Tags			Shift
//	@Description	get shift by id
//	@Accept			json
//	@Produce		json
//	@Success		200	{shift}	domain.Shift
//	@Router			/shifts [post]
func (h *shiftHandler) CreateShift() gin.HandlerFunc {
	return func(c *gin.Context) {
		var shift domain.Shift
		if err := c.BindJSON(&shift); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		shift, err := h.service.CreateShift(shift)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, shift)
	}

}

// Put godoc
//
//	@Param			Authorization	header	string			true	"token"
//	@Param			shift			body	domain.Shift	true	"Shift"
//	@Summary		update a shift
//	@Tags			Shift
//	@Description	update shift
//	@Accept			json
//	@Produce		json
//	@Success		200	{shift}	domain.Shift
//	@Router			/shifts [put]
func (handler *shiftHandler) PutShift() gin.HandlerFunc {
	return func(c *gin.Context) {
		var shift domain.Shift
		if err := c.BindJSON(&shift); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		shift, err := handler.service.UpdateShift(shift)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, shift)
	}
}

// DeleteShift godoc
//
//	@Param			Authorization	header	string	true	"token"
//	@Param			id				path	string	true	"Shift"
//	@Summary		deletes a shift
//	@Tags			Shift
//	@Description	deletes shift by id
//	@Accept			json
//	@Produce		json
//	@Success		200	{shift}	domain.Shift
//	@Router			/shifts/{id} [delete]
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

// PatchShift godoc
//
//	@Param			Authorization	header	string			true	"token"
//	@Param			shift			body	domain.Shift	true	"Shift"
//	@Summary		update a field shift
//	@Tags			Shift
//	@Description	update a field shift
//	@Accept			json
//	@Produce		json
//	@Success		200	{shift}	domain.Shift
//	@Router			/shifts [patch]
func (handler *shiftHandler) Patch() gin.HandlerFunc {
	return func(c *gin.Context) {
		var shift domain.Shift
		if err := c.BindJSON(&shift); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		shift, err := handler.service.PatchShift(shift)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, shift)
	}
}

// CreateShiftByDNIAndRegistration godoc
//	@Param			Authorization	header	string			true	"token"
//	@Param			shift			body	domain.Shift	true	"Shift"
//	@Param			dni				path	string			true	"dni"
//	@Param			registration	path	string			true	"registration"
//	@Summary		create a shift by dni and registration
//	@Tags			Shift
//	@Description	create a shift by dni and registration
//	@Accept			json
//	@Produce		json
//	@Success		200	{shift}	domain.Shift
//	@Router			/shifts/{dni}/{registration} [post]
func (handler *shiftHandler) CreateShiftByDNIAndRegistration() gin.HandlerFunc {
	return func(c *gin.Context) {
		dni := c.Param("dni")
		registration := c.Param("registration")

		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to read request body"})
			return
		}

		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		fmt.Println("Request body: ", string(bodyBytes))

		var shift domain.Shift

		if err := c.ShouldBindJSON(&shift); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("juajuajua", dni, registration, shift)

		shift, err = handler.service.CreateShiftByDNIAndRegistration(dni, registration, shift)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, shift)
	}
}

// ReadShiftByDNI godoc
//	@Param			dni	path	string	true	"dni"
//	@Summary		get a shift by dni
//	@Tags			Shift
//	@Description	get a shift by dni
//	@Accept			json
//	@Produce		json
//	@Success		200	{shift}	domain.Shift
//	@Router			/shifts/{dni} [get]
func (handler *shiftHandler) ReadShiftByDNI() gin.HandlerFunc {
	return func(c *gin.Context) {
		dni := c.Param("dni")
		shift, err := handler.service.ReadShiftByDNI(dni)
		if err != nil {
			web.Failure(c, 404, errors.New("shift not found"))
			return
		}
		web.Success(c, 200, shift)
	}
}
