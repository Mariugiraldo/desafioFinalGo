package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"repositoryapi/internal/dentist"
	"repositoryapi/internal/domain"
	"repositoryapi/pkg/web"
	"strconv"
)

type dentistHandler struct {
	s dentist.DentistService
}

func NewProductHandler(s dentist.DentistService) *dentistHandler {
	return &dentistHandler{
		s: s,
	}
}

// GetById godoc
//	@Param			id	path	int	true	"Dentist id"
//	@Summary		get a dentist
//	@Tags			Dentist
//	@Description	get dentist by id
//	@Accept			json
//	@Produce		json
//	@Success		200	{dentist}	domain.Dentist
//	@Router			/dentists/{id} [get]
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

// Post godoc
//	@Param			Authorization	header	string			true	"token"
//	@Param			dentist			body	domain.Dentist	true	"Dentist"
//	@Summary		create a dentist
//	@Tags			Dentist
//	@Description	create dentist
//	@Accept			json
//	@Produce		json
//	@Success		200	{dentist}	domain.Dentist
//	@Router			/dentists [post]
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

// Put godoc
//	@Param			Authorization	header	string			true	"token"
//	@Param			dentist			body	domain.Dentist	true	"Dentist"
//	@Summary		update a dentist
//	@Tags			Dentist
//	@Description	update dentist
//	@Accept			json
//	@Produce		json
//	@Success		200	{dentist}	domain.Dentist
//	@Router			/dentists [put]
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

// Patch godoc
//	@Param			Authorization	header	string			true	"token"
//	@Param			dentist			body	domain.Dentist	true	"Dentist"
//	@Summary		update a field dentist
//	@Tags			Dentist
//	@Description	update a field dentist
//	@Accept			json
//	@Produce		json
//	@Success		200	{dentist}	domain.Dentist
//	@Router			/dentists [patch]
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

// DeleteDentist godoc
//	@Param			Authorization	header	string	true	"token"
//	@Param			id				path	int		true	"Dentist id"
//	@Summary		deletes a dentist
//	@Tags			Dentist
//	@Description	deletes dentist by id
//	@Accept			json
//	@Produce		json
//	@Success		200	{dentist}	domain.Dentist
//	@Router			/dentists/{id} [delete]
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
