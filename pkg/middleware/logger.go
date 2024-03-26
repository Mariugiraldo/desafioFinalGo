package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		verb := context.Request.Method
		time := time.Now()
		patch := context.Request.URL.Path

		context.Next()

		var size int
		if context.Writer != nil {
			size = context.Writer.Size()

		}

		fmt.Printf("verb: %v\ntime: %v\npatch: %v\nsize: %v\n", verb, time, patch, size)

	}
}
