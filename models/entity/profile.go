package entity

import "time"

var ProfileCollection = "profiles"

type Profile struct {
	UserId string `idx:"{user_id},unique" json:"user_id" binding:"required"`
	Email string `idx:"{email},unique" json:"email" binding:"required"`
	Name string `json:"name"`
	Career string `json:"career"`
	Image string `json:"image"`
	Description string `json:"description"`
	Address string `json:"address"`
	Language string `json:"language"`
	Birthday time.Time `json:"birthday"`
}
