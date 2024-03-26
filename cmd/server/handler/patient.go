package handler

import (
	"errors"
	"net/http"
	"repositoryapi/internal/domain"
	"repositoryapi/internal/patient"
	"repositoryapi/pkg/web"
	"strconv"

	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	s patient.PatientServiceInterface
}

func NewPatientHandler(s patient.PatientServiceInterface) *patientHandler {
	return &patientHandler{
		s: s,
	}
}

func (h *patientHandler) FindAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		patients, _ := h.s.ReadAllPatient()
		c.JSON(200, patients)
	}
}

func (h *patientHandler) FindPatientById() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		patient, err := h.s.FindPatientById(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		web.Success(c, http.StatusOK, patient)

	}
}

func (handler *patientHandler) CreatePatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		var patient domain.Patient
		if err := c.BindJSON(&patient); err != nil {
			web.Failure(c, 400, err)
			return
		}

		patient, err := handler.s.CreatePatient(patient)
		if err != nil {
			web.Failure(c, 500, err)
			return
		}
		web.Success(c, http.StatusOK, patient)
	}
}

func (handler *patientHandler) UpdatePatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		var patient domain.Patient
		if err := c.BindJSON(&patient); err != nil {
			web.Failure(c, 400, err)
			return
		}

		patient, err := handler.s.UpdatePatient(patient)
		if err != nil {
			web.Failure(c, 500, err)
			return
		}
		c.JSON(200, patient)
	}
}

func (handler *patientHandler) PatchPatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		var patient domain.Patient
		if err := c.BindJSON(&patient); err != nil {
			web.Failure(c, 400, err)
			return
		}

		patient, err := handler.s.PatchPatient(patient)
		if err != nil {
			web.Failure(c, 500, err)
			return
		}
		c.JSON(200, patient)
	}
}

func (handler *patientHandler) DeletePatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.Failure(c, 400, errors.New("invalid id"))
		return
	}
	err = handler.s.DeletePatient(int(id))
	if err != nil {
		web.Failure(c, 500, err)
		return
	}
	web.Success(c, http.StatusOK, nil)
}
