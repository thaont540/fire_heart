package controllers

import (
	"fire_heart/models/entity"
	"fire_heart/models/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type EducationController struct{}
var educationService = new(service.EducationService)

func (educationController *EducationController) Store(c *gin.Context) {
	type CreateEducationInput struct {
		StartDate string `json:"start_date" binding:"required"`
		EndDate string `json:"end_date" binding:"required"`
		School string `json:"school" binding:"required"`
		Major string `json:"major"`
		Description string `json:"description"`
		Certificate string `json:"certificate"`
	}

	var input CreateEducationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Please input all fields"})
		return
	}

	user := c.MustGet("user").(entity.User)

	education := entity.Education{}
	education.UserId = user.ID.Hex()
	education.Description = input.Description
	startDate, _ := time.Parse("2006-01-02", input.StartDate)
	education.StartDate = startDate
	endDate, _ := time.Parse("2006-01-02", input.EndDate)
	education.EndDate = endDate
	education.School = input.School
	education.Major = input.Major
	education.Certificate = input.Certificate

	insertResult, err := educationService.Create(education)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"message": "Can not insert new education"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "data": insertResult.InsertedID})
	}
}

func (educationController *EducationController) Show(c *gin.Context) {
	userId := c.Param("userId")

	educations, err := educationService.FindByUserId(userId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Education not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "data": educations})
	}
}

func (educationController *EducationController) Delete(c *gin.Context) {
	user := c.MustGet("user").(entity.User)

	DeletedCount, err := userService.DeleteById(user.ID.Hex())

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "User not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "data": DeletedCount})
	}
}
