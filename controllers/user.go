package controllers

import (
	"bitbucket.org/Milinel/golangContainer/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func PostUser(c *gin.Context) {
	message, err := c.GetRawData()
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"title":   "Incorrect raw",
			"message": err.Error(),
		})
	}

	err = services.PushJSON(message)

	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"title":   "Incorrect data",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":   "Successful",
		"message": "Successful post user",
	})

	logrus.Println("Success")
}

func GetUsers(c *gin.Context) {
	messages, err := services.GetUsers(time.Duration(time.Hour*2))

	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"title":   "Incorrect raw",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":   "Successful",
		"messages": messages,
	})
}
