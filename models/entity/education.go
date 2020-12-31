package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var EducationCollection = "educations"

type Education struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	UserId string `idx:"{user_id},unique" json:"user_id" binding:"required"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	School string `json:"school"`
	Major string `json:"major"`
	Description string `json:"description"`
	Certificate string `json:"certificate"`
}
