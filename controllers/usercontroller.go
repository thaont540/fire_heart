package controllers

import (
	"fire_heart/models/entity"
	"fire_heart/models/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type UserController struct{}
var userService = new(service.UserService)

func (userController *UserController) Store(c *gin.Context) {
	type CreateUserInput struct {
		UserName string `idx:"{user_name},unique" json:"user_name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Please input all fields"})
		return
	}

	user := entity.User{}
	user.UserName = input.UserName
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
		return
	}
	user.Password = string(hash)

	insertResult, err := userService.Create(user)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"message": "Can not insert new user"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "data": insertResult.InsertedID})
	}
}

func (userController *UserController) Show(c *gin.Context) {
	id := c.Param("id")

	user, err := userService.FindById(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "User not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "data": user})
	}
}

func (userController *UserController) Delete(c *gin.Context) {
	user := c.MustGet("user").(entity.User)

	DeletedCount, err := userService.DeleteById(user.ID.Hex())

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "User not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "data": DeletedCount})
	}
}
