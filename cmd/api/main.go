package main

import (
	"bitbucket.org/Milinel/golangContainer/controllers"
	"github.com/gin-gonic/gin"
)

const address = "localhost:8080"

func main() {
	router := gin.Default()

	router.POST("persons/", controllers.PostUser)
	router.GET("persons/", controllers.GetUsers)

	router.Run(address)
}
