package user

import (
	"bitbucket.org/Milinel/golangContainer/models"
	"bitbucket.org/Milinel/golangContainer/services/databaseClient"
	"encoding/json"
	"testing"
	"time"
)

func TestSortUser(t *testing.T) {
	period := time.Hour

	var messages []databaseClient.SetMember
	users := []models.UserUI{
		{TimeStamp: time.Now().Add(-time.Hour * 2)},
		{TimeStamp: time.Now().Add(-time.Minute * 30)},
		{TimeStamp: time.Now()},
	}

	for _, v := range users {
		bytes, err := json.Marshal(v)
		if err != nil {
			t.Error(err)
		}

		messages = append(messages, databaseClient.SetMember{Member: string(bytes), Score: float64(v.TimeStamp.In(time.Local).Unix())})
	}

	result := sortUsers(period, messages)

	for _, v := range result {
		if v.TimeStamp == users[0].TimeStamp {
			t.Error("user with old timestamp was not deleted")
		}
	}
}

func TestSortUserWithInvalidUser(t *testing.T) {
	period := time.Hour

	var messages []databaseClient.SetMember

	messages = append(messages, databaseClient.SetMember{Score: float64(time.Now().Unix()), Member: "invalid data"})
	result := sortUsers(period, messages)

	if len(result) > 0 {
		t.Error("unmarshaled invalid data")
	}
}
