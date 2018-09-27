package consumer

import (
	"bitbucket.org/Milinel/golangContainer/models"
	"encoding/json"
	"testing"
	"time"
)

func TestPrepareMessage(t *testing.T) {
	message := models.User{FirstName: "Maksim", LastName: "Tretiak", TimeStamp: time.Now()}

	result, err := prepareMessage(message)
	if err != nil {
		t.Error(err)
	}

	if result.Member == nil {
		t.Error("invalid Z member")
	}

	user := models.UserUI{}
	if err := json.Unmarshal(result.Member.([]byte), &user); err != nil {
		t.Error(err)
	}

	if (message.FirstName + " " + message.LastName) != user.Name {
		t.Error("invalid name in result")
	}
}

func TestPrepareMessageWithoutName(t *testing.T) {
	message := models.User{LastName: "Tretiak", TimeStamp: time.Now()}

	_, err := prepareMessage(message)
	if err == nil {
		t.Error("prepare message couldn't parse user without first/last name, but it has parsed it")
	}
}

func TestPrepareMessageWithOldTime(t *testing.T) {
	message := models.User{FirstName: "Maksim", LastName: "Tretiak", TimeStamp: time.Now().Add(-time.Hour * 2)}

	_, err := prepareMessage(message)
	if err == nil {
		t.Error("prepare message couldn't parse user with old timestamp, but it has parsed it")
	}
}
