package main

import (
	"database/sql"
	"repositoryapi/cmd/server/handler"
	"repositoryapi/internal/dentist"
	"repositoryapi/internal/patient"
	"repositoryapi/internal/shift"
	"repositoryapi/pkg/store"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:Jeifer05@tcp(localhost:3306)/turnos-odontologia")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	storage := store.NewSqlStore(db)

	repo := dentist.NewRepository(storage)
	service := dentist.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	storagePat := store.NewSqlStorePatient(db)
	repoPat := patient.NewPatientRepository(storagePat)
	servicePat := patient.NewPatientService(repoPat)
	patientHandler := handler.NewPatientHandler(servicePat)

	storageshi := store.NewSqlStoreShift(db)
	repoShi := shift.NewShiftRepository(storageshi)
	serviceShi := shift.NewShiftService(repoShi)
	shiftHandler := handler.NewShiftHandler(serviceShi)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	dentists := r.Group("/dentists")
	{

		dentists.GET("/:id", productHandler.GetByID())
		/* dentists.GET("", productHandler.GetAll()) */
		dentists.POST("", productHandler.Post())
		dentists.PUT("", productHandler.Put())
		dentists.PATCH("", productHandler.Patch())
		dentists.DELETE("/:id", productHandler.Delete())
	}

	patients := r.Group("/patients")
	{
		patients.GET("/all", patientHandler.FindAll())
		patients.GET("/:id", patientHandler.FindPatientById())
		patients.POST("", patientHandler.CreatePatient())
		patients.PUT("", patientHandler.UpdatePatient())
		patients.PATCH("", patientHandler.PatchPatient())
		patients.DELETE("/:id", patientHandler.DeletePatient)
	}

	shifts := r.Group("/shifts")
	{
		shifts.GET("/all", shiftHandler.FindAll())
		shifts.GET("/:id", shiftHandler.FindShiftById())
		shifts.POST("", shiftHandler.CreateShift())
		shifts.PUT("", shiftHandler.UpdateShift())
		shifts.PATCH("", shiftHandler.PatchShift())
		shifts.DELETE("/:id", shiftHandler.DeleteShift())
	}

	r.Run(":8080")

}
