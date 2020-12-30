package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthController struct {}

func (authController *AuthController)Login(c *gin.Context) {
	type LoginInput struct {
		UserName string `json:"user_name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Please input all fields"})
		return
	}

	user, err := userService.FindByUsername(input.UserName)

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Not found"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.AbortWithStatusJSON(402, gin.H{"error": "Email or password is invalid."})
		return
	}

	token, err := user.GetJwtToken()

	if err != nil {
		c.AbortWithStatusJSON(402, gin.H{"error": "Something went wrong."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK", "token": token, "data": user.ID})
}
