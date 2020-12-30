package routers

import (
	"fire_heart/controllers"
	"fire_heart/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func setAuthRoute(router *gin.Engine) {
	authController := new(controllers.AuthController)
	router.POST("/login", authController.Login)
}

func setUserRoute(router *gin.Engine)  {
	userController := new (controllers.UserController)
	routeGroup := router.Group("/users")
	routeGroup.Use(middlewares.Authentication())
	routeGroup.POST("/", userController.Store)
	routeGroup.GET("/:id", userController.Show)
	routeGroup.DELETE("/:id", userController.Delete)
}

func setProfileRoute(router *gin.Engine) {
	profileController := new(controllers.ProfileController)
	router.POST("/profile/:userId", profileController.Store)
	router.GET("/profile/:userId", profileController.Show)
}

func setExperienceRoute(router *gin.Engine) {
	experienceController := new(controllers.ExperienceController)
	router.POST("/experience/:userId", experienceController.Store)
	router.GET("/experience/:userId", experienceController.Show)
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
	setUserRoute(router)
	setProfileRoute(router)
	setExperienceRoute(router)

	return router
}
