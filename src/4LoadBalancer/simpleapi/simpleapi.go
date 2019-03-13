package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	port := os.Args[1]

	ngn := gin.Default()
	ngn.GET("/ping",
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "pong from "+port)
		},
	)

	_ = ngn.Run(":" + port)
}
