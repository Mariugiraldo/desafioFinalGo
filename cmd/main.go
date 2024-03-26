package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"repositoryapi/cmd/docs"
	"repositoryapi/cmd/server/handler"
	"repositoryapi/internal/dentist"
	"repositoryapi/internal/patient"
	"repositoryapi/internal/shift"
	"repositoryapi/pkg/middleware"
	"repositoryapi/pkg/store"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/turnos-odontologia")
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
	dentistHandler := handler.NewProductHandler(service)

	storagePat := store.NewSqlStorePatient(db)
	repoPat := patient.NewPatientRepository(storagePat)
	servicePat := patient.NewPatientService(repoPat)
	patientHandler := handler.NewPatientHandler(servicePat)

	shiftRepo := shift.NewRepositoryShift(storage)
	shiftService := shift.NewServiceShift(shiftRepo)
	shiftHandler := handler.NewShiftHandler(shiftService)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	docs.SwaggerInfo.Host = "localhost:8080"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	dentists := r.Group("/dentists")
	{
		dentists.GET("/:id", dentistHandler.GetByID())
		dentists.POST("", middleware.Authentication(), dentistHandler.Post())
		dentists.PUT("", middleware.Authentication(), dentistHandler.Put())
		dentists.PATCH("", middleware.Authentication(), dentistHandler.Patch())
		dentists.DELETE("/:id", middleware.Authentication(), dentistHandler.Delete())
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
		shifts.POST("", middleware.Authentication(), shiftHandler.CreateShift())
		shifts.PUT("", middleware.Authentication(), shiftHandler.PutShift())
		shifts.DELETE("/:id", middleware.Authentication(), shiftHandler.DeleteShift())
		shifts.PATCH("", middleware.Authentication(), shiftHandler.Patch())

	}

	r.Run(":8080")

}
