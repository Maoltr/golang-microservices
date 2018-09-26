package main

import (
	"bitbucket.org/Milinel/golangContainer/controllers/user"
	"github.com/gin-gonic/gin"
)

const address = "localhost:8080"

func main() {
	router := gin.Default()

	router.POST("persons/", user.PostUser)
	router.GET("persons/:period", user.GetUsers)

	router.Run(address)
}
