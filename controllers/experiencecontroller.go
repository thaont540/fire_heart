package controllers

import (
	"fire_heart/models/entity"
	"fire_heart/models/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ExperienceController struct {}
var experienceService = new(service.ExperienceService)

func (experienceController *ExperienceController)Store(c *gin.Context) {
	type CreateExperienceInput struct {
		StartDate string `json:"start_date" binding:"required"`
		EndDate string `json:"end_date" binding:"required"`
		Position string `json:"position" binding:"required"`
		Project string `json:"project" binding:"required"`
		Description string `json:"description" binding:"required"`
		Technical string `json:"technical" binding:"required"`
		TeamSize string `json:"team_size" binding:"required"`
		Effort string `json:"effort" binding:"required"`
	}

	var input CreateExperienceInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Please input all fields"})
		return
	}

	user := c.MustGet("user").(entity.User)

	experience := entity.Experience{}
	experience.UserId = user.ID.Hex()
	experience.Position = input.Position
	experience.Project = input.Project
	experience.Description = input.Description
	experience.Technical = input.Technical
	experience.TeamSize = input.TeamSize
	experience.Effort = input.Effort
	startDate, _ := time.Parse("2006-01-02", input.StartDate)
	experience.StartDate = startDate
	endDate, _ := time.Parse("2006-01-02", input.EndDate)
	experience.EndDate = endDate

	insertResult, err := experienceService.Create(experience)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Can not insert new experience"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "data": insertResult.InsertedID})
	}
}

func (experienceController *ExperienceController) Show(c *gin.Context) {
	userId := c.Param("userId")
	experience, err := experienceService.FindByUserId(userId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "User not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "data": experience})
	}
}