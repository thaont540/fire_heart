package entity

import "time"

var UserCollection = "users"

type User struct {
	Email string `idx:"{email},unique" json:"email" binding:"required"`
	//Password string `json:"password" binding:"required"`
	Name string `json:"name"`
	//VerifiedAt *time.Time
	//Created time.Time `bson:"_created" json:"_created"`
	//UpdatedAt time.Time `bson:"_modified" json:"_modified"`
	Career string `json:"career"`
	Image string `json:"image"`
	Description string `json:"description"`
	Address string `json:"address"`
	Language string `json:"language"`
	Birthday time.Time `json:"birthday"`
}
