package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"DELETE", "PUT", "PATCH"}, // default allow GET and POST
		AllowHeaders:    []string{"authorization"},
	}))

	router.GET("/get", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "OK"})
	})
	router.POST("/post", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "OK"})
	})
	router.PUT("/put", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "OK"})
	})
	router.PATCH("/patch", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "OK"})
	})
	router.DELETE("/delete", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "OK"})
	})

	router.Run("0.0.0.0:8080")
}
