package main

import (
	"database/sql"
	"repositoryapi/cmd/server/handler"
	"repositoryapi/internal/product"
	"repositoryapi/pkg/store"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:29mayo1973@tcp(localhost:3306)/clinica_odontologica")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	storage := store.NewSqlStore(db)

	repo := product.NewRepository(storage)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	dentists := r.Group("/dentists")

	{

		dentists.GET("/:id", productHandler.GetByID())
		/* dentists.GET("", productHandler.GetAll()) */

	}
	r.Run(":8080")

}
