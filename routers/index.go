package routers

import (
	"fire_heart/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func setAuthRoute(router *gin.Engine)  {
	userController := new (controllers.UserController)
	router.POST("/", userController.Store)
	router.GET("/:email", userController.Show)
	router.DELETE("/:email", userController.Delete)
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

	return router
}
