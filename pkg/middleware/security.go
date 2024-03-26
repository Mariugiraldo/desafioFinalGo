package middleware

import (
	"repositoryapi/pkg/web"
	"errors"
	"os"
	"github.com/gin-gonic/gin"

)

func Authentication()gin.HandlerFunc{
	return func(context *gin.Context){
		token := context.GetHeader("TOKEN")
		if token == ""{
			web.Failure(context, 401, errors.New("token not found"))
			context.Abort()
			return
		}
		if token != os.Getenv("TOKEN"){
			web.Failure(context, 401, errors.New("invalid token"))
			context.Abort()
			return
		}
		context.Next()
	}
}
