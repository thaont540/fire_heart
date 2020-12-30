package routers

import (
	"fire_heart/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func setAuthRoute(router *gin.Engine) {
	userController := new (controllers.UserController)
	router.POST("/users", userController.Store)
	router.GET("/users/:id", userController.Show)
	router.DELETE("/users/:id", userController.Delete)
}

func setProfileRoute(router *gin.Engine) {
	profileController := new(controllers.ProfileController)
	router.POST("/users/:userId/profile", profileController.Store)
	router.GET("/users/:id/profile", profileController.Show)
}

func setExperienceRoute(router *gin.Engine) {
	experienceController := new(controllers.ExperienceController)
	router.POST("/users/:userId/experience", experienceController.Store)
	router.GET("/users/:id/experience", experienceController.Show)
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	setAuthRoute(router)
	setProfileRoute(router)
	setExperienceRoute(router)

	return router
}
