package middleware

import (
	"repositoryapi/pkg/web"
	"errors"
	"github.com/gin-gonic/gin"

)

func Authentication()gin.HandlerFunc{
	return func(context *gin.Context){
		token := context.GetHeader("Authorization")
		if token == ""{
			web.Failure(context, 401, errors.New("token not found"))
			context.Abort()
			return
		}
		if token != "Bearer dG9rZW4tdmFsaWRvLWZpbmFsLWdv"{
			web.Failure(context, 401, errors.New("invalid token"))
			context.Abort()
			return
		}
		context.Next()
	}
}
