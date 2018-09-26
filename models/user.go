package models

import "time"

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Address   string    `json:"address"`
	Gender    string    `json:"gender"`
	TimeStamp time.Time `json:"time_stamp"`
}

type UserUI struct {
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Gender    string    `json:"gender"`
	TimeStamp time.Time `json:"time_stamp"`
}

func (user User) TransformUser() UserUI {
	return UserUI{Name: user.FirstName + " " + user.LastName, Address: user.Address, Gender: user.Gender, TimeStamp: user.TimeStamp}
}
