package main

import (
	"database/sql"
	"repositoryapi/cmd/docs"
	"repositoryapi/cmd/server/handler"
	"repositoryapi/internal/dentist"
	"repositoryapi/internal/patient"
	"repositoryapi/internal/shift"
	"repositoryapi/pkg/store"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	shiftRepo := shift.NewRepositoryShift(storage)
	shiftService := shift.NewServiceShift(shiftRepo)
	shiftHandler := handler.NewShiftHandler(shiftService)

	r := gin.Default()

	docs.SwaggerInfo.Host = "localhost"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	dentists := r.Group("/dentists")
	{
		dentists.GET("/:id", productHandler.GetByID())
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
		shifts.GET("/:id", shiftHandler.GetByIDShift())
		shifts.POST("", shiftHandler.CreateShift())
		shifts.PUT("", shiftHandler.PutShift())
		shifts.DELETE("/:id", shiftHandler.DeleteShift())
		shifts.PATCH("", shiftHandler.Patch())
	}

	r.Run(":8080")

}
