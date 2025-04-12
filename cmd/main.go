package main

import (
	"fmt"
	"gin_http/cmd/middleware"
	"gin_http/cmd/routes"
	"gin_http/cmd/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//Middleware global
	r.Use(middleware.LoggerMiddleware())
	//Services
	userService := services.NewUserService()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  "ok",
		})
	})

	routes.SetupUserRoutes(r, userService)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	fmt.Println("Listening at port 3000")

}
