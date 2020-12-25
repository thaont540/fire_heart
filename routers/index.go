package routers

import (
	"fire_heart/controllers"
	"github.com/gin-gonic/gin"
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

	setAuthRoute(router)

	return router
}
