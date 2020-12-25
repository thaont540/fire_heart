package controllers

import (
	"fire_heart/models/entity"
	"fire_heart/models/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type UserController struct{}
var userService = new(service.UserService)

func (userController *UserController) Store(c *gin.Context) {
	type CreateUserInput struct {
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Name string `json:"name" binding:"required"`
	}

	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Please input all fields"})
		return
	}

	user := entity.User{}
	user.Email = input.Email
	user.Name = input.Name
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
		return
	}
	user.Password = string(hash)

	insertResult, err := userService.Create(user)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Can not insert new user"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "data": insertResult.InsertedID})
	}
}

func (userController *UserController) Show(c *gin.Context) {
	type FindUserInput struct {
		Email string `json:"email" binding:"required"`
	}

	var input FindUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Please input all fields"})
		return
	}

	user, err := userService.FindByEmail(input.Email)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "User not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "data": user})
	}
}

func (userController *UserController) Delete(c *gin.Context) {
	type DeleteUserInput struct {
		Email string `json:"email" binding:"required"`
	}

	var input DeleteUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Please input all fields"})
		return
	}

	deleteResult, err := userService.DeleteByEmail(input.Email)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "User not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "data": deleteResult.DeletedCount})
	}
}
