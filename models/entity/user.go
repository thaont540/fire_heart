package entity

import (
	"fire_heart/utils"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UserCollection = "users"

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	UserName string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
	//VerifiedAt *time.Time
	//Created time.Time `bson:"_created" json:"_created"`
	//UpdatedAt time.Time `bson:"_modified" json:"_modified"`
}

func (user *User)GetJwtToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.UserName,
	})
	secretKey := utils.Env("JWT_SECRET")

	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}
