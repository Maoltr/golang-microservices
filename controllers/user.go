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
	period, ok := c.Get("period")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"title":   "Incorrect raw",
			"message": "Enter period please",
		})

		return
	}

	duration, err := time.ParseDuration(period.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"title":   "Incorrect raw",
			"message": "Enter correct period",
		})

		return
	}

	messages, err := services.GetUsers(duration)

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
