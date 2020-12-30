package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var ExperienceCollection = "experiences"

type Experience struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	UserId string `json:"user_id"`
	StartDate time.Time `json:"start_date" bson:"startdate,omitempty"`
	EndDate time.Time `json:"end_date"`
	Position string `json:"position"`
	Project string `json:"project"`
	Description string `json:"description"`
	Technical string `json:"technical"`
	TeamSize string `json:"team_size"`
	Effort string `json:"effort"`
}