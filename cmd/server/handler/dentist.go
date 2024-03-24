package handler

import (
	"errors"
	"repositoryapi/internal/dentist"
	"repositoryapi/internal/domain"
	"repositoryapi/pkg/web"
	"strconv"

	"github.com/gin-gonic/gin"
)

type dentistHandler struct {
	s dentist.DentistService
}

func NewProductHandler(s dentist.DentistService) *dentistHandler {
	return &dentistHandler{
		s: s,
	}
}

/* func (h *dentistHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		dentist, _ := h.s.GetAll()
		c.JSON(200, dentist)
	}
} */

func (h *dentistHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		dentist, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		web.Success(c, 200, dentist)

	}

}

func (handler *dentistHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentist domain.Dentist
		if err := c.BindJSON(&dentist); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		dentist, err := handler.s.CreateDentist(dentist)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, dentist)
	}
}

func (handler *dentistHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentist domain.Dentist
		if err := c.BindJSON(&dentist); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		dentist, err := handler.s.UpdateDentist(dentist)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, dentist)
	}
}

func (handler *dentistHandler) Patch() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentist domain.Dentist
		if err := c.BindJSON(&dentist); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		dentist, err := handler.s.PatchDentist(dentist)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, dentist)
	}
}

func (h *dentistHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		dentist, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		h.s.Delete(dentist.Id)
		web.Success(c, 200, true)

	}

}
