package user

import (
	"bitbucket.org/Milinel/golangContainer/services/user"
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

	err = user.PushJSON(message)

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
}

func GetUsers(c *gin.Context) {
	period := c.Param("period")

	duration, err := time.ParseDuration(period)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"title":   "Incorrect raw",
			"message": "Enter correct period, err: " + err.Error(),
		})

		return
	}

	messages, err := user.GetUsers(duration)

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
