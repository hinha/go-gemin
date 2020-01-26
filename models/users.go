package models

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Country string `json:"country"`
	DateOfBirth time.Time `json:"DateOfBirth"`
	Profile []Profiles `json:"profile"`
	Joined_On timestamp.Timestamp `json:"joined_on"`
}

type Profiles struct {
	Update_On timestamp.Timestamp `json:"update_on"`
	ImageUrl string `json:"ImageUrl"`
	Confirmed bool `json:"confirmed"`
	Status int64 `json:"status"`
}

var (
	ctx = context.Background()
	ErrInvalidName = errors.New("invalid product name")
)

func RegisterUser(name, username, password,
	phone, email, country, imgurl, status string, confirmed bool) (user *Users, profile *Profiles, err error) {

	if name == "" && username == "" {
		err = ErrInvalidName
		return
	}
	return
}