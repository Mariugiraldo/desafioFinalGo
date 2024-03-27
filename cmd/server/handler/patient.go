package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"repositoryapi/internal/domain"
	"repositoryapi/internal/patient"
	"repositoryapi/pkg/web"
	"strconv"
)

type patientHandler struct {
	s patient.PatientServiceInterface
}

func NewPatientHandler(s patient.PatientServiceInterface) *patientHandler {
	return &patientHandler{
		s: s,
	}
}

// Findall godoc
//
//	@Summary		get all patients
//	@Tags			Patient
//	@Description	get all patients
//	@Accept			json
//	@Produce		json
//	@Success		200	{patients}	domain.Patient
//	@Router			/patients  [get]
func (h *patientHandler) FindAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		patients, _ := h.s.ReadAllPatient()
		c.JSON(200, patients)
	}
}

// FindPatientById godoc
//
//	@Param			id	path	int	true	"Patient id"
//	@Summary		get a patient
//	@Tags			Patient
//	@Description	get patients by id
//	@Accept			json
//	@Produce		json
//	@Success		200	{patients}	domain.Patient
//	@Router			/patients/{id} [get]
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

// CreatePatient godoc
//
//	@Param			Authorization	header	string			true	"token"
//	@Param			patient			body	domain.Patient	true	"patient"
//	@Summary		create a patient
//	@Tags			Patient
//	@Description	get patient by id
//	@Accept			json
//	@Produce		json
//	@Success		200	{patient}	domain.Patient
//	@Router			/patients [post]
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

// UpdatePatient godoc
//
//	@Param			Authorization	header	string			true	"token"
//	@Param			patient			body	domain.Patient	true	"Patient"
//	@Summary		update a patient
//	@Tags			Patient
//	@Description	update patient
//	@Accept			json
//	@Produce		json
//	@Success		200	{patient}	domain.Patient
//	@Router			/patients [put]
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

// PatchPatient godoc
//
//	@Param			Authorization	header	string			true	"token"
//	@Param			patient			body	domain.Patient	true	"Patient"
//	@Summary		update a field patient
//	@Tags			Patient
//	@Description	update a field patient
//	@Accept			json
//	@Produce		json
//	@Success		200	{patient}	domain.patient
//	@Router			/patients [patch]
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

// DeletePatient godoc
//
//	@Param			Authorization	header	string	true	"token"
//	@Param			id			path		string	true	"Patient"
//	@Summary		deletes a patient
//	@Tags			Patient
//	@Description	deletes patient by id
//	@Accept			json
//	@Produce		json
//	@Success		200	{patient}	domain.Patient
//	@Router			/patients/{id} [delete]
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
