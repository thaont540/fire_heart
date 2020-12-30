package controllers

import (
	"fire_heart/models/entity"
	"fire_heart/models/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ProfileController struct{}
var profileService = new(service.ProfileService)

func (profileController *ProfileController) Store(c *gin.Context) {
	type CreateProfileInput struct {
		Email string `json:"email" binding:"required"`
		Name string `json:"name" binding:"required"`
		Career string `json:"career" binding:"required"`
		Image string `json:"image"`
		Description string `json:"description"`
		Address string `json:"address"`
		Language string `json:"language"`
		Birthday string `json:"birthday"`
	}

	var input CreateProfileInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Please input all fields"})
		return
	}

	profile := entity.Profile{}
	profile.UserId = c.Param("userId")
	profile.Email = input.Email
	profile.Name = input.Name
	profile.Career = input.Career
	profile.Image = input.Image
	profile.Description = input.Description
	profile.Address = input.Address
	profile.Language = input.Language
	t, _ := time.Parse("2006-01-02", input.Birthday)
	profile.Birthday = t

	insertResult, err := profileService.Create(profile)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Can not insert new profile"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "data": insertResult.InsertedID})
	}
}

func (profileController *ProfileController) Show(c *gin.Context) {
	userId := c.Param("id")
	profile, err := profileService.FindByUserId(userId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "User not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "data": profile})
	}
}

//func (profileController *ProfileController) Delete(c *gin.Context) {
//	type DeleteUserInput struct {
//		Email string `json:"email" binding:"required"`
//	}
//
//	var input DeleteUserInput
//
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.AbortWithStatusJSON(401, gin.H{"error": "Please input all fields"})
//		return
//	}
//
//	deleteResult, err := userService.DeleteByEmail(input.Email)
//
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{"message": "User not found"})
//	} else {
//		c.JSON(http.StatusOK, gin.H{"message": "OK", "data": deleteResult.DeletedCount})
//	}
//}
