package main

import (
	"database/sql"
	"repositoryapi/cmd/docs"
	"repositoryapi/cmd/server/handler"
	"repositoryapi/internal/dentist"
	"repositoryapi/internal/shift"
	"repositoryapi/pkg/store"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	shiftRepo := shift.NewRepositoryShift(storage)
	shiftService := shift.NewServiceShift(shiftRepo)
	shiftHandler := handler.NewShiftHandler(shiftService)

	r := gin.Default()

	docs.SwaggerInfo.Host = "localhost"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	dentists := r.Group("/dentists")

	{
		dentists.GET("/:id", dentistHandler.GetByID())
		dentists.POST("", dentistHandler.Post())
		dentists.PUT("", dentistHandler.Put())
		dentists.PATCH("", dentistHandler.Patch())
		dentists.DELETE("/:id", dentistHandler.Delete())

	}

	shifts := r.Group("/shifts")
	{
		shifts.GET("/:id", shiftHandler.GetByIDShift())
		shifts.POST("", shiftHandler.CreateShift())
		shifts.PUT("", shiftHandler.PutShift())
		shifts.DELETE("/:id", shiftHandler.DeleteShift())
		shifts.PATCH("", shiftHandler.Patch())
		shifts.POST("", shiftHandler.CreateShiftByDni())
		/* shifts.GET("/shifts", shiftHandler.GetShiftsByPatientDNI()) */
	}

	r.Run(":8082")

}
